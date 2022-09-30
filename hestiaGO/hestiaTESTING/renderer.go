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
	"fmt"
	"hestia/hestiaERROR"
	"unicode"
)

// common
const (
	char_NEW_LINE = "\n"
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
	header_STRING = "\n" +
		"╔═══════════╗" + "\n" +
		"║TEST REPORT║" + "\n" +
		"╚═══════════╝" + "\n"

	title_ID_STRING = DATA_LABEL_ID + "\t\t: "
	end_ID_STRING   = "\n"

	title_NAME_STRING = DATA_LABEL_NAME + "\t\t: "
	end_NAME_STRING   = "\n"

	title_VERDICT_STRING = DATA_LABEL_VERDICT + "\t\t: "
	end_VERDICT_STRING   = "\n"

	title_DESCRIPTION_STRING = DATA_LABEL_DESCRIPTION + "\t:\n"
	end_DESCRIPTION_STRING   = "\n\n\n"

	title_SWITCHES_STRING = DATA_LABEL_SWITCHES + "\t:\n"
	end_SWITCHES_STRING   = "\n\n"
	open_SWITCH_STRING    = "[ "
	close_SWITCH_STRING   = " ]\t"
	end_SWITCH_STRING     = "\n"

	title_LOG_STRING = DATA_LABEL_LOG + "\t\t:" + "\n"
	open_LOG_STRING  = "[ "
	close_LOG_STRING = " ]\n"
	end_LOG_STRING   = "\n\n"

	footer_STRING = "═══[ END ]═══" + "\n" + "\n"
)

// toml
const (
	escapeQUOTE_TOML = "'''"

	header_TOML = "[" + DATA_LABEL_GROUP + "]" + "\n"

	title_ID_TOML = DATA_LABEL_ID + " = "
	end_ID_TOML   = "\n"

	title_VERDICT_TOML = DATA_LABEL_VERDICT + " = "
	end_VERDICT_TOML   = "\n"

	title_NAME_TOML = DATA_LABEL_NAME + " = '''\n"
	end_NAME_TOML   = "\n'''\n"

	title_DESCRIPTION_TOML = DATA_LABEL_DESCRIPTION + " = '''\n"
	end_DESCRIPTION_TOML   = "\n'''\n\n"

	title_SWITCHES_TOML = "[" + DATA_LABEL_GROUP + "." + DATA_LABEL_SWITCHES + "]" + "\n"
	field_SWITCH_TOML   = escapeQUOTE_TOML + " = "

	title_LOG_EMPTY_TOML = DATA_LABEL_LOG + " = []\n"
	title_LOG_TOML       = "\n" +
		"[[" + DATA_LABEL_GROUP + "." + DATA_LABEL_LOG + "]]" + "\n" +
		DATA_LABEL_VALUE + " = " + escapeQUOTE_TOML + "\n"
	end_LOG_TOML = "\n'''\n"
)

func _checkBeforeRender(s *Scenario, name string) {
	if s == nil {
		panic("calling hestiaTESTING.To" + name + " without providing Scenario!")
	}

	if s.Init() != hestiaERROR.OK {
		panic("calling hestiaTESTING.To" + name + " with unregistered/faulty Scenario!")
	}
}

func _renderString(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

func _renderBool(input bool) string {
	if input {
		return "true"
	}

	return "false"
}

func _renderNumber(number uint64) (out string) {
	var i int
	var x uint64
	var buffer [64]byte

	if number == 0 {
		return "0"
	}

	i = len(buffer)
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
	var setRight, next bool

	if source == "" {
		return source
	}

	setRight = false
	for i, char := range source {
		if !unicode.IsSpace(char) {
			if i == 0 {
				setRight = true
			}

			if !setRight {
				leftIndex = uint64(i)
			}

			rightIndex = uint64(i)
			next = true
			setRight = true
			continue
		}

		if next {
			rightIndex = uint64(i)
			next = false
		}
	}

	if next {
		rightIndex = uint64(len(source)) // complete last character
	}

	return source[leftIndex:rightIndex]
}
