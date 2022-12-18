// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day9

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForTwo(t *testing.T) {
	const sample = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

	b := bytes.NewBufferString(sample)
	got := tailVisits(b, 10)
	exp := 36
	if got != exp {
		t.Errorf("Wrong number visited positions for the tail, expected %d, got %d", exp, got)
	}

}

func TestRealForTwo(t *testing.T) {

	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := tailVisits(f, 10)

	// Your answer is probably different
	exp := 2536

	if got != exp {
		t.Errorf("Wrong number visited positions for the tail, expected %d, got %d", exp, got)
	}

}
