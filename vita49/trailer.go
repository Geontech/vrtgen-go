package vita49

const (
	TrailerBits = uint8(32)
)

type Trailer struct {
	CalibratedTime    EnableIndicator
	ValidData         EnableIndicator
	ReferenceLock     EnableIndicator
	AgcMgc            EnableIndicator
	DetectedSignal    EnableIndicator
	SpectralInversion EnableIndicator
	OverRange         EnableIndicator
	SampleLoss        EnableIndicator
}

func (t *Trailer) Pack(buf []byte) {
	t.CalibratedTime.Pack(buf, 31, 19)
	t.ValidData.Pack(buf, 30, 18)
	t.ReferenceLock.Pack(buf, 29, 17)
	t.AgcMgc.Pack(buf, 28, 16)
	t.DetectedSignal.Pack(buf, 27, 15)
	t.SpectralInversion.Pack(buf, 26, 14)
	t.OverRange.Pack(buf, 25, 13)
	t.SampleLoss.Pack(buf, 24, 12)
}

func (t *Trailer) Unpack(buf []byte) {
	t.CalibratedTime.Unpack(buf, 31, 19)
	t.ValidData.Unpack(buf, 30, 18)
	t.ReferenceLock.Unpack(buf, 29, 17)
	t.AgcMgc.Unpack(buf, 28, 16)
	t.DetectedSignal.Unpack(buf, 27, 15)
	t.SpectralInversion.Unpack(buf, 26, 14)
	t.OverRange.Unpack(buf, 25, 13)
	t.SampleLoss.Unpack(buf, 24, 12)
}

func (t *Trailer) Bits() uint8 {
	return TrailerBits
}
