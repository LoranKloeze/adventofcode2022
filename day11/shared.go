// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day11

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parseMonkeys(r io.Reader) (map[int]*Monkey, int) {
	monkeys := map[int]*Monkey{}
	totalMod := 1

	s := bufio.NewScanner(r)

	type connection struct {
		from      int
		to        int
		condition string
	}

	connections := []connection{}

	for {
		if !s.Scan() {
			break
		}
		monkey := Monkey{Id: parseId(s.Text())}

		s.Scan()
		monkey.Items = parseItems(s.Text())

		s.Scan()
		monkey.Operation = parseOperation(s.Text())

		s.Scan()
		monkey.DivisibleBy = parseDivisible(s.Text())
		totalMod *= monkey.DivisibleBy

		s.Scan()
		connections = append(connections, connection{monkey.Id, parseTrueCase(s.Text()), "ifTrue"})

		s.Scan()
		connections = append(connections, connection{monkey.Id, parseFalseCase(s.Text()), "ifFalse"})

		s.Scan() // Skip newline

		monkeys[monkey.Id] = &monkey
	}

	for _, c := range connections {
		switch c.condition {
		case "ifTrue":
			monkeys[c.from].TrueTo = monkeys[c.to]
		case "ifFalse":
			monkeys[c.from].FalseTo = monkeys[c.to]
		}
	}

	return monkeys, totalMod
}

func parseId(s string) int {
	res := regexp.MustCompile(`Monkey (\d)`).FindStringSubmatch(s)
	id, err := strconv.Atoi(res[1])
	if err != nil {
		log.Fatalf("Expected a number: %v", res[1])
	}
	return id
}

func parseItems(s string) []int {
	res := regexp.MustCompile(`Starting items: (.*)`).FindStringSubmatch(s)
	spl := strings.Split(res[1], ", ")
	items := []int{}
	for _, sp := range spl {
		id, err := strconv.Atoi(sp)
		if err != nil {
			log.Fatalf("Expected a number: %v", res[1])
		}
		items = append(items, id)

	}
	return items
}

func parseOperation(s string) string {
	res := regexp.MustCompile(`Operation: new = old (.*)`).FindStringSubmatch(s)
	return res[1]
}

func parseDivisible(s string) int {
	res := regexp.MustCompile(`Test: divisible by (\d*)`).FindStringSubmatch(s)
	d, err := strconv.Atoi(res[1])
	if err != nil {
		log.Fatalf("Expected a number: %v", res[1])
	}
	return d
}

func parseTrueCase(s string) int {
	res := regexp.MustCompile(`If true: throw to monkey (\d)`).FindStringSubmatch(s)
	id, err := strconv.Atoi(res[1])
	if err != nil {
		log.Fatalf("Expected a number: %v", res[1])
	}
	return id
}

func parseFalseCase(s string) int {
	res := regexp.MustCompile(`If false: throw to monkey (\d)`).FindStringSubmatch(s)
	id, err := strconv.Atoi(res[1])
	if err != nil {
		log.Fatalf("Expected a number: %v", res[1])
	}
	return id
}
