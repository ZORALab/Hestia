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

func Test_M32_FormatFLOAT32(t *testing.T) {
	scenarios := test_cases_SN_FormatFLOAT64()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/M32_FormatFLOAT32 API"
		s.Description = `
test hestiaSTRING/M32_FormatFLOAT32 is able to process float32 value with the
attributes listed in the Switches section.`

		// prepare
		subject := create_float32_subject(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: %e", subject))

		base := uint32(create_base(s))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Base	: %d", base))

		lettercase := create_lettercase(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Lettercase	: %d", lettercase))

		precision := uint32(create_precision(s))
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
			output, err = M32_FormatFLOAT32(subject,
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

		if !assert_S32_FormatFLOAT32_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S32_FormatFLOAT32_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S32_FormatFLOAT32_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S32_FormatFLOAT32_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
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

func assert_S32_FormatFLOAT32_output(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0),
		hestiaTESTING.HasCondition(s, cond_BASE_1),
		hestiaTESTING.HasCondition(s, cond_BASE_37):
		return output == ""
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
		hestiaTESTING.HasCondition(s, cond_EXPONENT_ZERO):
		return _assert_S32_FormatFLOAT32_output_zero(s, output)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		return _assert_S32_FormatFLOAT32_output_nan(s, output)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE),
		hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return _assert_S32_FormatFLOAT32_output_inf(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_2):
		return _assert_S32_FormatFLOAT32_output_base2(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_5):
		return _assert_S32_FormatFLOAT32_output_base5(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_8):
		return _assert_S32_FormatFLOAT32_output_base8(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_10):
		return _assert_S32_FormatFLOAT32_output_base10(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_12):
		return _assert_S32_FormatFLOAT32_output_base12(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_16):
		return _assert_S32_FormatFLOAT32_output_base16(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_22):
		return _assert_S32_FormatFLOAT32_output_base22(s, output)
	case hestiaTESTING.HasCondition(s, cond_BASE_36):
		return _assert_S32_FormatFLOAT32_output_base36(s, output)
	}

	return false
}

//nolint:gocognit,goconst
func _assert_S32_FormatFLOAT32_output_base36(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "7.clzj00*36+32"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.azd5m0*36+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.clxd00*36+32"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "7.clzi00*36+15"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "7.clzj00*36+4"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.clzj00*36-28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "7.clzi00*36-11"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "7.clzj00"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.clzj00*36+32"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "7.CLZJ00*36+32"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.azd5m0*36+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.clxd00*36+32"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "7.clzi00*36+15"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "7.clzj00*36+4"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7.clzj00*36-28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "7.clzi00*36-11"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "7.clzj00"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7clzj0000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1azd5m0000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "7clxd0000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "7clzi00000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "7clzj.000000"
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
			verdict = "7.clzj00"
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
func _assert_S32_FormatFLOAT32_output_base22(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "2.8f9d50*22+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "f.6j00i0*22+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.8f99f0*22+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.8f9d40*22+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.8f9d50*22+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.8f9d50*22-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.8f9d40*22-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.8f9d50*22+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.8f9d50*22+33"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "2.8F9D50*22+33"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "f.6j00i0*22+27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.8f99f0*22+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.8f9d40*22+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.8f9d50*22+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.8f9d50*22-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "2.8f9d40*22-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "2.8f9d50*22+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "28f9d50000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "f6j00i0000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "28f99f0000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "28f9d400000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "28f9d5.000000"
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
			verdict = "28.f9d500"
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
func _assert_S32_FormatFLOAT32_output_base16(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "b.c614f0*16+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.b41b5a*16+28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "b.c61010*16+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "b.c614e0*16+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "b.c614f0*16+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "b.c614f0*16-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "b.c614e0*16-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "b.c614f0*16+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0xb.c614f0p+33"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "0XB.C614F0P+33"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x4.b41b5ap+28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0xb.c61010p+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0xb.c614e0p+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0xb.c614f0p+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0xb.c614f0p-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0xb.c614e0p-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0xb.c614f0p+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "b.c614f0*16+33"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "B.C614F0*16+33"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.b41b5a*16+28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "b.c61010*16+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "b.c614e0*16+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "b.c614f0*16+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "b.c614f0*16-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "b.c614e0*16-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "b.c614f0*16+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0xb.c614f0p+33"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "0XB.C614F0P+33"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0x4.b41b5ap+28"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0xb.c61010p+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0xb.c614e0p+16"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0xb.c614f0p+5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0xb.c614f0p-27"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0xb.c614e0p-10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "0xb.c614f0p+1"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "bc614f0000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4b41b5a0000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "bc61010000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "bc614e00000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "bc614f.000000"
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
			verdict = "bc.614f00"
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
func _assert_S32_FormatFLOAT32_output_base12(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "4.1745a7*12+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.2516a5*12+29"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.174541*12+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "4.1745a6*12+17"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "4.1745a7*12+6"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.1745a7*12-26"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "4.1745a6*12-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "4.1745a7*12+2"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.1745a7*12+34"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "4.1745A7*12+34"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "2.2516a5*12+29"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.174541*12+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "4.1745a6*12+17"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "4.1745a7*12+6"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.1745a7*12-26"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "4.1745a6*12-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "4.1745a7*12+2"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "41745a70000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "22516a4a0000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "41745410000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "41745a600000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "41745a7.000000"
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
			verdict = "417.45a700"
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
func _assert_S32_FormatFLOAT32_output_base10(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "7.891235*10+29"
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
			verdict = "7.891235e+29"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234560e+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e+18"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568e-25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e-8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+3"
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
			verdict = "7.891235*10+29"
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
			verdict = "7.891235e+29"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234560e+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e+18"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.234568e-25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.234568e-8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.234568e+3"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "123456790000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "789123460000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "123456010000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1234567800000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "12345679.000000"
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
			verdict = "1234.567900"
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
func _assert_S32_FormatFLOAT32_output_base8(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "5.706052*8+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.550155*8+30"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5.706040*8+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "5.706052*8+18"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "5.706052*8+7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5.706052*8-25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "5.706052*8-8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "5.706052*8+3"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5.706052*8+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4.550155*8+30"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5.706040*8+35"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "5.706052*8+18"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "5.706052*8+7"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "5.706052*8-25"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "5.706052*8-8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "5.706052*8+3"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "570605170000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "4550155320000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "570604010000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "5706051600000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "57060517.000000"
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
			verdict = "5706.051700"
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
func _assert_S32_FormatFLOAT32_output_base5(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "1.113003*5+38"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.302002*5+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.113003*5+38"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.113003*5+21"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.113003*5+10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.113003*5-22"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.113003*5-5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.113003*5+6"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.113003*5+38"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.302002*5+33"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.113003*5+38"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.113003*5+21"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.113003*5+10"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.113003*5-22"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.113003*5-5"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.113003*5+6"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "111300302040000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1302001433410000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "111300244010000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1113003020300000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "11130030204.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "0.000011"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1113003.020400"
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
func _assert_S32_FormatFLOAT32_output_base2(s *hestiaTESTING.Scenario, output string) bool {
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
			verdict = "1.011110*2+51"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001011*2+48"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110*2+51"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110*2+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110*2+23"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110*2-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110*2+8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110*2+19"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110p+51"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.011110P+51"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001011p+48"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110p+51"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110p+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110p+23"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110p-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110p+8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110p+19"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110*2+51"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001011*2+48"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110*2+51"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110*2+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110*2+23"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110*2-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110*2+8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110*2+19"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110p+51"
			if hestiaTESTING.HasCondition(s, cond_UPPERCASE) {
				verdict = "1.011110P+51"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.001011p+48"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110p+51"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110p+34"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110p+23"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1.011110p-9"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "1.011110p+8"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "1.011110p+19"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1011110001100001010011110000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1001011010000011011010110100000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "1011110001100001000000010000000000000000000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "10111100011000010100111000000000000.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "101111000110000101001111.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			verdict = "0.000000"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			verdict = "101111000.110001"
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			verdict = "10111100011000010100.111100"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "11111001101111100011011100011110"
			} else {
				verdict = "01111001101111100011011100011110"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "11110001000111110101110010111010"
			} else {
				verdict = "01110001000111110101110010111010"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "11111001101111100011011011001111"
			} else {
				verdict = "01111001101111100011011011001111"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "11011101100010010001000010000111"
			} else {
				verdict = "01011101100010010001000010000111"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "11001011001111000110000101001111"
			} else {
				verdict = "01001011001111000110000101001111"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "10010110000110001101010100000101"
			} else {
				verdict = "00010110000110001101010100000101"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "10110010010101000001100011011110"
			} else {
				verdict = "00110010010101000001100011011110"
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) &&
			hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				verdict = "11000100100110100101001000101100"
			} else {
				verdict = "01000100100110100101001000101100"
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

//nolint:goconst
func _assert_S32_FormatFLOAT32_output_zero(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == "10000000000000000000000000000000"
		}

		return output == "00000000000000000000000000000000"
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

func _assert_S32_FormatFLOAT32_output_nan(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return output == "" // error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		return output == "01111111111111111111111111111111"
	default:
		return output == string_FLOAT_NAN
	}
}

func _assert_S32_FormatFLOAT32_output_inf(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return output == "" // error is raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		if hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE) {
			return output == "11111111100000000000000000000000"
		}
		return output == "01111111100000000000000000000000"
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		return output == string_FLOAT_POSITIVE_INF
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return output == string_FLOAT_NEGATIVE_INF
	default:
		return false
	}
}

func assert_S32_FormatFLOAT32_panick(panick string) bool {
	return panick == ""
}
