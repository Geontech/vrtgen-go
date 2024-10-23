package vita49

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolarizationBytes(t *testing.T) {
	p := Polarization{}
	assert.Equal(t, polarizationBytes, p.Size())
}

func TestPolarizationDefault(t *testing.T) {
	p := Polarization{}
	assert.Equal(t, float64(0), p.TiltAngle)
	assert.Equal(t, float64(0), p.EllipticityAngle)
	// Pack
	packed := p.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	p.Unpack(packed)
	assert.Equal(t, float64(0), p.TiltAngle)
	assert.Equal(t, float64(0), p.EllipticityAngle)
}

func TestPolarization(t *testing.T) {
	cases := []struct {
		name             string
		tiltAngle        float64
		ellipticityAngle float64
		expected         []byte
	}{
		{
			name:             "Rule 9.4.8-1",
			tiltAngle:        1,
			ellipticityAngle: 1,
			expected:         []byte{0x20, 0, 0x20, 0},
		},
		{
			name:             "Rule 9.4.8-2/1",
			tiltAngle:        1,
			ellipticityAngle: 0,
			expected:         []byte{0x20, 0, 0, 0},
		},
		{
			name:             "Rule 9.4.8-2/2",
			tiltAngle:        -1,
			ellipticityAngle: 0,
			expected:         []byte{0xE0, 0, 0, 0},
		},
		{
			name:             "Rule 9.4.8-2/3",
			tiltAngle:        1. / (1 << 13), //radix=13
			ellipticityAngle: 0,
			expected:         []byte{0, 0x01, 0, 0},
		},
		{
			name:             "Rule 9.4.8-2/4",
			tiltAngle:        -1. / (1 << 13), //radix=13
			ellipticityAngle: 0,
			expected:         []byte{0xFF, 0xFF, 0, 0},
		},
		{
			name:             "Rule 9.4.8-3/1",
			tiltAngle:        0,
			ellipticityAngle: 1,
			expected:         []byte{0, 0, 0x20, 0},
		},
		{
			name:             "Rule 9.4.8-3/2",
			tiltAngle:        0,
			ellipticityAngle: -1,
			expected:         []byte{0, 0, 0xE0, 0},
		},
		{
			name:             "Rule 9.4.8-3/3",
			tiltAngle:        0,
			ellipticityAngle: 1. / (1 << 13), //radix=13
			expected:         []byte{0, 0, 0, 0x01},
		},
		{
			name:             "Rule 9.4.8-3/4",
			tiltAngle:        0,
			ellipticityAngle: -1. / (1 << 13), //radix=13
			expected:         []byte{0, 0, 0xFF, 0xFF},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := Polarization{}
			p.TiltAngle = tc.tiltAngle
			p.EllipticityAngle = tc.ellipticityAngle

			assert.Equal(t, tc.tiltAngle, p.TiltAngle)
			assert.Equal(t, tc.ellipticityAngle, p.EllipticityAngle)

			// Pack
			packed := p.Pack()

			assert.Equal(t, tc.expected, packed)
			// Unpack
			p.Unpack(packed)

			assert.Equal(t, tc.tiltAngle, p.TiltAngle)
			assert.Equal(t, tc.ellipticityAngle, p.EllipticityAngle)
		})
	}
}

func TestPointingVectorBytes(t *testing.T) {
	p := PointingVector{}
	assert.Equal(t, pointingVectorBytes, p.Size())
}

func TestPointingVectorDefault(t *testing.T) {
	p := PointingVector{}
	assert.Equal(t, float64(0), p.Elevation)
	assert.Equal(t, float64(0), p.Azimuthal)
	// Pack
	packed := p.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	p.Unpack(packed)
	assert.Equal(t, float64(0), p.Elevation)
	assert.Equal(t, float64(0), p.Azimuthal)
}

func TestPointingVector(t *testing.T) {
	cases := []struct {
		name      string
		elevation float64
		azimuthal float64
		expected  []byte
	}{
		{
			name:      "Rule 9.4.1.1-1",
			elevation: 0,
			azimuthal: 0,
			expected:  []byte{0, 0, 0, 0},
		},
		{
			name:      "Rule 9.4.1.1-2/1",
			elevation: 0,
			azimuthal: 1,
			expected:  []byte{0, 0, 0, 0x80},
		},
		// {
		// 	name:      "Rule 9.4.1.1-2/2",
		// 	elevation: 0,
		// 	azimuthal: 511.9921875,
		// 	expected:  []byte{0, 0, 0xFF, 0xFF},
		// },
		{
			name:      "Rule 9.4.1.1-2/3",
			elevation: 0,
			azimuthal: .0078125,
			expected:  []byte{0, 0, 0, 0x01},
		},
		{
			name:      "Rule 9.4.1.1-3/1",
			elevation: 1,
			azimuthal: 0,
			expected:  []byte{0, 0x80, 0, 0},
		},
		{
			name:      "Rule 9.4.1.1-3/2",
			elevation: -1,
			azimuthal: 0,
			expected:  []byte{0xFF, 0x80, 0, 0},
		},
		{
			name:      "Rule 9.4.1.1-3/3",
			elevation: .0078125,
			azimuthal: 0,
			expected:  []byte{0, 0x01, 0, 0},
		},
		{
			name:      "Rule 9.4.1.1-3/4",
			elevation: -.0078125,
			azimuthal: 0,
			expected:  []byte{0xFF, 0xFF, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := PointingVector{}
			p.Elevation = tc.elevation
			p.Azimuthal = tc.azimuthal

			assert.Equal(t, tc.elevation, p.Elevation)
			assert.Equal(t, tc.azimuthal, p.Azimuthal)

			// Pack
			packed := p.Pack()

			assert.Equal(t, tc.expected, packed)
			// Unpack
			p.Unpack(packed)
			assert.Equal(t, tc.elevation, p.Elevation)
			assert.Equal(t, tc.azimuthal, p.Azimuthal)
		})
	}
}

func TestSpatialReferenceTypeBytes(t *testing.T) {
	s := SpatialReferenceType{}
	assert.Equal(t, spatialReferenceTypeBytes, s.Size())
}

func TestSpatialReferenceTypeDefault(t *testing.T) {
	s := SpatialReferenceType{}
	assert.Equal(t, uint16(0), s.SpatialIdentifier)
	assert.Equal(t, uint8(0), s.DefinedReference)
	assert.Equal(t, uint8(0), s.BeamType)
	// Pack
	packed := s.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	s.Unpack(packed)
	assert.Equal(t, uint16(0), s.SpatialIdentifier)
	assert.Equal(t, uint8(0), s.DefinedReference)
	assert.Equal(t, uint8(0), s.BeamType)
}

