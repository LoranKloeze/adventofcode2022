// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day1

import (
	"bytes"
	"testing"
)

func TestSampleForTwo(t *testing.T) {
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
	got := TopCaloriesSum(b, 3)
	exp := 45000

	if got != exp {
		t.Errorf("Wrong total calorie sum, expected %d, got %d", exp, got)
	}

}

func TestDifferentOrderSampleForTwo(t *testing.T) {
	const sample = `1000
2000
3000

10000

4000

5000
6000

7000
8000
9000
`

	b := bytes.NewBufferString(sample)
	got := TopCaloriesSum(b, 3)
	exp := 45000

	if got != exp {
		t.Errorf("Wrong total calorie sum, expected %d, got %d", exp, got)
	}

}
