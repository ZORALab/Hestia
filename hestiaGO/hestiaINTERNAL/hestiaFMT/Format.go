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

package hestiaFMT

func Format(statement string, args ...any) string {
	engine := &engine{
		buffer:    []rune{},
		arguments: args,
		pos:       0,
	}
	engine.to_PARSE_NORMAL()

	for _, c := range statement {
		switch c {
		case '%':
			_formatVerb(engine)
		case 't':
			_formatBool(engine)
		case 'c':
			_formatChar(engine)
		case 's':
			_formatString(engine)
		case 'd':
			_formatDecimal(engine)
		case 'o', 'O':
			_formatOctet(engine, c)
		case 'x', 'X':
			_formatHexadecimal(engine, c)
		case 'b':
			_formatBinary(engine)
		case 'f', 'F':
			_formatFloat(engine, c, format_DECIMAL_NO_EXPONENT)
		case 'e', 'E':
			_formatFloat(engine, c, format_SCIENTIFIC)
		case 'g', 'G':
			_formatFloat(engine, c, format_SCIENTIFIC_AUTO_EXPONENT)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			_formatDigits(engine, c)
		case '.':
			_formatPeriod(engine)
		default:
			_formatDefault(engine, c)
		}
	}

	return string(engine.buffer)
}

func _formatDefault(engine *engine, c rune) {
	switch engine.mode {
	case engine_PARSE_VERB:
		// it is an unknown verb '%?'
		engine.buffer = append(engine.buffer, '(', 'E', 'R', 'R', 'O', 'R', '=')
		switch c {
		case '\a':
			engine.buffer = append(engine.buffer, '\\', 'a')
		case '\b':
			engine.buffer = append(engine.buffer, '\\', 'b')
		case '\t':
			engine.buffer = append(engine.buffer, '\\', 't')
		case '\n':
			engine.buffer = append(engine.buffer, '\\', 'n')
		case '\v':
			engine.buffer = append(engine.buffer, '\\', 'v')
		case '\f':
			engine.buffer = append(engine.buffer, '\\', 'f')
		case '\r':
			engine.buffer = append(engine.buffer, '\\', 'r')
		default:
			engine.buffer = append(engine.buffer, c)
		}
		engine.buffer = append(engine.buffer, ' ', '?', '?', '?', ')')

		engine.pos++
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular character
		engine.buffer = append(engine.buffer, c)
	}
}

func _formatPeriod(engine *engine) {
	switch engine.mode {
	case engine_PARSE_VERB:
		// change to precision assignment 'p' in '%{w.p}f' like '2' in '%10.2f'
		if engine.setPrecision {
			// error - multiple period denotes ('%{w..p...}f')
			engine.buffer = append(engine.buffer,
				'(', 'E', 'R', 'R', 'O', 'R', '=', '%', '.', '.', ')',
			)
			engine.pos++
			engine.to_PARSE_NORMAL()
		}

		engine.setPrecision = true
	case engine_PARSE_NORMAL:
		// just a regular '.' character
		engine.buffer = append(engine.buffer, '.')
	}
}

func _formatDigits(engine *engine, c rune) {
	switch engine.mode {
	case engine_PARSE_VERB:
		// assign to width or precision '%{w.p}f' like '10' or '2' in '%10.2f'
		if engine.setPrecision {
			engine.precision = append(engine.precision, c)
		} else {
			engine.width = append(engine.width, c)
		}
	case engine_PARSE_NORMAL:
		// just a regular digit character '0', '1', '2', ...
		engine.buffer = append(engine.buffer, c)
	}
}

func _formatFloat(engine *engine, c rune, format numberType) {
	var ret []rune

	// configure default charset
	charset := lettercase_LOWER
	if c == 'F' {
		charset = lettercase_UPPER
	}

	// perform formatting
	switch engine.mode {
	case engine_PARSE_VERB:
		// it's a float with no exponent verb '%F' or '%f'
		ret = _parseNumber(&numberSTR{
			arg:       engine.arg(),
			base:      10,
			charset:   charset,
			width:     engine.width,
			precision: engine.precision,
			format:    format,
		})
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 'F' or 'f' character
		engine.buffer = append(engine.buffer, c)
	}
}

func _formatBinary(engine *engine) {
	var ret []rune

	switch engine.mode {
	case engine_PARSE_VERB:
		// it's a binary verb '%b'
		ret = _parseNumber(&numberSTR{
			arg:       engine.arg(),
			base:      2,
			charset:   lettercase_LOWER,
			width:     engine.width,
			precision: engine.precision,
			format:    format_DECIMALLESS,
		})
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 'b' character
		engine.buffer = append(engine.buffer, 'b')
	}
}

func _formatHexadecimal(engine *engine, c rune) {
	var ret []rune

	// configure default charset
	charset := lettercase_LOWER
	if c == 'X' {
		charset = lettercase_UPPER
	}

	// perform formatting
	switch engine.mode {
	case engine_PARSE_VERB:
		// it's a hex verb '%X' or '%x'
		ret = _parseNumber(&numberSTR{
			arg:       engine.arg(),
			base:      16,
			charset:   charset,
			width:     engine.width,
			precision: engine.precision,
			format:    format_DECIMALLESS,
		})
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 'X' or 'x' character
		engine.buffer = append(engine.buffer, c)
	}
}

func _formatOctet(engine *engine, c rune) {
	var ret []rune

	// configure default charset
	charset := lettercase_LOWER
	if c == 'O' {
		charset = lettercase_UPPER
	}

	// perform formatting
	switch engine.mode {
	case engine_PARSE_VERB:
		// it's an octet verb '%O' or '%o'
		ret = _parseNumber(&numberSTR{
			arg:       engine.arg(),
			base:      8,
			charset:   charset,
			width:     engine.width,
			precision: engine.precision,
			format:    format_DECIMALLESS,
		})
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 'O' or 'o' character
		engine.buffer = append(engine.buffer, c)
	}
}

func _formatDecimal(engine *engine) {
	var ret []rune

	switch engine.mode {
	case engine_PARSE_VERB:
		// it's a decimal verb '%d'
		ret = _parseNumber(&numberSTR{
			arg:       engine.arg(),
			base:      10,
			charset:   lettercase_LOWER,
			width:     engine.width,
			precision: engine.precision,
			format:    format_DECIMALLESS,
		})
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 'd' character
		engine.buffer = append(engine.buffer, 'd')
	}
}

func _formatString(engine *engine) {
	var ret []rune

	switch engine.mode {
	case engine_PARSE_VERB:
		// it's a string verb '%s'
		ret = _parseString(engine.arg())
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 's' character
		engine.buffer = append(engine.buffer, 's')
	}
}

func _formatChar(engine *engine) {
	var ret []rune

	switch engine.mode {
	case engine_PARSE_VERB:
		// it's a char verb '%t'
		ret = _parseChar(engine.arg())
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 'c' character
		engine.buffer = append(engine.buffer, 'c')
	}
}

func _formatBool(engine *engine) {
	var ret []rune

	switch engine.mode {
	case engine_PARSE_VERB:
		// it's a bool verb '%t'
		ret = _parseBool(engine.arg())
		engine.buffer = append(engine.buffer, ret...)
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		// just a regular 't' character
		engine.buffer = append(engine.buffer, 't')
	}
}

func _formatVerb(engine *engine) {
	switch engine.mode {
	case engine_PARSE_VERB:
		// literal percent character ('%%')
		engine.buffer = append(engine.buffer, '%')
		engine.to_PARSE_NORMAL()
	case engine_PARSE_NORMAL:
		engine.mode = engine_PARSE_VERB
	}
}
