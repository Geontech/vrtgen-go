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

const (
	headerBytes = uint32(4)
)

type PacketType uint8

const (
	SignalData PacketType = iota
	SignalDataStreamId
	ExtensionData
	ExtensionDataStreamId
	Context
	ExtensionContext
	Command
	ExtensionCommand
)

type Tsi uint8

const (
	NoneTsi Tsi = iota
	Utc
	Gps
	Other
)

type Tsf uint8

const (
	NoneTsf Tsf = iota
	SampleCount
	Picoseconds
	FreeRunning
)

type Header struct {
	PacketType    PacketType
	ClassIdEnable bool
	Tsi           Tsi
	Tsf           Tsf
	PacketCount   uint8
	PacketSize    uint16
}

func (h Header) Size() uint32 {
	return headerBytes
}

func (h *Header) Pack() []byte {
	buf := make([]byte, h.Size())
	var classIdEnableVal uint8
	if h.ClassIdEnable {
		classIdEnableVal = 1
	}
	buf[0] = (uint8(h.PacketType) << 4) | (classIdEnableVal << 3)
	buf[1] = (uint8(h.Tsi) << 6) | (uint8(h.Tsf) << 4) | (h.PacketCount % 16)
	binary.BigEndian.PutUint16(buf[2:], uint16(h.PacketSize))
	return buf
}

func (h *Header) Unpack(buf []byte) {
	h.PacketType = PacketType((buf[0] & 0xF0) >> 4)
	h.ClassIdEnable = (buf[0] & 0x08) != 0
	h.Tsi = Tsi((buf[1] & 0xC0) >> 6)
	h.Tsf = Tsf((buf[1] & 0x30) >> 4)
	h.PacketCount = buf[1] & 0x0F
	h.PacketSize = binary.BigEndian.Uint16(buf[2:])
}

type DataHeader struct {
	Header
	TrailerIncluded bool
	NotV490         bool
	Spectrum        bool
}

func (h *DataHeader) Pack() []byte {
	buf := h.Header.Pack()
	if h.TrailerIncluded {
		buf[0] |= (uint8(1) << 2)
	}
	if h.NotV490 {
		buf[0] |= (uint8(1) << 1)
	}
	if h.Spectrum {
		buf[0] |= uint8(1)
	}
	return buf
}

func (h *DataHeader) Unpack(buf []byte) {
	h.Header.Unpack(buf)
	h.TrailerIncluded = (buf[0] & 0x04) != 0
	h.NotV490 = (buf[0] & 0x02) != 0
	h.Spectrum = (buf[0] & 0x01) != 0
}

type Tsm uint8

const (
	Fine Tsm = iota
	Coarse
)

type ContextHeader struct {
	Header
	NotV490 bool
	Tsm     Tsm
}

func (h *ContextHeader) Pack() []byte {
	buf := h.Header.Pack()
	if h.NotV490 {
		buf[0] |= (uint8(1) << 1)
	}
	buf[0] |= uint8(h.Tsm)
	return buf
}

func (h *ContextHeader) Unpack(buf []byte) {
	h.Header.Unpack(buf)
	h.NotV490 = (buf[0] & 0x2) != 0
	h.Tsm = Tsm(buf[0] & 0x01)
}

type CommandHeader struct {
	Header
	Acknowledge  bool
	Cancellation bool
}

func (h *CommandHeader) Pack() []byte {
	buf := h.Header.Pack()
	if h.Acknowledge {
		buf[0] |= (uint8(1) << 2)
	}
	if h.Cancellation {
		buf[0] |= uint8(1)
	}
	return buf
}

func (h *CommandHeader) Unpack(buf []byte) {
	h.Header.Unpack(buf)
	h.Acknowledge = (buf[0] & 0x4) != 0
	h.Cancellation = (buf[0] & 0x1) != 0
}
