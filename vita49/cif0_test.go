package vita49

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCif0Size(t *testing.T) {
	c := Cif0{}
	assert.Equal(t, indicatorFieldBytes, c.Size())
}

func TestGainSize(t *testing.T) {
	g := Gain{}
	assert.Equal(t, gainBytes, g.Size())
}

func TestGainDefault(t *testing.T) {
	g := Gain{}
	assert.Equal(t, float64(0), g.Stage1)
	assert.Equal(t, float64(0), g.Stage2)
	// Pack
	packed := g.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	g.Unpack(packed)
	assert.Equal(t, float64(0), g.Stage1)
	assert.Equal(t, float64(0), g.Stage2)
}

func TestGain(t *testing.T) {
	cases := []struct {
		name     string
		stage1   float64
		stage2   float64
		expected []byte
	}{
		{
			name:     "Rule 9.5.3-3",
			stage1:   1.0,
			stage2:   1.0,
			expected: []byte{0, 0x80, 0, 0x80},
		},
		{
			name:     "Rule 9.5.3-4/1",
			stage1:   1.0,
			stage2:   0.0,
			expected: []byte{0, 0, 0, 0x80},
		},
		{
			name:     "Rule 9.5.3-4/2",
			stage1:   -1.0,
			stage2:   0.0,
			expected: []byte{0, 0, 0xFF, 0x80},
		},
		{
			name:     "Rule 9.5.3-4/3",
			stage1:   0.0078125,
			stage2:   0.0,
			expected: []byte{0, 0, 0, 0x01},
		},
		{
			name:     "Rule 9.5.3-4/4",
			stage1:   -0.0078125,
			stage2:   0.0,
			expected: []byte{0, 0, 0xFF, 0xFF},
		},
		{
			name:     "Rule 9.5.3-5/1",
			stage1:   1.0,
			stage2:   1.0,
			expected: []byte{0, 0x80, 0, 0x80},
		},
		{
			name:     "Rule 9.5.3-5/2",
			stage1:   -1.0,
			stage2:   -1.0,
			expected: []byte{0xFF, 0x80, 0xFF, 0x80},
		},
		{
			name:     "Rule 9.5.3-5/3",
			stage1:   0.0078125,
			stage2:   0.0078125,
			expected: []byte{0, 0x01, 0, 0x01},
		},
		{
			name:     "Rule 9.5.3-5/4",
			stage1:   -0.0078125,
			stage2:   -0.0078125,
			expected: []byte{0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			name:     "Rule 9.5.3-6",
			stage1:   1.0,
			stage2:   0.0,
			expected: []byte{0, 0, 0, 0x80},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Gain{}
			g.Stage1 = tc.stage1
			g.Stage2 = tc.stage2
			assert.Equal(t, tc.stage1, g.Stage1)
			assert.Equal(t, tc.stage2, g.Stage2)
			// Pack
			packed := g.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			g.Unpack(packed)
			assert.Equal(t, tc.stage1, g.Stage1)
			assert.Equal(t, tc.stage2, g.Stage2)
		})
	}
}

func TestDeviceIdentifierSize(t *testing.T) {
	d := DeviceIdentifier{}
	assert.Equal(t, deviceIdentifierBytes, d.Size())
}

func TestDeviceIdentifierDefault(t *testing.T) {
	d := DeviceIdentifier{}
	assert.Equal(t, uint32(0), d.ManufacturerOui)
	assert.Equal(t, uint16(0), d.DeviceCode)
	// Pack
	packed := d.Pack()
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	d.Unpack(packed)
	assert.Equal(t, uint32(0), d.ManufacturerOui)
	assert.Equal(t, uint16(0), d.DeviceCode)
}

func TestDeviceIdentifier(t *testing.T) {
	cases := []struct {
		name            string
		manufacturerOui uint32
		deviceCode      uint16
		expected        []byte
	}{
		{
			name:            "Rule 9.10.1-1",
			manufacturerOui: 0xABCDEF,
			deviceCode:      0x123,
			expected:        []byte{0, 0xAB, 0xCD, 0xEF, 0, 0, 0x01, 0x23},
		},
		{
			name:            "Rule 9.10.1-2",
			manufacturerOui: 0x12AB34,
			deviceCode:      0,
			expected:        []byte{0, 0x12, 0xAB, 0x34, 0, 0, 0, 0},
		},
		{
			name:            "Rule 9.10.1-3",
			manufacturerOui: 0,
			deviceCode:      0xABC,
			expected:        []byte{0, 0, 0, 0, 0, 0, 0x0A, 0xBC},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := DeviceIdentifier{}
			d.ManufacturerOui = tc.manufacturerOui
			d.DeviceCode = tc.deviceCode
			assert.Equal(t, tc.manufacturerOui, d.ManufacturerOui)
			assert.Equal(t, tc.deviceCode, d.DeviceCode)
			// Pack
			packed := d.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			d.Unpack(packed)
			assert.Equal(t, tc.manufacturerOui, d.ManufacturerOui)
			assert.Equal(t, tc.deviceCode, d.DeviceCode)
		})
	}
}

// Ephemeris

