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

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := probablyMyScore(f)

	fmt.Printf("The answer is %d\n", res)

}

func probablyMyScore(r io.Reader) (totalScore int) {
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
