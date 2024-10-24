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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCAMSize(t *testing.T) {
	c := CAM{}
	assert.Equal(t, uint32(4), c.Size())
}

func TestCAMDefault(t *testing.T) {
	c := CAM{}
	assert.Equal(t, false, c.ControlleeEnable)
	assert.Equal(t, Word, c.ControlleeFormat)
	assert.Equal(t, false, c.ControllerEnable)
	assert.Equal(t, Word, c.ControllerFormat)
	assert.Equal(t, false, c.PermitPartial)
	assert.Equal(t, false, c.PermitWarnings)
	assert.Equal(t, false, c.PermitErrors)
	assert.Equal(t, NoAction, c.ActionMode)
	assert.Equal(t, false, c.NackOnly)
	assert.Equal(t, TimestampControlMode(0), c.TimingControl)
	// Pack
	packed := c.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	c.Unpack(packed)
	assert.Equal(t, false, c.ControlleeEnable)
	assert.Equal(t, Word, c.ControlleeFormat)
	assert.Equal(t, false, c.ControllerEnable)
	assert.Equal(t, Word, c.ControllerFormat)
	assert.Equal(t, false, c.PermitPartial)
	assert.Equal(t, false, c.PermitWarnings)
	assert.Equal(t, NoAction, c.ActionMode)
	assert.Equal(t, false, c.NackOnly)
	assert.Equal(t, TimestampControlMode(0), c.TimingControl)
}

func TestCAM(t *testing.T) {
	cases := []struct {
		name             string
		controlleeEnable bool
		controlleeFormat IdentifierFormat
		controllerEnable bool
		controllerFormat IdentifierFormat
		permitPartial    bool
		permitWarnings   bool
		permitErrors     bool
		actionMode       ActionMode
		nackOnly         bool
		timingControl    TimestampControlMode
		expected         []byte
	}{
		{
			name:             "controlleeEnable True",
			controlleeEnable: true,
			expected:         []byte{0x80, 0, 0, 0},
		},
		{
			name:             "ControlleeFormat True",
			controlleeFormat: UUID,
			expected:         []byte{0x40, 0, 0, 0}, // Bit 30 set
		},
		{
			name:             "ControllerEnable True",
			controllerEnable: true,
			expected:         []byte{0x20, 0, 0, 0}, // Bit 29 set
		},
		{
			name:             "ControllerFormat True",
			controllerFormat: UUID,
			expected:         []byte{0x10, 0, 0, 0}, // Bit 28 set
		},
		{
			name:          "PermitPartial True",
			permitPartial: true,
			expected:      []byte{0x08, 0, 0, 0}, // Bit 27 set
		},
		{
			name:           "PermitWarnings True",
			permitWarnings: true,
			expected:       []byte{0x04, 0, 0, 0}, // Bit 26 set
		},
		{
			name:         "PermitErrors True",
			permitErrors: true,
			expected:     []byte{0x02, 0, 0, 0}, // Bit 25 set
		},
		{
			name:       "ActionMode Set",
			actionMode: 3,
			expected:   []byte{0x1, 0x80, 0, 0}, // Bits 23-24 set (3 << 23)
		},
		{
			name:     "NackOnly True",
			nackOnly: true,
			expected: []byte{0, 0x40, 0, 0}, // Bit 22 set
		},
		{
			name:          "TimingControl Set",
			timingControl: 7,
			expected:      []byte{0, 0, 0x70, 0}, // Bits 12 set (7 << 12)
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := CAM{
				ControlleeEnable: tc.controlleeEnable,
				ControlleeFormat: tc.controlleeFormat,
				ControllerEnable: tc.controllerEnable,
				ControllerFormat: tc.controllerFormat,
				PermitPartial:    tc.permitPartial,
				PermitWarnings:   tc.permitWarnings,
				PermitErrors:     tc.permitErrors,
				ActionMode:       tc.actionMode,
				NackOnly:         tc.nackOnly,
				TimingControl:    tc.timingControl,
			}
			assert.Equal(t, tc.controlleeEnable, c.ControlleeEnable)
			// Pack
			packed := c.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			c.Unpack(packed)
			assert.Equal(t, tc.controlleeEnable, c.ControlleeEnable)
			assert.Equal(t, tc.controlleeFormat, c.ControlleeFormat)
			assert.Equal(t, tc.controllerEnable, c.ControllerEnable)
			assert.Equal(t, tc.controllerFormat, c.ControllerFormat)
			assert.Equal(t, tc.permitPartial, c.PermitPartial)
			assert.Equal(t, tc.permitWarnings, c.PermitWarnings)
			assert.Equal(t, tc.permitErrors, c.PermitErrors)
			assert.Equal(t, tc.actionMode, c.ActionMode)
			assert.Equal(t, tc.nackOnly, c.NackOnly)
			assert.Equal(t, tc.timingControl, c.TimingControl)
		})
	}
}

