// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day6

import "bufio"

func scanGrouped(groupSize int) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if len(data) < groupSize {
			return 0, nil, bufio.ErrFinalToken
		}
		return 1, data[:groupSize], nil
	}
}

func unique(s []byte) bool {
	seen := map[byte]bool{}

	for _, b := range s {
		if _, ok := seen[b]; ok {
			return false
		}
		seen[b] = true
	}

	return true
}
