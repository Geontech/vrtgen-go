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

func TestTrailerDefault(t *testing.T) {
	trailer := Trailer{}
	packed := trailer.Pack()
	expected := []byte{0, 0, 0, 0}
	if !reflect.DeepEqual(packed, expected) {
		t.Errorf("Default packed bits = %v; want %v", packed, expected)
	}
}

func TestTrailerCalibratedTime(t *testing.T) {
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
			trailer := Trailer{}
			// Assign
			trailer.CalibratedTime.Enable = tc.enable
			if trailer.CalibratedTime.Enable != tc.enable {
				t.Errorf("CalibratedTime.Enable = %t; want %t", trailer.CalibratedTime.Enable, tc.enable)
			}
			trailer.CalibratedTime.Value = tc.value
			if trailer.CalibratedTime.Value != tc.value {
				t.Errorf("CalibratedTime.Value = %t; want %t", trailer.CalibratedTime.Value, tc.value)
			}
			// Pack
			packed := trailer.Pack()
			expected := []byte{enableVal << 7, valueVal << 3, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			trailer.Unpack(packed)
			if trailer.CalibratedTime.Enable != tc.enable {
				t.Errorf("CalibratedTime.Enable = %t; want %t", trailer.CalibratedTime.Enable, tc.enable)
			}
			if tc.enable {
				if trailer.CalibratedTime.Value != tc.value {
					t.Errorf("CalibratedTime.Value = %t; want %t", trailer.CalibratedTime.Value, tc.value)
				}
			} else if trailer.CalibratedTime.Value != false {
				t.Errorf("CalibratedTime.Value = %t; want %t", trailer.CalibratedTime.Value, false)
			}
		})
	}
}

func TestTrailerValidData(t *testing.T) {
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
			trailer := Trailer{}
			// Assign
			trailer.ValidData.Enable = tc.enable
			if trailer.ValidData.Enable != tc.enable {
				t.Errorf("ValidData.Enable = %t; want %t", trailer.ValidData.Enable, tc.enable)
			}
			trailer.ValidData.Value = tc.value
			if trailer.ValidData.Value != tc.value {
				t.Errorf("ValidData.Value = %t; want %t", trailer.ValidData.Value, tc.value)
			}
			// Pack
			packed := trailer.Pack()
			expected := []byte{enableVal << 6, valueVal << 2, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			trailer.Unpack(packed)
			if trailer.ValidData.Enable != tc.enable {
				t.Errorf("ValidData.Enable = %t; want %t", trailer.ValidData.Enable, tc.enable)
			}
			if tc.enable {
				if trailer.ValidData.Value != tc.value {
					t.Errorf("ValidData.Value = %t; want %t", trailer.ValidData.Value, tc.value)
				}
			} else if trailer.ValidData.Value != false {
				t.Errorf("ValidData.Value = %t; want %t", trailer.ValidData.Value, false)
			}
		})
	}
}

func TestTrailerReferenceLock(t *testing.T) {
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
			trailer := Trailer{}
			// Assign
			trailer.ReferenceLock.Enable = tc.enable
			if trailer.ReferenceLock.Enable != tc.enable {
				t.Errorf("ReferenceLock.Enable = %t; want %t", trailer.ReferenceLock.Enable, tc.enable)
			}
			trailer.ReferenceLock.Value = tc.value
			if trailer.ReferenceLock.Value != tc.value {
				t.Errorf("ReferenceLock.Value = %t; want %t", trailer.ReferenceLock.Value, tc.value)
			}
			// Pack
			packed := trailer.Pack()
			expected := []byte{enableVal << 5, valueVal << 1, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			trailer.Unpack(packed)
			if trailer.ReferenceLock.Enable != tc.enable {
				t.Errorf("ReferenceLock.Enable = %t; want %t", trailer.ReferenceLock.Enable, tc.enable)
			}
			if tc.enable {
				if trailer.ReferenceLock.Value != tc.value {
					t.Errorf("ReferenceLock.Value = %t; want %t", trailer.ReferenceLock.Value, tc.value)
				}
			} else if trailer.ReferenceLock.Value != false {
				t.Errorf("ReferenceLock.Value = %t; want %t", trailer.ReferenceLock.Value, false)
			}
		})
	}
}

func TestTrailerAgcMgc(t *testing.T) {
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
			trailer := Trailer{}
			// Assign
			trailer.AgcMgc.Enable = tc.enable
			if trailer.AgcMgc.Enable != tc.enable {
				t.Errorf("AgcMgc.Enable = %t; want %t", trailer.AgcMgc.Enable, tc.enable)
			}
			trailer.AgcMgc.Value = tc.value
			if trailer.AgcMgc.Value != tc.value {
				t.Errorf("AgcMgc.Value = %t; want %t", trailer.AgcMgc.Value, tc.value)
			}
			// Pack
			packed := trailer.Pack()
			expected := []byte{enableVal << 4, valueVal, 0, 0}
			if !reflect.DeepEqual(packed, expected) {
				t.Errorf("Packed bits = %v; want %v", packed, expected)
			}
			// Unpack
			trailer.Unpack(packed)
			if trailer.AgcMgc.Enable != tc.enable {
				t.Errorf("AgcMgc.Enable = %t; want %t", trailer.AgcMgc.Enable, tc.enable)
			}
			if tc.enable {
				if trailer.AgcMgc.Value != tc.value {
					t.Errorf("AgcMgc.Value = %t; want %t", trailer.AgcMgc.Value, tc.value)
				}
			} else if trailer.AgcMgc.Value != false {
				t.Errorf("AgcMgc.Value = %t; want %t", trailer.AgcMgc.Value, false)
			}
		})
	}
}
