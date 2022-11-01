// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 Caleb Spare <cespare@gmail.com>
// Copyright 2018 Ulf Adams <ulfjack@google.com>
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
// (1)
// The code in this file is a Go translation of the Java code written by
// Ulf Adams which may be found at
//
//     https://github.com/ulfjack/ryu/analysis/PrintFloatLookupTable.java
//     https://github.com/ulfjack/ryu/analysis/PrintDoubleLookupTable.java
//
// That source code is licensed under Apache 2.0 and this code is derivative
// work thereof.
//
// (2)
// This macro program generates hestiaGO/hestiaINTERNAL/hestiaFMT/tablesFLOAT.go
// file.

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"math/big"
)

const (
	posTableSize32               = 47
	negTableSize32               = 31
	value_POWER_5_BITS32         = 61 // max 63
	value_POWER_5_INVERSE_BITS32 = 59 // max 63

	posTableSize64               = 326
	negTableSize64               = 291 + 1
	value_POWER_5_BITS64         = 121 // max 127
	value_POWER_5_INVERSE_BITS64 = 122 // max 126
)

func main() {
	b := bytes.NewBuffer([]byte(`
// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 Caleb Spare <cespare@gmail.com>
// Copyright 2018 Ulf Adams <ulfjack@google.com>
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

// WARNING: Auto-generated codes = DO NOT EDIT!

package hestiaFMT

import (
	"hestia/hestiaINTERNAL/hestiaMATH"
)`))

	fmt.Fprintf(b, `
const (
	value_POWER_5_BITS32         = %d
	value_POWER_5_INVERSE_BITS32 = %d
	value_POWER_5_BITS64         = %d
	value_POWER_5_INVERSE_BITS64 = %d
)
`,
		value_POWER_5_BITS32,
		value_POWER_5_INVERSE_BITS32,
		value_POWER_5_BITS64,
		value_POWER_5_INVERSE_BITS64,
	)
	fmt.Fprintln(b,
		`
func __m32_5Power_Divide_By_2Power_Multiplier(q uint32) uint64 {
	switch q {`)
	for i := int64(0); i < posTableSize32; i++ {
		pow5 := big.NewInt(5)
		pow5.Exp(pow5, big.NewInt(i), nil)
		shift := pow5.BitLen() - value_POWER_5_BITS32
		rsh(pow5, shift)
		if i != 0 {
			fmt.Fprintf(b, "\n")
		}
		fmt.Fprintf(b,
			`	case %d:
		return %d`, i, pow5.Uint64())
	}
	fmt.Fprintln(b,
		`
	default:
		return 0
	}
}
`)

	fmt.Fprintln(b,
		`
func __m32_5PowerInverse_Divide_By_2Power_Multiplier(q uint32) uint64 {
	switch q {`)
	for i := int64(0); i < negTableSize32; i++ {
		pow5 := big.NewInt(5)
		pow5.Exp(pow5, big.NewInt(i), nil)
		shift := pow5.BitLen() - 1 + value_POWER_5_INVERSE_BITS32
		inv := big.NewInt(1)
		rsh(inv, -shift)
		inv.Quo(inv, pow5)
		inv.Add(inv, big.NewInt(1))
		if i != 0 {
			fmt.Fprintf(b, "\n")
		}
		fmt.Fprintf(b,
			`	case %d:
		return %d`, i, inv.Uint64())
	}
	fmt.Fprintln(b,
		`
	default:
		return 0
	}
}
`)

	mask64 := big.NewInt(1)
	mask64.Lsh(mask64, 64)
	mask64.Sub(mask64, big.NewInt(1))

	fmt.Fprintln(b,
		`
func __m64_5Power_Divide_By_2Power_Multiplier(q uint64) *hestiaMATH.UINT128 {
	switch q {`)
	for i := int64(0); i < posTableSize64; i++ {
		pow5 := big.NewInt(5)
		pow5.Exp(pow5, big.NewInt(i), nil)
		shift := pow5.BitLen() - value_POWER_5_BITS64
		rsh(pow5, shift)
		lo := new(big.Int).And(pow5, mask64)
		hi := new(big.Int).Rsh(pow5, 64)
		if i != 0 {
			fmt.Fprintf(b, "\n")
		}
		fmt.Fprintf(b,
			`	case %d:
		return &hestiaMATH.UINT128{ %d, %d }`, i, lo.Uint64(), hi.Uint64())
	}
	fmt.Fprintln(b,
		`
	default:
		return &hestiaMATH.UINT128{}
	}
}
`)

	fmt.Fprintln(b,
		`
func __m64_5PowerInverse_Divide_By_2Power_Multiplier(q uint64) *hestiaMATH.UINT128 {
	switch q {`)
	for i := int64(0); i < negTableSize64; i++ {
		pow5 := big.NewInt(5)
		pow5.Exp(pow5, big.NewInt(i), nil)
		// We want floor(log_2 5^q) here, which is pow5.BitLen() - 1.
		shift := pow5.BitLen() - 1 + value_POWER_5_INVERSE_BITS64
		inv := big.NewInt(1)
		rsh(inv, -shift)
		inv.Quo(inv, pow5)
		inv.Add(inv, big.NewInt(1))
		lo := new(big.Int).And(inv, mask64)
		hi := new(big.Int).Rsh(inv, 64)
		if i != 0 {
			fmt.Fprintf(b, "\n")
		}
		fmt.Fprintf(b,
			`	case %d:
		return &hestiaMATH.UINT128{ %d, %d }`, i, lo.Uint64(), hi.Uint64())
	}
	fmt.Fprintln(b,
		`
	default:
		return &hestiaMATH.UINT128{}
	}
}
`)

	text, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("../../tablesFLOAT.go", text, 0644); err != nil {
		log.Fatal(err)
	}
}

func rsh(x *big.Int, n int) {
	if n < 0 {
		x.Lsh(x, uint(-n))
	} else {
		x.Rsh(x, uint(n))
	}
}
