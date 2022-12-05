package day3

import (
	"bufio"
	"io"
	"strings"

	"golang.org/x/exp/slices"
)

func SumPrioCompartments(r io.Reader) int {
	s := bufio.NewScanner(r)
	var tot int

	for s.Scan() {
		str := s.Text()

		lCompartment := str[:len(str)/2]
		rCompartment := str[len(str)/2:]

		lSlice := strings.Split(lCompartment, "")
		rSlice := strings.Split(rCompartment, "")

		for _, l := range lSlice {
			if slices.Contains(rSlice, l) {
				tot += prioForItem(l)
				break
			}
		}

	}
	return tot
}
