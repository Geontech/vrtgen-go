package vita49

import (
	"encoding/binary"
	"fmt"
)

const (
	polarizationBytes         = uint32(4)
	pointingVectorBytes       = uint32(4)
	spatialReferenceTypeBytes = uint32(8)
	beamWidthBytes            = uint32(4)
	ebNoBERBytes              = uint32(4)
	thresholdBytes            = uint32(4)
	interceptPointsBytes      = uint32(4)
	sNRNoiseBytes             = uint32(4)
	spectrumTypeBytes         = uint32(4)
	windowTypeBytes           = uint32(8)
	spectrumF1F2IndiciesBytes = uint32(8)
	spectrumBytes             = uint32(8)
	sectorStepScanCIFBytes    = uint32(4)
	sectorStepScanRecordBytes = uint32(16)
	sectorStepScanBytes       = uint32(8) // Variable
	versionInformationBytes   = uint32(8)
)

type Cif1 struct {
	IndicatorField1
	PhaseOffset          uint32
	Polarization         Polarization
	PointingVector       PointingVector
	SpatialScanType      uint32
	SpatialReferenceType SpatialReferenceType
	BeamWidth            BeamWidth
	Range                uint32
	EbNoBER              EbNoBER
	Threshold            Threshold
	CompressionPoint     uint32
	InterceptPoints      InterceptPoints
	SNRNoise             SNRNoise
	AuxFrequency         uint32
	AuxGain              Gain
	AuxBandwith          uint32
	ArrayOfCifs          uint32
	Spectrum             Spectrum
	SectorStepScan       SectorStepScan
	IndexList            IndexList
	DiscreteIo32         uint32
	DiscreteIo64         uint64
	HealthStatus         uint32
	V49SpecCompliance    uint32
	VersionInformation   VersionInformation
	BufferSize           uint32
}

// Polarization
// Represents antenna polarization with tilt (inclination) and ellipticity angles
type Polarization struct {
	TiltAngle        float64
	EllipticityAngle float64
}

func (p Polarization) Size() uint32 {
	return polarizationBytes
}

func (p *Polarization) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	binary.BigEndian.PutUint16(retval[2:], uint16(ToFixed16(p.EllipticityAngle, 13)))
	binary.BigEndian.PutUint16(retval[0:], uint16(ToFixed16(p.TiltAngle, 13)))

	return retval
}

func (p *Polarization) Unpack(buf []byte) {
	p.EllipticityAngle = FromFixed(int16(binary.BigEndian.Uint16(buf[2:])), 13)
	p.TiltAngle = FromFixed(int16(binary.BigEndian.Uint16(buf[0:])), 13)
}

// PointingVector
// Allows for reporting or controlling the direction of RF energy from a system
type PointingVector struct {
	Elevation float64
	Azimuthal float64
}

func (p PointingVector) Size() uint32 {
	return pointingVectorBytes
}

func (p *PointingVector) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	binary.BigEndian.PutUint16(retval[2:], uint16(ToFixed16(p.Azimuthal, 7)))
	binary.BigEndian.PutUint16(retval[0:], uint16(ToFixed16(p.Elevation, 7)))
	return retval
}

func (p *PointingVector) Unpack(buf []byte) {
	p.Azimuthal = FromFixed(int16(binary.BigEndian.Uint16(buf[2:])), 7)
	p.Elevation = FromFixed(int16(binary.BigEndian.Uint16(buf[0:])), 7)
}

// Spatial Reference Type
// Describes the reference point for the antenna scan
type SpatialReferenceType struct {
	SpatialIdentifier uint16
	DefinedReference  uint8 // Reference Point for the antenna scan
	BeamType          uint8 // Type of antenna scan pattern being used
}

func (s SpatialReferenceType) Size() uint32 {
	return spatialReferenceTypeBytes
}

func (s *SpatialReferenceType) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	word1 := uint32(0)
	word1 |= uint32(s.SpatialIdentifier) << 16    // Bits 16-31
	word1 |= uint32(s.DefinedReference&0x03) << 2 // Bits 2 & 3
	word1 |= uint32(s.BeamType & 0x03)            // Bits 0 & 1
	binary.BigEndian.PutUint32(retval[0:], word1)

	return retval
}

