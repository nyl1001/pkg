// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fifoutils

import (
	"errors"

	"github.com/nyl1001/log"
)

var ErrEmpty error

func init() {
	ErrEmpty = errors.New("fifo is empty")
}

type FIFO struct {
	array []interface{}
	len   int
}

func NewFIFO() *FIFO {
	fifo := FIFO{array: make([]interface{}, 0), len: 0}
	return &fifo
}

func (f *FIFO) Push(ele interface{}) {
	if f.len < len(f.array) {
		f.array[f.len] = ele
	} else {
		f.array = append(f.array, ele)
	}
	f.len += 1
}

func (f *FIFO) Pop() interface{} {
	if f.len <= 0 {
		return nil
	}
	ele := f.array[0]
	f.len -= 1
	for i := 0; i < f.len; i += 1 {
		f.array[i] = f.array[i+1]
	}
	f.array[f.len] = nil
	return ele
}

func (f *FIFO) Len() int {
	return f.len
}

func (f *FIFO) ElementAt(idx int) interface{} {
	if idx >= 0 && idx < f.len {
		return f.array[idx]
	} else {
		log.Fatalf("Out of index")
		return nil
	}
}
