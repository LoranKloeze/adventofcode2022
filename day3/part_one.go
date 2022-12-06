package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/exp/slices"
)

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := sumPrioCompartments(f)

	fmt.Printf("The answer is %d\n", res)

}

func sumPrioCompartments(r io.Reader) int {
	s := bufio.NewScanner(r)
	var tot int

	for s.Scan() {
		rucksack := s.Text()

		lCompartment := rucksack[:len(rucksack)/2]
		rCompartment := rucksack[len(rucksack)/2:]

		for _, item := range lCompartment {
			if slices.Contains([]rune(rCompartment), item) {
				tot += prioForItem(item)
				break
			}
		}

	}
	return tot
}
