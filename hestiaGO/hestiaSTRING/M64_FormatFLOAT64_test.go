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

func Test_M64_FormatFLOAT64(t *testing.T) {
	scenarios := test_cases_SN_FormatFLOAT64()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/M64_FormatFLOAT64 API"
		s.Description = `
test hestiaSTRING/M64_FormatFLOAT64 is able to process float64 value with the
attributes listed in the Switches section.`

		// prepare
		subject := create_float64_subject(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: %e", subject))

		base := create_base(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Base	: %d", base))

		lettercase := create_lettercase(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Lettercase	: %d", lettercase))

		precision := create_precision(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Precision	: %d", precision))

		notation := create_notation(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Notation	: %d", notation))

		// test
		var output string
		err := hestiaERROR.OK
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			output, err = M64_FormatFLOAT64(subject,
				base,
				precision,
				lettercase,
				notation,
			)
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: '%s'", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Error	: %d", err))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		if !assert_S64_FormatFLOAT64_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S64_FormatFLOAT64_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S64_FormatFLOAT64_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S64_FormatFLOAT64_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0),
		hestiaTESTING.HasCondition(s, cond_BASE_1),
		hestiaTESTING.HasCondition(s, cond_BASE_37):
		return err == hestiaERROR.DATA_INVALID
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) &&
		hestiaTESTING.HasCondition(s, cond_BASE_5):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) &&
		hestiaTESTING.HasCondition(s, cond_BASE_8):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) &&
		hestiaTESTING.HasCondition(s, cond_BASE_10):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) &&
		hestiaTESTING.HasCondition(s, cond_BASE_12):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) &&
		hestiaTESTING.HasCondition(s, cond_BASE_16):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) &&
		hestiaTESTING.HasCondition(s, cond_BASE_22):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3) &&
		!(hestiaTESTING.HasCondition(s, cond_BASE_2) ||
			hestiaTESTING.HasCondition(s, cond_BASE_10) ||
			hestiaTESTING.HasCondition(s, cond_BASE_16)):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO) &&
		!(hestiaTESTING.HasCondition(s, cond_BASE_2) ||
			hestiaTESTING.HasCondition(s, cond_BASE_10) ||
			hestiaTESTING.HasCondition(s, cond_BASE_16)):
		fallthrough
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) &&
		hestiaTESTING.HasCondition(s, cond_BASE_36):
		return err == hestiaERROR.BAD_DESCRIPTOR
	}

	return err == hestiaERROR.OK
}

func assert_S64_FormatFLOAT64_output(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0),
		hestiaTESTING.HasCondition(s, cond_BASE_1),
		hestiaTESTING.HasCondition(s, cond_BASE_37):
		return output == ""
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
		hestiaTESTING.HasCondition(s, cond_EXPONENT_ZERO):
		return _assert_S64_FormatFLOAT64_output_zero(s, output)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		return _assert_S64_FormatFLOAT64_output_nan(s, output)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE),
		hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return _assert_S64_FormatFLOAT64_output_inf(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_2):
		return _assert_S64_FormatFLOAT64_output_base2(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_5):
		return _assert_S64_FormatFLOAT64_output_base5(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_8):
		return _assert_S64_FormatFLOAT64_output_base8(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_10):
		return _assert_S64_FormatFLOAT64_output_base10(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_12):
		return _assert_S64_FormatFLOAT64_output_base12(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_16):
		return _assert_S64_FormatFLOAT64_output_base16(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_22):
		return _assert_S64_FormatFLOAT64_output_base22(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_36):
		return _assert_S64_FormatFLOAT64_output_base36(s, output)
	}

	return false
}

