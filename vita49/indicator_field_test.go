package vita49

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func packAndUnpack(t *testing.T, i *IndicatorField0, v uint32, b uint8) {
	// Pack
	packed := []byte{0xFF, 0xFF, 0xFF, 0xFF}
	i.Pack(packed)
	expected := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(expected, v<<b)
	assert.Equal(t, expected, packed)
	// Unpack
	i.Unpack(packed)
}

func TestIndicatorFieldSize(t *testing.T) {
	indField := IndicatorField0{}
	assert.Equal(t, indField.Size(), indicatorFieldBytes)
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

func TestReferencePointId(t *testing.T) {
	// Default is false
	indField := IndicatorField0{}
	assert.Equal(t, false, indField.ReferencePointId)
	packAndUnpack(t, &indField, 0, 30)
	assert.Equal(t, false, indField.ReferencePointId)
	// Set true
	indField.ReferencePointId = true
	assert.Equal(t, true, indField.ReferencePointId)
	packAndUnpack(t, &indField, 1, 30)
	assert.Equal(t, true, indField.ReferencePointId)
	// Set back to false
	indField.ReferencePointId = false
	assert.Equal(t, false, indField.ReferencePointId)
	packAndUnpack(t, &indField, 0, 30)
	assert.Equal(t, false, indField.ReferencePointId)
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
