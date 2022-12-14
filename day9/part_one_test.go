// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day9

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	const sample = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	b := bytes.NewBufferString(sample)
	got := uniqueTailVisits(b)
	exp := 13
	if got != exp {
		t.Errorf("Wrong number unique tail visitis, expected %d, got %d", exp, got)
	}

}

func TestRealForOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := uniqueTailVisits(f)

	// Your answer is probably different
	exp := 6367

	if got != exp {
		t.Errorf("Wrong number unique tail visitis, expected %d, got %d", exp, got)
	}

}
