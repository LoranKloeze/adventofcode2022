// Copyright 2022 Loran Kloeze. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package day7

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"
)

type EntryType int

const (
	FileEntry EntryType = iota
	DirEntry
)

type Entry struct {
	Name     string
	Type     EntryType
	Size     int
	Parent   *Entry
	Children []*Entry
}

func (e Entry) String() string {
	var t string
	switch e.Type {
	case DirEntry:
		t = "[D]"
	case FileEntry:
		t = "[F]"
	}
	return fmt.Sprintf("%s %s", t, e.Name)
}

func (e *Entry) findSizeAtMost(atMost int, cb func(e *Entry, sz int)) {
	sz := e.fullSize()
	if sz <= atMost {
		cb(e, sz)
	}
	for _, c := range e.Children {
		if c.Type == DirEntry {
			c.findSizeAtMost(atMost, cb)
		}
	}
}

func (e *Entry) findSizeAtLeast(atLeast int, cb func(e *Entry, sz int)) {
	sz := e.fullSize()
	if sz >= atLeast {
		cb(e, sz)
	}
	for _, c := range e.Children {
		if c.Type == DirEntry {
			c.findSizeAtLeast(atLeast, cb)
		}
	}
}

func (e *Entry) fullSize() int {
	size := 0
	for _, c := range e.Children {
		switch c.Type {
		case FileEntry:
			size += c.Size
		case DirEntry:
			size += c.fullSize()
		}
	}
	return size
}

func (e *Entry) find(path string) (*Entry, bool) {
	if path == "/" {
		return e, true
	}

	spl := strings.Split(path, "/")[1:] // Ignore leading slash

	segments := len(spl)

	current := e
outer:
	for _, p := range spl {
		for _, c := range current.Children {
			if c.Name == p {
				current = c
				segments--
				continue outer
			}
		}
	}

	if segments == 0 {
		return current, true
	} else {
		return nil, false
	}

}

func parseTree(r io.Reader) (root *Entry, err error) {
	s := bufio.NewScanner(r)

	root = &Entry{Type: DirEntry, Name: "/"}
	parentDir := root
	pwd := "/"
	for s.Scan() {
		isDir := strings.HasPrefix(s.Text(), "dir ")
		if isDir {
			spl := strings.Split(s.Text(), " ")
			entry := Entry{Name: spl[1], Type: DirEntry, Parent: parentDir}
			parentDir.Children = append(parentDir.Children, &entry)
		}

		isFile := !strings.HasPrefix(s.Text(), "dir ") && !strings.HasPrefix(s.Text(), "$")
		if isFile {
			spl := strings.Split(s.Text(), " ")
			size, err := strconv.Atoi(spl[0])
			if err != nil {
				return nil, fmt.Errorf("tried to parse a non-number: %v", err)
			}
			entry := Entry{Name: spl[1], Type: FileEntry, Size: size, Parent: parentDir}
			parentDir.Children = append(parentDir.Children, &entry)
		}

		if strings.HasPrefix(s.Text(), "$ cd ") {
			spl := strings.Split(s.Text(), " ")
			pwd = filepath.Clean(pwd + "/" + spl[2])
			var ok bool
			parentDir, ok = root.find(pwd)
			if !ok {
				return nil, fmt.Errorf("tried to cd to a non-existing dir: %q", pwd)
			}
		}

	}
	return root, nil
}
