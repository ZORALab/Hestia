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

package hestiaBITS

import (
	"hestia/hestiaNUMBER"
)

func Len(x uint) uint {
	var i, a, b uint

	// loop through bit shifting to find the highest bit
	for i = 1; i <= hestiaNUMBER.MAX_UINT64; i <<= 1 {
		x |= (x >> i)

		a = x ^ (x >> 1)
		if a == b {
			// max bit reached
			break
		}

		b = a
	}

	// convert back to decimal counting value
	a = 0
	for b != 0 {
		a++
		b >>= 1
	}

	return a
}

func TrailingZero(x uint) uint {
	var count uint

	for (x & 1) == 0 {
		x >>= 1
		count++
	}

	return count
}
