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
const SUITE_NAME: &str = "hestia_testing::conclude API";

// test conditions

// test values

// test libs
fn assert_panic(s: &hestia_testing::Scenario) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_PROPER_VERDICT) {
		return true;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_FAIL_VERDICT) {
		return true;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_SKIP_VERDICT) {
		return true;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_UNKNOWN_VERDICT) {
		return false;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_INITIAL_VERDICT) {
		return false;
	}

	return false;
}

fn assert_output(s: &hestia_testing::Scenario, output: hestia_testing::Error) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_PROPER_VERDICT) {
		return output == hestia_testing::ERROR_OK;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_FAIL_VERDICT) {
		return output == hestia_testing::ERROR_OK;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_SKIP_VERDICT) {
		return output == hestia_testing::ERROR_OK;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_UNKNOWN_VERDICT) {
		return output != hestia_testing::ERROR_OK;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_INITIAL_VERDICT) {
		return output != hestia_testing::ERROR_OK;
	}

	return false;
}

fn test_conclude_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;
	hestia_testing::conclude(s, hestia_testing::VERDICT_PASS);

	// test
	let ts: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	let verdict: hestia_testing::Verdict = testlibs_test::create_verdict(s);
	hestia_testing::log(s, format!("Given verdict	: '''\n{}\n'''", verdict));

	if hestia_testing::has_condition(s, testlibs_test::COND_UNKNOWN_VERDICT)
		|| hestia_testing::has_condition(s, testlibs_test::COND_INITIAL_VERDICT)
	{
		println!("{}", hestia_testing::to_string(s));
	}
	let output: hestia_testing::Error = hestia_testing::conclude(ts, verdict);

	hestia_testing::log(s, format!("Got output	: '''\n{}\n'''", output));

	// assert
	if !assert_panic(s) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	if !assert_output(s, output) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	// report
	if !(hestia_testing::has_condition(s, testlibs_test::COND_UNKNOWN_VERDICT)
		|| hestia_testing::has_condition(s, testlibs_test::COND_INITIAL_VERDICT))
	{
		println!("{}", hestia_testing::to_string(s));
	}
	assert!(hestia_testing::conclusion(s) == hestia_testing::VERDICT_PASS);
}

// test suites
hestia_testing_exec!(test_conclude_initial_verdict, true, {
	test_conclude_algorithm(
		4,
		"\
test hestia_testing::conclude() is able to panic with initial (0) verdict value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_INITIAL_VERDICT)],
	)
});

hestia_testing_exec!(test_conclude_unknown_verdict, true, {
	test_conclude_algorithm(
		3,
		"\
test hestia_testing::conclude() is able to panic with unknown verdict value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_UNKNOWN_VERDICT)],
	)
});

hestia_testing_exec!(test_conclude_skip_verdict, {
	test_conclude_algorithm(
		2,
		"\
test hestia_testing::conclude() is able to run with VERDICT_SKIP.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_SKIP_VERDICT)],
	)
});

hestia_testing_exec!(test_conclude_fail_verdict, {
	test_conclude_algorithm(
		1,
		"\
test hestia_testing::conclude() is able to run with VERDICT_FAIL.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_FAIL_VERDICT)],
	)
});

hestia_testing_exec!(test_conclude_proper_verdict, {
	test_conclude_algorithm(
		0,
		"\
test hestia_testing::conclude() is able to run with VERDICT_PASS.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_PROPER_VERDICT)],
	)
});