func TestEphemerisSize(t *testing.T) {
	e := Ephemeris{}
	assert.Equal(t, ephemerisBytes, e.Size())
}

func TestEphemerisDefault(t *testing.T) {
	// Construct
	e := NewEphemeris()
	assert.Equal(t, uint8(0), e.Tsi)
	assert.Equal(t, uint8(0), e.Tsf)
	assert.Equal(t, uint32(0), e.ManufacturerOui)
	assert.Equal(t, uint32(^uint32(0)), e.IntegerTimestamp)
	assert.Equal(t, uint64(^uint64(0)), e.FractionalTimestamp)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), e.PositionX)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), e.PositionY)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), e.PositionZ)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), e.AttitudeAlpha)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), e.AttitudeBeta)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), e.AttitudePhi)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), e.VelocityDx)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), e.VelocityDy)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), e.VelocityDz)
	// Pack
	packed := e.Pack()

	word1 := make([]byte, 4)
	tsBytes := bytes.Repeat([]byte{0xFF}, 12)
	fields := make([]byte, 4)
	binary.BigEndian.PutUint32(fields, 0x7FFFFFFF)
	expected := append(append(word1, tsBytes...), bytes.Repeat(fields, 9)...)

	assert.Equal(t, expected, packed)
	// Unpack
	e.Unpack(packed)
	assert.Equal(t, uint8(0), e.Tsi)
	assert.Equal(t, uint8(0), e.Tsf)
	assert.Equal(t, uint32(0), e.ManufacturerOui)
	assert.Equal(t, uint32(^uint32(0)), e.IntegerTimestamp)
	assert.Equal(t, uint64(^uint64(0)), e.FractionalTimestamp)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), e.PositionX)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), e.PositionY)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), e.PositionZ)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), e.AttitudeAlpha)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), e.AttitudeBeta)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), e.AttitudePhi)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), e.VelocityDx)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), e.VelocityDy)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), e.VelocityDz)
}

func TestEphemeris(t *testing.T) {
	cases := []struct {
		name                string
		offset              uint16
		tsi                 uint8
		tsf                 uint8
		manufactureroui     uint32
		integertimestamp    uint32
		fractionaltimestamp uint64
		positionx           float64
		positiony           float64
		positionz           float64
		attitudealpha       float64
		attitudebeta        float64
		attitudephi         float64
		velocitydx          float64
		velocitydy          float64
		velocitydz          float64
		expected            []byte
	}{
		{
			name:     "Rule 9.4.3-4",
			tsi:      3,
			expected: []byte{0x0C},
			offset:   0,
		},
		{
			name:     "Rule 9.4.3-4",
			tsf:      3,
			expected: []byte{0x03},
			offset:   0,
		},
		{
			name:            "Rule 9.4.3-4",
			manufactureroui: 0xF11111,
			expected:        []byte{0, 0xF1, 0x11, 0x11},
			offset:          0,
		},
		{
			name:             "Rule 9.4.3-4",
			integertimestamp: 0xFFFFFFFF,
			expected:         []byte{0xFF, 0xFF, 0xFF, 0xFF},
			offset:           4,
		},
		{
			name:                "Rule 9.4.3-4",
			fractionaltimestamp: 0xFFFFFFFFFFFFFFFF,
			expected:            bytes.Repeat([]byte{0xFF}, 8),
			offset:              8,
		},
		{
			name:      "Rule 9.4.3-5 Smallest Fraction",
			positionx: -0.03125,
			expected:  []byte{0xFF, 0xFF, 0xFF, 0xFF},
			offset:    16,
		},
		{
			name:      "Rule 9.4.3-5 X=1",
			positionx: 1,
			expected:  []byte{0, 0, 0, 0x20},
			offset:    16,
		},
		{
			name:      "Rule 9.4.3-5 Y=1",
			positiony: 1,
			expected:  []byte{0, 0, 0, 0x20},
			offset:    20,
		},
		{
			name:      "Rule 9.4.3-5 Z=1",
			positionz: 1,
			expected:  []byte{0, 0, 0, 0x20},
			offset:    24,
		},
		{
			name:          "Rule 9.4.3-6 alpha=1",
			attitudealpha: 1,
			expected:      []byte{0, 0x40, 0, 0},
			offset:        28,
		},
		{
			name:         "Rule 9.4.3-6 beta=1",
			attitudebeta: 1,
			expected:     []byte{0, 0x40, 0, 0},
			offset:       32,
		},
		{
			name:        "Rule 9.4.3-6 phi=1",
			attitudephi: 1,
			expected:    []byte{0, 0x40, 0, 0},
			offset:      36,
		},
		{
			name:       "Rule 9.4.3-7 dx=1",
			velocitydx: 1,
			expected:   []byte{0, 0x01, 0, 0},
			offset:     40,
		},
		{
			name:       "Rule 9.4.3-7 dy=1",
			velocitydy: 1,
			expected:   []byte{0, 0x01, 0, 0},
			offset:     44,
		},
		{
			name:       "Rule 9.4.3-7 dz=1",
			velocitydz: 1,
			expected:   []byte{0, 0x01, 0, 0},
			offset:     48,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := NewEphemeris()
			e.Tsi = tc.tsi
			e.Tsf = tc.tsf
			e.ManufacturerOui = tc.manufactureroui
			e.IntegerTimestamp = tc.integertimestamp
			e.FractionalTimestamp = tc.fractionaltimestamp
			e.PositionX = tc.positionx
			e.PositionY = tc.positiony
			e.PositionZ = tc.positionz
			e.AttitudeAlpha = tc.attitudealpha
			e.AttitudeBeta = tc.attitudebeta
			e.AttitudePhi = tc.attitudephi
			e.VelocityDx = tc.velocitydx
			e.VelocityDy = tc.velocitydy
			e.VelocityDz = tc.velocitydz

			assert.Equal(t, tc.tsi, e.Tsi)
			assert.Equal(t, tc.tsf, e.Tsf)
			assert.Equal(t, tc.manufactureroui, e.ManufacturerOui)
			assert.Equal(t, tc.integertimestamp, e.IntegerTimestamp)
			assert.Equal(t, tc.fractionaltimestamp, e.FractionalTimestamp)
			assert.Equal(t, tc.positionx, e.PositionX)
			assert.Equal(t, tc.positiony, e.PositionY)
			assert.Equal(t, tc.positionz, e.PositionZ)
			assert.Equal(t, tc.attitudealpha, e.AttitudeAlpha)
			assert.Equal(t, tc.attitudebeta, e.AttitudeBeta)
			assert.Equal(t, tc.attitudephi, e.AttitudePhi)
			assert.Equal(t, tc.velocitydx, e.VelocityDx)
			assert.Equal(t, tc.velocitydy, e.VelocityDy)
			assert.Equal(t, tc.velocitydz, e.VelocityDz)

			// Pack
			packed := e.Pack()
			testSlice := make([]byte, 52)
			copy(testSlice[tc.offset:], tc.expected)
			assert.Equal(t, testSlice, packed)
			// Unpack
			e.Unpack(packed)
			assert.Equal(t, tc.tsi, e.Tsi)
			assert.Equal(t, tc.tsf, e.Tsf)
			assert.Equal(t, tc.manufactureroui, e.ManufacturerOui)
			assert.Equal(t, tc.integertimestamp, e.IntegerTimestamp)
			assert.Equal(t, tc.fractionaltimestamp, e.FractionalTimestamp)
			assert.Equal(t, tc.positionx, e.PositionX)
			assert.Equal(t, tc.positiony, e.PositionY)
			assert.Equal(t, tc.positionz, e.PositionZ)
			assert.Equal(t, tc.attitudealpha, e.AttitudeAlpha)
			assert.Equal(t, tc.attitudebeta, e.AttitudeBeta)
			assert.Equal(t, tc.attitudephi, e.AttitudePhi)
			assert.Equal(t, tc.velocitydx, e.VelocityDx)
			assert.Equal(t, tc.velocitydy, e.VelocityDy)
			assert.Equal(t, tc.velocitydz, e.VelocityDz)

		})
	}

}