func TestControlCAMSize(t *testing.T) {
	c := ControlCAM{}
	assert.Equal(t, uint32(4), c.Size())
}

func TestControlCAMDefault(t *testing.T) {
	c := ControlCAM{}
	assert.Equal(t, false, c.ReqV)
	assert.Equal(t, false, c.ReqX)
	assert.Equal(t, false, c.ReqS)
	assert.Equal(t, false, c.ReqW)
	assert.Equal(t, false, c.ReqEr)
	// Pack
	packed := c.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	c.Unpack(packed)
	assert.Equal(t, false, c.ReqV)
	assert.Equal(t, false, c.ReqX)
	assert.Equal(t, false, c.ReqS)
	assert.Equal(t, false, c.ReqW)
	assert.Equal(t, false, c.ReqEr)
}

func TestControlCAM(t *testing.T) {
	cases := []struct {
		name     string
		reqV     bool
		reqX     bool
		reqS     bool
		reqW     bool
		reqEr    bool
		expected []byte
	}{
		{
			name:     "reqV True",
			reqV:     true,
			expected: []byte{0, 0x10, 0, 0},
		},
		{
			name:     "reqX True",
			reqX:     true,
			expected: []byte{0, 0x08, 0, 0},
		},
		{
			name:     "reqS True",
			reqS:     true,
			expected: []byte{0, 0x04, 0, 0},
		},
		{
			name:     "reqW True",
			reqW:     true,
			expected: []byte{0, 0x02, 0, 0},
		},
		{
			name:     "reqEr True",
			reqEr:    true,
			expected: []byte{0, 0x01, 0, 0},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := ControlCAM{
				ReqV:  tc.reqV,
				ReqX:  tc.reqX,
				ReqS:  tc.reqS,
				ReqW:  tc.reqW,
				ReqEr: tc.reqEr,
			}
			assert.Equal(t, tc.reqV, c.ReqV)
			assert.Equal(t, tc.reqX, c.ReqX)
			assert.Equal(t, tc.reqS, c.ReqS)
			assert.Equal(t, tc.reqW, c.ReqW)
			assert.Equal(t, tc.reqEr, c.ReqEr)
			// Pack
			packed := c.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			c.Unpack(packed)
			assert.Equal(t, tc.reqV, c.ReqV)
			assert.Equal(t, tc.reqX, c.ReqX)
			assert.Equal(t, tc.reqS, c.ReqS)
			assert.Equal(t, tc.reqW, c.ReqW)
			assert.Equal(t, tc.reqEr, c.ReqEr)
		})
	}
}

func TestAcknowledgeCAM(t *testing.T) {
	cases := []struct {
		name                string
		ackV                bool
		ackX                bool
		ackS                bool
		ackW                bool
		ackEr               bool
		partialAction       bool
		scheduledOrExecuted bool
		expected            []byte
	}{
		{
			name:     "ackV True",
			ackV:     true,
			expected: []byte{0, 0x10, 0, 0},
		},
		{
			name:     "ackX True",
			ackX:     true,
			expected: []byte{0, 0x08, 0, 0},
		},
		{
			name:     "ackS True",
			ackS:     true,
			expected: []byte{0, 0x04, 0, 0},
		},
		{
			name:     "ackW True",
			ackW:     true,
			expected: []byte{0, 0x02, 0, 0},
		},
		{
			name:     "ackEr True",
			ackEr:    true,
			expected: []byte{0, 0x01, 0, 0},
		},
		{
			name:          "partialAction True",
			partialAction: true,
			expected:      []byte{0, 0, 0x80, 0},
		},
		{
			name:                "scheduledOrExecuted True",
			scheduledOrExecuted: true,
			expected:            []byte{0, 0, 0x40, 0},
		},
	}
	for _, ta := range cases {
		t.Run(ta.name, func(t *testing.T) {
			a := AcknowledgeCAM{
				AckV:                ta.ackV,
				AckX:                ta.ackX,
				AckS:                ta.ackS,
				AckW:                ta.ackW,
				AckEr:               ta.ackEr,
				ScheduledOrExecuted: ta.scheduledOrExecuted,
				PartialAction:       ta.partialAction,
			}
			assert.Equal(t, ta.ackV, a.AckV)
			assert.Equal(t, ta.ackX, a.AckX)
			assert.Equal(t, ta.ackS, a.AckS)
			assert.Equal(t, ta.ackW, a.AckW)
			assert.Equal(t, ta.ackEr, a.AckEr)
			assert.Equal(t, ta.scheduledOrExecuted, a.ScheduledOrExecuted)
			assert.Equal(t, ta.partialAction, a.PartialAction)
			// Pack
			packed := a.Pack()
			assert.Equal(t, ta.expected, packed)
			// Unpack
			a.Unpack(packed)
			assert.Equal(t, ta.ackV, a.AckV)
			assert.Equal(t, ta.ackX, a.AckX)
			assert.Equal(t, ta.ackS, a.AckS)
			assert.Equal(t, ta.ackW, a.AckW)
			assert.Equal(t, ta.ackEr, a.AckEr)
			assert.Equal(t, ta.scheduledOrExecuted, a.ScheduledOrExecuted)
			assert.Equal(t, ta.partialAction, a.PartialAction)
		})
	}
}

