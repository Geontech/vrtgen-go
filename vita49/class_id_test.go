package vita49

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassIdSize(t *testing.T) {
	d := ClassId{}
	assert.Equal(t, classIdBytes, d.Size())
}

func TestClassIdDefault(t *testing.T) {
	d := ClassId{}
	assert.Equal(t, uint8(0), d.PadBitCount)
	assert.Equal(t, uint32(0), d.Oui)
	assert.Equal(t, uint16(0), d.InformationCode)
	assert.Equal(t, uint16(0), d.PacketCode)

	// Pack
	packed := d.Pack()
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0} // Adjust based on the packing logic
	assert.Equal(t, expected, packed)

	// Unpack
	d.Unpack(packed)
	assert.Equal(t, uint8(0), d.PadBitCount)
	assert.Equal(t, uint32(0), d.Oui)
	assert.Equal(t, uint16(0), d.InformationCode)
	assert.Equal(t, uint16(0), d.PacketCode)
}

func TestClassId(t *testing.T) {
	cases := []struct {
		name            string
		padBitCount     uint8
		oui             uint32
		informationCode uint16
		packetCode      uint16
		expected        []byte
	}{
		{
			name:        "Rule 5.1.3-4",
			padBitCount: 0x1F,
			expected:    []byte{0xF8, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Rule 5.1.3-1",
			oui:      0x00FFFFFF,
			expected: []byte{0, 0xFF, 0xFF, 0xFF, 0, 0, 0, 0},
		},
		{
			name:            "Rule 5.1.3-2",
			informationCode: 0xFFFF,
			expected:        []byte{0, 0, 0, 0, 0xFF, 0xFF, 0, 0},
		},
		{
			name:       "Rule 5.1.3-3",
			packetCode: 0xFFFF,
			expected:   []byte{0, 0, 0, 0, 0, 0, 0xFF, 0xFF},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := ClassId{}
			d.PadBitCount = tc.padBitCount
			d.Oui = tc.oui
			d.InformationCode = tc.informationCode
			d.PacketCode = tc.packetCode

			assert.Equal(t, tc.padBitCount, d.PadBitCount)
			assert.Equal(t, tc.oui, d.Oui)
			assert.Equal(t, tc.informationCode, d.InformationCode)
			assert.Equal(t, tc.packetCode, d.PacketCode)

			// Pack
			packed := d.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			d.Unpack(packed)
			assert.Equal(t, tc.padBitCount, d.PadBitCount)
			assert.Equal(t, tc.oui, d.Oui)
			assert.Equal(t, tc.informationCode, d.InformationCode)
			assert.Equal(t, tc.packetCode, d.PacketCode)
		})
	}
}
