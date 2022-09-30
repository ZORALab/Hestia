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
	"hestia/hestiaTESTING"
)

// test suites label
const (
	suite_LEN_API           = "hestiaBITS Len(...) API"
	suite_TRAILING_ZERO_API = "hestiaBITS TrailingZero(...) API"
	suite_MATCH_SIZE_API    = "hestiaBITS MatchSize(...) API"
)

// test switches
const (
	cond_GENERATE_UINT_64_BITS = "generate 64-bit uint number"
	cond_GENERATE_UINT_52_BITS = "generate 52-bit uint number"
	cond_GENERATE_UINT_32_BITS = "generate 32-bit uint number"
	cond_GENERATE_UINT_22_BITS = "generate 22-bit uint number"
	cond_GENERATE_UINT_16_BITS = "generate 16-bit uint number"
	cond_GENERATE_UINT_12_BITS = "generate 12-bit uint number"
	cond_GENERATE_UINT_8_BITS  = "generate 8-bit uint number"
	cond_GENERATE_UINT_5_BITS  = "generate 5-bit uint number"

	cond_SEEK_1024_BITS = "seek 1024-bits sized number"
	cond_SEEK_64_BITS   = "seek 64-bits sized number"
	cond_SEEK_52_BITS   = "seek 52-bits sized number"
	cond_SEEK_32_BITS   = "seek 32-bits sized number"
	cond_SEEK_22_BITS   = "seek 22-bits sized number"
	cond_SEEK_16_BITS   = "seek 16-bits sized number"
	cond_SEEK_12_BITS   = "seek 12-bits sized number"
	cond_SEEK_8_BITS    = "seek 8-bits sized number"
	cond_SEEK_5_BITS    = "seek 5-bits sized number"

	cond_SEEK_POSITIVE = "seek positive number"
	cond_SEEK_NEGATIVE = "seek negative number"
)

// test values
const (
	value_64_BITS      = 64
	value_UINT_64_BITS = 1<<value_64_BITS - 1

	value_52_BITS      = 52
	value_UINT_52_BITS = 1<<value_52_BITS - 1

	value_32_BITS      = 32
	value_UINT_32_BITS = 1<<value_32_BITS - 1

	value_22_BITS      = 22
	value_UINT_22_BITS = 1<<value_22_BITS - 1

	value_16_BITS      = 16
	value_UINT_16_BITS = 1<<value_16_BITS - 1

	value_12_BITS      = 12
	value_UINT_12_BITS = 1<<value_12_BITS - 1

	value_8_BITS      = 8
	value_UINT_8_BITS = 1<<value_8_BITS - 1

	value_5_BITS      = 5
	value_UINT_5_BITS = 1<<value_5_BITS - 1
)

func testlibs_AssertMatchSize(s *hestiaTESTING.Scenario,
	output uint64, err hestiaERROR.Error) bool {
	switch output {
	case value_UINT_5_BITS:
		switch {
		case s.Switches[cond_GENERATE_UINT_5_BITS] && s.Switches[cond_SEEK_1024_BITS]:
			return err != hestiaERROR.OK
		case s.Switches[cond_GENERATE_UINT_5_BITS]:
			return true
		}
	case hestiaNUMBER.MAX_INT8:
		if !s.Switches[cond_SEEK_NEGATIVE] {
			return false
		}

		fallthrough
	case hestiaNUMBER.MAX_UINT8:
		switch {
		case s.Switches[cond_SEEK_5_BITS], s.Switches[cond_SEEK_8_BITS]:
			return true
		}
	case hestiaNUMBER.MAX_INT16:
		if !s.Switches[cond_SEEK_NEGATIVE] {
			return false
		}

		fallthrough
	case hestiaNUMBER.MAX_UINT16:
		switch {
		case s.Switches[cond_GENERATE_UINT_64_BITS] && s.Switches[cond_SEEK_12_BITS],
			s.Switches[cond_GENERATE_UINT_64_BITS] && s.Switches[cond_SEEK_16_BITS]:
			return true
		}
	case hestiaNUMBER.MAX_INT32:
		if !s.Switches[cond_SEEK_NEGATIVE] {
			return false
		}

		fallthrough
	case hestiaNUMBER.MAX_UINT32:
		switch {
		case s.Switches[cond_GENERATE_UINT_64_BITS] && s.Switches[cond_SEEK_22_BITS],
			s.Switches[cond_GENERATE_UINT_64_BITS] && s.Switches[cond_SEEK_32_BITS]:
			return true
		}
	case hestiaNUMBER.MAX_INT64:
		if !s.Switches[cond_SEEK_NEGATIVE] {
			return false
		}

		fallthrough
	case hestiaNUMBER.MAX_UINT64:
		switch {
		case s.Switches[cond_GENERATE_UINT_64_BITS] && s.Switches[cond_SEEK_1024_BITS]:
			return err != hestiaERROR.OK
		default:
			return s.Switches[cond_SEEK_64_BITS]
		}
	}

	return false
}

func testlibs_GenerateTargetSign(s *hestiaTESTING.Scenario) bool {
	switch {
	case s.Switches[cond_SEEK_NEGATIVE]:
		return false
	default:
		return true
	}
}

