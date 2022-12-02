// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := MyActualScore(f)

	fmt.Printf("The answer is %d\n", res)

}

func MyActualScore(r io.Reader) (totalScore int) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		spl := strings.Split(s.Text(), " ")
		opponent := spl[0]
		me := spl[1]

		score, err := parseActualMove(opponent, me)
		if err != nil {
			fmt.Printf("Warning: cannot parse move: %v\n", err)
		}

		totalScore += score
	}
	return totalScore
}

func parseActualMove(opponentCode, meCode string) (int, error) {

	opponent, ok := shapeForCode[opponentCode]
	if !ok {
		return 0, fmt.Errorf("code %q is not mappable to a shape", opponentCode)
	}

	// Win
	if meCode == "Z" {
		if opponent == "rock" {
			return scoreForShape["paper"] + scoreForOutcome["won"], nil
		}
		if opponent == "paper" {
			return scoreForShape["scissors"] + scoreForOutcome["won"], nil
		}
		if opponent == "scissors" {
			return scoreForShape["rock"] + scoreForOutcome["won"], nil
		}
	}

	// Draw
	if meCode == "Y" {
		return scoreForShape[opponent] + scoreForOutcome["draw"], nil
	}

	// Lose
	if meCode == "X" {
		if opponent == "rock" {
			return scoreForShape["scissors"] + scoreForOutcome["lose"], nil
		}
		if opponent == "paper" {
			return scoreForShape["rock"] + scoreForOutcome["lose"], nil
		}
		if opponent == "scissors" {
			return scoreForShape["paper"] + scoreForOutcome["lose"], nil
		}
	}

	return 0, fmt.Errorf("cannot parse opponent %q and me %q", opponentCode, meCode)

}
