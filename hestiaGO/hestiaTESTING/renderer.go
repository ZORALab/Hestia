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
	"hestia/hestiaERROR"
)

// common
const (
	char_NEW_LINE = "\n"
	char_TAB      = "\t"
	char_QUOTE    = "\""
)

// verdict
const (
	string_PASS    = "PASSED"
	string_FAIL    = "FAILED"
	string_SKIP    = "SKIPPED"
	string_UNKNOWN = "UNKNOWN"
)

// string
const (
	header_STRING                   = "╔═══════════╗\n" + "║TEST REPORT║\n" + "╚═══════════╝\n"
	titleEndQuote_STRING            = char_TAB + char_TAB + ": "
	titleDescriptionEndQuote_STRING = char_TAB + ": "

	titleStartSwitch_STRING   = char_NEW_LINE
	titleEndSwitch_STRING     = char_TAB + ":" + char_NEW_LINE
	fieldSwitchOpening_STRING = "[ "
	fieldSwitchClosing_STRING = "]\t"

	titleStartLog_STRING   = char_NEW_LINE
	titleEndLog_STRING     = char_TAB + char_TAB + ":" + char_NEW_LINE
	fieldLogOpening_STRING = "("
	fieldLogClosing_STRING = ") "

	footer_STRING = "═══[ END ]═══\n\n"
)

// json
const (
	indentLV2_JSON       = char_TAB + char_TAB
	indentLV3_JSON       = char_TAB + char_TAB + char_TAB
	titleStartQuote_JSON = indentLV2_JSON + char_QUOTE
	titleEndQuote_JSON   = char_QUOTE + ": "
	fieldEnd_JSON        = ","
	fieldEndLine_JSON    = fieldEnd_JSON + char_NEW_LINE
	fieldEndString_JSON  = "\"" + fieldEndLine_JSON

	header_JSON = "{\n\t\"" + DATA_LABEL_GROUP + char_QUOTE + ": {" + char_NEW_LINE

	titleID_JSON = titleStartQuote_JSON +
		DATA_LABEL_ID +
		titleEndQuote_JSON

	titleName_JSON = titleStartQuote_JSON +
		DATA_LABEL_NAME +
		titleEndQuote_JSON +
		char_QUOTE

	titleVerdict_JSON = titleStartQuote_JSON +
		DATA_LABEL_VERDICT +
		titleEndQuote_JSON +
		char_QUOTE

	titleDescription_JSON = titleStartQuote_JSON +
		DATA_LABEL_DESCRIPTION +
		titleEndQuote_JSON +
		char_QUOTE

	titleSwitches_JSON = titleStartQuote_JSON +
		DATA_LABEL_SWITCHES +
		titleEndQuote_JSON
	titleSwitchesOpen_JSON  = titleSwitches_JSON + "{"
	titleSwitchesClose_JSON = "\n\t\t}," + char_NEW_LINE
	titleSwitchesEmpty_JSON = titleSwitches_JSON + "{}," + char_NEW_LINE

	titleLog_JSON = titleStartQuote_JSON +
		DATA_LABEL_LOG +
		titleEndQuote_JSON
	titleLogOpen_JSON  = titleLog_JSON + "["
	titleLogClose_JSON = char_NEW_LINE + "\t\t]" + char_NEW_LINE
	titleLogEmpty_JSON = titleLog_JSON + "[]" + char_NEW_LINE

	footer_JSON = "\t}\n}"
)

// toml
const (
	fieldEndString_TOML  = char_QUOTE + char_NEW_LINE
	fieldEnd_TOML        = "," + char_NEW_LINE
	titleStartQuote_TOML = char_QUOTE
	titleEndQuote_TOML   = char_QUOTE + " = "

	header_TOML = "[" + DATA_LABEL_GROUP + "]" + char_NEW_LINE

	titleID_TOML = titleStartQuote_TOML + DATA_LABEL_ID + titleEndQuote_TOML

	titleName_TOML = titleStartQuote_TOML + DATA_LABEL_NAME + titleEndQuote_TOML

	titleVerdict_TOML = titleStartQuote_TOML + DATA_LABEL_VERDICT + titleEndQuote_TOML

	titleDescription_TOML = titleStartQuote_TOML +
		DATA_LABEL_DESCRIPTION +
		titleEndQuote_TOML

	titleLog_TOML      = char_QUOTE + DATA_LABEL_LOG + titleEndQuote_TOML
	titleLogOpen_TOML  = titleLog_TOML + "[" + char_NEW_LINE
	titleLogClose_TOML = "]" + char_NEW_LINE + char_NEW_LINE
	titleLogEmpty_TOML = titleLog_TOML + "[]" + char_NEW_LINE + char_NEW_LINE

	titleSwitches_TOML = "[" +
		DATA_LABEL_GROUP + "." +
		DATA_LABEL_SWITCHES + "]" +
		char_NEW_LINE
)

// yaml
const (
	tab_YAML                = "  "
	listEmpty_YAML          = "[]"
	titleStartQuote_YAML    = tab_YAML + char_QUOTE
	titleLV3StartQuote_YAML = tab_YAML + titleStartQuote_YAML
	titleEndQuote_YAML      = char_QUOTE + ": "

	header_YAML = "---" + char_NEW_LINE + DATA_LABEL_GROUP + ":" + char_NEW_LINE

	titleID_YAML = titleStartQuote_YAML + DATA_LABEL_ID + titleEndQuote_YAML

	titleName_YAML = titleStartQuote_YAML + DATA_LABEL_NAME + titleEndQuote_YAML

	titleVerdict_YAML = titleStartQuote_YAML + DATA_LABEL_VERDICT + titleEndQuote_YAML

	titleDescription_YAML = titleStartQuote_YAML +
		DATA_LABEL_DESCRIPTION +
		titleEndQuote_YAML

	titleSwitches_YAML      = titleStartQuote_YAML + DATA_LABEL_SWITCHES + char_QUOTE + ":"
	titleSwitchesOpen_YAML  = titleSwitches_YAML + char_NEW_LINE
	titleSwitchesClose_YAML = char_NEW_LINE
	titleSwitchesEmpty_YAML = titleSwitches_YAML + " " + listEmpty_YAML + char_NEW_LINE

	titleLog_YAML        = titleStartQuote_YAML + DATA_LABEL_LOG + titleEndQuote_YAML
	titleLogOpen_YAML    = titleLog_YAML + char_NEW_LINE
	titleLogEmpty_YAML   = titleLog_YAML + listEmpty_YAML + char_NEW_LINE
	arrayLogOpening_YAML = tab_YAML + tab_YAML + "- " + char_QUOTE
)

func _checkBeforeRender(s *Scenario, name string) {
	if s == nil {
		panic("calling hestiaTESTING.To" + name + " without providing Scenario!")
	}

	if s.Init() != hestiaERROR.OK {
		panic("calling hestiaTESTING.To" + name + " with unregistered/faulty Scenario!")
	}
}
