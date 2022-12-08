package day5

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStacksPop(t *testing.T) {
	sp := stacksPlatform{
		{"A", "B", "C"},
		{"X", "Y", "Z"},
	}

	gotPopped := sp.pop(1, 2)
	expPopped := []string{"Y", "Z"}
	if !cmp.Equal(gotPopped, expPopped) {
		t.Errorf("Wrong popped crates returned from pop(), expected %v, got %v", expPopped, gotPopped)
	}

	gotLeftInStack := sp[1]
	expLeftInStack := []string{"X"}
	if !cmp.Equal(gotLeftInStack, expLeftInStack) {
		t.Errorf("Wrong crates left in popped stack after using pop(), expected %v, got %v", expLeftInStack, gotLeftInStack)
	}

}

func TestStacksParseInitialCrates(t *testing.T) {
	var sp stacksPlatform
	var got, exp []string

	sp = make(stacksPlatform, 3)
	sp.parseInitalCrates("    [D]    ")
	got = sp[2]
	exp = []string{"D"}
	if !cmp.Equal(got, exp) {
		t.Errorf("Incorrect parsing of initial crates row, expected %v, got %v", exp, got)
	}

	sp = make(stacksPlatform, 10)
	sp.parseInitalCrates("    [C] [R] [Z]     [R]     [H] [Z]")
	got = []string{sp[2][0], sp[3][0], sp[4][0], sp[6][0], sp[8][0], sp[9][0]}
	exp = []string{"C", "R", "Z", "R", "H", "Z"}
	if !cmp.Equal(got, exp) {
		t.Errorf("Incorrect parsing of initial crates row, expected %v, got %v", exp, got)
	}

}

func TestStacksMoveCratesPerCrate(t *testing.T) {
	sp := stacksPlatform{
		{},
		{"A", "B", "C"},
		{"D", "E", "F"},
		{"G", "H", "I"},
	}
	sp.moveCrates("move 2 from 1 to 3", perCrate)

	exp := stacksPlatform{
		{},
		{"A"},
		{"D", "E", "F"},
		{"G", "H", "I", "C", "B"},
	}
	got := sp
	if !cmp.Equal(got, exp) {
		t.Errorf("Incorrect new stacks platform after moving per crate, expected %v, got %v", exp, got)
	}
}

func TestStacksMoveCratesPerGroup(t *testing.T) {
	sp := stacksPlatform{
		{},
		{"A", "B", "C"},
		{"D", "E", "F"},
		{"G", "H", "I"},
	}
	sp.moveCrates("move 2 from 1 to 3", perGroup)

	exp := stacksPlatform{
		{},
		{"A"},
		{"D", "E", "F"},
		{"G", "H", "I", "B", "C"},
	}
	got := sp
	if !cmp.Equal(got, exp) {
		t.Errorf("Incorrect new stacks platform after moving per group, expected %v, got %v", exp, got)
	}
}

func TestStacksTopCrates(t *testing.T) {
	sp := stacksPlatform{
		{"A", "B", "C"},
		{"X", "Y", "Z"},
		{"T", "C", "V", "B"},
		{"Z"},
	}

	got := sp.topCrates()
	exp := "CZBZ"

	if got != exp {
		t.Errorf("Wrong top crates string returned, expected %v, got %v", exp, got)
	}

}
