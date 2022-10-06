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
const SUITE_NAME: &str = "hestia_testing Interpret API";

// test libs
fn assert_string_verdict(output: &str, expect: &str) -> bool {
	if output.eq(expect) {
		return true;
	}

	return false;
}

fn test_interpret_algorithm(id: u64, desc: String, switches: Vec<String>, expect: &str) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	let verdict = testlibs_test::create_verdict(s);
	let output = hestia_testing::interpret(verdict);
	hestia_testing::log(s, format!("Given verdict	: {}\n", verdict));
	hestia_testing::log(s, format!("Got output	: {}\n", output));

	// assert
	hestia_testing::conclude(s, hestia_testing::VERDICT_PASS);

	if !assert_string_verdict(&output, expect) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	// report
	println!("{}", hestia_testing::to_string(s));
	assert!(hestia_testing::conclusion(s) == hestia_testing::VERDICT_PASS);
}

// test suites
hestia_testing_exec!(interpret_verdict_unknown, {
	test_interpret_algorithm(
		3,
		"\
Test Interpret() is able to work properly when verdict is VERDICT_UNKNOWN.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_UNKNOWN_VERDICT)],
		hestia_testing::STRING_VERDICT_UNKNOWN,
	)
});

hestia_testing_exec!(interpret_verdict_skip, {
	test_interpret_algorithm(
		2,
		"\
Test Interpret() is able to work properly when verdict is VERDICT_SKIP.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_SKIP_VERDICT)],
		hestia_testing::STRING_VERDICT_SKIP,
	)
});

hestia_testing_exec!(interpret_verdict_fail, {
	test_interpret_algorithm(
		1,
		"\
Test Interpret() is able to work properly when verdict is VERDICT_FAIL.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_FAIL_VERDICT)],
		hestia_testing::STRING_VERDICT_FAIL,
	)
});

hestia_testing_exec!(interpret_verdict_pass, {
	test_interpret_algorithm(
		0,
		"\
Test Interpret() is able to work properly when verdict is VERDICT_PASS.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_PROPER_VERDICT)],
		hestia_testing::STRING_VERDICT_PASS,
	)
});
