// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
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

// NOTE:
//   1. Test codes are in the public hestiaSTRING package.
//   2. Ryū algorithm paper -> https://doi.org/10.1145/3192366.3192369

package hestiaFMT

type omizuData struct {
	round            []rune
	partial          []rune
	exponent         []rune
	base             []rune
	notation         Notation
	lettercase       Lettercase
	negativeValue    bool
	negativeExponent bool
}

func __sN_Omizu_Render_FLOAT(data *omizuData) (out []rune) {
	var base string

	if data == nil {
		panic("__sN_Hadōken_Render receiving nil data structure!")
	}

	out = []rune{}
	base = string(data.base)

	if data.negativeValue {
		out = append(out, '-')
	}

	switch {
	case data.notation == NOTATION_ISO6093NR3,
		data.notation == NOTATION_ISO6093NR3_AUTO:
		if base == "16" {
			if data.lettercase == LETTERCASE_UPPER {
				out = append(out, '0', 'X')
			} else {
				out = append(out, '0', 'x')
			}
		}
	}

	out = append(out, data.round...)
	out = append(out, '.')
	out = append(out, data.partial...)

	if len(data.exponent) > 0 {
		switch {
		case data.notation == NOTATION_ISO6093NR3,
			data.notation == NOTATION_ISO6093NR3_AUTO:
			switch {
			case base == "2", base == "16":
				if data.lettercase == LETTERCASE_UPPER {
					out = append(out, 'P')
				} else {
					out = append(out, 'p')
				}
			case base == "10":
				if data.lettercase == LETTERCASE_UPPER {
					out = append(out, 'E')
				} else {
					out = append(out, 'e')
				}
			}
		default:
			out = append(out, '*')
			out = append(out, data.base...)
		}

		if data.negativeExponent {
			out = append(out, '-')
		} else {
			out = append(out, '+')
		}

		out = append(out, data.exponent...)
	}

	return out
}

func __sN_Omizu_Render_INF_FLOAT(sign bool) []rune {
	if sign {
		return []rune{'+', '∞'}
	}
	return []rune{'-', '∞'}
}

func __sN_Omizu_Render_NaN_FLOAT() []rune {
	return []rune{'N', 'a', 'N'}
}

func __SN_Valid_Base_Number(base uint8) {
	if base < 2 || base > 36 {
		panic("base number smaller than 2 or larger than 36!")
	}
}

func __SN_Valid_Notation(base uint8, notation Notation) {
	switch {
	case notation == NOTATION_ISO6093NR3_AUTO && !(base == 2 || base == 10 || base == 16):
		fallthrough
	case notation == NOTATION_ISO6093NR3 && !(base == 2 || base == 10 || base == 16):
		panic("requested ISO6093 NR3 notation but base is not 2/10/16!")
	case notation == NOTATION_IEEE754 && base != 2:
		panic("requested IEEE754 notation but base is not 2!")
	default:
	}
}

func __s32_LOG2_5PowerOf(input int32) int32 {
	if input > 3528 {
		panic("Format_FLOAT: Log2_5PowerOf has big input (>3528).")
	}

	return int32((uint32(input)*1217359)>>19 + 1)
}

func __s32_LOG10_2PowerOf(input int32) uint32 {
	if input > 1650 {
		panic("Format_FLOAT: Log10_2PowerOf has big input (>1650).")
	}

	return (uint32(input) * 78913) >> 18
}

func __s32_LOG10_5PowerOf(input int32) uint32 {
	if input > 2620 {
		panic("Format_FLOAT: Log10_5PowerOf has big input (>2620).")
	}

	return (uint32(input) * 732923) >> 20
}
