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

import "encoding/binary"

const (
	classIdBytes = uint32(8)
)

type ClassID struct {
	PadBitCount     uint8
	Oui             uint32 `yaml:"oui,omitempty"`
	InformationCode uint16 `yaml:"informationCode,omitempty"`
	PacketCode      uint16 `yaml:"packetCode,omitempty"`
}

func (c *ClassID) Size() uint32 {
	return classIdBytes
}

func (c *ClassID) Pack() []byte {
	buf := make([]byte, c.Size())
	word1 := (uint32(c.PadBitCount&0x1F) << 27) | (c.Oui & 0x00FFFFFF)
	binary.BigEndian.PutUint32(buf[0:], word1)
	binary.BigEndian.PutUint16(buf[4:], c.InformationCode)
	binary.BigEndian.PutUint16(buf[6:], c.PacketCode)
	return buf
}

func (c *ClassID) Unpack(buf []byte) {
	c.PadBitCount = uint8(binary.BigEndian.Uint32(buf[0:]) >> 27)
	c.Oui = binary.BigEndian.Uint32(buf[0:]) & 0x00FFFFFF
	c.InformationCode = binary.BigEndian.Uint16(buf[4:])
	c.PacketCode = binary.BigEndian.Uint16(buf[6:])
}
