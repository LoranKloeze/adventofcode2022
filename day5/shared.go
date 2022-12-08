// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Let's assume all stacks are placed on a platform.
// In real life we never assume but need to ask our users, here we don't.
type stacksPlatform [][]string

// Constants for the moving strategy
const perCrate = 1
const perGroup = 2

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

func (stacks stacksPlatform) moveCrates(row string, strategy int) error {
	spl := strings.Split(row, " ")

	amount, err := strconv.Atoi(spl[1])
	if err != nil {
		return fmt.Errorf("[stack amount] Cannot convert %s to a number: %v", spl[1], err)
	}

	from, err := strconv.Atoi(spl[3])
	if err != nil {
		return fmt.Errorf("[stack from] Cannot convert %s to a number: %v", spl[3], err)
	}

	to, err := strconv.Atoi(spl[5])
	if err != nil {
		return fmt.Errorf("[stack to] Cannot convert %s to a number: %v", spl[5], err)
	}

	switch strategy {
	case perCrate:
		for i := 0; i < amount; i++ {
			crate := stacks.pop(from, 1)[0]
			stacks[to] = append(stacks[to], crate)
		}
	case perGroup:
		crates := stacks.pop(from, amount)
		stacks[to] = append(stacks[to], crates...)
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
