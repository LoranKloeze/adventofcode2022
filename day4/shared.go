// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day4

import (
	"fmt"
	"regexp"
	"strconv"
)

type pair struct {
	begin int
	end   int
}

func extractPairs(s string) (pair, pair, error) {
	re := regexp.MustCompile(`(\d*)-(\d*),(\d*)-(\d*)`)
	sm := re.FindStringSubmatch(s)

	if len(sm) != 5 {
		return pair{}, pair{}, fmt.Errorf("cannot extract all numbers from %q", s)
	}

	pairOne, pairTwo := pair{}, pair{}

	// Loop over 4 submatches in re
	for i := 1; i < 5; i++ {
		nr, err := strconv.Atoi(sm[i])
		if err != nil {
			return pair{}, pair{}, fmt.Errorf("%q contains non-numbers", s)
		}
		switch i {
		case 1:
			pairOne.begin = nr
		case 2:
			pairOne.end = nr
		case 3:
			pairTwo.begin = nr
		case 4:
			pairTwo.end = nr
		}
	}

	return pairOne, pairTwo, nil

}
