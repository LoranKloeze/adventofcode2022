package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := sumPrioGroups(f)

	fmt.Printf("The answer is %d\n", res)

}

func sumPrioGroups(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(groupSplit)

	var tot int
	for s.Scan() {
		rucksacks := strings.Split(s.Text(), "\n")

		for _, item := range rucksacks[0] {
			if slices.Contains([]rune(rucksacks[1]), item) && slices.Contains([]rune(rucksacks[2]), item) {
				tot += prioForItem(item)
				break
			}
		}

	}
	return tot
}

func groupSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var groupSize = 3

	var rucksackCnt int
	for advance = 0; advance < len(data); advance++ {
		endOfData := len(data[advance+1:]) == 0
		endOfRucksack := data[advance] == '\n' || (endOfData && atEOF)

		if endOfRucksack {
			rucksackCnt++
		}

		if rucksackCnt == groupSize {
			// + 1 because we skip the last \n
			return advance + 1, token, nil
		}

		if rucksackCnt < groupSize && endOfData {
			// No more data left but we haven't found all rucksacks yet, ask for a larger data slice
			return 0, nil, nil
		}

		token = append(token, data[advance])
	}

	return advance, token, nil

}