// Geolocation
func TestGeolocationSize(t *testing.T) {
	g := Geolocation{}
	assert.Equal(t, geolocationBytes, g.Size())
}

func TestGeolocationDefault(t *testing.T) {
	// Construct
	g := NewGeolocation()
	assert.Equal(t, uint8(0), g.Tsi)
	assert.Equal(t, uint8(0), g.Tsf)
	assert.Equal(t, uint32(0), g.ManufacturerOui)
	assert.Equal(t, uint32(^uint32(0)), g.IntegerTimestamp)
	assert.Equal(t, uint64(^uint64(0)), g.FractionalTimestamp)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.Latitude)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.Longitude)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), g.Altitude)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), g.SpeedOverGround)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.HeadingAngle)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.TrackAngle)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.MagneticVariation)
	// Pack
	packed := g.Pack()
	word1 := make([]byte, 4)
	tsBytes := bytes.Repeat([]byte{0xFF}, 12)
	fields := make([]byte, 4)
	binary.BigEndian.PutUint32(fields, 0x7FFFFFFF)
	expected := append(append(word1, tsBytes...), bytes.Repeat(fields, 7)...)
	assert.Equal(t, expected, packed)
	// Unpack
	g.Unpack(packed)
	assert.Equal(t, uint8(0), g.Tsi)
	assert.Equal(t, uint8(0), g.Tsf)
	assert.Equal(t, uint32(0), g.ManufacturerOui)
	assert.Equal(t, uint32(^uint32(0)), g.IntegerTimestamp)
	assert.Equal(t, uint64(^uint64(0)), g.FractionalTimestamp)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.Latitude)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.Longitude)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 5), g.Altitude)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 16), g.SpeedOverGround)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.HeadingAngle)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.TrackAngle)
	assert.Equal(t, FromFixed(int32(0x7FFFFFFF), 22), g.MagneticVariation)
}

