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

import (
	"unsafe"
)

func S8_FormatPOINTER[ANY any](input *ANY, base uint8, lettercase Lettercase) []rune {
	return FormatUINT8(uint8(uintptr(unsafe.Pointer(input))), base, lettercase)
}

func S16_FormatPOINTER[ANY any](input *ANY, base uint16, lettercase Lettercase) []rune {
	return FormatUINT16(uint16(uintptr(unsafe.Pointer(input))), base, lettercase)
}

func S32_FormatPOINTER[ANY any](input *ANY, base uint32, lettercase Lettercase) []rune {
	return FormatUINT32(uint32(uintptr(unsafe.Pointer(input))), base, lettercase)
}

func S64_FormatPOINTER[ANY any](input *ANY, base uint64, lettercase Lettercase) []rune {
	return FormatUINT64(uint64(uintptr(unsafe.Pointer(input))), base, lettercase)
}
