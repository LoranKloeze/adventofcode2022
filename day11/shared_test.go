// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day11

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseMonkeys(t *testing.T) {
	input := `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
	b := bytes.NewBufferString(input)

	gotMonkeys, gotTotalMod := parseMonkeys(b)

	expMonkeys := map[int]*Monkey{
		0: {Id: 0, Items: []int{79, 98}, Operation: "* 19", DivisibleBy: 23, InspectedCnt: 0},
		1: {Id: 1, Items: []int{54, 65, 75, 74}, Operation: "+ 6", DivisibleBy: 19, InspectedCnt: 0},
		2: {Id: 2, Items: []int{79, 60, 97}, Operation: "* old", DivisibleBy: 13, InspectedCnt: 0},
		3: {Id: 3, Items: []int{74}, Operation: "+ 3", DivisibleBy: 17, InspectedCnt: 0},
	}

	expMonkeys[0].TrueTo = expMonkeys[2]
	expMonkeys[0].FalseTo = expMonkeys[3]
	expMonkeys[1].TrueTo = expMonkeys[2]
	expMonkeys[1].FalseTo = expMonkeys[0]
	expMonkeys[2].TrueTo = expMonkeys[1]
	expMonkeys[2].FalseTo = expMonkeys[3]
	expMonkeys[3].TrueTo = expMonkeys[0]
	expMonkeys[3].FalseTo = expMonkeys[1]

	for k, e := range expMonkeys {
		g := gotMonkeys[k]
		if e.Id != g.Id {
			t.Errorf("Monkey [%d] - Id, expected %d, got %d", e.Id, e.Id, g.Id)
		}
		if !cmp.Equal(e.Items, g.Items) {
			t.Errorf("Monkey [%d] - Items, expected %d, got %d", e.Id, e.Items, g.Items)
		}
		if e.Operation != g.Operation {
			t.Errorf("Monkey [%d] - Operation, expected %q, got %q", e.Id, e.Operation, g.Operation)
		}
		if e.DivisibleBy != g.DivisibleBy {
			t.Errorf("Monkey [%d] - DivisibleBy, expected %d, got %d", e.Id, e.DivisibleBy, g.DivisibleBy)
		}
		if e.InspectedCnt != g.InspectedCnt {
			t.Errorf("Monkey [%d] - InspectedCnt, expected %d, got %d", e.Id, e.InspectedCnt, g.InspectedCnt)
		}
		if e.TrueTo.Id != g.TrueTo.Id {
			t.Errorf("Monkey [%d] - TrueTo.Id, expected %d, got %d", e.Id, e.TrueTo.Id, g.TrueTo.Id)
		}
		if e.FalseTo.Id != g.FalseTo.Id {
			t.Errorf("Monkey [%d] - FalseTo.Id, expected %d, got %d", e.Id, e.FalseTo.Id, g.FalseTo.Id)
		}
	}

	expTotalMod := 96577

	if gotTotalMod != expTotalMod {
		t.Errorf("Total modulo is incorrect, expected %d, got %d", expTotalMod, gotTotalMod)
	}
}
