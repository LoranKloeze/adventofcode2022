// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day1

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForOne(t *testing.T) {
	const sample = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

	b := bytes.NewBufferString(sample)
	got := maxCaloriesSum(b)
	exp := 24000

	if got != exp {
		t.Errorf("Wrong maximum calorie sum, expected %d, got %d", exp, got)
	}

}

func TestRealForOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := maxCaloriesSum(f)

	// Your answer is probably different
	answer := 71506

	if got != answer {
		t.Errorf("Wrong maximum calorie sum, expected %d, got %d", answer, got)
	}

}
