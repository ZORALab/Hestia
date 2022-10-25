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

package hestiaMATH

func S64_TrailingZeros_64Bits(x uint64) uint64 {
	var count uint64

	if x == 0 {
		return 0
	}

	for (x & 1) == 0 {
		x >>= 1
		count++
	}

	return count
}

func S32_TrailingZeros_32Bits(x uint32) uint32 {
	var count uint32

	if x == 0 {
		return 0
	}

	for (x & 1) == 0 {
		x >>= 1
		count++
	}

	return count
}

func S16_TrailingZeros_16Bits(x uint16) uint16 {
	var count uint16

	if x == 0 {
		return 0
	}

	for (x & 1) == 0 {
		x >>= 1
		count++
	}

	return count
}

func S8_TrailingZeros_8Bits(x uint8) uint8 {
	var count uint8

	if x == 0 {
		return 0
	}

	for (x & 1) == 0 {
		x >>= 1
		count++
	}

	return count
}
