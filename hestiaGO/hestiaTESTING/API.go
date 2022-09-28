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

import (
	"testing"

	"hestia/hestiaSTRING"
)

func Register(s *Scenario, t *testing.T) {
	if s == nil {
		panic("calling hestiaTESTING.Register without providing Scenario!")
	}

	if t == nil {
		panic("calling hestiaTESTING.Register without providing testing.T!")
	}

	s.controller = t
	s.skip = s.controller.SkipNow
	s.fail = s.controller.FailNow
}

func Logf(s *Scenario, format string, args ...any) {
	if s == nil {
		panic("calling hestiaTESTING.Logf without providing Scenario!")
	}

	if format == "" {
		panic("calling hestiaTESTING.Logf without providing log string!")
	}

	if s.Log == nil {
		s.Init()
	}

	s.Log = append(s.Log, hestiaSTRING.Printf(format, args...))
}

func Logln(s *Scenario, args ...any) {
	if s == nil {
		panic("calling hestiaTESTING.Logln without providing Scenario!")
	}

	if s.Log == nil {
		s.Init()
	}

	s.Log = append(s.Log, hestiaSTRING.Println(args...))
}

func Exec(function func() any) (out any) {
	defer func() {
		err := recover()
		if err != nil {
			out = err
		}
	}()

	out = function()

	return out
}

func HasExecuted(s *Scenario) bool {
	if s == nil {
		panic("calling hestiaTESTING.HasExecuted without providing Scenario!")
	}

	return s.verdict != priv_VERDICT_UNKNOWN
}

func Conclusion(s *Scenario) Verdict {
	if s == nil {
		panic("calling hestiaTESTING.Verdict without providing Scenario!")
	}

	return s.verdict
}

func Conclude(s *Scenario, certification Verdict) {
	switch {
	case s == nil:
		panic("calling hestiaTESTING.Conclude without providing Scenario!")
	case s.skip == nil || s.fail == nil:
		panic("calling hestiaTESTING.Conclude without registering *testing.T!")
	case certification == VERDICT_PASS:
		s.verdict = VERDICT_PASS
	case certification == VERDICT_FAIL:
		s.verdict = VERDICT_FAIL
	case certification == VERDICT_SKIP:
		s.verdict = VERDICT_SKIP
	default:
		panic("calling hestiaTESTING.Conclude with unknown Verdict!")
	}
}