func TestSpatialReferenceType(t *testing.T) {
	cases := []struct {
		name              string
		SpatialIdentifier uint16
		DefinedReference  uint8
		BeamType          uint8
		expected          []byte
		msg               string
	}{
		{
			name:              "Rule 9.4.12-1",
			SpatialIdentifier: 0,
			DefinedReference:  0,
			BeamType:          0,
			expected:          []byte{0, 0, 0, 0},
			msg:               "Rule does not apply in library",
		},
		{
			name:              "Rule 9.4.12-2",
			SpatialIdentifier: 0,
			DefinedReference:  0,
			BeamType:          0,
			expected:          []byte{0, 0, 0, 0},
			msg:               "Rule does not apply in library",
		},
		{
			name:              "Rule 9.4.12-3",
			SpatialIdentifier: 0,
			DefinedReference:  1, // ECEF
			BeamType:          0, // Not specified
			expected:          []byte{0, 0, 0, 0x04},
		},
		{
			name:              "Rule 9.4.12-4/1",
			SpatialIdentifier: 0,
			DefinedReference:  0, // Not specified
			BeamType:          2, // Null
			expected:          []byte{0, 0, 0, 0x02},
		},
		{
			name:              "Rule 9.4.12-4/2",
			SpatialIdentifier: 0,
			DefinedReference:  1, // ECEF
			BeamType:          1, // Beam, or signal direction
			expected:          []byte{0, 0, 0, 0x05},
		},
		{
			name:              "Rule 9.4.12-4/3",
			SpatialIdentifier: 0,
			DefinedReference:  2, // Platform centered
			BeamType:          1, //Beam, or signal direction
			expected:          []byte{0, 0, 0, 0x09},
		},
		{
			name:              "Rule 9.4.12-4/4",
			SpatialIdentifier: 0,
			DefinedReference:  3, // Array centered
			BeamType:          1, // Beam, or signal direction
			expected:          []byte{0, 0, 0, 0x0D},
		},
		{
			name:              "Rule 9.4.12-4/5",
			SpatialIdentifier: 0,
			DefinedReference:  3, // Array centered
			BeamType:          3, // Reserved
			expected:          []byte{0, 0, 0, 0x0F},
		},
		{
			name:              "Rule 9.4.12-4/6",
			SpatialIdentifier: 0,
			DefinedReference:  1, // ECEF
			BeamType:          1, // Beam, or signal direction
			expected:          []byte{0, 0, 0, 0x05},
		},
		{
			name:              "Rule 9.4.12-4/7",
			SpatialIdentifier: 0,
			DefinedReference:  1, // ECEF
			BeamType:          2, // Null
			expected:          []byte{0, 0, 0, 0x06},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			s := SpatialReferenceType{}
			s.SpatialIdentifier = tc.SpatialIdentifier
			s.DefinedReference = tc.DefinedReference
			s.BeamType = tc.BeamType

			assert.Equal(t, tc.SpatialIdentifier, s.SpatialIdentifier)
			assert.Equal(t, tc.DefinedReference, s.DefinedReference)
			assert.Equal(t, tc.BeamType, s.BeamType)

			// Pack
			packed := s.Pack()

			assert.Equal(t, tc.expected, packed)
			// Unpack
			s.Unpack(packed)

			assert.Equal(t, tc.SpatialIdentifier, s.SpatialIdentifier)
			assert.Equal(t, tc.DefinedReference, s.DefinedReference)
			assert.Equal(t, tc.BeamType, s.BeamType)
		})
	}
}

func TestBeamWidthBytes(t *testing.T) {
	b := BeamWidth{}
	assert.Equal(t, beamWidthBytes, b.Size())
}

func TestBeamWidthDefault(t *testing.T) {
	b := BeamWidth{}
	assert.Equal(t, float64(0), b.Horizontal)
	assert.Equal(t, float64(0), b.Vertical)
	// Pack
	packed := b.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	b.Unpack(packed)
	assert.Equal(t, float64(0), b.Horizontal)
	assert.Equal(t, float64(0), b.Vertical)
}

func TestBeamWidth(t *testing.T) {
	cases := []struct {
		name       string
		Horizontal float64
		Vertical   float64
		expected   []byte
	}{
		{
			name:       "Rule 9.4.2-1",
			Horizontal: 0,
			Vertical:   0,
			expected:   []byte{0, 0, 0, 0},
		},
		{
			name:       "Rule 9.4.2-2/1",
			Horizontal: 1,
			Vertical:   0,
			expected:   []byte{0, 0x80, 0, 0},
		},
		{
			name:       "Rule 9.4.2-2/2",
			Horizontal: 0,
			Vertical:   1,
			expected:   []byte{0, 0, 0, 0x80},
		},
		{
			name:       "Rule 9.4.2-3/1",
			Horizontal: 1,
			Vertical:   1,
			expected:   []byte{0, 0x80, 0, 0x80},
		},
		{
			name:       "Rule 9.4.2-3/2",
			Horizontal: 0.0078125,
			Vertical:   0,
			expected:   []byte{0, 0x01, 0, 0},
		},
		{
			name:       "Rule 9.4.2-3/3",
			Horizontal: 0,
			Vertical:   0.0078125,
			expected:   []byte{0, 0, 0, 0x01},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := BeamWidth{}
			b.Horizontal = tc.Horizontal
			b.Vertical = tc.Vertical
			assert.Equal(t, tc.Horizontal, b.Horizontal)
			assert.Equal(t, tc.Vertical, b.Vertical)
			// Pack
			packed := b.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			b.Unpack(packed)
			assert.Equal(t, tc.Horizontal, b.Horizontal)
			assert.Equal(t, tc.Vertical, b.Vertical)
		})
	}
}

func TestEbNoBERBytes(t *testing.T) {
	e := EbNoBER{}
	assert.Equal(t, ebNoBERBytes, e.Size())
}

func TestEbNoBERDefault(t *testing.T) {
	e := NewEbNoBer()
	assert.Equal(t, float64(32767), e.Ebno)
	assert.Equal(t, float64(32767), e.Ber)
	// Pack
	packed := e.Pack()
	expected := []byte{0xFF, 0x80, 0xFF, 0x80}
	assert.Equal(t, expected, packed)
	// Unpack
	e.Unpack(packed)
	assert.Equal(t, float64(-1), e.Ebno)
	assert.Equal(t, float64(-1), e.Ber)
}

func TestEbNoBER(t *testing.T) {
	cases := []struct {
		name     string
		Ebno     float64
		Ber      float64
		expected []byte
	}{
		{
			name:     "Rule 9.5.17-1",
			Ebno:     0,
			Ber:      0,
			expected: []byte{0, 0, 0, 0},
		},
		{
			name:     "Rule 9.5.17-2",
			Ebno:     0,
			Ber:      0,
			expected: []byte{0, 0, 0, 0},
			// msg: "Not applicable for test"
		},
		{
			name:     "Rule 9.5.17-3",
			Ebno:     0, // radix test
			Ber:      0,
			expected: []byte{0, 0, 0, 0},
		},
		{
			name:     "Rule 9.5.17-4/1",
			Ebno:     -256,
			Ber:      0,
			expected: []byte{0x80, 0, 0, 0},
		},
		{
			name:     "Rule 9.5.17-4/2",
			Ebno:     255.984375,
			Ber:      0,
			expected: []byte{0x7f, 0xfe, 0, 0},
		},
		{
			name:     "Rule 9.5.17-4/3",
			Ebno:     255.9921875, // max possible value 0x7FFF
			Ber:      0,
			expected: []byte{0x7f, 0xff, 0, 0},
		},
		{
			name:     "Rule 9.5.17-5/1",
			Ebno:     0.0078125,
			Ber:      0,
			expected: []byte{0, 0x01, 0, 0},
		},
		{
			name:     "Rule 9.5.17-5/2",
			Ebno:     1 - 0.0078125,
			Ber:      0,
			expected: []byte{0, 0x7F, 0, 0},
		},
		{
			name:     "Rule 9.5.17-6",
			Ebno:     1.0,
			Ber:      0,
			expected: []byte{0, 0x80, 0, 0},
			// msg: "n/a"
		},
		{
			name:     "Rule 9.5.17-7",
			Ebno:     1,
			Ber:      0,
			expected: []byte{0, 0x80, 0, 0},
		},
		{
			name:     "Rule 9.5.17-8/1",
			Ebno:     0,
			Ber:      -256,
			expected: []byte{0, 0, 0x80, 0},
		},
		{
			name:     "Rule 9.5.17-8/2",
			Ebno:     0,
			Ber:      255.984375,
			expected: []byte{0, 0, 0x7f, 0xfe},
		},
		{
			name:     "Rule 9.5.17-8/3",
			Ebno:     0,
			Ber:      0.0078125,
			expected: []byte{0, 0, 0, 0x01},
		},
		{
			name:     "Rule 9.5.17-8/4",
			Ebno:     0,
			Ber:      1 - 0.0078125,
			expected: []byte{0, 0, 0, 0x7F},
		},
		{
			name:     "Rule 9.5.17-9",
			Ebno:     0,
			Ber:      1,
			expected: []byte{0, 0, 0, 0x80},
		},
		{
			name:     "Rule 9.5.17-10",
			Ebno:     1,
			Ber:      0,
			expected: []byte{0, 0x80, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := EbNoBER{}
			e.Ebno = tc.Ebno
			e.Ber = tc.Ber
			assert.Equal(t, tc.Ebno, e.Ebno)
			assert.Equal(t, tc.Ber, e.Ber)
			// Pack
			packed := e.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			e.Unpack(packed)
			assert.Equal(t, tc.Ebno, e.Ebno)
			assert.Equal(t, tc.Ber, e.Ber)
		})
	}
}

