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

pub struct Scenario {
	pub id: u64,
	pub name: String,
	pub description: String,
	pub switches: Vec<String>,
	pub logs: Vec<String>,

	verdict: data::Verdict,
}

pub fn new_scenario() -> Scenario {
	return Scenario {
		id: 0,
		name: String::from(""),
		description: String::from(""),
		switches: Vec::new(),
		logs: Vec::new(),
		verdict: 0,
	};
}

pub fn has_condition(s: &Scenario, condition: &str) -> bool {
	if s.switches.iter().any(|i| i == condition) {
		return true;
	}

	return false;
}

pub fn log(s: &mut Scenario, statement: String) {
	if statement.chars().count() == 0 {
		panic!("calling hestiaTESTING.Logf without providing formatting string!");
	}

	s.logs.push(statement);
}

pub fn conclude(s: &mut Scenario, certification: data::Verdict) -> data::Error {
	match certification {
		| data::VERDICT_PASS => s.verdict = data::VERDICT_PASS,
		| data::VERDICT_FAIL => s.verdict = data::VERDICT_FAIL,
		| data::VERDICT_SKIP => s.verdict = data::VERDICT_SKIP,
		| _ => {
			panic!("calling hestia_testing.conclude an with unknown Verdict!")
		},
	}

	return data::ERROR_OK;
}

pub fn conclusion(s: &Scenario) -> data::Verdict {
	return s.verdict;
}

pub fn has_executed(s: &Scenario) -> bool {
	return s.verdict != 0;
}

pub fn to_string(s: &Scenario) -> String {
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

	// render name
	let mut name: String = String::from("''");
	if !s.name.is_empty() {
		name = String::from(s.name.trim());
	}

	// render description
	let mut description: String = String::from("''");
	if !s.description.is_empty() {
		description = String::from(s.description.trim());
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
		verdict = data::interpret(conclusion(s)),
		label_name = data::LABEL_NAME,
		name = name,
		label_description = data::LABEL_DESCRIPTION,
		description = description,
		label_switches = data::LABEL_SWITCHES,
		switches = switches,
		label_log = data::LABEL_LOG,
		logs = logs,
	);
}

pub fn to_toml(s: &Scenario) -> String {
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
		logs.push_str(&format!(
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
		if s.switches.len() == 0 {
			logs.push_str(&format!(
				"\
{label_log} = []",
				label_log = data::LABEL_LOG,
			));
		} else {
			logs.push_str(&format!(
				"\
[{label_group}]
{label_log} = []",
				label_group = data::LABEL_GROUP,
				label_log = data::LABEL_LOG,
			));
		}
	}

	// render name
	let mut name = String::from("''");
	if s.name.trim().chars().count() > 0 {
		name = format!("'''\n{}\n'''", s.name.trim());
	}

	// render description
	let mut description = String::from("''");
	if s.description.trim().chars().count() > 0 {
		description = format!("'''\n{}\n'''", s.description.trim());
	}

	// render output
	return format!(
		"\
[{label_group}]
{label_id} = {id}
{label_verdict} = '{verdict}'
{label_name} = {name}
{label_description} = {description}
{switches}
{logs}
",
		label_group = data::LABEL_GROUP,
		label_id = data::LABEL_ID,
		id = s.id,
		label_verdict = data::LABEL_VERDICT,
		verdict = data::interpret(conclusion(s)),
		label_name = data::LABEL_NAME,
		name = name,
		label_description = data::LABEL_DESCRIPTION,
		description = description,
		switches = switches,
		logs = logs,
	);
}

#[cfg(test)]
pub fn new_scenario_sample_v1(verdict: data::Verdict) -> Scenario {
	let mut s: Scenario = new_scenario();
	s.verdict = verdict;
	return s;
}
