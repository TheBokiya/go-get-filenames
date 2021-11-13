package main

import "testing"

func TestReadFiles(t *testing.T) {
	f := readFilenames("pre-wedding")

	if len(f) != 61 {
		t.Errorf("Expected number of photos of 60, but got %v", len(f))
	}
}
