// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/lorankloeze/advent2022/day1"
	"github.com/lorankloeze/advent2022/day2"
)

func main() {
	challenges := map[string]func(){
		"day1-1": day1.PartOne,
		"day1-2": day1.PartTwo,
		"day2-1": day2.PartOne,
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <challenge e.g. day1-1>\n", os.Args[0])
		os.Exit(1)
	}

	err := runChallenge(os.Args[1], challenges)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running challenge: %v\n", err)
		os.Exit(1)
	}
}

func runChallenge(challenge string, challenges map[string]func()) error {
	if fn, ok := challenges[challenge]; ok {
		fn()
		return nil
	} else {
		return fmt.Errorf("challenge %q not found", challenge)
	}
}
