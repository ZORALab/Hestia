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
	"hestia/hestiaNUMBER/hestiaBITS"
)

func s_FormatBits(input uint64, base uint64, isNegative bool) (out string) {
	var i int
	var x uint64
	var buffer [64 + 1]byte // +1 for base-2 signed value

	if base < 2 || base > uint64(len(DIGITS)) {
		panic("base conversion must be 2 ≥ x ≥ 36")
	}

	i = len(buffer)

	if isNegative {
		// modulus back to positive number for division
		input = -input
	}

	switch {
	case base&(base-1) == 0: // power of 2
		// x is shift
		x = hestiaBITS.TrailingZero(uint(base)) & 7
		for input >= base {
			i--
			buffer[i] = DIGITS[uint(input)&uint(base-1)]
			input >>= x
		}

		i--
		buffer[i] = DIGITS[uint(input)]
	default:
		// x is quotient
		for input >= base {
			i--

			x = input / base
			buffer[i] = DIGITS[uint(input-x*base)]
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
