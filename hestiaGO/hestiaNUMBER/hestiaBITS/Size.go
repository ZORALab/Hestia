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
	"hestia/hestiaNUMBER"
)

func MatchSize(input *uint64, size uint16, withSign bool) hestiaERROR.Error {
	var mask uint64

	switch {
	case input == nil:
		return hestiaERROR.INVALID_ARGUMENT
	case size > 64:
		return hestiaERROR.OUT_OF_RANGE
	case size == 0:
		*input = 0
		return hestiaERROR.OK
	case size <= 8:
		if withSign {
			mask = hestiaNUMBER.MAX_INT8
		} else {
			mask = hestiaNUMBER.MAX_UINT8
		}
	case size <= 16:
		if withSign {
			mask = hestiaNUMBER.MAX_INT16
		} else {
			mask = hestiaNUMBER.MAX_UINT16
		}
	case size <= 32:
		if withSign {
			mask = hestiaNUMBER.MAX_INT32
		} else {
			mask = hestiaNUMBER.MAX_UINT32
		}
	case size <= 64:
		if withSign {
			mask = hestiaNUMBER.MAX_INT64
		} else {
			mask = hestiaNUMBER.MAX_UINT64
		}
	}

	*input &= mask

	return hestiaERROR.OK
}