func TestThresholdBytes(t *testing.T) {
	th := Threshold{}
	assert.Equal(t, thresholdBytes, th.Size())
}

func TestThresholdDefault(t *testing.T) {
	th := Threshold{}
	assert.Equal(t, float64(0), th.Stage1)
	assert.Equal(t, float64(0), th.Stage2)
	// Pack
	packed := th.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	th.Unpack(packed)
	assert.Equal(t, float64(0), th.Stage1)
	assert.Equal(t, float64(0), th.Stage2)
}

func TestThreshold(t *testing.T) {
	cases := []struct {
		name     string
		Stage1   float64
		Stage2   float64
		expected []byte
	}{
		{
			name:     "Rule 9.6.2.9-1",
			Stage1:   0,
			Stage2:   0,
			expected: []byte{0, 0, 0, 0},
		},
		{
			name:     "Rule 9.6.2.9-2/1",
			Stage1:   1,
			Stage2:   1,
			expected: []byte{0, 0x80, 0, 0x80},
		},
		{
			name:     "Rule 9.6.2.9-2/2",
			Stage1:   -1,
			Stage2:   -1,
			expected: []byte{0xff, 0x80, 0xff, 0x80},
		},
		{
			name:     "Rule 9.6.2.9-2/3",
			Stage1:   0.0078125,
			Stage2:   0,
			expected: []byte{0, 0x01, 0, 0},
		},
		{
			name:     "Rule 9.6.2.9-2/4",
			Stage1:   -.0078125,
			Stage2:   0,
			expected: []byte{0xff, 0xff, 0, 0},
		},
		{
			name:     "Rule 9.6.2.9-2/5",
			Stage1:   .0078125,
			Stage2:   .0078125,
			expected: []byte{0, 0x01, 0, 0x01},
		},
		{
			name:     "Rule 9.6.2.9-2/6",
			Stage1:   -.0078125,
			Stage2:   -.0078125,
			expected: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			name:     "Rule 9.6.2.9-3/1",
			Stage1:   1,
			Stage2:   0,
			expected: []byte{0, 0x80, 0, 0},
		},
		{
			name:     "Rule 9.6.2.9-3/2",
			Stage1:   -1,
			Stage2:   0,
			expected: []byte{0xff, 0x80, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			th := Threshold{}
			th.Stage1 = tc.Stage1
			th.Stage2 = tc.Stage2
			assert.Equal(t, tc.Stage1, th.Stage1)
			assert.Equal(t, tc.Stage2, th.Stage2)
			// Pack
			packed := th.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			th.Unpack(packed)
			assert.Equal(t, tc.Stage1, th.Stage1)
			assert.Equal(t, tc.Stage2, th.Stage2)
		})
	}
}

func TestInterceptPointsBytes(t *testing.T) {
	ip := InterceptPoints{}
	assert.Equal(t, interceptPointsBytes, ip.Size())
}

func TestInterceptPointsDefault(t *testing.T) {
	ip := InterceptPoints{}
	assert.Equal(t, float64(0), ip.SecondOrder)
	assert.Equal(t, float64(0), ip.ThirdOrder)
	// Pack
	packed := []byte{0xFF, 0xFF, 0xFF, 0xFF}
	packed = ip.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	ip.Unpack(packed)
	assert.Equal(t, float64(0), ip.SecondOrder)
	assert.Equal(t, float64(0), ip.ThirdOrder)
}

func TestInterceptPoints(t *testing.T) {
	cases := []struct {
		name        string
		SecondOrder float64
		ThirdOrder  float64
		expected    []byte
	}{
		{
			name:        "Rule 9.5.6-1",
			SecondOrder: 0,
			ThirdOrder:  1,
			expected:    []byte{0, 0, 0, 0x80},
		},
		{
			name:        "Rule 9.5.6-2",
			SecondOrder: 1,
			ThirdOrder:  0,
			expected:    []byte{0, 0x80, 0, 0},
		},
		{
			name:        "Rule 9.5.6-3",
			SecondOrder: 1,
			ThirdOrder:  1,
			expected:    []byte{0, 0x80, 0, 0x80},
		},
		{
			name:        "Rule 9.5.6-4",
			SecondOrder: 1,
			ThirdOrder:  1,
			expected:    []byte{0, 0x80, 0, 0x80},
		},
		{
			name:        "Rule 9.5.6-5/1",
			SecondOrder: -256,
			ThirdOrder:  -256,
			expected:    []byte{0x80, 0, 0x80, 0},
		},
		{
			name:        "Rule 9.5.6-5/2",
			SecondOrder: 255,
			ThirdOrder:  255,
			expected:    []byte{0x7f, 0x80, 0x7f, 0x80},
		},
		{
			name:        "Rule 9.5.6-6/1",
			SecondOrder: 255.9921875,
			ThirdOrder:  0,
			expected:    []byte{0x7f, 0xff, 0, 0},
		},
		{
			name:        "Rule 9.5.6-6/2",
			SecondOrder: 0,
			ThirdOrder:  255.9921875,
			expected:    []byte{0, 0, 0x7f, 0xff},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ip := InterceptPoints{}
			ip.SecondOrder = tc.SecondOrder
			ip.ThirdOrder = tc.ThirdOrder
			assert.Equal(t, tc.SecondOrder, ip.SecondOrder)
			assert.Equal(t, tc.ThirdOrder, ip.ThirdOrder)
			// Pack
			packed := ip.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			ip.Unpack(packed)
			assert.Equal(t, tc.SecondOrder, ip.SecondOrder)
			assert.Equal(t, tc.ThirdOrder, ip.ThirdOrder)
		})
	}
}

func TestSNRNoiseBytes(t *testing.T) {
	s := SNRNoise{}
	assert.Equal(t, snrNoiseBytes, s.Size())
}

