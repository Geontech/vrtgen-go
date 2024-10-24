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

type EnableIndicator struct {
	Enable bool
	Value  bool
}

func (ei *EnableIndicator) Reset() {
	ei.Enable = false
	ei.Value = false
}

func (ei *EnableIndicator) Pack(buf []byte, enPos uint8, inPos uint8) {
	var enableVal uint8
	var indicatorVal uint8
	if ei.Enable {
		enableVal = 1
		if ei.Value {
			indicatorVal = 1
		}
	}
	bytePos := uint8((31 - enPos) / 8)
	mask := ^(uint8(1) << (enPos % 8))
	buf[bytePos] = (buf[bytePos] & mask) | (enableVal << (enPos % 8))
	bytePos = uint8((31 - inPos) / 8)
	mask = ^(uint8(1) << (inPos % 8))
	buf[bytePos] = (buf[bytePos] & mask) | (indicatorVal << (inPos % 8))
}

func (ei *EnableIndicator) Unpack(buf []byte, enPos uint8, inPos uint8) {
	bytePos := uint8((31 - enPos) / 8)
	ei.Enable = (buf[bytePos] & (1 << (enPos % 8))) != 0
	bytePos = uint8((31 - inPos) / 8)
	ei.Value = (buf[bytePos] & (1 << (inPos % 8))) != 0
}
