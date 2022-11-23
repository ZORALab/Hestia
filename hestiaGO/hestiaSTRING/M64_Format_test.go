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
	"hestia/hestiaTESTING"
	"testing"
)

//nolint:goconst
func generate_M64_Format_params(caseID uint64) (desc, subject, expect string, args []any) {
	switch caseID {
	case 0:
		desc = `
test hestiaSTRING/M64_Format can process basic string.
`
		subject = "it's a basic string!"
		expect = subject
		args = nil
	case 1:
		desc = `
test hestiaSTRING/M64_Format can process string with arguments.
`
		subject = "a very %s string %s.\n"
		expect = "a very normal string as well.\n"
		args = []any{"normal", "as well"}
	case 2:
		desc = `
test hestiaSTRING/M64_Format can process string with proper fixed width string arguments.
`
		subject = "a very %5s string %10s.\n"
		expect = "a very normal string as well   .\n"
		args = []any{"normal", "as well"}
	case 3:
		desc = `
test hestiaSTRING/M64_Format can process string with proper precision string arguments.
`
		subject = "a very %.5s string %.10s.\n"
		expect = "a very normal string as well.\n"
		args = []any{"normal", "as well"}
	case 4:
		desc = `
test hestiaSTRING/M64_Format can process string with insufficient string arguments.
`
		subject = "a very %s string %s.\n"
		expect = "a very normal string (STRING=bad).\n"
		args = []any{"normal"}
	case 5:
		desc = `
test hestiaSTRING/M64_Format can process string with invalid string arguments.
`
		subject = "a very %s string %s.\n"
		expect = "a very (STRING=bad) string (STRING=bad).\n"
		args = []any{0, 123.4567}
	case 6:
		desc = `
test hestiaSTRING/M64_Format can process string with proper fixed width arguments.
`
		subject = "a very %5s string %10s.\n"
		expect = "a very normal string as well   .\n"
		args = []any{"normal", "as well"}
	case 7:
		desc = `
test hestiaSTRING/M64_Format can process string with oversupplied string arguments.
`
		subject = "a very %s string %s.\n"
		expect = "a very normal string as well.\n"
		args = []any{"normal", "as well", "over?"}
	case 8:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint8 normal arguments.
`
		subject = "the uint8 is %d and %d.\n"
		expect = "the uint8 is 1 and 2.\n"
		args = []any{uint8(1), uint8(2)}
	case 9:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint8 invalid arguments.
`
		subject = "the uint8 is %d and %d.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 10:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint8 empty argument.
`
		subject = "the uint8 is %d and %d.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 11:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint8 proper fixed width arguments.
`
		subject = "the uint8 is %5d and %10d.\n"
		expect = "the uint8 is 1     and 2         .\n"
		args = []any{uint8(1), uint8(2)}
	case 12:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint8 with proper precision arguments.
`
		subject = "the uint8 is %.5d and %.10d.\n"
		expect = "the uint8 is 1 and 2.\n"
		args = []any{uint8(1), uint8(2)}
	case 13:
		desc = `
test hestiaSTRING/M64_Format can process float uint8 normal arguments.
`
		subject = "the uint8 is %f and %f.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint8(1), uint8(2)}
	case 14:
		desc = `
test hestiaSTRING/M64_Format can process float uint8 invalid arguments.
`
		subject = "the uint8 is %f and %f.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 15:
		desc = `
test hestiaSTRING/M64_Format can process float uint8 empty argument.
`
		subject = "the uint8 is %f and %f.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 16:
		desc = `
test hestiaSTRING/M64_Format can process float uint8 proper fixed width arguments.
`
		subject = "the uint8 is %5f and %10f.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint8(1), uint8(2)}
	case 17:
		desc = `
test hestiaSTRING/M64_Format can process float uint8 with proper precision arguments.
`
		subject = "the uint8 is %.5f and %.10f.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint8(1), uint8(2)}
	case 18:
		desc = `
test hestiaSTRING/M64_Format can process octal uint8 normal arguments.
`
		subject = "the uint8 is %o and %o.\n"
		expect = "the uint8 is 1 and 2.\n"
		args = []any{uint8(1), uint8(2)}
	case 19:
		desc = `
test hestiaSTRING/M64_Format can process octal uint8 invalid arguments.
`
		subject = "the uint8 is %o and %o.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 20:
		desc = `
test hestiaSTRING/M64_Format can process octal uint8 empty argument.
`
		subject = "the uint8 is %o and %o.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 21:
		desc = `
test hestiaSTRING/M64_Format can process octal uint8 proper fixed width arguments.
`
		subject = "the uint8 is %5o and %10o.\n"
		expect = "the uint8 is 1     and 2         .\n"
		args = []any{uint8(1), uint8(2)}
	case 22:
		desc = `
test hestiaSTRING/M64_Format can process octal uint8 with proper precision arguments.
`
		subject = "the uint8 is %.5o and %.10o.\n"
		expect = "the uint8 is 1 and 2.\n"
		args = []any{uint8(1), uint8(2)}
	case 23:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint8 normal arguments.
`
		subject = "the uint8 is %x and %x.\n"
		expect = "the uint8 is 1 and 2.\n"
		args = []any{uint8(1), uint8(2)}
	case 24:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint8 invalid arguments.
`
		subject = "the uint8 is %x and %x.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 25:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint8 empty argument.
`
		subject = "the uint8 is %x and %x.\n"
		expect = "the uint8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 26:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint8 proper fixed width arguments.
`
		subject = "the uint8 is %5x and %10x.\n"
		expect = "the uint8 is 1     and 2         .\n"
		args = []any{uint8(1), uint8(2)}
	case 27:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint8 with proper precision arguments.
`
		subject = "the uint8 is %.5x and %.10x.\n"
		expect = "the uint8 is 1 and 2.\n"
		args = []any{uint8(1), uint8(2)}
	case 28:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint16 normal arguments.
`
		subject = "the uint16 is %d and %d.\n"
		expect = "the uint16 is 1 and 2.\n"
		args = []any{uint16(1), uint16(2)}
	case 29:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint16 invalid arguments.
`
		subject = "the uint16 is %d and %d.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 30:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint16 empty argument.
`
		subject = "the uint16 is %d and %d.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 31:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint16 proper fixed width arguments.
`
		subject = "the uint16 is %5d and %10d.\n"
		expect = "the uint16 is 1     and 2         .\n"
		args = []any{uint16(1), uint16(2)}
	case 32:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint16 with proper precision arguments.
`
		subject = "the uint16 is %.5d and %.10d.\n"
		expect = "the uint16 is 1 and 2.\n"
		args = []any{uint16(1), uint16(2)}
	case 33:
		desc = `
test hestiaSTRING/M64_Format can process float uint16 normal arguments.
`
		subject = "the uint16 is %f and %f.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint16(1), uint16(2)}
	case 34:
		desc = `
test hestiaSTRING/M64_Format can process float uint16 invalid arguments.
`
		subject = "the uint16 is %f and %f.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 35:
		desc = `
test hestiaSTRING/M64_Format can process float uint16 empty argument.
`
		subject = "the uint16 is %f and %f.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 36:
		desc = `
test hestiaSTRING/M64_Format can process float uint16 proper fixed width arguments.
`
		subject = "the uint16 is %5f and %10f.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint16(1), uint16(2)}
	case 37:
		desc = `
test hestiaSTRING/M64_Format can process float uint16 with proper precision arguments.
`
		subject = "the uint16 is %.5f and %.10f.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint16(1), uint16(2)}
	case 38:
		desc = `
test hestiaSTRING/M64_Format can process octal uint16 normal arguments.
`
		subject = "the uint16 is %o and %o.\n"
		expect = "the uint16 is 1 and 2.\n"
		args = []any{uint16(1), uint16(2)}
	case 39:
		desc = `
test hestiaSTRING/M64_Format can process octal uint16 invalid arguments.
`
		subject = "the uint16 is %o and %o.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 40:
		desc = `
test hestiaSTRING/M64_Format can process octal uint16 empty argument.
`
		subject = "the uint16 is %o and %o.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 41:
		desc = `
test hestiaSTRING/M64_Format can process octal uint16 proper fixed width arguments.
`
		subject = "the uint16 is %5o and %10o.\n"
		expect = "the uint16 is 1     and 2         .\n"
		args = []any{uint16(1), uint16(2)}
	case 42:
		desc = `
test hestiaSTRING/M64_Format can process octal uint16 with proper precision arguments.
`
		subject = "the uint16 is %.5o and %.10o.\n"
		expect = "the uint16 is 1 and 2.\n"
		args = []any{uint16(1), uint16(2)}
	case 43:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint16 normal arguments.
`
		subject = "the uint16 is %x and %x.\n"
		expect = "the uint16 is 1 and 2.\n"
		args = []any{uint16(1), uint16(2)}
	case 44:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint16 invalid arguments.
`
		subject = "the uint16 is %x and %x.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 45:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint16 empty argument.
`
		subject = "the uint16 is %x and %x.\n"
		expect = "the uint16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 46:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint16 proper fixed width arguments.
`
		subject = "the uint16 is %5x and %10x.\n"
		expect = "the uint16 is 1     and 2         .\n"
		args = []any{uint16(1), uint16(2)}
	case 47:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint16 with proper precision arguments.
`
		subject = "the uint16 is %.5x and %.10x.\n"
		expect = "the uint16 is 1 and 2.\n"
		args = []any{uint16(1), uint16(2)}
	case 48:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint32 normal arguments.
`
		subject = "the uint32 is %d and %d.\n"
		expect = "the uint32 is 1 and 2.\n"
		args = []any{uint32(1), uint32(2)}
	case 49:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint32 invalid arguments.
`
		subject = "the uint32 is %d and %d.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 50:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint32 empty argument.
`
		subject = "the uint32 is %d and %d.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 51:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint32 proper fixed width arguments.
`
		subject = "the uint32 is %5d and %10d.\n"
		expect = "the uint32 is 1     and 2         .\n"
		args = []any{uint32(1), uint32(2)}
	case 52:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint32 with proper precision arguments.
`
		subject = "the uint32 is %.5d and %.10d.\n"
		expect = "the uint32 is 1 and 2.\n"
		args = []any{uint32(1), uint32(2)}
	case 53:
		desc = `
test hestiaSTRING/M64_Format can process float uint32 normal arguments.
`
		subject = "the uint32 is %f and %f.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint32(1), uint32(2)}
	case 54:
		desc = `
test hestiaSTRING/M64_Format can process float uint32 invalid arguments.
`
		subject = "the uint32 is %f and %f.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 55:
		desc = `
test hestiaSTRING/M64_Format can process float uint32 empty argument.
`
		subject = "the uint32 is %f and %f.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 56:
		desc = `
test hestiaSTRING/M64_Format can process float uint32 proper fixed width arguments.
`
		subject = "the uint32 is %5f and %10f.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint32(1), uint32(2)}
	case 57:
		desc = `
test hestiaSTRING/M64_Format can process float uint32 with proper precision arguments.
`
		subject = "the uint32 is %.5f and %.10f.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint32(1), uint32(2)}
	case 58:
		desc = `
test hestiaSTRING/M64_Format can process octal uint32 normal arguments.
`
		subject = "the uint32 is %o and %o.\n"
		expect = "the uint32 is 1 and 2.\n"
		args = []any{uint32(1), uint32(2)}
	case 59:
		desc = `
test hestiaSTRING/M64_Format can process octal uint32 invalid arguments.
`
		subject = "the uint32 is %o and %o.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 60:
		desc = `
test hestiaSTRING/M64_Format can process octal uint32 empty argument.
`
		subject = "the uint32 is %o and %o.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 61:
		desc = `
test hestiaSTRING/M64_Format can process octal uint32 proper fixed width arguments.
`
		subject = "the uint32 is %5o and %10o.\n"
		expect = "the uint32 is 1     and 2         .\n"
		args = []any{uint32(1), uint32(2)}
	case 62:
		desc = `
test hestiaSTRING/M64_Format can process octal uint32 with proper precision arguments.
`
		subject = "the uint32 is %.5o and %.10o.\n"
		expect = "the uint32 is 1 and 2.\n"
		args = []any{uint32(1), uint32(2)}
	case 63:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint32 normal arguments.
`
		subject = "the uint32 is %x and %x.\n"
		expect = "the uint32 is 1 and 2.\n"
		args = []any{uint32(1), uint32(2)}
	case 64:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint32 invalid arguments.
`
		subject = "the uint32 is %x and %x.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 65:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint32 empty argument.
`
		subject = "the uint32 is %x and %x.\n"
		expect = "the uint32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 66:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint32 proper fixed width arguments.
`
		subject = "the uint32 is %5x and %10x.\n"
		expect = "the uint32 is 1     and 2         .\n"
		args = []any{uint32(1), uint32(2)}
	case 67:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint32 with proper precision arguments.
`
		subject = "the uint32 is %.5x and %.10x.\n"
		expect = "the uint32 is 1 and 2.\n"
		args = []any{uint32(1), uint32(2)}
	case 68:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint64 normal arguments.
`
		subject = "the uint64 is %d and %d.\n"
		expect = "the uint64 is 1 and 2.\n"
		args = []any{uint64(1), uint64(2)}
	case 69:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint64 invalid arguments.
`
		subject = "the uint64 is %d and %d.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 70:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint64 empty argument.
`
		subject = "the uint64 is %d and %d.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 71:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint64 proper fixed width arguments.
`
		subject = "the uint64 is %5d and %10d.\n"
		expect = "the uint64 is 1     and 2         .\n"
		args = []any{uint64(1), uint64(2)}
	case 72:
		desc = `
test hestiaSTRING/M64_Format can process decimal uint64 with proper precision arguments.
`
		subject = "the uint64 is %.5d and %.10d.\n"
		expect = "the uint64 is 1 and 2.\n"
		args = []any{uint64(1), uint64(2)}
	case 73:
		desc = `
test hestiaSTRING/M64_Format can process float uint64 normal arguments.
`
		subject = "the uint64 is %f and %f.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint64(1), uint64(2)}
	case 74:
		desc = `
test hestiaSTRING/M64_Format can process float uint64 invalid arguments.
`
		subject = "the uint64 is %f and %f.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 75:
		desc = `
test hestiaSTRING/M64_Format can process float uint64 empty argument.
`
		subject = "the uint64 is %f and %f.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 76:
		desc = `
test hestiaSTRING/M64_Format can process float uint64 proper fixed width arguments.
`
		subject = "the uint64 is %5f and %10f.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint64(1), uint64(2)}
	case 77:
		desc = `
test hestiaSTRING/M64_Format can process float uint64 with proper precision arguments.
`
		subject = "the uint64 is %.5f and %.10f.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{uint64(1), uint64(2)}
	case 78:
		desc = `
test hestiaSTRING/M64_Format can process octal uint64 normal arguments.
`
		subject = "the uint64 is %o and %o.\n"
		expect = "the uint64 is 1 and 2.\n"
		args = []any{uint64(1), uint64(2)}
	case 79:
		desc = `
test hestiaSTRING/M64_Format can process octal uint64 invalid arguments.
`
		subject = "the uint64 is %o and %o.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 80:
		desc = `
test hestiaSTRING/M64_Format can process octal uint64 empty argument.
`
		subject = "the uint64 is %o and %o.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 81:
		desc = `
test hestiaSTRING/M64_Format can process octal uint64 proper fixed width arguments.
`
		subject = "the uint64 is %5o and %10o.\n"
		expect = "the uint64 is 1     and 2         .\n"
		args = []any{uint64(1), uint64(2)}
	case 82:
		desc = `
test hestiaSTRING/M64_Format can process octal uint64 with proper precision arguments.
`
		subject = "the uint64 is %.5o and %.10o.\n"
		expect = "the uint64 is 1 and 2.\n"
		args = []any{uint64(1), uint64(2)}
	case 83:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint64 normal arguments.
`
		subject = "the uint64 is %x and %x.\n"
		expect = "the uint64 is 1 and 2.\n"
		args = []any{uint64(1), uint64(2)}
	case 84:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint64 invalid arguments.
`
		subject = "the uint64 is %x and %x.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 85:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint64 empty argument.
`
		subject = "the uint64 is %x and %x.\n"
		expect = "the uint64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 86:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint64 proper fixed width arguments.
`
		subject = "the uint64 is %5x and %10x.\n"
		expect = "the uint64 is 1     and 2         .\n"
		args = []any{uint64(1), uint64(2)}
	case 87:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal uint64 with proper precision arguments.
`
		subject = "the uint64 is %.5x and %.10x.\n"
		expect = "the uint64 is 1 and 2.\n"
		args = []any{uint64(1), uint64(2)}
	case 88:
		desc = `
test hestiaSTRING/M64_Format can process decimal int8 normal arguments.
`
		subject = "the int8 is %d and %d.\n"
		expect = "the int8 is 1 and 2.\n"
		args = []any{int8(1), int8(2)}
	case 89:
		desc = `
test hestiaSTRING/M64_Format can process decimal int8 invalid arguments.
`
		subject = "the int8 is %d and %d.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 90:
		desc = `
test hestiaSTRING/M64_Format can process decimal int8 empty argument.
`
		subject = "the int8 is %d and %d.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 91:
		desc = `
test hestiaSTRING/M64_Format can process decimal int8 proper fixed width arguments.
`
		subject = "the int8 is %5d and %10d.\n"
		expect = "the int8 is 1     and 2         .\n"
		args = []any{int8(1), int8(2)}
	case 92:
		desc = `
test hestiaSTRING/M64_Format can process decimal int8 with proper precision arguments.
`
		subject = "the int8 is %.5d and %.10d.\n"
		expect = "the int8 is 1 and 2.\n"
		args = []any{int8(1), int8(2)}
	case 93:
		desc = `
test hestiaSTRING/M64_Format can process float int8 normal arguments.
`
		subject = "the int8 is %f and %f.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int8(1), int8(2)}
	case 94:
		desc = `
test hestiaSTRING/M64_Format can process float int8 invalid arguments.
`
		subject = "the int8 is %f and %f.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 95:
		desc = `
test hestiaSTRING/M64_Format can process float int8 empty argument.
`
		subject = "the int8 is %f and %f.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 96:
		desc = `
test hestiaSTRING/M64_Format can process float int8 proper fixed width arguments.
`
		subject = "the int8 is %5f and %10f.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int8(1), int8(2)}
	case 97:
		desc = `
test hestiaSTRING/M64_Format can process float int8 with proper precision arguments.
`
		subject = "the int8 is %.5f and %.10f.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int8(1), int8(2)}
	case 98:
		desc = `
test hestiaSTRING/M64_Format can process octal int8 normal arguments.
`
		subject = "the int8 is %o and %o.\n"
		expect = "the int8 is 1 and 2.\n"
		args = []any{int8(1), int8(2)}
	case 99:
		desc = `
test hestiaSTRING/M64_Format can process octal int8 invalid arguments.
`
		subject = "the int8 is %o and %o.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 100:
		desc = `
test hestiaSTRING/M64_Format can process octal int8 empty argument.
`
		subject = "the int8 is %o and %o.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 101:
		desc = `
test hestiaSTRING/M64_Format can process octal int8 proper fixed width arguments.
`
		subject = "the int8 is %5o and %10o.\n"
		expect = "the int8 is 1     and 2         .\n"
		args = []any{int8(1), int8(2)}
	case 102:
		desc = `
test hestiaSTRING/M64_Format can process octal int8 with proper precision arguments.
`
		subject = "the int8 is %.5o and %.10o.\n"
		expect = "the int8 is 1 and 2.\n"
		args = []any{int8(1), int8(2)}
	case 103:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int8 normal arguments.
`
		subject = "the int8 is %x and %x.\n"
		expect = "the int8 is 1 and 2.\n"
		args = []any{int8(1), int8(2)}
	case 104:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int8 invalid arguments.
`
		subject = "the int8 is %x and %x.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 105:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int8 empty argument.
`
		subject = "the int8 is %x and %x.\n"
		expect = "the int8 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 106:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int8 proper fixed width arguments.
`
		subject = "the int8 is %5x and %10x.\n"
		expect = "the int8 is 1     and 2         .\n"
		args = []any{int8(1), int8(2)}
	case 107:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int8 with proper precision arguments.
`
		subject = "the int8 is %.5x and %.10x.\n"
		expect = "the int8 is 1 and 2.\n"
		args = []any{int8(1), int8(2)}
	case 108:
		desc = `
test hestiaSTRING/M64_Format can process decimal int16 normal arguments.
`
		subject = "the int16 is %d and %d.\n"
		expect = "the int16 is 1 and 2.\n"
		args = []any{int16(1), int16(2)}
	case 109:
		desc = `
test hestiaSTRING/M64_Format can process decimal int16 invalid arguments.
`
		subject = "the int16 is %d and %d.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 110:
		desc = `
test hestiaSTRING/M64_Format can process decimal int16 empty argument.
`
		subject = "the int16 is %d and %d.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 111:
		desc = `
test hestiaSTRING/M64_Format can process decimal int16 proper fixed width arguments.
`
		subject = "the int16 is %5d and %10d.\n"
		expect = "the int16 is 1     and 2         .\n"
		args = []any{int16(1), int16(2)}
	case 112:
		desc = `
test hestiaSTRING/M64_Format can process decimal int16 with proper precision arguments.
`
		subject = "the int16 is %.5d and %.10d.\n"
		expect = "the int16 is 1 and 2.\n"
		args = []any{int16(1), int16(2)}
	case 113:
		desc = `
test hestiaSTRING/M64_Format can process float int16 normal arguments.
`
		subject = "the int16 is %f and %f.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int16(1), int16(2)}
	case 114:
		desc = `
test hestiaSTRING/M64_Format can process float int16 invalid arguments.
`
		subject = "the int16 is %f and %f.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 115:
		desc = `
test hestiaSTRING/M64_Format can process float int16 empty argument.
`
		subject = "the int16 is %f and %f.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 116:
		desc = `
test hestiaSTRING/M64_Format can process float int16 proper fixed width arguments.
`
		subject = "the int16 is %5f and %10f.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int16(1), int16(2)}
	case 117:
		desc = `
test hestiaSTRING/M64_Format can process float int16 with proper precision arguments.
`
		subject = "the int16 is %.5f and %.10f.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int16(1), int16(2)}
	case 118:
		desc = `
test hestiaSTRING/M64_Format can process octal int16 normal arguments.
`
		subject = "the int16 is %o and %o.\n"
		expect = "the int16 is 1 and 2.\n"
		args = []any{int16(1), int16(2)}
	case 119:
		desc = `
test hestiaSTRING/M64_Format can process octal int16 invalid arguments.
`
		subject = "the int16 is %o and %o.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 120:
		desc = `
test hestiaSTRING/M64_Format can process octal int16 empty argument.
`
		subject = "the int16 is %o and %o.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 121:
		desc = `
test hestiaSTRING/M64_Format can process octal int16 proper fixed width arguments.
`
		subject = "the int16 is %5o and %10o.\n"
		expect = "the int16 is 1     and 2         .\n"
		args = []any{int16(1), int16(2)}
	case 122:
		desc = `
test hestiaSTRING/M64_Format can process octal int16 with proper precision arguments.
`
		subject = "the int16 is %.5o and %.10o.\n"
		expect = "the int16 is 1 and 2.\n"
		args = []any{int16(1), int16(2)}
	case 123:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int16 normal arguments.
`
		subject = "the int16 is %x and %x.\n"
		expect = "the int16 is 1 and 2.\n"
		args = []any{uint16(1), uint16(2)}
	case 124:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int16 invalid arguments.
`
		subject = "the int16 is %x and %x.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 125:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int16 empty argument.
`
		subject = "the int16 is %x and %x.\n"
		expect = "the int16 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 126:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int16 proper fixed width arguments.
`
		subject = "the int16 is %5x and %10x.\n"
		expect = "the int16 is 1     and 2         .\n"
		args = []any{int16(1), int16(2)}
	case 127:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int16 with proper precision arguments.
`
		subject = "the int16 is %.5x and %.10x.\n"
		expect = "the int16 is 1 and 2.\n"
		args = []any{int16(1), int16(2)}
	case 128:
		desc = `
test hestiaSTRING/M64_Format can process decimal int32 normal arguments.
`
		subject = "the int32 is %d and %d.\n"
		expect = "the int32 is 1 and 2.\n"
		args = []any{int32(1), int32(2)}
	case 129:
		desc = `
test hestiaSTRING/M64_Format can process decimal int32 invalid arguments.
`
		subject = "the int32 is %d and %d.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 130:
		desc = `
test hestiaSTRING/M64_Format can process decimal int32 empty argument.
`
		subject = "the int32 is %d and %d.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 131:
		desc = `
test hestiaSTRING/M64_Format can process decimal int32 proper fixed width arguments.
`
		subject = "the int32 is %5d and %10d.\n"
		expect = "the int32 is 1     and 2         .\n"
		args = []any{int32(1), int32(2)}
	case 132:
		desc = `
test hestiaSTRING/M64_Format can process decimal int32 with proper precision arguments.
`
		subject = "the int32 is %.5d and %.10d.\n"
		expect = "the int32 is 1 and 2.\n"
		args = []any{int32(1), int32(2)}
	case 133:
		desc = `
test hestiaSTRING/M64_Format can process float int32 normal arguments.
`
		subject = "the int32 is %f and %f.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int32(1), int32(2)}
	case 134:
		desc = `
test hestiaSTRING/M64_Format can process float int32 invalid arguments.
`
		subject = "the int32 is %f and %f.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 135:
		desc = `
test hestiaSTRING/M64_Format can process float int32 empty argument.
`
		subject = "the int32 is %f and %f.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 136:
		desc = `
test hestiaSTRING/M64_Format can process float int32 proper fixed width arguments.
`
		subject = "the int32 is %5f and %10f.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int32(1), int32(2)}
	case 137:
		desc = `
test hestiaSTRING/M64_Format can process float int32 with proper precision arguments.
`
		subject = "the int32 is %.5f and %.10f.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int32(1), int32(2)}
	case 138:
		desc = `
test hestiaSTRING/M64_Format can process octal int32 normal arguments.
`
		subject = "the int32 is %o and %o.\n"
		expect = "the int32 is 1 and 2.\n"
		args = []any{int32(1), int32(2)}
	case 139:
		desc = `
test hestiaSTRING/M64_Format can process octal int32 invalid arguments.
`
		subject = "the int32 is %o and %o.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 140:
		desc = `
test hestiaSTRING/M64_Format can process octal int32 empty argument.
`
		subject = "the int32 is %o and %o.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 141:
		desc = `
test hestiaSTRING/M64_Format can process octal int32 proper fixed width arguments.
`
		subject = "the int32 is %5o and %10o.\n"
		expect = "the int32 is 1     and 2         .\n"
		args = []any{int32(1), int32(2)}
	case 142:
		desc = `
test hestiaSTRING/M64_Format can process octal int32 with proper precision arguments.
`
		subject = "the int32 is %.5o and %.10o.\n"
		expect = "the int32 is 1 and 2.\n"
		args = []any{int32(1), int32(2)}
	case 143:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int32 normal arguments.
`
		subject = "the int32 is %x and %x.\n"
		expect = "the int32 is 1 and 2.\n"
		args = []any{int32(1), int32(2)}
	case 144:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int32 invalid arguments.
`
		subject = "the int32 is %x and %x.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 145:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int32 empty argument.
`
		subject = "the int32 is %x and %x.\n"
		expect = "the int32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 146:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int32 proper fixed width arguments.
`
		subject = "the int32 is %5x and %10x.\n"
		expect = "the int32 is 1     and 2         .\n"
		args = []any{int32(1), int32(2)}
	case 147:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int32 with proper precision arguments.
`
		subject = "the int32 is %.5x and %.10x.\n"
		expect = "the int32 is 1 and 2.\n"
		args = []any{int32(1), int32(2)}
	case 148:
		desc = `
test hestiaSTRING/M64_Format can process decimal int64 normal arguments.
`
		subject = "the int64 is %d and %d.\n"
		expect = "the int64 is 1 and 2.\n"
		args = []any{int64(1), int64(2)}
	case 149:
		desc = `
test hestiaSTRING/M64_Format can process decimal int64 invalid arguments.
`
		subject = "the int64 is %d and %d.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 150:
		desc = `
test hestiaSTRING/M64_Format can process decimal int64 empty argument.
`
		subject = "the int64 is %d and %d.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 151:
		desc = `
test hestiaSTRING/M64_Format can process decimal int64 proper fixed width arguments.
`
		subject = "the int64 is %5d and %10d.\n"
		expect = "the int64 is 1     and 2         .\n"
		args = []any{int64(1), int64(2)}
	case 152:
		desc = `
test hestiaSTRING/M64_Format can process decimal int64 with proper precision arguments.
`
		subject = "the int64 is %.5d and %.10d.\n"
		expect = "the int64 is 1 and 2.\n"
		args = []any{int64(1), int64(2)}
	case 153:
		desc = `
test hestiaSTRING/M64_Format can process float int64 normal arguments.
`
		subject = "the int64 is %f and %f.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int64(1), int64(2)}
	case 154:
		desc = `
test hestiaSTRING/M64_Format can process float int64 invalid arguments.
`
		subject = "the int64 is %f and %f.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 155:
		desc = `
test hestiaSTRING/M64_Format can process float int64 empty argument.
`
		subject = "the int64 is %f and %f.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 156:
		desc = `
test hestiaSTRING/M64_Format can process float int64 proper fixed width arguments.
`
		subject = "the int64 is %5f and %10f.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int64(1), int64(2)}
	case 157:
		desc = `
test hestiaSTRING/M64_Format can process float int64 with proper precision arguments.
`
		subject = "the int64 is %.5f and %.10f.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{int64(1), int64(2)}
	case 158:
		desc = `
test hestiaSTRING/M64_Format can process octal int64 normal arguments.
`
		subject = "the int64 is %o and %o.\n"
		expect = "the int64 is 1 and 2.\n"
		args = []any{int64(1), int64(2)}
	case 159:
		desc = `
test hestiaSTRING/M64_Format can process octal int64 invalid arguments.
`
		subject = "the int64 is %o and %o.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 160:
		desc = `
test hestiaSTRING/M64_Format can process octal int64 empty argument.
`
		subject = "the int64 is %o and %o.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 161:
		desc = `
test hestiaSTRING/M64_Format can process octal int64 proper fixed width arguments.
`
		subject = "the int64 is %5o and %10o.\n"
		expect = "the int64 is 1     and 2         .\n"
		args = []any{int64(1), int64(2)}
	case 162:
		desc = `
test hestiaSTRING/M64_Format can process octal int64 with proper precision arguments.
`
		subject = "the int64 is %.5o and %.10o.\n"
		expect = "the int64 is 1 and 2.\n"
		args = []any{int64(1), int64(2)}
	case 163:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int64 normal arguments.
`
		subject = "the int64 is %x and %x.\n"
		expect = "the int64 is 1 and 2.\n"
		args = []any{int64(1), int64(2)}
	case 164:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int64 invalid arguments.
`
		subject = "the int64 is %x and %x.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{"1", "2"}
	case 165:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int64 empty argument.
`
		subject = "the int64 is %x and %x.\n"
		expect = "the int64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{}
	case 166:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int64 proper fixed width arguments.
`
		subject = "the int64 is %5x and %10x.\n"
		expect = "the int64 is 1     and 2         .\n"
		args = []any{int64(1), int64(2)}
	case 167:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal int64 with proper precision arguments.
`
		subject = "the int64 is %.5x and %.10x.\n"
		expect = "the int64 is 1 and 2.\n"
		args = []any{int64(1), int64(2)}
	case 168:
		desc = `
test hestiaSTRING/M64_Format can process decimal float32 normal arguments.
`
		subject = "the float32 is %d and %d.\n"
		expect = "the float32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 169:
		desc = `
test hestiaSTRING/M64_Format can process no-exponent float32 normal arguments.
`
		subject = "the float32 is %f and %f.\n"
		expect = "the float32 is 10000000000.000000 and 20000000000.000000.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 170:
		desc = `
test hestiaSTRING/M64_Format can process no-exponent float32 normal arguments.
`
		subject = "the float32 is %F and %F.\n"
		expect = "the float32 is 10000000000.000000 and 20000000000.000000.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 171:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 float32 normal arguments.
`
		subject = "the float32 is %e and %e.\n"
		expect = "the float32 is 1.000000e+10 and 2.000000e+10.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 172:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 float32 normal arguments.
`
		subject = "the float32 is %E and %E.\n"
		expect = "the float32 is 1.000000E+10 and 2.000000E+10.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 173:
		desc = `
test hestiaSTRING/M64_Format can process auto-exponent float32 normal arguments.
`
		subject = "the float32 is %g and %g.\n"
		expect = "the float32 is 1.000000e+10 and 2.000000e+10.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 174:
		desc = `
test hestiaSTRING/M64_Format can process auto-exponent float32 normal arguments.
`
		subject = "the float32 is %G and %G.\n"
		expect = "the float32 is 1.000000E+10 and 2.000000E+10.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 175:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal float32 normal arguments.
`
		subject = "the float32 is %x and %x.\n"
		expect = "the float32 is 0x1.000000p+10 and 0x2.000000p+10.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 176:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal float32 normal arguments.
`
		subject = "the float32 is %X and %X.\n"
		expect = "the float32 is 0X1.000000P+10 and 0X2.000000P+10.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 177:
		desc = `
test hestiaSTRING/M64_Format can process binary float32 normal arguments.
`
		subject = "the float32 is %b and %b.\n"
		expect = "the float32 is 1.000000p+10 and 1.000000p+11.\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 178:
		desc = `
test hestiaSTRING/M64_Format can process string float32 normal arguments.
`
		subject = "the float32 is %s and %s.\n"
		expect = "the float32 is (STRING=bad) and (STRING=bad).\n"
		args = []any{float32(1.0e10), float32(2.0e10)}
	case 179:
		desc = `
test hestiaSTRING/M64_Format can process decimal float64 normal arguments.
`
		subject = "the float64 is %d and %d.\n"
		expect = "the float64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 180:
		desc = `
test hestiaSTRING/M64_Format can process no-exponent float64 normal arguments.
`
		subject = "the float64 is %f and %f.\n"
		expect = "the float64 is 10000000000.000000 and 20000000000.000000.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 181:
		desc = `
test hestiaSTRING/M64_Format can process no-exponent float64 normal arguments.
`
		subject = "the float64 is %F and %F.\n"
		expect = "the float64 is 10000000000.000000 and 20000000000.000000.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 182:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 float64 normal arguments.
`
		subject = "the float64 is %e and %e.\n"
		expect = "the float64 is 1.000000e+10 and 2.000000e+10.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 183:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 float64 normal arguments.
`
		subject = "the float64 is %E and %E.\n"
		expect = "the float64 is 1.000000E+10 and 2.000000E+10.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 184:
		desc = `
test hestiaSTRING/M64_Format can process auto-exponent float64 normal arguments.
`
		subject = "the float64 is %g and %g.\n"
		expect = "the float64 is 1.000000e+10 and 2.000000e+10.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 185:
		desc = `
test hestiaSTRING/M64_Format can process auto-exponent float64 normal arguments.
`
		subject = "the float64 is %G and %G.\n"
		expect = "the float64 is 1.000000E+10 and 2.000000E+10.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 186:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal float64 normal arguments.
`
		subject = "the float64 is %x and %x.\n"
		expect = "the float64 is 0x1.000000p+10 and 0x2.000000p+10.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 187:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal float64 normal arguments.
`
		subject = "the float64 is %X and %X.\n"
		expect = "the float64 is 0X1.000000P+10 and 0X2.000000P+10.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 188:
		desc = `
test hestiaSTRING/M64_Format can process binary float64 normal arguments.
`
		subject = "the float64 is %b and %b.\n"
		expect = "the float64 is 1.000000p+10 and 1.000000p+11.\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 189:
		desc = `
test hestiaSTRING/M64_Format can process string float64 normal arguments.
`
		subject = "the float64 is %s and %s.\n"
		expect = "the float64 is (STRING=bad) and (STRING=bad).\n"
		args = []any{float64(1.0e10), float64(2.0e10)}
	case 190:
		desc = `
test hestiaSTRING/M64_Format can process bool boolean normal arguments.
`
		subject = "the bool is %t and %t.\n"
		expect = "the bool is true and false.\n"
		args = []any{true, false}
	case 191:
		desc = `
test hestiaSTRING/M64_Format can process string boolean normal arguments.
`
		subject = "the bool is %s and %s.\n"
		expect = "the bool is (STRING=bad) and (STRING=bad).\n"
		args = []any{true, false}
	case 192:
		desc = `
test hestiaSTRING/M64_Format can process decimal boolean normal arguments.
`
		subject = "the bool is %d and %d.\n"
		expect = "the bool is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{true, false}
	case 193:
		desc = `
test hestiaSTRING/M64_Format can process no-exponent boolean normal arguments.
`
		subject = "the bool is %f and %F.\n"
		expect = "the bool is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{true, false}
	case 194:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 boolean normal arguments.
`
		subject = "the bool is %e and %E.\n"
		expect = "the bool is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{true, false}
	case 195:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 boolean normal arguments.
`
		subject = "the bool is %g and %G.\n"
		expect = "the bool is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{true, false}
	case 196:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal boolean normal arguments.
`
		subject = "the bool is %x and %X.\n"
		expect = "the bool is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{true, false}
	case 197:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal uint8 normal arguments.
`
		subject = "the uint8 is %O and %O.\n"
		expect = "the uint8 is 0o1 and 0o2.\n"
		args = []any{uint8(1), uint8(2)}
	case 198:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal uint16 normal arguments.
`
		subject = "the uint16 is %O and %O.\n"
		expect = "the uint16 is 0o1 and 0o2.\n"
		args = []any{uint16(1), uint16(2)}
	case 199:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal uint32 normal arguments.
`
		subject = "the uint32 is %O and %O.\n"
		expect = "the uint32 is 0o1 and 0o2.\n"
		args = []any{uint32(1), uint32(2)}
	case 200:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal uint64 normal arguments.
`
		subject = "the uint64 is %O and %O.\n"
		expect = "the uint64 is 0o1 and 0o2.\n"
		args = []any{uint64(1), uint64(2)}
	case 201:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal int8 normal arguments.
`
		subject = "the int8 is %O and %O.\n"
		expect = "the int8 is 0o1 and 0o2.\n"
		args = []any{int8(1), int8(2)}
	case 202:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal int16 normal arguments.
`
		subject = "the int16 is %O and %O.\n"
		expect = "the int16 is 0o1 and 0o2.\n"
		args = []any{int16(1), int16(2)}
	case 203:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal int32 normal arguments.
`
		subject = "the int32 is %O and %O.\n"
		expect = "the int32 is 0o1 and 0o2.\n"
		args = []any{int32(1), int32(2)}
	case 204:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal int64 normal arguments.
`
		subject = "the int64 is %O and %O.\n"
		expect = "the int64 is 0o1 and 0o2.\n"
		args = []any{int64(1), int64(2)}
	case 205:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal float32 normal arguments.
`
		subject = "the float32 is %O and %O.\n"
		expect = "the float32 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{float32(2e10), float32(2e10)}
	case 206:
		desc = `
test hestiaSTRING/M64_Format can process prefixed-octal float64 normal arguments.
`
		subject = "the float64 is %O and %O.\n"
		expect = "the float64 is (NUMBER=bad) and (NUMBER=bad).\n"
		args = []any{float64(2e10), float64(2e10)}
	case 207:
		desc = `
test hestiaSTRING/M64_Format can process bool uint8 normal arguments.
`
		subject = "the uint8 is %t and %t.\n"
		expect = "the uint8 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{uint8(1), uint8(2)}
	case 208:
		desc = `
test hestiaSTRING/M64_Format can process bool uint16 normal arguments.
`
		subject = "the uint16 is %t and %t.\n"
		expect = "the uint16 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{uint16(1), uint16(2)}
	case 209:
		desc = `
test hestiaSTRING/M64_Format can process bool uint32 normal arguments.
`
		subject = "the uint32 is %t and %t.\n"
		expect = "the uint32 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{uint32(1), uint32(2)}
	case 210:
		desc = `
test hestiaSTRING/M64_Format can process bool uint64 normal arguments.
`
		subject = "the uint64 is %t and %t.\n"
		expect = "the uint64 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{uint64(1), uint64(2)}
	case 211:
		desc = `
test hestiaSTRING/M64_Format can process bool int8 normal arguments.
`
		subject = "the int8 is %t and %t.\n"
		expect = "the int8 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{int8(1), int8(2)}
	case 212:
		desc = `
test hestiaSTRING/M64_Format can process bool int16 normal arguments.
`
		subject = "the int16 is %t and %t.\n"
		expect = "the int16 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{int16(1), int16(2)}
	case 213:
		desc = `
test hestiaSTRING/M64_Format can process bool int32 normal arguments.
`
		subject = "the int32 is %t and %t.\n"
		expect = "the int32 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{int32(1), int32(2)}
	case 214:
		desc = `
test hestiaSTRING/M64_Format can process bool int64 normal arguments.
`
		subject = "the int64 is %t and %t.\n"
		expect = "the int64 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{int64(1), int64(2)}
	case 215:
		desc = `
test hestiaSTRING/M64_Format can process bool float32 normal arguments.
`
		subject = "the float32 is %t and %t.\n"
		expect = "the float32 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{float32(2e10), float32(2e10)}
	case 216:
		desc = `
test hestiaSTRING/M64_Format can process bool float64 normal arguments.
`
		subject = "the float64 is %t and %t.\n"
		expect = "the float64 is (BOOL=bad) and (BOOL=bad).\n"
		args = []any{float64(2e10), float64(2e10)}
	case 217:
		desc = `
test hestiaSTRING/M64_Format can process percent symbol alongside normal arguments.
`
		subject = "the %% is %s.\n"
		expect = "the % is percent.\n"
		args = []any{"percent"}
	case 218:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 float32 normal width+precision arguments.
`
		subject = "the float32 is %20.10e.\n"
		expect = "the float32 is 2.0000000000e+10    .\n"
		args = []any{float32(2e10)}
	case 219:
		desc = `
test hestiaSTRING/M64_Format can process ISO6093NR3 float64 normal width+precision arguments.
`
		subject = "the float64 is %20.10e.\n"
		expect = "the float64 is 2.0000000000e+10    .\n"
		args = []any{float64(2e10)}
	case 220:
		desc = `
test hestiaSTRING/M64_Format can process auto-expo float32 normal width+precision arguments.
`
		subject = "the float32 is %20.10g.\n"
		expect = "the float32 is 2.0000000000e+10    .\n"
		args = []any{float32(2e10)}
	case 221:
		desc = `
test hestiaSTRING/M64_Format can process auto-expo float64 normal width+precision arguments.
`
		subject = "the float64 is %20.10g.\n"
		expect = "the float64 is 2.0000000000e+10    .\n"
		args = []any{float64(2e10)}
	case 222:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal float32 normal width+precision arguments.
`
		subject = "the float32 is %20.10x.\n"
		expect = "the float32 is 0x2.0000000000p+10  .\n"
		args = []any{float32(2e10)}
	case 223:
		desc = `
test hestiaSTRING/M64_Format can process hexadecimal float64 normal width+precision arguments.
`
		subject = "the float64 is %20.10x.\n"
		expect = "the float64 is 0x2.0000000000p+10  .\n"
		args = []any{float64(2e10)}
	case 224:
		desc = `
test hestiaSTRING/M64_Format can process char float64 normal argument.
`
		subject = "the float64 is %c.\n"
		expect = "the float64 is (CHAR=bad).\n"
		args = []any{float64(2e10)}
	case 225:
		desc = `
test hestiaSTRING/M64_Format can process char float32 normal argument.
`
		subject = "the float32 is %c.\n"
		expect = "the float32 is (CHAR=bad).\n"
		args = []any{float32(2e10)}
	case 226:
		desc = `
test hestiaSTRING/M64_Format can process char int64 normal argument.
`
		subject = "the int64 is %c.\n"
		expect = "the int64 is (CHAR=bad).\n"
		args = []any{int64(2)}
	case 227:
		desc = `
test hestiaSTRING/M64_Format can process char int32 normal argument.
`
		subject = "the int32 is %c.\n"
		expect = "the int32 is (CHAR=bad).\n"
		args = []any{int32(2)}
	case 228:
		desc = `
test hestiaSTRING/M64_Format can process char int16 normal argument.
`
		subject = "the int16 is %c.\n"
		expect = "the int16 is (CHAR=bad).\n"
		args = []any{int16(2)}
	case 229:
		desc = `
test hestiaSTRING/M64_Format can process char int8 normal argument.
`
		subject = "the int8 is %c.\n"
		expect = "the int8 is (CHAR=bad).\n"
		args = []any{int8(2)}
	case 230:
		desc = `
test hestiaSTRING/M64_Format can process char uint64 normal argument.
`
		subject = "the uint64 is %c.\n"
		expect = "the uint64 is (CHAR=bad).\n"
		args = []any{uint64(2)}
	case 231:
		desc = `
test hestiaSTRING/M64_Format can process char uint32 normal argument.
`
		subject = "the uint32 is %c.\n"
		expect = "the uint32 is (CHAR=bad).\n"
		args = []any{uint32(2)}
	case 232:
		desc = `
test hestiaSTRING/M64_Format can process char uint16 normal argument.
`
		subject = "the uint16 is %c.\n"
		expect = "the uint16 is (CHAR=bad).\n"
		args = []any{uint16(2)}
	case 233:
		desc = `
test hestiaSTRING/M64_Format can process char uint8 normal argument.
`
		subject = "the uint8 is %c.\n"
		expect = "the uint8 is (CHAR=bad).\n"
		args = []any{uint8(2)}
	default:
		// NOTE: if you added new cases, remember to update the suite's count below.
		desc = "test hestiaSTRING/M64Format test suite error: remove this immediately!"
		subject = "???"
		expect = "unknown test case"
		args = nil
	}

	return desc, subject, expect, args
}

func Test_M64_Format(t *testing.T) {
	for i := 0; i <= 233; i++ {
		var subject, expect string
		var args []any

		s := &hestiaTESTING.Scenario{
			ID:   uint64(i),
			Name: "hestiaSTRING/M64_Format API",
		}

		// prepare
		s.Description, subject, expect, args = generate_M64_Format_params(s.ID)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: '%s'", subject))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Expect	: '%s'", expect))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Argument	: '%s'", args))

		// test
		var output string
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			if args == nil {
				output = M64_Format(subject)
			} else {
				output = M64_Format(subject, args...)
			}

			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s, hestiaTESTING.Format("Got Output	: '%v'", output))
		hestiaTESTING.Log(s, hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		if panick != "" {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, "Failed by panick!")
			t.Fail()
		}

		if output != expect {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			hestiaTESTING.Log(s, "Failed by output!")
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}
