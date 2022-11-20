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

func Test_S64_ParseFLOAT64(t *testing.T) {
	scenarios := test_cases_SN_ParseNUMBER()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/S64_ParseFLOAT64 API"
		s.Description = `
test hestiaSTRING/S64_ParseFLOAT64 is able to process numeric string value with
the attributes listed in the Switches section.`

		// prepare
		subject := generate_float64_string(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: '%s'", subject))

		base := create_base(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Base	: %d", base))

		notation := create_notation(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Notation	: %d", notation))

		// test
		var output float64
		err := hestiaERROR.OK
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			output, err = S64_ParseFLOAT64(subject, float64(base), notation)
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
		if !assert_S64_ParseFLOAT64_error(s, err) {
			hestiaTESTING.Log(s, "failed due to err mismatched!")
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S64_ParseFLOAT64_panick(panick) {
			hestiaTESTING.Log(s, "failed due to panick exists!")
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S64_ParseFLOAT64_output(s, output) {
			hestiaTESTING.Log(s, "failed due to output mismatched!")
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S64_ParseFLOAT64_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
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

func assert_S64_ParseFLOAT64_output(s *hestiaTESTING.Scenario, output float64) bool {
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
			return hestiaMATH.S64_IEEE754_IsNaN_FLOAT64(output)
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
			return hestiaMATH.S64_IEEE754_IsINF_FLOAT64(output, 1)
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
			return hestiaMATH.S64_IEEE754_IsINF_FLOAT64(output, -1)
		default:
			return output == 0.0 // error raised
		}
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return output == 0.0 // error raised
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		return assert_S64_ParseFLOAT64_output_ieee754(s, output)
	case hestiaTESTING.HasCondition(s, cond_VALUE_NUMBER):
		return assert_S64_ParseFLOAT64_output_minmax(s, output)
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		return assert_S64_ParseFLOAT64_output_sci_auto(s, output)
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		return assert_S64_ParseFLOAT64_output_scientific(s, output)
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		return assert_S64_ParseFLOAT64_output_iso6093nr3(s, output)
	}

	return false
}

func assert_S64_ParseFLOAT64_output_minmax(s *hestiaTESTING.Scenario, output float64) bool {
	switch {
	case !hestiaTESTING.HasCondition(s, cond_BASE_10):
		return output == 0.0 // fail this case as we don't test other bases
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			return output == 0 // overly small so assume to 0
		case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
			return hestiaMATH.S64_IEEE754_IsINF_FLOAT64(output, -1)
		default:
			return hestiaMATH.S64_IEEE754_IsINF_FLOAT64(output, 1)
		}
	default:
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			switch {
			case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
				// rounded -hestiaMATH.SMALLEST_FLOAT64
				return output == -5e-324
			default:
				// rounded hestiaMATH.SMALLEST_FLOAT64
				return output == 5e-324
			}
		default:
			switch {
			case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
				// rounded -hestiaMATH.MIN_FLOAT64
				return output == -1.7976931348623145e+308
			default:
				// rounded hestiaMATH.MAX_FLOAT64
				return output == 1.7976931348623145e+308
			}
		}
	}
}

func assert_S64_ParseFLOAT64_output_ieee754(s *hestiaTESTING.Scenario, output float64) bool {
	if !hestiaTESTING.HasCondition(s, cond_BASE_2) {
		return output == 0.0 // error raised
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		return output == 0.0 // error raised
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
			return output == -1.234567891234e+35
		}

		return output == 1.234567891234e+35
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
			return output == -1.234567891234e-08
		}

		return output == 1.234567891234e-08
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
			return output == -1.234567891234e+07
		}

		return output == 1.234567891234e+07
	}

	return false
}

