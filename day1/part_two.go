// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	fmt.Printf("The answer is %d\n", topCaloriesSum(f, 3))

}

func topCaloriesSum(r io.Reader, top int) int {

	s := bufio.NewScanner(r)

	var sum, totalSum int
	var sums []int

	for s.Scan() {

		if s.Text() == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}

		calories, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("Expected a number: %v", s.Text())
		}

		sum += calories
	}

	// Anything left in sum?
	if sum > 0 {
		sums = append(sums, sum)
	}

	sort.Ints(sums)

	for _, n := range sums[len(sums)-top:] {
		totalSum += n
	}

	return totalSum
}
