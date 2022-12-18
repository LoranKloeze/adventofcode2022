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
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	X    int
	Y    int
	Next *Knot
}

func (k Knot) adjacentTo(other *Knot) bool {
	return math.Abs(float64(k.X)-float64(other.X)) <= 1 && math.Abs(float64(k.Y)-float64(other.Y)) <= 1
}

func (k Knot) key() string {
	return fmt.Sprintf("%d,%d", k.Y, k.X)
}

type Rope []Knot

// Main entry for part one of this day
func PartOne() {
	f, err := os.Open("day9/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := tailVisits(f, 2)

	fmt.Printf("The answer is %d\n", res)

}

// Main entry for part two of this day
func PartTwo() {
	f, err := os.Open("day9/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	res := tailVisits(f, 10)

	fmt.Printf("The answer is %d\n", res)

}

func initRope(knots int) Rope {
	rope := Rope{}
	var lastKnot *Knot

	for i := 0; i < knots; i++ {
		knot := Knot{}
		if lastKnot != nil {
			knot.Next = lastKnot
		}
		rope = append(Rope{knot}, rope...)
		lastKnot = &knot
	}
	return rope
}

func tailVisits(r io.Reader, knots int) int {
	rope := initRope(knots)
	visited := map[string]bool{}

	s := bufio.NewScanner(r)

	for s.Scan() {
		spl := strings.Split(s.Text(), " ")
		dir := spl[0]
		steps, err := strconv.Atoi(spl[1])
		if err != nil {
			log.Fatalf("Expected a number: %v", s.Text())
		}

		for i := 0; i < steps; i++ {
			head := &rope[0]
			switch dir {
			case "L":
				head.X--
			case "U":
				head.Y++
			case "R":
				head.X++
			case "D":
				head.Y--
			}

			for idx := range rope {
				head := rope[idx]
				if head.Next == nil {
					visited[head.key()] = true
				} else {
					tail := &rope[idx+1]
					if head.adjacentTo(tail) {
						continue
					}

					repositionTail(head, tail)
				}
			}

		}
	}

	return len(visited)
}

func repositionTail(head Knot, tail *Knot) {
	dX := head.X - tail.X
	dXAbs := int(math.Abs(float64(dX)))
	dY := head.Y - tail.Y
	dYAbs := int(math.Abs(float64(dY)))

	if dXAbs <= 1 {
		tail.Y += dY / dYAbs

		if dXAbs > 0 {
			tail.X += dX / dXAbs
		}
	} else {
		tail.X += dX / dXAbs

		if dYAbs > 0 {
			tail.Y += dY / dYAbs
		}

	}

}