func TestSNRNoiseDefault(t *testing.T) {
	s := SNRNoise{}
	assert.Equal(t, float64(0), s.Snr)
	assert.Equal(t, float64(0), s.Noise)
	// Pack
	packed := s.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	s.Unpack(packed)
	assert.Equal(t, float64(0), s.Snr)
	assert.Equal(t, float64(0), s.Noise)
}

func TestSNRNoise(t *testing.T) {
	cases := []struct {
		name     string
		Snr      float64
		Noise    float64
		expected []byte
	}{
		{
			name:     "Rule 9.5.7-1",
			Snr:      0,
			Noise:    0,
			expected: []byte{0, 0, 0, 0},
		},
		{
			name:     "Rule 9.5.7-2",
			Snr:      1,
			Noise:    0,
			expected: []byte{0, 0x80, 0, 0},
		},
		{
			name:     "Rule 9.5.7-3",
			Snr:      1,
			Noise:    0,
			expected: []byte{0, 0x80, 0, 0},
		},
		{
			name:     "Rule 9.5.7-4/1",
			Snr:      -256,
			Noise:    0,
			expected: []byte{0x80, 0, 0, 0},
		},
		{
			name:     "Rule 9.5.7-4/2",
			Snr:      255.984375,
			Noise:    0,
			expected: []byte{0x7f, 0xfe, 0, 0},
		},
		{
			name:     "Rule 9.5.7-5/1",
			Snr:      0.0078125,
			Noise:    0,
			expected: []byte{0, 0x01, 0, 0},
		},
		{
			name:     "Rule 9.5.7-5/2",
			Snr:      -.0078125,
			Noise:    0,
			expected: []byte{0xff, 0xff, 0, 0},
		},
		{
			name:     "Rule 9.5.7-6",
			Snr:      0,
			Noise:    0,
			expected: []byte{0, 0, 0, 0},
		},
		{
			name:     "Rule 9.5.7-7",
			Snr:      0,
			Noise:    255.9921875,
			expected: []byte{0, 0, 0x7f, 0xff},
		},
		{
			name:     "Rule 9.5.7-8/1",
			Snr:      0,
			Noise:    255.9921875,
			expected: []byte{0, 0, 0x7f, 0xff},
		},
		{
			name:     "Rule 9.5.7-8/2",
			Snr:      0,
			Noise:    0,
			expected: []byte{0, 0, 0, 0},
		},
		{
			name:     "Rule 9.5.7-9/1",
			Snr:      .0078125,
			Noise:    .0078125,
			expected: []byte{0, 0x01, 0, 0x01},
		},
		{
			name:     "Rule 9.5.7-9/2",
			Snr:      -.0078125,
			Noise:    -.0078125,
			expected: []byte{0xff, 0xff, 0xff, 0xff},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := SNRNoise{}
			s.Snr = tc.Snr
			s.Noise = tc.Noise
			assert.Equal(t, tc.Snr, s.Snr)
			assert.Equal(t, tc.Noise, s.Noise)
			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			s.Unpack(packed)
			assert.Equal(t, tc.Snr, s.Snr)
			assert.Equal(t, tc.Noise, s.Noise)
		})
	}
}

func TestSpectrumTypeBytes(t *testing.T) {
	s := SpectrumType{}
	assert.Equal(t, spectrumTypeBytes, s.Size())
}

func TestSpectrumTypeDefault(t *testing.T) {
	s := SpectrumType{}
	assert.Equal(t, uint8(0), s.SpectrumType)
	assert.Equal(t, uint8(0), s.AveragingType)
	assert.Equal(t, uint8(0), s.WindowTime)
	// Pack
	packed := s.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	s.Unpack(packed)
	assert.Equal(t, uint8(0), s.SpectrumType)
	assert.Equal(t, uint8(0), s.AveragingType)
	assert.Equal(t, uint8(0), s.WindowTime)
}

func TestSpectrumType(t *testing.T) {
	cases := []struct {
		name          string
		SpectrumType  uint8
		AveragingType uint8
		WindowTime    uint8
		expected      []byte
	}{
		// Spectrum Type
		{
			name:          "Rule 9.6.1.1.1-1/1",
			SpectrumType:  0, // Default
			AveragingType: 0,
			WindowTime:    0,
			expected:      []byte{0, 0, 0, 0},
		},
		{
			name:          "Rule 9.6.1.1.1-1/2",
			SpectrumType:  1, // Log power (dB)
			AveragingType: 0,
			WindowTime:    0,
			expected:      []byte{0, 0, 0, 0x1},
		},
		{
			name:          "Rule 9.6.1.1.1-1/3",
			SpectrumType:  2, // Cartesian (I, Q)
			AveragingType: 0,
			WindowTime:    0,
			expected:      []byte{0, 0, 0, 0x2},
		},
		{
			name:          "Rule 9.6.1.1.1-1/4",
			SpectrumType:  3, // Polar (magnitude, phase)
			AveragingType: 0,
			WindowTime:    0,
			expected:      []byte{0, 0, 0, 0x3},
		},
		{
			name:          "Rule 9.6.1.1.1-1/5",
			SpectrumType:  4, // Magnitude
			AveragingType: 0,
			WindowTime:    0,
			expected:      []byte{0, 0, 0, 0x4},
		},
		{
			name:          "Rule 9.6.1.1.1-5",
			SpectrumType:  0, // Default
			AveragingType: 0,
			WindowTime:    0,
			expected:      []byte{0, 0, 0, 0},
		},
		// Averaging Type
		{
			name:          "Rule 9.6.1.1.2-1/1",
			SpectrumType:  0,
			AveragingType: 0, // No averaging
			WindowTime:    0,
			expected:      []byte{0, 0, 0, 0},
		},
		{
			name:          "Rule 9.6.1.1.2-1/2",
			SpectrumType:  0,
			AveragingType: 1, // Linear Averaging
			WindowTime:    0,
			expected:      []byte{0, 0, 0x01, 0},
		},
		{
			name:          "Rule 9.6.1.1.2-1/3",
			SpectrumType:  0,
			AveragingType: 2, // Peak Hold
			WindowTime:    0,
			expected:      []byte{0, 0, 0x02, 0},
		},
		{
			name:          "Rule 9.6.1.1.2-1/4",
			SpectrumType:  0,
			AveragingType: 4, // Min Hold
			WindowTime:    0,
			expected:      []byte{0, 0, 0x04, 0},
		},
		{
			name:          "Rule 9.6.1.1.2-1/5",
			SpectrumType:  0,
			AveragingType: 8, // Exponential Averaging
			WindowTime:    0,
			expected:      []byte{0, 0, 0x8, 0},
		},
		{
			name:          "Rule 9.6.1.1.2-1/6",
			SpectrumType:  0,
			AveragingType: 16, //Median Averaging
			WindowTime:    0,
			expected:      []byte{0, 0, 0x10, 0},
		},
		{
			name:          "Rule 9.6.1.1.2-1/7",
			SpectrumType:  0,
			AveragingType: 32, // Smoothing
			WindowTime:    0,
			expected:      []byte{0, 0, 0x20, 0},
		},
		// Window Type
		{
			name:          "Rule 9.6.1.1.3-1/1",
			SpectrumType:  0,
			AveragingType: 0,
			WindowTime:    0, // Overlap is not controlled
			expected:      []byte{0, 0, 0, 0},
		},
		{
			name:          "Rule 9.6.1.1.3-1/2",
			SpectrumType:  0,
			AveragingType: 0,
			WindowTime:    1, // Percent overlap
			expected:      []byte{0, 0x01, 0, 0},
		},
		{
			name:          "Rule 9.6.1.1.3-1/3",
			SpectrumType:  0,
			AveragingType: 0,
			WindowTime:    2, // Samples
			expected:      []byte{0, 0x02, 0, 0},
		},
		{
			name:          "Rule 9.6.1.1.3-1/4",
			SpectrumType:  0,
			AveragingType: 0,
			WindowTime:    3, // Time
			expected:      []byte{0, 0x03, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := SpectrumType{}
			s.WindowTime = tc.WindowTime
			s.AveragingType = tc.AveragingType
			s.SpectrumType = tc.SpectrumType
			assert.Equal(t, tc.WindowTime, s.WindowTime)
			assert.Equal(t, tc.AveragingType, s.AveragingType)
			assert.Equal(t, tc.SpectrumType, s.SpectrumType)
			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			s.Unpack(packed)
			assert.Equal(t, tc.WindowTime, s.WindowTime)
			assert.Equal(t, tc.AveragingType, s.AveragingType)
			assert.Equal(t, tc.SpectrumType, s.SpectrumType)
		})
	}
}

