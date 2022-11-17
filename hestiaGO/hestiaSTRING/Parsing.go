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

package hestiaSTRING

import (
	"hestia/hestiaERROR"
	"hestia/hestiaINTERNAL/hestiaFMT"
)

func S32_ParseFLOAT32(input string, base float32,
	notation hestiaFMT.Notation) (out float32, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return 0, hestiaERROR.DATA_INVALID
	}

	out, e := hestiaFMT.S32_ParseFLOAT32(input, base, notation)
	switch e {
	case hestiaFMT.ERROR_OK:
		err = hestiaERROR.OK
	case hestiaFMT.ERROR_BASE_INVALID:
		err = hestiaERROR.DATA_MISMATCHED
	case hestiaFMT.ERROR_BASE_MISMATCHED:
		err = hestiaERROR.DATA_MISMATCHED
	case hestiaFMT.ERROR_INPUT_INVALID:
		err = hestiaERROR.INVALID_ARGUMENT
	default:
		err = hestiaERROR.BAD_EXEC
	}

	return out, err
}

func S64_ParseFLOAT64(input string, base float64,
	notation hestiaFMT.Notation) (out float64, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return 0, hestiaERROR.DATA_INVALID
	}

	out, e := hestiaFMT.S64_ParseFLOAT64(input, base, notation)
	switch e {
	case hestiaFMT.ERROR_OK:
		err = hestiaERROR.OK
	case hestiaFMT.ERROR_BASE_INVALID:
		err = hestiaERROR.DATA_MISMATCHED
	case hestiaFMT.ERROR_BASE_MISMATCHED:
		err = hestiaERROR.DATA_MISMATCHED
	case hestiaFMT.ERROR_INPUT_INVALID:
		err = hestiaERROR.INVALID_ARGUMENT
	default:
		err = hestiaERROR.BAD_EXEC
	}

	return out, err
}