//nolint:gocognit,goconst
func _assert_S64_FormatFLOAT64_output_base36(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "f.r5huhi*36+30"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.5p7iuq*36+24"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.n9c000*36+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "f.r5huhi*36+13"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "f.r5huhi*36+2"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "f.r5huhi*36-30"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "f.r5huhi*36-13"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "f.r5huhi*36-2"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "f.r5huhi*36+30"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "F.R5HUHI*36+30"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.5p7iuq*36+24"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.n9c000*36+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "f.r5huhi*36+13"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "f.r5huhi*36+2"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "f.r5huhi*36-30"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "f.r5huhi*36-13"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "f.r5huhi*36-2"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "fr5huhia00000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "25p7iupz4sh00000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2n9c000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "fr5huhia000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "fr5.huhia0"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0.0fr5hu"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		// empty as error is raised
	default:
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:gocognit,goconst
func _assert_S64_FormatFLOAT64_output_base22(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.0akglb*22+32"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "d.b29b9a*22+25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "b.d1e000*22+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.0akglb*22+15"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.0akglb*22+4"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.0akglb*22-28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.0akglb*22-11"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.0akglb"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.0akglb*22+32"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.0AKGLB*22+32"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "d.b29b9a*22+25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "b.d1e000*22+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.0akglb*22+15"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.0akglb*22+4"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.0akglb*22-28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.0akglb*22-11"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.0akglb"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "10akglb3ii00000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "db29b99kl4id00000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "bd1e000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "10akglb3ii000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "10akg.lb3ii0"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.0akglb"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		// empty as error is raised
	default:
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:gocognit,goconst
func _assert_S64_FormatFLOAT64_output_base16(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.1f71fb*16+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.c09091*16+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.e24000*16+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.1f71fb*16+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.1f71fb*16+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.1f71fb*16-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.1f71fb*16-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.1f71fb*16+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.1f71fbp+33"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "0X1.1F71FBP+33"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.c09091p+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.e24000p+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0x1.1f71fbp+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0x1.1f71fbp+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.1f71fbp-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0x1.1f71fbp-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0x1.1f71fbp+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.1f71fb*16+33"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.1F71FB*16+33"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.c09091*16+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.e24000*16+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.1f71fb*16+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.1f71fb*16+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.1f71fb*16-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.1f71fb*16-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.1f71fb*16+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.1f71fbp+33"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "0x1.1F71FBP+33"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.c09091p+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.e24000p+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0x1.1f71fbp+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0x1.1f71fbp+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x1.1f71fbp-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0x1.1f71fbp-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0x1.1f71fbp+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "11f71fb092200000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1c09091768d40100000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1e240000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "11f71fb0922000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "11f71f.b09220"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "11.f71fb1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		// empty as error is raised
	default:
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:gocognit,goconst
func _assert_S64_FormatFLOAT64_output_base12(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.7b3263*12+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "6.1907a0*12+28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5.b54000*12+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.7b3263*12+17"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.7b3263*12+6"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.7b3263*12-26"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.7b3263*12-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.7b3263*12+2"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.7b3263*12+34"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.7B3263*12+34"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "6.1907a0*12+28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5.b54000*12+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.7b3263*12+17"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.7b3263*12+6"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.7b3263*12-26"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.7b3263*12-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.7b3263*12+2"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "17b32635716a00000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "61907a02414071500000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5b540000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "17b32635716a000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "17b3263.5716a0"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "17b.326357"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		verdict = "" // error raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		verdict = "" // error raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		// empty as error is raised
	default:
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:gocognit,goconst
func _assert_S64_FormatFLOAT64_output_base10(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568*10+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.891234*10+29"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234560*10+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568*10+18"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568*10+7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568*10-25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568*10-8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568*10+3"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568e+35"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+35"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.891234e+29"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "7.891234E+29"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234560e+35"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234560E+35"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e+18"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+18"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+7"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+7"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568e-25"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E-25"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e-8"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E-8"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+3"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+3"
			}
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568*10+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.891234*10+29"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234560*10+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568*10+18"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568*10+7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568*10-25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568*10-8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568*10+3"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568e+35"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+35"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.891234e+29"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "7.891234E+29"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234560e+35"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234560E+35"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e+18"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+18"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+7"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+7"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568e-25"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E-25"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e-8"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E-8"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+3"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.234568E+3"
			}
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "123456789123400000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "789123400000000100000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "123456000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1234567891234000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "12345678.912340"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1234.567891"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		// empty as error is raised
	default:
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:gocognit,goconst
func _assert_S64_FormatFLOAT64_output_base8(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.175620*8+36"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "3.402205*8+31"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "3.611000*8+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.175620*8+19"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.175620*8+8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.175620*8-24"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.175620*8-7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.175620*8+4"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.175620*8+36"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "3.402205*8+31"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "3.611000*8+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.175620*8+19"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.175620*8+8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.175620*8-24"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.175620*8-7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.175620*8+4"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2175617660444200000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "34022044273215200100000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "361100000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "21756176604442000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "217561766.044420"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "21756.176605"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		// empty as error is raised
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:gocognit,goconst
func _assert_S64_FormatFLOAT64_output_base5(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.302114*5+40" // 1.234568e+35
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "3.123331*5+36" // 7.891234e+29
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.242231*5+37" // 1.234560e+35
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.302114*5+23" // 1.234560e+35
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.302114*5+12" // 1.234568e+18
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.302114*5-20" // 1.234568e-25
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.302114*5-3" // 1.234568e-08
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.302114*5+8" // 1.234568e+03
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.302114*5+40" // 1.234568e+35
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "3.123331*5+36" // 7.891234e+29
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.242231*5+37" // 1.234560e+35
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.302114*5+23" // 1.234568e+18
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.302114*5+12" // 1.234568e+07
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.302114*5-20" // 1.234568e-25
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.302114*5-3" // 1.234568e-08
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.302114*5+8" // 1.234568e+03
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "13021134334000441400000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "3123330443421300000000100000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "12422311000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "130211343340004414000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1302113433400.044140"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.001302"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "130211343.340010"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		// empty as error is raised
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:gocognit,goconst,lll
func _assert_S64_FormatFLOAT64_output_base2(s *hestiaTESTING.Scenario, output string) bool {
	verdict := ""

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		// empty as error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000*2+63"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.110000*2+66"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.111001*2+46"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000*2+46"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000*2+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000*2+3"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000*2+20"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000*2+31"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000p+63"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+63"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.110000p+66"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.110000P+66"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.111001p+46"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.111001P+46"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000p+46"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+46"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000p+35"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+35"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000p+3"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+3"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000p+20"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+20"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000p+31"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+31"
			}
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000*2+63"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.110000*2+66"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.111001*2+46"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000*2+46"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000*2+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000*2+3"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000*2+20"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000*2+31"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000p+63"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+63"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.110000p+66"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.110000P+66"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.111001p+46"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.111001P+46"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000p+46"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+46"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000p+35"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+35"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001000p+3"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+3"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.001000p+20"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+20"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.001000p+31"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.001000P+31"
			}
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1000111110111000111111011000010010010001000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1110000001001000010010001011101101000110101000000000100000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "11110001001000000000000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "10001111101110001111110110000100100100010000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "100011111011100011111101100001001001.000100"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1000.111111"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "100011111011100011111.101100"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "10001111101110001111110110000100.100100"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1100011100110111110001101110001111000000001100101110110010001010"
			} else {
				verdict = "0100011100110111110001101110001111000000001100101110110010001010"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1100011000100011111010111001011100100111100111111110101000101111"
			} else {
				verdict = "0100011000100011111010111001011100100111100111111110101000101111"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1100011100110111110001101101100111001010011001110101100010111010"
			} else {
				verdict = "0100011100110111110001101101100111001010011001110101100010111010"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1100001110110001001000100001000011110100110000000001101100001100"
			} else {
				verdict = "0100001110110001001000100001000011110100110000000001101100001100"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1100000101100111100011000010100111011101001100011110001110101000"
			} else {
				verdict = "0100000101100111100011000010100111011101001100011110001110101000"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1011101011000011000110101010000010010101001010100100010011011101"
			} else {
				verdict = "0011101011000011000110101010000010010101001010100100010011011101"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1011111001001010100000110001101111010111100110000001000001111111"
			} else {
				verdict = "0011111001001010100000110001101111010111100110000001000001111111"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "1100000010010011010010100100010110000101010001111001011011011100"
			} else {
				verdict = "0100000010010011010010100100010110000101010001111001011011011100"
			}
		}
	default:
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) &&
		!(hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) ||
			hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN)) {
		verdict = "-" + verdict
	}

	return verdict == output
}

