// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 Caleb Spare <cespare@gmail.com>
// Copyright 2018 Ulf Adams <ulfjack@google.com>
// Copyright 2009 The Go Authors <https://cs.opensource.google/go/go>
//
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//                  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

// IEEE754 64-bits float bit-definition: (1) sign (11) exponent (52) mantissa

pub const BITS_FLOAT64_MANTISSA: u64 = 52;
pub const BITS_FLOAT64_EXPONENT: u64 = 11;
pub const MASK_FLOAT64_SIGN: u64 = 1 << (BITS_FLOAT64_EXPONENT + BITS_FLOAT64_MANTISSA);
pub const MASK_FLOAT64_MANTISSA: u64 = 1 << BITS_FLOAT64_MANTISSA - 1;
pub const MASK_FLOAT64_EXPONENT: u64 = (1 << BITS_FLOAT64_EXPONENT - 1) << BITS_FLOAT64_MANTISSA;
pub const BIAS_FLOAT64_EXPONENT: u64 = 1023;
pub const MASK_FLOAT64_INFINITY_POSITIVE: u64 = MASK_FLOAT64_EXPONENT;
pub const MASK_FLOAT64_INFINITY_NEGATIVE: u64 = MASK_FLOAT64_SIGN | MASK_FLOAT64_EXPONENT;

pub struct Float64 {
	pub exponent: u64,
	pub mantissa: u64,
	pub base: u64,
	pub negative_value: bool,
	pub negative_exponent: bool,
}

pub fn s64_ieee754_decode_float64(input: f64, data: &mut Float64) {
	data.base = 2;
	data.mantissa = s64_ieee754_float_to_bits(input);
	data.negative_value = (data.mantissa & MASK_FLOAT64_SIGN) != 0;
	data.exponent = (data.mantissa & MASK_FLOAT64_EXPONENT) >> BITS_FLOAT64_MANTISSA;
	data.mantissa &= MASK_FLOAT64_MANTISSA;
}

pub fn s64_ieee754_is_nan_float64(input: f64) -> bool {
	let x = s64_ieee754_float_to_bits(input);
	return x & MASK_FLOAT64_MANTISSA != 0
		&& x & MASK_FLOAT64_EXPONENT == MASK_FLOAT64_EXPONENT;
}

pub fn s64_ieee754_is_inf_float64(input: f64, sign: i8) -> bool {
	let x = s64_ieee754_float_to_bits(input);
	return (sign >= 0 && x == MASK_FLOAT64_INFINITY_POSITIVE)
		|| (sign <= 0 && x == MASK_FLOAT64_INFINITY_NEGATIVE);
}

pub fn s64_ieee754_float_to_bits(input: f64) -> u64 {
	// note:
	//   1) Rust core team discovered some stability concerns that is hard
	//      to comprehend/test against at this time. Their discovery shall
	//      be revisited in near future as optimization efforts and
	//      direction alignment.
	unsafe {
		// Transmuting data type on memory is always unsafe in both Rust
		// and Go. We just have to make sure it is clearly specified and
		// only do what it supposedly be doing.
		return std::mem::transmute::<f64, u64>(input);
	}
}

pub fn s64_ieee754_bits_to_float(input: u64) -> f64 {
	// note:
	//   1) Rust core team discovered some stability concerns that is hard
	//      to comprehend/test against at this time. Their discovery shall
	//      be revisited in near future as optimization efforts and
	//      direction alignment.
	unsafe {
		// Transmuting data type on memory is always unsafe in both Rust
		// and Go. We just have to make sure it is clearly specified and
		// only do what it supposedly be doing.
		return std::mem::transmute::<u64, f64>(input);
	}
}

pub fn s64_floor_float64(input: f64) -> f64 {
	let fraction: f64;
	let mut round: f64;

	if input == 0.0 {
		return input;
	} else if s64_ieee754_is_nan_float64(input) {
		return input;
	} else if s64_ieee754_is_inf_float64(input, 0) {
		return input;
	} else if input < 0.0 {
		(round, fraction) = s64_modf_float64(-input);
		if fraction != 0.0 {
			round += 1.0;
		}

		return -round;
	} else {
		(round, _) = s64_modf_float64(input);
	}

	return round;
}

pub fn s64_frexp_float64(input: f64) -> (f64, i64) {
	if input == 0.0 {
		return (input, 0);
	} else if s64_ieee754_is_nan_float64(input) {
		return (input, 0);
	} else if s64_ieee754_is_inf_float64(input, 0) {
		return (input, 0);
	} else {
		// proceed to process fractional expression
	}

	let (float, mut exponent) = s64_normalize_float64(input);
	let mut x = s64_ieee754_float_to_bits(float);
	exponent += ((x & MASK_FLOAT64_EXPONENT) >> BITS_FLOAT64_MANTISSA) as i64
		- BIAS_FLOAT64_EXPONENT as i64
		+ 1;
	x &= !MASK_FLOAT64_EXPONENT;
	x |= (BIAS_FLOAT64_EXPONENT - 1) << BITS_FLOAT64_MANTISSA;
	let fraction = s64_ieee754_bits_to_float(x);

	return (fraction, exponent);
}

pub fn s64_modf_float64(input: f64) -> (f64, f64) {
	let fraction: f64;
	let round: f64;

	if input < 1.0 {
		if input < 0.0 {
			(round, fraction) = s64_modf_float64(-input);
			return (-round, -fraction);
		} else if input == 0.0 {
			return (input, input);
		} else {
			return (0.0, input); // 0, .0
		}
	}

	let mut x = s64_ieee754_float_to_bits(input);
	let e = ((x & MASK_FLOAT64_EXPONENT) >> BITS_FLOAT64_MANTISSA) - BIAS_FLOAT64_EXPONENT;

	if e < 64 - 12 {
		x &= !(1 << (64 - 12 - e) - 1);
	}

	round = s64_ieee754_bits_to_float(x);
	fraction = input - round;

	return (round, fraction);
}

pub fn s64_normalize_float64(input: f64) -> (f64, i64) {
	if s64_absolute_float64(input) < 2.2250738585072014e-308 {
		return (input * (1 << 52) as f64, -52);
	}

	return (input, 0);
}

pub fn s64_absolute_float64(input: f64) -> f64 {
	return s64_ieee754_bits_to_float(s64_ieee754_float_to_bits(input) & !MASK_FLOAT64_SIGN);
}
