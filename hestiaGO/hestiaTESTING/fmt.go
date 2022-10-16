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

package hestiaTESTING

import (
	"unicode"
)

const (
	char_DIGITS = "0123456789abcdefghijklmnopqrstuvwxyz"
)

func _renderNumber(number, base uint64) (out string) {
	var i int
	var x uint64
	var buffer [64]byte

	if number == 0 {
		return "0"
	}

	i = len(buffer)
	for number >= base {
		i--

		x = number / base
		buffer[i] = char_DIGITS[uint(number-x*base)]
		number = x
	}

	i--
	buffer[i] = char_DIGITS[uint(number)]

	return string(buffer[i:])
}

func __trimWhitespace(source string) string {
	var leftIndex, rightIndex uint64
	var setRight, next bool

	if source == "" {
		return source
	}

	setRight = false
	for i, char := range source {
		if !unicode.IsSpace(char) {
			if i == 0 {
				setRight = true
			}

			if !setRight {
				leftIndex = uint64(i)
			}

			rightIndex = uint64(i)
			next = true
			setRight = true
			continue
		}

		if next {
			rightIndex = uint64(i)
			next = false
		}
	}

	if next {
		rightIndex = uint64(len(source)) // complete last character
	}

	return source[leftIndex:rightIndex]
}
