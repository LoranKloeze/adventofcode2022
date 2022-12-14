// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day9

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) adjacentTo(a Point) bool {
	dX := math.Abs(float64(p.X - a.X))
	dY := math.Abs(float64(p.Y - a.Y))
	return dX <= 1 && dY <= 1
}

func (p Point) key() string {
	return fmt.Sprintf("%d,%d", p.Y, p.X)
}

func uniqueTailVisits(r io.Reader) int {
	head := Point{}
	tail := Point{}
	tailPositions := map[string]struct{}{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		spl := strings.Split(s.Text(), " ")
		direction := spl[0]

		steps, err := strconv.Atoi(spl[1])
		if err != nil {
			log.Fatalf("Expected a number: %v", spl[1])
		}

		for step := 0; step < steps; step++ {
			processStep(direction, &head, &tail)
			tailPositions[tail.key()] = struct{}{}
		}

	}
	return len(tailPositions)
}

func processStep(direction string, head, tail *Point) {
	switch direction {
	case "R":
		head.X++
		if !tail.adjacentTo(*head) {
			tail.X++
			tail.Y = head.Y
		}
	case "L":
		head.X--
		if !tail.adjacentTo(*head) {
			tail.X--
			tail.Y = head.Y
		}
	case "U":
		head.Y--
		if !tail.adjacentTo(*head) {
			tail.Y--
			tail.X = head.X
		}
	case "D":
		head.Y++
		if !tail.adjacentTo(*head) {
			tail.Y++
			tail.X = head.X
		}
	}
}