func (s *SpatialReferenceType) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf[0:])
	s.SpatialIdentifier = uint16(word1 >> 16)
	s.DefinedReference = uint8((word1 >> 2) & 0x03)
	s.BeamType = uint8(word1 & 0x03)
}

// Beam Width
// The 3dB width of the main lobe
type BeamWidth struct {
	Horizontal float64
	Vertical   float64
}

func (p BeamWidth) Size() uint32 {
	return beamWidthBytes
}

func (p *BeamWidth) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	binary.BigEndian.PutUint16(retval[2:], uint16(ToFixed16(p.Vertical, 7)))
	binary.BigEndian.PutUint16(retval[0:], uint16(ToFixed16(p.Horizontal, 7)))

	return retval
}

func (p *BeamWidth) Unpack(buf []byte) {
	p.Vertical = FromFixed(int16(binary.BigEndian.Uint16(buf[2:])), 7)
	p.Horizontal = FromFixed(int16(binary.BigEndian.Uint16(buf[0:])), 7)
}

// EbNoBER
// EbNo - Energy per bit to noise density ratio
// A measure of the energy per bit to naise power per hertz of the signal for the signal
// BER - bit error rate
// A measure of the ratio fo the number of bits recieved in error to the total number of
// bits recieved over some period of time
type EbNoBER struct {
	Ebno float64
	Ber  float64
}

// Constructor for EbNoBer
func NewEbNoBer() *EbNoBER {
	return &EbNoBER{
		Ebno: 32767, // Default to 'not used' 0x7FFF
		Ber:  32767, // Default to 'not used' 0x7FFF
	}
}

func (e EbNoBER) Size() uint32 {
	return ebNoBERBytes
}

func (e *EbNoBER) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	// ebnoFixed := uint16(ToFixed16(e.Ebno, 7))
	// berFixed := uint16(ToFixed16(e.Ber, 7))

	// // Create a 32-bit word with ebno in bits 16-31 and ber in bits 0-15
	// word := uint32(ebnoFixed)<<16 | uint32(berFixed)
	// binary.BigEndian.PutUint32(retval, word)

	binary.BigEndian.PutUint16(retval[2:], uint16(ToFixed16(e.Ber, 7)))
	binary.BigEndian.PutUint16(retval[0:], uint16(ToFixed16(e.Ebno, 7)))

	return retval
}

func (e *EbNoBER) Unpack(buf []byte) {

	// Read the 32-bit word from the byte array in big-endian order
	word := binary.BigEndian.Uint32(buf)

	// Extract ebno from bits 16-31
	e.Ebno = FromFixed(int16(word>>16), 7) // Shift right by 16 bits

	// Extract ber from bits 0-15
	e.Ber = FromFixed(int16(word&0xFFFF), 7) // Mask with 0xFFFF to get lower 16 bits

	// e.ber = FromFixed(int16(binary.BigEndian.Uint16(retval[2:])), 7)
	// e.ebno = FromFixed(int16(binary.BigEndian.Uint16(retval[0:])), 7)

}

// Threshold
// Provides the ability to set a signal threshold level in dB or dBm,
// to trigger some signal based action
type Threshold struct {
	Stage1 float64
	Stage2 float64
}

func (t Threshold) Size() uint32 {
	return thresholdBytes
}

func (t *Threshold) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	binary.BigEndian.PutUint16(retval[2:], uint16(ToFixed16(t.Stage2, 7)))
	binary.BigEndian.PutUint16(retval[0:], uint16(ToFixed16(t.Stage1, 7)))

	return retval
}

func (t *Threshold) Unpack(buf []byte) {
	t.Stage2 = FromFixed(int16(binary.BigEndian.Uint16(buf[2:])), 7)
	t.Stage1 = FromFixed(int16(binary.BigEndian.Uint16(buf[0:])), 7)
}

// InterceptPoints
// Second and third order intercept points are combined into a single word
// for efficiency; they are often considered together as measures of a tuners distortion performance
type InterceptPoints struct {
	SecondOrder float64
	ThirdOrder  float64
}

