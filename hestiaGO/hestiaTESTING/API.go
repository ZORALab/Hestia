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

func Log(s *Scenario, statement string) {
	if s == nil {
		panic("calling hestiaTESTING.Log without providing Scenario!")
	}

	if statement == "" {
		panic("calling hestiaTESTING.Log without providing formatting string!")
	}

	s.Init()
	s.Logs = append(s.Logs, statement)
}

func Exec(function func() any) (out any) {
	defer func() {
		err := recover()
		if err != nil {
			out = err
		}
	}()

	out = function()

	return out
}

func HasExecuted(s *Scenario) bool {
	if s == nil {
		panic("calling hestiaTESTING.HasExecuted without providing Scenario!")
	}

	return s.verdict != priv_VERDICT_UNKNOWN
}

func Conclude(s *Scenario, certification Verdict) {
	switch {
	case s == nil:
		panic("calling hestiaTESTING.Conclude without providing Scenario!")
	case certification == VERDICT_PASS:
		s.verdict = VERDICT_PASS
	case certification == VERDICT_FAIL:
		s.verdict = VERDICT_FAIL
	case certification == VERDICT_SKIP:
		s.verdict = VERDICT_SKIP
	default:
		panic("calling hestiaTESTING.Conclude an with unknown Verdict!")
	}
}

func Conclusion(s *Scenario) Verdict {
	if s == nil {
		panic("calling hestiaTESTING.Verdict without providing Scenario!")
	}

	return s.verdict
}

func Interpret(verdict Verdict) string {
	switch verdict {
	case VERDICT_PASS:
		return string_PASS
	case VERDICT_FAIL:
		return string_FAIL
	case VERDICT_SKIP:
		return string_SKIP
	default:
		return string_UNKNOWN
	}
}

func ToString(s *Scenario) (output string) {
	_checkBeforeRender(s, "String")

	output = header_STRING
	output += title_ID_STRING + _renderNumber(s.ID, 10) + end_ID_STRING
	output += title_NAME_STRING + s.Name + end_NAME_STRING
	output += title_VERDICT_STRING + Interpret(s.verdict) + end_VERDICT_STRING
	output += title_DESCRIPTION_STRING + s.Description + end_DESCRIPTION_STRING

	// render switches
	output += title_SWITCHES_STRING
	for k, v := range s.Switches {
		output += open_SWITCH_STRING + _renderBool(v) + close_SWITCH_STRING +
			__trimWhitespace(k) +
			end_SWITCH_STRING
	}
	output += end_SWITCHES_STRING

	// render log
	output += title_LOG_STRING
	for i, v := range s.Logs {
		output += open_LOG_STRING + _renderNumber(uint64(i), 10) + close_LOG_STRING +
			__trimWhitespace(v) +
			end_LOG_STRING
	}

	// render report footer
	output += footer_STRING

	return output
}

func ToTOML(s *Scenario) (output string) {
	var first bool

	_checkBeforeRender(s, "TOML")

	// render header
	output = header_TOML
	output += title_ID_TOML + _renderNumber(s.ID, 10) + end_ID_TOML
	output += title_VERDICT_TOML + Interpret(s.verdict) + end_VERDICT_TOML
	output += title_NAME_TOML + s.Name + end_NAME_TOML
	output += title_DESCRIPTION_TOML + s.Description + end_DESCRIPTION_TOML
	if len(s.Logs) == 0 {
		output += title_LOG_EMPTY_TOML
	}

	// render Switches
	output += title_SWITCHES_TOML
	first = true
	for k, v := range s.Switches {
		if !first {
			output += char_NEW_LINE
		}
		output += escapeQUOTE_TOML +
			__trimWhitespace(k) +
			field_SWITCH_TOML +
			_renderBool(v)
		first = false
	}

	// render Log if available
	if len(s.Logs) > 0 {
		output += char_NEW_LINE
		for _, v := range s.Logs {
			output += title_LOG_TOML + __trimWhitespace(v) + end_LOG_TOML
		}
	}

	// render report footer
	// -- HOORAY! Nothing to render --

	return output
}
