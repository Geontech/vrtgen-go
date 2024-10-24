package vita49

type StateEventIndicators struct {
	CalibratedTime    EnableIndicator
	ValidData         EnableIndicator
	ReferenceLock     EnableIndicator
	AgcMgc            EnableIndicator
	DetectedSignal    EnableIndicator
	SpectralInversion EnableIndicator
	OverRange         EnableIndicator
	SampleLoss        EnableIndicator
}

func (t *StateEventIndicators) Size() uint32 {
	return 4
}

func (s *StateEventIndicators) Pack() []byte {
	buf := make([]byte, s.Size())
	s.CalibratedTime.Pack(buf, 31, 19)
	s.ValidData.Pack(buf, 30, 18)
	s.ReferenceLock.Pack(buf, 29, 17)
	s.AgcMgc.Pack(buf, 28, 16)
	s.DetectedSignal.Pack(buf, 27, 15)
	s.SpectralInversion.Pack(buf, 26, 14)
	s.OverRange.Pack(buf, 25, 13)
	s.SampleLoss.Pack(buf, 24, 12)
	return buf
}

func (s *StateEventIndicators) Unpack(buf []byte) {
	s.CalibratedTime.Unpack(buf, 31, 19)
	s.ValidData.Unpack(buf, 30, 18)
	s.ReferenceLock.Unpack(buf, 29, 17)
	s.AgcMgc.Unpack(buf, 28, 16)
	s.DetectedSignal.Unpack(buf, 27, 15)
	s.SpectralInversion.Unpack(buf, 26, 14)
	s.OverRange.Unpack(buf, 25, 13)
	s.SampleLoss.Unpack(buf, 24, 12)
}
