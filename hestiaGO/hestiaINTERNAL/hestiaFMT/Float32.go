// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
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

// NOTE:
//   1. Test codes are in the public hestiaSTRING package.
//   2. Ryū algorithm paper -> https://doi.org/10.1145/3192366.3192369

package hestiaFMT

import (
	"hestia/hestiaINTERNAL/hestiaMATH"
)

const (
	value_POWER_5_BITS32         = 61
	value_POWER_5_INVERSE_BITS32 = 59
)

type ParamsFLOAT32 struct {
	Value      float32
	Precision  uint32
	Base       uint32
	Lettercase Lettercase
	Notation   Notation
}

func M32_FormatFLOAT32(input *ParamsFLOAT32) (out []rune) {
	var data *hestiaMATH.Float32
	var i uint32

	// start by panicking invalid inputs and picking off easy targets
	switch {
	case input == nil:
		panic("input data structure is nil!")
	case input.Base < 2 || input.Base > 36:
		panic("base number smaller than 2 or larger than 36!")
	case input.Notation == NOTATION_ISO6093NR3_AUTO &&
		!(input.Base == 2 || input.Base == 10 || input.Base == 16):
		fallthrough
	case input.Notation == NOTATION_ISO6093NR3 &&
		!(input.Base == 2 || input.Base == 10 || input.Base == 16):
		panic("requested ISO6093 NR3 notation but base is not 2/10/16!")
	case input.Notation == NOTATION_IEEE754 && input.Base != 2:
		panic("requested IEEE754 notation but base is not 2!")
	case input.Notation == NOTATION_IEEE754:
		i = hestiaMATH.S32_IEEE754_FloatToBits(input.Value)
		out = FormatUINT32(i, 2, LETTERCASE_LOWER)
		for i = 32 - uint32(len(out)); i > 0; i-- {
			out = append([]rune{'0'}, out...)
		}

		return out
	case hestiaMATH.S32_IEEE754_IsNaN_FLOAT32(input.Value):
		return []rune{'N', 'a', 'N'}
	case hestiaMATH.S32_IEEE754_IsINF_FLOAT32(input.Value, 1):
		return []rune{'+', '∞'}
	case hestiaMATH.S32_IEEE754_IsINF_FLOAT32(input.Value, -1):
		return []rune{'-', '∞'}
	case input.Value == 0:
		out = []rune{'0', '.', '0'}
		if input.Precision > 0 {
			for i = input.Precision - 1; i != 0; i-- {
				out = append(out, '0')
			}
		}

		return out
	default:
		data = &hestiaMATH.Float32{}
		hestiaMATH.S32_IEEE754_Decode_FLOAT32(input.Value, data)
	}

	// convert to base10 using Ryū algorithm with its lookup tables
	if !__s32_Ryū_Base10_Exact_INT_Float32(data) {
		__m32_Ryū_Base10_Float32(data)
	}

	// encode to string using Hadōken algorithm
	return __s32_Hadōken_Base10_Encode_Float32(input, data)
}

func __s32_Hadōken_Base10_Encode_Float32(setting *ParamsFLOAT32,
	data *hestiaMATH.Float32) (out []rune) {
	var length, exponentOUT uint32
	var round, partial []rune

	// convert mantissa into char and exponent output base count
	round = FormatUINT32(data.Mantissa, setting.Base, setting.Lettercase)
	exponentOUT = data.Exponent
	if data.Base != 10 {
		exponentOUT = uint32(float64(data.Exponent) /
			hestiaMATH.S64_LOG10(float64(data.Base)))
	}

	// normalize the mantissa value
	__s32_Hadōken_Normalize(&round, &partial, &exponentOUT, &data.NegativeExponent, setting)

	// adjust to output precision setting
	length = uint32(len(partial))
	if setting.Precision > 0 && length < setting.Precision {
		length = setting.Precision - length
		for length != 0 {
			partial = append(partial, '0')
			length--
		}
	}

	// construct output char list
	__s32_Hadōken_Render(&out, &round, &partial, &exponentOUT, setting, data)

	// return the formatted output
	return out
}

