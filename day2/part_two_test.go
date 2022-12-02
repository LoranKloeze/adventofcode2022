// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day2

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForTwo(t *testing.T) {
	const sample = `A Y
B X
C Z
`

	b := bytes.NewBufferString(sample)
	got := MyActualScore(b)
	exp := 12

	if got != exp {
		t.Errorf("Wrong total score, expected %d, got %d", exp, got)
	}

}

func TestRealForTwo(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := MyActualScore(f)

	// Your answer is probably different
	answer := 12411

	if got != answer {
		t.Errorf("Wrong total score, expected %d, got %d", answer, got)
	}

}
