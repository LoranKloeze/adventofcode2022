// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"bytes"
	"os"
	"testing"
)

func TestSampleForTwo(t *testing.T) {
	const sample = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

	b := bytes.NewBufferString(sample)
	got := sizeOfDirToDelete(b)
	exp := 24933642
	if got != exp {
		t.Errorf("Wrong size for dir to delete, expected %d, got %d", exp, got)
	}

}

func TestRealForTwo(t *testing.T) {

	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Failed to open input data: %v\n", err)
		return
	}
	defer f.Close()

	got := sizeOfDirToDelete(f)

	// Your answer is probably different
	exp := 5649896

	if got != exp {
		t.Errorf("Wrong size for dir to delete, expected %d, got %d", exp, got)
	}

}
