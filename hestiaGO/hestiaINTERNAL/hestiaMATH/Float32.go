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

// IEEE754 32-bits float bit-definition: (1) sign (8) exponent (23) mantissa

const (
	BITS_FLOAT32_MANTISSA = 23
	BITS_FLOAT32_EXPONENT = 8

	MASK_FLOAT32_SIGN              = 1 << (BITS_FLOAT32_EXPONENT + BITS_FLOAT32_MANTISSA)
	MASK_FLOAT32_MANTISSA          = 1<<BITS_FLOAT32_MANTISSA - 1
	MASK_FLOAT32_EXPONENT          = (1<<BITS_FLOAT32_EXPONENT - 1) << BITS_FLOAT32_MANTISSA
	BIAS_FLOAT32_EXPONENT          = 127
	MASK_FLOAT32_INFINITY_POSITIVE = MASK_FLOAT32_EXPONENT
	MASK_FLOAT32_INFINITY_NEGATIVE = MASK_FLOAT32_SIGN | MASK_FLOAT32_EXPONENT
)

type Float32 struct {
	Exponent         uint32
	Mantissa         uint32
	Base             uint32
	NegativeValue    bool
	NegativeExponent bool
}

func S32_IEEE754_Decode_FLOAT32(input float32, data *Float32) {
	if data == nil {
		panic("given data for IEEE754_Decode_FLOAT32 is nil!")
	}

	data.Base = 2
	data.Mantissa = S32_IEEE754_FloatToBits(input)
	data.NegativeValue = (data.Mantissa & MASK_FLOAT32_SIGN) != 0
	data.Exponent = (data.Mantissa & MASK_FLOAT32_EXPONENT) >> BITS_FLOAT32_MANTISSA
	data.Mantissa &= MASK_FLOAT32_MANTISSA
}

func S32_IEEE754_IsNaN_FLOAT32(input float32) bool {
	x := S32_IEEE754_FloatToBits(input)
	return x&MASK_FLOAT32_MANTISSA != 0 && x&MASK_FLOAT32_EXPONENT == MASK_FLOAT32_EXPONENT
}

func S32_IEEE754_IsINF_FLOAT32(input float32, sign int) bool {
	x := S32_IEEE754_FloatToBits(input)
	return (sign >= 0 && x == MASK_FLOAT32_INFINITY_POSITIVE) ||
		(sign <= 0 && x == MASK_FLOAT32_INFINITY_NEGATIVE)
}

func S32_IEEE754_FloatToBits(input float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&input))
}

func S32_IEEE754_BitsToFloat(input uint32) float32 {
	return *(*float32)(unsafe.Pointer(&input))
}

func S32_Floor_FLOAT32(input float32) (round float32) {
	return _s32_Floor_FLOAT32(input)
}

func S32_Modf_FLOAT32(input float32) (round float32, fraction float32) {
	return _s32_Modf_FLOAT32(input)
}
