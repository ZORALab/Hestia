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
pub const MAX_UINT8: u128 = (1 << 8) - 1;
pub const MAX_UINT16: u128 = (1 << 16) - 1;
pub const MAX_UINT32: u128 = (1 << 32) - 1;
pub const MAX_UINT64: u128 = (1 << 64) - 1;
pub const MAX_UINT128: u128 = !0;

pub const MAX_INT8: u128 = (1 << 7) - 1;
pub const MAX_INT16: u128 = (1 << 15) - 1;
pub const MAX_INT32: u128 = (1 << 31) - 1;
pub const MAX_INT64: u128 = (1 << 63) - 1;
pub const MAX_INT128: u128 = (1 << 127) - 1;

pub const MIN_INT8: i128 = -1 << 7;
pub const MIN_INT16: i128 = -1 << 15;
pub const MIN_INT32: i128 = -1 << 31;
pub const MIN_INT64: i128 = -1 << 63;
pub const MIN_INT128: i128 = -1 << 127;
