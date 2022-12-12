// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day8

import (
	"fmt"
	"io"
	"os"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := maxScenicScore(f)

	fmt.Printf("The answer is %d\n", res)

}

func maxScenicScore(r io.Reader) int {
	var maxScore int

	grid := parseGrid(r)
	for y, row := range grid {
		for x, _ := range row {
			s := grid.scenicScore(y, x)
			if s > maxScore {
				maxScore = s
			}
		}
	}
	return maxScore
}
