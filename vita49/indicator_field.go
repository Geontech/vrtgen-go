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

type IndicatorField struct{}

func (f *IndicatorField) Size() uint32 {
	return 4
}

type IndicatorField0 struct {
	IndicatorField
	ChangeIndicator          bool // bit position 31
	ReferencePointID         bool // bit position 30
	Bandwidth                bool // bit position 29
	IfRefFrequency           bool // bit position 28
	RfRefFrequency           bool // bit position 27
	RfRefFrequencyOffset     bool // bit position 26
	IfBandOffset             bool // bit position 25
	ReferenceLevel           bool // bit position 24
	Gain                     bool // bit position 23
	OverRangeCount           bool // bit position 22
	SampleRate               bool // bit position 21
	TimestampAdjustment      bool // bit position 20
	TimestampCalibrationTime bool // bit position 19
	Temperature              bool // bit position 18
	DeviceID                 bool // bit position 17
	StateEventIndicators     bool // bit position 16
	SignalDataFormat         bool // bit position 15
	FormattedGps             bool // bit position 14
	FormattedIns             bool // bit position 13
	EcefEphemeris            bool // bit position 12
	RelativeEphemeris        bool // bit position 11
	EphemerisRefID           bool // bit position 10
	GpsAscii                 bool // bit position 9
	ContextAssociationLists  bool // bit position 8
	If7Enable                bool // bit position 7
	If3Enable                bool // bit position 3
	If2Enable                bool // bit position 2
	If1Enable                bool // bit position 1
}

type IndicatorField1 struct {
	IndicatorField
	PhaseOffset             bool // bit position 31
	Polarization            bool // bit position 30
	PointingVector          bool // bit position 29
	PointingVectorStructure bool // bit position 28
	SpatialScanType         bool // bit position 27
	SpatialReferenceType    bool // bit position 26
	BeamWidth               bool // bit position 25
	Range                   bool // bit position 24
	EbnoBer                 bool // bit position 20
	Threshold               bool // bit position 19
	CompressionPoint        bool // bit position 18
	InterceptPoints         bool // bit position 17
	SnrNoiseFigure          bool // bit position 16
	AuxFrequency            bool // bit position 15
	AuxGain                 bool // bit position 14
	AuxBandwidth            bool // bit position 13
	ArrayOfCifs             bool // bit position 11
	Spectrum                bool // bit position 10
	SectorStepScan          bool // bit position 9
	IndexList               bool // bit position 7
	DiscreteIo32            bool // bit position 6
	DiscreteIo64            bool // bit position 5
	HealthStatus            bool // bit position 4
	V49SpecCompliance       bool // bit position 3
	VersionInformation      bool // bit position 2
	BufferSize              bool // bit position 1
}

type IndicatorField2 struct {
	IndicatorField
	Bind                    bool // bit position 31
	CitedSID                bool // bit position 30
	SiblingSID              bool // bit position 29
	ParentSID               bool // bit position 28
	ChildSID                bool // bit position 27
	CitedMessageID          bool // bit position 26
	ControlleeID            bool // bit position 25
	ControlleeUUID          bool // bit position 24
	ControllerID            bool // bit position 23
	ControllerUUID          bool // bit position 22
	InformationSource       bool // bit position 21
	TraceID                 bool // bit position 20
	CountryCode             bool // bit position 19
	Operator                bool // bit position 18
	PlatformClass           bool // bit position 17
	PlatformInstance        bool // bit position 16
	PlatformDisplay         bool // bit position 15
	EmsDeviceClass          bool // bit position 14
	EmsDeviceType           bool // bit position 13
	EmsDeviceInstance       bool // bit position 12
	ModulationClass         bool // bit position 11
	ModulationType          bool // bit position 10
	FunctionID              bool // bit position 9
	ModeID                  bool // bit position 8
	EventID                 bool // bit position 7
	FunctionPriorityID      bool // bit position 6
	CommunicationPriorityID bool // bit position 5
	RfFootprint             bool // bit position 4
	RfFootprintRange        bool // bit position 3
}

type IndicatorField3 struct {
	IndicatorField
	TimestampDetails     bool // bit position 31
	TimestampSkew        bool // bit position 30
	RiseTime             bool // bit position 27
	FallTime             bool // bit position 26
	OffsetTime           bool // bit position 25
	PulseWidth           bool // bit position 24
	Period               bool // bit position 23
	Duration             bool // bit position 22
	Dwell                bool // bit position 21
	Jitter               bool // bit position 20
	Age                  bool // bit position 17
	ShelfLife            bool // bit position 16
	AirTemperature       bool // bit position 7
	SeaGroundTemperature bool // bit position 6
	Humidity             bool // bit position 5
	BarometricPressure   bool // bit position 4
	SeaSwellState        bool // bit position 3
	TroposphericState    bool // bit position 2
	NetworkID            bool // bit position 1
}