func TestWIF0Default(t *testing.T) {
	w := WIF0{}
	assert.Equal(t, false, w.Wif7Enable)
	assert.Equal(t, false, w.Wif3Enable)
	assert.Equal(t, false, w.Wif2Enable)
	assert.Equal(t, false, w.Wif1Enable)

	// Pack
	packed := w.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)

	// Unpack
	w.Unpack(packed)
	assert.Equal(t, false, w.Wif7Enable)
	assert.Equal(t, false, w.Wif3Enable)
	assert.Equal(t, false, w.Wif2Enable)
	assert.Equal(t, false, w.Wif1Enable)
}

func TestWIF0(t *testing.T) {
	cases := []struct {
		name       string
		wif7Enable bool
		wif3Enable bool
		wif2Enable bool
		wif1Enable bool
		expected   []byte
	}{
		{
			name:       "Wif7Enable True",
			wif7Enable: true,
			expected:   []byte{0, 0, 0, 0x80},
		},
		{
			name:       "Wif3Enable True",
			wif3Enable: true,
			expected:   []byte{0, 0, 0, 0x08},
		},
		{
			name:       "Wif2Enable True",
			wif2Enable: true,
			expected:   []byte{0, 0, 0, 0x04},
		},
		{
			name:       "Wif1Enable True",
			wif1Enable: true,
			expected:   []byte{0, 0, 0, 0x02},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := WIF0{
				Wif7Enable: tc.wif7Enable,
				Wif3Enable: tc.wif3Enable,
				Wif2Enable: tc.wif2Enable,
				Wif1Enable: tc.wif1Enable,
			}
			assert.Equal(t, tc.wif7Enable, w.Wif7Enable)
			assert.Equal(t, tc.wif3Enable, w.Wif3Enable)
			assert.Equal(t, tc.wif2Enable, w.Wif2Enable)
			assert.Equal(t, tc.wif1Enable, w.Wif1Enable)

			// Pack
			packed := w.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			w.Unpack(packed)
			assert.Equal(t, tc.wif7Enable, w.Wif7Enable)
			assert.Equal(t, tc.wif3Enable, w.Wif3Enable)
			assert.Equal(t, tc.wif2Enable, w.Wif2Enable)
			assert.Equal(t, tc.wif1Enable, w.Wif1Enable)
		})
	}
}

func TestEIF0Default(t *testing.T) {
	e := EIF0{}
	assert.Equal(t, false, e.Eif7Enable)
	assert.Equal(t, false, e.Eif3Enable)
	assert.Equal(t, false, e.Eif2Enable)
	assert.Equal(t, false, e.Eif1Enable)

	// Pack
	packed := e.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)

	// Unpack
	e.Unpack(packed)
	assert.Equal(t, false, e.Eif7Enable)
	assert.Equal(t, false, e.Eif3Enable)
	assert.Equal(t, false, e.Eif2Enable)
	assert.Equal(t, false, e.Eif1Enable)
}

