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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixed16(t *testing.T) {
	cases := []struct {
		name            string
		radix           uint8
		negativeOneInt  uint16
		largestPosInt   uint16
		largestNegInt   uint16
		largestPosVal   uint16
		largestPosFrac  uint16
		largestNegFrac  uint16
		smallestNegFrac uint16
	}{
		{
			name:            "Q9.7",
			radix:           7,
			negativeOneInt:  uint16(0xFF80),
			largestPosInt:   uint16(0x7F80),
			largestNegInt:   uint16(0x8000),
			largestPosVal:   uint16(0x7FFF),
			largestPosFrac:  uint16(0x007F),
			largestNegFrac:  uint16(0xFF81),
			smallestNegFrac: uint16(0xFFFF),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Radix
			radixInt := int16(1) << tc.radix
			floatVal := float64(1)
			assert.Equal(t, radixInt, ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(radixInt, tc.radix))
			// Negative one
			floatVal = float64(-1)
			assert.Equal(t, int16(tc.negativeOneInt), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(tc.negativeOneInt), tc.radix))
			// Largest Positive Int
			floatVal = float64(uint64(1<<(16-tc.radix-1)) - 1)
			assert.Equal(t, int16(tc.largestPosInt), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(tc.largestPosInt), tc.radix))
			// Largest Negative Int
			floatVal = float64(-(int(1) << (16 - tc.radix - 1)))
			assert.Equal(t, int16(tc.largestNegInt), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(tc.largestNegInt), tc.radix))
			// Largest Positive Value
			floatVal = float64(float64(int(1)<<(16-tc.radix-1)) - (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int16(tc.largestPosVal), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(tc.largestPosVal), tc.radix))
			// Largest Positive Fraction
			floatVal = float64(1.0 - (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int16(tc.largestPosFrac), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(tc.largestPosFrac), tc.radix))
			// Smallest Positive Fraction
			floatVal = float64(1.0 / float64(int(1)<<tc.radix))
			assert.Equal(t, int16(1), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(1), tc.radix))
			// Largest Negative Fraction
			floatVal = float64(-1.0 + (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int16(tc.largestNegFrac), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(tc.largestNegFrac), tc.radix))
			// Smallest Negative Fraction
			floatVal = float64(-1.0 / float64(int(1)<<tc.radix))
			assert.Equal(t, int16(tc.smallestNegFrac), ToFixed16(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int16(tc.smallestNegFrac), tc.radix))
		})
	}
}

func TestFixed32(t *testing.T) {
	cases := []struct {
		name            string
		radix           uint8
		negativeOneInt  uint32
		largestPosInt   uint32
		largestNegInt   uint32
		largestPosVal   uint32
		largestPosFrac  uint32
		largestNegFrac  uint32
		smallestNegFrac uint32
	}{
		{
			name:            "Q10.22",
			radix:           22,
			negativeOneInt:  uint32(0xFFC00000),
			largestPosInt:   uint32(0x7FC00000),
			largestNegInt:   uint32(0x80000000),
			largestPosVal:   uint32(0x7FFFFFFF),
			largestPosFrac:  uint32(0x003FFFFF),
			largestNegFrac:  uint32(0xFFC00001),
			smallestNegFrac: uint32(0xFFFFFFFF),
		},
		{
			name:            "Q27.5",
			radix:           5,
			negativeOneInt:  uint32(0xFFFFFFE0),
			largestPosInt:   uint32(0x7FFFFFE0),
			largestNegInt:   uint32(0x80000000),
			largestPosVal:   uint32(0x7FFFFFFF),
			largestPosFrac:  uint32(0x0000001F),
			largestNegFrac:  uint32(0xFFFFFFE1),
			smallestNegFrac: uint32(0xFFFFFFFF),
		},
		{
			name:            "Q16.16",
			radix:           16,
			negativeOneInt:  uint32(0xFFFF0000),
			largestPosInt:   uint32(0x7FFF0000),
			largestNegInt:   uint32(0x80000000),
			largestPosVal:   uint32(0x7FFFFFFF),
			largestPosFrac:  uint32(0x0000FFFF),
			largestNegFrac:  uint32(0xFFFF0001),
			smallestNegFrac: uint32(0xFFFFFFFF),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Radix
			radixInt := int32(1) << tc.radix
			floatVal := float64(1)
			assert.Equal(t, radixInt, ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(radixInt, tc.radix))
			// Negative one
			floatVal = float64(-1)
			assert.Equal(t, int32(tc.negativeOneInt), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.negativeOneInt), tc.radix))
			// Largest Positive Int
			floatVal = float64(uint64(1<<(32-tc.radix-1)) - 1)
			assert.Equal(t, int32(tc.largestPosInt), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.largestPosInt), tc.radix))
			// Largest Negative Int
			floatVal = float64(-(int(1) << (32 - tc.radix - 1)))
			assert.Equal(t, int32(tc.largestNegInt), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.largestNegInt), tc.radix))
			// Largest Positive Value
			floatVal = float64(float64(int(1)<<(32-tc.radix-1)) - (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int32(tc.largestPosVal), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.largestPosVal), tc.radix))
			// Largest Positive Fraction
			floatVal = float64(1.0 - (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int32(tc.largestPosFrac), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.largestPosFrac), tc.radix))
			// Smallest Positive Fraction
			floatVal = float64(1.0 / float64(int(1)<<tc.radix))
			assert.Equal(t, int32(1), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(1), tc.radix))
			// Largest Negative Fraction
			floatVal = float64(-1.0 + (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int32(tc.largestNegFrac), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.largestNegFrac), tc.radix))
			// Smallest Negative Fraction
			floatVal = float64(-1.0 / float64(int(1)<<tc.radix))
			assert.Equal(t, int32(tc.smallestNegFrac), ToFixed32(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.smallestNegFrac), tc.radix))
		})
	}
}

