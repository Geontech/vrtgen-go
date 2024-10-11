package vita49

type EnableIndicator struct {
	Enable bool
	Value  bool
}

func (ei *EnableIndicator) Reset() {
	ei.Enable = false
	ei.Value = false
}

func (ei *EnableIndicator) Pack(buf []byte, enPos uint8, inPos uint8) {
	var enableVal uint8
	var indicatorVal uint8
	if ei.Enable {
		enableVal = 1
		if ei.Value {
			indicatorVal = 1
		}
	}
	bytePos := uint8((31 - enPos) / 8)
	mask := ^(uint8(1) << (enPos % 8))
	buf[bytePos] = (buf[bytePos] & mask) | (enableVal << (enPos % 8))
	bytePos = uint8((31 - inPos) / 8)
	mask = ^(uint8(1) << (inPos % 8))
	buf[bytePos] = (buf[bytePos] & mask) | (indicatorVal << (inPos % 8))
}

func (ei *EnableIndicator) Unpack(buf []byte, enPos uint8, inPos uint8) {
	bytePos := uint8((31 - enPos) / 8)
	ei.Enable = (buf[bytePos] & (1 << (enPos % 8))) != 0
	bytePos = uint8((31 - inPos) / 8)
	ei.Value = (buf[bytePos] & (1 << (inPos % 8))) != 0
}
