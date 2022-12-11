// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"fmt"
	"io"
	"math"
)

func sizeOfDirToDelete(r io.Reader) int {
	root, err := parseTree(r)
	if err != nil {
		fmt.Printf("Unexpected error while parsing tree: %v", err)
	}

	unusedSpace := 70000000 - root.fullSize()
	tofreeUp := 30000000 - unusedSpace
	fmt.Println("We need to free up", tofreeUp)

	candidate := math.MaxInt
	root.findSizeAtLeast(tofreeUp, func(e *Entry, sz int) {
		if sz < candidate {
			candidate = sz
		}
	})
	return candidate
}
