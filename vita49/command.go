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

type IdentifierFormat uint8

const (
	Word IdentifierFormat = iota
	UUID
)

type ActionMode uint8

const (
	NoAction ActionMode = iota
	DryRun
	Execute
)

type TimestampControlMode uint8

const (
	Ignore TimestampControlMode = iota
	Device
	Late
	Early
	EarlyLate
	TimingIssues = 7
)

type CAM struct {
	ControlleeEnable bool
	ControlleeFormat IdentifierFormat
	ControllerEnable bool
	ControllerFormat IdentifierFormat
	PermitPartial    bool
	PermitWarnings   bool
	PermitErrors     bool
	ActionMode       ActionMode
	NackOnly         bool
	TimingControl    TimestampControlMode
}

func (c *CAM) Size() uint32 {
	return 4
}

func (c *CAM) Pack() []byte {
	buf := make([]byte, c.Size())
	var bitmap uint32
	bitmap |= indicatorFieldUint(c.ControlleeEnable, 31)
	var ceFormat bool
	if c.ControlleeFormat == UUID {
		ceFormat = true
	}
	bitmap |= indicatorFieldUint(ceFormat, 30)
	bitmap |= indicatorFieldUint(c.ControllerEnable, 29)
	var crFormat bool
	if c.ControllerFormat == UUID {
		crFormat = true
	}
	bitmap |= indicatorFieldUint(crFormat, 28)
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
	if ceFormat := indicatorFieldBool(bitmap, 30); ceFormat {
		c.ControlleeFormat = UUID
	} else {
		c.ControlleeFormat = Word
	}
	c.ControllerEnable = indicatorFieldBool(bitmap, 29)
	if crFormat := indicatorFieldBool(bitmap, 28); crFormat {
		c.ControllerFormat = UUID
	} else {
		c.ControllerFormat = Word
	}
	c.PermitPartial = indicatorFieldBool(bitmap, 27)
	c.PermitWarnings = indicatorFieldBool(bitmap, 26)
	c.PermitErrors = indicatorFieldBool(bitmap, 25)
	c.ActionMode = ActionMode(uint8(bitmap>>23) & 3)
	c.NackOnly = indicatorFieldBool(bitmap, 22)
	c.TimingControl = TimestampControlMode(uint8(bitmap>>12) & 7)
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
