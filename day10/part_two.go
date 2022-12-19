// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day10

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day10/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := findLetters(f)

	fmt.Println("The answer is")
	fmt.Printf("%s\n", res)

}

func findLetters(r io.Reader) string {
	var sb strings.Builder
	instructions := parseInstructions(r)

	var ip int
	var inAddX bool

	maxCycles := 240
	cycle := 1
	spritePos := 1

	sb.WriteString(determinePixel(cycle, spritePos))

	for {

		if cycle%40 == 0 {
			sb.WriteString("\n")
		}

		if inAddX {
			spritePos += instructions[ip].Value
			inAddX = false
			ip++
		} else {
			if instructions[ip].Operation == "addx" {
				inAddX = true
			}
			if instructions[ip].Operation == "noop" {
				ip++
			}
		}

		sb.WriteString(determinePixel(cycle, spritePos))

		cycle++
		if cycle >= maxCycles {
			break
		}
	}
	return sb.String()
}

func determinePixel(cycle, spritePos int) string {

	if (cycle%40) >= spritePos-1 && (cycle%40) <= spritePos+1 {
		return "#"
	} else {
		return "."
	}
}
