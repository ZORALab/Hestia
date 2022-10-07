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
	open_SWITCH_STRING    = "["
	close_SWITCH_STRING   = "]\n"
	end_SWITCH_STRING     = "\n"
	end_SWITCHES_STRING   = "\n\n"

	title_LOG_STRING = DATA_LABEL_LOG + "\t\t:" + "\n"
	open_LOG_STRING  = "["
	close_LOG_STRING = "]\n"
	end_LOG_STRING   = "\n\n"

	footer_STRING = "═══[ END ]═══\n\n\n"
)

// toml
const (
	escapeQUOTE_TOML = "'''"

	header_TOML = "[" + DATA_LABEL_GROUP + "]" + "\n"

	title_ID_TOML = DATA_LABEL_ID + " = "
	end_ID_TOML   = "\n"

	title_VERDICT_TOML = DATA_LABEL_VERDICT + " = '"
	end_VERDICT_TOML   = "'\n"

	title_NAME_TOML = DATA_LABEL_NAME + " = '''\n"
	end_NAME_TOML   = "\n'''\n"

	title_DESCRIPTION_TOML = DATA_LABEL_DESCRIPTION + " = '''\n"
	end_DESCRIPTION_TOML   = "\n'''\n"

	empty_SWITCHES_TOML = DATA_LABEL_SWITCHES + " = []\n"
	title_SWITCH_TOML   = "\n[[" + DATA_LABEL_GROUP + "." + DATA_LABEL_SWITCHES + "]]" +
		"\n" +
		DATA_LABEL_VALUE + " = " + escapeQUOTE_TOML + "\n"
	end_SWITCH_TOML = "\n" + escapeQUOTE_TOML + "\n"

	empty_LOG_WITH_EMPTY_SWITCHES_TOML = DATA_LABEL_LOG + " = []\n"
	empty_LOG_TOML                     = "\n[" + DATA_LABEL_GROUP + "]\n" +
		empty_LOG_WITH_EMPTY_SWITCHES_TOML
	title_LOG_TOML = "\n[[" + DATA_LABEL_GROUP + "." + DATA_LABEL_LOG + "]]\n" +
		DATA_LABEL_VALUE + " = " + escapeQUOTE_TOML + "\n"
	end_LOG_TOML = "\n" + escapeQUOTE_TOML + "\n"
)

func _checkBeforeRender(s *Scenario, name string) {
	if s == nil {
		panic("calling hestiaTESTING.To" + name + " without providing Scenario!")
	}

	s.Init()
}
