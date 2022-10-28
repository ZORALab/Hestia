// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2019 Caleb Spare <cespare@gmail.com>
// Copyright 2018 Ulf Adams <ulfjack@google.com>
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

func S32_Multiply_By_Bits(x, y uint32) (high, low uint32) {
	tmp := uint64(x) * uint64(y)
	high, low = uint32(tmp>>32), uint32(tmp)
	return high, low
}

func S64_Multiply_By_Bits(x, y uint64) (high, low uint64) {
	const mask32 = uint64(^uint32(0))

	x0 := x & mask32
	x1 := x >> 32

	y0 := y & mask32
	y1 := y >> 32

	w0 := x0 * y0
	t := x1*y0 + w0>>32

	w1 := (t & mask32)
	w1 += x0 * y1
	w2 := t >> 32

	high = x1*y1 + w2 + w1>>32
	low = x * y

	return high, low
}