func __s32_Hadōken_Render(buffer *[]rune, round *[]rune, partial *[]rune,
	exponent *uint32, setting *ParamsFLOAT32, data *hestiaMATH.Float32) {
	var ret []rune

	if data.NegativeValue {
		*buffer = append(*buffer, '-')
	}

	switch {
	case setting.Notation == NOTATION_ISO6093NR3,
		setting.Notation == NOTATION_ISO6093NR3_AUTO:
		if setting.Base == 16 {
			if setting.Lettercase == LETTERCASE_UPPER {
				*buffer = append(*buffer, '0', 'X')
			} else {
				*buffer = append(*buffer, '0', 'x')
			}
		}
	}

	*buffer = append(*buffer, *round...)
	*buffer = append(*buffer, '.')
	*buffer = append(*buffer, *partial...)

	if *exponent > 0 {
		switch {
		case setting.Notation == NOTATION_ISO6093NR3,
			setting.Notation == NOTATION_ISO6093NR3_AUTO:
			switch {
			case setting.Base == 2, setting.Base == 16:
				if setting.Lettercase == LETTERCASE_UPPER {
					*buffer = append(*buffer, 'P')
				} else {
					*buffer = append(*buffer, 'p')
				}
			case setting.Base == 10:
				if setting.Lettercase == LETTERCASE_UPPER {
					*buffer = append(*buffer, 'E')
				} else {
					*buffer = append(*buffer, 'e')
				}
			}
		default:
			*buffer = append(*buffer, '*')
			ret = FormatUINT32(setting.Base, 10, setting.Lettercase)
			*buffer = append(*buffer, ret...)
		}

		if data.NegativeExponent {
			*buffer = append(*buffer, '-')
		} else {
			*buffer = append(*buffer, '+')
		}

		ret = FormatUINT32(*exponent, 10, setting.Lettercase)
		*buffer = append(*buffer, ret...)
	}
}

func __s32_Hadōken_Normalize(round, partial *[]rune,
	exponentOUT *uint32, negativeExponent *bool, setting *ParamsFLOAT32) {
	rLength := uint32(len(*round))
	switch {
	case !*negativeExponent && *exponentOUT >= rLength: // huge positive number
		switch {
		case setting.Notation == NOTATION_DECIMAL_NO_EXPONENT:
			for *exponentOUT != 0 {
				*round = append(*round, '0')
				*exponentOUT--
			}
		default:
			// always scientific 1 decimal notation (X.XXXX^B+YYYY)
			*exponentOUT += rLength - 1
			*partial = (*round)[1:]
			*round = (*round)[0:1]
		}
	case !*negativeExponent && *exponentOUT < rLength: // no exponent positive number
		switch {
		case setting.Notation == NOTATION_SCIENTIFIC,
			setting.Notation == NOTATION_ISO6093NR3:
			// scientific 1 decimal notation (X.XXXX^B+YYYY)
			*exponentOUT += rLength - 1
			*partial = (*round)[1:]
			*round = (*round)[0:1]
		default:
			// always decimal with no exponent notation (XXXXXXXX.0)
			for *exponentOUT != 0 {
				*round = append(*round, '0')
				*exponentOUT--
			}
		}
	case *negativeExponent && *exponentOUT <= rLength: // no exponent negative number
		switch {
		case setting.Notation == NOTATION_SCIENTIFIC,
			setting.Notation == NOTATION_ISO6093NR3:
			// scientific 1 decimal notation (X.XXXX^B+YYYY)
			rLength--
			*partial = (*round)[1:]
			*round = (*round)[0:1]

			*negativeExponent = false
			*exponentOUT = rLength - *exponentOUT
		default:
			// always decimal with no exponent notation (XXXX.XXXX)
			*partial = (*round)[*exponentOUT:]
			*round = (*round)[0:*exponentOUT]
			*exponentOUT = 0
		}
	case *negativeExponent && *exponentOUT > rLength: // very small number
		switch {
		case setting.Notation == NOTATION_DECIMAL_NO_EXPONENT:
			*exponentOUT -= rLength
			for *exponentOUT != 0 {
				*partial = append(*partial, '0')
				*exponentOUT--
			}
			*partial = append(*partial, *round...)
			*round = []rune{'0'}
		default:
			// always scientific 1 decimal notation (X.XXXX^B-YYYY)
			*partial = (*round)[1:]
			*round = (*round)[0:1]
			*exponentOUT -= rLength - 1
		}
	}

	if len(*partial) == 0 {
		*partial = []rune{'0'}
	}
}

