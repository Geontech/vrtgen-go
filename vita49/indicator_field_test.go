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
	"testing"

	"github.com/stretchr/testify/assert"
)

func packAndUnpack(t *testing.T, i *IndicatorField0, v uint32, b uint8) {
	// Pack
	packed := i.Pack()
	expected := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(expected, v<<b)
	assert.Equal(t, expected, packed)
	// Unpack
	i.Unpack(packed)
}

func TestIndicatorFieldSize(t *testing.T) {
	indField := IndicatorField0{}
	assert.Equal(t, uint32(4), indField.Size())
}

func TestChangeIndicator(t *testing.T) {
	// Default is false
	indField := IndicatorField0{}
	assert.Equal(t, false, indField.ChangeIndicator)
	packAndUnpack(t, &indField, 0, 31)
	assert.Equal(t, false, indField.ChangeIndicator)
	// Set true
	indField.ChangeIndicator = true
	assert.Equal(t, true, indField.ChangeIndicator)
	packAndUnpack(t, &indField, 1, 31)
	assert.Equal(t, true, indField.ChangeIndicator)
	// Set back to false
	indField.ChangeIndicator = false
	assert.Equal(t, false, indField.ChangeIndicator)
	packAndUnpack(t, &indField, 0, 31)
	assert.Equal(t, false, indField.ChangeIndicator)
}

func TestReferencePointID(t *testing.T) {
	// Default is false
	indField := IndicatorField0{}
	assert.Equal(t, false, indField.ReferencePointID)
	packAndUnpack(t, &indField, 0, 30)
	assert.Equal(t, false, indField.ReferencePointID)
	// Set true
	indField.ReferencePointID = true
	assert.Equal(t, true, indField.ReferencePointID)
	packAndUnpack(t, &indField, 1, 30)
	assert.Equal(t, true, indField.ReferencePointID)
	// Set back to false
	indField.ReferencePointID = false
	assert.Equal(t, false, indField.ReferencePointID)
	packAndUnpack(t, &indField, 0, 30)
	assert.Equal(t, false, indField.ReferencePointID)
}

func TestBandwidth(t *testing.T) {
	// Default is false
	indField := IndicatorField0{}
	assert.Equal(t, false, indField.Bandwidth)
	packAndUnpack(t, &indField, 0, 29)
	assert.Equal(t, false, indField.Bandwidth)
	// Set true
	indField.Bandwidth = true
	assert.Equal(t, true, indField.Bandwidth)
	packAndUnpack(t, &indField, 1, 29)
	assert.Equal(t, true, indField.Bandwidth)
	// Set back to false
	indField.Bandwidth = false
	assert.Equal(t, false, indField.Bandwidth)
	packAndUnpack(t, &indField, 0, 29)
	assert.Equal(t, false, indField.Bandwidth)
}

func TestIfRefFrequency(t *testing.T) {
	// Default is false
	indField := IndicatorField0{}
	assert.Equal(t, false, indField.IfRefFrequency)
	packAndUnpack(t, &indField, 0, 28)
	assert.Equal(t, false, indField.IfRefFrequency)
	// Set true
	indField.IfRefFrequency = true
	assert.Equal(t, true, indField.IfRefFrequency)
	packAndUnpack(t, &indField, 1, 28)
	assert.Equal(t, true, indField.IfRefFrequency)
	// Set back to false
	indField.IfRefFrequency = false
	assert.Equal(t, false, indField.IfRefFrequency)
	packAndUnpack(t, &indField, 0, 28)
	assert.Equal(t, false, indField.IfRefFrequency)
}

func TestRfRefFrequency(t *testing.T) {
	// Default is false
	indField := IndicatorField0{}
	assert.Equal(t, false, indField.RfRefFrequency)
	packAndUnpack(t, &indField, 0, 27)
	assert.Equal(t, false, indField.RfRefFrequency)
	// Set true
	indField.RfRefFrequency = true
	assert.Equal(t, true, indField.RfRefFrequency)
	packAndUnpack(t, &indField, 1, 27)
	assert.Equal(t, true, indField.RfRefFrequency)
	// Set back to false
	indField.RfRefFrequency = false
	assert.Equal(t, false, indField.RfRefFrequency)
	packAndUnpack(t, &indField, 0, 27)
	assert.Equal(t, false, indField.RfRefFrequency)
}
