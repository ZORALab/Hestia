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

package hestiaTESTING

const (
	DATA_LABEL_GROUP       = "Result"
	DATA_LABEL_NAME        = "Name"
	DATA_LABEL_ID          = "ID"
	DATA_LABEL_VERDICT     = "Verdict"
	DATA_LABEL_DESCRIPTION = "Description"
	DATA_LABEL_SWITCHES    = "Switches"
	DATA_LABEL_LOG         = "Log"

	DATA_LABEL_VALUE = "Value"
)

type Error uint8

const (
	ERROR_OK Error = iota
	ERROR_BAD_EXCHANGE
)

type Verdict uint8

const (
	priv_VERDICT_UNKNOWN Verdict = iota
	VERDICT_PASS
	VERDICT_SKIP
	VERDICT_FAIL
)

const (
	string_PASS    = "PASSED"
	string_FAIL    = "FAILED"
	string_SKIP    = "SKIPPED"
	string_UNKNOWN = "UNKNOWN"
)