func testlibs_GenerateTargetSize(s *hestiaTESTING.Scenario) uint16 {
	switch {
	case s.Switches[cond_SEEK_1024_BITS]:
		return 1024
	case s.Switches[cond_SEEK_64_BITS]:
		return 64
	case s.Switches[cond_SEEK_52_BITS]:
		return 52
	case s.Switches[cond_SEEK_32_BITS]:
		return 32
	case s.Switches[cond_SEEK_22_BITS]:
		return 22
	case s.Switches[cond_SEEK_16_BITS]:
		return 16
	case s.Switches[cond_SEEK_12_BITS]:
		return 12
	case s.Switches[cond_SEEK_8_BITS]:
		return 8
	case s.Switches[cond_SEEK_5_BITS]:
		return 5
	default:
		return 0
	}
}

func testlibs_AssertTrailingZero(s *hestiaTESTING.Scenario, output uint) bool {
	switch {
	case s.Switches[cond_GENERATE_UINT_64_BITS]:
		return output == value_64_BITS-1
	case s.Switches[cond_GENERATE_UINT_52_BITS]:
		return output == value_52_BITS-1
	case s.Switches[cond_GENERATE_UINT_32_BITS]:
		return output == value_32_BITS-1
	case s.Switches[cond_GENERATE_UINT_22_BITS]:
		return output == value_22_BITS-1
	case s.Switches[cond_GENERATE_UINT_16_BITS]:
		return output == value_16_BITS-1
	case s.Switches[cond_GENERATE_UINT_12_BITS]:
		return output == value_12_BITS-1
	case s.Switches[cond_GENERATE_UINT_8_BITS]:
		return output == value_8_BITS-1
	case s.Switches[cond_GENERATE_UINT_5_BITS]:
		return output == value_5_BITS-1
	default:
		return false
	}
}

func testlibs_GenerateTrailingZero(s *hestiaTESTING.Scenario) uint {
	switch {
	case s.Switches[cond_GENERATE_UINT_64_BITS]:
		return 1 << (value_64_BITS - 1)
	case s.Switches[cond_GENERATE_UINT_52_BITS]:
		return 1 << (value_52_BITS - 1)
	case s.Switches[cond_GENERATE_UINT_32_BITS]:
		return 1 << (value_32_BITS - 1)
	case s.Switches[cond_GENERATE_UINT_22_BITS]:
		return 1 << (value_22_BITS - 1)
	case s.Switches[cond_GENERATE_UINT_16_BITS]:
		return 1 << (value_16_BITS - 1)
	case s.Switches[cond_GENERATE_UINT_12_BITS]:
		return 1 << (value_12_BITS - 1)
	case s.Switches[cond_GENERATE_UINT_8_BITS]:
		return 1 << (value_8_BITS - 1)
	case s.Switches[cond_GENERATE_UINT_5_BITS]:
		return 1 << (value_5_BITS - 1)
	default:
		return 1
	}
}

func testlibs_assertLen(s *hestiaTESTING.Scenario, output uint) bool {
	switch {
	case s.Switches[cond_GENERATE_UINT_64_BITS]:
		return output == value_64_BITS
	case s.Switches[cond_GENERATE_UINT_52_BITS]:
		return output == value_52_BITS
	case s.Switches[cond_GENERATE_UINT_32_BITS]:
		return output == value_32_BITS
	case s.Switches[cond_GENERATE_UINT_22_BITS]:
		return output == value_22_BITS
	case s.Switches[cond_GENERATE_UINT_16_BITS]:
		return output == value_16_BITS
	case s.Switches[cond_GENERATE_UINT_12_BITS]:
		return output == value_12_BITS
	case s.Switches[cond_GENERATE_UINT_8_BITS]:
		return output == value_8_BITS
	case s.Switches[cond_GENERATE_UINT_5_BITS]:
		return output == value_5_BITS
	default:
		return false
	}
}

func testlibs_assertTestableHost(s *hestiaTESTING.Scenario) bool {
	limit := hestiaNUMBER.CPU()
	switch {
	case s.Switches[cond_GENERATE_UINT_64_BITS]:
		return limit >= value_64_BITS
	case s.Switches[cond_GENERATE_UINT_52_BITS]:
		return limit >= value_64_BITS
	case s.Switches[cond_GENERATE_UINT_32_BITS]:
		return limit >= value_32_BITS
	case s.Switches[cond_GENERATE_UINT_22_BITS]:
		return limit >= value_32_BITS
	case s.Switches[cond_GENERATE_UINT_16_BITS]:
		return limit >= value_16_BITS
	case s.Switches[cond_GENERATE_UINT_12_BITS]:
		return limit >= value_16_BITS
	case s.Switches[cond_GENERATE_UINT_8_BITS]:
		return limit >= value_8_BITS
	case s.Switches[cond_GENERATE_UINT_5_BITS]:
		return limit >= value_8_BITS
	}

	return false
}

func testlibs_GenerateUINT(s *hestiaTESTING.Scenario) uint {
	switch {
	case s.Switches[cond_GENERATE_UINT_64_BITS]:
		return value_UINT_64_BITS
	case s.Switches[cond_GENERATE_UINT_52_BITS]:
		return value_UINT_52_BITS
	case s.Switches[cond_GENERATE_UINT_32_BITS]:
		return value_UINT_32_BITS
	case s.Switches[cond_GENERATE_UINT_22_BITS]:
		return value_UINT_22_BITS
	case s.Switches[cond_GENERATE_UINT_16_BITS]:
		return value_UINT_16_BITS
	case s.Switches[cond_GENERATE_UINT_12_BITS]:
		return value_UINT_12_BITS
	case s.Switches[cond_GENERATE_UINT_8_BITS]:
		return value_UINT_8_BITS
	case s.Switches[cond_GENERATE_UINT_5_BITS]:
		return value_UINT_5_BITS
	default:
		return 0
	}
}
