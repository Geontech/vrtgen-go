/*
 * Copyright (C) 2024 Geon Technologies, LLC
 *
 * This file is part of vrtgen-go.
 *
 * vrtgen-go is free software: you can redistribute it and/or modify it under the
 * terms of the GNU Lesser General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * vrtgen-go is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
 * FOR A PARTICULAR PURPOSE.  See the GNU Lesser General Public License for
 * more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */

package vita49

import (
	"encoding/binary"
)

type Cif7 struct {
	IndicatorField7
	CurrentValue      uint32
	AverageValue      uint32
	MedianValue       uint32
	StandardDeviation uint32
	MaxValue          uint32
	MinValue          uint32
	Precision         uint32
	Accuracy          uint32
	Velocity          uint32
	Acceleration      uint32
	ThirdDerivative   uint32
	Probability       Probability
	Belief            Belief
}

// Represents the 2nd order probability that the standard probability
// (1st order probability) is correct.
type Belief struct {
	BeliefPercent uint32
}

func (b Belief) Size() uint32 {
	return 4
}

func (b *Belief) Pack() []byte {
	// 1 word
	retval := make([]byte, b.Size())

	word1 := uint32(0)
	word1 |= uint32(b.BeliefPercent & 0xFF) // Bits 0-7
	binary.BigEndian.PutUint32(retval[0:], word1)
	return retval
}

func (b *Belief) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf[0:])
	b.BeliefPercent = uint32(word1 & 0xFF)
}

// Represents the probability that the the selected field in the same Packet
// Structure Level and array index (if appropriate) is accurate or true.
type Probability struct {
	ProbabilityFunction uint32
	ProbabilityPercent  uint32
}

func (p Probability) Size() uint32 {
	return 4
}

func (p *Probability) Pack() []byte {
	// 1 word
	retval := make([]byte, p.Size())

	word1 := uint32(0)
	word1 |= uint32(p.ProbabilityFunction&0xFF) << 8 // Bits 8-15
	word1 |= uint32(p.ProbabilityPercent & 0xFF)     // Bits 0-7
	binary.BigEndian.PutUint32(retval[0:], word1)
	return retval
}

func (p *Probability) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf[0:])
	p.ProbabilityFunction = uint32(word1>>8) & 0xFF
	p.ProbabilityPercent = uint32(word1 & 0xFF)
}
