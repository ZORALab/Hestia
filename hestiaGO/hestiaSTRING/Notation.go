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

package hestiaSTRING

import (
	"hestia/hestiaERROR"
	"hestia/hestiaINTERNAL/hestiaFMT"
)

const (
	NOTATION_SCIENTIFIC_AUTO     = hestiaFMT.NOTATION_SCIENTIFIC_AUTO
	NOTATION_ISO6093NR3_AUTO     = hestiaFMT.NOTATION_ISO6093NR3_AUTO
	NOTATION_ISO6093NR3          = hestiaFMT.NOTATION_ISO6093NR3
	NOTATION_SCIENTIFIC          = hestiaFMT.NOTATION_SCIENTIFIC
	NOTATION_DECIMAL_NO_EXPONENT = hestiaFMT.NOTATION_DECIMAL_NO_EXPONENT
	NOTATION_IEEE754             = hestiaFMT.NOTATION_IEEE754
)

func _processNotation(notation *hestiaFMT.Notation) hestiaERROR.Error {
	switch *notation {
	case NOTATION_SCIENTIFIC,
		NOTATION_SCIENTIFIC_AUTO,
		NOTATION_DECIMAL_NO_EXPONENT,
		NOTATION_IEEE754,
		NOTATION_ISO6093NR3_AUTO,
		NOTATION_ISO6093NR3:
		return hestiaERROR.OK
	default:
		return hestiaERROR.DATA_INVALID
	}
}
