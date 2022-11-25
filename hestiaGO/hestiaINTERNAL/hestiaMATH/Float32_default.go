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

package hestiaMATH

func _s32_Floor_FLOAT32(input float32) (round float32) {
	var fraction float32

	switch {
	case input == 0:
		fallthrough
	case S32_IEEE754_IsNaN_FLOAT32(input):
		fallthrough
	case S32_IEEE754_IsINF_FLOAT32(input, 0):
		return input
	case input < 0:
		round, fraction = S32_Modf_FLOAT32(-input)
		if fraction != 0.0 {
			round++
		}
		return -round
	default:
		round, _ = S32_Modf_FLOAT32(input)
		return round
	}
}

func _s32_Modf_FLOAT32(input float32) (round float32, fraction float32) {
	if input < 1 {
		switch {
		case input < 0:
			round, fraction = S32_Modf_FLOAT32(-input)
			return -round, -fraction
		case input == 0:
			return input, input // 0, .0
		}

		return 0, input
	}

	x := S32_IEEE754_FloatToBits(input)
	e := ((x & MASK_FLOAT32_EXPONENT) >> BITS_FLOAT32_MANTISSA) - BIAS_FLOAT32_EXPONENT

	if e < 32-8 {
		x &^= 1<<(32-8-e) - 1
	}

	round = S32_IEEE754_BitsToFloat(x)
	fraction = input - round
	return round, fraction
}
