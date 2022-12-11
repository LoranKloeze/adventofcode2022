// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day5

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day5/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := resultOfCrateMover9001(f)

	fmt.Printf("The answer is %q\n", res)

}

func resultOfCrateMover9001(r io.Reader) string {
	s := bufio.NewScanner(r)

	stacks := make(stacksPlatform, maxStacks)

	for s.Scan() {
		switch {
		case strings.Contains(s.Text(), "["):
			stacks.parseInitalCrates(s.Text())
		case strings.HasPrefix(s.Text(), "move"):
			err := stacks.moveCrates(s.Text(), perGroup)
			if err != nil {
				log.Fatalf("Error parsing rows with move instructions: %v", err)
			}
		}
	}

	return stacks.topCrates()
}
