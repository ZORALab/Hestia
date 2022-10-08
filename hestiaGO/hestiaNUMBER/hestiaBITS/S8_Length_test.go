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

func test_cases_S8_Length() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaNUMBER/hestiaBITS/S8_Length is able to process 0-bits value.
`,
			Switches: []string{
				cond_BITS_0,
			},
		}, {
			Description: `
test hestiaNUMBER/hestiaBITS/S8_Length is able to process 8-bits value.
`,
			Switches: []string{
				cond_BITS_8,
			},
		},
	}
}

func Test_S8_Length(t *testing.T) {
	scenarios := test_cases_S8_Length()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaNUMBER/hestiaBITS/S8_Length API"

		// prepare
		subject := uint8(create_sample(s))
		hestiaTESTING.Log(s, _format("Given Subject	: 0b%b", subject))

		// test
		output := S8_Length(subject)
		hestiaTESTING.Log(s, _format("Got Length	: %d", output))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		if !assert_S8_Length_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S8_Length_output(s *hestiaTESTING.Scenario, output uint8) bool {
	if hestiaTESTING.HasCondition(s, cond_BITS_8) {
		return output == uint8(value_BITS_8_COUNT)
	}

	if hestiaTESTING.HasCondition(s, cond_BITS_0) {
		return output == 0
	}

	return false
}