func __s32_Ryū_Base10_Exact_INT_Float32(data *hestiaMATH.Float32) bool {
	var e, shift int32
	var x, y uint32

	// check if exponent is smaller than mantissa where roundiness is performable
	e = int32(data.Exponent) - hestiaMATH.BIAS_FLOAT32_EXPONENT
	if e > hestiaMATH.MASK_FLOAT32_MANTISSA_MSB {
		return false // base-10 value is too big - bail out
	}

	// check for mantissa integrity
	shift = hestiaMATH.MASK_FLOAT32_MANTISSA_MSB - e
	x = data.Mantissa | 1<<hestiaMATH.MASK_FLOAT32_MANTISSA_MSB // implicit '1' (1.000)
	y = x >> shift
	if (y << shift) != x {
		return false // data loss detected - bail out
	}
	data.Mantissa = y

	// good case - begin calculation in base10
	e = 0
	for data.Mantissa%10 == 0 {
		data.Mantissa /= 10
		e++
	}

	if e < 0 {
		data.NegativeExponent = true
		data.Exponent = uint32(e * -1)
	} else {
		data.NegativeExponent = false
		data.Exponent = uint32(e)
	}

	data.Base = 10

	return true
}

//nolint:gocognit
func __m32_Ryū_Base10_Float32(data *hestiaMATH.Float32) {
	var e2, e10, i, j, k, l, u, v, w, removed int32
	var m2, q, a, b, c, out uint32
	var lastRemovedDigit uint8
	var even, uShift, acceptBounds, aTrailingZeros, bTrailingZeros bool

	// Step 1 - parse and preparations
	if data.Exponent == 0 {
		e2 = 1 -
			hestiaMATH.BIAS_FLOAT32_EXPONENT -
			hestiaMATH.MASK_FLOAT32_MANTISSA_MSB -
			2
		m2 = data.Mantissa
	} else {
		e2 = int32(data.Exponent) -
			hestiaMATH.BIAS_FLOAT32_EXPONENT -
			hestiaMATH.MASK_FLOAT32_MANTISSA_MSB -
			2
		m2 = uint32(1)<<hestiaMATH.MASK_FLOAT32_MANTISSA_MSB | data.Mantissa
	}
	even = m2&1 == 0
	acceptBounds = even

	// Step 2- determine valid decimal representations interval
	u = 1
	if data.Mantissa != 0 || data.Exponent <= 1 {
		u = 2
		uShift = true
	}

	u = 4*int32(m2) - u // mm
	v = 4 * int32(m2)   // mv
	w = 4*int32(m2) + 2 // mp

	// Step 3 - Convert u, v, w to decimal base
	if e2 >= 0 {
		q = __s32_LOG10_2PowerOf(uint32(e2))
		e10 = int32(q)
		k = value_POWER_5_INVERSE_BITS32 + int32(__s32_LOG2_5PowerOf(q)) - 1
		i = -e2 + int32(q) + k
		a = __m32_Multiply_5PowerInverse_Divide_By_2Power(uint32(u), q, i) // vm
		b = __m32_Multiply_5PowerInverse_Divide_By_2Power(uint32(v), q, i) // vr
		c = __m32_Multiply_5PowerInverse_Divide_By_2Power(uint32(w), q, i) // vp

		if q != 0 && (c-1)/10 <= a/10 {
			// learn at least 1 removed digit
			l = value_POWER_5_INVERSE_BITS32 + int32(__s32_LOG2_5PowerOf(q-1)) - 1
			lastRemovedDigit = uint8(
				__m32_Multiply_5PowerInverse_Divide_By_2Power(
					uint32(v),
					q-1,
					-e2+int32(q-1)+l,
				) % 10,
			)
		}

		if q <= 9 {
			// largest 5 power = 5^10 that fits in 24-bits number.
			// Only 1 of u, v, and w can be a multiple 5, if any.
			switch {
			case v%5 == 0:
				bTrailingZeros = __s32_Factor_Of_5(uint32(v), q)
			case acceptBounds:
				aTrailingZeros = __s32_Factor_Of_5(uint32(u), q)
			case __s32_Factor_Of_5(uint32(w), q):
				c--
			}
		}
	} else {
		q = __s32_LOG10_5PowerOf(uint32(-1 * e2)) // modulus the negative e2
		e10 = e2 + int32(q)
		i = -e2 - int32(q)
		k = int32(__s32_LOG2_5PowerOf(uint32(i))) - value_POWER_5_BITS32
		j = int32(q) - k

		a = __m32_Multiply_5Power_Divide_By_2Power(uint32(u), uint32(i), j) // vm
		b = __m32_Multiply_5Power_Divide_By_2Power(uint32(v), uint32(i), j) // vr
		c = __m32_Multiply_5Power_Divide_By_2Power(uint32(w), uint32(i), j) // vp

		if q != 0 && (c-1)/10 <= a/10 {
			j = int32(q) -
				1 -
				(int32(__s32_LOG2_5PowerOf(uint32(i))) - value_POWER_5_BITS32)
			lastRemovedDigit = uint8(
				__m32_Multiply_5Power_Divide_By_2Power(
					uint32(v),
					uint32(i+1),
					j,
				) % 10,
			)
		}

		if q <= 1 {
			// {a,b,c} is trailing zeros if:
			//   1) {u,v,w} has at least q trailing 0-bits

			// for 'v = 4 * m2', v always has at least 2 trailing 0-bits
			bTrailingZeros = true

			if acceptBounds {
				// for 'u = m2 - 2', it has 1 trailing 0-bit;
				// for 'u = m2 - 1', it has 0 trailing 0-bit;
				// Therefore, u always have 1 trailing 0-bit if shifted.
				aTrailingZeros = uShift
			} else {
				// for 'w = m2 + 2', it always has at least 1 trailing 0-bit
				c--
			}
		} else if q < 31 {
			bTrailingZeros = __s32_Factor_Of_2(uint32(v), q-1)
		}
	}

	// Step 4 - Find shortest, correctly rounded decimal representaton
	if aTrailingZeros || bTrailingZeros {
		// General cases (probability: ~4.0%)
		for c/10 > a/10 {
			aTrailingZeros = aTrailingZeros && a%10 == 0
			bTrailingZeros = bTrailingZeros && lastRemovedDigit == 0
			lastRemovedDigit = uint8(b % 10)
			a /= 10
			b /= 10
			c /= 10
			removed++
		}

		if aTrailingZeros {
			for a%10 == 0 {
				bTrailingZeros = bTrailingZeros && lastRemovedDigit == 0
				lastRemovedDigit = uint8(b % 10)
				a /= 10
				b /= 10
				c /= 10
				removed++
			}
		}

		if bTrailingZeros && lastRemovedDigit == 5 && b%2 == 0 {
			// Round even for exact number (....50..0.)
			lastRemovedDigit = 4
		}

		out = b
		if (b == a && (!acceptBounds || !aTrailingZeros)) || lastRemovedDigit >= 5 {
			// b is outside bound OR round up is required -> out + 1
			out++
		}
	} else {
		// common case (~96.0%) - Relative percentage per loop iteration
		// approximately:
		//   0: 13.6%
		//   1: 70.7%
		//   2: 14.1%
		//   3: 1.39%
		//   4: 0.14%
		//   5+: 0.01%
		for c/10 > a/10 {
			lastRemovedDigit = uint8(b % 10)
			a /= 10
			b /= 10
			c /= 10
			removed++
		}

		out = b
		if b == a || lastRemovedDigit >= 5 {
			// b is outside bound OR round up is required -> out += 1
			out++
		}
	}

	// Step 5 - Update data structure with processed value
	data.Base = 10
	data.Mantissa = out
	e10 += removed
	data.NegativeExponent = false
	if e10 < 0 {
		e10 *= -1
		data.NegativeExponent = true
	}
	data.Exponent = uint32(e10)
}

