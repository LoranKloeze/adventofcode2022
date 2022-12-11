// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"fmt"
	"io"
)

func sumOfDirsUnder100000(r io.Reader) int {
	root, err := parseTree(r)
	if err != nil {
		fmt.Printf("Unexpected error while parsing tree: %v", err)
	}

	sum := 0
	root.findSizeAtMost(100000, func(e *Entry, sz int) {
		sum += sz
	})
	return sum
}
