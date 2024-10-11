package vita49

import (
	"bytes"
	"encoding/binary"
)

const (
	gainBytes             = uint32(4)
	deviceIdentifierBytes = uint32(8)
	ephemerisBytes        = uint32(52)
	geolocationBytes      = uint32(44)
	payloadFormatBytes    = uint32(8)
)

type Cif0 struct {
	IndicatorField0
}

// Gain
type Gain struct {
	Stage1 float64
	Stage2 float64
}

func (g Gain) Size() uint32 {
	return gainBytes
}

func (g *Gain) Pack() []byte {
	buf := make([]byte, g.Size())
	binary.BigEndian.PutUint16(buf[2:], uint16(ToFixed16(g.Stage1, 7)))
	binary.BigEndian.PutUint16(buf[0:], uint16(ToFixed16(g.Stage2, 7)))
	return buf
}

func (g *Gain) Unpack(buf []byte) {
	g.Stage1 = FromFixed(int16(binary.BigEndian.Uint16(buf[2:])), 7)
	g.Stage2 = FromFixed(int16(binary.BigEndian.Uint16(buf[0:])), 7)
}

// Device ID
type DeviceIdentifier struct {
	ManufacturerOui uint32
	DeviceCode      uint16
}

func (d DeviceIdentifier) Size() uint32 {
	return deviceIdentifierBytes
}

func (d *DeviceIdentifier) Pack() []byte {
	buf := make([]byte, d.Size())
	d.ManufacturerOui &= 0x00FFFFFF
	binary.BigEndian.PutUint32(buf[0:], d.ManufacturerOui)
	binary.BigEndian.PutUint16(buf[4:], 0) // reserved bits
	binary.BigEndian.PutUint16(buf[6:], d.DeviceCode)
	return buf
}

func (d *DeviceIdentifier) Unpack(buf []byte) {
	d.ManufacturerOui = binary.BigEndian.Uint32(buf[0:])
	d.DeviceCode = binary.BigEndian.Uint16(buf[6:])
}

// Ephemeris
type Ephemeris struct {
	Tsi                 uint8
	Tsf                 uint8
	ManufacturerOui     uint32
	IntegerTimestamp    uint32
	FractionalTimestamp uint64
	PositionX           float64
	PositionY           float64
	PositionZ           float64
	AttitudeAlpha       float64
	AttitudeBeta        float64
	AttitudePhi         float64
	VelocityDx          float64
	VelocityDy          float64
	VelocityDz          float64
}

func (e Ephemeris) Size() uint32 {
	return ephemerisBytes
}

func NewEphemeris() Ephemeris {
	return Ephemeris{
		Tsi:                 0,
		Tsf:                 0,
		ManufacturerOui:     0,
		IntegerTimestamp:    ^uint32(0),
		FractionalTimestamp: ^uint64(0),
		PositionX:           FromFixed(int32(0x7FFFFFFF), 5),
		PositionY:           FromFixed(int32(0x7FFFFFFF), 5),
		PositionZ:           FromFixed(int32(0x7FFFFFFF), 5),
		AttitudeAlpha:       FromFixed(int32(0x7FFFFFFF), 22),
		AttitudeBeta:        FromFixed(int32(0x7FFFFFFF), 22),
		AttitudePhi:         FromFixed(int32(0x7FFFFFFF), 22),
		VelocityDx:          FromFixed(int32(0x7FFFFFFF), 16),
		VelocityDy:          FromFixed(int32(0x7FFFFFFF), 16),
		VelocityDz:          FromFixed(int32(0x7FFFFFFF), 16),
	}
}

