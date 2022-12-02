// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day2

import (
	"bytes"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	const sample = `A Y
B X
C Z
`

	b := bytes.NewBufferString(sample)
	got := MyScore(b)
	exp := 15

	if got != exp {
		t.Errorf("Wrong total score, expected %d, got %d", exp, got)
	}

}

func TestParseMove(t *testing.T) {

	type testCase struct {
		opponentPlays string
		mePlays       string
		expScore      int
	}

	tests := []testCase{
		{opponentPlays: "A", mePlays: "X", expScore: 4},
		{opponentPlays: "A", mePlays: "Y", expScore: 8},
		{opponentPlays: "A", mePlays: "Z", expScore: 3},

		{opponentPlays: "B", mePlays: "X", expScore: 1},
		{opponentPlays: "B", mePlays: "Y", expScore: 5},
		{opponentPlays: "B", mePlays: "Z", expScore: 9},

		{opponentPlays: "C", mePlays: "X", expScore: 7},
		{opponentPlays: "C", mePlays: "Y", expScore: 2},
		{opponentPlays: "C", mePlays: "Z", expScore: 6},
	}

	for _, c := range tests {
		if got, _ := parseMove(c.opponentPlays, c.mePlays); got != c.expScore {
			t.Errorf("Expected parseMove(%q,%q) to return %d, got %d", c.opponentPlays, c.mePlays, c.expScore, got)
		}
	}
}

func TestUnknownMove(t *testing.T) {
	opponent, me := "foo", "bar"
	_, err := parseMove(opponent, me)
	if err == nil {
		t.Errorf("Expected call parseMove(%q,%q) to return error", opponent, me)
	}
}
