package vita49

import "encoding/binary"

type Cif2 struct {
	IndicatorField2
	Bind                    uint32
	CitedSid                uint32
	SiblingSid              uint32
	ParentSid               uint32
	ChildSid                uint32
	CitedMessageId          uint32
	ControlleeId            uint32
	ControlleeUuid          uint32
	ControllerId            uint32
	ControllerUuid          uint32
	InformationSource       uint32
	TrackId                 uint32
	CountryCode             uint32
	Operator                uint32
	PlatformClass           uint32
	PlatformInstance        uint32
	PlatformDisplay         uint32
	EMSDeviceClass          uint32
	EMSDeviceType           uint32
	EMSDeviceInstance       uint32
	ModulationClass         uint32
	ModulationType          uint32
	FunctionId              uint32
	ModeId                  uint32
	EventId                 uint32
	FunctionPriorityId      uint32
	CommunicationPriorityId uint32
	RfFootprint             uint32
	RfFootprintRange        uint32
}

func (c Cif2) Size() uint32 {
	return c.IndicatorField2.Size()
}

func NewCif2() *Cif2 {
	return &Cif2{}
}

func (s *Cif2) Pack() []byte {
	// 1 word
	retval := make([]byte, 4)

	// Initialize all bits to zero
	var bits uint32
	if s.IndicatorField2.Bind {
		bits |= 1 << 31 // Bit 31
	}
	if s.IndicatorField2.CitedSid {
		bits |= 1 << 30 // Bit 30
	}
	if s.IndicatorField2.SiblingSid {
		bits |= 1 << 29 // Bit 29
	}
	if s.IndicatorField2.ParentSid {
		bits |= 1 << 28 // Bit 28
	}
	if s.IndicatorField2.ChildSid {
		bits |= 1 << 27 // Bit 27
	}
	if s.IndicatorField2.CitedMessageId {
		bits |= 1 << 26 // Bit 26
	}
	if s.IndicatorField2.ControlleeId {
		bits |= 1 << 25 // Bit 25
	}
	if s.IndicatorField2.ControlleeUuid {
		bits |= 1 << 24 // Bit 24
	}
	if s.IndicatorField2.ControllerId {
		bits |= 1 << 23 // Bit 23
	}
	if s.IndicatorField2.ControllerUuid {
		bits |= 1 << 22 // Bit 22
	}
	if s.IndicatorField2.InformationSource {
		bits |= 1 << 21 // Bit 21
	}
	if s.IndicatorField2.TraceId {
		bits |= 1 << 20 // Bit 20
	}
	if s.IndicatorField2.CountryCode {
		bits |= 1 << 19 // Bit 19
	}
	if s.IndicatorField2.Operator {
		bits |= 1 << 18 // Bit 18
	}
	if s.IndicatorField2.PlatformClass {
		bits |= 1 << 17 // Bit 17
	}
	if s.IndicatorField2.PlatformInstance {
		bits |= 1 << 16 // Bit 16
	}
	if s.IndicatorField2.PlatformDisplay {
		bits |= 1 << 15 // Bit 15
	}
	if s.IndicatorField2.EmsDeviceClass {
		bits |= 1 << 14 // Bit 14
	}
	if s.IndicatorField2.EmsDeviceType {
		bits |= 1 << 13 // Bit 13
	}
	if s.IndicatorField2.EmsDeviceInstance {
		bits |= 1 << 12 // Bit 12
	}
	if s.IndicatorField2.ModulationClass {
		bits |= 1 << 11 // Bit 11
	}
	if s.IndicatorField2.ModulationType {
		bits |= 1 << 10 // Bit 10
	}
	if s.IndicatorField2.FunctionId {
		bits |= 1 << 9 // Bit 9
	}
	if s.IndicatorField2.ModeId {
		bits |= 1 << 8 // Bit 8
	}
	if s.IndicatorField2.EventId {
		bits |= 1 << 7 // Bit 7
	}
	if s.IndicatorField2.FunctionPriorityId {
		bits |= 1 << 6 // Bit 6
	}
	if s.IndicatorField2.CommunicationPriorityId {
		bits |= 1 << 5 // Bit 5
	}
	if s.IndicatorField2.RfFootprint {
		bits |= 1 << 4 // Bit 4
	}
	if s.IndicatorField2.RfFootprintRange {
		bits |= 1 << 3 // Bit 3
	}

	// Store the packed bits in retval
	binary.BigEndian.PutUint32(retval, bits)

	return retval
}

func (s *Cif2) Unpack(buf []byte) {
	// Read the packed bits
	bits := binary.BigEndian.Uint32(buf)
	// Unpack each boolean field based on its bit position
	s.IndicatorField2.Bind = (bits & (1 << 31)) != 0                   // Bit 31
	s.IndicatorField2.CitedSid = (bits & (1 << 30)) != 0               // Bit 30
	s.IndicatorField2.SiblingSid = (bits & (1 << 29)) != 0             // Bit 29
	s.IndicatorField2.ParentSid = (bits & (1 << 28)) != 0              // Bit 28
	s.IndicatorField2.ChildSid = (bits & (1 << 27)) != 0               // Bit 27
	s.IndicatorField2.CitedMessageId = (bits & (1 << 26)) != 0         // Bit 26
	s.IndicatorField2.ControlleeId = (bits & (1 << 25)) != 0           // Bit 25
	s.IndicatorField2.ControlleeUuid = (bits & (1 << 24)) != 0         // Bit 24
	s.IndicatorField2.ControllerId = (bits & (1 << 23)) != 0           // Bit 23
	s.IndicatorField2.ControllerUuid = (bits & (1 << 22)) != 0         // Bit 22
	s.IndicatorField2.InformationSource = (bits & (1 << 21)) != 0      // Bit 21
	s.IndicatorField2.TraceId = (bits & (1 << 20)) != 0                // Bit 20
	s.IndicatorField2.CountryCode = (bits & (1 << 19)) != 0            // Bit 19
	s.IndicatorField2.Operator = (bits & (1 << 18)) != 0               // Bit 18
	s.IndicatorField2.PlatformClass = (bits & (1 << 17)) != 0          // Bit 17
	s.IndicatorField2.PlatformInstance = (bits & (1 << 16)) != 0       // Bit 16
	s.IndicatorField2.PlatformDisplay = (bits & (1 << 15)) != 0        // Bit 15
	s.IndicatorField2.EmsDeviceClass = (bits & (1 << 14)) != 0         // Bit 14
	s.IndicatorField2.EmsDeviceType = (bits & (1 << 13)) != 0          // Bit 13
	s.IndicatorField2.EmsDeviceInstance = (bits & (1 << 12)) != 0      // Bit 12
	s.IndicatorField2.ModulationClass = (bits & (1 << 11)) != 0        // Bit 11
	s.IndicatorField2.ModulationType = (bits & (1 << 10)) != 0         // Bit 10
	s.IndicatorField2.FunctionId = (bits & (1 << 9)) != 0              // Bit 9
	s.IndicatorField2.ModeId = (bits & (1 << 8)) != 0                  // Bit 8
	s.IndicatorField2.EventId = (bits & (1 << 7)) != 0                 // Bit 7
	s.IndicatorField2.FunctionPriorityId = (bits & (1 << 6)) != 0      // Bit 6
	s.IndicatorField2.CommunicationPriorityId = (bits & (1 << 5)) != 0 // Bit 5
	s.IndicatorField2.RfFootprint = (bits & (1 << 4)) != 0             // Bit 4
	s.IndicatorField2.RfFootprintRange = (bits & (1 << 3)) != 0        // Bit 3
}
