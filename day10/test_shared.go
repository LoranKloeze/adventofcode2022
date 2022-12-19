package day10

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInstructions(t *testing.T) {
	input :=
		"addx 5\n" +
			"addx 12\n" +
			"noop\n" +
			"addx 15"
	b := bytes.NewBufferString(input)
	got := parseInstructions(b)

	exp := []Instruction{
		{Operation: "addx", Value: 5},
		{Operation: "addx", Value: 12},
		{Operation: "noop", Value: 0},
		{Operation: "addx", Value: 15},
	}

	if !cmp.Equal(got, exp) {
		t.Errorf("Expected \n%v\ngot \n%v\n", exp, got)
	}

}
