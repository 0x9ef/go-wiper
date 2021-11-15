// Copyright (c) 2021 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package wipe

import (
	crand "crypto/rand"
	"math/big"
	rand "math/rand"
	"time"
)

const EOF = 0

// GetRandomByte returns a random byte with selected cryptogrphically mode.
// If flag is RandomCrypto, we use a crypto/rand module for generation strongly cryprographically byte.
// In otherwise flags we use a standart cryptographically library
func GetRandomByte(n uint8, flag Flag) (byte, error) {
	if n == 0 {
		return EOF, nil
	}
	if flag&FlagCrypto != 0 {
		bigN, err := crand.Int(crand.Reader, big.NewInt(int64(n)))
		if err != nil {
			return EOF, err
		}
		return byte(bigN.Int64()), nil
	}
	rand.Seed(time.Now().UnixNano())
	return byte(rand.Intn(int(n))), nil
}
