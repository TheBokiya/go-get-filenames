package main

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	filenames := readFilenames("pre-wedding")

	saveToFile("selected-photos.txt", toString(filenames))
}

func toString(filenames []string) string {
	return strings.Join([]string(filenames), "\n")
}

func saveToFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)
}

func readFilenames(folderName string) []string {
	var filenames []string
	err := filepath.WalkDir(folderName, func(path string, d fs.DirEntry, err error) error {
		filenames = append(filenames, d.Name())
		return nil
	})

	if err != nil {
		panic(err)
	}
	return filenames
}
