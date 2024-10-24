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

// Timestamp Details
func TestTimeStampDetailsSize(t *testing.T) {
	tsd := TimestampDetails{}
	assert.Equal(t, uint32(8), tsd.Size())
}

func TestTimeStampDetailsDefault(t *testing.T) {
	tsd := TimestampDetails{}
	assert.Equal(t, uint8(0), tsd.UserDefined)
	assert.Equal(t, false, tsd.Global)
	assert.Equal(t, uint8(0), tsd.TseCode)
	assert.Equal(t, uint8(0), tsd.LshCode)
	assert.Equal(t, uint8(0), tsd.LspCode)
	assert.Equal(t, uint8(0), tsd.TimeSource)
	assert.Equal(t, false, tsd.EnablePosixTimeOffset)
	assert.Equal(t, uint8(0), tsd.PosixTimeOffset)
	assert.Equal(t, uint32(0), tsd.TimestampEpoch)
	// Pack
	packed := tsd.Pack()
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, expected, packed)
	// Unpack
	tsd.Unpack(packed)
	assert.Equal(t, false, tsd.Global)
}

func TestTimeStampDetails(t *testing.T) {
	cases := []struct {
		name                  string
		userDefined           uint8
		global                bool
		tseCode               uint8
		lshCode               uint8
		lspCode               uint8
		timeSource            uint8
		enablePosixTimeOffset bool
		posixTimeOffset       uint8
		timestampEpoch        uint32
		expected              []byte
	}{
		{
			name:        "userDefined",
			userDefined: 0xFF,
			expected:    []byte{0xFF, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.7.3.4-4 Global True",
			global:   true,
			expected: []byte{0, 0x04, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.7.3.4-6 TSE",
			tseCode:  3,
			expected: []byte{0, 0x3, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.7.3.4-7 LSH",
			lshCode:  3,
			expected: []byte{0, 0, 0xC0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Rule 9.7.3.4-13 LSP",
			lspCode:  3,
			expected: []byte{0, 0, 0x30, 0, 0, 0, 0, 0},
		},
		{
			name:       "Rule 9.7.3.4-18 TimeSource",
			timeSource: 7,
			expected:   []byte{0, 0, 0x0E, 0, 0, 0, 0, 0},
		},
		{
			name:                  "Rule 9.7.3.4-19 Enable Posix Offset True",
			enablePosixTimeOffset: true,
			expected:              []byte{0, 0, 0x01, 0, 0, 0, 0, 0},
		},
		{
			name:            "Posix Offset False",
			posixTimeOffset: 0xFF,
			expected:        []byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:                  "Posix Offset True",
			posixTimeOffset:       0xFF,
			enablePosixTimeOffset: true,
			expected:              []byte{0, 0, 0x01, 0xFF, 0, 0, 0, 0},
		},
		{
			name:           "Rule 9.7.3.4-20 TS Epoch",
			timestampEpoch: 0xFFFFFFFF,
			expected:       []byte{0, 0, 0, 0, 0xFF, 0xFF, 0xFF, 0xFF},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tsd := TimestampDetails{
				UserDefined:           tc.userDefined,
				Global:                tc.global,
				TseCode:               tc.tseCode,
				LshCode:               tc.lshCode,
				LspCode:               tc.lspCode,
				TimeSource:            tc.timeSource,
				EnablePosixTimeOffset: tc.enablePosixTimeOffset,
				PosixTimeOffset:       tc.posixTimeOffset,
				TimestampEpoch:        tc.timestampEpoch,
			}
			assert.Equal(t, tc.userDefined, tsd.UserDefined)
			assert.Equal(t, tc.global, tsd.Global)
			assert.Equal(t, tc.tseCode, tsd.TseCode)
			assert.Equal(t, tc.lshCode, tsd.LshCode)
			assert.Equal(t, tc.lspCode, tsd.LspCode)
			assert.Equal(t, tc.timeSource, tsd.TimeSource)
			assert.Equal(t, tc.enablePosixTimeOffset, tsd.EnablePosixTimeOffset)
			assert.Equal(t, tc.posixTimeOffset, tsd.PosixTimeOffset)
			assert.Equal(t, tc.timestampEpoch, tsd.TimestampEpoch)
			// Pack
			packed := tsd.Pack()
			assert.Equal(t, tc.expected, packed)
			// Unpack
			tsd.Unpack(packed)
			assert.Equal(t, tc.userDefined, tsd.UserDefined)
			assert.Equal(t, tc.global, tsd.Global)
			assert.Equal(t, tc.tseCode, tsd.TseCode)
			assert.Equal(t, tc.lshCode, tsd.LshCode)
			assert.Equal(t, tc.lspCode, tsd.LspCode)
			assert.Equal(t, tc.timeSource, tsd.TimeSource)
			assert.Equal(t, tc.enablePosixTimeOffset, tsd.EnablePosixTimeOffset)
			togglePosixOffset := uint8(0)
			if tsd.EnablePosixTimeOffset {
				togglePosixOffset = 1
			}
			assert.Equal(t, tc.posixTimeOffset*(togglePosixOffset), tsd.PosixTimeOffset)
			assert.Equal(t, tc.timestampEpoch, tsd.TimestampEpoch)
		})
	}
}

// Sea Swell
func TestSeaSwellStateSize(t *testing.T) {
	s := SeaSwellState{}
	assert.Equal(t, uint32(4), s.Size())
}

func TestSeaSwellStateDefault(t *testing.T) {
	s := SeaSwellState{}
	assert.Equal(t, uint8(0), s.UserDefined)
	assert.Equal(t, uint8(0), s.SwellState)
	assert.Equal(t, uint8(0), s.SeaState)

	// Pack
	packed := s.Pack()
	expected := []byte{0, 0, 0, 0} // Adjust based on your packing logic
	assert.Equal(t, expected, packed)

	// Unpack
	s.Unpack(packed)
	assert.Equal(t, uint8(0), s.UserDefined)
	assert.Equal(t, uint8(0), s.SwellState)
	assert.Equal(t, uint8(0), s.SeaState)
}

func TestSeaSwellState(t *testing.T) {
	cases := []struct {
		name        string
		userDefined uint8
		swellState  uint8
		seaState    uint8
		expected    []byte
	}{
		{
			name:        "9.9.1-3 UserDefined",
			userDefined: 63,
			expected:    []byte{0, 0, 0xFC, 0}, // Adjust based on your packing logic
		},
		{
			name:       "9.9.1-2 Swell State",
			swellState: 31,
			expected:   []byte{0, 0, 0x03, 0xE0}, // Adjust based on your packing logic
		},
		{
			name:     "9.9.1-1 Sea State",
			seaState: 31,
			expected: []byte{0, 0, 0, 0x1F}, // Adjust based on your packing logic
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := SeaSwellState{
				UserDefined: tc.userDefined,
				SwellState:  tc.swellState,
				SeaState:    tc.seaState,
			}

			assert.Equal(t, tc.userDefined, s.UserDefined)
			assert.Equal(t, tc.swellState, s.SwellState)
			assert.Equal(t, tc.seaState, s.SeaState)

			// Pack
			packed := s.Pack()
			assert.Equal(t, tc.expected, packed)

			// Unpack
			s.Unpack(packed)
			assert.Equal(t, tc.userDefined, s.UserDefined)
			assert.Equal(t, tc.swellState, s.SwellState)
			assert.Equal(t, tc.seaState, s.SeaState)
		})
	}
}