func TestWindowTypeBytes(t *testing.T) {
	ip := WindowType{}
	assert.Equal(t, windowTypeBytes, ip.Size())
}

func TestWindowTypeDefault(t *testing.T) {
	wt := WindowType{}
	assert.Equal(t, uint8(0), wt.WindowType)
	// Pack
	packed := wt.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	wt.Unpack(packed)
	assert.Equal(t, uint8(0), wt.WindowType)
}

func TestWindowType(t *testing.T) {
	cases := []struct {
		name       string
		WindowType uint8
		expected   []byte
	}{
		{
			name:       "Rule 9.6.1.2-1",
			WindowType: 0,
			expected:   []byte{0, 0, 0, 0},
		},
		{
			name:       "Rule 9.6.1.2-2/1",
			WindowType: 2, // Hanning
			expected:   []byte{0, 0, 0, 0x2},
		},
		{
			name:       "Rule 9.6.1.2-2/2",
			WindowType: 14, // Poisson
			expected:   []byte{0, 0, 0, 0x0e},
		},
		{
			name:       "Rule 9.6.1.2-2/3",
			WindowType: 38, // Blackman
			expected:   []byte{0, 0, 0, 0x26},
		},
		{
			name:       "Rule 9.6.1.2-3",
			WindowType: 0, // Default (Rectangular)
			expected:   []byte{0, 0, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			wt := WindowType{}
			wt.WindowType = tc.WindowType
			assert.Equal(t, tc.WindowType, wt.WindowType)
			// Pack
			packed := wt.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			wt.Unpack(packed)
			assert.Equal(t, tc.WindowType, wt.WindowType)
		})
	}
}

func TestSpectrumF1F2IndiciesBytes(t *testing.T) {
	s := SpectrumF1F2Indicies{}
	assert.Equal(t, spectrumF1F2IndiciesBytes, s.Size())
}

func TestSpectrumF1F2IndiciesDefault(t *testing.T) {
	s := SpectrumF1F2Indicies{}
	assert.Equal(t, uint32(0), s.F1Index)
	assert.Equal(t, uint32(0), s.F2Index)
	// Pack
	packed := s.Pack()
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	s.Unpack(packed)
	assert.Equal(t, uint32(0), s.F1Index)
	assert.Equal(t, uint32(0), s.F2Index)
}