func TestEIF0(t *testing.T) {
	cases := []struct {
		name       string
		eif7Enable bool
		eif3Enable bool
		eif2Enable bool
		eif1Enable bool
		expected   []byte
	}{
		{
			name:       "Eif7Enable True",
			eif7Enable: true,
			expected:   []byte{0, 0, 0, 0x80},
		},
		{
			name:       "Eif3Enable True",
			eif3Enable: true,
			expected:   []byte{0, 0, 0, 0x08},
		},
		{
			name:       "Eif2Enable True",
			eif2Enable: true,
			expected:   []byte{0, 0, 0, 0x04},
		},
		{
			name:       "Eif1Enable True",
			eif1Enable: true,
			expected:   []byte{0, 0, 0, 0x02},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := EIF0{
				Eif7Enable: tc.eif7Enable,
				Eif3Enable: tc.eif3Enable,
				Eif2Enable: tc.eif2Enable,
				Eif1Enable: tc.eif1Enable,
			}
			assert.Equal(t, tc.eif7Enable, e.Eif7Enable)
			assert.Equal(t, tc.eif3Enable, e.Eif3Enable)
			assert.Equal(t, tc.eif2Enable, e.Eif2Enable)
			assert.Equal(t, tc.eif1Enable, e.Eif1Enable)

			// Pack
			packed := e.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			e.Unpack(packed)
			assert.Equal(t, tc.eif7Enable, e.Eif7Enable)
			assert.Equal(t, tc.eif3Enable, e.Eif3Enable)
			assert.Equal(t, tc.eif2Enable, e.Eif2Enable)
			assert.Equal(t, tc.eif1Enable, e.Eif1Enable)
		})
	}
}

func TestWarningErrorFieldsDefault(t *testing.T) {
	w := WarningErrorFields{}
	assert.Equal(t, false, w.FieldNotExecuted)
	assert.Equal(t, false, w.DeviceFailure)
	assert.Equal(t, false, w.ErroneousField)
	assert.Equal(t, false, w.ParamOutOfRange)
	assert.Equal(t, false, w.ParamUnsupportedPrecision)
	assert.Equal(t, false, w.FieldValueInvalid)
	assert.Equal(t, false, w.TimestampProblem)
	assert.Equal(t, false, w.HazardousPowerLevels)
	assert.Equal(t, false, w.Distortion)
	assert.Equal(t, false, w.InBandPowerCompliance)
	assert.Equal(t, false, w.OutOfBandPowerCompliance)
	assert.Equal(t, false, w.CositeInterference)
	assert.Equal(t, false, w.RegionalInterference)

	// Pack
	packed := w.Pack()
	expected := []byte{0, 0, 0, 0}
	assert.Equal(t, expected, packed)

	// Unpack
	w.Unpack(packed)
	assert.Equal(t, false, w.FieldNotExecuted)
	assert.Equal(t, false, w.DeviceFailure)
	assert.Equal(t, false, w.ErroneousField)
	assert.Equal(t, false, w.ParamOutOfRange)
	assert.Equal(t, false, w.ParamUnsupportedPrecision)
	assert.Equal(t, false, w.FieldValueInvalid)
	assert.Equal(t, false, w.TimestampProblem)
	assert.Equal(t, false, w.HazardousPowerLevels)
	assert.Equal(t, false, w.Distortion)
	assert.Equal(t, false, w.InBandPowerCompliance)
	assert.Equal(t, false, w.OutOfBandPowerCompliance)
	assert.Equal(t, false, w.CositeInterference)
	assert.Equal(t, false, w.RegionalInterference)
}

