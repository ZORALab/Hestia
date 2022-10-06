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
const SUITE_NAME: &str = "hestia_testing::to_string API";

// test conditions

// test values

// test libs
fn assert_output(_s: &hestia_testing::Scenario, output: String) -> bool {
	if output.chars().count() > 0 {
		return true;
	}

	return false;
}

fn test_to_string_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	let ts: &mut hestia_testing::Scenario = &mut testlibs_test::create_scenario(s);
	let output: String = hestia_testing::to_string(ts);

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
hestia_testing_exec!(test_to_string_proper_unknown, {
	test_to_string_algorithm(
		8,
		"\
test hestia_testing::to_string() is able to generate output with proper values + unknown verdict.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_PROPER_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_UNKNOWN_VERDICT),
		],
	)
});

hestia_testing_exec!(test_to_string_proper_skip, {
	test_to_string_algorithm(
		7,
		"\
test hestia_testing::to_string() is able to generate output with proper values + skip verdict.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_PROPER_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_SKIP_VERDICT),
		],
	)
});

hestia_testing_exec!(test_to_string_proper_fail, {
	test_to_string_algorithm(
		6,
		"\
test hestia_testing::to_string() is able to generate output with proper values + fail verdict.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_PROPER_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_FAIL_VERDICT),
		],
	)
});

hestia_testing_exec!(test_to_string_empty_logs, {
	test_to_string_algorithm(
		5,
		"\
test hestia_testing::to_string() is able to generate output with empty logs.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_PROPER_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_EMPTY_LOGS),
			String::from(testlibs_test::COND_PROPER_VERDICT),
		],
	)
});

hestia_testing_exec!(test_to_string_empty_switches, {
	test_to_string_algorithm(
		4,
		"\
test hestia_testing::to_string() is able to generate output with empty switches.
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

hestia_testing_exec!(test_to_string_empty_description, {
	test_to_string_algorithm(
		3,
		"\
test hestia_testing::to_string() is able to generate output with empty description.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_ZERO_ID),
			String::from(testlibs_test::COND_EMPTY_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_PROPER_VERDICT),
		],
	)
});

hestia_testing_exec!(test_to_string_zero_id, {
	test_to_string_algorithm(
		2,
		"\
test hestia_testing::to_string() is able to generate output with empty ID.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_PROPER_NAME),
			String::from(testlibs_test::COND_ZERO_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_PROPER_VERDICT),
		],
	)
});

hestia_testing_exec!(test_to_string_empty_name, {
	test_to_string_algorithm(
		1,
		"\
test hestia_testing::to_string() is able to generate output with empty name.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_EMPTY_NAME),
			String::from(testlibs_test::COND_PROPER_ID),
			String::from(testlibs_test::COND_PROPER_DESC),
			String::from(testlibs_test::COND_PROPER_SWITCHES),
			String::from(testlibs_test::COND_PROPER_LOGS),
			String::from(testlibs_test::COND_PROPER_VERDICT),
		],
	)
});

hestia_testing_exec!(test_to_string_proper, {
	test_to_string_algorithm(
		0,
		"\
test hestia_testing::to_string() is able to generate output with proper values.
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
