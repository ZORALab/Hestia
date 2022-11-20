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

package hestiaBITS

import (
	"hestia/hestiaERROR"
	"hestia/hestiaINTERNAL/hestiaMATH"
)

func S64_Resize(input *uint64, size uint16, withSign bool) hestiaERROR.Error {
	switch hestiaMATH.S64_Resize(input, size, withSign) {
	case hestiaMATH.ERROR_OK:
		return hestiaERROR.OK
	case hestiaMATH.ERROR_INPUT_INVALID:
		return hestiaERROR.INVALID_ARGUMENT
	case hestiaMATH.ERROR_INPUT_OUT_OF_RANGE:
		return hestiaERROR.OUT_OF_RANGE
	default:
		return hestiaERROR.BAD_EXEC
	}
}

func S32_Resize(input *uint32, size uint16, withSign bool) hestiaERROR.Error {
	switch hestiaMATH.S32_Resize(input, size, withSign) {
	case hestiaMATH.ERROR_OK:
		return hestiaERROR.OK
	case hestiaMATH.ERROR_INPUT_INVALID:
		return hestiaERROR.INVALID_ARGUMENT
	case hestiaMATH.ERROR_INPUT_OUT_OF_RANGE:
		return hestiaERROR.OUT_OF_RANGE
	default:
		return hestiaERROR.BAD_EXEC
	}
}

func S16_Resize(input *uint16, size uint16, withSign bool) hestiaERROR.Error {
	switch hestiaMATH.S16_Resize(input, size, withSign) {
	case hestiaMATH.ERROR_OK:
		return hestiaERROR.OK
	case hestiaMATH.ERROR_INPUT_INVALID:
		return hestiaERROR.INVALID_ARGUMENT
	case hestiaMATH.ERROR_INPUT_OUT_OF_RANGE:
		return hestiaERROR.OUT_OF_RANGE
	default:
		return hestiaERROR.BAD_EXEC
	}
}

func S8_Resize(input *uint8, size uint16, withSign bool) hestiaERROR.Error {
	switch hestiaMATH.S8_Resize(input, size, withSign) {
	case hestiaMATH.ERROR_OK:
		return hestiaERROR.OK
	case hestiaMATH.ERROR_INPUT_INVALID:
		return hestiaERROR.INVALID_ARGUMENT
	case hestiaMATH.ERROR_INPUT_OUT_OF_RANGE:
		return hestiaERROR.OUT_OF_RANGE
	default:
		return hestiaERROR.BAD_EXEC
	}
}
