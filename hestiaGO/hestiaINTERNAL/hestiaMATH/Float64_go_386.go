// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2010 The Go Authors <https://cs.opensource.google/go/go>
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

//go:build !tinygo && 386

package hestiaMATH

func _s64_Modf_FLOAT64(input float64) (round float64, fraction float64) {
	if input < 1 {
		switch {
		case input < 0:
			round, fraction = S64_Modf_FLOAT64(-input)
			return -round, -fraction
		case input == 0:
			return input, input // 0, .0
		}

		return 0, input
	}

	x := S64_IEEE754_FloatToBits(input)
	e := ((x & MASK_FLOAT64_EXPONENT) >> BITS_FLOAT64_MANTISSA) - BIAS_FLOAT64_EXPONENT

	if e < 64-12 {
		x &^= 1<<(64-12-e) - 1
	}

	round = S64_IEEE754_BitsToFloat(x)
	fraction = input - round
	return round, fraction
}

func _s64_Floor_FLOAT64(input float64) (round float64)
