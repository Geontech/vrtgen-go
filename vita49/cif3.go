package vita49

import (
	"encoding/binary"
)

const (
	TimestampDetailsBytes = uint32(8)
	SeaSwellStateBytes    = uint32(4)
)

type TimestampDetails struct {
	UserDefined           uint8
	Global                bool
	TseCode               uint8
	LshCode               uint8
	LspCode               uint8
	TimeSource            uint8
	EnablePosixTimeOffset bool
	PosixTimeOffset       uint8
	TimestampEpoch        uint32
}

func (t TimestampDetails) Size() uint32 {
	return TimestampDetailsBytes
}

func (t *TimestampDetails) Pack() []byte {
	buf := make([]byte, t.Size())
	word1 := uint32(0)
	word1 |= uint32(t.UserDefined) << 24
	if t.Global {
		word1 |= uint32(1) << 18
	}
	word1 |= uint32(t.TseCode) << 16
	word1 |= uint32(t.LshCode) << 14
	word1 |= uint32(t.LspCode) << 12
	word1 |= uint32(t.TimeSource) << 9
	if t.EnablePosixTimeOffset {
		word1 |= uint32(1) << 8
		word1 |= uint32(t.PosixTimeOffset)
	}
	binary.BigEndian.PutUint32(buf, word1)
	binary.BigEndian.PutUint32(buf[4:], t.TimestampEpoch)
	return buf
}

func (t *TimestampDetails) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf)
	t.UserDefined = uint8(word1 >> 24)
	t.Global = ((word1 >> 18) & 1) != 0
	t.TseCode = uint8((word1 >> 16) & 3)
	t.LshCode = uint8((word1 >> 14) & 3)
	t.LspCode = uint8((word1 >> 12) & 3)
	t.TimeSource = uint8((word1 >> 9) & 7)
	t.EnablePosixTimeOffset = ((word1 >> 8) & 1) != 0
	t.PosixTimeOffset = uint8(word1)
	t.TimestampEpoch = binary.BigEndian.Uint32(buf[4:])
}

type SeaSwellState struct {
	UserDefined uint8
	SwellState  uint8
	SeaState    uint8
}

func (s SeaSwellState) Size() uint32 {
	return SeaSwellStateBytes
}

func (s *SeaSwellState) Pack() []byte {
	buf := make([]byte, s.Size())
	word := uint32(0)
	word |= uint32(s.UserDefined) << 10
	word |= uint32(s.SwellState) << 5
	word |= uint32(s.SeaState)
	binary.BigEndian.PutUint32(buf, word)
	return buf
}

func (s *SeaSwellState) Unpack(buf []byte) {
	word := binary.BigEndian.Uint32(buf)
	s.UserDefined = uint8(word>>10) & 63
	s.SwellState = uint8(word>>5) & 31
	s.SeaState = uint8(word) & 31
}
