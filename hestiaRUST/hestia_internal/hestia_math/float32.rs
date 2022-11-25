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

// IEEE754 32-bits float bit-definition: (1) sign (8) exponent (23) mantissa

pub const BITS_FLOAT32_MANTISSA: u32 = 23;
pub const BITS_FLOAT32_EXPONENT: u32 = 8;
pub const MASK_FLOAT32_SIGN: u32 = 1 << (BITS_FLOAT32_EXPONENT + BITS_FLOAT32_MANTISSA);
pub const MASK_FLOAT32_MANTISSA: u32 = 1 << BITS_FLOAT32_MANTISSA - 1;
pub const MASK_FLOAT32_EXPONENT: u32 = (1 << BITS_FLOAT32_EXPONENT - 1) << BITS_FLOAT32_MANTISSA;
pub const BIAS_FLOAT32_EXPONENT: u32 = 127;
pub const MASK_FLOAT32_INFINITY_POSITIVE: u32 = MASK_FLOAT32_EXPONENT;
pub const MASK_FLOAT32_INFINITY_NEGATIVE: u32 = MASK_FLOAT32_SIGN | MASK_FLOAT32_EXPONENT;

pub struct Float32 {
	pub exponent: u32,
	pub mantissa: u32,
	pub base: u32,
	pub negative_value: bool,
	pub negative_exponent: bool,
}

pub fn s32_ieee754_decode_float32(input: f32, data: &mut Float32) {
	data.base = 2;
	data.mantissa = s32_ieee754_float_to_bits(input);
	data.negative_value = (data.mantissa & MASK_FLOAT32_SIGN) != 0;
	data.exponent = (data.mantissa & MASK_FLOAT32_EXPONENT) >> BITS_FLOAT32_MANTISSA;
	data.mantissa &= MASK_FLOAT32_MANTISSA;
}

pub fn s32_ieee754_is_nan_float32(input: f32) -> bool {
	let x = s32_ieee754_float_to_bits(input);
	return x & MASK_FLOAT32_MANTISSA != 0
		&& x & MASK_FLOAT32_EXPONENT == MASK_FLOAT32_EXPONENT;
}

pub fn s32_ieee754_is_inf_float32(input: f32, sign: i8) -> bool {
	let x = s32_ieee754_float_to_bits(input);
	return (sign >= 0 && x == MASK_FLOAT32_INFINITY_POSITIVE)
		|| (sign <= 0 && x == MASK_FLOAT32_INFINITY_NEGATIVE);
}

pub fn s32_ieee754_float_to_bits(input: f32) -> u32 {
	// note:
	//   1) Rust core team discovered some stability concerns that is hard
	//      to comprehend/test against at this time. Their discovery shall
	//      be revisited in near future as optimization efforts and
	//      direction alignment.
	unsafe {
		// Transmuting data type on memory is always unsafe in both Rust
		// and Go. We just have to make sure it is clearly specified and
		// only do what it supposedly be doing.
		return std::mem::transmute::<f32, u32>(input);
	}
}

pub fn s32_ieee754_bits_to_float(input: u32) -> f32 {
	// note:
	//   1) Rust core team discovered some stability concerns that is hard
	//      to comprehend/test against at this time. Their discovery shall
	//      be revisited in near future as optimization efforts and
	//      direction alignment.
	unsafe {
		// Transmuting data type on memory is always unsafe in both Rust
		// and Go. We just have to make sure it is clearly specified and
		// only do what it supposedly be doing.
		return std::mem::transmute::<u32, f32>(input);
	}
}

pub fn s32_floor_float32(input: f32) -> f32 {
	let fraction: f32;
	let mut round: f32;

	if input == 0.0 {
		return input;
	} else if s32_ieee754_is_nan_float32(input) {
		return input;
	} else if s32_ieee754_is_inf_float32(input, 0) {
		return input;
	} else if input < 0.0 {
		(round, fraction) = s32_modf_float32(-input);
		if fraction != 0.0 {
			round += 1.0;
		}

		return -round;
	} else {
		(round, _) = s32_modf_float32(input);
	}

	return round;
}

pub fn s32_modf_float32(input: f32) -> (f32, f32) {
	let fraction: f32;
	let round: f32;

	if input < 1.0 {
		if input < 0.0 {
			(round, fraction) = s32_modf_float32(-input);
			return (-round, -fraction);
		} else if input == 0.0 {
			return (input, input);
		} else {
			return (0.0, input); // 0, .0
		}
	}

	let mut x = s32_ieee754_float_to_bits(input);
	let e = ((x & MASK_FLOAT32_EXPONENT) >> BITS_FLOAT32_MANTISSA) - BIAS_FLOAT32_EXPONENT;

	if e < 32 - 8 {
		x = x & !(1 << (32 - 8 - e) - 1);
	}

	round = s32_ieee754_bits_to_float(x);
	fraction = input - round;

	return (round, fraction);
}