func (i InterceptPoints) Size() uint32 {
	return interceptPointsBytes
}

func (i *InterceptPoints) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	binary.BigEndian.PutUint16(retval[2:], uint16(ToFixed16(i.ThirdOrder, 7)))
	binary.BigEndian.PutUint16(retval[0:], uint16(ToFixed16(i.SecondOrder, 7)))

	return retval
}

func (i *InterceptPoints) Unpack(buf []byte) {
	i.ThirdOrder = FromFixed(int16(binary.BigEndian.Uint16(buf[2:])), 7)
	i.SecondOrder = FromFixed(int16(binary.BigEndian.Uint16(buf[0:])), 7)
}

// SNRNoise
// Signal to noise ratio - a measure of the signal power to noise power (dB)
type SNRNoise struct {
	Snr   float64
	Noise float64
}

func (s SNRNoise) Size() uint32 {
	return sNRNoiseBytes
}

func (s *SNRNoise) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	binary.BigEndian.PutUint16(retval[2:], uint16(ToFixed16(s.Noise, 7)))
	binary.BigEndian.PutUint16(retval[0:], uint16(ToFixed16(s.Snr, 7)))

	return retval
}

func (s *SNRNoise) Unpack(buf []byte) {
	s.Noise = FromFixed(int16(binary.BigEndian.Uint16(buf[2:])), 7)
	s.Snr = FromFixed(int16(binary.BigEndian.Uint16(buf[0:])), 7)
}

// SpectrumType
// Describes or sets the basic characteristics of the spectral data
type SpectrumType struct {
	SpectrumType  uint8 // Type of spectral data being presented
	AveragingType uint8 // Indicates averaging type being performed
	WindowTime    uint8
}

func (s SpectrumType) Size() uint32 {
	return spectrumTypeBytes
}

func (s *SpectrumType) Pack() []byte {
	retval := make([]byte, 4)

	word1 := uint32(0)
	word1 |= uint32(s.WindowTime&0x0F) << 16
	word1 |= uint32(s.AveragingType&0xFF) << 8
	word1 |= uint32(s.SpectrumType & 0xFF)
	binary.BigEndian.PutUint32(retval[0:], word1)

	return retval
}

func (s *SpectrumType) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf)
	s.WindowTime = uint8((word1 >> 16) & 0x0F)
	s.AveragingType = uint8((word1 >> 8) & 0xFF)
	s.SpectrumType = uint8(word1 & 0xFF)
}

// WindowType
// Indicates the time-domain window that was used on the time-domain data before
// being transformed to the frequency domain
type WindowType struct {
	WindowType uint8
}

func (w WindowType) Size() uint32 {
	return windowTypeBytes
}

func (w *WindowType) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	word1 := uint32(0)
	word1 |= uint32(w.WindowType) // Byte 3
	binary.BigEndian.PutUint32(retval[0:], word1)

	return retval
}

func (w *WindowType) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf[0:])
	w.WindowType = uint8((word1 & 0xFF))
}

// SpectrumF1F2Indicies
// Used to indicate if only a subset of the total spectrum
// points that were computed are provided in the data packet
type SpectrumF1F2Indicies struct {
	F1Index uint32
	F2Index uint32
}

func (s SpectrumF1F2Indicies) Size() uint32 {
	return spectrumF1F2IndiciesBytes
}

func (s *SpectrumF1F2Indicies) Pack() []byte {
	// 2 words
	retval := make([]byte, 2*4)

	binary.BigEndian.PutUint32(retval[4:], s.F1Index)
	binary.BigEndian.PutUint32(retval[0:], s.F2Index)

	return retval
}

func (s *SpectrumF1F2Indicies) Unpack(buf []byte) {
	s.F1Index = binary.BigEndian.Uint32(buf[4:])
	s.F2Index = binary.BigEndian.Uint32(buf[0:])
}

// Spectrum
// Describes control or context for spectral information
type Spectrum struct {
	SpectrumType          SpectrumType
	WindowType            WindowType
	NumberTransformPoints uint32
	NumberWindowPoints    uint32
	Resolution            uint64
	Span                  uint64
	NumberAverages        uint32
	WeightingFactor       uint32
	SpectrumF1F2Indicies  SpectrumF1F2Indicies
	WindowTimeDelta       uint32
}

