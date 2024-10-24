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
	"reflect"
	"testing"
)

func TestStateEventIndicatorsDefault(t *testing.T) {
	sei := StateEventIndicators{}
	packed := sei.Pack()
	expected := []byte{0, 0, 0, 0}
	if !reflect.DeepEqual(packed, expected) {
		t.Errorf("Default packed bits = %v; want %v", packed, expected)
	}
}

func TestStateEventIndicatorsCalibratedTime(t *testing.T) {
	cases := []struct {
		name   string
		enable bool
		value  bool
	}{
		{
			name:   "FalseFalse",
			enable: false,
			value:  false,
		},
		{
			name:   "TrueFalse",
			enable: true,
			value:  false,
		},
		{
			name:   "FalseTrue",
			enable: false,
			value:  true,
		},
		{
			name:   "TrueTrue",
			enable: true,
			value:  true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var enableVal uint8
			if tc.enable {
				enableVal = 1
			}
			var valueVal uint8
			if tc.enable && tc.value {
				valueVal = 1
			}
			sei := StateEventIndicators{}
			// Assign
			sei.CalibratedTime.Enable = tc.enable
			if sei.CalibratedTime.Enable != tc.enable {
				t.Errorf("CalibratedTime.Enable = %t; want %t", sei.CalibratedTime.Enable, tc.enable)
			}
			sei.CalibratedTime.Value = tc.value
			if sei.CalibratedTime.Value != tc.value {
				t.Errorf("CalibratedTime.Value = %t; want %t", sei.CalibratedTime.Value, tc.value)
			}
			// Pack
			packed := sei.Pack()
			expected := []byte{enableVal << 7, valueVal << 3, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			sei.Unpack(packed)
			if sei.CalibratedTime.Enable != tc.enable {
				t.Errorf("CalibratedTime.Enable = %t; want %t", sei.CalibratedTime.Enable, tc.enable)
			}
			if tc.enable {
				if sei.CalibratedTime.Value != tc.value {
					t.Errorf("CalibratedTime.Value = %t; want %t", sei.CalibratedTime.Value, tc.value)
				}
			} else if sei.CalibratedTime.Value != false {
				t.Errorf("CalibratedTime.Value = %t; want %t", sei.CalibratedTime.Value, false)
			}
		})
	}
}

func TestStateEventIndicatorsValidData(t *testing.T) {
	cases := []struct {
		name   string
		enable bool
		value  bool
	}{
		{
			name:   "FalseFalse",
			enable: false,
			value:  false,
		},
		{
			name:   "TrueFalse",
			enable: true,
			value:  false,
		},
		{
			name:   "FalseTrue",
			enable: false,
			value:  true,
		},
		{
			name:   "TrueTrue",
			enable: true,
			value:  true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var enableVal uint8
			if tc.enable {
				enableVal = 1
			}
			var valueVal uint8
			if tc.enable && tc.value {
				valueVal = 1
			}
			sei := StateEventIndicators{}
			// Assign
			sei.ValidData.Enable = tc.enable
			if sei.ValidData.Enable != tc.enable {
				t.Errorf("ValidData.Enable = %t; want %t", sei.ValidData.Enable, tc.enable)
			}
			sei.ValidData.Value = tc.value
			if sei.ValidData.Value != tc.value {
				t.Errorf("ValidData.Value = %t; want %t", sei.ValidData.Value, tc.value)
			}
			// Pack
			packed := sei.Pack()
			expected := []byte{enableVal << 6, valueVal << 2, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			sei.Unpack(packed)
			if sei.ValidData.Enable != tc.enable {
				t.Errorf("ValidData.Enable = %t; want %t", sei.ValidData.Enable, tc.enable)
			}
			if tc.enable {
				if sei.ValidData.Value != tc.value {
					t.Errorf("ValidData.Value = %t; want %t", sei.ValidData.Value, tc.value)
				}
			} else if sei.ValidData.Value != false {
				t.Errorf("ValidData.Value = %t; want %t", sei.ValidData.Value, false)
			}
		})
	}
}

