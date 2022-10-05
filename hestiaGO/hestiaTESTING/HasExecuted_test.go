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

func _testHasExecutedScenarios() []*Scenario {
	return []*Scenario{
		{
			Description: `
Test HasExecuted() is able to work properly with proper Scenario settings.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:        true,
				cond_PROPER_DESCRIPTION: true,
				cond_PROPER_LOG:         true,
				cond_PROPER_SWITCHES:    true,
				cond_PROPER_VERDICT:     true,
				expect_PANIC:            false,
			},
		}, {
			Description: `
Test HasExecuted() is able to work properly with empty Name setting.
`,
			Switches: map[string]bool{
				cond_EMPTY_NAME:         true,
				cond_PROPER_DESCRIPTION: true,
				cond_PROPER_LOG:         true,
				cond_PROPER_SWITCHES:    true,
				cond_PROPER_VERDICT:     true,
				expect_PANIC:            false,
			},
		}, {
			Description: `
Test HasExecuted() is able to work properly with empty Switches setting.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:        true,
				cond_PROPER_DESCRIPTION: true,
				cond_PROPER_LOG:         true,
				cond_EMPTY_SWITCHES:     true,
				cond_PROPER_VERDICT:     true,
				expect_PANIC:            false,
			},
		}, {
			Description: `
Test HasExecuted() is able to work properly with nil Switches setting.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:        true,
				cond_PROPER_DESCRIPTION: true,
				cond_PROPER_LOG:         true,
				cond_NIL_SWITCHES:       true,
				cond_PROPER_VERDICT:     true,
				expect_PANIC:            false,
			},
		}, {
			Description: `
Test HasExecuted() is able to work properly with empty log setting.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:        true,
				cond_PROPER_DESCRIPTION: true,
				cond_EMPTY_LOG:          true,
				cond_PROPER_SWITCHES:    true,
				cond_PROPER_VERDICT:     true,
				expect_PANIC:            false,
			},
		}, {
			Description: `
Test HasExecuted() is able to work properly with nil log setting.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:        true,
				cond_PROPER_DESCRIPTION: true,
				cond_NIL_LOG:            true,
				cond_PROPER_SWITCHES:    true,
				cond_PROPER_VERDICT:     true,
				expect_PANIC:            false,
			},
		}, {
			Description: `
Test HasExecuted() is able to work properly with empty description setting.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:       true,
				cond_EMPTY_DESCRIPTION: true,
				cond_NIL_LOG:           true,
				cond_PROPER_SWITCHES:   true,
				cond_PROPER_VERDICT:    true,
				expect_PANIC:           false,
			},
		}, {
			Description: `
Test HasExecuted() is able to panic when nil Scenario is supplied.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_SUPPLY_NIL_SCENARIO: true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             true,
			},
		}, {
			Description: `
Test HasExecuted() is able to work properly when verdict is set to unknown.
`,
			Switches: map[string]bool{
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_SUPPLY_NIL_SCENARIO: true,
				cond_UNKNOWN_VERDICT:     true,
				expect_PANIC:             true,
			},
		},
	}
}

func TestHasExecutedAPI(t *testing.T) {
	scenarios := _testHasExecutedScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = suite_HAS_EXECUTED_API

		// prepare
		ts := &Scenario{}
		testlib_ConfigureName(s, ts)
		testlib_ConfigureDescription(s, ts)
		testlib_ConfigureLog(s, ts)
		testlib_ConfigureSwitches(s, ts)
		testlib_ConfigureVerdict(s, ts)

		// test
		output := false
		_panick := Exec(func() any {
			if !s.Switches[cond_SUPPLY_NIL_SCENARIO] {
				output = HasExecuted(ts)
			} else {
				output = HasExecuted(nil)
			}
			return ""
		})
		panick := _panick.(string)

		// log output
		Logf(s, "Test Scenario's Verdict	= %#v", ts.verdict)
		Logf(s, "Got Output			= %v", output)
		Logf(s, "Got Panic			= %q", panick)

		// assert
		if !testlib_AssertPanic(s, panick) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if !testlib_AssertHasExecuted(s, output) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if Conclusion(s) != VERDICT_FAIL {
			Conclude(s, VERDICT_PASS)
		}

		// report
		t.Logf("\n%s\n\n\n", ToString(s))
	}
}
