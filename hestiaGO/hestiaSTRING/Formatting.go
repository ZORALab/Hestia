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
	"strings"
	"unicode"
)

type CharsMap uint8

const (
	CHARSMAP_DEFAULT CharsMap = iota
	CHARSMAP_TURKISH
	CHARSMAP_AZERI
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
