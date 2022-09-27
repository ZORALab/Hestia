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

const (
	MAX_UINT   = ^uint(0)
	MAX_UINT8  = 1<<8 - 1
	MAX_UINT16 = 1<<16 - 1
	MAX_UINT32 = 1<<32 - 1
	MAX_UINT64 = 1<<64 - 1

	MAX_INT   = int(^uint(0) >> 1)
	MAX_INT8  = 1<<7 - 1
	MAX_INT16 = 1<<15 - 1
	MAX_INT32 = 1<<31 - 1
	MAX_INT64 = 1<<63 - 1

	MIN_INT   = -int(^uint(0)>>1) - 1
	MIN_INT8  = -1 << 7
	MIN_INT16 = -1 << 15
	MIN_INT32 = -1 << 31
	MIN_INT64 = -1 << 63
)

//nolint:gochecknoglobals
var (
	cpu_size uint64
)

func CPU() uint64 {
	if cpu_size > 0 {
		return cpu_size
	}

	i := ^uint(0) ^ (^uint(0) >> 1)
	for i != 0 {
		cpu_size++
		i >>= 1
	}

	return cpu_size
}

func Len(x uint) uint64 {
	var i, a, b uint64

	// loop through bit shifting to find the highest bit
	for i = 1; i <= MAX_UINT64; i <<= 1 {
		x |= (x >> i)

		a = uint64(x ^ (x >> 1))
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

func TrailingZero(x uint) uint64 {
	var count uint64

	for (x & 1) == 0 {
		x >>= 1
		count++
	}

	return count
}
