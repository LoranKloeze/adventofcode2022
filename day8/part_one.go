// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day8

import (
	"fmt"
	"io"
	"os"
)

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := visibleFromOutside(f)

	fmt.Printf("The answer is %d\n", res)

}

func visibleFromOutside(r io.Reader) int {
	var visible int
	grid := parseGrid(r)
	for y, row := range grid {
		for x, _ := range row {
			if grid.isVisible(y, x) {
				visible++
			}
		}
	}
	return visible
}