func TestSpectrumF1F2Indicies(t *testing.T) {
	cases := []struct {
		name     string
		F1Index  uint32
		F2Index  uint32
		expected []byte
	}{
		{
			name:     "Rule 9.6.1.9-1",
			F1Index:  0,
			F2Index:  0,
			expected: []byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.6.1.9-2/1",
			F1Index:  0,
			F2Index:  49, // N == 100 points (arbitrary)
			expected: []byte{0, 0, 0, 0x31, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.6.1.9-2/2",
			F1Index:  0,
			F2Index:  95, // N == 100 points
			expected: []byte{0, 0, 0, 0x5f, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.6.1.9-3",
			F1Index:  0,
			F2Index:  50, // N == 100 points
			expected: []byte{0, 0, 0, 0x32, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.6.1.9-4",
			F1Index:  0,  // Default
			F2Index:  50, // N == 100 points
			expected: []byte{0, 0, 0, 0x32, 0, 0, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := SpectrumF1F2Indicies{}
			s.F1Index = tc.F1Index
			s.F2Index = tc.F2Index
			assert.Equal(t, tc.F1Index, s.F1Index)
			assert.Equal(t, tc.F2Index, s.F2Index)
			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			s.Unpack(packed)
			assert.Equal(t, tc.F1Index, s.F1Index)
			assert.Equal(t, tc.F2Index, s.F2Index)
		})
	}
}

func TestSpectrumBytes(t *testing.T) {
	s := Spectrum{}
	assert.Equal(t, spectrumBytes, s.Size())
}

func TestSpectrumDefault(t *testing.T) {
	s := Spectrum{}
	assert.Equal(t, SpectrumType{}, s.SpectrumType)
	assert.Equal(t, WindowType{}, s.WindowType)
	assert.Equal(t, uint32(0), s.NumberTransformPoints)
	assert.Equal(t, uint32(0), s.NumberWindowPoints)
	assert.Equal(t, uint64(0), s.Resolution)
	assert.Equal(t, uint64(0), s.Span)
	assert.Equal(t, uint32(0), s.NumberAverages)
	assert.Equal(t, uint32(0), s.WeightingFactor)
	assert.Equal(t, SpectrumF1F2Indicies{}, s.SpectrumF1F2Indicies)
	assert.Equal(t, uint32(0), s.WindowTimeDelta)
	// Pack
	packed := s.Pack()
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	s.Unpack(packed)
	assert.Equal(t, SpectrumType{}, s.SpectrumType)
	assert.Equal(t, WindowType{}, s.WindowType)
	assert.Equal(t, uint32(0), s.NumberTransformPoints)
	assert.Equal(t, uint32(0), s.NumberWindowPoints)
	assert.Equal(t, uint64(0), s.Resolution)
	assert.Equal(t, uint64(0), s.Span)
	assert.Equal(t, uint32(0), s.NumberAverages)
	assert.Equal(t, uint32(0), s.WeightingFactor)
	assert.Equal(t, SpectrumF1F2Indicies{}, s.SpectrumF1F2Indicies)
	assert.Equal(t, uint32(0), s.WindowTimeDelta)
}

func TestSpectrum(t *testing.T) {
	cases := []struct {
		name                  string
		SpectrumType          SpectrumType
		WindowType            WindowType
		NumberTransformPoints uint32
		NumberWindowPoints    uint32
		Resolution            uint64
		Span                  uint64
		NumberAverages        uint32
		WeightingFactor       uint32
		SpectrumF1F2Indicies  SpectrumF1F2Indicies
		WindowTimeDelta       uint32
		expected              []byte
	}{
		{
			name:                  "Rule 9.6.1-1",
			SpectrumType:          SpectrumType{1, 0, 0},
			WindowType:            WindowType{14},
			NumberTransformPoints: 100,
			NumberWindowPoints:    0,
			Resolution:            0,
			Span:                  0,
			NumberAverages:        1,
			WeightingFactor:       0,
			SpectrumF1F2Indicies:  SpectrumF1F2Indicies{49, 0},
			WindowTimeDelta:       0,
			expected:              []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x64, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x1},
		},
		{
			name:                  "Rule 9.6.1-2",
			SpectrumType:          SpectrumType{},
			WindowType:            WindowType{},
			NumberTransformPoints: 0,
			NumberWindowPoints:    0,
			Resolution:            0,
			Span:                  0,
			NumberAverages:        0,
			WeightingFactor:       0,
			SpectrumF1F2Indicies:  SpectrumF1F2Indicies{},
			WindowTimeDelta:       0,
			expected:              []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:                  "Rule 9.6.1-3",
			SpectrumType:          SpectrumType{},
			WindowType:            WindowType{},
			NumberTransformPoints: 0,
			NumberWindowPoints:    0,
			Resolution:            0,
			Span:                  0,
			NumberAverages:        0,
			WeightingFactor:       0,
			SpectrumF1F2Indicies:  SpectrumF1F2Indicies{},
			WindowTimeDelta:       0,
			expected:              []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Spectrum{}
			s.SpectrumType = tc.SpectrumType
			s.WindowType = tc.WindowType
			s.NumberTransformPoints = tc.NumberTransformPoints
			s.NumberWindowPoints = tc.NumberWindowPoints
			s.Resolution = tc.Resolution
			s.Span = tc.Span
			s.NumberAverages = tc.NumberAverages
			s.WeightingFactor = tc.WeightingFactor
			s.SpectrumF1F2Indicies = tc.SpectrumF1F2Indicies
			s.WindowTimeDelta = tc.WindowTimeDelta
			assert.Equal(t, tc.SpectrumType, s.SpectrumType)
			assert.Equal(t, tc.WindowType, s.WindowType)
			assert.Equal(t, tc.NumberTransformPoints, s.NumberTransformPoints)
			assert.Equal(t, tc.NumberWindowPoints, s.NumberWindowPoints)
			assert.Equal(t, tc.Resolution, s.Resolution)
			assert.Equal(t, tc.Span, s.Span)
			assert.Equal(t, tc.NumberAverages, s.NumberAverages)
			assert.Equal(t, tc.WeightingFactor, s.WeightingFactor)
			assert.Equal(t, tc.SpectrumF1F2Indicies, s.SpectrumF1F2Indicies)
			assert.Equal(t, tc.WindowTimeDelta, s.WindowTimeDelta)
			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			s.Unpack(packed)
			assert.Equal(t, tc.SpectrumType, s.SpectrumType)
			assert.Equal(t, tc.WindowType, s.WindowType)
			assert.Equal(t, tc.NumberTransformPoints, s.NumberTransformPoints)
			assert.Equal(t, tc.NumberWindowPoints, s.NumberWindowPoints)
			assert.Equal(t, tc.Resolution, s.Resolution)
			assert.Equal(t, tc.Span, s.Span)
			assert.Equal(t, tc.NumberAverages, s.NumberAverages)
			assert.Equal(t, tc.WeightingFactor, s.WeightingFactor)
			assert.Equal(t, tc.SpectrumF1F2Indicies, s.SpectrumF1F2Indicies)
			assert.Equal(t, tc.WindowTimeDelta, s.WindowTimeDelta)
		})
	}
}

func TestIndexListBytes(t *testing.T) {
	s := IndexList{}
	expectedBytes := uint32(8)
	if s.NumEntries != 0 {
		expectedBytes += uint32(4*s.NumEntries) / 8
	}
	assert.Equal(t, expectedBytes, s.Size())
}

func TestIndexListDefault(t *testing.T) {
	s := IndexList{}
	assert.Equal(t, uint32(0), s.TotalSize)
	assert.Equal(t, uint8(0), s.EntrySize)
	assert.Equal(t, uint32(0), s.NumEntries)
	assert.Empty(t, s.Entries)

	// Pack
	packed := s.Pack()
	expected := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	assert.Equal(t, expected, packed)

	// Unpack
	s.Unpack(packed)
	assert.Equal(t, uint32(0), s.TotalSize)
	assert.Equal(t, uint8(0), s.EntrySize)
	assert.Equal(t, uint32(0), s.NumEntries)
	assert.Empty(t, s.Entries)
}

func TestIndexList(t *testing.T) {
	cases := []struct {
		name       string
		TotalSize  uint32
		EntrySize  uint8
		NumEntries uint32
		Entries    []uint32
		expected   []byte
	}{
		{
			name:       "Rule 9.3.2.4/1",
			TotalSize:  1,
			EntrySize:  0xF,
			NumEntries: 0,
			Entries:    []uint32{},
			expected:   []byte{0x0, 0x0, 0x0, 0x1, 0xF0, 0x0, 0x0, 0x0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := IndexList{
				TotalSize:  tc.TotalSize,
				EntrySize:  tc.EntrySize,
				NumEntries: tc.NumEntries,
				Entries:    tc.Entries,
			}

			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			var unpacked IndexList
			unpacked.Unpack(packed)
			assert.Equal(t, s.TotalSize, unpacked.TotalSize)
			assert.Equal(t, s.EntrySize, unpacked.EntrySize)
			assert.Equal(t, s.NumEntries, unpacked.NumEntries)
			assert.Equal(t, s.Entries, unpacked.Entries)
		})
	}
}

func TestSectorStepScanCIFBytes(t *testing.T) {
	s := SectorStepScanCIF{}
	assert.Equal(t, sectorStepScanCIFBytes, s.Size())
}

func TestSectorStepScanCIFDefault(t *testing.T) {
	s := SectorStepScanCIF{}
	assert.Equal(t, false, s.SectorNumber)
	assert.Equal(t, false, s.F1StartFrequency)
	assert.Equal(t, false, s.F2StartFrequency)
	assert.Equal(t, false, s.ResolutionBandwidth)
	assert.Equal(t, false, s.TuneStepSize)
	assert.Equal(t, false, s.NumberPoints)
	assert.Equal(t, false, s.DefaultGain)
	assert.Equal(t, false, s.Threshold)
	assert.Equal(t, false, s.DwellTime)
	assert.Equal(t, false, s.StartTime)
	assert.Equal(t, false, s.Time3)
	assert.Equal(t, false, s.Time4)
	// Pack
	packed := s.Pack()
	expected := []byte{0x0, 0x0, 0x0, 0x0} // All bits are zero
	assert.Equal(t, expected, packed)
	// Unpack
	s.Unpack(packed)
	assert.Equal(t, false, s.SectorNumber)
	assert.Equal(t, false, s.F1StartFrequency)
	assert.Equal(t, false, s.F2StartFrequency)
	assert.Equal(t, false, s.ResolutionBandwidth)
	assert.Equal(t, false, s.TuneStepSize)
	assert.Equal(t, false, s.NumberPoints)
	assert.Equal(t, false, s.DefaultGain)
	assert.Equal(t, false, s.Threshold)
	assert.Equal(t, false, s.DwellTime)
	assert.Equal(t, false, s.StartTime)
	assert.Equal(t, false, s.Time3)
	assert.Equal(t, false, s.Time4)
}

func TestSectorStepScanCIF(t *testing.T) {
	cases := []struct {
		name                string
		SectorNumber        bool
		F1StartFrequency    bool
		F2StartFrequency    bool
		ResolutionBandwidth bool
		TuneStepSize        bool
		NumberPoints        bool
		DefaultGain         bool
		Threshold           bool
		DwellTime           bool
		StartTime           bool
		Time3               bool
		Time4               bool
		expected            []byte
	}{
		{
			name:                "Rule 9.6.2.14-2/1",
			SectorNumber:        true,
			F1StartFrequency:    true,
			F2StartFrequency:    false,
			ResolutionBandwidth: true,
			TuneStepSize:        false,
			NumberPoints:        true,
			DefaultGain:         true,
			Threshold:           false,
			DwellTime:           true,
			StartTime:           false,
			Time3:               true,
			Time4:               false,
			expected:            []byte{0xd6, 0xa0, 0x0, 0x0}, // Packed representation
		},
		{
			name:                "Rule 9.6.2.14-2/2",
			SectorNumber:        false,
			F1StartFrequency:    false,
			F2StartFrequency:    false,
			ResolutionBandwidth: false,
			TuneStepSize:        false,
			NumberPoints:        false,
			DefaultGain:         false,
			Threshold:           false,
			DwellTime:           false,
			StartTime:           false,
			Time3:               false,
			Time4:               false,
			expected:            []byte{0x0, 0x0, 0x0, 0x0}, // All bits are zero
		},
		{
			name:                "Rule 9.6.2.14-2/3",
			SectorNumber:        true,
			F1StartFrequency:    true,
			F2StartFrequency:    true,
			ResolutionBandwidth: true,
			TuneStepSize:        true,
			NumberPoints:        true,
			DefaultGain:         true,
			Threshold:           true,
			DwellTime:           true,
			StartTime:           true,
			Time3:               true,
			Time4:               true,
			expected:            []byte{0xFF, 0xF0, 0x0, 0x0}, // Packed representation for all true
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := SectorStepScanCIF{}
			s.SectorNumber = tc.SectorNumber
			s.F1StartFrequency = tc.F1StartFrequency
			s.F2StartFrequency = tc.F2StartFrequency
			s.ResolutionBandwidth = tc.ResolutionBandwidth
			s.TuneStepSize = tc.TuneStepSize
			s.NumberPoints = tc.NumberPoints
			s.DefaultGain = tc.DefaultGain
			s.Threshold = tc.Threshold
			s.DwellTime = tc.DwellTime
			s.StartTime = tc.StartTime
			s.Time3 = tc.Time3
			s.Time4 = tc.Time4
			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			s.Unpack(packed)
			assert.Equal(t, tc.SectorNumber, s.SectorNumber)
			assert.Equal(t, tc.F1StartFrequency, s.F1StartFrequency)
			assert.Equal(t, tc.F2StartFrequency, s.F2StartFrequency)
			assert.Equal(t, tc.ResolutionBandwidth, s.ResolutionBandwidth)
			assert.Equal(t, tc.TuneStepSize, s.TuneStepSize)
			assert.Equal(t, tc.NumberPoints, s.NumberPoints)
			assert.Equal(t, tc.DefaultGain, s.DefaultGain)
			assert.Equal(t, tc.Threshold, s.Threshold)
			assert.Equal(t, tc.DwellTime, s.DwellTime)
			assert.Equal(t, tc.StartTime, s.StartTime)
			assert.Equal(t, tc.Time3, s.Time3)
			assert.Equal(t, tc.Time4, s.Time4)
		})
	}
}

