// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res, _ := startOfMessageMarkerPos(f)

	fmt.Printf("The answer is %d\n", res)

}

func startOfMessageMarkerPos(r io.Reader) (int, error) {
	groupSize := 14

	s := bufio.NewScanner(r)
	s.Split(scanGrouped(groupSize))

	endPos := groupSize
	for s.Scan() {
		if len(s.Bytes()) > 0 && unique(s.Bytes()) {
			return endPos, nil
		}
		endPos++
	}

	return 0, fmt.Errorf("no start-of-message found")
}
