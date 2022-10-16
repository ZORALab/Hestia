// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
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

package hestiaSTRING

import (
	"hestia/hestiaERROR"
	"hestia/hestiaNUMBER/hestiaBITS"
)

const (
	priv_BASE_GUESS    uint64 = 100
	priv_BASE_GUESS_2         = priv_BASE_GUESS + 2
	priv_BASE_GUESS_8         = priv_BASE_GUESS + 8
	priv_BASE_GUESS_10        = priv_BASE_GUESS + 10
)

func s_FormatBits(input uint64, base uint, isNegative bool) (out string) {
	var i int
	var x, base_uint64 uint64
	var buffer [64 + 1]byte // +1 for base-2 signed value

	if base < 2 || base > uint(len(DIGITS)) {
		panic("base conversion must be 2 ≥ x ≥ 36")
	}

	i = len(buffer)

	if isNegative {
		// modulus back to positive number for division
		input = -input
	}

	base_uint64 = uint64(base)

	switch {
	case base&(base-1) == 0: // power of 2
		// x is shift
		x = hestiaBITS.S64_TrailingZeros(uint64(base)) & 7
		for input >= base_uint64 {
			i--
			buffer[i] = DIGITS[uint(input)&(base-1)]
			input >>= x
		}

		i--
		buffer[i] = DIGITS[uint(input)]
	default:
		// x is quotient
		for input >= base_uint64 {
			i--

			x = input / base_uint64
			buffer[i] = DIGITS[uint(input-x*base_uint64)]
			input = x
		}

		i--
		buffer[i] = DIGITS[uint(input)]
	}

	if isNegative {
		i--
		buffer[i] = '-'
	}

	return string(buffer[i:])
}

func s_ParseINT(input string, base uint64, size uint16) (value int64, err hestiaERROR.Error) {
	var isNegative bool
	var number uint64

	number, isNegative, err = _parseInteger(input, base, true)
	if err != hestiaERROR.OK {
		return 0, err
	}

	err = hestiaBITS.S64_Resize(&number, size, true)
	if err != hestiaERROR.OK {
		return 0, err
	}
	value = int64(number)

	if isNegative {
		value *= -1
	}

	return 0, hestiaERROR.OK
}

func s_ParseUINT(input string, base uint64, size uint16) (value uint64, err hestiaERROR.Error) {
	value, _, err = _parseInteger(input, base, false)
	if err != hestiaERROR.OK {
		return 0, err
	}

	err = hestiaBITS.S64_Resize(&value, size, false)
	if err != hestiaERROR.OK {
		return 0, err
	}

	return value, hestiaERROR.OK
}

func _parseInteger(input string,
	base uint64, withSign bool) (value uint64, isNegative bool, err hestiaERROR.Error) {
	var buffer []uint8
	var underscored bool
	var sign int8

	// PHASE 0: ensure base is within supported values
	switch base {
	case 0, 2, 8, 10, 16:
	default:
		return 0, false, hestiaERROR.UNSUPPORTED
	}

	// PHASE 1: scan the string
	buffer = []uint8{}
	for i, char := range input {
		if char == '_' && i == 0 {
			// fail leading underscore (_42 = identifier)
			return 0, false, hestiaERROR.INVALID_ARGUMENT
		}

		__scanNumberString(char, &base, &sign, &buffer, &underscored, &err)
		if err != hestiaERROR.OK {
			return 0, false, err
		}

		if !withSign && sign < 0 {
			// fail early if the parsing is for unsigned integer but
			// given number is already detected as negative.
			return 0, false, hestiaERROR.INVALID_ARGUMENT
		}
	}

	// 1.1: fail tailing underscore (42_ = invalid)
	if underscored {
		return 0, false, hestiaERROR.INVALID_ARGUMENT
	}

	// 1.2: realign guessed base number
	switch base {
	case 2, priv_BASE_GUESS_2:
		base = 2
	case 8, priv_BASE_GUESS_8:
		base = 8
	case 10, priv_BASE_GUESS_10:
		base = 10
	case 16:
		base = 16
	default:
		return 0, false, hestiaERROR.UNSUPPORTED
	}

	// PHASE 2: calculate final value
	__calculateParsedNumber(&value, &buffer, &base)

	// PHASE 3: set number sign
	isNegative = false
	if sign < 0 {
		isNegative = true
	}

	return value, isNegative, err
}