func (s Spectrum) Size() uint32 {
	return spectrumBytes
}

func (s *Spectrum) Pack() []byte {
	// 13 words total == 52 bytes
	retval := make([]byte, 13*4)

	// Spectrum type - 4 bytes
	spectrumWord := uint32(0)
	spectrumWord |= uint32(s.SpectrumType.SpectrumType) << 0     // Bits 0-7
	spectrumWord |= uint32(s.SpectrumType.AveragingType) << 8    // Bits 8-15
	spectrumWord |= uint32(s.SpectrumType.WindowTime&0x0F) << 16 // Bits 16-19
	binary.BigEndian.PutUint32(retval[48:], spectrumWord)        // Start at offset 48

	// Window type - 4 bytes
	windowWord := uint32(0)
	windowWord |= uint32(s.WindowType.WindowType)
	binary.BigEndian.PutUint32(retval[44:], windowWord) // Start at offset 44

	// No. of transform points - 4 bytes
	binary.BigEndian.PutUint32(retval[40:], s.NumberTransformPoints) // Start at offset 40

	// No. of window points - 4 bytes
	binary.BigEndian.PutUint32(retval[36:], s.NumberWindowPoints) // Start at offset 36

	// Resolution - 8 bytes
	binary.BigEndian.PutUint64(retval[28:], s.Resolution) // Start at offset 28

	// Span - 8 bytes
	binary.BigEndian.PutUint64(retval[20:], s.Span) // Start at offset 20

	// No. of averages - 4 bytes
	binary.BigEndian.PutUint32(retval[16:], s.NumberAverages) // Start at offset 16

	// Weighting factor - 4 bytes
	binary.BigEndian.PutUint32(retval[12:], uint32(s.WeightingFactor)) // Start at offset 12

	// Spectrum F1-F2 indices - 8 bytes
	binary.BigEndian.PutUint32(retval[8:], s.SpectrumF1F2Indicies.F1Index) // Start at offset 8
	binary.BigEndian.PutUint32(retval[4:], s.SpectrumF1F2Indicies.F2Index) // Start at offset 4

	// Window time-delta - 4 bytes
	binary.BigEndian.PutUint32(retval[0:], s.WindowTimeDelta) // Start at offset 0

	return retval
}

func (s *Spectrum) Unpack(buf []byte) {

	// Window time-delta - 4 bytes
	s.WindowTimeDelta = binary.BigEndian.Uint32(buf[0:])

	// Spectrum F1-F2 indices - 8 bytes
	s.SpectrumF1F2Indicies.F2Index = binary.BigEndian.Uint32(buf[4:])
	s.SpectrumF1F2Indicies.F1Index = binary.BigEndian.Uint32(buf[8:])

	// Weighting factor - 4 bytes
	s.WeightingFactor = uint32(binary.BigEndian.Uint32(buf[12:]))

	// No. of averages - 4 bytes
	s.NumberAverages = binary.BigEndian.Uint32(buf[16:])

	// Span - 8 bytes
	s.Span = binary.BigEndian.Uint64(buf[20:])

	// Resolution - 8 bytes
	s.Resolution = binary.BigEndian.Uint64(buf[28:])

	// No. of window points - 4 bytes
	s.NumberWindowPoints = binary.BigEndian.Uint32(buf[36:])

	// No. of transform points - 4 bytes
	s.NumberTransformPoints = binary.BigEndian.Uint32(buf[40:])

	// Window type - 4 bytes
	windowWord := binary.BigEndian.Uint32(buf[44:])
	s.WindowType.WindowType = uint8(windowWord & 0xFFFF) // Bits 0-15

	// Spectrum type - 4 bytes
	spectrumWord := binary.BigEndian.Uint32(buf[48:])
	s.SpectrumType.SpectrumType = uint8(spectrumWord & 0xFF)         // Bits 0-7
	s.SpectrumType.AveragingType = uint8((spectrumWord >> 8) & 0xFF) // Bits 8-15
	s.SpectrumType.WindowTime = uint8((spectrumWord >> 16) & 0x0F)   // Bits 16-19
}

