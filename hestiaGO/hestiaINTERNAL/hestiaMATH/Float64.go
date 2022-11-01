// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 Caleb Spare <cespare@gmail.com>
// Copyright 2018 Ulf Adams <ulfjack@google.com>
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

package hestiaMATH

import (
	"unsafe"
)

// IEEE754 64-bits float bit-definition: (1) sign (11) exponent (52) mantissa

const (
	BITS_FLOAT64_MANTISSA = 52
	BITS_FLOAT64_EXPONENT = 11

	MASK_FLOAT64_SIGN              = 1 << (BITS_FLOAT64_EXPONENT + BITS_FLOAT64_MANTISSA)
	MASK_FLOAT64_MANTISSA          = 1<<BITS_FLOAT64_MANTISSA - 1
	MASK_FLOAT64_EXPONENT          = (1<<BITS_FLOAT64_EXPONENT - 1) << BITS_FLOAT64_MANTISSA
	BIAS_FLOAT64_EXPONENT          = 1023
	MASK_FLOAT64_INFINITY_POSITIVE = MASK_FLOAT64_EXPONENT
	MASK_FLOAT64_INFINITY_NEGATIVE = MASK_FLOAT64_SIGN | MASK_FLOAT64_EXPONENT
)

type Float64 struct {
	Exponent         uint64
	Mantissa         uint64
	Base             uint64
	NegativeValue    bool
	NegativeExponent bool
}

func S64_IEEE754_Decode_FLOAT64(input float64, data *Float64) {
	if data == nil {
		panic("given data for IEEE754_Decode_FLOAT64 is nil!")
	}

	data.Base = 2
	data.Mantissa = S64_IEEE754_FloatToBits(input)
	data.NegativeValue = (data.Mantissa & MASK_FLOAT64_SIGN) != 0
	data.Exponent = (data.Mantissa & MASK_FLOAT64_EXPONENT) >> BITS_FLOAT64_MANTISSA
	data.Mantissa &= MASK_FLOAT64_MANTISSA
}

func S64_IEEE754_IsNaN_FLOAT64(input float64) bool {
	x := S64_IEEE754_FloatToBits(input)
	return x&MASK_FLOAT64_MANTISSA != 0 && x&MASK_FLOAT64_EXPONENT == MASK_FLOAT64_EXPONENT
}

func S64_IEEE754_IsINF_FLOAT64(input float64, sign int) bool {
	x := S64_IEEE754_FloatToBits(input)
	return (sign >= 0 && x == MASK_FLOAT64_INFINITY_POSITIVE) ||
		(sign <= 0 && x == MASK_FLOAT64_INFINITY_NEGATIVE)
}

func S64_IEEE754_FloatToBits(input float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&input))
}

func S64_IEEE754_BitsToFloat(input uint64) float64 {
	return *(*float64)(unsafe.Pointer(&input))
}
