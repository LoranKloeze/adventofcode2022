// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/lorankloeze/advent2022/day1"
	"github.com/lorankloeze/advent2022/day10"
	"github.com/lorankloeze/advent2022/day11"
	"github.com/lorankloeze/advent2022/day2"
	"github.com/lorankloeze/advent2022/day3"
	"github.com/lorankloeze/advent2022/day4"
	"github.com/lorankloeze/advent2022/day5"
	"github.com/lorankloeze/advent2022/day6"
	"github.com/lorankloeze/advent2022/day7"
	"github.com/lorankloeze/advent2022/day8"
	"github.com/lorankloeze/advent2022/day9"
)

func main() {
	challenges := map[string]func(){
		"day1-1":  day1.PartOne,
		"day1-2":  day1.PartTwo,
		"day2-1":  day2.PartOne,
		"day2-2":  day2.PartTwo,
		"day3-1":  day3.PartOne,
		"day3-2":  day3.PartTwo,
		"day4-1":  day4.PartOne,
		"day4-2":  day4.PartTwo,
		"day5-1":  day5.PartOne,
		"day5-2":  day5.PartTwo,
		"day6-1":  day6.PartOne,
		"day6-2":  day6.PartTwo,
		"day7-1":  day7.PartOne,
		"day7-2":  day7.PartTwo,
		"day8-1":  day8.PartOne,
		"day8-2":  day8.PartTwo,
		"day9-1":  day9.PartOne,
		"day9-2":  day9.PartTwo,
		"day10-1": day10.PartOne,
		"day10-2": day10.PartTwo,
		"day11-1": day11.PartOne,
		"day11-2": day11.PartTwo,
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
