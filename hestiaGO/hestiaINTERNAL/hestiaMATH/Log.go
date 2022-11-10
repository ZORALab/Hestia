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

func S32_LOG2(input float32) float32 {
	return float32(S64_LOG2(float64(input))) // use 64-bits for lower math error rate
}

func S64_LOG2(input float64) float64 {
	return _s64_LOG2(input)
}

func S32_LOG10(input float32) float32 {
	return float32(S64_LOG10(float64(input)))
}

func S64_LOG10(input float64) float64 {
	return _s64_LOG10(input)
}

func S32_LOG(input float32) float32 {
	return float32(S64_LOG(float64(input))) // use 64-bits for lower math error rate
}

func S64_LOG(input float64) float64 {
	switch {
	case S64_IEEE754_IsNaN_FLOAT64(input):
		return input
	case S64_IEEE754_IsINF_FLOAT64(input, 1):
		return input
	case input == 0:
		return S64_IEEE754_BitsToFloat(
			MASK_FLOAT64_SIGN | MASK_FLOAT64_EXPONENT,
		)
	}

	return _s64_LOG(input)
}
