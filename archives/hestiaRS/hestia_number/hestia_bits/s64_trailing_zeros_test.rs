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
use crate::hestia_number::hestia_bits;
use crate::hestia_number::hestia_bits::testlibs_test;
use crate::hestia_testing;

// test suites
const SUITE_NAME: &str = "hestia_number::hestia_bits::s64_trailing_zeros API";

// test libs
fn assert_output(s: &hestia_testing::Scenario, output: u64) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_64) {
		return output == testlibs_test::VALUE_BITS_64_COUNT as u64 - 1;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_32) {
		return output == testlibs_test::VALUE_BITS_32_COUNT as u64 - 1;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_16) {
		return output == testlibs_test::VALUE_BITS_16_COUNT as u64 - 1;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_8) {
		return output == testlibs_test::VALUE_BITS_8_COUNT as u64 - 1;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_0) {
		return output == testlibs_test::VALUE_BITS_0_COUNT as u64;
	}

	return false;
}

fn test_s64_trailing_zeros_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	let subject: u64 = testlibs_test::create_sample(s) as u64;
	let output: u64 = hestia_bits::s64_trailing_zeros(subject);

	hestia_testing::log(s, format!("Given subject: '{}'", subject));
	hestia_testing::log(s, format!("Got output: '{}'", output));

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
hestia_testing_exec!(test_s64_trailing_zeros_bits_0, {
	test_s64_trailing_zeros_algorithm(
		4,
		"\
test hestia_number::hestia_bits::s64_trailing_zeros() is able to process 0-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_0)],
	)
});

hestia_testing_exec!(test_s64_trailing_zeros_bits_8, {
	test_s64_trailing_zeros_algorithm(
		3,
		"\
test hestia_number::hestia_bits::s64_trailing_zeros() is able to process 8-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_8)],
	)
});

hestia_testing_exec!(test_s64_trailing_zeros_bits_16, {
	test_s64_trailing_zeros_algorithm(
		2,
		"\
test hestia_number::hestia_bits::s64_trailing_zeros() is able to process 16-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_16)],
	)
});

hestia_testing_exec!(test_s64_trailing_zeros_bits_32, {
	test_s64_trailing_zeros_algorithm(
		1,
		"\
test hestia_number::hestia_bits::s64_trailing_zeros() is able to process 32-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_32)],
	)
});

hestia_testing_exec!(test_s64_trailing_zeros_bits_64, {
	test_s64_trailing_zeros_algorithm(
		0,
		"\
test hestia_number::hestia_bits::s64_trailing_zeros() is able to process 64-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_64)],
	)
});
