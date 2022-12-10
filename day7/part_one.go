package day7

import (
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

// func sumOfDirsUnder100000(r io.Reader) int {
// 	s := bufio.NewScanner(r)

// 	root := Entry{Type: DirEntry}
// 	parentDir := &root
// 	pwd := "/"
// 	for s.Scan() {
// 		if strings.HasPrefix("dir ", s.Text()) {
// 			spl := strings.Split(s.Text(), " ")
// 			entry := Entry{Name: spl[1], Type: DirEntry, Parent: parentDir}
// 			parentDir.Children = append(parentDir.Children, &entry)
// 		}

// 		if strings.HasPrefix("cd ", s.Text()) {
// 			spl := strings.Split(s.Text(), " ")
// 			pwd = filepath.Clean(pwd + "/" + spl[1])
// 			parentDir, _ = findDir(&root, pwd)
// 		}

// 	}
// 	return 0
// }

func findDir(root *Entry, path string) (*Entry, bool) {
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
