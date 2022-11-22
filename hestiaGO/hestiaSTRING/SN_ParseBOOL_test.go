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

package hestiaSTRING

import (
	"hestia/hestiaERROR"
	"hestia/hestiaTESTING"
	"testing"
)

func test_cases_SN_BOOL() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Switches: []string{
				cond_ROUND_NORMAL,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_NORMAL,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_NORMAL,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_NORMAL,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_UNKNOWNCASE,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_UNKNOWNCASE,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_UNKNOWNCASE,
				cond_NEGATIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_UNKNOWNCASE,
				cond_NEGATIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_POSITIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_NEGATIVE,
			},
		}, {
			Switches: []string{
				cond_ROUND_ZERO,
				cond_NEGATIVE,
			},
		},
	}
}

func Test_SN_BOOL(t *testing.T) {
	scenarios := test_cases_SN_BOOL()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/SN_ParseBOOL API"
		s.Description = `
test hestiaSTRING/SN_ParseBOOL is able to process the given string under the
following conditions.
`

		// prepare
		subject := generate_bool_string(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: '%s'", subject))

		// test
		var output bool
		var err hestiaERROR.Error
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			output, err = SN_ParseBOOL(subject)
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: %v", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Error	: %v", err))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		if !assert_SN_ParseBOOL_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_SN_ParseBOOL_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_SN_ParseBOOL_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_SN_ParseBOOL_panick(panick string) bool {
	return panick == ""
}

func assert_SN_ParseBOOL_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
			return err == hestiaERROR.OK
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
			switch {
			case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
				return err == hestiaERROR.OK
			case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
				return err == hestiaERROR.OK
			case hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
				return err != hestiaERROR.OK
			default:
				return err == hestiaERROR.OK
			}
		}
	default:
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
			return err == hestiaERROR.OK
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
			switch {
			case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
				return err == hestiaERROR.OK
			case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
				return err == hestiaERROR.OK
			case hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
				return err != hestiaERROR.OK
			default:
				return err == hestiaERROR.OK
			}
		}
	}

	return err != hestiaERROR.OK
}

func assert_SN_ParseBOOL_output(s *hestiaTESTING.Scenario, out bool) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
			return out == false
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
			switch {
			case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
				return out == false
			case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
				return out == false
			case hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
				return out == false // error raised
			default:
				return out == false
			}
		}
	default:
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
			return out == true
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
			switch {
			case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
				return out == true
			case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
				return out == true
			case hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
				return out == false // error raised
			default:
				return out == true
			}
		}
	}

	return out == false
}

func generate_bool_string(s *hestiaTESTING.Scenario) string {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
			return "0"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
			switch {
			case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
				return "FALSE"
			case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
				return "false"
			case hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
				return "WhAtEvEr-"
			default:
				return "False"
			}
		}
	default:
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
			return "1"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
			switch {
			case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
				return "TRUE"
			case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
				return "true"
			case hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
				return "WhAtEvEr+"
			default:
				return "True"
			}
		}
	}

	return ""
}