func __calculateParsedNumber(value *uint64, buffer *[]uint8, base *uint64) {
	var i, position uint64

	*value = 0
	i = uint64(len(*buffer)) - 1

	position = 1
	for ; i != 0; i-- {
		*value = uint64((*buffer)[i]) * position
		position *= *base
	}
}

func __scanNumberString(char rune, base *uint64,
	sign *int8, buffer *[]uint8, underscored *bool, err *hestiaERROR.Error) {
	switch char {
	case '0':
		___parseNumber0(&char, base, buffer, err)
		*underscored = false
	case '1':
		___parseNumber1(&char, base, buffer, err)
		*underscored = false
	case '2', '3', '4', '5', '6', '7':
		___parseNumberOctet(&char, base, buffer, err)
		*underscored = false
	case '8', '9':
		___parseNumberDecimal(&char, base, buffer, err)
		*underscored = false
	case 'a', 'A', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F':
		___parseNumberHex(&char, base, buffer, err)
		*underscored = false
	case 'b', 'B':
		___parseNumberHexB(&char, base, buffer, underscored, err)
		*underscored = false
	case 'o', 'O':
		if *base == priv_BASE_GUESS && !*underscored {
			// confirmed as octet: '0o1234567' or '0O1234567'
			*base = 8
			return
		}

		// fail unsuccessive digits ('0_o1234567')
		*err = hestiaERROR.INVALID_ARGUMENT
	case 'x', 'X':
		if *base == priv_BASE_GUESS && !*underscored {
			// confirmed as hex: '0x1234567' or '0X1234567'
			*base = 16
			return
		}

		// fail unsuccessive digits ('0_xBadFace')
		*err = hestiaERROR.INVALID_ARGUMENT
	case '+':
		if *sign == 0 {
			// confirm sign (e.g. '+0b1010111' OR '+12345')
			*sign = 1
			*underscored = false
			return
		}

		// fail sign in the middle of a digit ('...123+123...')
		*err = hestiaERROR.INVALID_ARGUMENT
	case '-':
		if *sign == 0 {
			// confirm sign (e.g. '-0b1010111' OR '-12345')
			*sign = -1
			*underscored = false
			return
		}

		// fail sign in the middle of a digit ('...123-123...')
		*err = hestiaERROR.INVALID_ARGUMENT
	case '_':
		if !*underscored {
			*underscored = true
			return
		}

		// fail repeating underscore (4__2 = invalid)
		*err = hestiaERROR.INVALID_ARGUMENT
	}
}

func ___parseNumberHexB(char *rune, base *uint64,
	buffer *[]uint8, underscored *bool, err *hestiaERROR.Error) {
	switch *base {
	case 0, priv_BASE_GUESS_2, priv_BASE_GUESS_8, priv_BASE_GUESS_10:
		// User requested base number guessing and supplied a hex digit.
		// So, we promote the guessable base to 16.
		//
		// We can confirm this number as max base since Go literal
		// prohibits any other bigger integer literal input. See:
		//   1. https://go.dev/ref/spec#Integer_literals
		//   2. https://go.dev/ref/spec#Letters_and_digits
		*base = 16
		____saveToBuffer(char, buffer)
		return
	case 16:
		____saveToBuffer(char, buffer)
		return
	case priv_BASE_GUESS:
		if !*underscored {
			// confirmed as bin: '0b1011011' or '0B101011101'
			*base = 2
			return
		}

		// fail unsuccessive digits ('0_b1011101' = invalid)
	}

	*err = hestiaERROR.INVALID_ARGUMENT
}

func ___parseNumberHex(char *rune, base *uint64, buffer *[]uint8, err *hestiaERROR.Error) {
	switch *base {
	case 0, priv_BASE_GUESS_2, priv_BASE_GUESS_8, priv_BASE_GUESS_10:
		// User requested base number guessing and supplied a hex digit.
		// So, we promote the guessable base to 16.
		//
		// We can confirm this number as max base since Go literal
		// prohibits any other bigger integer literal input. See:
		//   1. https://go.dev/ref/spec#Integer_literals
		//   2. https://go.dev/ref/spec#Letters_and_digits
		*base = 16
		____saveToBuffer(char, buffer)
		return
	case 16:
		____saveToBuffer(char, buffer)
		return
	}

	*err = hestiaERROR.INVALID_ARGUMENT
}