// SectorStepScanCIF
// Provides a way to setup & report a multi-sectored scanning reciever
// Also provides for stepping individual frequencies, vs scanning
type SectorStepScanCIF struct {
	SectorNumber        bool // Required
	F1StartFrequency    bool // Required
	F2StartFrequency    bool
	ResolutionBandwidth bool
	TuneStepSize        bool
	NumberPoints        bool
	DefaultGain         bool
	Threshold           bool
	DwellTime           bool
	StartTime           bool
	Time3               bool
	Time4               bool
}

func (s SectorStepScanCIF) Size() uint32 {
	return sectorStepScanCIFBytes
}

func (s *SectorStepScanCIF) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	// Initialize all bits to zero
	var bits uint32

	// Set the bits according to the boolean values
	if s.SectorNumber {
		bits |= 1 << 31 // Bit 31
	}
	if s.F1StartFrequency {
		bits |= 1 << 30 // Bit 30
	}
	if s.F2StartFrequency {
		bits |= 1 << 29 // Bit 29
	}
	if s.ResolutionBandwidth {
		bits |= 1 << 28 // Bit 28
	}
	if s.TuneStepSize {
		bits |= 1 << 27 // Bit 27
	}
	if s.NumberPoints {
		bits |= 1 << 26 // Bit 26
	}
	if s.DefaultGain {
		bits |= 1 << 25 // Bit 25
	}
	if s.Threshold {
		bits |= 1 << 24 // Bit 24
	}
	if s.DwellTime {
		bits |= 1 << 23 // Bit 23
	}
	if s.StartTime {
		bits |= 1 << 22 // Bit 22
	}
	if s.Time3 {
		bits |= 1 << 21 // Bit 21
	}
	if s.Time4 {
		bits |= 1 << 20 // Bit 20
	}

	// Store the packed bits in retval
	binary.BigEndian.PutUint32(retval, bits)

	return retval
}

func (s *SectorStepScanCIF) Unpack(buf []byte) {
	// Read the packed bits
	bits := binary.BigEndian.Uint32(buf)
	// Unpack each boolean field based on its bit position
	s.SectorNumber = (bits & (1 << 31)) != 0        // Bit 31
	s.F1StartFrequency = (bits & (1 << 30)) != 0    // Bit 30
	s.F2StartFrequency = (bits & (1 << 29)) != 0    // Bit 29
	s.ResolutionBandwidth = (bits & (1 << 28)) != 0 // Bit 28
	s.TuneStepSize = (bits & (1 << 27)) != 0        // Bit 27
	s.NumberPoints = (bits & (1 << 26)) != 0        // Bit 26
	s.DefaultGain = (bits & (1 << 25)) != 0         // Bit 25
	s.Threshold = (bits & (1 << 24)) != 0           // Bit 24
	s.DwellTime = (bits & (1 << 23)) != 0           // Bit 23
	s.StartTime = (bits & (1 << 22)) != 0           // Bit 22
	s.Time3 = (bits & (1 << 21)) != 0               // Bit 21
	s.Time4 = (bits & (1 << 20)) != 0               // Bit 20
}

// SectorStepScanRecord
type SectorStepScanRecord struct {
	SectorNumber        uint32
	F1StartFrequency    uint64
	F2StopFrequency     uint64
	ResolutionBandwidth uint64
	TuneStepSize        uint64
	NumberPoints        uint32
	DefaultGain         Gain
	Threshold           Gain
	DwellTime           uint32
	StartTime           uint32
	Time3               uint32
	Time4               uint32
}

func (s SectorStepScanRecord) Size() uint32 {
	return sectorStepScanRecordBytes
}

