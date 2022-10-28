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

package hestiaMATH

func S32_FactorOf(input int32, factor int32) (count uint32, quotient int32, remainder int32) {
	if factor == 0 {
		panic("factor is 0 where the input is divided by 0")
	}

	for count = 0; ; count++ {
		remainder = input % factor
		input /= factor

		if remainder != 0 {
			return count, input, remainder
		}
	}
}

func S64_FactorOf(input int64, factor int64) (count uint64, quotient int64, remainder int64) {
	if factor == 0 {
		panic("factor is 0 where the input is divided by 0")
	}

	for count = 0; ; count++ {
		remainder = input % factor
		input /= factor

		if remainder != 0 {
			return count, input, remainder
		}
	}
}
