package day5

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const maxStacks = 10

// Let's assume all stacks are placed on a platform.
// In real life we never assume but need to ask our users, here we don't.
type stacksPlatform [][]string

func (stacks stacksPlatform) pop(stack int, amount int) []string {

	l := len(stacks[stack])
	crates := make([]string, amount)
	copy(crates, stacks[stack][l-amount:])
	stacks[stack] = stacks[stack][:l-amount]

	return crates

}

func (stacks stacksPlatform) parseInitalCrates(row string) {
	re := regexp.MustCompile(`(\[[A-Z]\]|\s{3})\s?`)
	p := re.FindAllStringSubmatch(row, -1)

	for i, m := range p {
		if !strings.HasPrefix(m[1], "[") {
			continue
		}

		crate := string(m[1][1]) // Extract X from "[X]"
		stackIdx := i + 1        // Stacks start with number 1

		stacks[stackIdx] = append([]string{crate}, stacks[stackIdx]...) // Prepending, not appending
	}
}

func (stacks stacksPlatform) moveCrates(row string) error {

	spl := strings.Split(row, " ")

	amount, err := strconv.Atoi(spl[1])
	if err != nil {
		return fmt.Errorf("[stack amount] Cannot convert %s to a number", spl[1])
	}

	from, err := strconv.Atoi(spl[3])
	if err != nil {
		return fmt.Errorf("[stack from] Cannot convert %s to a number ", spl[3])
	}

	to, err := strconv.Atoi(spl[5])
	if err != nil {
		return fmt.Errorf("[stack to] Cannot convert %s to a number", spl[5])
	}

	for i := 0; i < amount; i++ {
		crate := stacks.pop(from, 1)[0]
		stacks[to] = append(stacks[to], crate)
	}

	return nil
}

func (stacks stacksPlatform) topCrates() (res string) {
	for _, stack := range stacks {
		if len(stack) > 0 {
			res += stack[len(stack)-1:][0]
		}
	}
	return res
}

func resultOfCrateMover9000(r io.Reader) string {
	s := bufio.NewScanner(r)

	stacks := make(stacksPlatform, maxStacks)

	for s.Scan() {
		switch {
		case strings.Contains(s.Text(), "["):
			stacks.parseInitalCrates(s.Text())
		case strings.HasPrefix(s.Text(), "move"):
			err := stacks.moveCrates(s.Text())
			if err != nil {
				log.Fatalf("Error parsing rows with move instructions: %v", err)
			}
		}
	}

	return stacks.topCrates()
}
