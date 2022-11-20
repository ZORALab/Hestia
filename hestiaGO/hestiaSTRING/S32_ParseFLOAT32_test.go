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

func Test_S32_ParseFLOAT32(t *testing.T) {
	scenarios := test_cases_SN_ParseNUMBER()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/S32_ParseFLOAT32 API"
		s.Description = `
test hestiaSTRING/S32_ParseFLOAT32 is able to process numeric string value with
the attributes listed in the Switches section.`

		// prepare
		subject := generate_float32_string(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: '%s'", subject))

		base := create_base(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Base	: %d", base))

		notation := create_notation(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Notation	: %d", notation))

		// test
		var output float32
		err := hestiaERROR.OK
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			output, err = S32_ParseFLOAT32(subject, float32(base), notation)
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: %g", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Error	: %d", err))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		if !assert_S32_ParseFLOAT32_error(s, err) {
			hestiaTESTING.Log(s, "failed due to err mismatched!")
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S32_ParseFLOAT32_panick(panick) {
			hestiaTESTING.Log(s, "failed due to panick exists!")
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S32_ParseFLOAT32_output(s, output) {
			hestiaTESTING.Log(s, "failed due to output mismatched!")
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S32_ParseFLOAT32_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_10):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			return err == hestiaERROR.OK
		default:
			return err != hestiaERROR.OK
		}
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_10):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			return err == hestiaERROR.OK
		default:
			return err != hestiaERROR.OK
		}
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_10):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			return err == hestiaERROR.OK
		default:
			return err != hestiaERROR.OK
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return err != hestiaERROR.OK
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		switch {
		case !hestiaTESTING.HasCondition(s, cond_BASE_2):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BROKEN):
			return err != hestiaERROR.OK
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			return err == hestiaERROR.INVALID_ARGUMENT
		}
	default:
	}

	return err == hestiaERROR.OK
}

func assert_S32_ParseFLOAT32_output(s *hestiaTESTING.Scenario, output float32) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_10):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			return hestiaMATH.S32_IEEE754_IsNaN_FLOAT32(output)
		default:
			return output == 0.0 // error raised
		}
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_10):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			return hestiaMATH.S32_IEEE754_IsINF_FLOAT32(output, 1)
		default:
			return output == 0.0 // error raised
		}
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_10):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			fallthrough
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			return hestiaMATH.S32_IEEE754_IsINF_FLOAT32(output, -1)
		default:
			return output == 0.0 // error raised
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return output == 0.0 // error raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		return assert_S32_ParseFLOAT32_output_ieee754(s, output)
	case hestiaTESTING.HasCondition(s, cond_VALUE_NUMBER):
		return assert_S32_ParseFLOAT32_output_minmax(s, output)
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		return assert_S32_ParseFLOAT32_output_sci_auto(s, output)
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		return assert_S32_ParseFLOAT32_output_scientific(s, output)
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		return assert_S32_ParseFLOAT32_output_iso6093nr3(s, output)
	}

	return false
}

func assert_S32_ParseFLOAT32_output_minmax(s *hestiaTESTING.Scenario, output float32) bool {
	switch {
	case !hestiaTESTING.HasCondition(s, cond_BASE_10):
		return output == 0.0 // fail this case as we don't test other bases
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			return output == 0 // overly small so assume to 0
		case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
			return hestiaMATH.S32_IEEE754_IsINF_FLOAT32(output, -1)
		default:
			return hestiaMATH.S32_IEEE754_IsINF_FLOAT32(output, 1)
		}
	default:
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			switch {
			case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
				return output == -1.4e-44 // hestiaMATH.SMALLEST_FLOAT32
			default:
				return output == 1.4e-44 // hestiaMATH.SMALLEST_FLOAT32
			}
		default:
			switch {
			case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
				return output == hestiaMATH.MIN_FLOAT32
			default:
				return output == 3.4028235e+38 // round hestiaMATH.MAX_FLOAT32
			}
		}
	}
}

func assert_S32_ParseFLOAT32_output_ieee754(s *hestiaTESTING.Scenario, output float32) bool {
	if !hestiaTESTING.HasCondition(s, cond_BASE_2) {
		return output == 0.0 // error raised
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		return output == 0.0 // error raised
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == -1.2345679e+35
		}

		return output == 1.2345679e+35
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == -1.2345678e+18
		}

		return output == 1.2345678e+18
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == -1.2345679e+07
		}

		return output == 1.2345679e+07
	}

	return false
}

