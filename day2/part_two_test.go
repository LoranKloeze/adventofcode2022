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
	got := myActualScore(b)
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

	got := myActualScore(f)

	// Your answer is probably different
	answer := 12411

	if got != answer {
		t.Errorf("Wrong total score, expected %d, got %d", answer, got)
	}

}

func TestParseActualMove(t *testing.T) {

	type testCase struct {
		opponentPlays string
		mePlays       string
		expScore      int
	}

	tests := []testCase{
		{opponentPlays: "A", mePlays: "X", expScore: 3},
		{opponentPlays: "A", mePlays: "Y", expScore: 4},
		{opponentPlays: "A", mePlays: "Z", expScore: 8},

		{opponentPlays: "B", mePlays: "X", expScore: 1},
		{opponentPlays: "B", mePlays: "Y", expScore: 5},
		{opponentPlays: "B", mePlays: "Z", expScore: 9},

		{opponentPlays: "C", mePlays: "X", expScore: 2},
		{opponentPlays: "C", mePlays: "Y", expScore: 6},
		{opponentPlays: "C", mePlays: "Z", expScore: 7},
	}

	for _, c := range tests {
		if got, _ := parseActualMove(c.opponentPlays, c.mePlays); got != c.expScore {
			t.Errorf("Expected parseActualMove(%q,%q) to return %d, got %d", c.opponentPlays, c.mePlays, c.expScore, got)
		}
	}
}

func TestUnknownActualMove(t *testing.T) {
	opponent, me := "foo", "bar"
	_, err := parseActualMove(opponent, me)
	if err == nil {
		t.Errorf("Expected call parseMove(%q,%q) to return error", opponent, me)
	}
}
