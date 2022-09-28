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

func TestLoglnAPI(t *testing.T) {
	x := &Scenario{}

	scenarios := []struct {
		controller  *testing.T
		subject     *Scenario
		setNil      bool
		format      string
		arg         []byte
		expectPanic bool
		expectLog   bool
	}{
		{
			controller:  t,
			setNil:      true,
			subject:     x,
			format:      "test %s",
			arg:         []byte("logger"),
			expectPanic: false,
			expectLog:   true,
		}, {
			controller:  t,
			setNil:      false,
			subject:     x,
			format:      "test %s",
			arg:         []byte("logger"),
			expectPanic: false,
			expectLog:   true,
		}, {
			controller:  nil,
			setNil:      true,
			subject:     x,
			format:      "test %s",
			arg:         []byte("logger"),
			expectPanic: false,
			expectLog:   true,
		}, {
			controller:  nil,
			setNil:      false,
			subject:     x,
			format:      "test %s",
			arg:         []byte("logger"),
			expectPanic: false,
			expectLog:   true,
		}, {
			controller:  t,
			setNil:      true,
			subject:     x,
			format:      "",
			arg:         []byte("logger"),
			expectPanic: false,
			expectLog:   true,
		}, {
			controller:  t,
			setNil:      true,
			subject:     nil,
			format:      "test %s",
			arg:         []byte("logger"),
			expectPanic: true,
			expectLog:   false,
		}, {
			controller:  t,
			setNil:      true,
			subject:     nil,
			format:      "",
			arg:         []byte("logger"),
			expectPanic: true,
			expectLog:   false,
		}, {
			controller:  t,
			setNil:      true,
			subject:     x,
			format:      "test %s",
			arg:         nil,
			expectPanic: false,
			expectLog:   true,
		}, {
			controller:  t,
			setNil:      false,
			subject:     x,
			format:      "test %s",
			arg:         nil,
			expectPanic: false,
			expectLog:   true,
		},
	}

	for _, s := range scenarios {
		// prepare
		if s.subject != nil {
			if s.setNil {
				s.subject.Log = nil
			} else {
				s.subject.Log = []string{}
			}
			s.subject.controller = s.controller
		}

		// execute
		_out := Exec(func() any {
			Logln(s.subject, s.format, s.arg)
			return ""
		})

		out, _ := _out.(string)
		entry := 0
		if s.subject != nil {
			entry = len(s.subject.Log)
		}

		// verdict
		if s.expectPanic && out == "" || !s.expectPanic && out != "" {
			t.Errorf("FAILED 1: expected panic '%v' got '%v'", s.expectPanic, out)
		}
		t.Logf("PASSED 1: expected panic '%v' got '%v'", s.expectPanic, out)

		if s.subject != nil {
			if s.expectLog && entry == 0 || !s.expectLog && entry != 0 {
				t.Errorf("FAILED 2: expect log '%v' got '%v'",
					s.expectLog,
					entry == 0,
				)
			} else {
				t.Logf("PASSED 2: expect log '%v' got '%v'",
					s.expectLog,
					entry == 0,
				)
			}
		} else {
			if s.expectLog {
				t.Errorf("ERROR 2: missing subject but expect registration")
			} else {
				t.Logf("PASSED 2: expect log '%v' got 'log'",
					s.expectLog,
				)
			}
		}
	}
}
