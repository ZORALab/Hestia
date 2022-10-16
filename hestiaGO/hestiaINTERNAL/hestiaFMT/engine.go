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

type none bool

type engineMode uint

const (
	engine_PARSE_NORMAL engineMode = iota
	engine_PARSE_VERB
)

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
		return none(false)
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
