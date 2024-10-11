package vita49

import "encoding/binary"

const (
	classIdBytes = uint32(8)
)

type ClassId struct {
	PadBitCount     uint8
	Oui             uint32 `yaml:"oui,omitempty"`
	InformationCode uint16 `yaml:"informationCode,omitempty"`
	PacketCode      uint16 `yaml:"packetCode,omitempty"`
}

func (c ClassId) Size() uint32 {
	return classIdBytes
}

func (c *ClassId) Pack() []byte {
	buf := make([]byte, c.Size())
	var word1 uint32
	word1 = (uint32(c.PadBitCount&0x1F) << 27) | (c.Oui & 0x00FFFFFF)
	binary.BigEndian.PutUint32(buf[0:], word1)
	binary.BigEndian.PutUint16(buf[4:], c.InformationCode)
	binary.BigEndian.PutUint16(buf[6:], c.PacketCode)
	return buf
}

func (c *ClassId) Unpack(buf []byte) {
	c.PadBitCount = uint8(binary.BigEndian.Uint32(buf[0:]) >> 27)
	c.Oui = binary.BigEndian.Uint32(buf[0:]) & 0x00FFFFFF
	c.InformationCode = binary.BigEndian.Uint16(buf[4:])
	c.PacketCode = binary.BigEndian.Uint16(buf[6:])
}
