package day11

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Id           int
	Items        []int
	Operation    string
	DivisibleBy  int
	TrueTo       *Monkey
	FalseTo      *Monkey
	InspectedCnt int
}

func (m *Monkey) Throw(divisor string, totalMod int) {
	for len(m.Items) > 0 {
		// Shift first item from monkeys items
		level := m.Items[0]
		m.Items = m.Items[1:]

		m.InspectedCnt++

		// Split up monkeys operation string to make it usable
		spl := strings.Split(m.Operation, " ")
		operator, operandStr, operandNr := spl[0], spl[1], 0
		if operandStr != "old" {
			var err error
			operandNr, err = strconv.Atoi(operandStr)
			if err != nil {
				log.Fatalf("Expected a number: %v", operandStr)
			}
		}

		switch operator {
		case "*":
			if operandStr == "old" {
				level *= level
			} else {
				level *= operandNr
			}
		case "+":
			if operandStr == "old" {
				level += level
			} else {
				level += operandNr
			}
		}

		// Needed to cater for part 1 or part 2 of the challenge
		if divisor == "part1" {
			level /= 3
		} else {
			level = level % totalMod
		}

		isDivisible := level%m.DivisibleBy == 0
		if isDivisible {
			m.TrueTo.Items = append(m.TrueTo.Items, level)
		} else {
			m.FalseTo.Items = append(m.FalseTo.Items, level)
		}

	}
}

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day11/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := monkeyBusiness(f, 20, "part1")

	fmt.Printf("The answer is %d\n", res)

}

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day11/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := monkeyBusiness(f, 10000, "part2")

	fmt.Printf("The answer is %d\n", res)

}

func monkeyBusiness(r io.Reader, rounds int, divisor string) int {
	monkeys, totalMod := parseMonkeys(r)

	for i := 0; i < rounds; i++ {
		for mI := 0; mI < len(monkeys); mI++ {
			monkey := monkeys[mI]
			monkey.Throw(divisor, totalMod)
		}
	}

	inspected := []int{}
	for _, m := range monkeys {
		inspected = append(inspected, m.InspectedCnt)
	}
	sort.Ints(inspected)

	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}
