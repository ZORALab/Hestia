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

package hestiaTESTING

import (
	"testing"
)

func _testToTOMLScenarios() []*Scenario {
	return []*Scenario{
		{
			Description: `
Test ToTOML() is able to work properly with proper Scenario settings.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				expect_OUTPUT_STRING,
			},
		}, {
			Description: `
Test ToTOML() is able to work properly with empty Name setting.
`,
			Switches: []string{
				cond_EMPTY_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				expect_OUTPUT_STRING,
			},
		}, {
			Description: `
Test ToTOML() is able to work properly with empty Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_EMPTY_SWITCHES,
				expect_OUTPUT_STRING,
			},
		}, {
			Description: `
Test ToTOML() is able to work properly with nil Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_NIL_SWITCHES,
				expect_OUTPUT_STRING,
			},
		}, {
			Description: `
Test ToTOML() is able to work properly with empty log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_EMPTY_LOG,
				cond_PROPER_SWITCHES,
				expect_OUTPUT_STRING,
			},
		}, {
			Description: `
Test ToTOML() is able to work properly with nil log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				expect_OUTPUT_STRING,
			},
		}, {
			Description: `
Test ToTOML() is able to work properly with empty description setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_EMPTY_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				expect_OUTPUT_STRING,
			},
		}, {
			Description: `
Test ToTOML() is able to panic when nil Scenario is supplied.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_SUPPLY_NIL_SCENARIO,
				expect_PANIC,
			},
		}, {
			Description: `
Test ToTOML() is able to work properly with empty Switches and empty Log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_EMPTY_LOG,
				cond_EMPTY_SWITCHES,
				expect_OUTPUT_STRING,
			},
		},
	}
}

func TestToTOMLAPI(t *testing.T) {
	scenarios := _testToTOMLScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = suite_TO_TOML_API

		// prepare
		ts := &Scenario{}
		testlib_ConfigureName(s, ts)
		testlib_ConfigureDescription(s, ts)
		testlib_ConfigureLog(s, ts)
		testlib_ConfigureSwitches(s, ts)

		// test
		output := ""
		_panick := Exec(func() any {
			if !HasCondition(s, cond_SUPPLY_NIL_SCENARIO) {
				output = ToTOML(ts)
			} else {
				output = ToTOML(nil)
			}
			return ""
		})
		panick := _panick.(string)

		// log output
		Log(s, _renderString("Test Scenario's ID		= %#v", ts.ID))
		Log(s, _renderString("Test Scenario's Name		= %#v", ts.Name))
		Log(s, _renderString("Test Scenario's Switches	= %#v", ts.Switches))
		Log(s, _renderString("Test Scenario's Log		= %#v", ts.Logs))
		Log(s, _renderString("Got Output			= %q", output))
		Log(s, _renderString("Got Panic			= %q", panick))

		// assert
		if !testlib_AssertPanic(s, panick) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if !testlib_AssertOutputString(s, output) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if Conclusion(s) != VERDICT_FAIL {
			Conclude(s, VERDICT_PASS)
		}

		// report
		t.Logf("\n%s\n\n\n", ToTOML(s))
	}
}
