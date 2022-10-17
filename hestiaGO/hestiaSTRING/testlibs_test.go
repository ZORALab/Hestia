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
	"hestia/hestiaINTERNAL/hestiaFMT"
	"hestia/hestiaTESTING"
)

// test condition labels
const (
	cond_BOOL_TRUE  = "format boolean 'true' value"
	cond_BOOL_FALSE = "format boolean 'false' value"

	cond_BASE_0  = "format value into Base-0 number"
	cond_BASE_1  = "format value into Base-1 number"
	cond_BASE_2  = "format value into Base-2 number"
	cond_BASE_5  = "format value into Base-5 number"
	cond_BASE_8  = "format value into Base-8 number"
	cond_BASE_10 = "format value into Base-10 number"
	cond_BASE_12 = "format value into Base-12 number"
	cond_BASE_16 = "format value into Base-16 number"
	cond_BASE_22 = "format value into Base-22 number"
	cond_BASE_36 = "format value into Base-36 number"
	cond_BASE_37 = "format value into Base-37 number"

	cond_VALUE_UINT64 = "provide max uint64 as sample value"
	cond_VALUE_INT64  = "provide min/max int64 as sample value"
	cond_VALUE_UINT32 = "provide max uint32 as sample value"
	cond_VALUE_INT32  = "provide min/max int32 as sample value"
	cond_VALUE_UINT16 = "provide max uint16 as sample value"
	cond_VALUE_UINT8  = "provide max uint8 as sample value"
	cond_VALUE_ZERO   = "provide 0 as sample value"

	cond_POSITIVE = "provide positive value"
	cond_NEGATIVE = "provide positive negative"

	cond_UNKNOWNCASE = "formate value into unknown case characters"
	cond_LOWERCASE   = "formate value into lowercase characters"
	cond_UPPERCASE   = "formate value into uppercase characters"
)

func create_lettercase(s *hestiaTESTING.Scenario) hestiaFMT.Lettercase {
	switch {
	case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
		return LETTERCASE_UPPER
	case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
		return LETTERCASE_LOWER
	}

	return 100
}

func create_base(s *hestiaTESTING.Scenario) uint64 {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0):
		return 0
	case hestiaTESTING.HasCondition(s, cond_BASE_1):
		return 1
	case hestiaTESTING.HasCondition(s, cond_BASE_2):
		return 2
	case hestiaTESTING.HasCondition(s, cond_BASE_5):
		return 5
	case hestiaTESTING.HasCondition(s, cond_BASE_8):
		return 8
	case hestiaTESTING.HasCondition(s, cond_BASE_10):
		return 10
	case hestiaTESTING.HasCondition(s, cond_BASE_12):
		return 12
	case hestiaTESTING.HasCondition(s, cond_BASE_16):
		return 16
	case hestiaTESTING.HasCondition(s, cond_BASE_22):
		return 22
	case hestiaTESTING.HasCondition(s, cond_BASE_36):
		return 36
	case hestiaTESTING.HasCondition(s, cond_BASE_37):
		return 37
	default:
		return 100
	}
}
