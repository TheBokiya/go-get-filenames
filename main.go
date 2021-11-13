package main

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	var filenames []string

	folder := "pre-wedding"
	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		filenames = append(filenames, d.Name())
		return nil
	})

	if err != nil {
		panic(err)
	}

	saveToFile("selected-photos.txt", toString(filenames))
}

func toString(filenames []string) string {
	return strings.Join([]string(filenames), "\n")
}

func saveToFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)
}
