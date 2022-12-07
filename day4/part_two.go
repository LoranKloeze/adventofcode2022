// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day4

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Main entry for part one of this day
func PartTwo() {
	f, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := totalOverlapping(f)

	fmt.Printf("The answer is %d\n", res)

}

func totalOverlapping(r io.Reader) int {

	s := bufio.NewScanner(r)
	var cnt int
	for s.Scan() {
		p1, p2, err := extractPairs(s.Text())
		if err != nil {
			log.Fatalf("Error parsing pairs: %v", err)
		}

		p1OverlapsP2 := p1.begin <= p2.end && p1.begin >= p2.begin
		p2OverlapsP1 := p2.begin <= p1.end && p2.begin >= p1.begin
		if p1OverlapsP2 || p2OverlapsP1 {
			cnt++
		}

	}
	return cnt
}
