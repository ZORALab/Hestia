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

package hestiaIO

import (
	"hestia/hestiaERROR"
)

type Whence uint8

const (
	FROM_START Whence = iota
	FROM_CURRENT
	FROM_END
)

type Init interface {
	Init() (err hestiaERROR.Error)
}

type String interface {
	String() (err hestiaERROR.Error)
}

type Destroy interface {
	Destroy() (err hestiaERROR.Error)
}

type Reader interface {
	Read(target *[]byte) (err hestiaERROR.Error)
}

type Writer interface {
	Write(source *[]byte) (err hestiaERROR.Error)
}

type Seeker interface {
	Seek(offset int64, whence Whence) (err hestiaERROR.Error)
}

type Closer interface {
	Close() (err hestiaERROR.Error)
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadCloser interface {
	Reader
	Closer
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type ReadSeeker interface {
	Reader
	Seeker
}

type WriteSeeker interface {
	Writer
	Seeker
}

type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}
