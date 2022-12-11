// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"bytes"
	"testing"
)

func TestSampleForOne(t *testing.T) {
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
	got := sumOfDirsUnder100000(b)
	exp := 95437
	if got != exp {
		t.Errorf("Wrong size sum of dirs with size < 100,000, expected %d, got %d", exp, got)
	}

}

func TestFindDir(t *testing.T) {

	r := Entry{Name: "", Type: DirEntry}
	rA := Entry{Name: "a", Type: DirEntry, Parent: &r}
	rB := Entry{Name: "b", Type: DirEntry, Parent: &r}
	rC := Entry{Name: "c", Type: DirEntry, Parent: &r}
	r.Children = append(r.Children, &rA, &rB, &rC)

	rAX := Entry{Name: "x", Type: DirEntry, Parent: &rA}
	rAY := Entry{Name: "y", Type: DirEntry, Parent: &rA}
	rA.Children = append(rA.Children, &rAX, &rAY)

	rAYC := Entry{Name: "c", Type: DirEntry, Parent: &rAY}
	rAY.Children = append(rAY.Children, &rAYC)

	rBP := Entry{Name: "p", Type: DirEntry, Parent: &rB}
	rB.Children = append(rB.Children, &rBP)

	tests := []struct {
		path     string
		expEntry *Entry
		expOk    bool
	}{
		{path: "/a/y/c", expEntry: &rAYC, expOk: true},
		{path: "/b/p", expEntry: &rBP, expOk: true},
		{path: "/non/exist", expEntry: nil, expOk: false},
		{path: "/a/non/exist", expEntry: nil, expOk: false},
		{path: "/", expEntry: &r, expOk: true},
	}
	for _, tc := range tests {
		gotDir, gotOk := findDir(&r, tc.path)

		if gotDir != tc.expEntry {
			t.Errorf("Wrong dir returned by findDir() for path %q, expected %v, got %v", tc.path, tc.expEntry, gotDir)
		}

		if gotOk != tc.expOk {
			t.Errorf("Wronk ok value returned by findDir() for path %q, expected %v, got %v", tc.path, tc.expOk, gotOk)
		}
	}

}

func TestTreeParse(t *testing.T) {
	const input = `$ cd /
$ ls
dir y
100 blabla.xslx
dir x
$ cd y
$ ls
dir a
dir b
1337 afile
$ cd a
42 log.txt
$ cd ..
$ cd ..
$ cd x
dir v
135533 setup.exe
48 autoexec.bat
`

	b := bytes.NewBufferString(input)
	root := parseTree(b)

	got := len(root.Children)
	exp := 3
	if got != exp {
		t.Errorf("Wrong number of children for root, expected %d, got %d", exp, got)
	}

}