func TestWarningErrorFields(t *testing.T) {
	cases := []struct {
		name                      string
		fieldNotExecuted          bool
		deviceFailure             bool
		erroneousField            bool
		paramOutOfRange           bool
		paramUnsupportedPrecision bool
		fieldValueInvalid         bool
		timestampProblem          bool
		hazardousPowerLevels      bool
		distortion                bool
		inBandPowerCompliance     bool
		outOfBandPowerCompliance  bool
		cositeInterference        bool
		regionalInterference      bool
		expected                  []byte
	}{
		{
			name:             "FieldNotExecuted True",
			fieldNotExecuted: true,
			expected:         []byte{0x80, 0, 0, 0},
		},
		{
			name:          "DeviceFailure True",
			deviceFailure: true,
			expected:      []byte{0x40, 0, 0, 0},
		},
		{
			name:           "ErroneousField True",
			erroneousField: true,
			expected:       []byte{0x20, 0, 0, 0},
		},
		{
			name:            "ParamOutOfRange True",
			paramOutOfRange: true,
			expected:        []byte{0x10, 0, 0, 0},
		},
		{
			name:                      "ParamUnsupportedPrecision True",
			paramUnsupportedPrecision: true,
			expected:                  []byte{0x08, 0, 0, 0},
		},
		{
			name:              "FieldValueInvalid True",
			fieldValueInvalid: true,
			expected:          []byte{0x04, 0, 0, 0},
		},
		{
			name:             "TimestampProblem True",
			timestampProblem: true,
			expected:         []byte{0x02, 0, 0, 0},
		},
		{
			name:                 "HazardousPowerLevels True",
			hazardousPowerLevels: true,
			expected:             []byte{0x01, 0, 0, 0},
		},
		{
			name:       "Distortion True",
			distortion: true,
			expected:   []byte{0, 0x80, 0, 0},
		},
		{
			name:                  "InBandPowerCompliance True",
			inBandPowerCompliance: true,
			expected:              []byte{0, 0x40, 0, 0},
		},
		{
			name:                     "OutOfBandPowerCompliance True",
			outOfBandPowerCompliance: true,
			expected:                 []byte{0, 0x20, 0, 0},
		},
		{
			name:               "CositeInterference True",
			cositeInterference: true,
			expected:           []byte{0, 0x10, 0, 0},
		},
		{
			name:                 "RegionalInterference True",
			regionalInterference: true,
			expected:             []byte{0, 0x08, 0, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := WarningErrorFields{
				FieldNotExecuted:          tc.fieldNotExecuted,
				DeviceFailure:             tc.deviceFailure,
				ErroneousField:            tc.erroneousField,
				ParamOutOfRange:           tc.paramOutOfRange,
				ParamUnsupportedPrecision: tc.paramUnsupportedPrecision,
				FieldValueInvalid:         tc.fieldValueInvalid,
				TimestampProblem:          tc.timestampProblem,
				HazardousPowerLevels:      tc.hazardousPowerLevels,
				Distortion:                tc.distortion,
				InBandPowerCompliance:     tc.inBandPowerCompliance,
				OutOfBandPowerCompliance:  tc.outOfBandPowerCompliance,
				CositeInterference:        tc.cositeInterference,
				RegionalInterference:      tc.regionalInterference,
			}
			// Assert initial values
			assert.Equal(t, tc.fieldNotExecuted, w.FieldNotExecuted)
			assert.Equal(t, tc.deviceFailure, w.DeviceFailure)
			assert.Equal(t, tc.erroneousField, w.ErroneousField)
			assert.Equal(t, tc.paramOutOfRange, w.ParamOutOfRange)
			assert.Equal(t, tc.paramUnsupportedPrecision, w.ParamUnsupportedPrecision)
			assert.Equal(t, tc.fieldValueInvalid, w.FieldValueInvalid)
			assert.Equal(t, tc.timestampProblem, w.TimestampProblem)
			assert.Equal(t, tc.hazardousPowerLevels, w.HazardousPowerLevels)
			assert.Equal(t, tc.distortion, w.Distortion)
			assert.Equal(t, tc.inBandPowerCompliance, w.InBandPowerCompliance)
			assert.Equal(t, tc.outOfBandPowerCompliance, w.OutOfBandPowerCompliance)
			assert.Equal(t, tc.cositeInterference, w.CositeInterference)
			assert.Equal(t, tc.regionalInterference, w.RegionalInterference)

			// Pack
			packed := w.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			w.Unpack(packed)
			// Assert unpacked values
			assert.Equal(t, tc.fieldNotExecuted, w.FieldNotExecuted)
			assert.Equal(t, tc.deviceFailure, w.DeviceFailure)
			assert.Equal(t, tc.erroneousField, w.ErroneousField)
			assert.Equal(t, tc.paramOutOfRange, w.ParamOutOfRange)
			assert.Equal(t, tc.paramUnsupportedPrecision, w.ParamUnsupportedPrecision)
			assert.Equal(t, tc.fieldValueInvalid, w.FieldValueInvalid)
			assert.Equal(t, tc.timestampProblem, w.TimestampProblem)
			assert.Equal(t, tc.hazardousPowerLevels, w.HazardousPowerLevels)
			assert.Equal(t, tc.distortion, w.Distortion)
			assert.Equal(t, tc.inBandPowerCompliance, w.InBandPowerCompliance)
			assert.Equal(t, tc.outOfBandPowerCompliance, w.OutOfBandPowerCompliance)
			assert.Equal(t, tc.cositeInterference, w.CositeInterference)
			assert.Equal(t, tc.regionalInterference, w.RegionalInterference)
		})
	}
}