//nolint:gocognit,goconst
func assert_S32_ParseFLOAT32_output_iso6093nr3(s *hestiaTESTING.Scenario, output float32) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.090625e+07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.090625e+07
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2345682e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2345682e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981855e+28
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981855e+28
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -1.0906251e-05
					}

					return output == 1.0906251e-05
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1090.625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1090.625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.234568e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.234568e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.398186e+12
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.398186e+12
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -0.10906251
					}

					return output == 0.10906251
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -10.90625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 10.90625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -123456.79
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 123456.79
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.398185911464691e+06
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.398185911464691e+06
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -10.90625
					}

					return output == 10.90625
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			}
		}
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -906250
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 906250
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -5.6780003e+19
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 5.6780003e+19
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -8.5911466e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 8.5911466e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -9.062501e-07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.062501e-07
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			// untested - redundant
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			// untested - redundant
		}
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1e+07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1e+07
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2339999e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2339999e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981e+28
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981e+28
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -9.999999e-06
					}

					return output == 9.999999e-06
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					return output == 0 // error raised
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			// untested - redundant
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			// untested - redundant
		}
	}

	return false
}

//nolint:gocognit,goconst
func assert_S32_ParseFLOAT32_output_sci_auto(s *hestiaTESTING.Scenario, output float32) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -10.90625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 10.90625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -194.93792
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 194.93792
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -668.7342529296875
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 668.7342529296875
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1234.5679
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1234.5679
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2056.4631
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2056.4631
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -43981.85911464691
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 43981.85911464691
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -11686.241
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 11686.241
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -49360.143673315804
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 49360.143673315804
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -10.90625
					}

					return output == 10.90625
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -10.90625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 10.90625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -194.93792
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 194.93792
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -668.7342529296875
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 668.7342529296875
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1234.5679
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1234.5679
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2056.4631
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2056.4631
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -43981.85911464691
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 43981.85911464691
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -11686.241
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 11686.241
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -49360.143673315804
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 49360.143673315804
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -10.90625
					}

					return output == 10.90625
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -10.90625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 10.90625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -194.93792
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 194.93792
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -668.7342529296875
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 668.7342529296875
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1234.5679
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1234.5679
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2056.4631
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2056.4631
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -43981.85911464691
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 43981.85911464691
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -11686.241
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 11686.241
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -49360.143673315804
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 49360.143673315804
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -10.90625
					}

					return output == 10.90625
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		}
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.90625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.90625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.93792003
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.93792003
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.7342529296875
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.7342529296875
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.56780005
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.56780005
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.462770061728395
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.462770061728395
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.8591146469116211
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.8591146469116211
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.24036097261116046
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.24036097261116046
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -0.14367333
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.14367333
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -0.90625
					}

					return output == 0.90625
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			// untested - redundant
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			// untested - redundant
		}
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -10
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 10
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -194
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 194
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -668
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 668
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1234
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1234
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2056
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2056
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -43981
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 43981
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -11686
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 11686
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -49360
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 49360
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -10
					}

					return output == 10
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			// untested - redundant
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			// untested - redundant
		}
	}

	return false
}

//nolint:gocognit
func assert_S32_ParseFLOAT32_output_scientific(s *hestiaTESTING.Scenario, output float32) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.090625e+07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.090625e+07
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.9493796e+14
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.9493796e+14
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -6.6873444e+20
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 6.6873444e+20
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2345682e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2345682e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.0564632e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.0564632e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981855e+28
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981855e+28
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.16862416e+30
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.16862416e+30
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.9360146e+34
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.9360146e+34
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -1.0906251e-05
					}

					return output == 1.0906251e-05
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1090.625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1090.625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.9493794e+06
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.9493794e+06
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -6.6873434e+08
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 6.6873434e+08
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.234568e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.234568e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.0564632e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.0564632e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.398186e+12
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.398186e+12
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.1686241e+14
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.1686241e+14
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.9360144e+16
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.9360144e+16
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -0.10906251
					}

					return output == 0.10906251
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -10.90625
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 10.90625
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -194.93793
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 194.93793
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -668.73425
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 668.73425
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -123456.79
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 123456.79
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -205646.31
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 205646.31
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.398185911464691e+06
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.398185911464691e+06
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.1686241e+06
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.1686241e+06
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.9360145e+06
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.9360145e+06
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -10.90625
					}

					return output == 10.90625
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		}
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -906250
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 906250
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -9.3792004e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.3792004e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -7.3425304e+17
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 7.3425304e+17
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -5.6780003e+19
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 5.6780003e+19
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.627701e+19
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.627701e+19
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -8.5911466e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 8.5911466e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.4036096e+25
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.4036096e+25
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.4367333e+29
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.4367333e+29
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -9.062501e-07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.062501e-07
					}

				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					// untested - redundant
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					// untested - redundant
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			// untested - redundant
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			// untested - redundant
		}
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
		hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			switch {
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1e+07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1e+07
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.94e+14
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.94e+14
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -6.68e+20
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 6.68e+20
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2339999e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2339999e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.0560002e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.0560002e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981e+28
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981e+28
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.1685998e+30
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.1685998e+30
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.9359987e+34
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.9359987e+34
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -9.999999e-06
					}

					return output == 9.999999e-06
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
				}
			}
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			// untested - redundant
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			// untested - redundant
		}
	}

	return false
}

