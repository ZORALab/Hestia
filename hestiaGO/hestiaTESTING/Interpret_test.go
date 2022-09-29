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

import (
	"testing"
)

func TestInterpretAPI(t *testing.T) {
	out := Interpret(VERDICT_PASS)
	t.Logf("Requested: 'VERDICT_PASS' Got: '%v'", out)
	if out != string_PASS {
		t.Errorf("FAILED")
	} else {
		t.Logf("PASSED")
	}

	out = Interpret(VERDICT_FAIL)
	t.Logf("Requested: 'VERDICT_FAIL' Got: '%v'", out)
	if out != string_FAIL {
		t.Errorf("FAILED")
	} else {
		t.Logf("PASSED")
	}

	out = Interpret(VERDICT_SKIP)
	t.Logf("Requested: 'VERDICT_SKIP' Got: '%v'", out)
	if out != string_SKIP {
		t.Errorf("FAILED")
	} else {
		t.Logf("PASSED")
	}

	out = Interpret(100)
	t.Logf("Requested: 'Unknown 100' Got: '%v'", out)
	if out != string_UNKNOWN {
		t.Errorf("FAILED")
	} else {
		t.Logf("PASSED")
	}
}
