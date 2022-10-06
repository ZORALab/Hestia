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

// test values
pub const VALUE_TYPE_128_BITS_128: u128 = 1 << 127;
pub const VALUE_TYPE_128_BITS_128_COUNT: u128 = 128;

pub const VALUE_TYPE_128_BITS_64: u128 = 1 << 63;
pub const VALUE_TYPE_128_BITS_64_COUNT: u128 = 64;
pub const VALUE_TYPE_64_BITS_64: u64 = 1 << 63;
pub const VALUE_TYPE_64_BITS_64_COUNT: u64 = 64;

pub const VALUE_TYPE_128_BITS_32: u128 = 1 << 31;
pub const VALUE_TYPE_128_BITS_32_COUNT: u128 = 32;
pub const VALUE_TYPE_64_BITS_32: u64 = 1 << 31;
pub const VALUE_TYPE_64_BITS_32_COUNT: u64 = 32;
pub const VALUE_TYPE_32_BITS_32: u32 = 1 << 31;
pub const VALUE_TYPE_32_BITS_32_COUNT: u32 = 32;

pub const VALUE_TYPE_128_BITS_16: u128 = 1 << 15;
pub const VALUE_TYPE_128_BITS_16_COUNT: u128 = 16;
pub const VALUE_TYPE_64_BITS_16: u64 = 1 << 15;
pub const VALUE_TYPE_64_BITS_16_COUNT: u64 = 16;
pub const VALUE_TYPE_32_BITS_16: u32 = 1 << 15;
pub const VALUE_TYPE_32_BITS_16_COUNT: u32 = 16;
pub const VALUE_TYPE_16_BITS_16: u16 = 1 << 15;
pub const VALUE_TYPE_16_BITS_16_COUNT: u16 = 16;

pub const VALUE_TYPE_128_BITS_8: u128 = 1 << 7;
pub const VALUE_TYPE_128_BITS_8_COUNT: u128 = 8;
pub const VALUE_TYPE_64_BITS_8: u64 = 1 << 7;
pub const VALUE_TYPE_64_BITS_8_COUNT: u64 = 8;
pub const VALUE_TYPE_32_BITS_8: u32 = 1 << 7;
pub const VALUE_TYPE_32_BITS_8_COUNT: u32 = 8;
pub const VALUE_TYPE_16_BITS_8: u16 = 1 << 7;
pub const VALUE_TYPE_16_BITS_8_COUNT: u16 = 8;
pub const VALUE_TYPE_8_BITS_8: u8 = 1 << 7;
pub const VALUE_TYPE_8_BITS_8_COUNT: u8 = 8;

pub const VALUE_TYPE_128_BITS_0: u128 = 1 << 0;
pub const VALUE_TYPE_128_BITS_0_COUNT: u128 = 1;
pub const VALUE_TYPE_64_BITS_0: u64 = 1 << 0;
pub const VALUE_TYPE_64_BITS_0_COUNT: u64 = 1;
pub const VALUE_TYPE_32_BITS_0: u32 = 1 << 0;
pub const VALUE_TYPE_32_BITS_0_COUNT: u32 = 1;
pub const VALUE_TYPE_16_BITS_0: u16 = 1 << 0;
pub const VALUE_TYPE_16_BITS_0_COUNT: u16 = 1;
pub const VALUE_TYPE_8_BITS_0: u8 = 1 << 0;
pub const VALUE_TYPE_8_BITS_0_COUNT: u8 = 1;

pub fn create_u128(s: &hestia_testing::Scenario) -> u128 {
	if hestia_testing::has_condition(s, COND_BITS_128) {
		return VALUE_TYPE_128_BITS_128;
	}

	if hestia_testing::has_condition(s, COND_BITS_64) {
		return VALUE_TYPE_128_BITS_64;
	}

	if hestia_testing::has_condition(s, COND_BITS_32) {
		return VALUE_TYPE_128_BITS_32;
	}

	if hestia_testing::has_condition(s, COND_BITS_16) {
		return VALUE_TYPE_128_BITS_16;
	}

	if hestia_testing::has_condition(s, COND_BITS_8) {
		return VALUE_TYPE_128_BITS_8;
	}

	if hestia_testing::has_condition(s, COND_BITS_0) {
		return VALUE_TYPE_128_BITS_0;
	}

	return 0;
}

pub fn create_u64(s: &hestia_testing::Scenario) -> u64 {
	if hestia_testing::has_condition(s, COND_BITS_64) {
		return VALUE_TYPE_64_BITS_64;
	}

	if hestia_testing::has_condition(s, COND_BITS_32) {
		return VALUE_TYPE_64_BITS_32;
	}

	if hestia_testing::has_condition(s, COND_BITS_16) {
		return VALUE_TYPE_64_BITS_16;
	}

	if hestia_testing::has_condition(s, COND_BITS_8) {
		return VALUE_TYPE_64_BITS_8;
	}

	if hestia_testing::has_condition(s, COND_BITS_0) {
		return VALUE_TYPE_64_BITS_0;
	}

	return 0;
}

pub fn create_u32(s: &hestia_testing::Scenario) -> u32 {
	if hestia_testing::has_condition(s, COND_BITS_32) {
		return VALUE_TYPE_32_BITS_32;
	}

	if hestia_testing::has_condition(s, COND_BITS_16) {
		return VALUE_TYPE_32_BITS_16;
	}

	if hestia_testing::has_condition(s, COND_BITS_8) {
		return VALUE_TYPE_32_BITS_8;
	}

	if hestia_testing::has_condition(s, COND_BITS_0) {
		return VALUE_TYPE_32_BITS_0;
	}

	return 0;
}

pub fn create_u16(s: &hestia_testing::Scenario) -> u16 {
	if hestia_testing::has_condition(s, COND_BITS_16) {
		return VALUE_TYPE_16_BITS_16;
	}

	if hestia_testing::has_condition(s, COND_BITS_8) {
		return VALUE_TYPE_16_BITS_8;
	}

	if hestia_testing::has_condition(s, COND_BITS_0) {
		return VALUE_TYPE_16_BITS_0;
	}

	return 0;
}

pub fn create_u8(s: &hestia_testing::Scenario) -> u8 {
	if hestia_testing::has_condition(s, COND_BITS_8) {
		return VALUE_TYPE_8_BITS_8;
	}

	if hestia_testing::has_condition(s, COND_BITS_0) {
		return VALUE_TYPE_8_BITS_0;
	}

	return 0;
}
