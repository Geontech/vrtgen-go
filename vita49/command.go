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

type ControlFormat bool

const (
	UUID ControlFormat = true
	ID   ControlFormat = false
)

type CAM struct {
	ControlleeEnable bool
	ControlleeFormat ControlFormat
	ControllerEnable bool
	ControllerFormat ControlFormat
	PermitPartial    bool
	PermitWarnings   bool
	PermitErrors     bool
	ActionMode       uint8
	NackOnly         bool
	TimingControl    uint8
}

func (c *CAM) Size() uint32 {
	return 4
}

func (c *CAM) Pack() []byte {
	buf := make([]byte, c.Size())
	var bitmap uint32
	bitmap |= indicatorFieldUint(c.ControlleeEnable, 31)
	bitmap |= indicatorFieldUint(bool(c.ControlleeFormat), 30)
	bitmap |= indicatorFieldUint(c.ControllerEnable, 29)
	bitmap |= indicatorFieldUint(bool(c.ControllerFormat), 28)
	bitmap |= indicatorFieldUint(c.PermitPartial, 27)
	bitmap |= indicatorFieldUint(c.PermitWarnings, 26)
	bitmap |= indicatorFieldUint(c.PermitErrors, 25)
	bitmap |= uint32(c.ActionMode) << 23
	bitmap |= indicatorFieldUint(c.NackOnly, 22)
	bitmap |= uint32(c.TimingControl) << 12
	binary.BigEndian.PutUint32(buf[0:], bitmap)
	return buf
}

func (c *CAM) Unpack(buf []byte) {
	bitmap := binary.BigEndian.Uint32(buf[0:])
	c.ControlleeEnable = indicatorFieldBool(bitmap, 31)
	c.ControlleeFormat = ControlFormat(indicatorFieldBool(bitmap, 30))
	c.ControllerEnable = indicatorFieldBool(bitmap, 29)
	c.ControllerFormat = ControlFormat(indicatorFieldBool(bitmap, 28))
	c.PermitPartial = indicatorFieldBool(bitmap, 27)
	c.PermitWarnings = indicatorFieldBool(bitmap, 26)
	c.PermitErrors = indicatorFieldBool(bitmap, 25)
	c.ActionMode = uint8(bitmap>>23) & 3
	c.NackOnly = indicatorFieldBool(bitmap, 22)
	c.TimingControl = uint8(bitmap>>12) & 7
}

type ControlCAM struct {
	CAM
	ReqV  bool
	ReqX  bool
	ReqS  bool
	ReqW  bool
	ReqEr bool
}

func (c *ControlCAM) Pack() []byte {
	buf := c.CAM.Pack()
	var bitmap uint32
	bitmap |= indicatorFieldUint(c.ReqV, 20)
	bitmap |= indicatorFieldUint(c.ReqX, 19)
	bitmap |= indicatorFieldUint(c.ReqS, 18)
	bitmap |= indicatorFieldUint(c.ReqW, 17)
	bitmap |= indicatorFieldUint(c.ReqEr, 16)
	buf[1] += uint8(bitmap >> 16)
	return buf
}

func (c *ControlCAM) Unpack(buf []byte) {
	c.CAM.Unpack(buf)
	bitmap := binary.BigEndian.Uint32(buf[0:])
	c.ReqV = indicatorFieldBool(bitmap, 20)
	c.ReqX = indicatorFieldBool(bitmap, 19)
	c.ReqS = indicatorFieldBool(bitmap, 18)
	c.ReqW = indicatorFieldBool(bitmap, 17)
	c.ReqEr = indicatorFieldBool(bitmap, 16)
}

type AcknowledgeCAM struct {
	CAM
	AckV                bool
	AckX                bool
	AckS                bool
	AckW                bool
	AckEr               bool
	PartialAction       bool
	ScheduledOrExecuted bool
}

func (a *AcknowledgeCAM) Pack() []byte {
	buf := a.CAM.Pack()
	var bitmap uint32
	bitmap |= indicatorFieldUint(a.AckV, 20)
	bitmap |= indicatorFieldUint(a.AckX, 19)
	bitmap |= indicatorFieldUint(a.AckS, 18)
	bitmap |= indicatorFieldUint(a.AckW, 17)
	bitmap |= indicatorFieldUint(a.AckEr, 16)
	bitmap |= indicatorFieldUint(a.PartialAction, 15)
	bitmap |= indicatorFieldUint(a.ScheduledOrExecuted, 14)
	buf[1] += uint8(bitmap >> 16)
	// Shift to third byte, then grab only bits 15 and 14
	buf[2] += uint8(bitmap>>8) & 0xC0
	return buf
}

func (a *AcknowledgeCAM) Unpack(buf []byte) {
	a.CAM.Unpack(buf)
	bitmap := binary.BigEndian.Uint32(buf[0:])
	a.AckV = indicatorFieldBool(bitmap, 20)
	a.AckX = indicatorFieldBool(bitmap, 19)
	a.AckS = indicatorFieldBool(bitmap, 18)
	a.AckW = indicatorFieldBool(bitmap, 17)
	a.AckEr = indicatorFieldBool(bitmap, 16)
	a.PartialAction = indicatorFieldBool(bitmap, 15)
	a.ScheduledOrExecuted = indicatorFieldBool(bitmap, 14)
}

