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
use crate::hestia_testing::data;
use crate::hestia_testing::scenario;

pub fn to_string(s: &scenario::Scenario) -> String {
	// render switches
	let mut switches = String::from("");
	for (i, v) in s.switches.iter().enumerate() {
		switches.push_str(&format!("[{}]\t{}\n", i, v));
	}

	// render log
	let mut logs = String::from("");
	for (i, v) in s.logs.iter().enumerate() {
		logs.push_str(&format!("[{}]\n{}\n", i, v));
	}

	// render output
	return format!(
		"
╔═══════════╗
║TEST REPORT║
╚═══════════╝
{label_id}		: {id}
{label_verdict}		: {verdict}
{label_name}		: {name}
{label_description}	:
{description}


{label_switches}\t	:
{switches}


{label_log}		:
{logs}
═══[ END ]═══
",
		label_id = data::LABEL_ID,
		id = s.id,
		label_verdict = data::LABEL_VERDICT,
		verdict = data::interpret(scenario::conclusion(s)),
		label_name = data::LABEL_NAME,
		name = s.name,
		label_description = data::LABEL_DESCRIPTION,
		description = s.description,
		label_switches = data::LABEL_SWITCHES,
		switches = switches,
		label_log = data::LABEL_LOG,
		logs = logs,
	);
}

pub fn to_toml(s: &scenario::Scenario) -> String {
	// render switches
	let mut switches = String::from("");
	for (_i, v) in s.switches.iter().enumerate() {
		switches.push_str(&format!(
			"
[[{label_group}.{label_switches}]]
{label_value} = '''
{switch}
'''
",
			label_group = data::LABEL_GROUP,
			label_switches = data::LABEL_SWITCHES,
			label_value = data::LABEL_VALUE,
			switch = v,
		));
	}

	if switches.chars().count() == 0 {
		switches.push_str(&format!(
			"{label_switches} = []",
			label_switches = data::LABEL_SWITCHES,
		));
	}

	// render log
	let mut logs = String::from("");
	for (_i, v) in s.logs.iter().enumerate() {
		switches.push_str(&format!(
			"
[[{label_group}.{label_log}]]
{label_value} = '''
{entry}
'''
",
			label_group = data::LABEL_GROUP,
			label_log = data::LABEL_LOG,
			label_value = data::LABEL_VALUE,
			entry = v,
		));
	}

	if logs.chars().count() == 0 {
		logs.push_str(&format!("{label_log} = []", label_log = data::LABEL_LOG,));
	}

	// render output
	return format!(
		"
[{label_group}]
{label_id} = {id}
{label_verdict} = '{verdict}'
{label_name} = '''
{name}
'''
{label_description} = '''
{description}
'''
{switches}
{logs}
",
		label_group = data::LABEL_GROUP,
		label_id = data::LABEL_ID,
		id = s.id,
		label_verdict = data::LABEL_VERDICT,
		verdict = data::interpret(scenario::conclusion(s)),
		label_name = data::LABEL_NAME,
		name = s.name,
		label_description = data::LABEL_DESCRIPTION,
		description = s.description,
		switches = switches,
		logs = logs,
	);
}
