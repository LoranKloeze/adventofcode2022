// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day3

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForTwo(t *testing.T) {
	const sample = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	b := bytes.NewBufferString(sample)
	got := sumPrioGroups(b)
	exp := 70

	if got != exp {
		t.Errorf("Wrong priority sum for groups of 3, expected %d, got %d", exp, got)
	}
}

func TestRealForTwo(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := sumPrioGroups(f)

	// Your answer is probably different
	exp := 2413

	if got != exp {
		t.Errorf("Wrong priority sum for groups of 3, expected %d, got %d", exp, got)
	}

}