func __m32_Multiply_5PowerInverse_Divide_By_2Power(m, q uint32, j int32) uint32 {
	var multiplier uint64

	switch q {
	case 0:
		multiplier = 576460752303423489
	case 1:
		multiplier = 461168601842738791
	case 2:
		multiplier = 368934881474191033
	case 3:
		multiplier = 295147905179352826
	case 4:
		multiplier = 472236648286964522
	case 5:
		multiplier = 377789318629571618
	case 6:
		multiplier = 302231454903657294
	case 7:
		multiplier = 483570327845851670
	case 8:
		multiplier = 386856262276681336
	case 9:
		multiplier = 309485009821345069
	case 10:
		multiplier = 495176015714152110
	case 11:
		multiplier = 396140812571321688
	case 12:
		multiplier = 316912650057057351
	case 13:
		multiplier = 507060240091291761
	case 14:
		multiplier = 405648192073033409
	case 15:
		multiplier = 324518553658426727
	case 16:
		multiplier = 519229685853482763
	case 17:
		multiplier = 415383748682786211
	case 18:
		multiplier = 332306998946228969
	case 19:
		multiplier = 531691198313966350
	case 20:
		multiplier = 425352958651173080
	case 21:
		multiplier = 340282366920938464
	case 22:
		multiplier = 544451787073501542
	case 23:
		multiplier = 435561429658801234
	case 24:
		multiplier = 348449143727040987
	case 25:
		multiplier = 557518629963265579
	case 26:
		multiplier = 446014903970612463
	case 27:
		multiplier = 356811923176489971
	case 28:
		multiplier = 570899077082383953
	case 29:
		multiplier = 456719261665907162
	case 30:
		multiplier = 365375409332725730
	}

	return __s32_Multiply_By_Shift_32_Bits(m, multiplier, j)
}

