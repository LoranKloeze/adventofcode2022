// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day8

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	const sample = `30373
25512
65332
33549
35390
`

	b := bytes.NewBufferString(sample)
	got := visibleFromOutside(b)
	exp := 21
	if got != exp {
		t.Errorf("Wrong number of visible trees from the outside, expected %d, got %d", exp, got)
	}

}

func TestRealForOne(t *testing.T) {

	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := visibleFromOutside(f)

	// Your answer is probably different
	exp := 1719

	if got != exp {
		t.Errorf("Wrong number of visible trees from the outside, expected %d, got %d", exp, got)
	}

}
