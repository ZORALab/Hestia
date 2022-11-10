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

//go:build !tinygo && amd64

package hestiaMATH

func _s64_LOG2(input float64) float64 {
	fraction, exponent := S64_Frexp_FLOAT64(input)
	if fraction == 0.5 {
		// return exact power of 2 without depending on
		// Log(0.5)*(1/Ln2)+exp being exactly exp-1
		return float64(exponent - 1)
	}

	return S64_LOG(fraction)*(1/LN_2) + float64(exponent)
}

func _s64_LOG10(input float64) float64 {
	return S64_LOG(input) * (1 / LN_10)
}

func _s64_LOG(input float64) float64
