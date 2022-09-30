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

	"hestia/hestiaERROR"
	"hestia/hestiaTESTING"
)

func _TestMatchSizeAPIScenarios() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
Test MatchSize() is able to match 5-bits positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_5_BITS:           true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 8-bits positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_8_BITS:           true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 12-bits positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_12_BITS:          true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 16-bits positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_16_BITS:          true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 22-bits positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_22_BITS:          true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 32-bits positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_32_BITS:          true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 64-bits positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_64_BITS:          true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to raise error and retain original value when matching
1024-bits positive number while given a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_1024_BITS:        true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 5-bits negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_5_BITS:           true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 8-bits negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_8_BITS:           true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 12-bits negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_12_BITS:          true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 16-bits negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_16_BITS:          true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 22-bits negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_22_BITS:          true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 32-bits negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_32_BITS:          true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 64-bits negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_64_BITS:          true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to raise error and retain orignal value when matching
1024-bits negative number while a given 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_1024_BITS:        true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 5-bits positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_5_BITS:          true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 8-bits positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_8_BITS:          true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 12-bits positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_12_BITS:         true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 16-bits positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_16_BITS:         true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 22-bits positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_22_BITS:         true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 32-bits positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_32_BITS:         true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 64-bits positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_64_BITS:         true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to raise error and retain original value when matching
1024-bits positive number while given a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_1024_BITS:       true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 5-bits negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_5_BITS:          true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 8-bits negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_8_BITS:          true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 12-bits negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_12_BITS:         true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 16-bits negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_16_BITS:         true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 22-bits negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_22_BITS:         true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 32-bits negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_32_BITS:         true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 64-bits negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_64_BITS:         true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to raise error and retain orignal value when matching
1024-bits negative number while a given 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_1024_BITS:       true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 0-bit negative number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_0_BIT:           true,
				cond_SEEK_NEGATIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 0-bit positive number properly when given
a 5-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_5_BITS: true,
				cond_SEEK_0_BIT:           true,
				cond_SEEK_POSITIVE:        true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 0-bit negative number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_0_BIT:            true,
				cond_SEEK_NEGATIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to match 0-bit positive number properly when given
a 64-bits number.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_0_BIT:            true,
				cond_SEEK_POSITIVE:         true,
			},
		}, {
			Description: `
Test MatchSize() is able to raise error when an nil input is supplied.
`,
			Switches: map[string]bool{
				cond_GENERATE_UINT_64_BITS: true,
				cond_SEEK_0_BIT:            true,
				cond_SEEK_POSITIVE:         true,
				cond_SUPPLY_NIL_POINTER:    true,
			},
		},
	}
}

func TestMatchSizeAPI(t *testing.T) {
	scenarios := _TestMatchSizeAPIScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = suite_MATCH_SIZE_API
		hestiaTESTING.Register(s, t)

		// prepare
		if !testlibs_assertTestableHost(s) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_SKIP)
			hestiaTESTING.Logf(s, "skip test case since host is incompatible.")
			t.Skip(hestiaTESTING.ToString(s))
		}

		subject := uint64(testlibs_GenerateUINT(s))
		size := testlibs_GenerateTargetSize(s)
		sign := testlibs_GenerateTargetSign(s)
		output := subject

		// execute
		err := hestiaERROR.OK
		_panick := hestiaTESTING.Exec(func() any {
			if s.Switches[cond_SUPPLY_NIL_POINTER] {
				err = MatchSize(nil, size, !sign)
			} else {
				err = MatchSize(&output, size, !sign)
			}
			return ""
		})
		panick, _ := _panick.(string)

		hestiaTESTING.Logf(s, "Subject	: 0B%b", subject)
		hestiaTESTING.Logf(s, "Output	: 0B%b", output)
		hestiaTESTING.Logf(s, "Status	: %d", err)
		hestiaTESTING.Logf(s, "Panick	: %q", panick)

		// assert
		if panick != "" {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !testlibs_AssertMatchSize(s, output, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		// report
		t.Logf(hestiaTESTING.ToString(s))
	}
}