//nolint:gocognit,goconst
func assert_S64_ParseFLOAT64_output_iso6093nr3(s *hestiaTESTING.Scenario, output float64) bool {
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
						return output == -1.090625e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.090625e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2345678e+38
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2345678e+38
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981859114646907e+46
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981859114646907e+46
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
						return output == -1.090625e-09
					}

					return output == 1.090625e-09
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
						return output == -1.2345678e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2345678e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981859114646914e+13
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981859114646914e+13
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
						return output == -0.10906249999999999
					}

					return output == 0.10906249999999999
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
						return output == -123456.78
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 123456.78
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
						return output == -9.0625e+09
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.0625e+09
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -5.678000000000001e+34
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 5.678000000000001e+34
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -8.591146469116213e+41
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 8.591146469116213e+41
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
						return output == -9.062500000000001e-11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.062500000000001e-11
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
						return output == -1e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2339999999999997e+38
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2339999999999997e+38
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					return output == 0 // error raised
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981e+46
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981e+46
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
						return output == -1.0000000000000003e-09
					}

					return output == 1.0000000000000003e-09
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
func assert_S64_ParseFLOAT64_output_sci_auto(s *hestiaTESTING.Scenario, output float64) bool {
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
						return output == -1234.5678
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1234.5678
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2056.4627700617284
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2056.4627700617284
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
						return output == -11686.240360972612
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 11686.240360972612
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
						return output == -1234.5678
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1234.5678
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2056.4627700617284
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2056.4627700617284
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
						return output == -11686.240360972612
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 11686.240360972612
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
						return output == -1234.5678
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1234.5678
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2056.4627700617284
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2056.4627700617284
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
						return output == -11686.240360972612
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 11686.240360972612
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
						return output == -0.9379200000000001
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.9379200000000001
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
						return output == -0.5678000000000001
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.5678000000000001
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
						return output == -0.14367331580551745
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 0.14367331580551745
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
func assert_S64_ParseFLOAT64_output_scientific(s *hestiaTESTING.Scenario, output float64) bool {
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
						return output == -1.090625e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.090625e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.9493792000000003e+26
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.9493792000000003e+26
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -6.687342529296873e+33
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 6.687342529296873e+33
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2345678e+38
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2345678e+38
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.0564627700617282e+40
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.0564627700617282e+40
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981859114646907e+46
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981859114646907e+46
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.1686240360972622e+50
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.1686240360972622e+50
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.9360143673315806e+58
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.9360143673315806e+58
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -1.090625e-09
					}

					return output == 1.090625e-09
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
						return output == -1.9493791999999996e+07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.9493791999999996e+07
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -6.687342529296875e+09
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 6.687342529296875e+09
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2345678e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2345678e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.056462770061728e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.056462770061728e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981859114646914e+13
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981859114646914e+13
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.1686240360972612e+14
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.1686240360972612e+14
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.936014367331581e+16
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.936014367331581e+16
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -0.10906249999999999
					}

					return output == 0.10906249999999999
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
						return output == -1949.3791999999999
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1949.3791999999999
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -6687.342529296875
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 6687.342529296875
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -123456.78
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 123456.78
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -205646.2770061728
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 205646.2770061728
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
						return output == -1.1686240360972611e+06
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.1686240360972611e+06
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.936014367331581e+07
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.936014367331581e+07
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
						return output == -9.0625e+09
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.0625e+09
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -9.379200000000002e+23
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.379200000000002e+23
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -7.342529296875e+30
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 7.342529296875e+30
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -5.678000000000001e+34
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 5.678000000000001e+34
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.627700617283948e+36
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.627700617283948e+36
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -8.591146469116213e+41
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 8.591146469116213e+41
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.4036097261116047e+45
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.4036097261116047e+45
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.4367331580551744e+53
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.4367331580551744e+53
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -9.062500000000001e-11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 9.062500000000001e-11
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
						return output == -1e+11
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1e+11
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_5):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.9399999999999998e+26
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.9399999999999998e+26
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_8):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -6.680000000000003e+33
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 6.680000000000003e+33
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_10):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.2339999999999997e+38
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.2339999999999997e+38
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_12):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -2.0560000000000001e+40
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 2.0560000000000001e+40
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_16):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.3981e+46
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.3981e+46
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_22):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -1.1685999999999999e+50
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 1.1685999999999999e+50
					}
				case hestiaTESTING.HasCondition(s, cond_BASE_36):
					switch {
					case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
						return output == -4.9359999999999976e+58
					case hestiaTESTING.HasCondition(s, cond_POSITIVE):
						return output == 4.9359999999999976e+58
					}
				}
			case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
				switch {
				case hestiaTESTING.HasCondition(s, cond_BASE_2):
					if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
						return output == -1.0000000000000003e-09
					}

					return output == 1.0000000000000003e-09
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

func assert_S64_ParseFLOAT64_panick(panick string) bool {
	return panick == ""
}

func generate_float64_string(s *hestiaTESTING.Scenario) (subject string) {
	if hestiaTESTING.HasCondition(s, cond_VALUE_NUMBER) {
		return configure_float64_minmax_string(s)
	}

	if hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754) {
		return configure_float64_ieee754_string(s)
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
			subject += "35"
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
			subject += "35"
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

func configure_float64_minmax_string(s *hestiaTESTING.Scenario) (subject string) {
	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
		subject = "-"
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
		subject += "4.9406564584124654417656879286822137236505980"
	default:
		subject += "1.79769313486231570814527423731704356798070"
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			subject += "e-325"
		default:
			subject += "e+309"
		}
	default:
		switch {
		case hestiaTESTING.HasCondition(s, cond_FLOAT_SMALLEST):
			subject += "e-324"
		default:
			subject += "e+308"
		}
	}

	return subject
}

func configure_float64_ieee754_string(s *hestiaTESTING.Scenario) (subject string) {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BROKEN):
		// do nothing - make it missing 1 bit
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_POSITIVE):
		subject = "0"
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE):
		subject = "1"
	default:
		return subject
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
		subject += "100011100110111110001101110001111000000001100101110110010001010"
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
		subject += "011111001001010100000110001101111010111100110000001000001111111"
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
		subject += "100000101100111100011000010100111011101001100011110001110101000"
	}

	return subject
}
