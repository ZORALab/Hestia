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

type lettercase uint

const (
	lettercase_LOWER lettercase = iota
	lettercase_UPPER
)

type numberType uint

const (
	format_DECIMALLESS              numberType = iota // (e.g. 123)
	format_SCIENTIFIC                                 // (e.g. -1.234456e+78)
	format_SCIENTIFIC_AUTO_EXPONENT                   // (e.g. -1.234 OR -1.234456e+78)
	format_DECIMAL_NO_EXPONENT                        // (e.g -123.456)
	format_HEX                                        // (e.g. -0x1.23abcp+20)
)

type numberSTR struct {
	arg       any
	base      uint16
	charset   lettercase
	width     []rune
	precision []rune
	format    numberType
}

func _parseNumber(str *numberSTR) (out []rune) {
	// to be developed later
	return []rune{'(', 'N', 'U', 'M', 'B', 'E', 'R', '=', 'b', 'a', 'd', ')'}
}
