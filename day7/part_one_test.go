// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"bytes"
	"fmt"
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

func TestFindEntry(t *testing.T) {

	r := Entry{Name: "", Type: DirEntry}
	rA := Entry{Name: "a", Type: DirEntry, Parent: &r}
	rB := Entry{Name: "b", Type: DirEntry, Parent: &r}
	rC := Entry{Name: "c", Type: DirEntry, Parent: &r}
	r.Children = append(r.Children, &rA, &rB, &rC)

	rAX := Entry{Name: "x", Type: DirEntry, Parent: &rA}
	rAY := Entry{Name: "y", Type: DirEntry, Parent: &rA}
	rAFile := Entry{Name: "setup.exe", Type: FileEntry, Parent: &rA}
	rA.Children = append(rA.Children, &rAX, &rAY, &rAFile)

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
		{path: "/a/setup.exe", expEntry: &rAFile, expOk: true},
		{path: "/non/exist", expEntry: nil, expOk: false},
		{path: "/a/non/exist", expEntry: nil, expOk: false},
		{path: "/", expEntry: &r, expOk: true},
	}
	for _, tc := range tests {
		gotDir, gotOk := findEntry(&r, tc.path)

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

	tests := []struct {
		path    string
		expType EntryType
		expSize int
		expName string
	}{
		{path: "/", expType: DirEntry, expSize: 0, expName: ""},
		{path: "/y", expType: DirEntry, expSize: 0, expName: "y"},
		{path: "/x", expType: DirEntry, expSize: 0, expName: "x"},
		{path: "/blabla.xslx", expType: FileEntry, expSize: 100, expName: "blabla.xslx"},
		{path: "/y/a", expType: DirEntry, expSize: 0, expName: "a"},
		{path: "/y/b", expType: DirEntry, expSize: 0, expName: "b"},
		{path: "/y/afile", expType: FileEntry, expSize: 1337, expName: "afile"},
		{path: "/y/a/log.txt", expType: FileEntry, expSize: 42, expName: "log.txt"},
		{path: "/x/v", expType: DirEntry, expSize: 0, expName: "v"},
		{path: "/x/setup.exe", expType: FileEntry, expSize: 135533, expName: "setup.exe"},
		{path: "/x/autoexec.bat", expType: FileEntry, expSize: 48, expName: "autoexec.bat"},
	}

	b := bytes.NewBufferString(input)
	root, err := parseTree(b)
	if err != nil {
		t.Fatalf("Unexpected error while parsing test tree")
	}

	for _, tc := range tests {
		entry, ok := findEntry(root, tc.path)
		if !ok {
			t.Errorf("Expected path %q to exist but it doesn't", tc.path)
		}

		if entry.Size != tc.expSize {
			t.Errorf("Expected %q to have size %d but got %d", tc.path, tc.expSize, entry.Size)
		}

		if entry.Type != tc.expType {
			t.Errorf("Expected %q to have type %v but got %v", tc.path, tc.expType, entry.Type)
		}

		if entry.Name != tc.expName {
			t.Errorf("Expected %q to have name %q but got %q", tc.path, tc.expName, entry.Name)
		}

	}

}

func TestInvalidTreeParse(t *testing.T) {
	const input = `$ cd /
$ ls
dqdqw blabla.xslx
`
	b := bytes.NewBufferString(input)
	_, err := parseTree(b)
	if err == nil {
		t.Errorf("Expected error parsing an invalid tree, got none")
	}
	fmt.Println(err)

}

func TestUnkownDirInTreeParse(t *testing.T) {
	const input = `$ cd /
$ ls
dir y
dir x
$ cd iamnothere
$ ls
dir a
dir b
`
	b := bytes.NewBufferString(input)
	_, err := parseTree(b)
	if err == nil {
		t.Errorf("Expected error trying to cd to a non-existing dir, got none")
	}
	fmt.Println(err)

}
