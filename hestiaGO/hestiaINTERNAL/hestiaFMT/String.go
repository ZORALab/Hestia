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

import (
	"unicode"
)

type Lettercase uint8

const (
	LETTERCASE_LOWER Lettercase = iota
	LETTERCASE_UPPER
)

type Notation uint8

const (
	NOTATION_SCIENTIFIC_AUTO = iota
	NOTATION_ISO6093NR3_AUTO
	NOTATION_DECIMALLESS
	NOTATION_SCIENTIFIC
	NOTATION_ISO6093NR3
	NOTATION_DECIMAL_NO_EXPONENT
	NOTATION_HEX
	NOTATION_IEEE754
)

func SN_IsPrintable_RUNE(c rune) bool {
	return unicode.IsPrint(c)
}
