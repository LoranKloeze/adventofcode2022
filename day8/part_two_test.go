// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day8

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForTwo(t *testing.T) {
	const sample = `30373
25512
65332
33549
35390
`

	b := bytes.NewBufferString(sample)
	got := maxScenicScore(b)
	exp := 8
	if got != exp {
		t.Errorf("Wrong max scenic score, expected %d, got %d", exp, got)
	}

}

func TestRealForTwo(t *testing.T) {

	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := maxScenicScore(f)

	// Your answer is probably different
	exp := 590824

	if got != exp {
		t.Errorf("Wrong max scenic score, expected %d, got %d", exp, got)
	}

}