func TestFixed64(t *testing.T) {
	cases := []struct {
		name            string
		radix           uint8
		negativeOneInt  uint64
		largestPosInt   uint64
		largestNegInt   uint64
		largestPosVal   uint64
		largestPosFrac  uint64
		largestNegFrac  uint64
		smallestNegFrac uint64
	}{
		{
			name:            "Q44.20",
			radix:           20,
			negativeOneInt:  uint64(0xFFFFFFFFFFF00000),
			largestPosInt:   uint64(0x7FFFFFFFFFF00000),
			largestNegInt:   uint64(0x8000000000000000),
			largestPosVal:   uint64(0x7FFFFFFFFFFFFFFF),
			largestPosFrac:  uint64(0x00000000000FFFFF),
			largestNegFrac:  uint64(0xFFFFFFFFFFF00001),
			smallestNegFrac: uint64(0xFFFFFFFFFFFFFFFF),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Radix
			radixInt := int64(1) << tc.radix
			floatVal := float64(1)
			assert.Equal(t, radixInt, ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(radixInt, tc.radix))
			// Negative one
			floatVal = float64(-1)
			assert.Equal(t, int64(tc.negativeOneInt), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int32(tc.negativeOneInt), tc.radix))
			// Largest Positive Int
			floatVal = float64(uint64(1<<(64-tc.radix-1)) - 1)
			assert.Equal(t, int64(tc.largestPosInt), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int64(tc.largestPosInt), tc.radix))
			// Largest Negative Int
			floatVal = float64(-(int(1) << (64 - tc.radix - 1)))
			assert.Equal(t, int64(tc.largestNegInt), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int64(tc.largestNegInt), tc.radix))
			// Largest Positive Value
			floatVal = float64(float64(int(1)<<(64-tc.radix-1)) - (1.0 / float64(int(1)<<tc.radix)))
			// TODO - assess below with float64
			// assert.Equal(t, int64(tc.largestPosVal), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int64(tc.largestPosVal), tc.radix))
			// Largest Positive Fraction
			floatVal = float64(1.0 - (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int64(tc.largestPosFrac), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int64(tc.largestPosFrac), tc.radix))
			// Smallest Positive Fraction
			floatVal = float64(1.0 / float64(int(1)<<tc.radix))
			assert.Equal(t, int64(1), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int64(1), tc.radix))
			// Largest Negative Fraction
			floatVal = float64(-1.0 + (1.0 / float64(int(1)<<tc.radix)))
			assert.Equal(t, int64(tc.largestNegFrac), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int64(tc.largestNegFrac), tc.radix))
			// Smallest Negative Fraction
			floatVal = float64(-1.0 / float64(int(1)<<tc.radix))
			assert.Equal(t, int64(tc.smallestNegFrac), ToFixed64(floatVal, tc.radix))
			assert.Equal(t, floatVal, FromFixed(int64(tc.smallestNegFrac), tc.radix))
		})
	}
}
