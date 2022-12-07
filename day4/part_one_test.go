// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day4

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	const sample = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	b := bytes.NewBufferString(sample)
	got := totalFullyContain(b)
	exp := 2

	if got != exp {
		t.Errorf("Wrong total of fully containing array pairs, expected %d, got %d", exp, got)
	}
}

func TestRealForOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := totalFullyContain(f)

	// Your answer is probably different
	exp := 536

	if got != exp {
		t.Errorf("Wrong total of fully containing array pairs, expected %d, got %d", exp, got)
	}

}
