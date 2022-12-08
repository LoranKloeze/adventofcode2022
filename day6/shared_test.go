// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day6

import (
	"testing"
)

func TestUnique(t *testing.T) {

	tests := []struct {
		input []byte
		exp   bool
	}{
		{[]byte{'a', 'b', 'c'}, true},
		{[]byte{'a', 'b', 'b'}, false},
	}

	for _, tc := range tests {
		got := unique(tc.input)
		if got != tc.exp {
			t.Errorf("Wrong result for unique(%v), expected %v, got %v", tc.input, tc.exp, got)
		}
	}
}
