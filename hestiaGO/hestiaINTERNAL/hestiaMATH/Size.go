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

package hestiaMATH

func S64_Resize(input *uint64, size uint16, signValue bool) Error {
	var mask uint64

	// mask to register
	switch {
	case input == nil:
		return ERROR_INPUT_INVALID
	case size > 64:
		return ERROR_INPUT_OUT_OF_RANGE
	case size == 0:
		*input = 0
		return ERROR_OK
	case size <= 8:
		if signValue {
			mask = MAX_INT8
		} else {
			mask = MAX_UINT8
		}
	case size <= 16:
		if signValue {
			mask = MAX_INT16
		} else {
			mask = MAX_UINT16
		}
	case size <= 32:
		if signValue {
			mask = MAX_INT32
		} else {
			mask = MAX_UINT32
		}
	default:
		if signValue {
			mask = MAX_INT64
		} else {
			mask = MAX_UINT64
		}
	}

	*input &= mask
	switch size {
	case 8, 16, 32, 64:
		return ERROR_OK
	default:
	}

	// mask to precision
	mask = (1 << size) - 1
	*input &= mask

	return ERROR_OK
}

func S32_Resize(input *uint32, size uint16, signValue bool) Error {
	var mask uint32

	// mask to register
	switch {
	case input == nil:
		return ERROR_INPUT_INVALID
	case size > 32:
		return ERROR_INPUT_OUT_OF_RANGE
	case size == 0:
		*input = 0
		return ERROR_OK
	case size <= 8:
		if signValue {
			mask = MAX_INT8
		} else {
			mask = MAX_UINT8
		}
	case size <= 16:
		if signValue {
			mask = MAX_INT16
		} else {
			mask = MAX_UINT16
		}
	default:
		if signValue {
			mask = MAX_INT32
		} else {
			mask = MAX_UINT32
		}
	}

	*input &= mask
	switch size {
	case 8, 16, 32:
		return ERROR_OK
	default:
	}

	// mask to precision
	mask = (1 << size) - 1
	*input &= mask

	return ERROR_OK
}

func S16_Resize(input *uint16, size uint16, signValue bool) Error {
	var mask uint16

	// mask to register
	switch {
	case input == nil:
		return ERROR_INPUT_INVALID
	case size > 32:
		return ERROR_INPUT_OUT_OF_RANGE
	case size == 0:
		*input = 0
		return ERROR_OK
	case size <= 8:
		if signValue {
			mask = MAX_INT8
		} else {
			mask = MAX_UINT8
		}
	default:
		if signValue {
			mask = MAX_INT16
		} else {
			mask = MAX_UINT16
		}
	}

	*input &= mask
	switch size {
	case 8, 16:
		return ERROR_OK
	default:
	}

	// mask to precision
	mask = (1 << size) - 1
	*input &= mask

	return ERROR_OK
}

func S8_Resize(input *uint8, size uint16, signValue bool) Error {
	var mask uint8

	// mask to register
	switch {
	case input == nil:
		return ERROR_INPUT_INVALID
	case size > 32:
		return ERROR_INPUT_OUT_OF_RANGE
	case size == 0:
		*input = 0
		return ERROR_OK
	default:
		if signValue {
			mask = MAX_INT8
		} else {
			mask = MAX_UINT8
		}
	}

	*input &= mask
	if size == 8 {
		return ERROR_OK
	}

	// mask to precision
	mask = (1 << size) - 1
	*input &= mask

	return ERROR_OK
}
