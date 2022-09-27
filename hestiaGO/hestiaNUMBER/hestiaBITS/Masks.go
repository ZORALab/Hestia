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

package hestiaBITS

const (
	MASK_UINT   = ^uint(0)
	MASK_UINT8  = 1<<8 - 1
	MASK_UINT16 = 1<<16 - 1
	MASK_UINT32 = 1<<32 - 1
	MASK_UINT64 = 1<<64 - 1

	MASK_INT8  = 1<<7 - 1
	MASK_INT16 = 1<<15 - 1
	MASK_INT32 = 1<<31 - 1
	MASK_INT64 = 1<<63 - 1
)
