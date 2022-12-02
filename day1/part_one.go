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
	"strconv"
)

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	fmt.Printf("The answer is %d\n", maxCaloriesSum(f))

}

func maxCaloriesSum(r io.Reader) int {

	s := bufio.NewScanner(r)

	var sum, maxSum int
	for s.Scan() {

		if s.Text() == "" {
			if sum > maxSum {
				maxSum = sum
			}
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
	if sum > maxSum {
		maxSum = sum
	}

	return maxSum
}
