// Copyright (c) 2021 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package wipe

import (
	crand "crypto/rand"
	rand "math/rand"
	"time"
)

// End Of File
const EOF = 0

// GetRandomByte returns a random byte with selected cryptogrphically mode.
// If flag is RandomCrypto, we use a crypto/rand module for generation strongly cryprographically byte.
// In otherwise flags we use a standart cryptographically library
func GetRandomByte(n int, flag Flag) (byte, error) {
	if flag&FlagNative != 0 {
		b := []byte{0}
		_, err := crand.Reader.Read(b)
		if err != nil {
			return EOF, err
		}
		return b[0], nil
	}

	rand.Seed(time.Now().UnixNano())
	if n > 255 {
		return 0, nil
	}
	return byte(rand.Intn(n)), nil
}
