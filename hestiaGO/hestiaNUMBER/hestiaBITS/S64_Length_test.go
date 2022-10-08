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

	"hestia/hestiaTESTING"
)

func test_cases_S64_Length() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaNUMBER/hestiaBITS/S64_Length is able to process 0-bits value.
`,
			Switches: []string{
				cond_BITS_0,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S64_Length is able to process 8-bits value.
`,
			Switches: []string{
				cond_BITS_8,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S64_Length is able to process 16-bits value.
`,
			Switches: []string{
				cond_BITS_16,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S64_Length is able to process 32-bits value.
`,
			Switches: []string{
				cond_BITS_32,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S64_Length is able to process 64-bits value.
`,
			Switches: []string{
				cond_BITS_64,
			},
		},
	}
}

func Test_S64_Length(t *testing.T) {
	scenarios := test_cases_S64_Length()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaNUMBER/hestiaBITS/S64_Length API"

		// prepare
		subject := create_sample(s)
		hestiaTESTING.Log(s, _format("Given Subject	: 0b%b", subject))

		// test
		output := S64_Length(subject)
		hestiaTESTING.Log(s, _format("Got Length	: %d", output))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		if !assert_S64_Length_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S64_Length_output(s *hestiaTESTING.Scenario, output uint64) bool {
	if hestiaTESTING.HasCondition(s, cond_BITS_64) {
		return output == value_BITS_64_COUNT
	}

	if hestiaTESTING.HasCondition(s, cond_BITS_32) {
		return output == value_BITS_32_COUNT
	}

	if hestiaTESTING.HasCondition(s, cond_BITS_16) {
		return output == value_BITS_16_COUNT
	}

	if hestiaTESTING.HasCondition(s, cond_BITS_8) {
		return output == value_BITS_8_COUNT
	}

	if hestiaTESTING.HasCondition(s, cond_BITS_0) {
		return output == 0
	}

	return false
}