func (s *SectorStepScanRecord) Pack() []byte {
	buf1 := make([]byte, 28)
	binary.BigEndian.PutUint32(buf1[0:], s.Time4)
	binary.BigEndian.PutUint32(buf1[8:], s.Time3)
	binary.BigEndian.PutUint32(buf1[16:], s.StartTime)
	binary.BigEndian.PutUint32(buf1[24:], s.DwellTime)
	// Appending the values this way sucks,
	// but the go compiler thinks these .Pack() functions return an int and won't accept using copy()
	buf1 = append(buf1, s.Threshold.Pack()...)
	buf1 = append(buf1, s.DefaultGain.Pack()...)
	buf2 := make([]byte, 32)
	binary.BigEndian.PutUint32(buf2[0:], s.NumberPoints)
	binary.BigEndian.PutUint64(buf2[4:], s.TuneStepSize)
	binary.BigEndian.PutUint64(buf2[12:], s.ResolutionBandwidth)
	binary.BigEndian.PutUint64(buf2[20:], s.F2StopFrequency)
	binary.BigEndian.PutUint64(buf2[28:], s.F1StartFrequency)
	binary.BigEndian.PutUint32(buf2[36:], s.SectorNumber)
	retval := append(buf1, buf2...)
	return retval
}

func (s *SectorStepScanRecord) Unpack(buf []byte) {

	// Unpack the SectorNumber (32 bits)
	s.SectorNumber = binary.BigEndian.Uint32(buf[76:])

	// Unpack F1StartFrequency (64 bits)
	s.F1StartFrequency = binary.BigEndian.Uint64(buf[68:])

	// Unpack F2StartFrequency (64 bits)
	s.F2StopFrequency = binary.BigEndian.Uint64(buf[60:])

	// Unpack ResolutionBandwidth (64 bits)
	s.ResolutionBandwidth = binary.BigEndian.Uint64(buf[52:])

	// Unpack TuneStepSize (64 bits)
	s.TuneStepSize = binary.BigEndian.Uint64(buf[44:])

	// Unpack NumberPoints (32 bits)
	s.NumberPoints = binary.BigEndian.Uint32(buf[40:])

	// Unpack DefaultGain (32 bits)
	s.DefaultGain.Unpack(buf[36:40])

	// Unpack Threshold (32 bits)
	s.Threshold.Unpack(buf[32:36])

	// Unpack DwellTime (32 bits)
	s.DwellTime = binary.BigEndian.Uint32(buf[24:])

	// Unpack StartTime (32 bits)
	s.StartTime = binary.BigEndian.Uint32(buf[16:])

	// Unpack Time3 (32 bits)
	s.Time3 = binary.BigEndian.Uint32(buf[8:])

	// Unpack Time4 (32 bits)
	s.Time4 = binary.BigEndian.Uint32(buf[0:])

}

// SectorStepScan
type SectorStepScan struct {
	ArraySize      uint32
	HeaderSize     uint8
	NumWordsRecord uint16
	NumRecords     uint16
	SubfieldCif    SectorStepScanCIF
	Records        []SectorStepScanRecord
}

func (s SectorStepScan) Size() uint32 {
	return sectorStepScanBytes
}

func (s *SectorStepScan) Pack() []byte {
	// 3 words + the number of words required for each record (NumRecords)
	retval := make([]byte, 80+(4*int(s.NumRecords)))

	// Array Size - 4 bytes
	arraySizeWord := uint32(0)
	arraySizeWord |= uint32(s.ArraySize)
	binary.BigEndian.PutUint32(retval[0:], arraySizeWord) // word 1 in the buffer

	// Header Size/NumWordsRecord/NumRecords
	headerSizeWord := uint32(0)
	headerSizeWord |= (uint32(s.HeaderSize) << 24)
	headerSizeWord |= (uint32(s.NumWordsRecord) << 12)
	headerSizeWord |= uint32(s.NumRecords)
	binary.BigEndian.PutUint32(retval[32:], headerSizeWord) // word 2 in the buffer

	// CIF
	cifWord := uint32(0)
	cifWord |= binary.BigEndian.Uint32(s.SubfieldCif.Pack())
	binary.BigEndian.PutUint32(retval[63:], cifWord) // word 3 in the buffer

	// Append N records to the slice - 8 bytes each record
	for _, record := range s.Records {
		retval = append(retval, record.Pack()...)
	}

	return retval
}

