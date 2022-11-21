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
	"hestia/hestiaERROR"
	"hestia/hestiaTESTING"
	"testing"
)

func test_cases_SN_Atoi() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_BROKEN,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_BROKEN,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_2,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_2,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_2,
				cond_BROKEN,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_2,
				cond_BROKEN,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_8,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_8,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_8,
				cond_BROKEN,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_8,
				cond_BROKEN,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_10,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_10,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_10,
				cond_BROKEN,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_10,
				cond_BROKEN,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_16,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_16,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_16,
				cond_BROKEN,
				cond_POSITIVE,
				cond_BASE_10,
			},
		}, {
			Switches: []string{
				cond_VALUE_NUMBER,
				cond_RESIZE_16,
				cond_BROKEN,
				cond_NEGATIVE,
				cond_BASE_10,
			},
		},
	}
}

func Test_SN_Atoi(t *testing.T) {
	scenarios := test_cases_SN_Atoi()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/SN_Atoi API"
		s.Description = `
test hestiaSTRING/SN_Atoi is able to process the given string under the
following conditions.
`

		// prepare
		subject := generate_INT_string(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: '%s'", subject))

		sizes := create_sizes(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Sizes	: %d", sizes))

		// test
		var output int64
		panick := ""
		err := hestiaERROR.OK

		_panick := hestiaTESTING.Exec(func() any {
			output, err = SN_Atoi(subject, sizes)
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: %v", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Error	: %v", err))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		if !assert_SN_ParseINT_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, hestiaTESTING.Format("Failed by panick!"))
			t.Fail()
		}

		if !assert_SN_ParseINT_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, hestiaTESTING.Format("Failed by output!"))
			t.Fail()
		}

		if !assert_SN_ParseINT_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, hestiaTESTING.Format("Failed by error!"))
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}
