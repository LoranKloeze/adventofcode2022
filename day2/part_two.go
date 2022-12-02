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

var winsFrom map[string]string = map[string]string{
	"rock":     "paper",
	"paper":    "scissors",
	"scissors": "rock",
}

var losesTo map[string]string = map[string]string{
	"rock":     "scissors",
	"paper":    "rock",
	"scissors": "paper",
}

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

	if meCode == "Z" {
		shapeToWin := winsFrom[opponent]
		return scoreForShape[shapeToWin] + scoreForOutcome["won"], nil
	}

	if meCode == "Y" {
		return scoreForShape[opponent] + scoreForOutcome["draw"], nil
	}

	if meCode == "X" {
		shapeToLose := losesTo[opponent]
		return scoreForShape[shapeToLose] + scoreForOutcome["lost"], nil
	}

	return 0, fmt.Errorf("cannot parse opponent %q and me %q", opponentCode, meCode)

}
