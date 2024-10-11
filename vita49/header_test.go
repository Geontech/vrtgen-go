package vita49

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderSize(t *testing.T) {
	h := Header{}
	assert.Equal(t, h.Size(), headerBytes)
}

func TestHeaderDefault(t *testing.T) {
	h := Header{}
	assert.Equal(t, SignalData, h.PacketType)
	assert.Equal(t, false, h.ClassIdEnable)
	assert.EqualValues(t, NoneTsi, h.Tsi)
	assert.EqualValues(t, NoneTsf, h.Tsf)
	assert.EqualValues(t, 0, h.PacketCount)
	assert.EqualValues(t, 0, h.PacketSize)
	// Pack
	packed := make([]byte, 4)
	h.Pack(packed)
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	h.Unpack(packed)
	assert.Equal(t, SignalData, h.PacketType)
	assert.Equal(t, false, h.ClassIdEnable)
	assert.EqualValues(t, NoneTsi, h.Tsi)
	assert.EqualValues(t, NoneTsf, h.Tsf)
	assert.EqualValues(t, 0, h.PacketCount)
	assert.EqualValues(t, 0, h.PacketSize)
}

func TestHeaderPacketType(t *testing.T) {
	cases := []struct {
		name       string
		packetType PacketType
	}{
		{
			name:       "SignalData",
			packetType: SignalData,
		},
		{
			name:       "SignalDataStreamId",
			packetType: SignalDataStreamId,
		},
		{
			name:       "ExtensionData",
			packetType: ExtensionData,
		},
		{
			name:       "ExtensionDataStreamId",
			packetType: ExtensionDataStreamId,
		},
		{
			name:       "Context",
			packetType: Context,
		},
		{
			name:       "ExtensionContext",
			packetType: ExtensionContext,
		},
		{
			name:       "Command",
			packetType: Command,
		},
		{
			name:       "ExtensionCommand",
			packetType: ExtensionCommand,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := Header{}
			// Assign
			h.PacketType = tc.packetType
			assert.Equal(t, tc.packetType, h.PacketType)
			// Pack
			packed := make([]byte, 4)
			h.Pack(packed)
			expected := []byte{uint8(tc.packetType) << 4, 0, 0, 0}
			assert.Equal(t, expected, packed)
			// Unpack
			h.Unpack(packed)
			assert.Equal(t, tc.packetType, h.PacketType)
		})
	}
}

func TestHeaderTsi(t *testing.T) {
	cases := []struct {
		name string
		tsi  Tsi
	}{
		{
			name: "None",
			tsi:  NoneTsi,
		},
		{
			name: "Utc",
			tsi:  Utc,
		},
		{
			name: "Gps",
			tsi:  Gps,
		},
		{
			name: "Other",
			tsi:  Other,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := Header{}
			// Assign
			h.Tsi = tc.tsi
			assert.Equal(t, tc.tsi, h.Tsi)
			// Pack
			packed := make([]byte, 4)
			h.Pack(packed)
			expected := []byte{0, uint8(tc.tsi) << 6, 0, 0}
			assert.Equal(t, expected, packed)
			// Unpack
			h.Unpack(packed)
			assert.Equal(t, tc.tsi, h.Tsi)
		})
	}
}

func TestHeaderTsf(t *testing.T) {
	cases := []struct {
		name string
		tsf  Tsf
	}{
		{
			name: "None",
			tsf:  NoneTsf,
		},
		{
			name: "SampleCount",
			tsf:  SampleCount,
		},
		{
			name: "Picoseconds",
			tsf:  Picoseconds,
		},
		{
			name: "Other",
			tsf:  FreeRunning,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := Header{}
			// Assign
			h.Tsf = tc.tsf
			assert.Equal(t, tc.tsf, h.Tsf)
			// Pack
			packed := make([]byte, 4)
			h.Pack(packed)
			expected := []byte{0, uint8(tc.tsf) << 4, 0, 0}
			assert.Equal(t, expected, packed)
			// Unpack
			h.Unpack(packed)
			assert.Equal(t, tc.tsf, h.Tsf)
		})
	}
}

func TestHeaderPacketCount(t *testing.T) {
	cases := []struct {
		name  string
		count uint8
	}{
		{
			name:  "zero",
			count: 0,
		},
		{
			name:  "one",
			count: 1,
		},
		{
			name:  "three",
			count: 3,
		},
		{
			name:  "seven",
			count: 7,
		},
		{
			name:  "thirteen",
			count: 13,
		},
		{
			name:  "fourty-two",
			count: 42,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := Header{}
			// Assign
			h.PacketCount = tc.count
			assert.Equal(t, tc.count, h.PacketCount)
			// Pack
			packed := make([]byte, 4)
			h.Pack(packed)
			expected := []byte{0, (tc.count % 16), 0, 0}
			assert.Equal(t, expected, packed)
			// Unpack
			h.Unpack(packed)
			assert.Equal(t, (tc.count % 16), h.PacketCount)
		})
	}
}

func TestHeaderPacketSize(t *testing.T) {
	cases := []struct {
		name string
		size uint16
	}{
		{
			name: "zero",
			size: 0,
		},
		{
			name: "fourty-two",
			size: 42,
		},
		{
			name: "max",
			size: 0xFFFF,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := Header{}
			// Assign
			h.PacketSize = tc.size
			assert.Equal(t, tc.size, h.PacketSize)
			// Pack
			packed := make([]byte, 4)
			h.Pack(packed)
			expected := []byte{0, 0, 0, 0}
			binary.BigEndian.PutUint16(expected[2:], tc.size)
			assert.Equal(t, expected, packed)
			// Unpack
			h.Unpack(packed)
			assert.Equal(t, tc.size, h.PacketSize)
		})
	}
}
