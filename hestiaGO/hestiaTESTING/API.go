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

	"hestia/hestiaERROR"
	"hestia/hestiaSTRING"
)

func Register(s *Scenario, t *testing.T) {
	if s == nil {
		panic("calling hestiaTESTING.Register without providing Scenario!")
	}

	if t == nil {
		panic("calling hestiaTESTING.Register without providing *testing.T!")
	}

	s.controller = t
	s.skip = s.controller.SkipNow
	s.fail = s.controller.FailNow
}

func Logf(s *Scenario, format string, args ...any) {
	if s == nil {
		panic("calling hestiaTESTING.Logf without providing Scenario!")
	}

	if format == "" {
		panic("calling hestiaTESTING.Logf without providing formatting string!")
	}

	if s.Init() != hestiaERROR.OK {
		panic("calling hestiaTESTING.Logf with unregistered/faulty Scenario!")
	}

	s.Log = append(s.Log, hestiaSTRING.Printf(format, args...))
}

func Logln(s *Scenario, args ...any) {
	if s == nil {
		panic("calling hestiaTESTING.Logln without providing Scenario!")
	}

	if s.Init() != hestiaERROR.OK {
		panic("calling hestiaTESTING.Logf with unregistered/faulty Scenario!")
	}

	s.Log = append(s.Log, hestiaSTRING.Println(args...))
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
	case s.controller == nil || s.skip == nil || s.fail == nil:
		panic("calling hestiaTESTING.Conclude without registering *testing.T!")
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

	// render report header
	output = header_STRING

	// render ID
	output += DATA_LABEL_ID + titleEndQuote_STRING
	output += hestiaSTRING.S_FormatUINT(s.ID, 10) + char_NEW_LINE

	// render Name
	output += DATA_LABEL_NAME + titleEndQuote_STRING
	output += s.Name + char_NEW_LINE

	// render Verdict
	output += DATA_LABEL_VERDICT + titleEndQuote_STRING
	output += Interpret(s.verdict) + char_NEW_LINE

	// render description
	output += DATA_LABEL_DESCRIPTION + titleDescriptionEndQuote_STRING
	output += s.Description + char_NEW_LINE

	// render switches
	output += titleStartSwitch_STRING + DATA_LABEL_SWITCHES + titleEndSwitch_STRING
	for k, v := range s.Switches {
		output += fieldSwitchOpening_STRING +
			hestiaSTRING.FormatBOOL(v) +
			fieldSwitchClosing_STRING +
			hestiaSTRING.TrimWhitespace(k) +
			char_NEW_LINE
	}
	output += char_NEW_LINE

	// render log
	output += titleStartLog_STRING + DATA_LABEL_LOG + titleEndLog_STRING
	for i, v := range s.Log {
		output += fieldLogOpening_STRING +
			hestiaSTRING.S_Itoa(int64(i)) +
			fieldLogClosing_STRING
		output += hestiaSTRING.TrimWhitespace(v) + char_NEW_LINE
	}

	// render report footer
	output += footer_STRING

	return output
}

func ToJSON(s *Scenario) (output string) {
	var first bool

	_checkBeforeRender(s, "JSON")

	// render header
	output = header_JSON

	// render ID
	output += titleID_JSON + hestiaSTRING.S_FormatUINT(s.ID, 10) + fieldEndLine_JSON

	// render Name
	output += titleName_JSON + s.Name + fieldEndString_JSON

	// render Verdict
	output += titleVerdict_JSON + Interpret(s.verdict) + fieldEndString_JSON

	// render Description
	output += titleDescription_JSON + s.Description + fieldEndString_JSON

	// render Switches
	if len(s.Switches) > 0 {
		output += titleSwitchesOpen_JSON
		first = true
		for k, v := range s.Switches {
			if !first {
				output += fieldEnd_JSON
			}

			// JSON escape
			k = hestiaSTRING.ReplaceAll(k, "\"", "\\\"")

			// render element
			output += char_NEW_LINE + char_TAB +
				titleStartQuote_JSON +
				hestiaSTRING.TrimWhitespace(k) +
				titleEndQuote_JSON +
				hestiaSTRING.FormatBOOL(v)

			first = false
		}
		output += titleSwitchesClose_JSON
	} else {
		output += titleSwitchesEmpty_JSON
	}

	// render Log
	if len(s.Log) > 0 {
		output += titleLogOpen_JSON
		first = true
		for _, v := range s.Log {
			if !first {
				output += fieldEnd_JSON
			}

			// json escape
			v = hestiaSTRING.ReplaceAll(v, "\"", "\\\"")

			// render element
			output += char_NEW_LINE + indentLV3_JSON +
				char_QUOTE + hestiaSTRING.TrimWhitespace(v) + char_QUOTE

			first = false
		}
		output += titleLogClose_JSON
	} else {
		output += titleLogEmpty_JSON
	}

	// render report footer
	output += footer_JSON

	return output
}

