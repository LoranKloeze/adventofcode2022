package day10

import (
	"fmt"
	"io"
	"os"
)

type Instruction struct {
	Operation string
	Value     int
}

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day10/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := sumSixSignalStrengths(f)

	fmt.Printf("The answer is %d\n", res)

}

func sumSixSignalStrengths(r io.Reader) int {

	instructions := parseInstructions(r)

	var sum int
	var ip int
	var inAddX bool

	maxCycles := 240
	regX := 1
	cycle := 1
	for {

		// Each 40th cycle, starting at 20
		if cycle%40 == 20 {
			sum += cycle * regX
		}

		if inAddX {
			regX += instructions[ip].Value
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

		cycle++
		if cycle >= maxCycles {
			break
		}
	}
	return sum
}
