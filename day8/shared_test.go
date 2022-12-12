package day8

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseGrid(t *testing.T) {
	input := `37487
83685
98423
`
	s := bytes.NewBufferString(input)
	got := parseGrid(s)
	exp := Grid{
		{3, 7, 4, 8, 7},
		{8, 3, 6, 8, 5},
		{9, 8, 4, 2, 3},
	}

	if !cmp.Equal(got, exp) {
		t.Errorf("Expected parseGrid() to produce a correct grid, expected %v, got %v", exp, got)
	}

}

func TestMaxX(t *testing.T) {
	const sample = `12345
12345
12345
`
	s := bytes.NewBufferString(sample)
	grid := parseGrid(s)

	exp := 4
	got := grid.maxX()
	if got != exp {
		t.Errorf("Expected maxX() to return %d, got %d", exp, got)
	}

}
func TestMaxY(t *testing.T) {
	const sample = `12345
12345
12345
`
	s := bytes.NewBufferString(sample)
	grid := parseGrid(s)

	exp := 2
	got := grid.maxY()
	if got != exp {
		t.Errorf("Expected maxY() to return %d, got %d", exp, got)
	}

}

func TestIsVisible(t *testing.T) {
	const sample = `30373
25512
65332
33549
35390
`
	s := bytes.NewBufferString(sample)
	grid := parseGrid(s)

	tests := []struct {
		x          int
		y          int
		expVisible bool
	}{
		{y: 0, x: 0, expVisible: true},
		{y: 0, x: 1, expVisible: true},
		{y: 0, x: 2, expVisible: true},
		{y: 0, x: 3, expVisible: true},
		{y: 0, x: 4, expVisible: true},
		{y: 1, x: 0, expVisible: true},
		{y: 1, x: 1, expVisible: true},
		{y: 1, x: 2, expVisible: true},
		{y: 1, x: 3, expVisible: false},
		{y: 1, x: 4, expVisible: true},
		{y: 2, x: 0, expVisible: true},
		{y: 2, x: 1, expVisible: true},
		{y: 2, x: 2, expVisible: false},
		{y: 2, x: 3, expVisible: true},
		{y: 2, x: 4, expVisible: true},
		{y: 3, x: 0, expVisible: true},
		{y: 3, x: 1, expVisible: false},
		{y: 3, x: 2, expVisible: true},
		{y: 3, x: 3, expVisible: false},
		{y: 3, x: 4, expVisible: true},
		{y: 4, x: 0, expVisible: true},
		{y: 4, x: 1, expVisible: true},
		{y: 4, x: 2, expVisible: true},
		{y: 4, x: 3, expVisible: true},
		{y: 4, x: 4, expVisible: true},
	}

	for _, tc := range tests {
		got := grid.isVisible(tc.y, tc.x)
		if got != tc.expVisible {
			t.Errorf("Expected tree @ %d,%d visibility to be %t but it's %t", tc.y, tc.x, tc.expVisible, got)
		}
	}

}
