package vita49

import (
	"math"
)

func ToFixed16(v float64, r uint8) int16 {
	scale := float64(uint32(1) << r)
	return int16(math.Round(v * scale))
}

func ToFixed32(v float64, r uint8) int32 {
	scale := float64(int(1) << r)
	return int32(math.Round(v * scale))
}

func ToFixed64(v float64, r uint8) int64 {
	scale := float64(uint64(1) << r)
	return int64(math.Round(v * scale))
}

func FromFixed[V int16 | int32 | int64](v V, r uint8) float64 {
	scale := float64(uint64(1) << r)
	return float64(v) / scale
}
