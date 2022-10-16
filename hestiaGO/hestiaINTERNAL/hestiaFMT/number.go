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

type numberType uint

const (
	number_DECIMALLESS              numberType = iota // (e.g. 123)
	number_SCIENTIFIC                                 // (e.g. -1.234456e+78)
	number_SCIENTIFIC_AUTO_EXPONENT                   // (e.g. -1.234 OR -1.234456e+78)
	number_DECIMAL_NO_EXPONENT                        // (e.g -123.456)
	number_HEX                                        // (e.g. -0x1.23abcp+20)
)

const (
	char_DIGITS_UPPER = "0123456789ABCDEFGHIJKLMONPQRSTUVWXYZ"
	char_DIGITS_LOWER = "0123456789abcdefghijklmnopqrstuvwxyz"
)

type numberSTR struct {
	arg        any
	base       uint16
	lettercase Lettercase
	width      []rune
	precision  []rune
	format     numberType
}

func _parseNumber(str *numberSTR) (out []rune) {
	// to be developed later
	return []rune{'(', 'N', 'U', 'M', 'B', 'E', 'R', '=', 'b', 'a', 'd', ')'}
}

func FormatUINT16(number uint16, base uint16, lettercase Lettercase) (out []rune) {
	var i, x uint16
	var charset []rune

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
	switch lettercase {
	case LETTERCASE_UPPER:
		charset = []rune(char_DIGITS_UPPER)
	case LETTERCASE_LOWER:
		charset = []rune(char_DIGITS_LOWER)
	}

	// process number according to base
	for number >= base {
		i--

		x = number / base
		out[i] = charset[number-x*base]
		number = x
	}

	// process last remainder
	i--
	out[i] = charset[number]

	// done conversion return output
	return out[i:]
}

func FormatUINT32(number uint32, base uint32, lettercase Lettercase) (out []rune) {
	var i, x uint32
	var charset []rune

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
	switch lettercase {
	case LETTERCASE_UPPER:
		charset = []rune(char_DIGITS_UPPER)
	case LETTERCASE_LOWER:
		charset = []rune(char_DIGITS_LOWER)
	}

	// process number according to base
	for number >= base {
		i--

		x = number / base
		out[i] = charset[number-x*base]
		number = x
	}

	// process last remainder
	i--
	out[i] = charset[number]

	// done conversion return output
	return out[i:]
}

func FormatUINT64(number uint64, base uint64, lettercase Lettercase) (out []rune) {
	var i, x uint64
	var charset []rune

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 64 // MAX: 64-bits in base-2
	out = make([]rune, i)
	switch lettercase {
	case LETTERCASE_UPPER:
		charset = []rune(char_DIGITS_UPPER)
	case LETTERCASE_LOWER:
		charset = []rune(char_DIGITS_LOWER)
	}

	// process number according to base
	for number >= base {
		i--

		x = number / base
		out[i] = charset[number-x*base]
		number = x
	}

	// process last remainder
	i--
	out[i] = charset[number]

	// done conversion return output
	return out[i:]
}
