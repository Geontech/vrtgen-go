/*
 * Copyright (C) 2024 Geon Technologies, LLC
 *
 * This file is part of vrtgen-go.
 *
 * vrtgen-go is free software: you can redistribute it and/or modify it under the
 * terms of the GNU Lesser General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * vrtgen-go is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
 * FOR A PARTICULAR PURPOSE.  See the GNU Lesser General Public License for
 * more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */

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
