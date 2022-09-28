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

func TestHasExecutedAPI(t *testing.T) {
	x := &Scenario{}

	scenarios := []struct {
		subject     *Scenario
		verdict     Verdict
		expectPanic bool
	}{
		{
			subject:     x,
			verdict:     priv_VERDICT_UNKNOWN,
			expectPanic: false,
		}, {
			subject:     x,
			verdict:     VERDICT_PASS,
			expectPanic: false,
		}, {
			subject:     x,
			verdict:     VERDICT_FAIL,
			expectPanic: false,
		}, {
			subject:     x,
			verdict:     VERDICT_SKIP,
			expectPanic: false,
		}, {
			subject:     nil,
			verdict:     priv_VERDICT_UNKNOWN,
			expectPanic: true,
		}, {
			subject:     nil,
			verdict:     VERDICT_PASS,
			expectPanic: true,
		}, {
			subject:     nil,
			verdict:     VERDICT_FAIL,
			expectPanic: true,
		}, {
			subject:     nil,
			verdict:     VERDICT_SKIP,
			expectPanic: true,
		},
	}

	for _, s := range scenarios {
		// prepare
		if s.subject != nil {
			s.subject.verdict = s.verdict
		}

		// execute
		verdict := false
		_out := Exec(func() any {
			verdict = HasExecuted(s.subject)
			return ""
		})
		out, _ := _out.(string)

		// verdict
		if s.expectPanic && out == "" || !s.expectPanic && out != "" {
			t.Errorf("FAILED 1: expected panic '%v' got '%v'", s.expectPanic, out)
		}
		t.Logf("PASSED 1: expected panic '%v' got '%v'", s.expectPanic, out)

		if !s.expectPanic && s.subject != nil {
			t.Logf("LOG 2: verdict is '%d' got '%v'", s.verdict, verdict)

			switch s.verdict {
			case priv_VERDICT_UNKNOWN:
				if verdict {
					t.Errorf("FAILED 2")
				}
			case VERDICT_PASS:
				if !verdict {
					t.Errorf("FAILED 2")
				}
			case VERDICT_FAIL:
				if !verdict {
					t.Errorf("FAILED 2")
				}
			case VERDICT_SKIP:
				if !verdict {
					t.Errorf("FAILED 2")
				}
			}
		}
	}
}
