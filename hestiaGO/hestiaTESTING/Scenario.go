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

package hestiaTESTING

type Scenario struct {
	ID          uint64
	Name        string
	Description string
	Switches    map[string]bool
	Log         []string

	verdict Verdict
}

func (s *Scenario) Init() {
	if s.Switches == nil {
		s.Switches = map[string]bool{}
	}

	if s.Log == nil {
		s.Log = []string{}
	}

	s.Name = __trimWhitespace(s.Name)
	s.Description = __trimWhitespace(s.Description)
}

func (s *Scenario) Exec(function func() any) (out any) {
	return Exec(function)
}
