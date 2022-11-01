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

	// start by panicking invalid inputs
	if input == nil {
		panic("input data structure for FormatFLOAT32 is nil!")
	}
	__SN_Valid_Base_Number(uint8(input.Base))
	__SN_Valid_Notation(uint8(input.Base), input.Notation)

	// picking off easy targets
	switch {
	case input.Notation == NOTATION_IEEE754:
		i = hestiaMATH.S32_IEEE754_FloatToBits(input.Value)
		out = S32_FormatUINT32(i, 2, LETTERCASE_LOWER)
		for i = 32 - uint32(len(out)); i > 0; i-- {
			out = append([]rune{'0'}, out...)
		}

		return out
	case hestiaMATH.S32_IEEE754_IsNaN_FLOAT32(input.Value):
		return __sN_Omizu_Render_NaN_FLOAT()
	case hestiaMATH.S32_IEEE754_IsINF_FLOAT32(input.Value, 1):
		return __sN_Omizu_Render_INF_FLOAT(true)
	case hestiaMATH.S32_IEEE754_IsINF_FLOAT32(input.Value, -1):
		return __sN_Omizu_Render_INF_FLOAT(false)
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
	return __s32_Omizu_Base10_Encode_Float32(input, data)
}

func __s32_Omizu_Base10_Encode_Float32(setting *ParamsFLOAT32,
	data *hestiaMATH.Float32) (out []rune) {
	var param *omizuData
	var length, exponentOUT uint32

	// initialize output rendering data structure
	param = &omizuData{
		round:            []rune{},
		partial:          []rune{},
		exponent:         []rune{},
		base:             []rune{},
		notation:         setting.Notation,
		lettercase:       setting.Lettercase,
		negativeValue:    data.NegativeValue,
		negativeExponent: data.NegativeExponent,
	}

	// convert mantissa into char and exponent output base count
	param.round = S32_FormatUINT32(data.Mantissa, setting.Base, param.lettercase)
	exponentOUT = data.Exponent
	if data.Base != 10 {
		exponentOUT = uint32(float64(data.Exponent) /
			hestiaMATH.S64_LOG10(float64(data.Base)))
	}

	// normalize the mantissa value
	__s32_Omizu_Normalize(param, &exponentOUT)

	// adjust to output precision setting
	length = uint32(len(param.partial))
	if setting.Precision > 0 && length < setting.Precision {
		length = setting.Precision - length
		for length != 0 {
			param.partial = append(param.partial, '0')
			length--
		}
	}

	// convert data into []rune output
	if exponentOUT != 0 {
		param.exponent = S32_FormatUINT32(exponentOUT, 10, param.lettercase)
	}
	param.base = S32_FormatUINT32(setting.Base, 10, param.lettercase)

	// construct output char list
	return __sN_Omizu_Render_FLOAT(param)
}

func __s32_Omizu_Normalize(data *omizuData, exponentOUT *uint32) {
	rLength := uint32(len(data.round))
	switch {
	case !data.negativeExponent && *exponentOUT >= rLength: // huge positive number
		switch {
		case data.notation == NOTATION_DECIMAL_NO_EXPONENT:
			for *exponentOUT != 0 {
				data.round = append(data.round, '0')
				*exponentOUT--
			}
		default:
			// always scientific 1 decimal notation (X.XXXX^B+YYYY)
			*exponentOUT += rLength - 1
			data.partial = (data.round)[1:]
			data.round = (data.round)[0:1]
		}
	case !data.negativeExponent && *exponentOUT < rLength: // no exponent positive number
		switch {
		case data.notation == NOTATION_SCIENTIFIC,
			data.notation == NOTATION_ISO6093NR3:
			// scientific 1 decimal notation (X.XXXX^B+YYYY)
			*exponentOUT += rLength - 1
			data.partial = (data.round)[1:]
			data.round = (data.round)[0:1]
		default:
			// always decimal with no exponent notation (XXXXXXXX.0)
			for *exponentOUT != 0 {
				data.round = append(data.round, '0')
				*exponentOUT--
			}
		}
	case data.negativeExponent && *exponentOUT <= rLength: // no exponent negative number
		switch {
		case data.notation == NOTATION_SCIENTIFIC,
			data.notation == NOTATION_ISO6093NR3:
			// scientific 1 decimal notation (X.XXXX^B+YYYY)
			rLength--
			data.partial = (data.round)[1:]
			data.round = (data.round)[0:1]

			data.negativeExponent = false
			*exponentOUT = rLength - *exponentOUT
		default:
			// always decimal with no exponent notation (XXXX.XXXX)
			data.partial = (data.round)[*exponentOUT:]
			data.round = (data.round)[0:*exponentOUT]
			*exponentOUT = 0
		}
	case data.negativeExponent && *exponentOUT > rLength: // very small number
		switch {
		case data.notation == NOTATION_DECIMAL_NO_EXPONENT:
			*exponentOUT -= rLength
			for *exponentOUT != 0 {
				data.partial = append(data.partial, '0')
				*exponentOUT--
			}
			data.partial = append(data.partial, data.round...)
			data.round = []rune{'0'}
		default:
			// always scientific 1 decimal notation (X.XXXX^B-YYYY)
			data.partial = (data.round)[1:]
			data.round = (data.round)[0:1]
			*exponentOUT -= rLength - 1
		}
	}

	if len(data.partial) == 0 {
		data.partial = []rune{'0'}
	}
}

