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

func Format(statement string, args ...any) string {
	var ret []rune

	engine := &engine{
		buffer:    []rune{},
		arguments: args,
		pos:       0,
	}
	engine.to_PARSE_NORMAL()

	for _, c := range statement {
		if engine.mode == engine_PARSE_NORMAL {
			if c == '%' {
				engine.mode = engine_PARSE_VERB
				continue
			}

			// parse as normal character otherwise and move on
			engine.buffer = append(engine.buffer, c)
			continue
		}

		// | beyond this point is verb mode (%c, %s, %d, ...) | //
		switch c {
		case '%': // literal percent character ('%%')
			ret = []rune{'%'}
		case 't': // bool verb '%t'
			ret = _formatBool(engine.arg())
		case 'c': // char verb '%t'
			ret = _formatChar(engine.arg())
		case 's': // string verb '%s'
			ret = _formatString(engine.arg())
		case 'd': // decimal verb '%d'
			ret = _formatNumber(&numberSTR{
				arg:       engine.arg(),
				base:      10,
				width:     engine.width,
				precision: engine.precision,
				notation:  NOTATION_DECIMALLESS,
				verb:      c,
			})
		case 'o', 'O': // octal verb '%o' or '%O'
			ret = _formatNumber(&numberSTR{
				arg:       engine.arg(),
				base:      10,
				width:     engine.width,
				precision: engine.precision,
				notation:  NOTATION_DECIMALLESS,
				verb:      c,
			})
		case 'x', 'X': // hexdecimal verb '%x or '%X'
			ret = _formatNumber(&numberSTR{
				arg:       engine.arg(),
				base:      16,
				width:     engine.width,
				precision: engine.precision,
				notation:  NOTATION_DECIMAL_NO_EXPONENT,
				verb:      c,
			})
		case 'b': // binary verb '%b'
			ret = _formatNumber(&numberSTR{
				arg:       engine.arg(),
				base:      2,
				width:     engine.width,
				precision: engine.precision,
				notation:  NOTATION_DECIMAL_NO_EXPONENT,
				verb:      c,
			})
		case 'f', 'F': // float with no exponent verb '%F' or '%f'
			ret = _formatNumber(&numberSTR{
				arg:       engine.arg(),
				base:      10,
				width:     engine.width,
				precision: engine.precision,
				notation:  NOTATION_DECIMAL_NO_EXPONENT,
				verb:      c,
			})
		case 'e', 'E': // float in scientific notation verb '%e' or '%E'
			ret = _formatNumber(&numberSTR{
				arg:       engine.arg(),
				base:      10,
				width:     engine.width,
				precision: engine.precision,
				notation:  NOTATION_ISO6093NR3,
				verb:      c,
			})
		case 'g', 'G': // float in auto-scientific notation verb '%g' or '%G'
			ret = _formatNumber(&numberSTR{
				arg:       engine.arg(),
				base:      10,
				width:     engine.width,
				precision: engine.precision,
				notation:  NOTATION_ISO6093NR3_AUTO,
				verb:      c,
			})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// assign to width/precision '%{w.p}f' like '10' or '2' in '%10.2f'
			if engine.setPrecision {
				engine.precision = append(engine.precision, c)
			} else {
				engine.width = append(engine.width, c)
			}

			continue
		case '.':
			// change to precision 'p' in '%{w.p}f' like '2' in '%10.2f'
			if !engine.setPrecision {
				engine.setPrecision = true
				continue
			}

			// bad verb - multiple period ('%{w..p...}f')
			ret = []rune{'(', 'E', 'R', 'R', 'O', 'R', '=', '%', '.', '.', ')'}
			_ = engine.arg()
		default:
			// unknown verb - '%?'
			ret = []rune{'(', 'E', 'R', 'R', 'O', 'R', '=', ' ', '\'', '%'}
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
			_ = engine.arg()
		}

		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	}

	return string(engine.buffer)
}

type numberSTR struct {
	arg       any
	base      uint8
	width     []rune
	precision []rune
	notation  Notation
	verb      rune
}

func _formatNumber(str *numberSTR) (out []rune) {
	// to be developed later
	return []rune{'(', 'N', 'U', 'M', 'B', 'E', 'R', '=', 'b', 'a', 'd', ')'}
}

func _formatString(arg any) (out []rune) {
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

func _formatChar(arg any) (out []rune) {
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
		return []rune{*v}
	case rune:
		return []rune{v}
	case *[]rune:
		return (*v)[:1]
	case []rune:
		return v[:1]
	default:
		return []rune{'(', 'C', 'H', 'A', 'R', '=', 'b', 'a', 'd', ')'}
	}
}

func _formatBool(arg any) (out []rune) {
	switch v := arg.(type) {
	case *bool:
		return SN_FormatBOOL(*v, LETTERCASE_LOWER)
	case bool:
		return SN_FormatBOOL(v, LETTERCASE_LOWER)
	default:
		return []rune{'(', 'B', 'O', 'O', 'L', '=', 'b', 'a', 'd', ')'}
	}
}

type engine struct {
	buffer       []rune
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
