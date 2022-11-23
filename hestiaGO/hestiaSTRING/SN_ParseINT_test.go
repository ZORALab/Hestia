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
	"hestia/hestiaINTERNAL/hestiaMATH"
	"hestia/hestiaTESTING"
	"testing"
)

func Test_SN_ParseINT(t *testing.T) {
	scenarios := test_cases_SN_ParseNUMBER()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/SN_ParseINT API"
		s.Description = `
test hestiaSTRING/SN_ParseINT is able to process the given string under the
following conditions.
`

		// prepare
		subject := generate_INT_string(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: '%s'", subject))

		base := create_base(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Base	: %d", base))

		sizes := create_sizes(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Sizes	: %d", sizes))

		// test
		var output int64
		panick := ""
		err := hestiaERROR.OK

		_panick := hestiaTESTING.Exec(func() any {
			output, err = SN_ParseINT(subject, base, sizes)
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

		if !assert_SN_ParseINT_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, hestiaTESTING.Format("Failed by panick!"))
			t.Fail()
		}

		if !assert_SN_ParseINT_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, hestiaTESTING.Format("Failed by output!"))
			t.Fail()
		}

		if !assert_SN_ParseINT_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, hestiaTESTING.Format("Failed by error!"))
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_SN_ParseINT_panick(panick string) bool {
	return panick == ""
}

func assert_SN_ParseINT_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		return err != hestiaERROR.OK
	case hestiaTESTING.HasCondition(s, cond_RESIZE_65):
		return err != hestiaERROR.OK
	case hestiaTESTING.HasCondition(s, cond_VALUE_NUMBER):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BROKEN):
			return err != hestiaERROR.OK
		default:
			return err == hestiaERROR.OK
		}
	default:
		return err == hestiaERROR.OK
	}
}

func assert_SN_ParseINT_output(s *hestiaTESTING.Scenario, out int64) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		return out == 0 // error raised
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		return out == 0 // error raised
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		return out == 0 // error raised
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return out == 0 // error raised
	case hestiaTESTING.HasCondition(s, cond_VALUE_NUMBER):
		switch {
		case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
			switch {
			case hestiaTESTING.HasCondition(s, cond_BROKEN):
				return out == 0 // error raised
			case hestiaTESTING.HasCondition(s, cond_RESIZE_0):
				return out == 0
			case hestiaTESTING.HasCondition(s, cond_RESIZE_2):
				return out == 0
			case hestiaTESTING.HasCondition(s, cond_RESIZE_5):
				return out == 31
			case hestiaTESTING.HasCondition(s, cond_RESIZE_8):
				return out == 0
			case hestiaTESTING.HasCondition(s, cond_RESIZE_10):
				return out == 0
			case hestiaTESTING.HasCondition(s, cond_RESIZE_12):
				return out == 4095
			case hestiaTESTING.HasCondition(s, cond_RESIZE_16):
				return out == 0
			case hestiaTESTING.HasCondition(s, cond_RESIZE_22):
				return out == 4194303
			case hestiaTESTING.HasCondition(s, cond_RESIZE_36):
				return out == 68719476735
			case hestiaTESTING.HasCondition(s, cond_RESIZE_65):
				return out == 0 // error raised
			default:
				return out == hestiaMATH.MIN_INT64
			}
		default:
			switch {
			case hestiaTESTING.HasCondition(s, cond_BROKEN):
				return out == 0 // error raised
			case hestiaTESTING.HasCondition(s, cond_RESIZE_0):
				return out == 0
			case hestiaTESTING.HasCondition(s, cond_RESIZE_2):
				return out == 3
			case hestiaTESTING.HasCondition(s, cond_RESIZE_5):
				return out == 31
			case hestiaTESTING.HasCondition(s, cond_RESIZE_8):
				return out == 127
			case hestiaTESTING.HasCondition(s, cond_RESIZE_10):
				return out == 1023
			case hestiaTESTING.HasCondition(s, cond_RESIZE_12):
				return out == 4095
			case hestiaTESTING.HasCondition(s, cond_RESIZE_16):
				return out == 32767
			case hestiaTESTING.HasCondition(s, cond_RESIZE_22):
				return out == 4194303
			case hestiaTESTING.HasCondition(s, cond_RESIZE_36):
				return out == 68719476735
			case hestiaTESTING.HasCondition(s, cond_RESIZE_65):
				return out == 0 // error raised
			default:
				return out == hestiaMATH.MAX_INT64
			}
		}
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_2):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -10
		}

		return out == 10
	case hestiaTESTING.HasCondition(s, cond_BASE_5):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -194
		}

		return out == 194
	case hestiaTESTING.HasCondition(s, cond_BASE_8):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -668
		}

		return out == 668
	case hestiaTESTING.HasCondition(s, cond_BASE_10):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -1234
		}

		return out == 1234
	case hestiaTESTING.HasCondition(s, cond_BASE_12):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -2056
		}

		return out == 2056
	case hestiaTESTING.HasCondition(s, cond_BASE_16):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -43981
		}

		return out == 43981
	case hestiaTESTING.HasCondition(s, cond_BASE_22):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -11686
		}

		return out == 11686
	case hestiaTESTING.HasCondition(s, cond_BASE_36):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return out == -49360
		}

		return out == 49360
	default:
		return out == 0 // unknown case - fail it
	}
}

//nolint:goconst
func generate_INT_string(s *hestiaTESTING.Scenario) (subject string) {
	if hestiaTESTING.HasCondition(s, cond_VALUE_NUMBER) {
		return configure_int64_minmax_string(s)
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
		subject = "-"
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			subject += "1010"
		case hestiaTESTING.HasCondition(s, cond_BASE_16):
			subject += "ABCD"
		default:
			subject += "1234"
		}
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
		subject += "0"
	default:
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		subject += ".11001"
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
		// do nothing - exactly what we want as decimaless
	default:
		subject += ".0"
	}

	return subject
}

func configure_int64_minmax_string(s *hestiaTESTING.Scenario) (subject string) {
	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
		subject = "-"

		switch {
		case hestiaTESTING.HasCondition(s, cond_BROKEN):
			subject += "9223372036854775809"
		default:
			subject += "9223372036854775808"
		}

		return subject
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		subject += "9223372036854775810"
	default:
		subject += "9223372036854775807"
	}

	return subject
}
