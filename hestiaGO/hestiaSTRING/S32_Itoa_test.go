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
	"hestia/hestiaNUMBER/hestiaBITS"
	"hestia/hestiaTESTING"
	"testing"
)

func test_cases_S32_Itoa() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaSTRING/S32_Itoa is able to process lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT32,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S32_Itoa is able to process uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT32,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S32_Itoa is able to process lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT32,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S32_Itoa is able to process uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT32,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S32_Itoa is able to process lowercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S32_Itoa is able to process uppercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S32_Itoa is able to process lowercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S32_Itoa is able to process uppercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		},
	}
}

func Test_S32_Itoa(t *testing.T) {
	scenarios := test_cases_S32_Itoa()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/S32_Itoa API"

		// prepare
		subject := int32(hestiaBITS.MAX_INT32)
		switch {
		case hestiaTESTING.HasCondition(s, cond_VALUE_ZERO):
			subject = 0
		case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
			subject = int32(hestiaBITS.MIN_INT32)
		}
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: %d", subject))

		// test
		var output string
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			output = S32_Itoa(subject)
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: '%s'", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		if !assert_S32_Itoa_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S32_Itoa_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S32_Itoa_output(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_VALUE_ZERO):
		return output == "0"
	case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
		return output == "-2147483648"
	case hestiaTESTING.HasCondition(s, cond_POSITIVE):
		return output == "2147483647"
	default:
		return false
	}
}

func assert_S32_Itoa_panick(panick string) bool {
	return panick == ""
}
