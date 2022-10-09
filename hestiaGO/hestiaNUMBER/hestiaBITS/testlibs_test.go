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
	"fmt"

	"hestia/hestiaTESTING"
)

// test conditions
const (
	cond_BITS_64 = "provide 64-bits value"
	cond_BITS_32 = "provide 32-bits value"
	cond_BITS_16 = "provide 16-bits value"
	cond_BITS_8  = "provide 8-bits value"
	cond_BITS_0  = "provide 0-bit value"

	cond_TO_BITS_1000 = "convert to 1000-bits value"
	cond_TO_BITS_64   = "convert to 64-bits value"
	cond_TO_BITS_35   = "convert to 35-bits value"
	cond_TO_BITS_32   = "convert to 32-bits value"
	cond_TO_BITS_22   = "convert to 22-bits value"
	cond_TO_BITS_16   = "convert to 16-bits value"
	cond_TO_BITS_12   = "convert to 12-bits value"
	cond_TO_BITS_8    = "convert to 8-bits value"
	cond_TO_BITS_5    = "convert to 5-bits value"
	cond_TO_BITS_0    = "convert to 0-bit value"

	cond_TO_UNSIGNED = "convert to unsigned value"
	cond_TO_SIGNED   = "convert to signed value"

	cond_NIL_INPUT = "provide nil input"
)

// test values
const (
	value_MASKED_BITS_35 = 0x7FFFFFFFF
	value_MASKED_BITS_22 = 0x3FFFFF
	value_MASKED_BITS_12 = 0xFFF
	value_MASKED_BITS_5  = 0x1F

	value_BITS_64_COUNT = 64
	value_BITS_32_COUNT = 32
	value_BITS_16_COUNT = 16
	value_BITS_8_COUNT  = 8
)

func create_sign(s *hestiaTESTING.Scenario) bool {
	if hestiaTESTING.HasCondition(s, cond_TO_SIGNED) {
		return true
	}

	if hestiaTESTING.HasCondition(s, cond_TO_UNSIGNED) {
		return false
	}

	return false
}

func create_size(s *hestiaTESTING.Scenario) uint16 {
	if hestiaTESTING.HasCondition(s, cond_TO_BITS_1000) {
		return 1000
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_64) {
		return 64
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_35) {
		return 35
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_32) {
		return 32
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_22) {
		return 22
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_16) {
		return 16
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_12) {
		return 12
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_8) {
		return 8
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_5) {
		return 5
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_0) {
		return 0
	}

	return 0
}

func create_sample(s *hestiaTESTING.Scenario) uint64 {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BITS_64):
		return 1 << (64 - 1)
	case hestiaTESTING.HasCondition(s, cond_BITS_32):
		return 1 << (32 - 1)
	case hestiaTESTING.HasCondition(s, cond_BITS_16):
		return 1 << (16 - 1)
	case hestiaTESTING.HasCondition(s, cond_BITS_8):
		return 1 << (8 - 1)
	case hestiaTESTING.HasCondition(s, cond_BITS_0):
		return 0
	default:
	}

	return 0
}

func _format(f string, args ...any) string {
	return fmt.Sprintf(f, args...)
}
