// Copyright (c) 2021 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package wipe

import (
	"fmt"
)

// Flag represent state of random generation by two library: rand.Rand and crypto.Rand
type Flag uint8

const (
	// FlagNative flag thats indicates of using non-cryptographically string random generation
	FlagNative Flag = iota + 1

	// FlagCrypto flag thats indicates of using cryptographically strong random generation from crypto.Rand module
	FlagCrypto
)

// Pass represents data which will be written to file as payload,
// lenght of this data and random flag, if random != 0 will be generated random byte 0-255.
type Pass struct {
	Data   []byte
	Len    uint8
	Random Flag
}

// NewPass returns a qualified Passage with setuped payload data and random flag
func NewPass(data []byte, random Flag) *Pass {
	return &Pass{Data: data, Len: uint8(len(data)), Random: random}
}

// Rule represents a state of passes thats can be used by Wipe function for safely wiping data.
// You can implement own alghoritm with 1-255 passes and data set.
type Rule []Pass

// NewRule returns
func NewRule(passes []Pass) *Rule {
	p := Rule(passes)
	return &p
}

// Add adding a new passage to rule
func (r *Rule) Add(pass Pass) { *r = append(*r, pass) }

// Clear clearing all passages from rule
func (r *Rule) Clear() { *r = (*r)[0:] }

// String stringify the rule
func (r *Rule) String() string {
	var msg string
	for i, pass := range *r {
		var decoded string
		if pass.Random&FlagNative != 0 {
			decoded = "RandomNative"
		} else if pass.Random&FlagCrypto != 0 {
			decoded = "RandomCrypto"
		} else {
			decoded = "None"
		}
		msg += fmt.Sprintf("PASS #%d: .data=bytes%v, .len=%d, .flag=%s\n", i, pass.Data, pass.Len, decoded)
	}
	return msg
}

var (
	// RuleFast data will be overwrited with zero bytes
	RuleFast = &Rule{
		{[]byte{0}, 1, 0},
	}

	// RuleVSITR ...
	// The German standard calls for each sector to be
	// overwritten with three alternating patterns of zeroes and ones
	// and in the last pass with 10101010
	RuleVSITR = &Rule{
		{[]byte("\x00"), 1, 0},
		{[]byte("\xFF"), 1, 0},
		{[]byte("\x00"), 1, 0},
		{[]byte("\xFF"), 1, 0},
		{[]byte("\x00"), 1, 0},
		{[]byte("\xFF"), 1, 0},
		{[]byte("\xAA"), 1, 0},
	}

	// RuleUsDod5220_22_M ...
	// US Department of Defense DoD 5220.22-M (3 passes)
	// DoD 5220.22-M is three pass overwriting algorithm: first pass
	// with zeroes, second pass with ones and the last pass with random bytes.
	// With all passes verification.
	RuleUsDod5220_22_M = &Rule{
		{[]byte("\x00"), 1, 0},
		{[]byte("\xFF"), 1, 0},
		{[]byte{0}, 1, FlagNative},
	}

	// RuleGutmann ...
	// Peter Gutmann alghoritm (35 passes)
	// The selection of patterns assumes that the user does not know the encoding mechanism
	// used by the drive, so it includes patterns designed specifically for three types of drives.
	//  A user who knows which type of encoding the drive uses can choose only those patterns intended
	// for their drive. A drive with a different encoding mechanism would need different patterns.
	//
	// An overwrite session consists of a lead-in of four random write patterns,
	// followed by patterns 5 to 31 (see rows of table below), executed in a random order,
	// and a lead-out of four more random patterns.
	// Each of patterns 5 to 31 was designed with a specific magnetic media encoding scheme in mind,
	// which each pattern targets. The drive is written to for all the passes even though the table below
	// only shows the bit patterns for the passes that are specifically targeted at each encoding scheme.
	// The end result should obscure any data on the drive so that only the most
	// advanced physical scanning (e.g., using a magnetic force microscope) of the drive is likely
	// to be able to recover any data. The series of patterns is as follows:
	RuleGutmann = &Rule{
		{[]byte{0}, 1, FlagNative},
		{[]byte{0}, 1, FlagNative},
		{[]byte{0}, 1, FlagNative},
		{[]byte{0}, 1, FlagNative},
		{[]byte("\x55"), 1, 0},         // 	01010101
		{[]byte("\xAA"), 1, 0},         //	10101010
		{[]byte("\x92\x49\x24"), 3, 0}, //	10010010 01001001 00100100
		{[]byte("\x49\x24\x92"), 3, 0}, //	01001001 00100100 10010010
		{[]byte("\x24\x92\x49"), 3, 0}, //	00100100 10010010 01001001
		{[]byte("\x00"), 1, 0},
		{[]byte("\x11"), 1, 0},
		{[]byte("\x22"), 1, 0},
		{[]byte("\x33"), 1, 0},
		{[]byte("\x44"), 1, 0},
		{[]byte("\x55"), 1, 0},
		{[]byte("\x66"), 1, 0},
		{[]byte("\x77"), 1, 0},
		{[]byte("\x88"), 1, 0},
		{[]byte("\x99"), 1, 0},
		{[]byte("\xAA"), 1, 0},
		{[]byte("\xBB"), 1, 0},
		{[]byte("\xCC"), 1, 0},
		{[]byte("\xDD"), 1, 0},
		{[]byte("\xEE"), 1, 0},
		{[]byte("\xFF"), 1, 0},
		{[]byte("\x92\x49\x24"), 3, 0}, // 	10010010 01001001 00100100
		{[]byte("\x49\x24\x92"), 3, 0}, //	01001001 00100100 10010010
		{[]byte("\x24\x92\x49"), 3, 0}, //	00100100 10010010 01001001
		{[]byte("\x6D\xB6\xDB"), 3, 0}, //	01101101 10110110 11011011
		{[]byte("\xB6\xDB\x6D"), 3, 0}, //	10110110 11011011 01101101
		{[]byte("\xDB\x92\x49"), 3, 0}, //	11011011 01101101 10110110
		{[]byte{0}, 1, FlagNative},
		{[]byte{0}, 1, FlagNative},
		{[]byte{0}, 1, FlagNative},
		{[]byte{0}, 1, FlagNative},
	}
)