func ToTOML(s *Scenario) (output string) {
	var first bool

	_checkBeforeRender(s, "TOML")

	// render header
	output = header_TOML

	// render ID field
	output += titleID_TOML + hestiaSTRING.S_FormatUINT(s.ID, 10) + char_NEW_LINE

	// render Name field
	output += titleName_TOML + char_QUOTE + s.Name + fieldEndString_TOML

	// render Verdict field
	output += titleVerdict_TOML + Interpret(s.verdict) + fieldEndString_TOML

	// render Description field
	output += titleDescription_TOML + s.Description + fieldEndString_TOML

	// render Log
	if len(s.Log) > 0 {
		output += titleLogOpen_TOML
		for _, v := range s.Log {
			// TOML escape
			v = hestiaSTRING.ReplaceAll(v, "\"", "\\\"")

			// render element
			output += char_TAB +
				char_QUOTE + hestiaSTRING.TrimWhitespace(v) + char_QUOTE +
				fieldEnd_TOML
		}
		output += titleLogClose_TOML
	} else {
		output += titleLogEmpty_TOML
	}

	// render Switches
	output += titleSwitches_TOML
	first = true
	for k, v := range s.Switches {
		if !first {
			output += char_NEW_LINE
		}

		// TOML escape
		k = hestiaSTRING.ReplaceAll(k, "\"", "\\\"")

		// render element
		output += titleStartQuote_TOML +
			hestiaSTRING.TrimWhitespace(k) +
			titleEndQuote_TOML +
			hestiaSTRING.FormatBOOL(v)

		first = false
	}

	// render report footer
	// -- HOORAY! Nothing to render --

	return output
}

func ToYAML(s *Scenario) (output string) {
	var first bool

	_checkBeforeRender(s, "YAML")

	// render header
	output = header_YAML

	// render ID field
	output += titleID_YAML +
		hestiaSTRING.S_FormatUINT(s.ID, 10) + char_NEW_LINE

	// render Name field
	output += titleName_YAML +
		char_QUOTE + s.Name + char_QUOTE + char_NEW_LINE

	// render Verdict field
	output += titleVerdict_YAML +
		char_QUOTE + Interpret(s.verdict) + char_QUOTE + char_NEW_LINE

	// render Description field
	output += titleDescription_YAML +
		char_QUOTE + s.Description + char_QUOTE + char_NEW_LINE

	// render Switches
	if len(s.Switches) > 0 {
		output += titleSwitchesOpen_YAML
		first = true
		for k, v := range s.Switches {
			if !first {
				output += char_NEW_LINE
			}

			// YAML escape
			k = hestiaSTRING.ReplaceAll(k, "\"", "\\\"")

			// render element
			output += titleLV3StartQuote_YAML +
				hestiaSTRING.TrimWhitespace(k) +
				titleEndQuote_YAML +
				hestiaSTRING.FormatBOOL(v)

			first = false
		}
		output += titleSwitchesClose_YAML
	} else {
		output += titleSwitchesEmpty_YAML
	}

	// render Log
	if len(s.Log) > 0 {
		output += titleLogOpen_YAML
		first = true
		for _, v := range s.Log {
			if !first {
				output += char_NEW_LINE
			}

			// YAML escape
			v = hestiaSTRING.ReplaceAll(v, "\"", "\\\"")

			// render element
			output += arrayLogOpening_YAML +
				hestiaSTRING.TrimWhitespace(v) +
				char_QUOTE

			first = false
		}
	} else {
		output += titleLogEmpty_YAML
	}

	// render report footer
	// -- HOORAY! Nothing to render --

	return output
}