//nolint:lll
func _assert_S64_FormatFLOAT64_output_zero(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == "1000000000000000000000000000000000000000000000000000000000000000"
		}

		return output == "0000000000000000000000000000000000000000000000000000000000000000"
	case hestiaTESTING.HasCondition(s, cond_PRECISION_0):
		return output == "0.0"
	case hestiaTESTING.HasCondition(s, cond_PRECISION_1):
		return output == "0.0"
	case hestiaTESTING.HasCondition(s, cond_PRECISION_4):
		return output == "0.0000"
	case hestiaTESTING.HasCondition(s, cond_PRECISION_5):
		return output == "0.00000"
	case hestiaTESTING.HasCondition(s, cond_PRECISION_10):
		return output == "0.0000000000"
	default:
		return output == "0.0"
	}
}

func _assert_S64_FormatFLOAT64_output_nan(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return output == "" // error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		return output == "0111111111111111111111111111111111111111111111111111111111111111"
	default:
		return output == string_FLOAT_NAN
	}
}

//nolint:lll
func _assert_S64_FormatFLOAT64_output_inf(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return output == "" // error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		if hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE) {
			return output == "1111111111110000000000000000000000000000000000000000000000000000"
		}
		return output == "0111111111110000000000000000000000000000000000000000000000000000"
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		return output == string_FLOAT_POSITIVE_INF
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return output == string_FLOAT_NEGATIVE_INF
	default:
		return false
	}
}

func assert_S64_FormatFLOAT64_panick(panick string) bool {
	return panick == ""
}
