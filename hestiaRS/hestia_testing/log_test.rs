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
use crate::hestia_testing::testlibs_test;

// test suites
const SUITE_NAME: &str = "hestia_testing::log API";

// test libs
fn assert_panic(s: &hestia_testing::Scenario) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_PROPER_LOG_DATA) {
		return true;
	}

	return false;
}

fn assert_log(s: &hestia_testing::Scenario, ts: &hestia_testing::Scenario) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_PROPER_LOG_DATA) {
		if ts.logs.iter().any(|i| i == testlibs_test::VALUE_LOG_DATA) {
			return true;
		}
	}

	return false;
}

fn test_log_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	let ts: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();

	let statement = testlibs_test::create_statement(s);
	hestia_testing::log(s, format!("Given string	: {}\n", statement));

	hestia_testing::conclude(s, hestia_testing::VERDICT_PASS);
	if hestia_testing::has_condition(s, testlibs_test::COND_EMPTY_LOG_DATA) {
		println!("{}", hestia_testing::to_string(s));
	}

	hestia_testing::log(ts, statement);

	// assert
	if !assert_panic(&s) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	if !assert_log(&s, &ts) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	// report
	if !hestia_testing::has_condition(s, testlibs_test::COND_EMPTY_LOG_DATA) {
		println!("{}", hestia_testing::to_string(s));
	}
	assert!(hestia_testing::conclusion(s) == hestia_testing::VERDICT_PASS);
}

// test suites
hestia_testing_exec!(log_empty_statement, true, {
	test_log_algorithm(
		1,
		"\
Test Log() is able to panic when an empty string statement is provided.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_EMPTY_LOG_DATA)],
	)
});

hestia_testing_exec!(log_proper_statement, {
	test_log_algorithm(
		0,
		"\
Test Log() is able to process properly when a proper string statement is provided.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_PROPER_LOG_DATA)],
	)
});
