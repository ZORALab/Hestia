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

type Lettercase uint8

const (
	LETTERCASE_LOWER Lettercase = iota
	LETTERCASE_UPPER
)

func _parseChar(arg any) (out []rune) {
	// to be developed later
	return []rune{'(', 'C', 'H', 'A', 'R', '=', 'b', 'a', 'd', ')'}
}

func _parseString(arg any) (out []rune) {
	// to be developed later
	return []rune{'(', 'S', 'T', 'R', 'I', 'N', 'G', '=', 'b', 'a', 'd', ')'}
}
