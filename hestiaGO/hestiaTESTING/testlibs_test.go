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

// test suite
const (
	suite_REGISTER_API     = "hestiaTESTING Register API"
	suite_LOGF_API         = "hestiaTESTING Logf API"
	suite_LOGLN_API        = "hestiaTESTING Logln API"
	suite_HAS_EXECUTED_API = "hestiaTESTING HasExecuted API"
	suite_CONCLUDE_API     = "hestiaTESTING Conclude API"
	suite_CONCLUSION_API   = "hestiaTESTING Conclusion API"
	suite_TO_STRING_API    = "hestiaTESTING ToString API"
	suite_TO_TOML_API      = "hestiaTESTING ToTOML API"
)

// all test switches
const (
	// scenario
	cond_SUPPLY_NIL_SCENARIO = "supply scenario parameter as nil"

	// register
	cond_PROPER_REGISTRATION      = "configure registration properly"
	cond_NIL_REGISTRATION         = "configure registration to register nil *testing.T"
	cond_FAULTY_SKIP_REGISTRATION = "configure registration with faulty skip"
	cond_FAULTY_FAIL_REGISTRATION = "configure registration with faulty fail"
	cond_FAULTY_BOTH_REGISTRATION = "configure registration with faulty skip and fail"

	// name
	cond_PROPER_NAME = "configure proper name"
	cond_EMPTY_NAME  = "configure empty name"

	// verdict
	cond_PROPER_VERDICT  = "configure VERDICT_PASS verdict"
	cond_FAIL_VERDICT    = "configure VERDICT_FAIL verdict"
	cond_SKIP_VERDICT    = "configure VERDICT_SKIP verdict"
	cond_UNKNOWN_VERDICT = "configure priv_VERDICT_UNKNOWN verdict"

	// description
	cond_PROPER_DESCRIPTION = "configure proper description"
	cond_EMPTY_DESCRIPTION  = "configure empty description"

	// log
	cond_PROPER_LOG = "configure proper log list"
	cond_EMPTY_LOG  = "configure empty log list"
	cond_NIL_LOG    = "configure nil log list"

	// switches
	cond_PROPER_SWITCHES = "configure proper switches list"
	cond_EMPTY_SWITCHES  = "configure empty switches list"
	cond_NIL_SWITCHES    = "configure nil switches list"

	// format
	cond_PROPER_STRING_FORMAT = "use proper string format"
	cond_EMPTY_STRING_FORMAT  = "use empty string format"

	// args
	cond_PROPER_STRING_ARGUMENTS = "use proper string arguments"
	cond_EMPTY_STRING_ARGUMENTS  = "use empty string arguments"
	cond_NIL_STRING_ARGUMENTS    = "use nil string arguments"

	// expectations
	expect_PANIC         = "expecting panic"
	expect_OUTPUT_STRING = "expecting string output"
)

// all test values
const (
	/// name
	t_NAME_PROPER = "Test Scenario Object"

	// description
	t_DESCRIPTION_PROPER = "A Test Scenario Object Description"

	// log
	t_LOG_LINE_1 = "Log Line 1"
	t_LOG_LINE_2 = "Log Line 1"

	// switches
	t_SWITCH_1 = "Switch 1"
	t_SWITCH_2 = "Switch 2"

	// string format
	t_FORMAT_1                 = "formatted %v %v %v"
	t_FORMAT_1_LOGF_SUCCESSFUL = "formatted proper1 5 true"
	t_FORMAT_2_LOGF_SUCCESSFUL = "formatted %!v(MISSING) %!v(MISSING) %!v(MISSING)"
	t_FORMAT_3_LOGF_SUCCESSFUL = "formatted <nil> %!v(MISSING) %!v(MISSING)"

	// string arguments
	t_ARG_1                  = "proper1"
	t_ARG_2                  = 5
	t_ARG_3                  = true
	t_ARG_1_LOGLN_SUCCESSFUL = "proper1 5 true\n"
	t_ARG_2_LOGLN_SUCCESSFUL = "\n"
	t_ARG_3_LOGLN_SUCCESSFUL = "<nil>\n"
)

func testlib_AssertOutputString(s *Scenario, output string) bool {
	if output != "" {
		return s.Switches[expect_OUTPUT_STRING]
	}

	return !s.Switches[expect_OUTPUT_STRING]
}

func testlib_AssertPanic(s *Scenario, panick string) bool {
	if panick != "" {
		return s.Switches[expect_PANIC]
	}

	return !s.Switches[expect_PANIC]
}

func testlib_AssertConclude(s, ts *Scenario) bool {
	switch {
	case ts.verdict == 125:
		return s.Switches[cond_SUPPLY_NIL_SCENARIO] ||
			s.Switches[expect_PANIC]
	case ts.verdict == VERDICT_PASS:
		return s.Switches[cond_PROPER_VERDICT]
	case ts.verdict == VERDICT_FAIL:
		return s.Switches[cond_FAIL_VERDICT]
	case ts.verdict == VERDICT_SKIP:
		return s.Switches[cond_SKIP_VERDICT]
	}

	return false
}

func testlib_GenerateVerdict(s *Scenario) Verdict {
	switch {
	case s.Switches[cond_PROPER_VERDICT]:
		return VERDICT_PASS
	case s.Switches[cond_FAIL_VERDICT]:
		return VERDICT_FAIL
	case s.Switches[cond_SKIP_VERDICT]:
		return VERDICT_SKIP
	default:
		return priv_VERDICT_UNKNOWN
	}
}

