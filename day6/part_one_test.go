// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day6

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range tests {
		b := bytes.NewBufferString(tc.input)
		got, _ := startOfPacketMarkerPos(b)
		if got != tc.exp {
			t.Errorf("Wrong start of packet marker position, expected %d, got %d", tc.exp, got)
		}
	}

}

func TestSampleForOneWithError(t *testing.T) {
	input := "abcbabcbabcbabcb"
	b := bytes.NewBufferString(input)

	_, err := startOfPacketMarkerPos(b)

	if err == nil { // err IS nil
		t.Fatalf("Expected error searching for start-of-packet from %q, got none", input)
	}

}

func TestRealForOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got, _ := startOfPacketMarkerPos(f)

	// Your answer is probably different
	exp := 1850

	if got != exp {
		t.Errorf("Wrong start of packet marker position, expected %d, got %d", exp, got)
	}

}
