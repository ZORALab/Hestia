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
	"hestia/hestiaTESTING"
	"testing"
)

func test_cases_SN_FormatBOOL() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Switches: []string{
				cond_BOOL_TRUE,
				cond_LOWERCASE,
			},
		}, {
			Switches: []string{
				cond_BOOL_TRUE,
				cond_UPPERCASE,
			},
		}, {
			Switches: []string{
				cond_BOOL_TRUE,
				cond_UNKNOWNCASE,
			},
		}, {
			Switches: []string{
				cond_BOOL_FALSE,
				cond_LOWERCASE,
			},
		}, {
			Switches: []string{
				cond_BOOL_FALSE,
				cond_UPPERCASE,
			},
		}, {
			Switches: []string{
				cond_BOOL_FALSE,
				cond_UNKNOWNCASE,
			},
		},
	}
}

func Test_SN_FormatBOOL(t *testing.T) {
	scenarios := test_cases_SN_FormatBOOL()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/SN_FormatBOOL API"
		s.Description = `
test hestiaSTRING/SN_FormatBOOL is able to process values with the following
switches.
`

		// prepare
		subject := false
		if hestiaTESTING.HasCondition(s, cond_BOOL_TRUE) {
			subject = true
		}

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: %t", subject))

		lettercase := create_lettercase(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Lettercase	: %d", lettercase))

		// test
		var output string
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			output = SN_FormatBOOL(subject, lettercase)
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: '%s'", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		if !assert_SN_FormatBOOL_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_SN_FormatBOOL_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_SN_FormatBOOL_output(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BOOL_TRUE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			return output == "TRUE"
		default:
			return output == "true"
		}
	case hestiaTESTING.HasCondition(s, cond_BOOL_FALSE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			return output == "FALSE"
		default:
			return output == "false"
		}
	}

	return false
}

func assert_SN_FormatBOOL_panick(panick string) bool {
	return panick == ""
}