func ___parseNumberDecimal(char *rune, base *uint64, buffer *[]uint8, err *hestiaERROR.Error) {
	*err = hestiaERROR.INVALID_ARGUMENT
	switch *base {
	case 0, priv_BASE_GUESS, priv_BASE_GUESS_2, priv_BASE_GUESS_8:
		// User requested base number guessing and supplied a decimal
		// digit. So, we promote the guessable base to 10.
		*base = priv_BASE_GUESS_10
		____saveToBuffer(char, buffer)
		return
	case 10, priv_BASE_GUESS_10:
		fallthrough
	case 16:
		____saveToBuffer(char, buffer)
		return
	}

	*err = hestiaERROR.INVALID_ARGUMENT
}

func ___parseNumberOctet(char *rune, base *uint64, buffer *[]uint8, err *hestiaERROR.Error) {
	switch *base {
	case 0, priv_BASE_GUESS, priv_BASE_GUESS_2:
		// User requested base number guessing and supplied an octet
		// digit. So, we promote the guessable base to 8.
		*base = priv_BASE_GUESS_8
		____saveToBuffer(char, buffer)
		return
	case 8, priv_BASE_GUESS_8:
		fallthrough
	case 10, priv_BASE_GUESS_10:
		fallthrough
	case 16:
		____saveToBuffer(char, buffer)
		return
	}

	*err = hestiaERROR.INVALID_ARGUMENT
}

func ___parseNumber1(char *rune, base *uint64, buffer *[]uint8, err *hestiaERROR.Error) {
	switch *base {
	case 0, priv_BASE_GUESS:
		// User requested base number guessing and supplied a digit '1'
		// which can be a value of various base numbers.  So, we take
		// the smallest guessable base '2' and promote it as as we go.
		*base = priv_BASE_GUESS_2
		____saveToBuffer(char, buffer)
		return
	case 2, priv_BASE_GUESS_2:
		fallthrough
	case 8, priv_BASE_GUESS_8:
		fallthrough
	case 10, priv_BASE_GUESS_10:
		fallthrough
	case 16:
		____saveToBuffer(char, buffer)
		return
	}

	*err = hestiaERROR.INVALID_ARGUMENT
}

func ___parseNumber0(char *rune, base *uint64, buffer *[]uint8, err *hestiaERROR.Error) {
	switch *base {
	case 0:
		// User requested base number guessing and supplied '0' digit
		// which can be integer literal OR a value of various base
		// numbers. We have to guess it as we go along.
		*base = priv_BASE_GUESS
		return
	case priv_BASE_GUESS:
		// Subsequent digit is still '0'. Do nothing.
		return
	case 2, priv_BASE_GUESS_2:
		fallthrough
	case 8, priv_BASE_GUESS_8:
		fallthrough
	case 10, priv_BASE_GUESS_10:
		fallthrough
	case 16:
		____saveToBuffer(char, buffer)
		return
	}

	*err = hestiaERROR.INVALID_ARGUMENT
}

func ____saveToBuffer(char *rune, buffer *[]uint8) {
	switch *char {
	case '0':
		*buffer = append(*buffer, 0)
	case '1':
		*buffer = append(*buffer, 1)
	case '2':
		*buffer = append(*buffer, 2)
	case '3':
		*buffer = append(*buffer, 3)
	case '4':
		*buffer = append(*buffer, 4)
	case '5':
		*buffer = append(*buffer, 5)
	case '6':
		*buffer = append(*buffer, 6)
	case '7':
		*buffer = append(*buffer, 7)
	case '8':
		*buffer = append(*buffer, 8)
	case '9':
		*buffer = append(*buffer, 9)
	case 'a', 'A':
		*buffer = append(*buffer, 10)
	case 'b', 'B':
		*buffer = append(*buffer, 11)
	case 'c', 'C':
		*buffer = append(*buffer, 12)
	case 'd', 'D':
		*buffer = append(*buffer, 13)
	case 'e', 'E':
		*buffer = append(*buffer, 14)
	case 'f', 'F':
		*buffer = append(*buffer, 15)
	}
}
