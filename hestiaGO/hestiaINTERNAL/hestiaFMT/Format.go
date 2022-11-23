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

// NOTE:
//   1. Test codes are in the public hestiaSTRING package.

package hestiaFMT

type engineMode uint

const (
	engine_PARSE_NORMAL engineMode = iota
	engine_PARSE_VERB
)

type engine struct {
	arguments    []any
	pos          uint16
	width        []rune
	precision    []rune
	mode         engineMode
	setPrecision bool
}

func (e *engine) arg() (out any) {
	if e.pos >= uint16(len(e.arguments)) {
		return None(true)
	}

	out = e.arguments[e.pos]
	e.pos++

	return out
}

func (e *engine) to_PARSE_NORMAL() {
	e.mode = engine_PARSE_NORMAL
	e.width = []rune{}
	e.precision = []rune{}
	e.setPrecision = false
}

type argument struct {
	arg       any
	base      uint8
	precision []rune
	notation  Notation
	verb      rune
}

func M64_Format(statement string, args ...any) *[]rune {
	var ret, output []rune
	var core *engine

	output = []rune{}
	core = &engine{
		arguments: args,
		pos:       0,
	}
	core.to_PARSE_NORMAL()

	for _, c := range statement {
		if core.mode == engine_PARSE_NORMAL {
			if c == '%' {
				core.mode = engine_PARSE_VERB
				continue
			}

			// parse as normal character otherwise and move on
			output = append(output, c)
			continue
		}

		// | beyond this point is verb mode (%c, %s, %d, ...) | //
		switch c {
		case '%': // literal percent character ('%%')
			ret = []rune{'%'}
		case 't': // bool verb '%t'
			ret = _M64_formatBool(core.arg())
		case 'c': // char verb '%t'
			ret = _M64_formatChar(core.arg())
		case 's': // string verb '%s'
			ret = _M64_formatString(core.arg())
		case 'd': // decimal verb '%d'
			ret = _M64_formatNumber(&argument{
				arg:       core.arg(),
				base:      10,
				precision: core.precision,
				notation:  NOTATION_DECIMALLESS,
				verb:      c,
			})
		case 'o', 'O': // octal verb '%o' or '%O'
			ret = _M64_formatNumber(&argument{
				arg:       core.arg(),
				base:      10,
				precision: core.precision,
				notation:  NOTATION_DECIMALLESS,
				verb:      c,
			})
		case 'x', 'X': // hexdecimal verb '%x or '%X'
			ret = _M64_formatNumber(&argument{
				arg:       core.arg(),
				base:      16,
				precision: core.precision,
				notation:  NOTATION_ISO6093NR3,
				verb:      c,
			})
		case 'b': // binary verb '%b'
			ret = _M64_formatNumber(&argument{
				arg:       core.arg(),
				base:      2,
				precision: core.precision,
				notation:  NOTATION_ISO6093NR3,
				verb:      c,
			})
		case 'f', 'F': // float with no exponent verb '%F' or '%f'
			ret = _M64_formatNumber(&argument{
				arg:       core.arg(),
				base:      10,
				precision: core.precision,
				notation:  NOTATION_DECIMAL_NO_EXPONENT,
				verb:      c,
			})
		case 'e', 'E': // float in scientific notation verb '%e' or '%E'
			ret = _M64_formatNumber(&argument{
				arg:       core.arg(),
				base:      10,
				precision: core.precision,
				notation:  NOTATION_ISO6093NR3,
				verb:      c,
			})
		case 'g', 'G': // float in auto-scientific notation verb '%g' or '%G'
			ret = _M64_formatNumber(&argument{
				arg:       core.arg(),
				base:      10,
				precision: core.precision,
				notation:  NOTATION_ISO6093NR3_AUTO,
				verb:      c,
			})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// assign to width/precision '%{w.p}f' like '10' or '2' in '%10.2f'
			if core.setPrecision {
				core.precision = append(core.precision, c)
			} else {
				core.width = append(core.width, c)
			}

			continue
		case '.':
			// change to precision 'p' in '%{w.p}f' like '2' in '%10.2f'
			if !core.setPrecision {
				core.setPrecision = true
				continue
			}

			// bad verb - multiple period ('%{w..p...}f')
			ret = []rune{'(', 'E', 'R', 'R', 'O', 'R', '=', '%', '.', '.', ')'}
			_ = core.arg()
		default:
			// unknown verb - '%?'
			ret = []rune{'(', 'E', 'R', 'R', 'O', 'R', '=', '\'', '%'}
			switch c {
			case '\a':
				ret = append(ret, '\\', 'a')
			case '\b':
				ret = append(ret, '\\', 'b')
			case '\t':
				ret = append(ret, '\\', 't')
			case '\n':
				ret = append(ret, '\\', 'n')
			case '\v':
				ret = append(ret, '\\', 'v')
			case '\f':
				ret = append(ret, '\\', 'f')
			case '\r':
				ret = append(ret, '\\', 'r')
			default:
				ret = append(ret, c)
			}
			ret = append(ret, '\'', '?', '?', '?', ')')
			_ = core.arg()
		}

		_S64_width_correction(&ret, &core.width)
		output = append(output, ret...)
		core.to_PARSE_NORMAL()
	}

	return &output
}

