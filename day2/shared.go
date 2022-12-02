// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day2

var shapeForCode map[string]string = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var scoreForShape map[string]int = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}
var scoreForOutcome map[string]int = map[string]int{
	"lost": 0,
	"draw": 3,
	"won":  6,
}