type IndicatorField7 struct {
	IndicatorField
	CurrentValue      bool // bit position 31
	AverageValue      bool // bit position 30
	MedianValue       bool // bit position 29
	StandardDeviation bool // bit position 28
	MaxValue          bool // bit position 27
	MinValue          bool // bit position 26
	Precision         bool // bit position 25
	Accuracy          bool // bit position 24
	FirstDerivative   bool // bit position 23
	SecondDerivative  bool // bit position 22
	ThirdDerivative   bool // bit position 21
	Probability       bool // bit position 20
	Belief            bool // bit position 19
}

func indicatorFieldUint(v bool, b uint32) uint32 {
	var value uint32
	if v {
		value = 1
	}
	return value << b
}

func indicatorFieldBool(v uint32, b uint32) bool {
	return (v & (uint32(1) << b)) != 0
}

func (f *IndicatorField0) Pack() []byte {
	buf := make([]byte, f.Size())
	var bitmap uint32
	bitmap |= indicatorFieldUint(f.ChangeIndicator, 31)
	bitmap |= indicatorFieldUint(f.ReferencePointID, 30)
	bitmap |= indicatorFieldUint(f.Bandwidth, 29)
	bitmap |= indicatorFieldUint(f.IfRefFrequency, 28)
	bitmap |= indicatorFieldUint(f.RfRefFrequency, 27)
	bitmap |= indicatorFieldUint(f.RfRefFrequencyOffset, 26)
	bitmap |= indicatorFieldUint(f.IfBandOffset, 25)
	bitmap |= indicatorFieldUint(f.ReferenceLevel, 24)
	bitmap |= indicatorFieldUint(f.Gain, 23)
	bitmap |= indicatorFieldUint(f.OverRangeCount, 22)
	bitmap |= indicatorFieldUint(f.SampleRate, 21)
	bitmap |= indicatorFieldUint(f.TimestampAdjustment, 20)
	bitmap |= indicatorFieldUint(f.TimestampCalibrationTime, 19)
	bitmap |= indicatorFieldUint(f.Temperature, 18)
	bitmap |= indicatorFieldUint(f.DeviceID, 17)
	bitmap |= indicatorFieldUint(f.StateEventIndicators, 16)
	bitmap |= indicatorFieldUint(f.SignalDataFormat, 15)
	bitmap |= indicatorFieldUint(f.FormattedGps, 14)
	bitmap |= indicatorFieldUint(f.FormattedIns, 13)
	bitmap |= indicatorFieldUint(f.EcefEphemeris, 12)
	bitmap |= indicatorFieldUint(f.RelativeEphemeris, 11)
	bitmap |= indicatorFieldUint(f.EphemerisRefID, 10)
	bitmap |= indicatorFieldUint(f.GpsAscii, 9)
	bitmap |= indicatorFieldUint(f.ContextAssociationLists, 8)
	bitmap |= indicatorFieldUint(f.If7Enable, 7)
	bitmap |= indicatorFieldUint(f.If3Enable, 3)
	bitmap |= indicatorFieldUint(f.If2Enable, 2)
	bitmap |= indicatorFieldUint(f.If1Enable, 1)
	binary.BigEndian.PutUint32(buf, bitmap)
	return buf
}

func (f *IndicatorField0) Unpack(buf []byte) {
	bitmap := binary.BigEndian.Uint32(buf)
	f.ChangeIndicator = indicatorFieldBool(bitmap, 31)
	f.ReferencePointID = indicatorFieldBool(bitmap, 30)
	f.Bandwidth = indicatorFieldBool(bitmap, 29)
	f.IfRefFrequency = indicatorFieldBool(bitmap, 28)
	f.RfRefFrequency = indicatorFieldBool(bitmap, 27)
	f.RfRefFrequencyOffset = indicatorFieldBool(bitmap, 26)
	f.IfBandOffset = indicatorFieldBool(bitmap, 25)
	f.ReferenceLevel = indicatorFieldBool(bitmap, 24)
	f.Gain = indicatorFieldBool(bitmap, 23)
	f.OverRangeCount = indicatorFieldBool(bitmap, 22)
	f.SampleRate = indicatorFieldBool(bitmap, 21)
	f.TimestampAdjustment = indicatorFieldBool(bitmap, 20)
	f.TimestampCalibrationTime = indicatorFieldBool(bitmap, 19)
	f.Temperature = indicatorFieldBool(bitmap, 18)
	f.DeviceID = indicatorFieldBool(bitmap, 17)
	f.StateEventIndicators = indicatorFieldBool(bitmap, 16)
	f.SignalDataFormat = indicatorFieldBool(bitmap, 15)
	f.FormattedGps = indicatorFieldBool(bitmap, 14)
	f.FormattedIns = indicatorFieldBool(bitmap, 13)
	f.EcefEphemeris = indicatorFieldBool(bitmap, 12)
	f.RelativeEphemeris = indicatorFieldBool(bitmap, 11)
	f.EphemerisRefID = indicatorFieldBool(bitmap, 10)
	f.GpsAscii = indicatorFieldBool(bitmap, 9)
	f.ContextAssociationLists = indicatorFieldBool(bitmap, 8)
	f.If7Enable = indicatorFieldBool(bitmap, 7)
	f.If3Enable = indicatorFieldBool(bitmap, 3)
	f.If2Enable = indicatorFieldBool(bitmap, 2)
	f.If1Enable = indicatorFieldBool(bitmap, 1)
}

