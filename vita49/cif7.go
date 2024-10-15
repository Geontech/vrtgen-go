package vita49

import (
	"encoding/binary"
)

const (
	beliefBytes      = uint32(4)
	probabilityBytes = uint32(4)
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
	return beliefBytes
}

func (b *Belief) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

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
	return probabilityBytes
}

func (p *Probability) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

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
