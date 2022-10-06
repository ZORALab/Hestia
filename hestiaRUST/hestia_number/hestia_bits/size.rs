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
use crate::hestia_number::hestia_bits::constants;

pub fn s128_resize(input: &mut u128, size: u16, with_sign: bool) -> Error {
	let mask: u128;

	if size > 128 {
		return Error::OutOfRange;
	} else if size == 0 {
		*input = 0;
		return Error::Ok;
	} else if size <= 8 {
		if with_sign {
			mask = constants::MAX_INT8;
		} else {
			mask = constants::MAX_UINT8;
		}
	} else if size <= 16 {
		if with_sign {
			mask = constants::MAX_INT16;
		} else {
			mask = constants::MAX_UINT16;
		}
	} else if size <= 32 {
		if with_sign {
			mask = constants::MAX_INT32;
		} else {
			mask = constants::MAX_UINT32;
		}
	} else if size <= 64 {
		if with_sign {
			mask = constants::MAX_INT64;
		} else {
			mask = constants::MAX_UINT64;
		}
	} else {
		if with_sign {
			mask = constants::MAX_INT128;
		} else {
			mask = constants::MAX_UINT128;
		}
	}

	*input = *input & mask;

	return Error::Ok;
}

pub fn s64_resize(input: &mut u64, size: u16, with_sign: bool) -> Error {
	let mask: u64;

	if size > 64 {
		return Error::OutOfRange;
	} else if size == 0 {
		*input = 0;
		return Error::Ok;
	} else if size <= 8 {
		if with_sign {
			mask = constants::MAX_INT8 as u64;
		} else {
			mask = constants::MAX_UINT8 as u64;
		}
	} else if size <= 16 {
		if with_sign {
			mask = constants::MAX_INT16 as u64;
		} else {
			mask = constants::MAX_UINT16 as u64;
		}
	} else if size <= 32 {
		if with_sign {
			mask = constants::MAX_INT32 as u64;
		} else {
			mask = constants::MAX_UINT32 as u64;
		}
	} else {
		if with_sign {
			mask = constants::MAX_INT64 as u64;
		} else {
			mask = constants::MAX_UINT64 as u64;
		}
	}

	*input = *input & mask;

	return Error::Ok;
}

pub fn s32_resize(input: &mut u32, size: u16, with_sign: bool) -> Error {
	let mask: u32;

	if size > 32 {
		return Error::OutOfRange;
	} else if size == 0 {
		*input = 0;
		return Error::Ok;
	} else if size <= 8 {
		if with_sign {
			mask = constants::MAX_INT8 as u32;
		} else {
			mask = constants::MAX_UINT8 as u32;
		}
	} else if size <= 16 {
		if with_sign {
			mask = constants::MAX_INT16 as u32;
		} else {
			mask = constants::MAX_UINT16 as u32;
		}
	} else {
		if with_sign {
			mask = constants::MAX_INT32 as u32;
		} else {
			mask = constants::MAX_UINT32 as u32;
		}
	}

	*input = *input & mask;

	return Error::Ok;
}

pub fn s16_resize(input: &mut u16, size: u16, with_sign: bool) -> Error {
	let mask: u16;

	if size > 16 {
		return Error::OutOfRange;
	} else if size == 0 {
		*input = 0;
		return Error::Ok;
	} else if size <= 8 {
		if with_sign {
			mask = constants::MAX_INT8 as u16;
		} else {
			mask = constants::MAX_UINT8 as u16;
		}
	} else {
		if with_sign {
			mask = constants::MAX_INT16 as u16;
		} else {
			mask = constants::MAX_UINT16 as u16;
		}
	}

	*input = *input & mask;

	return Error::Ok;
}

pub fn s8_resize(input: &mut u8, size: u16, with_sign: bool) -> Error {
	let mask: u8;

	if size > 8 {
		return Error::OutOfRange;
	} else if size == 0 {
		*input = 0;
		return Error::Ok;
	} else {
		if with_sign {
			mask = constants::MAX_INT8 as u8;
		} else {
			mask = constants::MAX_UINT8 as u8;
		}
	}

	*input = *input & mask;

	return Error::Ok;
}
