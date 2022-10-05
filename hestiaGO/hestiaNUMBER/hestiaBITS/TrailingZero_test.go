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

package hestiaBITS

import (
	"testing"

	"hestia/hestiaTESTING"
)

func _TestTrailingZeroAPIScenarios() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
Test TrailingZero() is able to identify 5-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
			},
		}, {
			Description: `
Test TrailingZero() is able to identify 8-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_8_BITS: true,
			},
		}, {
			Description: `
Test TrailingZero() is able to identify 12-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_12_BITS: true,
			},
		}, {
			Description: `
Test TrailingZero() is able to identify 16-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_16_BITS: true,
			},
		}, {
			Description: `
Test TrailingZero() is able to identify 22-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_22_BITS: true,
			},
		}, {
			Description: `
Test TrailingZero() is able to identify 32-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_32_BITS: true,
			},
		}, {
			Description: `
Test TrailingZero() is able to identify 52-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_52_BITS: true,
			},
		}, {
			Description: `
Test TrailingZero() is able to identify 64-bits value properly.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
			},
		},
	}
}

func TestTrailingZeroAPI(t *testing.T) {
	scenarios := _TestTrailingZeroAPIScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = suite_TRAILING_ZERO_API

		// prepare
		if !testlibs_assertTestableHost(s) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_SKIP)
			hestiaTESTING.Logf(s, "skip test case since host is incompatible.")
			t.Skip(hestiaTESTING.ToString(s))
		}

		subject := testlibs_GenerateTrailingZero(s)

		// execute
		var output uint
		_panick := hestiaTESTING.Exec(func() any {
			output = TrailingZero(subject)
			return ""
		})
		panick, _ := _panick.(string)

		hestiaTESTING.Logf(s, "Subject	: 0B%b", subject)
		hestiaTESTING.Logf(s, "Output	: %d zero(es)", output)
		hestiaTESTING.Logf(s, "Panick	: %q", panick)

		// assert
		if panick != "" {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !testlibs_AssertTrailingZero(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		// report
		t.Logf(hestiaTESTING.ToString(s))
	}
}
