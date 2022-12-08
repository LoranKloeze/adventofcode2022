package day6

import (
	"bufio"
	"fmt"
	"io"
)

func startOfPacketMarkerPos(r io.Reader) (int, error) {
	groupSize := 4

	s := bufio.NewScanner(r)
	s.Split(scanGrouped(groupSize))

	endPos := groupSize
	for s.Scan() {
		if len(s.Bytes()) > 0 && unique(s.Bytes()) {
			return endPos, nil
		}
		endPos++
	}

	return 0, fmt.Errorf("no start-of-packet found")
}

func scanGrouped(groupSize int) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if len(data) < groupSize {
			return 0, nil, bufio.ErrFinalToken
		}
		return 1, data[:groupSize], nil
	}
}

func unique(s []byte) bool {
	seen := map[byte]bool{}

	for _, b := range s {
		if _, ok := seen[b]; ok {
			return false
		}
		seen[b] = true
	}

	return true
}
