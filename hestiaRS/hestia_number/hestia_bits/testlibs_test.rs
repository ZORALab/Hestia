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
pub const COND_BITS_128: &str = "provide 128-bits value";
pub const COND_BITS_64: &str = "provide 64-bits value";
pub const COND_BITS_32: &str = "provide 32-bits value";
pub const COND_BITS_16: &str = "provide 16-bits value";
pub const COND_BITS_8: &str = "provide 8-bits value";
pub const COND_BITS_0: &str = "provide 0-bit value";

pub const COND_TO_BITS_1000: &str = "convert to 1000-bits value";
pub const COND_TO_BITS_127: &str = "convert to 127-bits value";
pub const COND_TO_BITS_72: &str = "convert to 72-bits value";
pub const COND_TO_BITS_64: &str = "convert to 64-bits value";
pub const COND_TO_BITS_35: &str = "convert to 35-bits value";
pub const COND_TO_BITS_32: &str = "convert to 32-bits value";
pub const COND_TO_BITS_22: &str = "convert to 22-bits value";
pub const COND_TO_BITS_16: &str = "convert to 16-bits value";
pub const COND_TO_BITS_12: &str = "convert to 12-bits value";
pub const COND_TO_BITS_8: &str = "convert to 8-bits value";
pub const COND_TO_BITS_5: &str = "convert to 5-bits value";
pub const COND_TO_BITS_0: &str = "convert to 0-bits value";

pub const COND_TO_UNSIGNED: &str = "convert to unsigned value";
pub const COND_TO_SIGNED: &str = "convert to signed value";

// test values
pub const VALUE_MASKED_BITS_72: u128 = 0xFFFFFFFFFFFFFFFFFF;
pub const VALUE_MASKED_BITS_35: u128 = 0x7FFFFFFFF;
pub const VALUE_MASKED_BITS_22: u128 = 0x3FFFFF;
pub const VALUE_MASKED_BITS_12: u128 = 0xFFF;
pub const VALUE_MASKED_BITS_5: u128 = 0x1F;

pub const VALUE_BITS_128_COUNT: u128 = 128;
pub const VALUE_BITS_64_COUNT: u128 = 64;
pub const VALUE_BITS_32_COUNT: u128 = 32;
pub const VALUE_BITS_16_COUNT: u128 = 16;
pub const VALUE_BITS_8_COUNT: u128 = 8;
pub const VALUE_BITS_0_COUNT: u128 = 0;

pub fn create_sign(s: &hestia_testing::Scenario) -> bool {
	if hestia_testing::has_condition(s, COND_TO_SIGNED) {
		return true;
	}

	if hestia_testing::has_condition(s, COND_TO_UNSIGNED) {
		return false;
	}

	return false;
}

pub fn create_size(s: &hestia_testing::Scenario) -> u16 {
	if hestia_testing::has_condition(s, COND_TO_BITS_1000) {
		return 1000;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_127) {
		return 127;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_72) {
		return 72;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_64) {
		return 64;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_35) {
		return 35;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_32) {
		return 32;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_22) {
		return 22;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_16) {
		return 16;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_12) {
		return 12;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_8) {
		return 8;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_5) {
		return 5;
	}

	if hestia_testing::has_condition(s, COND_TO_BITS_0) {
		return 0;
	}

	return 0;
}

pub fn create_sample(s: &hestia_testing::Scenario) -> u128 {
	if hestia_testing::has_condition(s, COND_BITS_128) {
		return 1 << (128 - 1);
	}

	if hestia_testing::has_condition(s, COND_BITS_64) {
		return 1 << (64 - 1);
	}

	if hestia_testing::has_condition(s, COND_BITS_32) {
		return 1 << (32 - 1);
	}

	if hestia_testing::has_condition(s, COND_BITS_16) {
		return 1 << (16 - 1);
	}

	if hestia_testing::has_condition(s, COND_BITS_8) {
		return 1 << (8 - 1);
	}

	if hestia_testing::has_condition(s, COND_BITS_0) {
		return 0;
	}

	return 0;
}
