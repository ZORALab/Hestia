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

func _testLogScenarios() []*Scenario {
	return []*Scenario{
		{
			Description: `
Test Log() is able to work properly with proper Scenario settings.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_STRING_STATEMENT,
			},
		}, {
			Description: `
Test Log() is able to work properly with empty Name setting.
`,
			Switches: []string{
				cond_EMPTY_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_STRING_STATEMENT,
			},
		}, {
			Description: `
Test Log() is able to work properly with empty Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_EMPTY_SWITCHES,
				cond_PROPER_STRING_STATEMENT,
			},
		}, {
			Description: `
Test Log() is able to work properly with nil Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_NIL_SWITCHES,
				cond_PROPER_STRING_STATEMENT,
			},
		}, {
			Description: `
Test Log() is able to work properly with empty log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_EMPTY_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_STRING_STATEMENT,
			},
		}, {
			Description: `
Test Log() is able to work properly with nil log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_STRING_STATEMENT,
			},
		}, {
			Description: `
Test Log() is able to work properly with empty description setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_EMPTY_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_STRING_STATEMENT,
			},
		}, {
			Description: `
Test Log() is able to panic when nil Scenario is supplied.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_SUPPLY_NIL_SCENARIO,
				cond_PROPER_STRING_STATEMENT,
				expect_PANIC,
			},
		}, {
			Description: `
Test Log() is able to panic with empty string statement setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_EMPTY_STRING_STATEMENT,
				expect_PANIC,
			},
		},
	}
}

func TestLogAPI(t *testing.T) {
	scenarios := _testLogScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = suite_LOG_API

		// prepare
		ts := &Scenario{}
		testlib_ConfigureName(s, ts)
		testlib_ConfigureDescription(s, ts)
		testlib_ConfigureLog(s, ts)
		testlib_ConfigureSwitches(s, ts)

		statement := testlib_CreateStringStatement(s)

		// test
		_panick := Exec(func() any {
			if !HasCondition(s, cond_SUPPLY_NIL_SCENARIO) {
				Log(ts, statement)
			} else {
				Log(nil, statement)
			}
			return ""
		})
		panick := _panick.(string)

		// log output
		Log(s, _format("Test Scenario's ID		= %#v", ts.ID))
		Log(s, _format("Test Scenario's Name		= %#v", ts.Name))
		Log(s, _format("Test Scenario's Switches	= %#v", ts.Switches))
		Log(s, _format("Test Scenario's Logs		= %#v", ts.Logs))
		Log(s, _format("Got Panic			= %q", panick))

		// assert
		if !testlib_AssertPanic(s, panick) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if !assert_Log(s, ts) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if Conclusion(s) != VERDICT_FAIL {
			Conclude(s, VERDICT_PASS)
		}

		// report
		t.Logf("%s", ToString(s))
	}
}

func assert_Log(s, ts *Scenario) bool {
	if ts == nil {
		return HasCondition(s, cond_SUPPLY_NIL_SCENARIO)
	}
	// ts is now available

	if HasCondition(s, expect_PANIC) {
		switch {
		case HasCondition(s, cond_EMPTY_STRING_STATEMENT),
			HasCondition(s, cond_SUPPLY_NIL_SCENARIO):
			return true
		}
	}
	// no longer in panic expectation.

	if ts.Logs == nil {
		return false
	}
	// ts.Log is now available

	for _, v := range ts.Logs {
		switch v {
		case value_STATEMENT_1:
			return HasCondition(s, cond_PROPER_STRING_STATEMENT)
		default:
			continue
		}
	}

	return false
}
