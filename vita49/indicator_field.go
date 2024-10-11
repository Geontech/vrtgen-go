package vita49

import "encoding/binary"

const (
	indicatorFieldBytes = uint32(4)
)

type IndicatorField struct{}

func (f IndicatorField) Size() uint32 {
	return indicatorFieldBytes
}

type IndicatorField0 struct {
	IndicatorField
	ChangeIndicator          bool // bit position 31
	ReferencePointId         bool // bit position 30
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
	DeviceId                 bool // bit position 17
	StateEventIndicators     bool // bit position 16
	SignalDataFormat         bool // bit position 15
	FormattedGps             bool // bit position 14
	FormattedIns             bool // bit position 13
	EcefEphemeris            bool // bit position 12
	RelativeEphemeris        bool // bit position 11
	EphemerisRefId           bool // bit position 10
	GpsAscii                 bool // bit position 9
	ContextAssociationLists  bool // bit position 8
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
	CitedSid                bool // bit position 30
	SiblingSid              bool // bit position 29
	ParentSid               bool // bit position 28
	ChildSid                bool // bit position 27
	CitedMessageId          bool // bit position 26
	ControlleeId            bool // bit position 25
	ControlleeUuid          bool // bit position 24
	ControllerId            bool // bit position 23
	ControllerUuid          bool // bit position 22
	InformationSource       bool // bit position 21
	TraceId                 bool // bit position 20
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
	FunctionId              bool // bit position 9
	ModeId                  bool // bit position 8
	EventId                 bool // bit position 7
	FunctionPriorityId      bool // bit position 6
	CommunicationPriorityId bool // bit position 5
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
	NetworkId            bool // bit position 1
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

func (f *IndicatorField0) Pack(buf []byte) {
	var bitmap uint32
	bitmap |= indicatorFieldUint(f.ChangeIndicator, 31)
	bitmap |= indicatorFieldUint(f.ReferencePointId, 30)
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
	bitmap |= indicatorFieldUint(f.DeviceId, 17)
	bitmap |= indicatorFieldUint(f.StateEventIndicators, 16)
	bitmap |= indicatorFieldUint(f.SignalDataFormat, 15)
	bitmap |= indicatorFieldUint(f.FormattedGps, 14)
	bitmap |= indicatorFieldUint(f.FormattedIns, 13)
	bitmap |= indicatorFieldUint(f.EcefEphemeris, 12)
	bitmap |= indicatorFieldUint(f.RelativeEphemeris, 11)
	bitmap |= indicatorFieldUint(f.EphemerisRefId, 10)
	bitmap |= indicatorFieldUint(f.GpsAscii, 9)
	bitmap |= indicatorFieldUint(f.ContextAssociationLists, 8)
	binary.BigEndian.PutUint32(buf, bitmap)
}

func (f *IndicatorField0) Unpack(buf []byte) {
	bitmap := binary.BigEndian.Uint32(buf)
	f.ChangeIndicator = indicatorFieldBool(bitmap, 31)
	f.ReferencePointId = indicatorFieldBool(bitmap, 30)
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
	f.DeviceId = indicatorFieldBool(bitmap, 17)
	f.StateEventIndicators = indicatorFieldBool(bitmap, 16)
	f.SignalDataFormat = indicatorFieldBool(bitmap, 15)
	f.FormattedGps = indicatorFieldBool(bitmap, 14)
	f.FormattedIns = indicatorFieldBool(bitmap, 13)
	f.EcefEphemeris = indicatorFieldBool(bitmap, 12)
	f.RelativeEphemeris = indicatorFieldBool(bitmap, 11)
	f.EphemerisRefId = indicatorFieldBool(bitmap, 10)
	f.GpsAscii = indicatorFieldBool(bitmap, 9)
	f.ContextAssociationLists = indicatorFieldBool(bitmap, 8)
}

func (f *IndicatorField1) Pack(buf []byte) {
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
