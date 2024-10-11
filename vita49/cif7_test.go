package vita49

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeliefBits(t *testing.T) {
	p := Belief{}
	assert.Equal(t, BeliefBits, p.Bits())
}

func TestBeliefDefault(t *testing.T) {
	b := Belief{}
	assert.Equal(t, uint32(0), b.BeliefPercent)
	// Pack
	packed := b.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	b.Unpack(packed)
	assert.Equal(t, uint32(0), b.BeliefPercent)
}

func TestBelief(t *testing.T) {
	cases := []struct {
		name          string
		BeliefPercent uint32
		expected      []byte
	}{
		{
			name:          "Rule 9.12-12",
			BeliefPercent: 50, // 50% confidence
			expected:      []byte{0, 0, 0, 0x32},
		},
		{
			name:          "Rule 9.12-13",
			BeliefPercent: 255, // 100% confidence
			expected:      []byte{0, 0, 0, 0xFF},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := Belief{}
			b.BeliefPercent = tc.BeliefPercent

			assert.Equal(t, tc.BeliefPercent, b.BeliefPercent)

			// Pack
			packed := b.Pack()

			assert.Equal(t, tc.expected, packed)
			// Unpack
			b.Unpack(packed)

			assert.Equal(t, tc.BeliefPercent, b.BeliefPercent)
		})
	}
}

func TestProbabilityBits(t *testing.T) {
	p := Probability{}
	assert.Equal(t, BeliefBits, p.Bits())
}

func TestProbabilityDefault(t *testing.T) {
	p := Probability{}
	assert.Equal(t, uint32(0), p.ProbabilityFunction)
	assert.Equal(t, uint32(0), p.ProbabilityPercent)
	// Pack
	packed := p.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	p.Unpack(packed)
	assert.Equal(t, uint32(0), p.ProbabilityFunction)
	assert.Equal(t, uint32(0), p.ProbabilityPercent)
}

func TestProbability(t *testing.T) {
	cases := []struct {
		name                string
		ProbabilityFunction uint32
		ProbabilityPercent  uint32
		expected            []byte
	}{
		{
			name:                "Rule 9.12-7/1",
			ProbabilityFunction: 1,   // Normal Distribution
			ProbabilityPercent:  255, // 100% confidence
			expected:            []byte{0, 0, 0x1, 0xFF},
		},
		{
			name:                "Rule 9.12-7/2",
			ProbabilityFunction: 0,   // Uniform Distribution
			ProbabilityPercent:  255, // 100% confidence
			expected:            []byte{0, 0, 0, 0xFF},
		},
		{
			name:                "Rule 9.12-9",
			ProbabilityFunction: 0,  // Uniformed Distribution
			ProbabilityPercent:  50, // 50% confidence
			expected:            []byte{0, 0, 0, 0x32},
		},
		{
			name:                "Rule 9.12-10",
			ProbabilityFunction: 1,  // Normal Distribution
			ProbabilityPercent:  50, // 50% confidence
			expected:            []byte{0, 0, 0x01, 0x32},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := Probability{}
			p.ProbabilityFunction = tc.ProbabilityFunction
			p.ProbabilityPercent = tc.ProbabilityPercent

			assert.Equal(t, tc.ProbabilityFunction, p.ProbabilityFunction)
			assert.Equal(t, tc.ProbabilityPercent, p.ProbabilityPercent)

			// Pack
			packed := p.Pack()

			assert.Equal(t, tc.expected, packed)
			// Unpack
			p.Unpack(packed)
			assert.Equal(t, tc.ProbabilityFunction, p.ProbabilityFunction)
			assert.Equal(t, tc.ProbabilityPercent, p.ProbabilityPercent)
		})
	}
}
