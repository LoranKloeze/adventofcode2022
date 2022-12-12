// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day8

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Grid [][]int

func (g Grid) maxY() int {
	return len(g) - 1
}

func (g Grid) maxX() int {
	if len(g) > 0 {
		return len(g[0]) - 1
	} else {
		return 0
	}
}

func (g Grid) isVisible(treeY, treeX int) bool {

	// Each tree around the edge of the grid is always visible
	if treeY == 0 || treeY == g.maxY() || treeX == 0 || treeX == g.maxX() {
		return true
	}

	var x, y int
	tree := g[treeY][treeX]
	visiblePositions := 4

	// Looking from the left
	for x = 0; x < treeX; x++ {
		if g[treeY][x] >= tree {
			visiblePositions--
			break
		}
	}

	// Looking from the right
	for x = treeX + 1; x <= g.maxX(); x++ {
		if g[treeY][x] >= tree {
			visiblePositions--
			break
		}
	}

	// Looking from the top
	for y = 0; y < treeY; y++ {
		if g[y][treeX] >= tree {
			visiblePositions--
			break
		}
	}
	// Looking from the bottom
	for y = treeY + 1; y <= g.maxY(); y++ {
		if g[y][treeX] >= tree {
			visiblePositions--
			break
		}
	}

	return visiblePositions > 0
}

func (g Grid) scenicScore(treeY, treeX int) int {
	var x, y, scoreL, scoreR, scoreT, scoreB int
	tree := g[treeY][treeX]

	// Looking towards the left
	for x = treeX - 1; x >= 0; x-- {
		scoreL++
		if g[treeY][x] >= tree {
			break
		}
	}

	// Looking towards the right
	for x = treeX + 1; x <= g.maxX(); x++ {
		scoreR++
		if g[treeY][x] >= tree {
			break
		}
	}

	// Looking towards the top
	for y = treeY - 1; y >= 0; y-- {
		scoreT++
		if g[y][treeX] >= tree {
			break
		}
	}

	// Looking towards the bottom
	for y = treeY + 1; y <= g.maxY(); y++ {
		scoreB++
		if g[y][treeX] >= tree {
			break
		}
	}

	return scoreL * scoreR * scoreB * scoreT
}

func parseGrid(r io.Reader) Grid {
	s := bufio.NewScanner(r)
	grid := Grid{}

	for s.Scan() {
		row := []int{}
		for _, c := range strings.Split(s.Text(), "") {
			nr, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("Expected a number: %v", s.Text())
			}
			row = append(row, nr)
		}
		grid = append(grid, row)
	}
	return grid
}