func (e *Ephemeris) Pack(buf []byte) {
	word1 := uint32(0)
	word1 |= uint32((e.Tsi & 3)) << 26
	word1 |= uint32((e.Tsf & 3)) << 24
	word1 |= uint32((e.ManufacturerOui & 0xFFFFFF))
	binary.BigEndian.PutUint32(buf[0:], word1)
	binary.BigEndian.PutUint32(buf[4:], e.IntegerTimestamp)
	binary.BigEndian.PutUint64(buf[8:], e.FractionalTimestamp)
	binary.BigEndian.PutUint32(buf[16:], uint32(ToFixed32(e.PositionX, 5)))
	binary.BigEndian.PutUint32(buf[20:], uint32(ToFixed32(e.PositionY, 5)))
	binary.BigEndian.PutUint32(buf[24:], uint32(ToFixed32(e.PositionZ, 5)))
	binary.BigEndian.PutUint32(buf[28:], uint32(ToFixed32(e.AttitudeAlpha, 22)))
	binary.BigEndian.PutUint32(buf[32:], uint32(ToFixed32(e.AttitudeBeta, 22)))
	binary.BigEndian.PutUint32(buf[36:], uint32(ToFixed32(e.AttitudePhi, 22)))
	binary.BigEndian.PutUint32(buf[40:], uint32(ToFixed32(e.VelocityDx, 16)))
	binary.BigEndian.PutUint32(buf[44:], uint32(ToFixed32(e.VelocityDy, 16)))
	binary.BigEndian.PutUint32(buf[48:], uint32(ToFixed32(e.VelocityDz, 16)))
}

func (e *Ephemeris) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf[0:])
	e.Tsi = uint8((word1 >> 26) & 3)
	e.Tsf = uint8((word1 >> 24) & 3)
	e.ManufacturerOui = uint32((word1 & 0xFFFFFF))
	e.IntegerTimestamp = binary.BigEndian.Uint32(buf[4:])
	e.FractionalTimestamp = binary.BigEndian.Uint64(buf[8:])
	e.PositionX = FromFixed(int32(binary.BigEndian.Uint32(buf[16:])), 5)
	e.PositionY = FromFixed(int32(binary.BigEndian.Uint32(buf[20:])), 5)
	e.PositionZ = FromFixed(int32(binary.BigEndian.Uint32(buf[24:])), 5)
	e.AttitudeAlpha = FromFixed(int32(binary.BigEndian.Uint32(buf[28:])), 22)
	e.AttitudeBeta = FromFixed(int32(binary.BigEndian.Uint32(buf[32:])), 22)
	e.AttitudePhi = FromFixed(int32(binary.BigEndian.Uint32(buf[36:])), 22)
	e.VelocityDx = FromFixed(int32(binary.BigEndian.Uint32(buf[40:])), 16)
	e.VelocityDy = FromFixed(int32(binary.BigEndian.Uint32(buf[44:])), 16)
	e.VelocityDz = FromFixed(int32(binary.BigEndian.Uint32(buf[48:])), 16)
}

// Geolocation
type Geolocation struct {
	Tsi                 uint8
	Tsf                 uint8
	ManufacturerOui     uint32
	IntegerTimestamp    uint32
	FractionalTimestamp uint64
	Latitude            float64
	Longitude           float64
	Altitude            float64
	SpeedOverGround     float64
	HeadingAngle        float64
	TrackAngle          float64
	MagneticVariation   float64
}

func (g Geolocation) Size() uint32 {
	return geolocationBytes
}

func NewGeolocation() Geolocation {
	return Geolocation{
		Tsi:                 0,
		Tsf:                 0,
		ManufacturerOui:     0,
		IntegerTimestamp:    ^uint32(0),
		FractionalTimestamp: ^uint64(0),
		Latitude:            FromFixed(int32(0x7FFFFFFF), 22),
		Longitude:           FromFixed(int32(0x7FFFFFFF), 22),
		Altitude:            FromFixed(int32(0x7FFFFFFF), 5),
		SpeedOverGround:     FromFixed(int32(0x7FFFFFFF), 16),
		HeadingAngle:        FromFixed(int32(0x7FFFFFFF), 22),
		TrackAngle:          FromFixed(int32(0x7FFFFFFF), 22),
		MagneticVariation:   FromFixed(int32(0x7FFFFFFF), 22),
	}
}

