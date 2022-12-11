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

func (e Entry) FullName() string {
	if e.Parent == nil {
		return e.Name
	}
	segments := []string{e.Name}
	p := e.Parent
	for p != nil {
		if p.Name != "/" {
			segments = append([]string{p.Name}, segments...) // ! Prepending, not appending
		}
		p = p.Parent
	}
	return "/" + strings.Join(segments, "/")

}

func (e *Entry) walkDown() {
	sz := e.FullSize()
	if sz <= 100000 {
		// fmt.Printf("Found %q with size %d\n", e.FullName(), e.FullSize())
		fmt.Printf("%d\n", e.FullSize())
	}
	for _, c := range e.Children {
		if c.Type == DirEntry {
			c.walkDown()
		}
	}
}

func (e *Entry) FullSize() int {
	size := 0
	for _, c := range e.Children {
		switch c.Type {
		case FileEntry:
			size += c.Size
		case DirEntry:
			size += c.FullSize()
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

func sumOfDirsUnder100000(r io.Reader) int {
	_, err := parseTree(r)
	if err != nil {
		fmt.Printf("Unexpected error while parsing tree: %v", err)
	}
	return 0
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