func testlib_AssertConclusion(s *Scenario, output Verdict) bool {
	switch {
	case output == priv_VERDICT_UNKNOWN:
		return s.Switches[cond_UNKNOWN_VERDICT]
	case output == VERDICT_PASS:
		return s.Switches[cond_PROPER_VERDICT]
	case output == VERDICT_SKIP:
		if !s.Switches[cond_PROPER_VERDICT] || s.Switches[cond_SUPPLY_NIL_SCENARIO] {
			return true
		}
	}

	return false
}

func testlib_AssertHasExecuted(s *Scenario, output bool) bool {
	if !output {
		switch {
		case s.Switches[cond_SUPPLY_NIL_SCENARIO],
			s.Switches[cond_UNKNOWN_VERDICT]:
			return true
		}
	}
	// is true

	return s.Switches[cond_PROPER_VERDICT]
}

func testlib_ConfigureVerdict(s, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_VERDICT]:
		ts.verdict = VERDICT_PASS
	case s.Switches[cond_UNKNOWN_VERDICT]:
		ts.verdict = priv_VERDICT_UNKNOWN
	default:
		ts.verdict = VERDICT_PASS
	}
}

func testlib_AssertRegistration(s, ts *Scenario) bool {
	if ts == nil {
		return false
	}
	// ts is now available

	switch {
	case s.Switches[expect_PANIC]:
		switch {
		case s.Switches[cond_SUPPLY_NIL_SCENARIO],
			s.Switches[cond_NIL_REGISTRATION]:
			return ts.controller == nil && ts.skip == nil && ts.fail == nil
		}
	case ts.controller != nil && ts.skip != nil && ts.fail != nil:
		return s.Switches[cond_PROPER_REGISTRATION] &&
			!s.Switches[cond_SUPPLY_NIL_SCENARIO]
	}

	return false
}

func testlib_AssertLog(s, ts *Scenario) bool {
	if ts == nil {
		return s.Switches[cond_SUPPLY_NIL_SCENARIO]
	}
	// ts is now available

	if s.Switches[expect_PANIC] {
		switch {
		case s.Switches[cond_EMPTY_STRING_FORMAT],
			s.Switches[cond_SUPPLY_NIL_SCENARIO],
			s.Switches[cond_FAULTY_FAIL_REGISTRATION],
			s.Switches[cond_FAULTY_SKIP_REGISTRATION],
			s.Switches[cond_FAULTY_BOTH_REGISTRATION]:
			return true
		}
	}
	// no longer in panic expectation.

	if ts.Log == nil {
		return false
	}
	// ts.Log is now available

	for _, v := range ts.Log {
		switch v {
		case t_FORMAT_1_LOGF_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_FORMAT] &&
				s.Switches[cond_PROPER_STRING_ARGUMENTS] {
				return true
			}
		case t_FORMAT_2_LOGF_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_FORMAT] &&
				s.Switches[cond_EMPTY_STRING_ARGUMENTS] {
				return true
			}
		case t_FORMAT_3_LOGF_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_FORMAT] &&
				s.Switches[cond_NIL_STRING_ARGUMENTS] {
				return true
			}
		case t_ARG_1_LOGLN_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_ARGUMENTS] {
				return true
			}
		case t_ARG_2_LOGLN_SUCCESSFUL:
			if s.Switches[cond_EMPTY_STRING_ARGUMENTS] {
				return true
			}
		case t_ARG_3_LOGLN_SUCCESSFUL:
			if s.Switches[cond_NIL_STRING_ARGUMENTS] {
				return true
			}
		}
	}

	return false
}

func testlib_CreateStringArguments(s *Scenario) []any {
	switch {
	case s.Switches[cond_PROPER_STRING_ARGUMENTS]:
		return []any{t_ARG_1, t_ARG_2, t_ARG_3}
	case s.Switches[cond_EMPTY_STRING_ARGUMENTS]:
		return []any{}
	case s.Switches[cond_NIL_STRING_ARGUMENTS]:
		return nil
	}

	return nil
}

func testlib_CreateStringFormat(s *Scenario) string {
	switch {
	case s.Switches[cond_PROPER_STRING_FORMAT]:
		return t_FORMAT_1
	case s.Switches[cond_EMPTY_STRING_FORMAT]:
		return ""
	}

	return ""
}

func testlib_ConfigureRegistrations(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = func() {}
		ts.fail = func() {}
	case s.Switches[cond_FAULTY_SKIP_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = nil
		ts.fail = func() {}
	case s.Switches[cond_FAULTY_FAIL_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = func() {}
		ts.fail = nil
	case s.Switches[cond_FAULTY_BOTH_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = nil
		ts.fail = nil
	}
}

func testlib_ConfigureName(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_NAME]:
		ts.Name = t_NAME_PROPER
	case s.Switches[cond_EMPTY_NAME]:
		ts.Name = ""
	}
}

func testlib_ConfigureDescription(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_DESCRIPTION]:
		ts.Description = t_DESCRIPTION_PROPER
	case s.Switches[cond_EMPTY_DESCRIPTION]:
		ts.Description = ""
	}
}

func testlib_ConfigureLog(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_LOG]:
		ts.Log = []string{t_LOG_LINE_1, t_LOG_LINE_2}
	case s.Switches[cond_EMPTY_LOG]:
		ts.Log = []string{}
	case s.Switches[cond_NIL_LOG]:
		ts.Log = nil
	}
}

func testlib_ConfigureSwitches(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_SWITCHES]:
		ts.Switches = map[string]bool{
			t_SWITCH_1: true,
			t_SWITCH_2: false,
		}
	case s.Switches[cond_EMPTY_SWITCHES]:
		ts.Switches = map[string]bool{}
	case s.Switches[cond_NIL_SWITCHES]:
		ts.Switches = nil
	}
}
