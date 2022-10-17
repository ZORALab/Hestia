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
//
// NOTE: test codes are inside hestiaSTRING package.

package hestiaFMT

//nolint:gochecknoglobals
var (
	char_DIGITS_UPPER = []rune("0123456789ABCDEFGHIJKLMONPQRSTUVWXYZ")
	char_DIGITS_LOWER = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
)

func FormatUINT8(number uint8, base uint8, lettercase Lettercase) (out []rune) {
	var i, x uint8
	var charset *[]rune

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
	_processLettercase(&lettercase, &charset)

	// process number according to base
	for number >= base {
		i--

		x = number / base
		out[i] = (*charset)[number-x*base]
		number = x
	}

	// process last remainder
	i--
	out[i] = (*charset)[number]

	// done conversion return output
	return out[i:]
}

func FormatUINT16(number uint16, base uint16, lettercase Lettercase) (out []rune) {
	var i, x uint16
	var charset *[]rune

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
	_processLettercase(&lettercase, &charset)

	// process number according to base
	for number >= base {
		i--

		x = number / base
		out[i] = (*charset)[number-x*base]
		number = x
	}

	// process last remainder
	i--
	out[i] = (*charset)[number]

	// done conversion return output
	return out[i:]
}

func FormatUINT32(number uint32, base uint32, lettercase Lettercase) (out []rune) {
	var i uint32
	var charset *[]rune

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
	_processLettercase(&lettercase, &charset)

	// process number according to base
	_s32_formatNumber(&out, &base, &number, &i, charset)

	// done conversion return output
	return out[i:]
}

func FormatINT32(input int32, base uint32, lettercase Lettercase) (out []rune) {
	var i, number uint32
	var charset *[]rune

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
	_processLettercase(&lettercase, &charset)

	// modulus negative number for division
	if input < 0 {
		number = uint32(input * -1)
	} else {
		number = uint32(input)
	}

	// process number according to base
	_s32_formatNumber(&out, &base, &number, &i, charset)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func FormatUINT64(number uint64, base uint64, lettercase Lettercase) (out []rune) {
	var i uint64
	var charset *[]rune

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
	_processLettercase(&lettercase, &charset)

	// process number according to base
	_s64_formatNumber(&out, &base, &number, &i, charset)

	// done conversion return output
	return out[i:]
}

func FormatINT64(input int64, base uint64, lettercase Lettercase) (out []rune) {
	var i uint64
	var charset *[]rune
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
	_processLettercase(&lettercase, &charset)

	// modulus negative number for division
	if input < 0 {
		number = uint64(input * -1)
	} else {
		number = uint64(input)
	}

	// process number according to base
	_s64_formatNumber(&out, &base, &number, &i, charset)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func _s32_formatNumber(buffer *[]rune,
	base *uint32, number *uint32, i *uint32, charset *[]rune) {
	var x uint32

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = (*charset)[*number-x*(*base)]
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = (*charset)[*number]
}

func _s64_formatNumber(buffer *[]rune,
	base *uint64, number *uint64, i *uint64, charset *[]rune) {
	var x uint64

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = (*charset)[*number-x*(*base)]
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = (*charset)[*number]
}

func _processLettercase(input *Lettercase, charset **[]rune) {
	switch *input {
	case LETTERCASE_UPPER:
		*charset = &char_DIGITS_UPPER
	case LETTERCASE_LOWER:
		*charset = &char_DIGITS_LOWER
	default:
		panic("unknown lettercase!")
	}
}
