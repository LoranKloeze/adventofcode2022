package day5

import (
	"bufio"
	"io"
	"log"
	"strings"
)

const maxStacks = 10

func resultOfCrateMover9000(r io.Reader) string {
	s := bufio.NewScanner(r)

	stacks := make(stacksPlatform, maxStacks)

	for s.Scan() {
		switch {
		case strings.Contains(s.Text(), "["):
			stacks.parseInitalCrates(s.Text())
		case strings.HasPrefix(s.Text(), "move"):
			err := stacks.moveCrates(s.Text(), perCrate)
			if err != nil {
				log.Fatalf("Error parsing rows with move instructions: %v", err)
			}
		}
	}

	return stacks.topCrates()
}
