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

func Trim(source string, cutset string) string {
	return strings.Trim(source, cutset)
}

func TrimLeft(source string, cutset string) string {
	return strings.TrimLeft(source, cutset)
}

func TrimRight(source string, cutset string) string {
	return strings.TrimRight(source, cutset)
}

func TrimWhitespace(source string) string {
	return strings.TrimSpace(source)
}

func TrimWhitespaceLeft(source string) string {
	return strings.TrimLeftFunc(source, unicode.IsSpace)
}

func TrimWhitespaceRight(source string) string {
	return strings.TrimRightFunc(source, unicode.IsSpace)
}

func TrimPrefix(source, prefix string) string {
	return strings.TrimPrefix(source, prefix)
}

func TrimSuffix(source, suffix string) string {
	return strings.TrimSuffix(source, suffix)
}
