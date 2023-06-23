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

// test suites
const SUITE_NAME: &str = "hestia_testing_exec! API";

// test conditions
const COND_EXPECT_PANIC: &str = "expect panic";
const COND_EXPECT_NO_PANIC: &str = "expect no panic";
const COND_EXPECT_PANIC_VALUE: &str = "expect panic with specific value";

// test values
const VALUE_PANIC: &str = "panic as expected!";

// test libs
fn assert_panics(s: &hestia_testing::Scenario) -> bool {
	if hestia_testing::has_condition(s, COND_EXPECT_PANIC) {
		return false;
	}

	if hestia_testing::has_condition(s, COND_EXPECT_NO_PANIC) {
		return true;
	}

	return false;
}

fn test_exec_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	hestia_testing::conclude(s, hestia_testing::VERDICT_PASS);
	if hestia_testing::has_condition(s, COND_EXPECT_PANIC) {
		println!("{}", hestia_testing::to_string(s));

		if hestia_testing::has_condition(s, COND_EXPECT_PANIC_VALUE) {
			panic!("{}", VALUE_PANIC);
		}

		panic!("{}", hestia_testing::to_string(s));
	}

	// assert
	if !assert_panics(s) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	// report
	println!("{}", hestia_testing::to_string(s));
	assert!(hestia_testing::conclusion(s) == hestia_testing::VERDICT_PASS);
}

// test suites
hestia_testing_exec!(test_exec, {
	test_exec_algorithm(
		2,
		"\
hestia_testing_exec!() is able to handle non-panic test properly.
"
		.to_string(),
		vec![String::from(COND_EXPECT_NO_PANIC)],
	)
});

hestia_testing_exec!(test_exec_panic, true, {
	test_exec_algorithm(
		1,
		"\
hestia_testing_exec!() is able to handle panicking test properly.
"
		.to_string(),
		vec![String::from(COND_EXPECT_PANIC)],
	)
});

hestia_testing_exec!(test_exec_panic_with_expect, true, "panic as expected!", {
	test_exec_algorithm(
		0,
		"\
hestia_testing_exec!() is able to handle panicking test with expected output properly.
"
		.to_string(),
		vec![
			String::from(COND_EXPECT_PANIC),
			String::from(COND_EXPECT_PANIC_VALUE),
		],
	)
});
