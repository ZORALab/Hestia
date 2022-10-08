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
	"testing"

	"hestia/hestiaERROR"
	"hestia/hestiaTESTING"
)

func test_cases_S32_Resize() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
signed 0-bits value.
`,
			Switches: []string{
				cond_TO_BITS_0,
				cond_TO_SIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
unsigned 0-bits value.
`,
			Switches: []string{
				cond_TO_BITS_0,
				cond_TO_UNSIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
signed 5-bits value.
`,
			Switches: []string{
				cond_TO_BITS_5,
				cond_TO_SIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
unsigned 5-bits value.
`,
			Switches: []string{
				cond_TO_BITS_5,
				cond_TO_UNSIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
signed 8-bits value.
`,
			Switches: []string{
				cond_TO_BITS_8,
				cond_TO_SIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
unsigned 8-bits value.
`,
			Switches: []string{
				cond_TO_BITS_8,
				cond_TO_UNSIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
signed 12-bits value.
`,
			Switches: []string{
				cond_TO_BITS_12,
				cond_TO_SIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
unsigned 12-bits value.
`,
			Switches: []string{
				cond_TO_BITS_12,
				cond_TO_UNSIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
signed 16-bits value.
`,
			Switches: []string{
				cond_TO_BITS_16,
				cond_TO_SIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
unsigned 16-bits value.
`,
			Switches: []string{
				cond_TO_BITS_16,
				cond_TO_UNSIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
signed 22-bits value.
`,
			Switches: []string{
				cond_TO_BITS_22,
				cond_TO_SIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
unsigned 22-bits value.
`,
			Switches: []string{
				cond_TO_BITS_22,
				cond_TO_UNSIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
signed 32-bits value.
`,
			Switches: []string{
				cond_TO_BITS_32,
				cond_TO_SIGNED,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S32_Resize is able to process 32-bits value to
unsigned 32-bits value.
`,
			Switches: []string{
				cond_TO_BITS_32,
				cond_TO_UNSIGNED,
			},
		},
	}
}

func Test_S32_Resize(t *testing.T) {
	scenarios := test_cases_S32_Resize()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaNUMBER/hestiaBITS/S32_Resize API"

		// prepare
		var subject = uint32(MAX_UINT32)
		hestiaTESTING.Log(s, _format("Given Subject	: 0b%b", subject))

		size := create_size(s)
		hestiaTESTING.Log(s, _format("Given Size	: %d", size))

		sign := create_sign(s)
		hestiaTESTING.Log(s, _format("Given Sign	: %v", sign))

		// test
		err := S32_Resize(&subject, size, sign)
		hestiaTESTING.Log(s, _format("Got Error	: %d", err))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		if !assert_S32_Resize_output(s, subject) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S32_Resize_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S32_Resize_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_TO_BITS_32),
		hestiaTESTING.HasCondition(s, cond_TO_BITS_22),
		hestiaTESTING.HasCondition(s, cond_TO_BITS_16),
		hestiaTESTING.HasCondition(s, cond_TO_BITS_12),
		hestiaTESTING.HasCondition(s, cond_TO_BITS_8),
		hestiaTESTING.HasCondition(s, cond_TO_BITS_5),
		hestiaTESTING.HasCondition(s, cond_TO_BITS_0):
		return err == hestiaERROR.OK
	case hestiaTESTING.HasCondition(s, cond_TO_BITS_1000):
		return err == hestiaERROR.OUT_OF_RANGE
	}

	return false
}

func assert_S32_Resize_output(s *hestiaTESTING.Scenario, output uint32) bool {
	if hestiaTESTING.HasCondition(s, cond_TO_BITS_1000) {
		return output == uint32(MAX_UINT32)
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_32) {
		if hestiaTESTING.HasCondition(s, cond_TO_SIGNED) {
			return output == uint32(MAX_INT32)
		}

		return output == uint32(MAX_UINT32)
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_22) {
		return output == uint32(value_MASKED_BITS_22)
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_16) {
		if hestiaTESTING.HasCondition(s, cond_TO_SIGNED) {
			return output == uint32(MAX_INT16)
		}

		return output == uint32(MAX_UINT16)
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_12) {
		return output == uint32(value_MASKED_BITS_12)
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_8) {
		if hestiaTESTING.HasCondition(s, cond_TO_SIGNED) {
			return output == uint32(MAX_INT8)
		}

		return output == uint32(MAX_UINT8)
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_5) {
		return output == uint32(value_MASKED_BITS_5)
	}

	if hestiaTESTING.HasCondition(s, cond_TO_BITS_0) {
		return output == 0
	}

	return false
}
