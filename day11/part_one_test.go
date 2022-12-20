// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day11

import (
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	f, err := os.Open("input_sample.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := monkeyBusiness(f, 20, "part1")
	exp := 10605
	if got != exp {
		t.Errorf("Wrong monkey business, expected %d, got %d", exp, got)
	}

}

func TestRealForOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := monkeyBusiness(f, 20, "part1")

	// Your answer is probably different
	exp := 50172

	if got != exp {
		t.Errorf("Wrong monkey business, expected %d, got %d", exp, got)
	}
}