func __m32_Multiply_5Power_Divide_By_2Power(m, q uint32, j int32) uint32 {
	var multiplier uint64

	switch q {
	case 0:
		multiplier = 1152921504606846976
	case 1:
		multiplier = 1441151880758558720
	case 2:
		multiplier = 1801439850948198400
	case 3:
		multiplier = 2251799813685248000
	case 4:
		multiplier = 1407374883553280000
	case 5:
		multiplier = 1759218604441600000
	case 6:
		multiplier = 2199023255552000000
	case 7:
		multiplier = 1374389534720000000
	case 8:
		multiplier = 1717986918400000000
	case 9:
		multiplier = 2147483648000000000
	case 10:
		multiplier = 1342177280000000000
	case 11:
		multiplier = 1677721600000000000
	case 12:
		multiplier = 2097152000000000000
	case 13:
		multiplier = 1310720000000000000
	case 14:
		multiplier = 1638400000000000000
	case 15:
		multiplier = 2048000000000000000
	case 16:
		multiplier = 1280000000000000000
	case 17:
		multiplier = 1600000000000000000
	case 18:
		multiplier = 2000000000000000000
	case 19:
		multiplier = 1250000000000000000
	case 20:
		multiplier = 1562500000000000000
	case 21:
		multiplier = 1953125000000000000
	case 22:
		multiplier = 1220703125000000000
	case 23:
		multiplier = 1525878906250000000
	case 24:
		multiplier = 1907348632812500000
	case 25:
		multiplier = 1192092895507812500
	case 26:
		multiplier = 1490116119384765625
	case 27:
		multiplier = 1862645149230957031
	case 28:
		multiplier = 1164153218269348144
	case 29:
		multiplier = 1455191522836685180
	case 30:
		multiplier = 1818989403545856475
	case 31:
		multiplier = 2273736754432320594
	case 32:
		multiplier = 1421085471520200371
	case 33:
		multiplier = 1776356839400250464
	case 34:
		multiplier = 2220446049250313080
	case 35:
		multiplier = 1387778780781445675
	case 36:
		multiplier = 1734723475976807094
	case 37:
		multiplier = 2168404344971008868
	case 38:
		multiplier = 1355252715606880542
	case 39:
		multiplier = 1694065894508600678
	case 40:
		multiplier = 2117582368135750847
	case 41:
		multiplier = 1323488980084844279
	case 42:
		multiplier = 1654361225106055349
	case 43:
		multiplier = 2067951531382569187
	case 44:
		multiplier = 1292469707114105741
	case 45:
		multiplier = 1615587133892632177
	case 46:
		multiplier = 2019483917365790221
	}
	return __s32_Multiply_By_Shift_32_Bits(m, multiplier, j)
}

func __s32_Multiply_By_Shift_32_Bits(m uint32, multiplier uint64, shift int32) uint32 {
	if shift <= 32 {
		panic("Format_FLOAT32: shift value must be >32")
	}

	high, low := hestiaMATH.S64_Multiply_By_Bits(uint64(m), multiplier)
	sum := (low >> uint(shift)) + (high << uint(64-shift))
	if sum >= hestiaMATH.MAX_UINT32 {
		panic("Format_FLOAT32: shifted sum overflowed (>=MAX_UINT32)")
	}

	return uint32(sum)
}

func __s32_Factor_Of_2(input uint32, minimum uint32) bool {
	return hestiaMATH.S32_TrailingZeros_32Bits(input) >= minimum
}

func __s32_Factor_Of_5(input uint32, power uint32) bool {
	nTimes, _, _ := hestiaMATH.S32_FactorOf(int32(input), 5)
	return nTimes >= power
}

func __s32_LOG2_5PowerOf(input uint32) uint32 {
	if input > 3528 {
		panic("Format_FLOAT32: Log2_5PowerOf has big input (>3528).")
	}

	return (input*1217359)>>19 + 1
}

func __s32_LOG10_5PowerOf(input uint32) uint32 {
	if input > 2620 {
		panic("Format_FLOAT32: Log10_5PowerOf has big input (>2620).")
	}

	return (input * 732923) >> 20
}

func __s32_LOG10_2PowerOf(input uint32) uint32 {
	if input > 1650 {
		panic("Format_FLOAT32: Log10_2PowerOf has big input (>1650).")
	}

	return (input * 78913) >> 18
}
