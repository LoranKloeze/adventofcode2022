package day3

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	const sample = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	b := bytes.NewBufferString(sample)
	got := SumPrioCompartments(b)
	exp := 157

	if got != exp {
		t.Errorf("Wrong priority sum for compartments, expected %d, got %d", exp, got)
	}
}

func TestRealForOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := SumPrioCompartments(f)

	// Your answer is probably different
	exp := 8394

	if got != exp {
		t.Errorf("Wrong priority sum for compartments, expected %d, got %d", exp, got)
	}

}
