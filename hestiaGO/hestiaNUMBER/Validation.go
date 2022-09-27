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

package hestiaNUMBER

func IsUINT(input any) (ok bool) {
	if _, ok = input.(uint); ok {
		return true
	}

	if _, ok = input.(*uint); ok {
		return true
	}

	return false
}

func IsUINT8(input any) (ok bool) {
	if _, ok = input.(uint8); ok {
		return true
	}

	if _, ok = input.(*uint8); ok {
		return true
	}

	return false
}

func IsUINT16(input any) (ok bool) {
	if _, ok = input.(uint16); ok {
		return true
	}

	if _, ok = input.(*uint16); ok {
		return true
	}

	return false
}

func IsUINT32(input any) (ok bool) {
	if _, ok = input.(uint32); ok {
		return true
	}

	if _, ok = input.(*uint32); ok {
		return true
	}

	return false
}

func IsUINT64(input any) (ok bool) {
	if _, ok = input.(uint64); ok {
		return true
	}

	if _, ok = input.(*uint64); ok {
		return true
	}

	return false
}

func IsINT(input any) (ok bool) {
	if _, ok = input.(int); ok {
		return true
	}

	if _, ok = input.(*int); ok {
		return true
	}

	return false
}

func IsINT8(input any) (ok bool) {
	if _, ok = input.(int8); ok {
		return true
	}

	if _, ok = input.(*int8); ok {
		return true
	}

	return false
}

func IsINT16(input any) (ok bool) {
	if _, ok = input.(int16); ok {
		return true
	}

	if _, ok = input.(*int16); ok {
		return true
	}

	return false
}

func IsINT32(input any) (ok bool) {
	if _, ok = input.(int32); ok {
		return true
	}

	if _, ok = input.(*int32); ok {
		return true
	}

	return false
}

func IsINT64(input any) (ok bool) {
	if _, ok = input.(int64); ok {
		return true
	}

	if _, ok = input.(*int64); ok {
		return true
	}

	return false
}

func IsFLOAT32(input any) (ok bool) {
	if _, ok = input.(float32); ok {
		return true
	}

	if _, ok = input.(*float32); ok {
		return true
	}

	return false
}

func IsFLOAT64(input any) (ok bool) {
	if _, ok = input.(float64); ok {
		return true
	}

	if _, ok = input.(*float64); ok {
		return true
	}

	return false
}

func IsCOMPLEX64(input any) (ok bool) {
	if _, ok = input.(complex64); ok {
		return true
	}

	if _, ok = input.(*complex64); ok {
		return true
	}

	return false
}

func IsCOMPLEX128(input any) (ok bool) {
	if _, ok = input.(complex128); ok {
		return true
	}

	if _, ok = input.(*complex128); ok {
		return true
	}

	return false
}
