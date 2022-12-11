// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"fmt"
	"io"
	"os"
)

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day7/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := sumOfDirsUnder100000(f)

	fmt.Printf("The answer is %d\n", res)

}

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
