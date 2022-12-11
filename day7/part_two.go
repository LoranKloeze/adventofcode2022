// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"fmt"
	"io"
	"math"
	"os"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day7/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := sizeOfDirToDelete(f)

	fmt.Printf("The answer is %d\n", res)
}

func sizeOfDirToDelete(r io.Reader) int {
	root, err := parseTree(r)
	if err != nil {
		fmt.Printf("Unexpected error while parsing tree: %v", err)
	}

	unusedSpace := 70000000 - root.fullSize()
	tofreeUp := 30000000 - unusedSpace

	candidate := math.MaxInt
	root.findSizeAtLeast(tofreeUp, func(e *Entry, sz int) {
		if sz < candidate {
			candidate = sz
		}
	})
	return candidate
}
