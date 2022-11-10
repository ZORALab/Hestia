// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2010 The Go Authors <https://cs.opensource.google/go/go>
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

#include "textflag.h"

#define Big		0x4330000000000000 // 2**52

// func _s64_Floor_FLOAT64(input float64) (round float64)
TEXT ·_s64_Floor_FLOAT64(SB),NOSPLIT,$0
	FMOVD   x+0(FP), F0  // F0=x
	FSTCW   -2(SP)       // save old Control Word
	MOVW    -2(SP), AX
	ANDW    $0xf3ff, AX
	ORW     $0x0400, AX  // Rounding Control set to -Inf
	MOVW    AX, -4(SP)   // store new Control Word
	FLDCW   -4(SP)       // load new Control Word
	FRNDINT              // F0=Floor(x)
	FLDCW   -2(SP)       // load old Control Word
	FMOVDP  F0, ret+8(FP)
	RET
