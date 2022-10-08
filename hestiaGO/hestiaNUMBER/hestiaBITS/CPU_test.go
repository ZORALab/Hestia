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

func test_cases_CPU() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaNUMBER/hestiaBITS/CPU is able to return value.
`,
			Switches: []string{},
		},
	}
}

func Test_CPU(t *testing.T) {
	scenarios := test_cases_CPU()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaNUMBER/hestiaBITS/CPU API"

		// prepare

		// test
		output := CPU()
		hestiaTESTING.Log(s, _format("Got Output	: %d", output))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		if !assert_CPU_output(output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_CPU_output(output uint64) bool {
	return output != 0
}
