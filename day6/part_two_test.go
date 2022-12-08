// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day6

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForTwo(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range tests {
		b := bytes.NewBufferString(tc.input)
		got, _ := startOfMessageMarkerPos(b)
		if got != tc.exp {
			t.Errorf("Wrong start of message marker position, expected %d, got %d", tc.exp, got)
		}
	}

}

func TestRealForTwo(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got, _ := startOfMessageMarkerPos(f)

	// Your answer is probably different
	exp := 2823

	if got != exp {
		t.Errorf("Wrong start of message marker position, expected %d, got %d", exp, got)
	}

}
