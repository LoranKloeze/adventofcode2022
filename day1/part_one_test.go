// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day1

import (
	"bytes"
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
	got := MaxCaloriesSum(b)
	exp := 24000

	if got != exp {
		t.Errorf("Wrong maximum calorie sum, expected %d, got %d", exp, got)
	}

}
