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

	// 1.79769313486231570814527423731704356798070e+308
	MAX_FLOAT64 = 0x1p1023 * (1 + (1 - 0x1p-52))
	// -4.9406564584124654417656879286822137236505980e-324
	MIN_FLOAT64 = -(0x1p-1022 * 0x1p-52)
)

//nolint:lll
const (
	SQRT_2 = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
)

//nolint:lll
const (
	LN_2  = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
	LN_10 = 2.30258509299404568401799145468436420760110148862877297603332790  // https://oeis.org/A002392
)
