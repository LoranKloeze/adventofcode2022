package day8

import (
	"io"
)

func visibleFromOutside(r io.Reader) int {
	var visible int
	grid := parseGrid(r)
	for y, row := range grid {
		for x, _ := range row {
			if grid.isVisible(y, x) {
				visible++
			}
		}
	}
	return visible
}
