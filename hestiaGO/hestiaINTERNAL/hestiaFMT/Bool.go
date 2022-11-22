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

package hestiaFMT

func SN_ParseBOOL(input string) (bool, Error) {
	switch input {
	case "1", "t", "true", "True", "TRUE":
		return true, ERROR_OK
	case "0", "f", "false", "False", "FALSE":
		return false, ERROR_OK
	default:
		return false, ERROR_INPUT_INVALID
	}
}

func SN_FormatBOOL(input bool, lettercase Lettercase) []rune {
	switch {
	case input && lettercase == LETTERCASE_UPPER:
		return []rune{'T', 'R', 'U', 'E'}
	case input:
		return []rune{'t', 'r', 'u', 'e'}
	case lettercase == LETTERCASE_UPPER:
		return []rune{'F', 'A', 'L', 'S', 'E'}
	default:
		return []rune{'f', 'a', 'l', 's', 'e'}
	}
}
