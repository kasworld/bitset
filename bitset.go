// Copyright 2014,2015,2016,2017,2018,2019,2020,2021 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bitset

import (
	"bytes"
	"fmt"
)

type BitContainer uint8

const FullBit = 0xff
const BitLen = 8
const MaskBit = BitLen - 1
const PosShift = 3

type BitSet struct {
	size     int
	bitsList []BitContainer
}

func New(size int) *BitSet {
	va := &BitSet{}
	va.size = size
	va.bitsList = make([]BitContainer, size/BitLen+1)
	return va
}

func (va *BitSet) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "BitSet[")
	for i := 0; i < va.size; i++ {
		if va.Get(i) {
			fmt.Fprintf(&buf, "%v", i)
		}
	}
	fmt.Fprintf(&buf, "]")
	return buf.String()
}

func (va *BitSet) Dup() *BitSet {
	rtn := New(va.size)
	copy(rtn.bitsList, va.bitsList)
	return rtn
}

func (va *BitSet) calcIndexAndBit(x int) (uint, BitContainer) {
	pos := uint(x)
	indexPos := pos >> PosShift
	bitPos := pos & MaskBit
	return indexPos, 1 << bitPos
}

func (va *BitSet) Set(x int) {
	index, bit := va.calcIndexAndBit(x)
	va.bitsList[index] |= bit
}

func (va *BitSet) Clear(x int) {
	index, bit := va.calcIndexAndBit(x)
	va.bitsList[index] &^= bit
}

func (va *BitSet) Get(x int) bool {
	index, bit := va.calcIndexAndBit(x)
	return va.bitsList[index]&bit != 0
}