type WIF0 struct {
	IndicatorField0
	Wif7Enable bool
	Wif3Enable bool
	Wif2Enable bool
	Wif1Enable bool
}

func (w *WIF0) Pack() []byte {
	w.IndicatorField0 = IndicatorField0{
		If7Enable: w.Wif7Enable,
		If3Enable: w.Wif3Enable,
		If2Enable: w.Wif2Enable,
		If1Enable: w.Wif1Enable,
	}
	return w.IndicatorField0.Pack()
}

func (w *WIF0) Unpack(buf []byte) {
	w.IndicatorField0.Unpack(buf)
	w.Wif7Enable = w.IndicatorField0.If7Enable
	w.Wif3Enable = w.IndicatorField0.If3Enable
	w.Wif2Enable = w.IndicatorField0.If2Enable
	w.Wif1Enable = w.IndicatorField0.If1Enable
}

type EIF0 struct {
	IndicatorField0
	Eif7Enable bool
	Eif3Enable bool
	Eif2Enable bool
	Eif1Enable bool
}

func (e *EIF0) Pack() []byte {
	e.IndicatorField0 = IndicatorField0{
		If7Enable: e.Eif7Enable,
		If3Enable: e.Eif3Enable,
		If2Enable: e.Eif2Enable,
		If1Enable: e.Eif1Enable,
	}
	return e.IndicatorField0.Pack()
}

func (e *EIF0) Unpack(buf []byte) {
	e.IndicatorField0.Unpack(buf)
	e.Eif7Enable = e.IndicatorField0.If7Enable
	e.Eif3Enable = e.IndicatorField0.If3Enable
	e.Eif2Enable = e.IndicatorField0.If2Enable
	e.Eif1Enable = e.IndicatorField0.If1Enable
}

type WEIF1 struct {
	IndicatorField1
}

type WEIF2 struct {
	IndicatorField2
}

type WEIF3 struct {
	IndicatorField3
}

type WEIF7 struct {
	IndicatorField7
}

type WarningErrorFields struct {
	FieldNotExecuted          bool
	DeviceFailure             bool
	ErroneousField            bool
	ParamOutOfRange           bool
	ParamUnsupportedPrecision bool
	FieldValueInvalid         bool
	TimestampProblem          bool
	HazardousPowerLevels      bool
	Distortion                bool
	InBandPowerCompliance     bool
	OutOfBandPowerCompliance  bool
	CositeInterference        bool
	RegionalInterference      bool
}

func (w *WarningErrorFields) Size() uint32 {
	return 4
}

func (w *WarningErrorFields) Pack() []byte {
	buf := make([]byte, w.Size())
	var bitmap uint32
	bitmap |= indicatorFieldUint(w.FieldNotExecuted, 31)
	bitmap |= indicatorFieldUint(w.DeviceFailure, 30)
	bitmap |= indicatorFieldUint(w.ErroneousField, 29)
	bitmap |= indicatorFieldUint(w.ParamOutOfRange, 28)
	bitmap |= indicatorFieldUint(w.ParamUnsupportedPrecision, 27)
	bitmap |= indicatorFieldUint(w.FieldValueInvalid, 26)
	bitmap |= indicatorFieldUint(w.TimestampProblem, 25)
	bitmap |= indicatorFieldUint(w.HazardousPowerLevels, 24)
	bitmap |= indicatorFieldUint(w.Distortion, 23)
	bitmap |= indicatorFieldUint(w.InBandPowerCompliance, 22)
	bitmap |= indicatorFieldUint(w.OutOfBandPowerCompliance, 21)
	bitmap |= indicatorFieldUint(w.CositeInterference, 20)
	bitmap |= indicatorFieldUint(w.RegionalInterference, 19)
	binary.BigEndian.PutUint32(buf, bitmap)
	return buf
}

func (w *WarningErrorFields) Unpack(buf []byte) {
	bitmap := binary.BigEndian.Uint32(buf)
	w.FieldNotExecuted = indicatorFieldBool(bitmap, 31)
	w.DeviceFailure = indicatorFieldBool(bitmap, 30)
	w.ErroneousField = indicatorFieldBool(bitmap, 29)
	w.ParamOutOfRange = indicatorFieldBool(bitmap, 28)
	w.ParamUnsupportedPrecision = indicatorFieldBool(bitmap, 27)
	w.FieldValueInvalid = indicatorFieldBool(bitmap, 26)
	w.TimestampProblem = indicatorFieldBool(bitmap, 25)
	w.HazardousPowerLevels = indicatorFieldBool(bitmap, 24)
	w.Distortion = indicatorFieldBool(bitmap, 23)
	w.InBandPowerCompliance = indicatorFieldBool(bitmap, 22)
	w.OutOfBandPowerCompliance = indicatorFieldBool(bitmap, 21)
	w.CositeInterference = indicatorFieldBool(bitmap, 20)
	w.RegionalInterference = indicatorFieldBool(bitmap, 19)
}
