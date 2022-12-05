package day3

import (
	"bufio"
	"io"
	"strings"

	"golang.org/x/exp/slices"
)

func SumPrioGroups(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(scanLineGroup(3))

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

func scanLineGroup(groupSize int) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		var newLCnt int
		for advance = 0; advance < len(data); advance++ {

			endOfRucksack := data[advance] == '\n' || (len(data[advance+1:]) == 0 && atEOF)
			if endOfRucksack {
				newLCnt++
			}

			if newLCnt == groupSize {
				return advance + 1, token, nil // + 1 because we skip the last \n
			}

			if newLCnt < groupSize && len(data[advance+1:]) == 0 {
				return 0, nil, nil
			}

			token = append(token, data[advance])
		}

		return advance, token, nil

	}
}