func (s *SectorStepScan) Unpack(buf []byte) {

	s.ArraySize = binary.BigEndian.Uint32(buf[0:4])
	fmt.Printf("Unpacking header word:\n")

	headerSizeWord := binary.BigEndian.Uint32(buf[32:36])

	fmt.Printf("Unpacking header size:\n")
	s.HeaderSize = uint8(headerSizeWord >> 24)

	fmt.Printf("Unpacking num words record:\n")
	s.NumWordsRecord = uint16((headerSizeWord >> 12) & 0xFFF)

	fmt.Printf("Unpacking num records:\n")
	s.NumRecords = uint16(headerSizeWord & 0xFFF)

	fmt.Printf("Unpacking records:\n")

	s.SubfieldCif.Unpack(buf[63:])

	// Unpack Entries (each entry is 4 bytes)
	s.Records = make([]SectorStepScanRecord, s.NumRecords)

	for i, record := range s.Records {
		record.Unpack(buf[79*(i*4):]) // Each record is 4 bytes
		s.Records[i] = record
	}
}

// IndexList
// Supports records containing an Index subfield.
// Specifies the indicies of the records (which could be a subset of the entire
// collection) upon which a device is to act.
type IndexList struct {
	TotalSize  uint32
	EntrySize  uint8
	NumEntries uint32
	Entries    []uint32
}

func (s IndexList) Size() uint32 {
	if s.NumEntries > 0 {
		return uint32(32+32+4*s.NumEntries) / 8
	}
	return uint32(8) // TotalSize + EntrySize and NumEntries
}

func (s *IndexList) Pack() []byte {
	totalSize := 8 + 4*int(s.NumEntries)
	retval := make([]byte, totalSize)

	// Total Size - 4 bytes
	binary.BigEndian.PutUint32(retval[0:], s.TotalSize)

	// Entry Size  - 4 bytes
	secondWord := uint32(s.EntrySize) << 28
	// Num Entries - 20 bits
	secondWord |= (s.NumEntries & 0x000FFFFF)

	binary.BigEndian.PutUint32(retval[4:], secondWord)

	// Entries - 1 byte each
	for i, entry := range s.Entries {
		binary.BigEndian.PutUint32(retval[8+i*4:], entry)
	}

	return retval
}

func (s *IndexList) Unpack(buf []byte) {
	// Unpack TotalSize (4 bytes)
	s.TotalSize = binary.BigEndian.Uint32(buf[0:4])

	secondWord := binary.BigEndian.Uint32(buf[4:8])
	s.EntrySize = uint8(secondWord >> 28)
	s.NumEntries = secondWord & 0x000FFFFF

	// Determine the number of entries
	numEntries := int(s.NumEntries)
	// Unpack Entries (each entry is 4 bytes)
	s.Entries = make([]uint32, numEntries)
	for i := 0; i < numEntries; i++ {
		s.Entries[i] = binary.BigEndian.Uint32(buf[8+i*4:])
	}
}

// VersionInformation
type VersionInformation struct {
	Year        uint8
	Day         uint16
	Revision    uint8
	UserDefined uint16
}

func (s VersionInformation) Size() uint32 {
	return versionInformationBytes
}

func (s *VersionInformation) Pack() []byte {
	// 1 word (4 bytes)
	retval := make([]byte, 4)

	// Create a 32-bit word (uint32)
	var word uint32

	// Pack the year into bits 25-31
	// binary.BigEndian.PutUint16(retval[0:], uint16(s.Year))
	year := uint32(s.Year) & 0x7F
	word |= year << 25

	// Pack the day into bits 16-24
	day := uint32(s.Day) & 0x1FF
	word |= day << 16

	// Pack the revision into bits 10-15
	revision := uint32(s.Revision) & 0x3F
	word |= revision << 10

	// Pack the user-defined value into bits 0-9
	userDefined := uint32(s.UserDefined) & 0x3FF
	word |= userDefined

	// Convert the 32-bit word to a byte slice
	binary.BigEndian.PutUint32(retval, word)

	return retval
}

func (s *VersionInformation) Unpack(buf []byte) {

	word := binary.BigEndian.Uint32(buf)

	s.Year = uint8((word >> 25) & 0x7F)

	s.Day = uint16((word >> 16) & 0x1FF)

	s.Revision = uint8((word >> 10) & 0x3F)

	s.UserDefined = uint16(word & 0x3FF)
}
