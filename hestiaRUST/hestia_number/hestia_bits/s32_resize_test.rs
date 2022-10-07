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
use crate::hestia_error::Error;
use crate::hestia_number::hestia_bits;
use crate::hestia_number::hestia_bits::constants;
use crate::hestia_number::hestia_bits::testlibs_test;
use crate::hestia_testing;

// test suites
const SUITE_NAME: &str = "hestia_number::hestia_bits::s32_resize API";

// test libs
fn assert_error(s: &hestia_testing::Scenario, err: Error) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_32)
		|| hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_22)
		|| hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_16)
		|| hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_12)
		|| hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_8)
		|| hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_5)
		|| hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_0)
	{
		return err == Error::Ok;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_1000) {
		return err == Error::OutOfRange;
	}

	return false;
}

fn assert_output(s: &hestia_testing::Scenario, output: u32) -> bool {
	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_1000) {
		return output == constants::MAX_UINT128 as u32;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_32) {
		if hestia_testing::has_condition(s, testlibs_test::COND_TO_SIGNED) {
			return output == constants::MAX_INT32 as u32;
		} else {
			return output == constants::MAX_UINT32 as u32;
		}
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_22) {
		return output == testlibs_test::VALUE_MASKED_BITS_22 as u32;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_16) {
		if hestia_testing::has_condition(s, testlibs_test::COND_TO_SIGNED) {
			return output == constants::MAX_INT16 as u32;
		} else {
			return output == constants::MAX_UINT16 as u32;
		}
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_12) {
		return output == testlibs_test::VALUE_MASKED_BITS_12 as u32;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_8) {
		if hestia_testing::has_condition(s, testlibs_test::COND_TO_SIGNED) {
			return output == constants::MAX_INT8 as u32;
		} else {
			return output == constants::MAX_UINT8 as u32;
		}
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_5) {
		return output == testlibs_test::VALUE_MASKED_BITS_5 as u32;
	}

	if hestia_testing::has_condition(s, testlibs_test::COND_TO_BITS_0) {
		return output == 0;
	}

	return false;
}

fn test_s32_resize_algorithm(id: u64, desc: String, switches: Vec<String>) {
	// prepare
	let mut s: &mut hestia_testing::Scenario = &mut hestia_testing::new_scenario();
	s.id = id;
	s.name = String::from(SUITE_NAME);
	s.description = desc;
	s.switches = switches;

	// test
	let mut subject: u32 = !0;
	hestia_testing::log(s, format!("Given subject: '0x{:X}'", subject));

	let size: u16 = testlibs_test::create_size(s);
	hestia_testing::log(s, format!("Given size: '{}'", size));

	let sign: bool = testlibs_test::create_sign(s);
	hestia_testing::log(s, format!("Given sign: '{}'", sign));

	let err: Error = hestia_bits::s32_resize(&mut subject, size, sign);
	hestia_testing::log(s, format!("Got subject: '0x{:X}'", subject));
	hestia_testing::log(s, format!("Got error: '{:?}'", err));

	// assert
	hestia_testing::conclude(s, hestia_testing::VERDICT_PASS);
	if !assert_output(s, subject) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	if !assert_error(s, err) {
		hestia_testing::conclude(s, hestia_testing::VERDICT_FAIL);
	}

	// report
	println!("{}", hestia_testing::to_string(s));
	assert!(hestia_testing::conclusion(s) == hestia_testing::VERDICT_PASS);
}

// test suites
hestia_testing_exec!(test_s32_resize_bits_32_to_signed_0, {
	test_s32_resize_algorithm(
		14,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to signed
0-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_0),
			String::from(testlibs_test::COND_TO_SIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_0, {
	test_s32_resize_algorithm(
		13,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
0-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_0),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_signed_5, {
	test_s32_resize_algorithm(
		12,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to signed
5-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_5),
			String::from(testlibs_test::COND_TO_SIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_5, {
	test_s32_resize_algorithm(
		11,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
5-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_5),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_signed_8, {
	test_s32_resize_algorithm(
		10,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
8-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_8),
			String::from(testlibs_test::COND_TO_SIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_8, {
	test_s32_resize_algorithm(
		9,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
8-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_8),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_signed_12, {
	test_s32_resize_algorithm(
		8,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to signed
12-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_12),
			String::from(testlibs_test::COND_TO_SIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_12, {
	test_s32_resize_algorithm(
		7,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
12-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_12),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_signed_16, {
	test_s32_resize_algorithm(
		6,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to signed
16-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_16),
			String::from(testlibs_test::COND_TO_SIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_16, {
	test_s32_resize_algorithm(
		5,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
16-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_16),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_signed_22, {
	test_s32_resize_algorithm(
		4,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to signed
22-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_22),
			String::from(testlibs_test::COND_TO_SIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_22, {
	test_s32_resize_algorithm(
		3,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
22-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_22),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_signed_32, {
	test_s32_resize_algorithm(
		2,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to signed
32-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_32),
			String::from(testlibs_test::COND_TO_SIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_32, {
	test_s32_resize_algorithm(
		1,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
32-bits value.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_32),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});

hestia_testing_exec!(test_s32_resize_bits_32_to_unsigned_1000, {
	test_s32_resize_algorithm(
		0,
		"\
test hestia_number::hestia_bits::s32_resize() is able to process 32-bits value to unsigned
1000-bits value by returning an error and leave the value in-tact.
"
		.to_string(),
		vec![
			String::from(testlibs_test::COND_TO_BITS_1000),
			String::from(testlibs_test::COND_TO_UNSIGNED),
		],
	)
});
