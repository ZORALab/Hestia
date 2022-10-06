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
const SUITE_NAME: &str = "hestia_testing::has_condition API";

// test conditions

// test values

// test libs
fn assert_output(s: &hestia_testing::Scenario, output: bool) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_PROPER_SWITCHES) {
		return output;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_EMPTY_SWITCHES) {
		return !output;
	}

	return false;
}

fn test_has_condition_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	let ts: &mut hestia_testing::Scenario = &mut testlibs_test::create_scenario(s);
	let output: bool = hestia_testing::has_condition(ts, testlibs_test::VALUE_SWITCH_2);

	hestia_testing::log(s, format!("Got output: '''\n{}\n'''", output));

	// assert
	hestia_testing::conclude(s, hestia_testing::VERDICT_PASS);
	if !assert_output(s, output) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	// report
	println!("{}", hestia_testing::to_string(s));
	assert!(hestia_testing::conclusion(s) == hestia_testing::VERDICT_PASS);
}

// test suites
hestia_testing_exec!(test_has_condition_empty, {
	test_has_condition_algorithm(
		0,
		"\
test hestia_testing::has_condition() is able to work properly with empty configurations.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_PROPER_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_EMPTY_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_PROPER_VERDICT),
		],
	)
});

hestia_testing_exec!(test_has_condition_proper, {
	test_has_condition_algorithm(
		0,
		"\
test hestia_testing::has_condition() is able to work properly with proper configurations.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_PROPER_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_PROPER_VERDICT),
		],
	)
});
