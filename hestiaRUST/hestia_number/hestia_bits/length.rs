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

use crate::hestia_number::hestia_bits::constants;

pub fn s128_length(x: u128) -> u128 {
	let mut i: u128;
	let mut a: u128;
	let mut b: u128;
	let mut c: u128;

	i = 0;
	a = x;
	b = 0;
	while i <= constants::MAX_UINT128 {
		a |= a >> i;
		c = a ^ (a >> 1);

		if b == c {
			// max bit reached
			break;
		}

		b = c;
		i = i + 1;
	}

	a = 0;
	while b != 0 {
		a = a + 1;
		b = b >> 1;
	}

	return a;
}

pub fn s64_length(x: u64) -> u64 {
	let mut i: u64;
	let mut a: u64;
	let mut b: u64;
	let mut c: u64;

	i = 0;
	a = x;
	b = 0;
	while i <= constants::MAX_UINT64 as u64 {
		a |= a >> i;
		c = a ^ (a >> 1);

		if b == c {
			// max bit reached
			break;
		}

		b = c;
		i = i + 1;
	}

	a = 0;
	while b != 0 {
		a = a + 1;
		b = b >> 1;
	}

	return a;
}

pub fn s32_length(x: u32) -> u32 {
	let mut i: u32;
	let mut a: u32;
	let mut b: u32;
	let mut c: u32;

	i = 0;
	a = x;
	b = 0;
	while i <= constants::MAX_UINT32 as u32 {
		a |= a >> i;
		c = a ^ (a >> 1);

		if b == c {
			// max bit reached
			break;
		}

		b = c;
		i = i + 1;
	}

	a = 0;
	while b != 0 {
		a = a + 1;
		b = b >> 1;
	}

	return a;
}

pub fn s16_length(x: u16) -> u16 {
	let mut i: u16;
	let mut a: u16;
	let mut b: u16;
	let mut c: u16;

	i = 0;
	a = x;
	b = 0;
	while i <= constants::MAX_UINT16 as u16 {
		a |= a >> i;
		c = a ^ (a >> 1);

		if b == c {
			// max bit reached
			break;
		}

		b = c;
		i = i + 1;
	}

	a = 0;
	while b != 0 {
		a = a + 1;
		b = b >> 1;
	}

	return a;
}

pub fn s8_length(x: u8) -> u8 {
	let mut i: u8;
	let mut a: u8;
	let mut b: u8;
	let mut c: u8;

	i = 0;
	a = x;
	b = 0;
	while i <= constants::MAX_UINT8 as u8 {
		a |= a >> i;
		c = a ^ (a >> 1);

		if b == c {
			// max bit reached
			break;
		}

		b = c;
		i = i + 1;
	}

	a = 0;
	while b != 0 {
		a = a + 1;
		b = b >> 1;
	}

	return a;
}
