// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day3

func prioForItem(item rune) int {

	if item > 96 {
		return int(item - 96)
	} else {
		return int(item - 38)
	}
}