//nolint:gocognit
func _M64_formatNumber(str *argument) (out []rune) {
	var lettercase Lettercase
	var precision uint64

	// process configurations
	lettercase = LETTERCASE_LOWER
	switch str.verb {
	case 'E', 'F', 'G', 'X':
		lettercase = LETTERCASE_UPPER
	}

	precision, _ = SN_ParseUINT(string(str.precision), 10, 32)

	// process argument by type
	out = []rune{'(', 'N', 'U', 'M', 'B', 'E', 'R', '=', 'b', 'a', 'd', ')'}
	switch v := str.arg.(type) {
	case uint8:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S8_FormatUINT8(v, str.base, lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S8_FormatUINT8(v, str.base, lettercase)
		}
	case *uint8:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S8_FormatUINT8(*v, str.base, lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S8_FormatUINT8(*v, str.base, lettercase)
		}
	case uint16:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S16_FormatUINT16(v, uint16(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S16_FormatUINT16(v, uint16(str.base), lettercase)
		}
	case *uint16:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S16_FormatUINT16(*v, uint16(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S16_FormatUINT16(*v, uint16(str.base), lettercase)
		}
	case uint32:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S32_FormatUINT32(v, uint32(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S32_FormatUINT32(v, uint32(str.base), lettercase)
		}
	case *uint32:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S32_FormatUINT32(*v, uint32(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S32_FormatUINT32(*v, uint32(str.base), lettercase)
		}
	case uint64:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S64_FormatUINT64(v, uint64(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S64_FormatUINT64(v, uint64(str.base), lettercase)
		}
	case *uint64:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S64_FormatUINT64(*v, uint64(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S64_FormatUINT64(*v, uint64(str.base), lettercase)
		}
	case int8:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S8_FormatINT8(v, str.base, lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S8_FormatINT8(v, str.base, lettercase)
		}
	case *int8:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S8_FormatINT8(*v, str.base, lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S8_FormatINT8(*v, str.base, lettercase)
		}
	case int16:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S16_FormatINT16(v, uint16(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S16_FormatINT16(v, uint16(str.base), lettercase)
		}
	case *int16:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S16_FormatINT16(*v, uint16(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S16_FormatINT16(*v, uint16(str.base), lettercase)
		}
	case int32:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S32_FormatINT32(v, uint32(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S32_FormatINT32(v, uint32(str.base), lettercase)
		}
	case *int32:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S32_FormatINT32(*v, uint32(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S32_FormatINT32(*v, uint32(str.base), lettercase)
		}
	case int64:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S64_FormatINT64(v, uint64(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S64_FormatINT64(v, uint64(str.base), lettercase)
		}
	case *int64:
		switch str.verb {
		case 'O':
			out = append([]rune{'0', 'o'},
				S64_FormatINT64(*v, uint64(str.base), lettercase)...)
		case 'b', 'd', 'o', 'x', 'X':
			out = S64_FormatINT64(*v, uint64(str.base), lettercase)
		}
	case float32:
		switch str.verb {
		case 'b', 'e', 'E', 'f', 'F', 'g', 'G', 'x', 'X':
			param := &ParamsFLOAT32{
				Value:      v,
				Precision:  uint32(precision),
				Base:       uint32(str.base),
				Lettercase: lettercase,
				Notation:   str.notation,
			}

			out = M32_FormatFLOAT32(param)
		}
	case *float32:
		switch str.verb {
		case 'b', 'e', 'E', 'f', 'F', 'g', 'G', 'x', 'X':
			param := &ParamsFLOAT32{
				Value:      *v,
				Precision:  uint32(precision),
				Base:       uint32(str.base),
				Lettercase: lettercase,
				Notation:   str.notation,
			}

			out = M32_FormatFLOAT32(param)
		}
	case float64:
		switch str.verb {
		case 'b', 'e', 'E', 'f', 'F', 'g', 'G', 'x', 'X':
			param := &ParamsFLOAT64{
				Value:      v,
				Precision:  precision,
				Base:       uint64(str.base),
				Lettercase: lettercase,
				Notation:   str.notation,
			}

			out = M64_FormatFLOAT64(param)
		}
	case *float64:
		switch str.verb {
		case 'b', 'e', 'E', 'f', 'F', 'g', 'G', 'x', 'X':
			param := &ParamsFLOAT64{
				Value:      *v,
				Precision:  precision,
				Base:       uint64(str.base),
				Lettercase: lettercase,
				Notation:   str.notation,
			}

			out = M64_FormatFLOAT64(param)
		}
	}

	return out
}

func _M64_formatString(arg any) (out []rune) {
	switch v := arg.(type) {
	case *string:
		return []rune(*v)
	case string:
		return []rune(v)
	case *rune:
		return []rune{*v}
	case rune:
		return []rune{v}
	case *[]rune:
		return *v
	case []rune:
		return v
	default:
		return []rune{'(', 'S', 'T', 'R', 'I', 'N', 'G', '=', 'b', 'a', 'd', ')'}
	}
}

func _M64_formatChar(arg any) (out []rune) {
	switch v := arg.(type) {
	case *string:
		switch len(*v) {
		case 0:
			return []rune{}
		case 1:
			return []rune(*v)
		default:
			return []rune((*v)[:1])
		}
	case string:
		switch len(v) {
		case 0:
			return []rune{}
		case 1:
			return []rune(v)
		default:
			return []rune((v)[:1])
		}
	case *rune:
		// int32 is treated as alias of rune
		if SN_IsPrintable_RUNE(*v) {
			return []rune{*v}
		}

		return []rune{'(', 'C', 'H', 'A', 'R', '=', 'b', 'a', 'd', ')'}
	case rune:
		// int32 is treated as alias of rune
		if SN_IsPrintable_RUNE(v) {
			return []rune{v}
		}

		return []rune{'(', 'C', 'H', 'A', 'R', '=', 'b', 'a', 'd', ')'}
	case *[]rune:
		return (*v)[:1]
	case []rune:
		return v[:1]
	default:
		return []rune{'(', 'C', 'H', 'A', 'R', '=', 'b', 'a', 'd', ')'}
	}
}

func _M64_formatBool(arg any) (out []rune) {
	switch v := arg.(type) {
	case *bool:
		return SN_FormatBOOL(*v, LETTERCASE_LOWER)
	case bool:
		return SN_FormatBOOL(v, LETTERCASE_LOWER)
	default:
		return []rune{'(', 'B', 'O', 'O', 'L', '=', 'b', 'a', 'd', ')'}
	}
}

func _S64_width_correction(output, width *[]rune) {
	x, _ := SN_ParseUINT(string((*width)), 10, 64)
	if uint64(len(*output)) >= x {
		return
	}

	diff := x - uint64(len(*output))
	for ; diff > 0; diff-- {
		*output = append(*output, ' ')
	}
}
