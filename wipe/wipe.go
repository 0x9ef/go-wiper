// Copyright (c) 2021 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package wipe

import (
	"bytes"
	"errors"
	"math"
	"os"
	"path/filepath"
)

// FileChunk size for more suitable purposes
const FileChunk = 2 * (1 << 20) // 2 MB

// Wipe safely wipes data with provided rule.
// Wipe function goes through the passes and overwriting data from last remembered position.
// If rule.Random has FlagNative or FlagCrypto, will be generated random byte.
func Wipe(path string, rule *Rule) error {
	// Get file size
	var lastPos int64
	fd, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	fstat, err := fd.Stat()
	if err != nil {
		return err
	}

	defer fd.Close() // defer after file processing
	fileSize := fstat.Size()
	fileParts := uint64(math.Ceil(float64(fileSize) / float64(FileChunk))) // file parts can't be zero

	// We are count number of passes from rule, then we range for
	// counted times of passes and ranging over file parts for more efficient speed times
	r := *rule
	passes := len(r)
	for pass := 0; pass < passes; pass++ {
		// In some cases, we have to fill in data randomly from 0-255,
		// so we resort to using GetRandomByte(255).
		var data []byte
		if r[pass].Random != 0 {
			flag := r[pass].Random
			b, err := GetRandomByte(255, flag)
			if err != nil {
				return err
			}
			data = []byte{b}
		} else {
			data = r[pass].Data
		}

		var dataSz = int(r[pass].Len)
		for i := uint64(0); i < fileParts; i++ {
			partSize := int64(math.Min(FileChunk, float64(fileSize-int64(i*FileChunk))))
			var overwriteData []byte
			if dataSz <= 1 { // 0 or 1
				overwriteData = bytes.Repeat(data, int(partSize))
			} else {
				overwriteData = make([]byte, partSize)
				for block := 0; block < int(partSize); block += dataSz {
					for index := 0; index < dataSz; index++ {
						overwriteData[block] = data[index] // fill block with selected byte
					}
				}
			}

			n, err := fd.WriteAt(overwriteData, lastPos)
			switch {
			case err != nil:
				return err
			case n == 0:
				return errors.New("wiper: written null bytes to a file")
			}

			lastPos += partSize // update last position
		}
	}
	return nil
}

// Remove gets absolute file path from file descriptor and removes a directory or file
func Remove(fd *os.File) error {
	if fd != nil && fd.Close() != nil {
		path, err := filepath.Abs(fd.Name())
		if err != nil {
			return err
		}
		return os.Remove(path)
	}
	return errors.New("wiper: provided invalid handle of file")
}
