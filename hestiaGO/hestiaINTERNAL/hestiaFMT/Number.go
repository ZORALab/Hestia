// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
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

package hestiaFMT

func S8_FormatUINT8(number uint8, base uint8, lettercase Lettercase) (out []rune) {
	var i uint8

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 8 // MAX: 8-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s8_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S8_FormatINT8(input int8, base uint8, lettercase Lettercase) (out []rune) {
	var i, number uint8

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 8 + 1 // MAX: 8-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint8(input * -1)
	} else {
		number = uint8(input)
	}

	// process number according to base
	_s8_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func S16_FormatUINT16(number uint16, base uint16, lettercase Lettercase) (out []rune) {
	var i uint16

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 16 // MAX: 16-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s16_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S16_FormatINT16(input int16, base uint16, lettercase Lettercase) (out []rune) {
	var i, number uint16

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 16 + 1 // MAX: 16-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint16(input * -1)
	} else {
		number = uint16(input)
	}

	// process number according to base
	_s16_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func S32_FormatUINT32(number uint32, base uint32, lettercase Lettercase) (out []rune) {
	var i uint32

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 32 // MAX: 32-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s32_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S32_FormatINT32(input int32, base uint32, lettercase Lettercase) (out []rune) {
	var i, number uint32

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 32 + 1 // MAX: 32-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint32(input * -1)
	} else {
		number = uint32(input)
	}

	// process number according to base
	_s32_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func S64_FormatUINT64(number uint64, base uint64, lettercase Lettercase) (out []rune) {
	var i uint64

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 64 // MAX: 64-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s64_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S64_FormatINT64(input int64, base uint64, lettercase Lettercase) (out []rune) {
	var i uint64
	var number uint64

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 64 + 1 // MAX: 64-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint64(input * -1)
	} else {
		number = uint64(input)
	}

	// process number according to base
	_s64_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

//nolint:goconst
func _SN_DIGIT_To_NUMBER(input rune) uint8 {
	switch input {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'a', 'A':
		return 10
	case 'b', 'B':
		return 11
	case 'c', 'C':
		return 12
	case 'd', 'D':
		return 13
	case 'e', 'E':
		return 14
	case 'f', 'F':
		return 15
	case 'g', 'G':
		return 16
	case 'h', 'H':
		return 17
	case 'i', 'I':
		return 18
	case 'j', 'J':
		return 19
	case 'k', 'K':
		return 20
	case 'l', 'L':
		return 21
	case 'm', 'M':
		return 22
	case 'n', 'N':
		return 23
	case 'o', 'O':
		return 24
	case 'p', 'P':
		return 25
	case 'q', 'Q':
		return 26
	case 'r', 'R':
		return 27
	case 's', 'S':
		return 28
	case 't', 'T':
		return 29
	case 'u', 'U':
		return 30
	case 'v', 'V':
		return 31
	case 'w', 'W':
		return 32
	case 'x', 'X':
		return 33
	case 'y', 'Y':
		return 34
	case 'z', 'Z':
		return 35
	default:
		panic("unknown digit character!")
	}
}

func _SN_NUMBER_To_DIGIT(number uint8, lettercase *Lettercase) rune {
	var charset []rune

	switch *lettercase {
	case LETTERCASE_UPPER:
		charset = []rune("0123456789ABCDEFGHIJKLMONPQRSTUVWXYZ")
	case LETTERCASE_LOWER:
		charset = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	default:
		panic("unknown lettercase!")
	}

	return charset[number]
}

func _s8_formatNumber(buffer *[]rune,
	base *uint8, number *uint8, i *uint8, lettercase *Lettercase) {
	var x uint8

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(*number-x*(*base), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(*number, lettercase)
}

func _s16_formatNumber(buffer *[]rune,
	base *uint16, number *uint16, i *uint16, lettercase *Lettercase) {
	var x uint16

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number-x*(*base)), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number), lettercase)
}

func _s32_formatNumber(buffer *[]rune,
	base *uint32, number *uint32, i *uint32, lettercase *Lettercase) {
	var x uint32

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number-x*(*base)), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number), lettercase)
}

func _s64_formatNumber(buffer *[]rune,
	base *uint64, number *uint64, i *uint64, lettercase *Lettercase) {
	var x uint64

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number-x*(*base)), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number), lettercase)
}