func (f *IndicatorField1) Pack() []byte {
	buf := make([]byte, f.Size())
	var bitmap uint32
	bitmap |= indicatorFieldUint(f.PhaseOffset, 31)
	bitmap |= indicatorFieldUint(f.Polarization, 30)
	bitmap |= indicatorFieldUint(f.PointingVector, 29)
	bitmap |= indicatorFieldUint(f.PointingVectorStructure, 28)
	bitmap |= indicatorFieldUint(f.SpatialScanType, 27)
	bitmap |= indicatorFieldUint(f.SpatialReferenceType, 26)
	bitmap |= indicatorFieldUint(f.BeamWidth, 25)
	bitmap |= indicatorFieldUint(f.Range, 24)
	bitmap |= indicatorFieldUint(f.EbnoBer, 20)
	bitmap |= indicatorFieldUint(f.Threshold, 19)
	bitmap |= indicatorFieldUint(f.CompressionPoint, 18)
	bitmap |= indicatorFieldUint(f.InterceptPoints, 17)
	bitmap |= indicatorFieldUint(f.SnrNoiseFigure, 16)
	bitmap |= indicatorFieldUint(f.AuxFrequency, 15)
	bitmap |= indicatorFieldUint(f.AuxGain, 14)
	bitmap |= indicatorFieldUint(f.AuxBandwidth, 13)
	bitmap |= indicatorFieldUint(f.ArrayOfCifs, 11)
	bitmap |= indicatorFieldUint(f.Spectrum, 10)
	bitmap |= indicatorFieldUint(f.SectorStepScan, 9)
	bitmap |= indicatorFieldUint(f.IndexList, 7)
	bitmap |= indicatorFieldUint(f.DiscreteIo32, 6)
	bitmap |= indicatorFieldUint(f.DiscreteIo64, 5)
	bitmap |= indicatorFieldUint(f.HealthStatus, 4)
	bitmap |= indicatorFieldUint(f.V49SpecCompliance, 3)
	bitmap |= indicatorFieldUint(f.VersionInformation, 2)
	bitmap |= indicatorFieldUint(f.BufferSize, 1)
	binary.BigEndian.PutUint32(buf, bitmap)
	return buf
}

func (f *IndicatorField1) Unpack(buf []byte) {
	bitmap := binary.BigEndian.Uint32(buf)
	f.PhaseOffset = indicatorFieldBool(bitmap, 31)
	f.Polarization = indicatorFieldBool(bitmap, 30)
	f.PointingVector = indicatorFieldBool(bitmap, 29)
	f.PointingVectorStructure = indicatorFieldBool(bitmap, 28)
	f.SpatialScanType = indicatorFieldBool(bitmap, 27)
	f.SpatialReferenceType = indicatorFieldBool(bitmap, 26)
	f.BeamWidth = indicatorFieldBool(bitmap, 25)
	f.Range = indicatorFieldBool(bitmap, 24)
	f.EbnoBer = indicatorFieldBool(bitmap, 20)
	f.Threshold = indicatorFieldBool(bitmap, 19)
	f.CompressionPoint = indicatorFieldBool(bitmap, 18)
	f.InterceptPoints = indicatorFieldBool(bitmap, 17)
	f.SnrNoiseFigure = indicatorFieldBool(bitmap, 16)
	f.AuxFrequency = indicatorFieldBool(bitmap, 15)
	f.AuxGain = indicatorFieldBool(bitmap, 14)
	f.AuxBandwidth = indicatorFieldBool(bitmap, 13)
	f.ArrayOfCifs = indicatorFieldBool(bitmap, 11)
	f.Spectrum = indicatorFieldBool(bitmap, 10)
	f.SectorStepScan = indicatorFieldBool(bitmap, 9)
	f.IndexList = indicatorFieldBool(bitmap, 7)
	f.DiscreteIo32 = indicatorFieldBool(bitmap, 6)
	f.DiscreteIo64 = indicatorFieldBool(bitmap, 5)
	f.HealthStatus = indicatorFieldBool(bitmap, 4)
	f.V49SpecCompliance = indicatorFieldBool(bitmap, 3)
	f.VersionInformation = indicatorFieldBool(bitmap, 2)
	f.BufferSize = indicatorFieldBool(bitmap, 1)
}