func TestSectorStepScanRecordBytes(t *testing.T) {
	s := SectorStepScanRecord{}
	assert.Equal(t, sectorStepScanRecordBytes, s.Size())
}

// func TestSectorStepScanRecordDefault(t *testing.T) {
// 	s := SectorStepScanRecord{}
// 	assert.Equal(t, uint32(0), s.SectorNumber)
// 	assert.Equal(t, uint64(0), s.F1StartFrequency)
// 	assert.Equal(t, uint64(0), s.F2StopFrequency)
// 	assert.Equal(t, uint64(0), s.ResolutionBandwidth)
// 	assert.Equal(t, uint64(0), s.TuneStepSize)
// 	assert.Equal(t, uint32(0), s.NumberPoints)
// 	assert.Equal(t, Gain{}, s.DefaultGain) // Assuming Gain is a struct
// 	assert.Equal(t, Gain{}, s.Threshold)   // Assuming Gain is a struct
// 	assert.Equal(t, uint32(0), s.DwellTime)
// 	assert.Equal(t, uint32(0), s.StartTime)
// 	assert.Equal(t, uint32(0), s.Time3)
// 	assert.Equal(t, uint32(0), s.Time4)

// 	// Pack
// 	packed := s.Pack()
// 	expected := make([]byte, 80)
// 	assert.Equal(t, expected, packed)

// 	// Unpack
// 	s.Unpack(packed)
// 	assert.Equal(t, uint32(0), s.SectorNumber)
// 	assert.Equal(t, uint64(0), s.F1StartFrequency)
// 	assert.Equal(t, uint64(0), s.F2StopFrequency)
// 	assert.Equal(t, uint64(0), s.ResolutionBandwidth)
// 	assert.Equal(t, uint64(0), s.TuneStepSize)
// 	assert.Equal(t, uint32(0), s.NumberPoints)
// 	assert.Equal(t, Gain{}, s.DefaultGain)
// 	assert.Equal(t, Gain{}, s.Threshold)
// 	assert.Equal(t, uint32(0), s.DwellTime)
// 	assert.Equal(t, uint32(0), s.StartTime)
// 	assert.Equal(t, uint32(0), s.Time3)
// 	assert.Equal(t, uint32(0), s.Time4)
// }