func (g *Geolocation) Pack(buf []byte) {
	word1 := uint32(0)
	word1 |= uint32((g.Tsi & 3)) << 26
	word1 |= uint32((g.Tsf & 3)) << 24
	word1 |= uint32((g.ManufacturerOui & 0xFFFFFF))

	binary.BigEndian.PutUint32(buf[0:], word1)
	binary.BigEndian.PutUint32(buf[4:], g.IntegerTimestamp)
	binary.BigEndian.PutUint64(buf[8:], g.FractionalTimestamp)
	binary.BigEndian.PutUint32(buf[16:], uint32(ToFixed32(g.Latitude, 22)))
	binary.BigEndian.PutUint32(buf[20:], uint32(ToFixed32(g.Longitude, 22)))
	binary.BigEndian.PutUint32(buf[24:], uint32(ToFixed32(g.Altitude, 5)))
	binary.BigEndian.PutUint32(buf[28:], uint32(ToFixed32(g.SpeedOverGround, 16)))
	binary.BigEndian.PutUint32(buf[32:], uint32(ToFixed32(g.HeadingAngle, 22)))
	binary.BigEndian.PutUint32(buf[36:], uint32(ToFixed32(g.TrackAngle, 22)))
	binary.BigEndian.PutUint32(buf[40:], uint32(ToFixed32(g.MagneticVariation, 22)))

}

func (g *Geolocation) Unpack(buf []byte) {
	word1 := binary.BigEndian.Uint32(buf[0:])
	g.Tsi = uint8((word1 >> 26) & 3)
	g.Tsf = uint8((word1 >> 24) & 3)
	g.ManufacturerOui = uint32((word1 & 0xFFFFFF))
	g.IntegerTimestamp = binary.BigEndian.Uint32(buf[4:])
	g.FractionalTimestamp = binary.BigEndian.Uint64(buf[8:])
	g.Latitude = FromFixed(int32(binary.BigEndian.Uint32(buf[16:])), 22)
	g.Longitude = FromFixed(int32(binary.BigEndian.Uint32(buf[20:])), 22)
	g.Altitude = FromFixed(int32(binary.BigEndian.Uint32(buf[24:])), 5)
	g.SpeedOverGround = FromFixed(int32(binary.BigEndian.Uint32(buf[28:])), 16)
	g.HeadingAngle = FromFixed(int32(binary.BigEndian.Uint32(buf[32:])), 22)
	g.TrackAngle = FromFixed(int32(binary.BigEndian.Uint32(buf[36:])), 22)
	g.MagneticVariation = FromFixed(int32(binary.BigEndian.Uint32(buf[40:])), 22)
}

// GPS ASCII
type GpsAscii struct {
	ManufacturerOui uint32
	NumberOfWords   uint32
	AsciiSentences  []uint8
}

func (g *GpsAscii) Size() uint32 {
	return uint32(8 + g.NumberOfWords*4)

}

func NewGpsAscii() GpsAscii {
	return GpsAscii{
		ManufacturerOui: 0,
		NumberOfWords:   0,
		AsciiSentences:  []uint8{},
	}
}
func (g *GpsAscii) Pack(buf []byte) {
	binary.BigEndian.PutUint32(buf[0:], g.ManufacturerOui)
	buf[0] = 0x00
	binary.BigEndian.PutUint32(buf[4:], g.NumberOfWords)
	copy(buf[8:], g.AsciiSentences)
	// Add padding when number of ascii bytes not divis by 4
	padding := 4*g.NumberOfWords - uint32(len(g.AsciiSentences))
	copy(buf[8+len(g.AsciiSentences):], make([]byte, padding))
}

func (g *GpsAscii) Unpack(buf []byte) {
	g.ManufacturerOui = binary.BigEndian.Uint32(buf[0:]) & 0x00FFFFFF
	g.NumberOfWords = binary.BigEndian.Uint32(buf[4:])
	g.AsciiSentences = buf[8:]
}

// Payload Format
type PayloadFormat struct {
	PackingMethod        bool
	RealComplexType      uint8
	DataItemFormat       uint8
	RepeatIndicator      bool
	EventTagSize         uint8
	ChannelTagSize       uint8
	DataItemFractionSize uint8
	ItemPackingFieldSize uint8
	DataItemSize         uint8
	RepeatCount          uint32
	VectorSize           uint32
}

