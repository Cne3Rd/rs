package rs

import (
	"fmt"
	"os"
	"path/filepath"
)

// rs recursively traverserse a directory
// visits every file and subdirectory contained
// within a directory, each file and subdirectory in those subdirectories, and so on
// returning a slice containing the files and directories

var list []string

func init() {
	fmt.Println("walking...\nthis may take a while")
}

func Walk(path string) []string {
	f, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range f {
		fpath := filepath.Join(path, file.Name())
		if file.IsDir() {
			Walk(fpath)
		}

		list = append(list, fpath)
	}
	return list
}
