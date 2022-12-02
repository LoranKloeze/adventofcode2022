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

var shapeForCode map[string]string = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var scoreForShape map[string]int = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}
var scoreForOutcome map[string]int = map[string]int{
	"lost": 0,
	"draw": 3,
	"won":  6,
}

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := ProbablyMyScore(f)

	fmt.Printf("The answer is %d\n", res)

}

func ProbablyMyScore(r io.Reader) (totalScore int) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		spl := strings.Split(s.Text(), " ")
		opponent := spl[0]
		me := spl[1]

		score, err := parseProbablyMove(opponent, me)
		if err != nil {
			fmt.Printf("Warning: cannot parse move: %v", err)
		}

		totalScore += score
	}
	return totalScore
}

func parseProbablyMove(opponentCode, meCode string) (int, error) {

	opponent, ok := shapeForCode[opponentCode]
	if !ok {
		return 0, fmt.Errorf("code %q is not mappable to a shape", opponentCode)
	}

	me, ok := shapeForCode[meCode]
	if !ok {
		return 0, fmt.Errorf("code %q is not mappable to a shape", opponentCode)
	}

	switch {
	case opponent == "rock" && me == "paper":
		return scoreForShape[me] + scoreForOutcome["won"], nil
	case opponent == "rock" && me == "scissors":
		return scoreForShape[me] + scoreForOutcome["lost"], nil
	case opponent == "rock" && me == "rock":
		return scoreForShape[me] + scoreForOutcome["draw"], nil
	case opponent == "scissors" && me == "paper":
		return scoreForShape[me] + scoreForOutcome["lost"], nil
	case opponent == "scissors" && me == "scissors":
		return scoreForShape[me] + scoreForOutcome["draw"], nil
	case opponent == "scissors" && me == "rock":
		return scoreForShape[me] + scoreForOutcome["won"], nil
	case opponent == "paper" && me == "paper":
		return scoreForShape[me] + scoreForOutcome["draw"], nil
	case opponent == "paper" && me == "scissors":
		return scoreForShape[me] + scoreForOutcome["won"], nil
	case opponent == "paper" && me == "rock":
		return scoreForShape[me] + scoreForOutcome["lost"], nil
	default:
		return 0, fmt.Errorf("cannot parse opponent %q and me %q", opponentCode, meCode)
	}

}
