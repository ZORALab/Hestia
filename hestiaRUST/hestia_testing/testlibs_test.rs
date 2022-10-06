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

use crate::hestia_testing;

// test conditions
pub const COND_PROPER_VERDICT: &str = "configure VERDICT_PASS verdict";
pub const COND_FAIL_VERDICT: &str = "configure VERDICT_FAIL verdict";
pub const COND_SKIP_VERDICT: &str = "configure VERDICT_SKIP verdict";
pub const COND_UNKNOWN_VERDICT: &str = "configure VERDICT_UNKNOWN verdict";
pub const COND_INITIAL_VERDICT: &str = "configure initial (0) verdict";

pub const COND_PROPER_NAME: &str = "configure proper Scenario name";
pub const COND_EMPTY_NAME: &str = "configure empty Scenario name";

pub const COND_PROPER_ID: &str = "configure proper Scenario id";
pub const COND_ZERO_ID: &str = "configure zero Scenario id";

pub const COND_PROPER_DESC: &str = "configure proper Scenario description";
pub const COND_EMPTY_DESC: &str = "configure empty Scenario description";

pub const COND_PROPER_SWITCHES: &str = "configure proper Scenario switches";
pub const COND_EMPTY_SWITCHES: &str = "configure empty Scenario switches";

pub const COND_PROPER_LOGS: &str = "configure proper Scenario logs";
pub const COND_EMPTY_LOGS: &str = "configure empty Scenario logs";

pub const COND_PROPER_LOG_DATA: &str = "configure proper log data";
pub const COND_EMPTY_LOG_DATA: &str = "configure empty log data";

// test values
pub const VALUE_UNKNOWN_VERDICT: hestia_testing::Verdict = 100;
pub const VALUE_NAME: &str = "test scenario";
pub const VALUE_DESCRIPTION: &str = "it is a test scenario.";
pub const VALUE_ID: u64 = 123;
pub const VALUE_SWITCH_1: &str = "switch data 1";
pub const VALUE_SWITCH_2: &str = "switch data 2";
pub const VALUE_LOG_1: &str = "log data 1";
pub const VALUE_LOG_2: &str = "log data 2";
pub const VALUE_LOG_DATA: &str = "target log data";

// factory functions
pub fn create_statement(s: &hestia_testing::Scenario) -> String {
	if hestia_testing::has_condition(s, COND_PROPER_LOG_DATA) {
		return VALUE_LOG_DATA.to_string();
	}

	if hestia_testing::has_condition(s, COND_EMPTY_LOG_DATA) {
		return "".to_string();
	}

	return "".to_string();
}

pub fn create_verdict(s: &hestia_testing::Scenario) -> hestia_testing::Verdict {
	if hestia_testing::has_condition(s, COND_PROPER_VERDICT) {
		return hestia_testing::VERDICT_PASS;
	}

	if hestia_testing::has_condition(s, COND_FAIL_VERDICT) {
		return hestia_testing::VERDICT_FAIL;
	}

	if hestia_testing::has_condition(s, COND_SKIP_VERDICT) {
		return hestia_testing::VERDICT_SKIP;
	}

	if hestia_testing::has_condition(s, COND_UNKNOWN_VERDICT) {
		return VALUE_UNKNOWN_VERDICT;
	}

	if hestia_testing::has_condition(s, COND_INITIAL_VERDICT) {
		return 0;
	}

	return 0;
}

pub fn configure_name(s: &hestia_testing::Scenario, ts: &mut hestia_testing::Scenario) {
	if hestia_testing::has_condition(s, COND_PROPER_NAME) {
		ts.name = VALUE_NAME.to_string();
	}

	if hestia_testing::has_condition(s, COND_EMPTY_NAME) {
		ts.name = "".to_string();
	}
}

pub fn configure_id(s: &hestia_testing::Scenario, ts: &mut hestia_testing::Scenario) {
	if hestia_testing::has_condition(s, COND_PROPER_ID) {
		ts.id = VALUE_ID;
	}

	if hestia_testing::has_condition(s, COND_ZERO_ID) {
		ts.id = 0;
	}
}

pub fn configure_description(s: &hestia_testing::Scenario, ts: &mut hestia_testing::Scenario) {
	if hestia_testing::has_condition(s, COND_PROPER_DESC) {
		ts.description = VALUE_DESCRIPTION.to_string();
	}

	if hestia_testing::has_condition(s, COND_EMPTY_DESC) {
		ts.description = "".to_string();
	}
}

pub fn configure_switches(s: &hestia_testing::Scenario, ts: &mut hestia_testing::Scenario) {
	if hestia_testing::has_condition(s, COND_PROPER_SWITCHES) {
		ts.switches = vec![VALUE_SWITCH_1.to_string(), VALUE_SWITCH_2.to_string()];
	}

	if hestia_testing::has_condition(s, COND_EMPTY_SWITCHES) {
		ts.switches = Vec::new();
	}
}

pub fn configure_logs(s: &hestia_testing::Scenario, ts: &mut hestia_testing::Scenario) {
	if hestia_testing::has_condition(s, COND_PROPER_LOGS) {
		ts.logs = vec![VALUE_LOG_1.to_string(), VALUE_LOG_2.to_string()];
	}

	if hestia_testing::has_condition(s, COND_EMPTY_LOGS) {
		ts.logs = Vec::new();
	}
}

pub fn create_scenario(s: &hestia_testing::Scenario) -> hestia_testing::Scenario {
	let verdict: hestia_testing::Verdict = create_verdict(s);
	let mut ts: hestia_testing::Scenario = hestia_testing::new_scenario_sample_v1(verdict);

	configure_name(s, &mut ts);
	configure_id(s, &mut ts);
	configure_description(s, &mut ts);
	configure_switches(s, &mut ts);
	configure_logs(s, &mut ts);

	return ts;
}