func assert_S32_ParseFLOAT32_panick(panick string) bool {
	return panick == ""
}

//nolint:goconst
func generate_float32_string(s *hestiaTESTING.Scenario) (subject string) {
	if hestiaTESTING.HasCondition(s, cond_VALUE_NUMBER) {
		return configure_float32_minmax_string(s)
	}

	if hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) {
		return configure_float32_ieee754_string(s)
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		return string_FLOAT_NAN
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		return "+inf"
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return string_FLOAT_NEGATIVE_INF
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			subject = "1010"
		case hestiaTESTING.HasCondition(s, cond_BASE_16):
			subject = "ABCD"
		default:
			subject = "1234"
		}
	default:
		subject = "0"
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			subject += ".11101"
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			subject += ".43211"
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			subject += ".56774"
		case hestiaTESTING.HasCondition(s, cond_BASE_16):
			subject += ".DBEEF"
		default:
			subject += ".5678"
		}
	default:
		subject += ".0"
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		return string_FLOAT_NAN
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		return "+inf"
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return string_FLOAT_NEGATIVE_INF
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		// don't provide any exponent
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2):
			subject += "*2"
		case hestiaTESTING.HasCondition(s, cond_BASE_5):
			subject += "*5"
		case hestiaTESTING.HasCondition(s, cond_BASE_8):
			subject += "*8"
		case hestiaTESTING.HasCondition(s, cond_BASE_10):
			subject += "*10"
		case hestiaTESTING.HasCondition(s, cond_BASE_12):
			subject += "*12"
		case hestiaTESTING.HasCondition(s, cond_BASE_16):
			subject += "*16"
		case hestiaTESTING.HasCondition(s, cond_BASE_22):
			subject += "*22"
		case hestiaTESTING.HasCondition(s, cond_BASE_36):
			subject += "*36"
		}

		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
			subject += "+"
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
			subject += "-"
		}

		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			subject += "20"
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			subject += "8"
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			subject += "2"
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		subject += "%2A^!31e"
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		switch {
		case hestiaTESTING.HasCondition(s, cond_BASE_2),
			hestiaTESTING.HasCondition(s, cond_BASE_16):
			subject += "p"
		default:
			subject += "e"
		}

		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
			subject += "+"
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
			subject += "-"
		}

		switch {
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
			subject += "20"
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
			subject += "8"
		case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
			subject += "2"
		}
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
		subject = "-" + subject
	}

	return subject
}

func configure_float32_minmax_string(s *hestiaTESTING.Scenario) (subject string) {
	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
		subject = "-"
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
		subject += "1.401298464324817070923729583289916131280"
	default:
		subject += "3.40282346638528859811704183484516925440"
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			subject += "e-46"
		default:
			subject += "e+40"
		}
	default:
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			subject += "e-45"
		default:
			subject += "e+38"
		}
	}

	return subject
}

func configure_float32_ieee754_string(s *hestiaTESTING.Scenario) (subject string) {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		// do nothing - make it missing 1 bit
	case hestiaTESTING.HasCondition(s, cond_POSITIVE):
		subject = "0"
	case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
		subject = "1"
	default:
		return ""
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
		subject += "1111001101111100011011100011110"
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
		subject += "1011101100010010001000010000111"
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
		subject += "1001011001111000110000101001111"
	}

	return subject
}