func TestStateEventIndicatorsReferenceLock(t *testing.T) {
	cases := []struct {
		name   string
		enable bool
		value  bool
	}{
		{
			name:   "FalseFalse",
			enable: false,
			value:  false,
		},
		{
			name:   "TrueFalse",
			enable: true,
			value:  false,
		},
		{
			name:   "FalseTrue",
			enable: false,
			value:  true,
		},
		{
			name:   "TrueTrue",
			enable: true,
			value:  true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var enableVal uint8
			if tc.enable {
				enableVal = 1
			}
			var valueVal uint8
			if tc.enable && tc.value {
				valueVal = 1
			}
			sei := StateEventIndicators{}
			// Assign
			sei.ReferenceLock.Enable = tc.enable
			if sei.ReferenceLock.Enable != tc.enable {
				t.Errorf("ReferenceLock.Enable = %t; want %t", sei.ReferenceLock.Enable, tc.enable)
			}
			sei.ReferenceLock.Value = tc.value
			if sei.ReferenceLock.Value != tc.value {
				t.Errorf("ReferenceLock.Value = %t; want %t", sei.ReferenceLock.Value, tc.value)
			}
			// Pack
			packed := sei.Pack()
			expected := []byte{enableVal << 5, valueVal << 1, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			sei.Unpack(packed)
			if sei.ReferenceLock.Enable != tc.enable {
				t.Errorf("ReferenceLock.Enable = %t; want %t", sei.ReferenceLock.Enable, tc.enable)
			}
			if tc.enable {
				if sei.ReferenceLock.Value != tc.value {
					t.Errorf("ReferenceLock.Value = %t; want %t", sei.ReferenceLock.Value, tc.value)
				}
			} else if sei.ReferenceLock.Value != false {
				t.Errorf("ReferenceLock.Value = %t; want %t", sei.ReferenceLock.Value, false)
			}
		})
	}
}

func TestStateEventIndicatorsAgcMgc(t *testing.T) {
	cases := []struct {
		name   string
		enable bool
		value  bool
	}{
		{
			name:   "FalseFalse",
			enable: false,
			value:  false,
		},
		{
			name:   "TrueFalse",
			enable: true,
			value:  false,
		},
		{
			name:   "FalseTrue",
			enable: false,
			value:  true,
		},
		{
			name:   "TrueTrue",
			enable: true,
			value:  true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var enableVal uint8
			if tc.enable {
				enableVal = 1
			}
			var valueVal uint8
			if tc.enable && tc.value {
				valueVal = 1
			}
			sei := StateEventIndicators{}
			// Assign
			sei.AgcMgc.Enable = tc.enable
			if sei.AgcMgc.Enable != tc.enable {
				t.Errorf("AgcMgc.Enable = %t; want %t", sei.AgcMgc.Enable, tc.enable)
			}
			sei.AgcMgc.Value = tc.value
			if sei.AgcMgc.Value != tc.value {
				t.Errorf("AgcMgc.Value = %t; want %t", sei.AgcMgc.Value, tc.value)
			}
			// Pack
			packed := sei.Pack()
			expected := []byte{enableVal << 4, valueVal, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			sei.Unpack(packed)
			if sei.AgcMgc.Enable != tc.enable {
				t.Errorf("AgcMgc.Enable = %t; want %t", sei.AgcMgc.Enable, tc.enable)
			}
			if tc.enable {
				if sei.AgcMgc.Value != tc.value {
					t.Errorf("AgcMgc.Value = %t; want %t", sei.AgcMgc.Value, tc.value)
				}
			} else if sei.AgcMgc.Value != false {
				t.Errorf("AgcMgc.Value = %t; want %t", sei.AgcMgc.Value, false)
			}
		})
	}
}
