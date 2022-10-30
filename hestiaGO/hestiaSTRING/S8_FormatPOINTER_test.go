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
	"hestia/hestiaERROR"
	"hestia/hestiaTESTING"
	"testing"
)

func Test_S8_FormatPOINTER_UINT8(t *testing.T) {
	// prepare
	s := &hestiaTESTING.Scenario{
		Name: "Test S8_FormatPOINTER API",
		Description: `
S8_FormatPOINTER is capable of process a given pointer of uint8 data type.`,
	}

	subject := uint8(sample_UINT64)
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Subject	: '%p'", &subject))

	// test
	var output string
	err := hestiaERROR.OK
	panick := ""

	_panick := hestiaTESTING.Exec(func() any {
		output, err = S8_FormatPOINTER(&subject, 16, LETTERCASE_LOWER)
		return ""
	})
	panick, _ = _panick.(string)

	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Output	: '%#v'", output))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Error	: '%#v'", err))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Panick	: '%#v'", panick))

	// assert
	hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
	if output == "" {
		t.Logf("output is empty!")
		t.Fail()
	}

	if err != hestiaERROR.OK {
		t.Logf("err is not OK!")
		t.Fail()
	}

	if panick != "" {
		t.Logf("was panicked!")
		t.Fail()
	}

	// report
	t.Logf("%v", hestiaTESTING.ToString(s))
}

func Test_S8_FormatPOINTER_INT8(t *testing.T) {
	// prepare
	s := &hestiaTESTING.Scenario{
		Name: "Test S8_FormatPOINTER API",
		Description: `
S8_FormatPOINTER is capable of process a given pointer of int16 data type.`,
	}

	subject := int8(sample_UINT64)
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Subject	: '%p'", &subject))

	// test
	var output string
	err := hestiaERROR.OK
	panick := ""

	_panick := hestiaTESTING.Exec(func() any {
		output, err = S8_FormatPOINTER(&subject, 16, LETTERCASE_LOWER)
		return ""
	})
	panick, _ = _panick.(string)

	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Output	: '%#v'", output))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Error	: '%#v'", err))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Panick	: '%#v'", panick))

	// assert
	hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
	if output == "" {
		t.Logf("output is empty!")
		t.Fail()
	}

	if err != hestiaERROR.OK {
		t.Logf("err is not OK!")
		t.Fail()
	}

	if panick != "" {
		t.Logf("was panicked!")
		t.Fail()
	}

	// report
	t.Logf("%v", hestiaTESTING.ToString(s))
}

func Test_S8_FormatPOINTER_BOOL(t *testing.T) {
	// prepare
	s := &hestiaTESTING.Scenario{
		Name: "Test S8_FormatPOINTER API",
		Description: `
S8_FormatPOINTER is capable of process a given pointer of bool data type.`,
	}

	subject := true
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Subject	: '%p'", &subject))

	// test
	var output string
	err := hestiaERROR.OK
	panick := ""

	_panick := hestiaTESTING.Exec(func() any {
		output, err = S8_FormatPOINTER(&subject, 16, LETTERCASE_LOWER)
		return ""
	})
	panick, _ = _panick.(string)

	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Output	: '%#v'", output))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Error	: '%#v'", err))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Panick	: '%#v'", panick))

	// assert
	hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
	if output == "" {
		t.Logf("output is empty!")
		t.Fail()
	}

	if err != hestiaERROR.OK {
		t.Logf("err is not OK!")
		t.Fail()
	}

	if panick != "" {
		t.Logf("was panicked!")
		t.Fail()
	}

	// report
	t.Logf("%v", hestiaTESTING.ToString(s))
}

func Test_S8_FormatPOINTER_STRING(t *testing.T) {
	// prepare
	s := &hestiaTESTING.Scenario{
		Name: "Test S8_FormatPOINTER API",
		Description: `
S8_FormatPOINTER is capable of process a given pointer of string data type.`,
	}

	subject := sample_STRING
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Subject	: '%p'", &subject))

	// test
	var output string
	err := hestiaERROR.OK
	panick := ""

	_panick := hestiaTESTING.Exec(func() any {
		output, err = S8_FormatPOINTER(&subject, 16, LETTERCASE_LOWER)
		return ""
	})
	panick, _ = _panick.(string)

	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Output	: '%#v'", output))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Error	: '%#v'", err))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Panick	: '%#v'", panick))

	// assert
	hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
	if output == "" {
		t.Logf("output is empty!")
		t.Fail()
	}

	if err != hestiaERROR.OK {
		t.Logf("err is not OK!")
		t.Fail()
	}

	if panick != "" {
		t.Logf("was panicked!")
		t.Fail()
	}

	// report
	t.Logf("%v", hestiaTESTING.ToString(s))
}

func Test_S8_FormatPOINTER_POINTER(t *testing.T) {
	// prepare
	s := &hestiaTESTING.Scenario{
		Name: "Test S8_FormatPOINTER API",
		Description: `
S8_FormatPOINTER is capable of process a given pointer of pointer data type.`,
	}

	x := sample_STRING
	subject := &x
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Subject	: '%p'", &subject))

	// test
	var output string
	err := hestiaERROR.OK
	panick := ""

	_panick := hestiaTESTING.Exec(func() any {
		output, err = S8_FormatPOINTER(&subject, 16, LETTERCASE_LOWER)
		return ""
	})
	panick, _ = _panick.(string)

	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Output	: '%#v'", output))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Error	: '%#v'", err))
	hestiaTESTING.Log(s, hestiaTESTING.Format("Given Panick	: '%#v'", panick))

	// assert
	hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
	if output == "" {
		t.Logf("output is empty!")
		t.Fail()
	}

	if err != hestiaERROR.OK {
		t.Logf("err is not OK!")
		t.Fail()
	}

	if panick != "" {
		t.Logf("was panicked!")
		t.Fail()
	}

	// report
	t.Logf("%v", hestiaTESTING.ToString(s))
}
