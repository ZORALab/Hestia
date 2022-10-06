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
const SUITE_NAME: &str = "hestia_number::hestia_bits::s128_length API";

// test conditions

// test values

// test libs
fn assert_output(s: &hestia_testing::Scenario, output: u128) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_128) {
		return output == testlibs_test::VALUE_TYPE_128_BITS_128_COUNT;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_64) {
		return output == testlibs_test::VALUE_TYPE_128_BITS_64_COUNT;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_32) {
		return output == testlibs_test::VALUE_TYPE_128_BITS_32_COUNT;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_16) {
		return output == testlibs_test::VALUE_TYPE_128_BITS_16_COUNT;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_8) {
		return output == testlibs_test::VALUE_TYPE_128_BITS_8_COUNT;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_BITS_0) {
		return output == testlibs_test::VALUE_TYPE_128_BITS_0_COUNT;
	}

	return false;
}

fn test_s128_length_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	let subject: u128 = testlibs_test::create_u128(s);
	let output: u128 = hestia_bits::s128_length(subject);

	hestia_testing::log(s, format!("Given subject: '''\n{}\n'''", subject));
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
hestia_testing_exec!(test_s128_len_bits_0, {
	test_s128_length_algorithm(
		5,
		"\
test hestia_number::hestia_bits::s128_length() is able to process 0-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_0)],
	)
});

hestia_testing_exec!(test_s128_len_bits_8, {
	test_s128_length_algorithm(
		4,
		"\
test hestia_number::hestia_bits::s128_length() is able to process 8-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_8)],
	)
});

hestia_testing_exec!(test_s128_len_bits_16, {
	test_s128_length_algorithm(
		3,
		"\
test hestia_number::hestia_bits::s128_length() is able to process 16-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_16)],
	)
});

hestia_testing_exec!(test_s128_len_bits_32, {
	test_s128_length_algorithm(
		2,
		"\
test hestia_number::hestia_bits::s128_length() is able to process 32-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_32)],
	)
});

hestia_testing_exec!(test_s128_len_bits_64, {
	test_s128_length_algorithm(
		1,
		"\
test hestia_number::hestia_bits::s128_length() is able to process 64-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_64)],
	)
});

hestia_testing_exec!(test_s128_len_bits_128, {
	test_s128_length_algorithm(
		0,
		"\
test hestia_number::hestia_bits::s128_length() is able to process 128-bits value.
"
		.to_string(),
		vec![String::from(testlibs_test::COND_BITS_128)],
	)
});
