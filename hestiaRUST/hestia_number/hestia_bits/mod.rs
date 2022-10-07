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
pub mod constants;
pub use constants::*;

pub mod length;
pub use length::*;

pub mod size;
pub use size::*;

#[cfg(test)]
mod testlibs_test;

#[cfg(test)]
mod s128_length_test;

#[cfg(test)]
mod s64_length_test;

#[cfg(test)]
mod s32_length_test;

#[cfg(test)]
mod s16_length_test;

#[cfg(test)]
mod s8_length_test;

#[cfg(test)]
mod s128_trailing_zeros_test;

#[cfg(test)]
mod s64_trailing_zeros_test;

#[cfg(test)]
mod s32_trailing_zeros_test;

#[cfg(test)]
mod s16_trailing_zeros_test;

#[cfg(test)]
mod s8_trailing_zeros_test;

#[cfg(test)]
mod s128_resize_test;

#[cfg(test)]
mod s64_resize_test;

#[cfg(test)]
mod s32_resize_test;

#[cfg(test)]
mod s16_resize_test;
