// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2012 Desmond Hume <https://stackoverflow.com/users/944687/desmond-hume>
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

import (
	"math"
)

const (
	const_LOG2_DESMOND_HUME_32_BITS uint32 = 0x07C4ACDD
	const_LOG2_DESMOND_HUME_64_BITS uint64 = 0x07EDD5E59A4E28C2
)

//nolint:gochecknoglobals
var (
	const_LOG2_DESMOND_HUME_TABLE_32_BITS = [32]int32{
		0, 9, 1, 10, 13, 21, 2, 29,
		11, 14, 16, 18, 22, 25, 3, 30,
		8, 12, 20, 28, 15, 17, 24, 7,
		19, 27, 23, 6, 26, 5, 4, 31,
	}

	const_LOG2_DESMOND_HUME_TABLE_64_BITS = [64]int64{
		63, 0, 58, 1, 59, 47, 53, 2,
		60, 39, 48, 27, 54, 33, 42, 3,
		61, 51, 37, 40, 49, 18, 28, 20,
		55, 30, 34, 11, 43, 14, 22, 4,
		62, 57, 46, 52, 38, 26, 32, 41,
		50, 36, 17, 19, 29, 10, 13, 21,
		56, 45, 25, 31, 35, 16, 9, 12,
		44, 24, 15, 8, 23, 7, 6, 5,
	}
)

// Doc: https://stackoverflow.com/questions/11376288/fast-computing-of-log2-for-64-bit-integers
func M32_DesmondFume_LOG2(input uint32) int32 {
	input |= input >> 1
	input |= input >> 2
	input |= input >> 4
	input |= input >> 8
	input |= input >> 16
	input = (input * const_LOG2_DESMOND_HUME_32_BITS) >> 27 // 32 - 5 rounds
	return const_LOG2_DESMOND_HUME_TABLE_32_BITS[input]
}

// Doc: https://stackoverflow.com/questions/11376288/fast-computing-of-log2-for-64-bit-integers
func M64_DesmondFume_LOG2(input uint64) int64 {
	input |= input >> 1
	input |= input >> 2
	input |= input >> 4
	input |= input >> 8
	input |= input >> 16
	input |= input >> 32
	input = (input * const_LOG2_DESMOND_HUME_64_BITS) >> 58 // 64 - 6 rounds
	return const_LOG2_DESMOND_HUME_TABLE_64_BITS[input]
}

func S32_LOG2(input float32) float32 {
	return float32(math.Log(float64(input)))
}

func S64_LOG2(input float64) float64 {
	return math.Log(input)
}

func S32_LOG10(input float32) float32 {
	return float32(math.Log10(float64(input)))
}

func S64_LOG10(input float64) float64 {
	return math.Log10(input)
}

func S32_LOG(input float32) float32 {
	return float32(math.Log(float64(input)))
}

func S64_LOG(input float64) float64 {
	return math.Log(input)
}