func NewPayloadFormat() PayloadFormat {
	return PayloadFormat{
		ItemPackingFieldSize: 1,
		DataItemSize:         1,
		RepeatCount:          1,
		VectorSize:           1,
	}
}

func (p PayloadFormat) Size() uint32 {
	return payloadFormatBytes
}
func (p *PayloadFormat) Pack(buf []byte) {
	word1 := uint32(0)
	packing := uint8(0)
	repeatindicator := uint8(0)
	if p.PackingMethod {
		packing = 1
	}
	if p.RepeatIndicator {
		repeatindicator = 1
	}
	word1 |= uint32((packing & 1)) << 31
	word1 |= uint32((p.RealComplexType & 3)) << 29
	word1 |= uint32((p.DataItemFormat & 31)) << 24
	word1 |= uint32((repeatindicator & 1)) << 23
	word1 |= uint32((p.EventTagSize & 7)) << 20
	word1 |= uint32((p.ChannelTagSize & 15)) << 16
	word1 |= uint32((p.DataItemFractionSize & 15)) << 12
	if p.ItemPackingFieldSize > 0 {
		word1 |= uint32((p.ItemPackingFieldSize-1)&63) << 6
	} else {
		word1 |= uint32(p.ItemPackingFieldSize&63) << 6
	}
	if p.DataItemSize > 0 {
		word1 |= uint32(p.DataItemSize - 1)
	} else {
		word1 |= uint32(p.DataItemSize)
	}
	binary.BigEndian.PutUint32(buf[0:], word1)

	if p.RepeatCount > 0 {
		binary.BigEndian.PutUint16(buf[4:], uint16(p.RepeatCount-1))
	} else {
		binary.BigEndian.PutUint16(buf[4:], 0)
	}

	if p.VectorSize > 0 {
		binary.BigEndian.PutUint16(buf[6:], uint16(p.VectorSize-1))
	} else {
		binary.BigEndian.PutUint16(buf[6:], 0)
	}

}

func (p *PayloadFormat) Unpack(buf []byte) {
	vectorSize := uint32(binary.BigEndian.Uint16(buf[6:]))
	if vectorSize < 0xFFFF {
		p.VectorSize = vectorSize + 1
	} else {
		p.VectorSize = 0x10000
	}

	repeatCount := uint32(binary.BigEndian.Uint16(buf[4:]))
	if repeatCount < 0xFFFF {
		p.RepeatCount = repeatCount + 1
	} else {
		p.RepeatCount = 0x10000
	}
	word1 := binary.BigEndian.Uint32(buf[0:])
	p.PackingMethod = ((word1 >> 31) & 1) != 0
	p.RealComplexType = uint8((word1 >> 29) & 3)
	p.DataItemFormat = uint8((word1 >> 24) & 31)
	p.RepeatIndicator = ((word1 >> 23) & 1) != 0
	p.EventTagSize = uint8((word1 >> 20) & 7)
	p.ChannelTagSize = uint8((word1 >> 16) & 15)
	p.DataItemFractionSize = uint8((word1 >> 12) & 15)

	itemPackingFieldSize := uint8((word1 >> 6) & 63)
	if itemPackingFieldSize < 64 {
		p.ItemPackingFieldSize = itemPackingFieldSize + 1
	} else {
		p.ItemPackingFieldSize = 64
	}

	dataItemSize := uint8(word1 & 63)
	if dataItemSize < 64 {
		p.DataItemSize = dataItemSize + 1
	} else {
		p.DataItemSize = 64
	}
}

// Context Association Lists
type ContextAssociationLists struct {
	SourceListSize     uint8
	SystemListSize     uint8
	VectorListSize     uint16
	AsyncTagListEnable bool
	AsyncListSize      uint16
	SourceList         []uint32
	SystemList         []uint32
	VectorList         []uint32
	AsyncList          []uint32
	AsyncTagList       []uint32
}

