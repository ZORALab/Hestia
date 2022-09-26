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

import (
	"hestia/hestiaNUMBER/hestiaBITS"
)

const (
	MAX_UINT = hestiaBITS.MAX_SIZE
	MAX_INT  = int64(^hestiaBITS.MAX_SIZE)
	MIN_INT  = -MAX_INT - 1

	MAX_UINT8  = uint64(1<<8 - 1)
	MAX_UINT16 = uint64(1<<16 - 1)
	MAX_UINT32 = uint64(1<<32 - 1)
	MAX_UINT64 = uint64(1<<64 - 1)

	MAX_INT8  = int64(1<<7 - 1)
	MAX_INT16 = int64(1<<15 - 1)
	MAX_INT32 = int64(1<<31 - 1)
	MAX_INT64 = int64(1<<63 - 1)

	MIN_INT8  = int64(-1 << 7)
	MIN_INT16 = int64(-1 << 15)
	MIN_INT32 = int64(-1 << 31)
	MIN_INT64 = int64(-1 << 63)
)