func TestGeolocation(t *testing.T) {
	cases := []struct {
		name                string
		tsi                 uint8
		tsf                 uint8
		manufactureroui     uint32
		integertimestamp    uint32
		fractionaltimestamp uint64
		latitude            float64
		longitude           float64
		altitude            float64
		speedoverground     float64
		headingangle        float64
		trackangle          float64
		magneticvariation   float64
		offset              uint16
		expected            []byte
	}{
		{
			name:            "Rule 9.4.5-2",
			manufactureroui: 0xFFFFFF,
			expected:        []byte{0, 0xFF, 0xFF, 0xFF},
			offset:          0,
		},
		{
			name:     "Rule 9.4.5-3",
			tsi:      3,
			expected: []byte{0x0C, 0, 0, 0},
			offset:   0,
		},
		{
			name:     "Rule 9.4.5-4",
			tsf:      3,
			expected: []byte{0x03, 0, 0, 0},
			offset:   0,
		},

		{
			name:             "Rule 9.4.5-6",
			integertimestamp: 0xFFFFFFFF,
			expected:         []byte{0xFF, 0xFF, 0xFF, 0xFF},
			offset:           4,
		},
		{
			name:                "Rule 9.4.5-6",
			fractionaltimestamp: 0xFFFFFFFFFFFFFFFF,
			expected:            bytes.Repeat([]byte{0xFF}, 8),
			offset:              8,
		},
		{
			name:     "Rule 9.4.5-7",
			latitude: 1,
			expected: []byte{0, 0x40, 0, 0},
			offset:   16,
		},
		{
			name:      "Rule 9.4.5-7",
			longitude: 1,
			expected:  []byte{0, 0x40, 0, 0},
			offset:    20,
		},
		{
			name:     "Rule 9.4.5-10",
			altitude: 1,
			expected: []byte{0, 0, 0, 0x20},
			offset:   24,
		},
		{
			name:            "Rule 9.4.5-11",
			speedoverground: 1,
			expected:        []byte{0, 0x01, 0, 0},
			offset:          28,
		},
		{
			name:         "Rule 9.4.5-12",
			headingangle: 1,
			expected:     []byte{0, 0x40, 0, 0},
			offset:       32,
		},
		{
			name:       "Rule 9.4.5-14",
			trackangle: 1,
			expected:   []byte{0, 0x40, 0, 0},
			offset:     36,
		},
		{
			name:              "Rule 9.4.5-16",
			magneticvariation: 1,
			expected:          []byte{0, 0x40, 0, 0},
			offset:            40,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGeolocation()
			g.Tsi = tc.tsi
			g.Tsf = tc.tsf
			g.ManufacturerOui = tc.manufactureroui
			g.IntegerTimestamp = tc.integertimestamp
			g.FractionalTimestamp = tc.fractionaltimestamp
			g.Latitude = tc.latitude
			g.Longitude = tc.longitude
			g.Altitude = tc.altitude
			g.SpeedOverGround = tc.speedoverground
			g.HeadingAngle = tc.headingangle
			g.TrackAngle = tc.trackangle
			g.MagneticVariation = tc.magneticvariation
			assert.Equal(t, tc.tsi, g.Tsi)
			assert.Equal(t, tc.tsf, g.Tsf)
			assert.Equal(t, tc.manufactureroui, g.ManufacturerOui)
			assert.Equal(t, tc.integertimestamp, g.IntegerTimestamp)
			assert.Equal(t, tc.fractionaltimestamp, g.FractionalTimestamp)
			assert.Equal(t, tc.latitude, g.Latitude)
			assert.Equal(t, tc.longitude, g.Longitude)
			assert.Equal(t, tc.altitude, g.Altitude)
			assert.Equal(t, tc.speedoverground, g.SpeedOverGround)
			assert.Equal(t, tc.headingangle, g.HeadingAngle)
			assert.Equal(t, tc.trackangle, g.TrackAngle)
			assert.Equal(t, tc.magneticvariation, g.MagneticVariation)
			// Pack
			packed := g.Pack()
			testSlice := make([]byte, 44)
			copy(testSlice[tc.offset:], tc.expected)
			assert.Equal(t, testSlice, packed)

			// Unpack
			g.Unpack(packed)
			assert.Equal(t, tc.tsi, g.Tsi)
			assert.Equal(t, tc.tsf, g.Tsf)
			assert.Equal(t, tc.manufactureroui, g.ManufacturerOui)
			assert.Equal(t, tc.integertimestamp, g.IntegerTimestamp)
			assert.Equal(t, tc.fractionaltimestamp, g.FractionalTimestamp)
			assert.Equal(t, tc.latitude, g.Latitude)
			assert.Equal(t, tc.longitude, g.Longitude)
			assert.Equal(t, tc.altitude, g.Altitude)
			assert.Equal(t, tc.speedoverground, g.SpeedOverGround)
			assert.Equal(t, tc.headingangle, g.HeadingAngle)
			assert.Equal(t, tc.trackangle, g.TrackAngle)
			assert.Equal(t, tc.magneticvariation, g.MagneticVariation)
		})
	}

}

// GPS ASCII
func TestGpsAsciiSize(t *testing.T) {
	g := GpsAscii{}
	assert.Equal(t, uint32(8), g.Size())
}

