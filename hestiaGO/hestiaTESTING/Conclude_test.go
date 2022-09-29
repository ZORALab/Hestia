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

func _testConcludeScenarios() []*Scenario {
	return []*Scenario{
		{
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly with proper Scenario settings.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to panic with fault fail function registration.
`,
			Switches: map[string]bool{
				cond_FAULTY_FAIL_REGISTRATION: true,
				cond_PROPER_NAME:              true,
				cond_PROPER_DESCRIPTION:       true,
				cond_PROPER_LOG:               true,
				cond_PROPER_SWITCHES:          true,
				cond_PROPER_VERDICT:           true,
				expect_PANIC:                  true,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to panic with fault skip function registration.
`,
			Switches: map[string]bool{
				cond_FAULTY_SKIP_REGISTRATION: true,
				cond_PROPER_NAME:              true,
				cond_PROPER_DESCRIPTION:       true,
				cond_PROPER_LOG:               true,
				cond_PROPER_SWITCHES:          true,
				cond_PROPER_VERDICT:           true,
				expect_PANIC:                  true,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to panic with faulty skip and fail functions registration.
`,
			Switches: map[string]bool{
				cond_FAULTY_BOTH_REGISTRATION: true,
				cond_PROPER_NAME:              true,
				cond_PROPER_DESCRIPTION:       true,
				cond_PROPER_LOG:               true,
				cond_PROPER_SWITCHES:          true,
				cond_PROPER_VERDICT:           true,
				expect_PANIC:                  true,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly with empty Name settings.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_EMPTY_NAME:          true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly with empty Switches settings.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_EMPTY_SWITCHES:      true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly with nil Switches settings.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_NIL_SWITCHES:        true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly with empty log settings.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_EMPTY_LOG:           true,
				cond_PROPER_SWITCHES:     true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly with nil log settings.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_NIL_LOG:             true,
				cond_PROPER_SWITCHES:     true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly with empty description settings.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_EMPTY_DESCRIPTION:   true,
				cond_NIL_LOG:             true,
				cond_PROPER_SWITCHES:     true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to panic when nil Scenario is supplied.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_SUPPLY_NIL_SCENARIO: true,
				cond_PROPER_VERDICT:      true,
				expect_PANIC:             true,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to panic when verdict is supplied with unknown.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_UNKNOWN_VERDICT:     true,
				expect_PANIC:             true,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly when verdict is supplied with fail.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_FAIL_VERDICT:        true,
				expect_PANIC:             false,
			},
		}, {
			Name: suite_CONCLUDE_API,
			Description: `
Test Conclude() able to work properly when verdict is supplied with skip.
`,
			Switches: map[string]bool{
				cond_PROPER_REGISTRATION: true,
				cond_PROPER_NAME:         true,
				cond_PROPER_DESCRIPTION:  true,
				cond_PROPER_LOG:          true,
				cond_PROPER_SWITCHES:     true,
				cond_SKIP_VERDICT:        true,
				expect_PANIC:             false,
			},
		},
	}
}

func TestConcludeAPI(t *testing.T) {
	scenarios := _testConcludeScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		Register(s, t)

		// prepare
		ts := &Scenario{}
		testlib_ConfigureRegistrations(s, ts)
		testlib_ConfigureName(s, ts)
		testlib_ConfigureDescription(s, ts)
		testlib_ConfigureLog(s, ts)
		testlib_ConfigureSwitches(s, ts)
		ts.verdict = Verdict(125)
		verdict := testlib_GenerateVerdict(s)

		// test
		_panick := Exec(func() any {
			if !s.Switches[cond_SUPPLY_NIL_SCENARIO] {
				Conclude(ts, verdict)
			} else {
				Conclude(nil, verdict)
			}
			return ""
		})
		panick := _panick.(string)

		// log output
		Logf(s, "Test Scenario's Verdict	= %#v", ts.verdict)
		Logf(s, "Got Panic			= %q", panick)

		// assert
		if s.Switches[expect_PANIC] && panick == "" ||
			!s.Switches[expect_PANIC] && panick != "" {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if !testlib_AssertConclude(s, ts) {
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