func (c *ContextAssociationLists) Size() uint32 {
	asyncTag := uint16(0)
	if c.AsyncTagListEnable {
		asyncTag = 1
	}
	return 8 + 4*(uint32(c.SourceListSize)+
		uint32(c.SystemListSize)+
		uint32(c.VectorListSize)+
		uint32(c.AsyncListSize)+
		uint32(c.AsyncListSize*asyncTag))
}

func NewContextAssociationLists() ContextAssociationLists {
	return ContextAssociationLists{
		SourceList:   []uint32{},
		SystemList:   []uint32{},
		VectorList:   []uint32{},
		AsyncList:    []uint32{},
		AsyncTagList: []uint32{},
	}
}

func (c *ContextAssociationLists) Pack(buf []byte) {
	asyncTag := uint16(0)
	offset := uint16(0)

	word1 := uint32(0)
	word1 |= uint32(c.SourceListSize) << 16
	word1 |= uint32(c.SystemListSize)
	binary.BigEndian.PutUint32(buf[0:], word1)

	word2 := uint32(0)
	word2 |= uint32(c.AsyncListSize & 0x7FFF)
	if c.AsyncTagListEnable {
		asyncTag = 1
	}
	word2 |= uint32(asyncTag&1) << 15
	word2 |= uint32(c.VectorListSize) << 16
	binary.BigEndian.PutUint32(buf[4:], word2)

	copy(buf[8:], uint32ToBytes(c.SourceList))
	offset += 8 + uint16(c.SourceListSize*4)
	copy(buf[offset:], uint32ToBytes(c.SystemList))
	offset += uint16(c.SystemListSize * 4)
	copy(buf[offset:], uint32ToBytes(c.VectorList))
	offset += c.VectorListSize * 4
	copy(buf[offset:], uint32ToBytes(c.AsyncList))
	offset += c.AsyncListSize * 4
	copy(buf[offset:offset+(c.AsyncListSize*asyncTag)], uint32ToBytes(c.AsyncTagList))
}

func (c *ContextAssociationLists) Unpack(buf []byte) {
	offset := uint16(8)
	nextOffset := uint16(0)
	word1 := binary.BigEndian.Uint32(buf[0:])
	c.SystemListSize = uint8(word1 & 0xFF)
	c.SourceListSize = uint8(word1>>16) & 0xFF
	word2 := binary.BigEndian.Uint32(buf[4:])
	c.VectorListSize = uint16(word2 >> 16)
	c.AsyncListSize = uint16(word2) & 0x7FFF
	asyncTag := uint16(word2&0x8000) >> 15
	c.AsyncTagListEnable = asyncTag != 0 // Eval as boolean

	nextOffset = offset + uint16(c.SourceListSize*4)
	c.SourceList = bytesToUint32Slice(buf[offset:nextOffset])
	offset = nextOffset
	nextOffset = offset + uint16(c.SystemListSize*4)
	c.SystemList = bytesToUint32Slice(buf[offset:nextOffset])
	offset = nextOffset
	nextOffset = offset + c.VectorListSize*4
	c.VectorList = bytesToUint32Slice(buf[offset:nextOffset])
	offset = nextOffset
	nextOffset = offset + c.AsyncListSize*4
	c.AsyncList = bytesToUint32Slice(buf[offset:nextOffset])
	offset = nextOffset
	nextOffset = offset + c.AsyncListSize*4*asyncTag // If enable tag is false, list is not stored
	c.AsyncTagList = bytesToUint32Slice(buf[offset:nextOffset])
}

// Convert Context Assocation lists into byte slices
func uint32ToBytes(list []uint32) []byte {
	var buf bytes.Buffer
	for _, i := range list {
		binary.Write(&buf, binary.BigEndian, i)
	}
	return buf.Bytes()
}

// Byte slices into CA lists
func bytesToUint32Slice(buf []byte) []uint32 {
	uint32Len := len(buf) / 4
	uint32Slice := make([]uint32, uint32Len)

	for i := 0; i < uint32Len; i++ {
		uint32Slice[i] = uint32(buf[i*4]) |
			(uint32(buf[i*4+1]) << 8) |
			(uint32(buf[i*4+2]) << 16) |
			(uint32(buf[i*4+3]) << 24)
	}

	return uint32Slice
}
