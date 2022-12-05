package day3

import (
	"bufio"
	"io"

	"golang.org/x/exp/slices"
)

func SumPrioCompartments(r io.Reader) int {
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
