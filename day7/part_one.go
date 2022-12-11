package day7

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func sumOfDirsUnder100000(r io.Reader) int {
	parseTree(r)
	return 0
}

func parseTree(r io.Reader) (root *Entry) {
	s := bufio.NewScanner(r)

	root = &Entry{Type: DirEntry}
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
				log.Fatalf("Expected a number, got none: %v", err)
			}
			entry := Entry{Name: spl[1], Type: FileEntry, Size: size, Parent: parentDir}
			parentDir.Children = append(parentDir.Children, &entry)
		}

		if strings.HasPrefix(s.Text(), "$ cd ") {
			spl := strings.Split(s.Text(), " ")
			pwd = filepath.Clean(pwd + "/" + spl[2])
			var ok bool
			parentDir, ok = findEntry(root, pwd)
			if !ok {
				fmt.Printf("In this program, a directory should always exist: %q does not exist\n", pwd)
				os.Exit(1)
			}
		}

	}
	return root
}

func findEntry(root *Entry, path string) (*Entry, bool) {
	if path == "/" {
		return root, true
	}

	spl := strings.Split(path, "/")[1:] // Ignore leading slash

	segments := len(spl)

	current := root
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
