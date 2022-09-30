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
	"unicode"
)

// common
const (
	char_NEW_LINE = "\n"
	char_TAB      = "\t"
	char_QUOTE    = "\""
	char_DIGITS   = "0123456789abcdefghijklmnopqrstuvwxyz"
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
	header_STRING = char_NEW_LINE +
		"╔═══════════╗" + char_NEW_LINE +
		"║TEST REPORT║" + char_NEW_LINE +
		"╚═══════════╝" + char_NEW_LINE
	titleEndQuote_STRING            = char_TAB + char_TAB + ": "
	titleDescriptionEndQuote_STRING = char_TAB + ": "

	titleStartSwitch_STRING   = char_NEW_LINE
	titleEndSwitch_STRING     = char_TAB + ":" + char_NEW_LINE
	fieldSwitchOpening_STRING = "[ "
	fieldSwitchClosing_STRING = " ]" + char_TAB

	titleStartLog_STRING   = char_NEW_LINE
	titleEndLog_STRING     = char_TAB + char_TAB + ":" + char_NEW_LINE
	fieldLogOpening_STRING = "("
	fieldLogClosing_STRING = ") "

	footer_STRING = "═══[ END ]═══" + char_NEW_LINE + char_NEW_LINE
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

func _checkBeforeRender(s *Scenario, name string) {
	if s == nil {
		panic("calling hestiaTESTING.To" + name + " without providing Scenario!")
	}

	if s.Init() != hestiaERROR.OK {
		panic("calling hestiaTESTING.To" + name + " with unregistered/faulty Scenario!")
	}
}

func _renderID(s *Scenario) (out string) {
	var i int
	var number, x uint64
	var buffer [64]byte

	if s.ID == 0 {
		return "0"
	}

	i = len(buffer)

	number = s.ID
	for number >= 10 {
		i--

		x = number / 10
		buffer[i] = char_DIGITS[uint(number-x*10)]
		number = x
	}

	i--
	buffer[i] = char_DIGITS[uint(number)]

	return string(buffer[i:])
}

func __trimWhitespace(source string) string {
	var leftIndex, rightIndex uint64
	var first, next bool

	if source == "" {
		return source
	}

	first = true
	for i, char := range source {
		if !unicode.IsSpace(char) {
			if first {
				leftIndex = uint64(i)
			}

			rightIndex = uint64(i)
			next = true
			first = false
			continue
		}

		if next {
			rightIndex = uint64(i)
			next = false
		}

		first = false
	}

	if next {
		rightIndex = uint64(len(source))
	}

	return source[leftIndex:rightIndex]
}
