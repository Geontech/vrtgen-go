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
	return 8
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
	return 4
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
