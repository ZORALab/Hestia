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

func _testConclusionScenarios() []*Scenario {
	return []*Scenario{
		{
			Description: `
Test Conclusion() is able to work properly with proper Scenario settings.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test Conclusion() is able to work properly with empty Name setting.
`,
			Switches: []string{
				cond_EMPTY_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test Conclusion() is able to work properly with empty Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_EMPTY_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test Conclusion() is able to work properly with nil Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_NIL_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test Conclusion() is able to work properly with empty log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_EMPTY_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test Conclusion() is able to work properly with nil log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test Conclusion() is able to work properly with empty description setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_EMPTY_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test Conclusion() is able to panic when nil Scenario is supplied.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_SUPPLY_NIL_SCENARIO,
				cond_PROPER_VERDICT,
				expect_PANIC,
			},
		}, {
			Description: `
Test Conclusion() is able to work properly when verict is set to unknown.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_SUPPLY_NIL_SCENARIO,
				cond_UNKNOWN_VERDICT,
				expect_PANIC,
			},
		},
	}
}

func TestConclusionAPI(t *testing.T) {
	scenarios := _testConclusionScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = suite_CONCLUSION_API

		// prepare
		ts := &Scenario{}
		testlib_ConfigureName(s, ts)
		testlib_ConfigureDescription(s, ts)
		testlib_ConfigureLog(s, ts)
		testlib_ConfigureSwitches(s, ts)
		testlib_ConfigureVerdict(s, ts)

		// test
		output := VERDICT_SKIP
		_panick := Exec(func() any {
			if !HasCondition(s, cond_SUPPLY_NIL_SCENARIO) {
				output = Conclusion(ts)
			} else {
				output = Conclusion(nil)
			}
			return ""
		})
		panick := _panick.(string)

		// log output
		Log(s, Format("Test Scenario's Verdict	= %#v", ts.verdict))
		Log(s, Format("Got Output			= %v", output))
		Log(s, Format("Got Panic			= %q", panick))

		// assert
		if !testlib_AssertPanic(s, panick) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if !assert_Conclusion(s, output) {
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

func assert_Conclusion(s *Scenario, output Verdict) bool {
	switch {
	case output == priv_VERDICT_UNKNOWN:
		return HasCondition(s, cond_UNKNOWN_VERDICT)
	case output == VERDICT_PASS:
		return HasCondition(s, cond_PROPER_VERDICT)
	case output == VERDICT_SKIP:
		if !HasCondition(s, cond_PROPER_VERDICT) ||
			HasCondition(s, cond_SUPPLY_NIL_SCENARIO) {
			return true
		}
	}

	return false
}
