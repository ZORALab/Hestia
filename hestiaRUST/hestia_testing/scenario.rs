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
