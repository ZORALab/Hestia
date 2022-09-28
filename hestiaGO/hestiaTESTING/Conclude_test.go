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

func TestConcludeAPI(t *testing.T) {
	x := &Scenario{}

	scenarios := []struct {
		skip        func()
		fail        func()
		subject     *Scenario
		verdict     Verdict
		expectPanic bool
	}{
		{
			skip:        func() {},
			fail:        func() {},
			subject:     x,
			verdict:     priv_VERDICT_UNKNOWN,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        func() {},
			subject:     x,
			verdict:     VERDICT_PASS,
			expectPanic: false,
		}, {
			skip:        func() {},
			fail:        func() {},
			subject:     x,
			verdict:     VERDICT_FAIL,
			expectPanic: false,
		}, {
			skip:        func() {},
			fail:        func() {},
			subject:     x,
			verdict:     VERDICT_SKIP,
			expectPanic: false,
		}, {
			skip:        func() {},
			fail:        func() {},
			subject:     nil,
			verdict:     priv_VERDICT_UNKNOWN,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        func() {},
			subject:     nil,
			verdict:     VERDICT_PASS,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        func() {},
			subject:     nil,
			verdict:     VERDICT_FAIL,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        func() {},
			subject:     nil,
			verdict:     VERDICT_SKIP,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        nil,
			subject:     x,
			verdict:     priv_VERDICT_UNKNOWN,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        nil,
			subject:     x,
			verdict:     VERDICT_PASS,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        nil,
			subject:     x,
			verdict:     VERDICT_FAIL,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        nil,
			subject:     x,
			verdict:     VERDICT_SKIP,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        func() {},
			subject:     x,
			verdict:     priv_VERDICT_UNKNOWN,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        func() {},
			subject:     x,
			verdict:     VERDICT_PASS,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        func() {},
			subject:     x,
			verdict:     VERDICT_FAIL,
			expectPanic: true,
		}, {
			skip:        nil,
			fail:        func() {},
			subject:     x,
			verdict:     VERDICT_SKIP,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        nil,
			subject:     x,
			verdict:     priv_VERDICT_UNKNOWN,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        nil,
			subject:     x,
			verdict:     VERDICT_PASS,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        nil,
			subject:     x,
			verdict:     VERDICT_FAIL,
			expectPanic: true,
		}, {
			skip:        func() {},
			fail:        nil,
			subject:     x,
			verdict:     VERDICT_SKIP,
			expectPanic: true,
		},
	}

	for _, s := range scenarios {
		// prepare
		if s.subject != nil {
			s.subject.verdict = priv_VERDICT_UNKNOWN
			s.subject.skip = s.skip
			s.subject.fail = s.fail
		}

		// execute
		_out := Exec(func() any {
			Conclude(s.subject, s.verdict)
			return ""
		})
		out, _ := _out.(string)

		// verdict
		if s.expectPanic && out == "" || !s.expectPanic && out != "" {
			t.Errorf("FAILED 1: expected panic '%v' got '%v'", s.expectPanic, out)
		}
		t.Logf("PASSED 1: expected panic '%v' got '%v'", s.expectPanic, out)

		if !s.expectPanic && s.subject != nil {
			t.Logf("LOG 2: verdict is '%d' got '%v'", s.verdict, s.subject.verdict)

			switch s.verdict {
			case priv_VERDICT_UNKNOWN:
				if s.subject.verdict != priv_VERDICT_UNKNOWN {
					t.Errorf("FAILED 2")
				}
			case VERDICT_PASS:
				if s.subject.verdict != VERDICT_PASS {
					t.Errorf("FAILED 2")
				}
			case VERDICT_FAIL:
				if s.subject.verdict != VERDICT_FAIL {
					t.Errorf("FAILED 2")
				}
			case VERDICT_SKIP:
				if s.subject.verdict != VERDICT_SKIP {
					t.Errorf("FAILED 2")
				}
			}
		}
	}
}