func TestGpsAsciiDefault(t *testing.T) {
	g := NewGpsAscii()
	assert.Equal(t, uint32(0), g.ManufacturerOui)
	assert.Equal(t, uint32(0), g.NumberOfWords)
	assert.Equal(t, []uint8{}, g.AsciiSentences)

	// Pack
	packed := g.Pack()
	expected := make([]byte, 8)
	assert.Equal(t, expected, packed)
	// Unpack
	g.Unpack(packed)
	assert.Equal(t, uint32(0), g.ManufacturerOui)
	assert.Equal(t, uint32(0), g.NumberOfWords)
	assert.Equal(t, []uint8{}, g.AsciiSentences)
}

func TestGpsAscii(t *testing.T) {
	cases := []struct {
		name            string
		manufactureroui uint32
		numberofwords   uint32
		asciisentences  []uint8
		offset          uint16
		expected        []byte
	}{
		{
			name:            "Rule 9.4.7-2",
			manufactureroui: 0x00FFFFFF,
			asciisentences:  []byte{},
			expected:        []byte{0, 0xFF, 0xFF, 0xFF, 0, 0, 0, 0},
			offset:          0,
		},

		{
			name:           "Rule 9.4.7-3",
			numberofwords:  1,
			asciisentences: []uint8{0xFF, 0xFF, 0xFF, 0xFF},
			expected:       []byte{0, 0, 0, 0x01, 0xFF, 0xFF, 0xFF, 0xFF},
			offset:         4,
		},
		{
			name:           "Rule 9.4.7-5",
			numberofwords:  3,
			asciisentences: []uint8{0xFF, 0xFF, 0xFF, 0xFF},
			expected:       append([]byte{0, 0, 0, 0x03, 0xFF, 0xFF, 0xFF, 0xFF}, make([]byte, 8)...),
			offset:         4,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGpsAscii()
			g.ManufacturerOui = tc.manufactureroui
			g.NumberOfWords = tc.numberofwords
			g.AsciiSentences = tc.asciisentences
			assert.Equal(t, tc.manufactureroui, g.ManufacturerOui)
			assert.Equal(t, tc.numberofwords, g.NumberOfWords)
			assert.Equal(t, tc.asciisentences, g.AsciiSentences)
			// Pack
			packed := g.Pack()
			testSlice := make([]byte, int((g.NumberOfWords+2)*4))
			copy(testSlice[tc.offset:], tc.expected)
			assert.Equal(t, testSlice, packed)
			// Unpack
			g.Unpack(packed)
			assert.Equal(t, tc.manufactureroui, g.ManufacturerOui)
			assert.Equal(t, tc.numberofwords, g.NumberOfWords)
			assert.Equal(t, tc.asciisentences, g.AsciiSentences[0:len(tc.asciisentences)])
		})
	}

}

// Payload Format

func TestPayloadFormatSize(t *testing.T) {
	p := PayloadFormat{}
	assert.Equal(t, payloadFormatBytes, p.Size())
}

func TestPayloadFormatDefault(t *testing.T) {
	p := NewPayloadFormat()
	assert.Equal(t, false, p.PackingMethod)
	assert.Equal(t, uint8(0), p.RealComplexType)
	assert.Equal(t, uint8(0), p.DataItemFormat)
	assert.Equal(t, false, p.RepeatIndicator)
	assert.Equal(t, uint8(0), p.EventTagSize)
	assert.Equal(t, uint8(0), p.ChannelTagSize)
	assert.Equal(t, uint8(0), p.DataItemFractionSize)
	assert.Equal(t, uint8(1), p.ItemPackingFieldSize)
	assert.Equal(t, uint8(1), p.DataItemSize)
	assert.Equal(t, uint32(1), p.RepeatCount)
	assert.Equal(t, uint32(1), p.VectorSize)

	// Pack
	packed := p.Pack()
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	p.Unpack(packed)
	assert.Equal(t, false, p.PackingMethod)
	assert.Equal(t, uint8(0), p.RealComplexType)
	assert.Equal(t, uint8(0), p.DataItemFormat)
	assert.Equal(t, false, p.RepeatIndicator)
	assert.Equal(t, uint8(0), p.EventTagSize)
	assert.Equal(t, uint8(0), p.ChannelTagSize)
	assert.Equal(t, uint8(0), p.DataItemFractionSize)
	assert.Equal(t, uint8(1), p.ItemPackingFieldSize)
	assert.Equal(t, uint8(1), p.DataItemSize)
	assert.Equal(t, uint32(1), p.RepeatCount)
	assert.Equal(t, uint32(1), p.VectorSize)
}

