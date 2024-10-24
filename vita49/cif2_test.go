package vita49

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCif2Size(t *testing.T) {
	c := Cif2{}
	assert.Equal(t, uint32(4), c.Size())
}