func __s32_Ryū_Base10_Exact_INT_Float32(data *hestiaMATH.Float32) bool {
	var e, shift int32
	var x uint32

	// check if exponent is smaller than mantissa where roundiness is performable
	e = int32(data.Exponent) - hestiaMATH.BIAS_FLOAT32_EXPONENT
	if e > hestiaMATH.BITS_FLOAT32_MANTISSA {
		return false // base-10 value is too big - bail out
	}

	// check for mantissa integrity
	shift = hestiaMATH.BITS_FLOAT32_MANTISSA - e
	x = data.Mantissa | 1<<hestiaMATH.BITS_FLOAT32_MANTISSA // implicit '1' (1.000)
	if ((x >> shift) << shift) != x {
		return false // data loss detected - bail out
	}
	data.Mantissa = x >> shift

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
	var e2, e10, i, j, k, l, removed int32
	var m2, q, a, u, b, v, c, out uint32
	var lastRemovedDigit uint8
	var mul uint64
	var even, acceptBounds, aTrailingZeros, bTrailingZeros bool

	// Step 1 - parse and preparations
	if data.Exponent == 0 {
		e2 = 1 -
			hestiaMATH.BIAS_FLOAT32_EXPONENT -
			hestiaMATH.BITS_FLOAT32_MANTISSA -
			2
		m2 = data.Mantissa
	} else {
		e2 = int32(data.Exponent) -
			hestiaMATH.BIAS_FLOAT32_EXPONENT -
			hestiaMATH.BITS_FLOAT32_MANTISSA -
			2
		m2 = uint32(1)<<hestiaMATH.BITS_FLOAT32_MANTISSA | data.Mantissa
	}
	even = m2&1 == 0
	acceptBounds = even

	// Step 2- determine valid decimal representations interval
	u = 1
	if data.Mantissa != 0 || data.Exponent <= 1 {
		u = 2
	}

	/* u = 4*m2 - u // mm -> compute as 'v-u' */
	v = 4 * m2 // mv
	/* w = 4*m2 + 2 // mp -> compute as '4*m2 + 2 */

	// Step 3 - Convert u, v, w to decimal base
	if e2 >= 0 {
		q = __s32_LOG10_2PowerOf(e2)
		e10 = int32(q)
		k = value_POWER_5_INVERSE_BITS32 + __s32_LOG2_5PowerOf(int32(q)) - 1
		i = -e2 + int32(q) + k
		mul = __m32_5PowerInverse_Divide_By_2Power_Multiplier(q)
		a = __s32_Multiply_By_Shift_32_Bits(4*m2-u, mul, i) // vm
		b = __s32_Multiply_By_Shift_32_Bits(v, mul, i)      // vr
		c = __s32_Multiply_By_Shift_32_Bits(4*m2+2, mul, i) // vp

		if q != 0 && (c-1)/10 <= a/10 {
			// learn at least 1 removed digit
			l = value_POWER_5_INVERSE_BITS32 + __s32_LOG2_5PowerOf(int32(q)-1) - 1
			mul = __m32_5PowerInverse_Divide_By_2Power_Multiplier(q - 1)
			lastRemovedDigit = uint8(
				__s32_Multiply_By_Shift_32_Bits(v, mul, -e2+int32(q-1)+l) % 10,
			)
		}

		if q <= 9 {
			// largest 5 power = 5^10 that fits in 24-bits number.
			// Only 1 of u, v, and w can be a multiple 5, if any.
			switch {
			case v%5 == 0:
				bTrailingZeros = __s32_Factor_Of_5(v, q)
			case acceptBounds:
				aTrailingZeros = __s32_Factor_Of_5(v-u, q)
			case __s32_Factor_Of_5(v+2, q):
				c--
			}
		}
	} else {
		q = __s32_LOG10_5PowerOf(-e2) // modulus the negative e2
		e10 = e2 + int32(q)
		i = -e2 - int32(q)
		k = __s32_LOG2_5PowerOf(i) - value_POWER_5_BITS32
		j = int32(q) - k

		mul = __m32_5Power_Divide_By_2Power_Multiplier(uint32(i))
		a = __s32_Multiply_By_Shift_32_Bits(4*m2-u, mul, j) // vm
		b = __s32_Multiply_By_Shift_32_Bits(4*m2, mul, j)   // vr
		c = __s32_Multiply_By_Shift_32_Bits(4*m2+2, mul, j) // vp

		if q != 0 && (c-1)/10 <= a/10 {
			j = int32(q) -
				1 -
				(__s32_LOG2_5PowerOf(i) - value_POWER_5_BITS32)
			mul = __m32_5Power_Divide_By_2Power_Multiplier(uint32(i + 1))
			lastRemovedDigit = uint8(
				__s32_Multiply_By_Shift_32_Bits(v, mul, j) % 10,
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
				aTrailingZeros = (u == 2)
			} else {
				// for 'w = m2 + 2', it always has at least 1 trailing 0-bit
				c--
			}
		} else if q < 31 {
			bTrailingZeros = __s32_Factor_Of_2(v, q-1)
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