func TestPayloadFormat(t *testing.T) {
	cases := []struct {
		name                 string
		packingmethod        bool
		realcomplextype      uint8
		dataitemformat       uint8
		repeatindicator      bool
		eventtagsize         uint8
		channeltagsize       uint8
		dataitemfractionsize uint8
		itempackingfieldsize uint8
		dataitemsize         uint8
		repeatcount          uint32
		vectorsize           uint32
		expected             []byte
	}{
		{
			name:          "Rule 9.13.3-2 false",
			packingmethod: false,
			expected:      []byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:          "Rule 9.13.3-2 true",
			packingmethod: true,
			expected:      []byte{0x80, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:            "Rule 9.13.3-3",
			realcomplextype: 3,
			expected:        []byte{0x60, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:           "Rule 9.13.3-4",
			dataitemformat: 31,
			expected:       []byte{0x1F, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:            "Rule 9.13.3-5",
			repeatindicator: false,
			expected:        []byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:            "Rule 9.13.3-5",
			repeatindicator: true,
			expected:        []byte{0, 0x80, 0, 0, 0, 0, 0, 0},
		},
		{
			name:         "Rule 9.13.3-6",
			eventtagsize: 7,
			expected:     []byte{0, 0x70, 0, 0, 0, 0, 0, 0},
		},
		{
			name:           "Rule 9.13.3-7",
			channeltagsize: 15,
			expected:       []byte{0, 0x0F, 0, 0, 0, 0, 0, 0},
		},
		{
			name:                 "Rule 9.13.3-8",
			dataitemfractionsize: 15,
			expected:             []byte{0, 0, 0xF0, 0, 0, 0, 0, 0},
		},
		{
			name:                 "Rule 9.13.3-12",
			itempackingfieldsize: 0x40,
			expected:             []byte{0, 0, 0x0F, 0xC0, 0, 0, 0, 0},
		},
		{
			name:         "Rule 9.13.3-13",
			dataitemsize: 0x40,
			expected:     []byte{0, 0, 0, 0x3F, 0, 0, 0, 0},
		},
		{
			name:        "Rule 9.13.3-14",
			repeatcount: 0x10000,
			expected:    []byte{0, 0, 0, 0, 0xFF, 0xFF, 0, 0},
		},
		{
			name:       "Rule 9.13.3-15",
			vectorsize: 0x10000,
			expected:   []byte{0, 0, 0, 0, 0, 0, 0xFF, 0xFF},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := PayloadFormat{}
			p.PackingMethod = tc.packingmethod
			p.RealComplexType = tc.realcomplextype
			p.DataItemFormat = tc.dataitemformat
			p.RepeatIndicator = tc.repeatindicator
			p.EventTagSize = tc.eventtagsize
			p.ChannelTagSize = tc.channeltagsize
			p.DataItemFractionSize = tc.dataitemfractionsize
			p.ItemPackingFieldSize = tc.itempackingfieldsize
			p.DataItemSize = tc.dataitemsize
			p.RepeatCount = tc.repeatcount
			p.VectorSize = tc.vectorsize

			assert.Equal(t, tc.packingmethod, p.PackingMethod)
			assert.Equal(t, tc.realcomplextype, p.RealComplexType)
			assert.Equal(t, tc.dataitemformat, p.DataItemFormat)
			assert.Equal(t, tc.repeatindicator, p.RepeatIndicator)
			assert.Equal(t, tc.eventtagsize, p.EventTagSize)
			assert.Equal(t, tc.channeltagsize, p.ChannelTagSize)
			assert.Equal(t, tc.dataitemfractionsize, p.DataItemFractionSize)
			assert.Equal(t, tc.itempackingfieldsize, p.ItemPackingFieldSize)
			assert.Equal(t, tc.dataitemsize, p.DataItemSize)
			assert.Equal(t, tc.repeatcount, p.RepeatCount)
			assert.Equal(t, tc.vectorsize, p.VectorSize)
			// Pack
			packed := p.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			p.Unpack(packed)
			assert.Equal(t, tc.packingmethod, p.PackingMethod)
			assert.Equal(t, tc.realcomplextype, p.RealComplexType)
			assert.Equal(t, tc.dataitemformat, p.DataItemFormat)
			assert.Equal(t, tc.repeatindicator, p.RepeatIndicator)
			assert.Equal(t, tc.eventtagsize, p.EventTagSize)
			assert.Equal(t, tc.channeltagsize, p.ChannelTagSize)
			assert.Equal(t, tc.dataitemfractionsize, p.DataItemFractionSize)

			if tc.vectorsize < 0xFFFF {
				assert.Equal(t, 1+tc.vectorsize, p.VectorSize)
			} else {
				assert.Equal(t, tc.vectorsize, p.VectorSize)
			}

			if tc.repeatcount < 0xFFFF {
				assert.Equal(t, 1+tc.repeatcount, p.RepeatCount)
			} else {
				assert.Equal(t, tc.repeatcount, p.RepeatCount)
			}

			if tc.itempackingfieldsize < 64 {
				assert.Equal(t, 1+tc.itempackingfieldsize, p.ItemPackingFieldSize)
			} else {
				assert.Equal(t, tc.itempackingfieldsize, p.ItemPackingFieldSize)
			}

			if tc.dataitemsize < 64 {
				assert.Equal(t, 1+tc.dataitemsize, p.DataItemSize)
			} else {
				assert.Equal(t, tc.dataitemsize, p.DataItemSize)
			}

		})
	}

}

// Context Association Lists
func TestContextAssociationListsSize(t *testing.T) {
	c := ContextAssociationLists{}
	assert.Equal(t, uint32(8), c.Size())
}

func TestContextAssociationListsDefault(t *testing.T) {
	c := NewContextAssociationLists()
	assert.Equal(t, uint8(0), c.SourceListSize)
	assert.Equal(t, uint8(0), c.SystemListSize)
	assert.Equal(t, uint16(0), c.VectorListSize)
	assert.Equal(t, false, c.AsyncTagListEnable)
	assert.Equal(t, uint16(0), c.AsyncListSize)
	assert.Equal(t, []uint32{}, c.SourceList)
	assert.Equal(t, []uint32{}, c.SystemList)
	assert.Equal(t, []uint32{}, c.VectorList)
	assert.Equal(t, []uint32{}, c.AsyncList)
	assert.Equal(t, []uint32{}, c.AsyncTagList)

	// Pack
	packed := c.Pack()
	expected := make([]byte, 8)
	assert.Equal(t, expected, packed)
	// Unpack
	c.Unpack(packed)
	assert.Equal(t, uint8(0), c.SourceListSize)
	assert.Equal(t, uint8(0), c.SystemListSize)
	assert.Equal(t, uint16(0), c.VectorListSize)
	assert.Equal(t, false, c.AsyncTagListEnable)
	assert.Equal(t, uint16(0), c.AsyncListSize)
	assert.Equal(t, []uint32{}, c.SourceList)
	assert.Equal(t, []uint32{}, c.SystemList)
	assert.Equal(t, []uint32{}, c.VectorList)
	assert.Equal(t, []uint32{}, c.AsyncList)
	assert.Equal(t, []uint32{}, c.AsyncTagList)
}

func TestContextAssociationLists(t *testing.T) {
	cases := []struct {
		name               string
		sourceListSize     uint8
		systemListSize     uint8
		vectorListSize     uint16
		asyncTagListEnable bool
		asyncListSize      uint16
		sourceList         []uint32
		systemList         []uint32
		vectorList         []uint32
		asyncList          []uint32
		asyncTagList       []uint32
		offset             uint16
		expected           []byte
	}{
		{
			name:           "Rule 9.13.2-1",
			sourceListSize: 1,
			sourceList:     []uint32{0xFFFFFFFF},
			systemList:     []uint32{},
			vectorList:     []uint32{},
			asyncList:      []uint32{},
			asyncTagList:   []uint32{},
			expected:       append([]byte{0, 0x01, 0, 0, 0, 0, 0, 0}, bytes.Repeat([]byte{0xFF}, 4)...),
			offset:         0,
		},
		{
			name:           "Rule 9.13.2-1",
			sourceListSize: 2,
			systemList:     []uint32{},
			vectorList:     []uint32{},
			asyncList:      []uint32{},
			asyncTagList:   []uint32{},
			sourceList:     []uint32{0xFFFFFFFF, 0xFFFFFFFF},
			expected:       append([]byte{0, 0x02, 0, 0, 0, 0, 0, 0}, bytes.Repeat([]byte{0xFF}, 8)...),
		},

		{
			name:           "Rule 9.13.2-2",
			systemListSize: 1,
			sourceList:     []uint32{},
			vectorList:     []uint32{},
			asyncList:      []uint32{},
			asyncTagList:   []uint32{},
			systemList:     []uint32{0xFFFFFFFF},
			expected:       append([]byte{0, 0, 0, 0x01, 0, 0, 0, 0}, bytes.Repeat([]byte{0xFF}, 4)...),
		},
		{
			name:           "Rule 9.13.2-2",
			sourceList:     []uint32{},
			vectorList:     []uint32{},
			asyncList:      []uint32{},
			asyncTagList:   []uint32{},
			systemListSize: 2,
			systemList:     []uint32{0xFFFFFFFF, 0xFFFFFFFF},
			expected:       append([]byte{0, 0, 0, 0x02, 0, 0, 0, 0}, bytes.Repeat([]byte{0xFF}, 8)...),
		},

		{
			name:           "Rule 9.13.2-3",
			sourceList:     []uint32{},
			systemList:     []uint32{},
			asyncList:      []uint32{},
			asyncTagList:   []uint32{},
			vectorListSize: 1,
			vectorList:     []uint32{0xFFFFFFFF},
			expected:       append([]byte{0, 0, 0, 0, 0, 0x01, 0, 0}, bytes.Repeat([]byte{0xFF}, 4)...),
		},
		{
			name:           "Rule 9.13.2-3",
			sourceList:     []uint32{},
			systemList:     []uint32{},
			asyncList:      []uint32{},
			asyncTagList:   []uint32{},
			vectorListSize: 2,
			vectorList:     []uint32{0xFFFFFFFF, 0xFFFFFFFF},
			expected:       append([]byte{0, 0, 0, 0, 0, 0x02, 0, 0}, bytes.Repeat([]byte{0xFF}, 8)...),
		},

		{
			sourceList:    []uint32{},
			systemList:    []uint32{},
			vectorList:    []uint32{},
			asyncTagList:  []uint32{},
			name:          "Rule 9.13.2-4",
			asyncListSize: 1,
			asyncList:     []uint32{0xFFFFFFFF},
			expected:      append([]byte{0, 0, 0, 0, 0, 0, 0, 0x01}, bytes.Repeat([]byte{0xFF}, 4)...),
		},
		{
			name:          "Rule 9.13.2-4",
			sourceList:    []uint32{},
			systemList:    []uint32{},
			vectorList:    []uint32{},
			asyncTagList:  []uint32{},
			asyncListSize: 2,
			asyncList:     []uint32{0xFFFFFFFF, 0xFFFFFFFF},
			expected:      append([]byte{0, 0, 0, 0, 0, 0, 0, 0x02}, bytes.Repeat([]byte{0xFF}, 8)...),
		},

		{
			name:               "Rule 9.13.2-5",
			sourceList:         []uint32{},
			systemList:         []uint32{},
			vectorList:         []uint32{},
			asyncListSize:      1,
			asyncList:          []uint32{0xFFFFFFFF},
			asyncTagListEnable: true,
			asyncTagList:       []uint32{0xFFFFFFFF},
			expected:           append([]byte{0, 0, 0, 0, 0, 0, 0x80, 0x01}, bytes.Repeat([]byte{0xFF}, 8)...),
		},
		{
			name:               "Rule 9.13.2-5",
			sourceList:         []uint32{},
			systemList:         []uint32{},
			vectorList:         []uint32{},
			asyncListSize:      2,
			asyncList:          []uint32{0xFFFFFFFF, 0xFFFFFFFF},
			asyncTagListEnable: true,
			asyncTagList:       []uint32{0xFFFFFFFF, 0xFFFFFFFF},
			expected:           append([]byte{0, 0, 0, 0, 0, 0, 0x80, 0x02}, bytes.Repeat([]byte{0xFF}, 16)...),
		},
		{
			name:               "Rule 9.13.2-6",
			sourceList:         []uint32{},
			systemList:         []uint32{},
			vectorList:         []uint32{},
			asyncListSize:      1,
			asyncList:          []uint32{0xFFFFFFFF},
			asyncTagListEnable: false,
			asyncTagList:       []uint32{0xFFFFFFFF},
			expected:           append([]byte{0, 0, 0, 0, 0, 0, 0, 0x01}, bytes.Repeat([]byte{0xFF}, 4)...),
		},
		{
			name:               "Rule 9.13.2-7",
			sourceListSize:     0,
			systemListSize:     0,
			vectorListSize:     0,
			asyncListSize:      0,
			asyncTagListEnable: false,
			sourceList:         []uint32{0xFFFFFFFF},
			systemList:         []uint32{0xFFFFFFFF},
			vectorList:         []uint32{0xFFFFFFFF},
			asyncList:          []uint32{0xFFFFFFFF},
			asyncTagList:       []uint32{0xFFFFFFFF},
			expected:           []byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := ContextAssociationLists{}

			c.SourceListSize = tc.sourceListSize
			c.SystemListSize = tc.systemListSize
			c.VectorListSize = tc.vectorListSize
			c.AsyncTagListEnable = tc.asyncTagListEnable
			c.AsyncListSize = tc.asyncListSize
			c.SourceList = tc.sourceList
			c.SystemList = tc.systemList
			c.VectorList = tc.vectorList
			c.AsyncList = tc.asyncList
			c.AsyncTagList = tc.asyncTagList

			assert.Equal(t, tc.sourceListSize, c.SourceListSize)
			assert.Equal(t, tc.systemListSize, c.SystemListSize)
			assert.Equal(t, tc.vectorListSize, c.VectorListSize)
			assert.Equal(t, tc.asyncTagListEnable, c.AsyncTagListEnable)
			assert.Equal(t, tc.asyncListSize, c.AsyncListSize)
			assert.Equal(t, tc.sourceList, c.SourceList)
			assert.Equal(t, tc.systemList, c.SystemList)
			assert.Equal(t, tc.vectorList, c.VectorList)
			assert.Equal(t, tc.asyncList, c.AsyncList)
			assert.Equal(t, tc.asyncTagList, c.AsyncTagList)

			// Pack
			totalSize := (2 +
				int(tc.sourceListSize) +
				int(tc.systemListSize) +
				int(tc.vectorListSize) +
				int(tc.asyncListSize)) * 4

			if tc.asyncTagListEnable {
				totalSize += int(tc.asyncListSize * 4)
			}

			packed := c.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			c.Unpack(packed)
			assert.Equal(t, tc.sourceListSize, c.SourceListSize)
			assert.Equal(t, tc.systemListSize, c.SystemListSize)
			assert.Equal(t, tc.vectorListSize, c.VectorListSize)
			assert.Equal(t, tc.asyncTagListEnable, c.AsyncTagListEnable)
			assert.Equal(t, tc.asyncListSize, c.AsyncListSize)
			// Ignore these for the last check
			if tc.sourceListSize != 0 || tc.systemListSize != 0 ||
				tc.vectorListSize != 0 ||
				tc.asyncListSize != 0 {
				assert.Equal(t, tc.sourceList, c.SourceList)
				assert.Equal(t, tc.systemList, c.SystemList)
				assert.Equal(t, tc.vectorList, c.VectorList)
				assert.Equal(t, tc.asyncList, c.AsyncList)
			}
			if tc.asyncTagListEnable {
				assert.Equal(t, tc.asyncTagList, c.AsyncTagList)
			}

		})
	}

}
