// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 Caleb Spare <cespare@gmail.com>
// Copyright 2018 Ulf Adams <ulfjack@google.com>
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

macro_rules! factor_of {
	($name:ident, $uType:ty, $iType:ty) => {
		pub fn $name(input: $iType, factor: $iType) -> ($uType, $iType, $iType) {
			if factor == 0 {
				panic!("factor is 0 where the input is divided by 0");
			}

			let mut count: $uType = 0;
			let mut quotient: $iType = input;
			let mut remainder: $iType;
			loop {
				remainder = quotient % factor;
				quotient /= factor;

				if remainder != 0 {
					return (count, quotient, remainder);
				}

				count += 1;
			}
		}
	};
}
factor_of!(s32_factor_of, u32, i32);
factor_of!(s64_factor_of, u64, i64);