func (f *IndicatorField2) Pack() []byte {
	buf := make([]byte, f.Size())
	var bitmap uint32
	bitmap |= indicatorFieldUint(f.Bind, 31)
	bitmap |= indicatorFieldUint(f.CitedSID, 30)
	bitmap |= indicatorFieldUint(f.SiblingSID, 29)
	bitmap |= indicatorFieldUint(f.ParentSID, 28)
	bitmap |= indicatorFieldUint(f.ChildSID, 27)
	bitmap |= indicatorFieldUint(f.CitedMessageID, 26)
	bitmap |= indicatorFieldUint(f.ControlleeID, 25)
	bitmap |= indicatorFieldUint(f.ControlleeUUID, 24)
	bitmap |= indicatorFieldUint(f.ControllerID, 23)
	bitmap |= indicatorFieldUint(f.ControllerUUID, 22)
	bitmap |= indicatorFieldUint(f.InformationSource, 21)
	bitmap |= indicatorFieldUint(f.TraceID, 20)
	bitmap |= indicatorFieldUint(f.CountryCode, 19)
	bitmap |= indicatorFieldUint(f.Operator, 18)
	bitmap |= indicatorFieldUint(f.PlatformClass, 17)
	bitmap |= indicatorFieldUint(f.PlatformInstance, 16)
	bitmap |= indicatorFieldUint(f.PlatformDisplay, 15)
	bitmap |= indicatorFieldUint(f.EmsDeviceClass, 14)
	bitmap |= indicatorFieldUint(f.EmsDeviceType, 13)
	bitmap |= indicatorFieldUint(f.EmsDeviceInstance, 12)
	bitmap |= indicatorFieldUint(f.ModulationClass, 11)
	bitmap |= indicatorFieldUint(f.ModulationType, 10)
	bitmap |= indicatorFieldUint(f.FunctionID, 9)
	bitmap |= indicatorFieldUint(f.ModeID, 8)
	bitmap |= indicatorFieldUint(f.EventID, 7)
	bitmap |= indicatorFieldUint(f.FunctionPriorityID, 6)
	bitmap |= indicatorFieldUint(f.CommunicationPriorityID, 5)
	bitmap |= indicatorFieldUint(f.RfFootprint, 4)
	bitmap |= indicatorFieldUint(f.RfFootprintRange, 3)
	binary.BigEndian.PutUint32(buf, bitmap)
	return buf
}

func (f *IndicatorField2) Unpack(buf []byte) {
	bitmap := binary.BigEndian.Uint32(buf)
	f.Bind = indicatorFieldBool(bitmap, 31)
	f.CitedSID = indicatorFieldBool(bitmap, 30)
	f.SiblingSID = indicatorFieldBool(bitmap, 29)
	f.ParentSID = indicatorFieldBool(bitmap, 28)
	f.ChildSID = indicatorFieldBool(bitmap, 27)
	f.CitedMessageID = indicatorFieldBool(bitmap, 26)
	f.ControlleeID = indicatorFieldBool(bitmap, 25)
	f.ControlleeUUID = indicatorFieldBool(bitmap, 24)
	f.ControllerID = indicatorFieldBool(bitmap, 23)
	f.ControllerUUID = indicatorFieldBool(bitmap, 22)
	f.InformationSource = indicatorFieldBool(bitmap, 21)
	f.TraceID = indicatorFieldBool(bitmap, 20)
	f.CountryCode = indicatorFieldBool(bitmap, 19)
	f.Operator = indicatorFieldBool(bitmap, 18)
	f.PlatformClass = indicatorFieldBool(bitmap, 17)
	f.PlatformInstance = indicatorFieldBool(bitmap, 16)
	f.PlatformDisplay = indicatorFieldBool(bitmap, 15)
	f.EmsDeviceClass = indicatorFieldBool(bitmap, 14)
	f.EmsDeviceType = indicatorFieldBool(bitmap, 13)
	f.EmsDeviceInstance = indicatorFieldBool(bitmap, 12)
	f.ModulationClass = indicatorFieldBool(bitmap, 11)
	f.ModulationType = indicatorFieldBool(bitmap, 10)
	f.FunctionID = indicatorFieldBool(bitmap, 9)
	f.ModeID = indicatorFieldBool(bitmap, 8)
	f.EventID = indicatorFieldBool(bitmap, 7)
	f.FunctionPriorityID = indicatorFieldBool(bitmap, 6)
	f.CommunicationPriorityID = indicatorFieldBool(bitmap, 5)
	f.RfFootprint = indicatorFieldBool(bitmap, 4)
	f.RfFootprintRange = indicatorFieldBool(bitmap, 3)
}
