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
#[macro_export]
macro_rules! hestia_testing_exec {
	($name:ident, $fx:block) => {
		#[test]
		fn $name() {
			$fx
		}
	};
	($name:ident, $panic:expr, $fx:block) => {
		#[test]
		#[should_panic]
		fn $name() {
			$fx
		}
	};
	($name:ident, $panic:expr, $expect:expr, $fx:block) => {
		#[test]
		#[should_panic(expected = $expect)]
		fn $name() {
			$fx
		}
	};
}