// func TestSectorStepScanRecord(t *testing.T) {
// 	cases := []struct {
// 		name                string
// 		SectorNumber        uint32
// 		F1StartFrequency    uint64
// 		F2StopFrequency     uint64
// 		ResolutionBandwidth uint64
// 		TuneStepSize        uint64
// 		NumberPoints        uint32
// 		DefaultGain         Gain
// 		Threshold           Gain
// 		DwellTime           uint32
// 		StartTime           uint32
// 		Time3               uint32
// 		Time4               uint32
// 		expected            []byte
// 	}{
// 		{
// 			name:                "Populate arbitrary record",
// 			SectorNumber:        1,
// 			F1StartFrequency:    1000,
// 			F2StopFrequency:     2000,
// 			ResolutionBandwidth: 100,
// 			TuneStepSize:        200,
// 			NumberPoints:        256,
// 			DefaultGain:         Gain{},
// 			Threshold:           Gain{},
// 			DwellTime:           3,
// 			StartTime:           10,
// 			Time3:               15,
// 			Time4:               20,
// 			expected:            []byte{0x0, 0x0, 0x0, 0x14, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x64, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xd0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe8, 0x0, 0x0, 0x0, 0x1},
// 		},
// 		{
// 			name:                "Record with only required fields",
// 			SectorNumber:        1,
// 			F1StartFrequency:    500,
// 			F2StopFrequency:     0,
// 			ResolutionBandwidth: 0,
// 			TuneStepSize:        0,
// 			NumberPoints:        0,
// 			DefaultGain:         Gain{},
// 			Threshold:           Gain{},
// 			DwellTime:           0,
// 			StartTime:           0,
// 			Time3:               5,
// 			Time4:               0,
// 			expected:            []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4, 0x0, 0x0, 0x0, 0x1},
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			s := SectorStepScanRecord{
// 				SectorNumber:        tc.SectorNumber,
// 				F1StartFrequency:    tc.F1StartFrequency,
// 				F2StopFrequency:     tc.F2StopFrequency,
// 				ResolutionBandwidth: tc.ResolutionBandwidth,
// 				TuneStepSize:        tc.TuneStepSize,
// 				NumberPoints:        tc.NumberPoints,
// 				DefaultGain:         tc.DefaultGain,
// 				Threshold:           tc.Threshold,
// 				DwellTime:           tc.DwellTime,
// 				StartTime:           tc.StartTime,
// 				Time3:               tc.Time3,
// 				Time4:               tc.Time4,
// 			}

// 			// Pack
// 			packed := s.Pack()
// 			assert.Equal(t, tc.expected, packed)

// 			// Unpack
// 			s.Unpack(packed)
// 			assert.Equal(t, tc.SectorNumber, s.SectorNumber)
// 			assert.Equal(t, tc.F1StartFrequency, s.F1StartFrequency)
// 			assert.Equal(t, tc.F2StopFrequency, s.F2StopFrequency)
// 			assert.Equal(t, tc.ResolutionBandwidth, s.ResolutionBandwidth)
// 			assert.Equal(t, tc.TuneStepSize, s.TuneStepSize)
// 			assert.Equal(t, tc.NumberPoints, s.NumberPoints)
// 			assert.Equal(t, tc.DefaultGain, s.DefaultGain)
// 			assert.Equal(t, tc.Threshold, s.Threshold)
// 			assert.Equal(t, tc.DwellTime, s.DwellTime)
// 			assert.Equal(t, tc.StartTime, s.StartTime)
// 			assert.Equal(t, tc.Time3, s.Time3)
// 			assert.Equal(t, tc.Time4, s.Time4)
// 		})
// 	}
// }

// func TestSectorStepScanBytes(t *testing.T) {
// 	s := SectorStepScan{}
// 	assert.Equal(t, sectorStepScanBytes, s.Size())
// }

// func TestSectorStepScanDefault(t *testing.T) {
// 	s := SectorStepScan{}
// 	assert.Equal(t, uint32(0), s.ArraySize)
// 	assert.Equal(t, uint8(0), s.HeaderSize)
// 	assert.Equal(t, uint16(0), s.NumWordsRecord)
// 	assert.Equal(t, uint16(0), s.NumRecords)
// 	assert.Equal(t, SectorStepScanCIF{}, s.SubfieldCif)
// 	assert.Empty(t, s.Records)

// 	// Pack
// 	packed := s.Pack()
// 	expected := make([]byte, 80) // Adjust based on expected size
// 	assert.Equal(t, expected, packed)

// 	// Unpack
// 	s.Unpack(packed)
// 	assert.Equal(t, uint32(0), s.ArraySize)
// 	assert.Equal(t, uint8(0), s.HeaderSize)
// 	assert.Equal(t, uint16(0), s.NumWordsRecord)
// 	assert.Equal(t, uint16(0), s.NumRecords)
// 	assert.Equal(t, SectorStepScanCIF{}, s.SubfieldCif)
// 	assert.Empty(t, s.Records)
// }

// func TestSectorStepScan(t *testing.T) {
// 	cif := SectorStepScanCIF{
// 		SectorNumber:        true,
// 		F1StartFrequency:    true,
// 		F2StartFrequency:    false,
// 		ResolutionBandwidth: false,
// 		TuneStepSize:        false,
// 		NumberPoints:        false,
// 		DefaultGain:         false,
// 		Threshold:           false,
// 		DwellTime:           false,
// 		StartTime:           false,
// 		Time3:               false,
// 		Time4:               false,
// 	}

// 	records := []SectorStepScanRecord{
// 		{
// 			SectorNumber:        1,
// 			F1StartFrequency:    1000,
// 			F2StopFrequency:     0,
// 			ResolutionBandwidth: 0,
// 			TuneStepSize:        0,
// 			NumberPoints:        0,
// 			DefaultGain:         Gain{Stage1: 0, Stage2: 0}, // Replace with actual Gain initialization
// 			Threshold:           Gain{Stage1: 0, Stage2: 0}, // Replace with actual Gain initialization
// 			DwellTime:           0,
// 			StartTime:           0,
// 			Time3:               0,
// 			Time4:               0,
// 		},
// 	}

// 	cases := []struct {
// 		name           string
// 		arraySize      uint32
// 		headerSize     uint8
// 		numWordsRecord uint16
// 		numRecords     uint16
// 		subfieldCif    SectorStepScanCIF
// 		records        []SectorStepScanRecord
// 		expected       []byte
// 	}{
// 		{
// 			name:           "Single Record",
// 			arraySize:      1,
// 			headerSize:     4,
// 			numWordsRecord: 8,
// 			numRecords:     1,
// 			subfieldCif:    cif,
// 			records:        records,
// 			expected:       []byte{0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x80, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			s := SectorStepScan{
// 				ArraySize:      tc.arraySize,
// 				HeaderSize:     tc.headerSize,
// 				NumWordsRecord: tc.numWordsRecord,
// 				NumRecords:     tc.numRecords,
// 				SubfieldCif:    tc.subfieldCif,
// 				Records:        tc.records,
// 			}

// 			// Pack
// 			packed := s.Pack()
// 			assert.Equal(t, tc.expected, packed)

// 			// Unpack
// 			s.Unpack(packed)

// 			assert.Equal(t, tc.arraySize, s.ArraySize)
// 			assert.Equal(t, tc.headerSize, s.HeaderSize)
// 			assert.Equal(t, tc.numWordsRecord, s.NumWordsRecord)
// 			assert.Equal(t, tc.numRecords, s.NumRecords)
// 			assert.Equal(t, tc.subfieldCif, s.SubfieldCif)
// 			assert.Equal(t, tc.records, s.Records)
// 		})
// 	}
// }

func TestVersionInformationBytes(t *testing.T) {
	s := VersionInformation{}
	assert.Equal(t, versionInformationBytes, s.Size())
}

func TestVersionInformationDefault(t *testing.T) {
	s := VersionInformation{}
	assert.Equal(t, uint8(0), s.Year)
	assert.Equal(t, uint16(0), s.Day)
	assert.Equal(t, uint8(0), s.Revision)
	assert.Equal(t, uint16(0), s.UserDefined)

	// Pack
	packed := s.Pack()
	expected := []byte{0x0, 0x0, 0x0, 0x0} // All bits are zero
	assert.Equal(t, expected, packed)

	// Unpack
	s.Unpack(packed)
	assert.Equal(t, uint8(0), s.Year)
	assert.Equal(t, uint16(0), s.Day)
	assert.Equal(t, uint8(0), s.Revision)
	assert.Equal(t, uint16(0), s.UserDefined)
}

func TestVersionInformation(t *testing.T) {
	cases := []struct {
		name        string
		Year        uint8
		Day         uint16
		Revision    uint8
		UserDefined uint16
		expected    []byte
	}{
		{
			name:        "Rule 9.10.4-1/1",
			Year:        127,
			Day:         511,
			Revision:    1,
			UserDefined: 1023,
			expected:    []byte{0xFF, 0xFF, 0x7, 0xFF},
		},
		{
			name:        "Rule 9.10.4-1/2",
			Year:        63,
			Day:         510,
			Revision:    31,
			UserDefined: 511,
			expected:    []byte{0x7F, 0xfe, 0x7D, 0xFF},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := VersionInformation{}
			s.Year = tc.Year
			s.Day = tc.Day
			s.Revision = tc.Revision
			s.UserDefined = tc.UserDefined

			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			s.Unpack(packed)
			assert.Equal(t, tc.Year, s.Year)
			assert.Equal(t, tc.Day, s.Day)
			assert.Equal(t, tc.Revision, s.Revision)
			assert.Equal(t, tc.UserDefined, s.UserDefined)
		})
	}
}
