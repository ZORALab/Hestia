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

func TestScenarioExecAPI(t *testing.T) {
	scenarios := []struct {
		doPanic bool
	}{
		{
			doPanic: false,
		}, {
			doPanic: true,
		},
	}

	for _, s := range scenarios {
		// prepare
		subject := &Scenario{}
		subject.Init()

		// execute
		_panick := subject.Exec(func() any {
			if s.doPanic {
				panic("got panic")
			}

			return ""
		})
		panick, _ := _panick.(string)

		// verdict
		if s.doPanic && panick == "" || !s.doPanic && panick != "" {
			t.Errorf("FAILED 1: expected panic '%v' got '%v'", s.doPanic, panick)
		}
		t.Logf("PASSED 1: expected panic '%v' got '%v'", s.doPanic, panick)
	}
}
