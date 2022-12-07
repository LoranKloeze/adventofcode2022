// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day5

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	const sample = `    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

	b := bytes.NewBufferString(sample)
	got := resultOfCrateMover9000(b)
	exp := "CMZ"

	if got != exp {
		t.Errorf("Wrong top rows of stacks using crate mover 9000, expected %q, got %q", exp, got)
	}
}

func TestRealForOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := resultOfCrateMover9000(f)

	// Your answer is probably different
	exp := "RFFFWBPNS"

	if got != exp {
		t.Errorf("Wrong top rows of stacks using crate mover 9000, expected %q, got %q", exp, got)
	}

}
