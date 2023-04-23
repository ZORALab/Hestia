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

// data labels
pub const LABEL_GROUP: &str = "Result";
pub const LABEL_NAME: &str = "Name";
pub const LABEL_ID: &str = "ID";
pub const LABEL_VERDICT: &str = "Verdict";
pub const LABEL_DESCRIPTION: &str = "Description";
pub const LABEL_SWITCHES: &str = "Switches";
pub const LABEL_LOG: &str = "Log";
pub const LABEL_VALUE: &str = "Value";

// errors
pub type Error = u8;

pub const ERROR_OK: Error = 0;
pub const ERROR_BAD_EXCHANGE: Error = 1;

// verdicts
pub type Verdict = u8;

pub const VERDICT_PASS: Verdict = 1;
pub const VERDICT_SKIP: Verdict = 2;
pub const VERDICT_FAIL: Verdict = 3;

pub const STRING_VERDICT_PASS: &str = "PASSED";
pub const STRING_VERDICT_FAIL: &str = "FAILED";
pub const STRING_VERDICT_SKIP: &str = "SKIPPED";
pub const STRING_VERDICT_UNKNOWN: &str = "UNKNOWN";

pub fn interpret(verdict: Verdict) -> String {
	match verdict {
		| VERDICT_PASS => return STRING_VERDICT_PASS.to_string(),
		| VERDICT_FAIL => return STRING_VERDICT_FAIL.to_string(),
		| VERDICT_SKIP => return STRING_VERDICT_SKIP.to_string(),
		| _ => return STRING_VERDICT_UNKNOWN.to_string(),
	}
}
