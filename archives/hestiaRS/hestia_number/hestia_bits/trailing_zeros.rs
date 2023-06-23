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
pub fn s128_trailing_zeros(x: u128) -> u128 {
	if x == 0 {
		return 0;
	}

	let mut count: u128;
	let mut buffer: u128 = x;

	count = 0;
	while (buffer & 1) == 0 {
		buffer = buffer >> 1;
		count = count + 1;
	}

	return count;
}

pub fn s64_trailing_zeros(x: u64) -> u64 {
	if x == 0 {
		return 0;
	}

	let mut count: u64;
	let mut buffer: u64 = x;

	count = 0;
	while (buffer & 1) == 0 {
		buffer = buffer >> 1;
		count = count + 1;
	}

	return count;
}

pub fn s32_trailing_zeros(x: u32) -> u32 {
	if x == 0 {
		return 0;
	}

	let mut count: u32;
	let mut buffer: u32 = x;

	count = 0;
	while (buffer & 1) == 0 {
		buffer = buffer >> 1;
		count = count + 1;
	}

	return count;
}

pub fn s16_trailing_zeros(x: u16) -> u16 {
	if x == 0 {
		return 0;
	}

	let mut count: u16;
	let mut buffer: u16 = x;

	count = 0;
	while (buffer & 1) == 0 {
		buffer = buffer >> 1;
		count = count + 1;
	}

	return count;
}

pub fn s8_trailing_zeros(x: u8) -> u8 {
	if x == 0 {
		return 0;
	}

	let mut count: u8;
	let mut buffer: u8 = x;

	count = 0;
	while (buffer & 1) == 0 {
		buffer = buffer >> 1;
		count = count + 1;
	}

	return count;
}
