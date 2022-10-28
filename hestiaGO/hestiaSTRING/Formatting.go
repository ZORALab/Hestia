// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2009 The Go Authors <https://cs.opensource.google/go/go>
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
	"hestia/hestiaINTERNAL/hestiaFMT"
	"strings"
	"unicode"
)

type CharsMap uint8

const (
	CHARSMAP_DEFAULT CharsMap = iota
	CHARSMAP_TURKISH
	CHARSMAP_AZERI
)

const (
	LETTERCASE_LOWER = hestiaFMT.LETTERCASE_LOWER
	LETTERCASE_UPPER = hestiaFMT.LETTERCASE_UPPER
)

const (
	NOTATION_SCIENTIFIC_AUTO_EXPONENT = hestiaFMT.NOTATION_SCIENTIFIC_AUTO_EXPONENT
	NOTATION_SCIENTIFIC               = hestiaFMT.NOTATION_SCIENTIFIC
	NOTATION_DECIMAL_NO_EXPONENT      = hestiaFMT.NOTATION_DECIMAL_NO_EXPONENT
	NOTATION_IEEE754                  = hestiaFMT.NOTATION_IEEE754
)

func ToUppercase(source string, charmap CharsMap) string {
	switch charmap {
	case CHARSMAP_TURKISH:
		return strings.ToUpperSpecial(unicode.TurkishCase, source)
	case CHARSMAP_AZERI:
		return strings.ToUpperSpecial(unicode.TurkishCase, source)
	default:
		return strings.ToUpper(source)
	}
}

func ToLowercase(source string, charmap CharsMap) string {
	switch charmap {
	case CHARSMAP_TURKISH:
		return strings.ToLowerSpecial(unicode.TurkishCase, source)
	case CHARSMAP_AZERI:
		return strings.ToLowerSpecial(unicode.TurkishCase, source)
	default:
		return strings.ToLower(source)
	}
}

func ToTitlecase(source string, charmap CharsMap) string {
	switch charmap {
	case CHARSMAP_TURKISH:
		return strings.ToTitleSpecial(unicode.TurkishCase, source)
	case CHARSMAP_AZERI:
		return strings.ToTitleSpecial(unicode.TurkishCase, source)
	default:
		return strings.ToTitle(source)
	}
}

func S8_FormatUINT8(input uint8, base uint8,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatUINT8(input, base, lettercase)), hestiaERROR.OK
}

func S8_FormatINT8(input int8, base uint8,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatINT8(input, base, lettercase)), hestiaERROR.OK
}

func S8_Itoa(input int8) (output string) {
	output, _ = S8_FormatINT8(input, 10, LETTERCASE_LOWER)
	return output
}

func S16_FormatUINT16(input uint16, base uint16,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatUINT16(input, base, lettercase)), hestiaERROR.OK
}

func S16_FormatINT16(input int16, base uint16,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatINT16(input, base, lettercase)), hestiaERROR.OK
}

func S16_Itoa(input int16) (output string) {
	output, _ = S16_FormatINT16(input, 10, LETTERCASE_LOWER)
	return output
}

func S32_FormatUINT32(input uint32, base uint32,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatUINT32(input, base, lettercase)), hestiaERROR.OK
}

func S32_FormatINT32(input int32, base uint32,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatINT32(input, base, lettercase)), hestiaERROR.OK
}

func S32_Itoa(input int32) (output string) {
	output, _ = S32_FormatINT32(input, 10, LETTERCASE_LOWER)
	return output
}

func S64_FormatUINT64(input uint64, base uint64,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatUINT64(input, base, lettercase)), hestiaERROR.OK
}

func S64_FormatINT64(input int64, base uint64,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatINT64(input, base, lettercase)), hestiaERROR.OK
}

func S64_Itoa(input int64) (output string) {
	output, _ = S64_FormatINT64(input, 10, LETTERCASE_LOWER)
	return output
}

func M32_FormatFLOAT32(input float32, base uint32, precision uint32,
	lettercase hestiaFMT.Lettercase,
	notation hestiaFMT.Notation) (out string, err hestiaERROR.Error) {
	// Step 1: process all parameters to prevent panic
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	_processLettercase(&lettercase)

	switch {
	case _processNotation(&notation) != hestiaERROR.OK:
		fallthrough
	case base != 2 && notation == NOTATION_IEEE754:
		return "", hestiaERROR.BAD_DESCRIPTOR
	}

	// Step 2: constructing the data structure for processing
	param := hestiaFMT.ParamsFLOAT32{
		Value:      input,
		Precision:  precision,
		Base:       base,
		Lettercase: lettercase,
		Notation:   notation,
	}

	// Step 3: perform formatting
	return string(hestiaFMT.M32_FormatFLOAT32(&param)), hestiaERROR.OK
}

func FormatBOOL(input bool, lettercase hestiaFMT.Lettercase) string {
	_processLettercase(&lettercase)
	return string(hestiaFMT.FormatBOOL(input, lettercase))
}

func S_ParseUINT(input string, base, size uint16) (out uint64, err hestiaERROR.Error) {
	return s_ParseUINT(input, uint64(base), size)
}

func S_ParseINT(input string, base, size uint16) (out int64, err hestiaERROR.Error) {
	return s_ParseINT(input, uint64(base), size)
}

func S_Atoi(input string, size uint16) (out int64, err hestiaERROR.Error) {
	return s_ParseINT(input, 10, size)
}

func ParseBOOL(input string) (bool, hestiaERROR.Error) {
	switch ToUppercase(input, CHARSMAP_DEFAULT) {
	case "1", "T", "TRUE":
		return true, hestiaERROR.OK
	case "0", "F", "FALSE":
		return false, hestiaERROR.OK
	default:
		return false, hestiaERROR.INVALID_ARGUMENT
	}
}

func _processLettercase(lettercase *hestiaFMT.Lettercase) {
	switch *lettercase {
	case LETTERCASE_UPPER, LETTERCASE_LOWER:
	default:
		*lettercase = LETTERCASE_LOWER
	}
}

func _processNotation(notation *hestiaFMT.Notation) hestiaERROR.Error {
	switch *notation {
	case NOTATION_SCIENTIFIC,
		NOTATION_SCIENTIFIC_AUTO_EXPONENT,
		NOTATION_DECIMAL_NO_EXPONENT,
		NOTATION_IEEE754:
		return hestiaERROR.OK
	default:
		return hestiaERROR.DATA_INVALID
	}
}
