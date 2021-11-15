// Copyright (c) 2021 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package wipe_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/0x9ef/go-wiper/wipe"
)

const _100KB = 100 * 1024
const _1MB = 1 * 1024 * 1024
const _500MB = 500 * _1MB
const _1GB = 1000 * _1MB

func generateFile(path string, size int) {
	fd, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer fd.Close()
	buf := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	rand.Read(buf)
	if _, err := fd.Write(buf); err != nil {
		panic(err)
	}
}

func TestSmall100KB(t *testing.T) {
	rule := wipe.RuleFast
	path := "tests/open_test_small100KB.bin"
	generateFile(path, _100KB)
	if err := wipe.Wipe(path, rule); err != nil {
		panic(err)
	}
}

func TestSmall1MB(t *testing.T) {
	rule := wipe.RuleFast
	path := "tests/open_test_small1MB.bin"
	generateFile(path, _1MB)
	if err := wipe.Wipe(path, rule); err != nil {
		panic(err)
	}
}

func TestLarge500MB(t *testing.T) {
	rule := wipe.RuleFast
	path := "tests/open_test_large500MB.bin"
	generateFile(path, _500MB)
	if err := wipe.Wipe(path, rule); err != nil {
		panic(err)
	}
}

func TestLarge1GB(t *testing.T) {
	rule := wipe.RuleFast
	path := "tests/open_test_large1GB.bin"
	generateFile(path, _1GB)
	if err := wipe.Wipe(path, rule); err != nil {
		panic(err)
	}
}
