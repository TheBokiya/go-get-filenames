package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filenames, err := readFilenames("pre-wedding")
	if err != nil {
		fmt.Println("Error reading filenames:", err)
		return
	}

	err = saveToFile("selected-photos.txt", toString(filenames))
	if err != nil {
		fmt.Println("Error saving to file:", err)
	}
}

func toString(filenames []string) string {
	return strings.Join(filenames, "\n")
}

func saveToFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0666)
}

func readFilenames(folderName string) ([]string, error) {
	var filenames []string
	err := filepath.WalkDir(folderName, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			filenames = append(filenames, d.Name())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return filenames, nil
}
