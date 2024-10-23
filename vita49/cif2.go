package vita49

type Cif2 struct {
	IndicatorField2
}

func (c *Cif2) Size() uint32 {
	return c.IndicatorField2.Size()
}

func (c *Cif2) Pack() []byte {
	return c.IndicatorField2.Pack()
}

func (c *Cif2) Unpack(buf []byte) {
	c.IndicatorField2.Unpack(buf)
}
