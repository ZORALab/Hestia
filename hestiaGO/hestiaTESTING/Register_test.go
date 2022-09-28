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

func TestRegisterAPI(t *testing.T) {
	x := &Scenario{}

	scenarios := []struct {
		subject        *Scenario
		controller     *testing.T
		expectPanic    bool
		expectRegister bool
	}{
		{
			subject:        x,
			controller:     t,
			expectPanic:    false,
			expectRegister: true,
		}, {
			subject:        x,
			controller:     nil,
			expectPanic:    true,
			expectRegister: false,
		}, {
			subject:        nil,
			controller:     t,
			expectPanic:    true,
			expectRegister: false,
		}, {
			subject:        nil,
			controller:     nil,
			expectPanic:    true,
			expectRegister: false,
		},
	}

	for _, s := range scenarios {
		// prepare
		if s.subject != nil {
			s.subject.controller = nil
		}

		// execute
		outExec := Exec(func() any {
			Register(s.subject, s.controller)
			return ""
		})

		out, _ := outExec.(string)

		// verdict
		if s.expectPanic && out == "" || !s.expectPanic && out != "" {
			t.Errorf("FAILED 1: expected panic '%v' got '%v'", s.expectPanic, out)
		}
		t.Logf("PASSED 1: expected panic '%v' got '%v'", s.expectPanic, out)

		if s.subject != nil {
			if s.expectRegister && s.subject.controller == nil ||
				!s.expectRegister && s.subject.controller != nil {
				t.Errorf("FAILED 2: expect register '%v' got '%v'",
					s.expectRegister,
					s.subject.controller == nil,
				)
			} else {
				t.Logf("PASSED 2: expect register '%v' got '%v'",
					s.expectRegister,
					s.subject.controller == nil,
				)
			}
		} else {
			if s.expectRegister {
				t.Errorf("ERROR 2: missing subject but expect registration")
			} else {
				t.Logf("PASSED 2: expect register '%v' got 'false'",
					s.expectRegister,
				)
			}
		}
	}
}
