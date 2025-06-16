package main

import (
	"os"
	"testing"
)

func cleanup(paths ...string) {
	for _, path := range paths {
		os.RemoveAll(path)
	}
}

func TestCreateDirs(t *testing.T) {
	testDirs := []string{"testdir1", "testdir2/nested"}

	createDirs(testDirs)

	for _, dir := range testDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Directory %s was not created", dir)
		}
	}

	cleanup("testdir1", "testdir2")
}

func TestWriteFile(t *testing.T) {
	path := "testfile.txt"
	content := "Hello, World!"

	writeFile(path, content)

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", path, err)
	}

	if string(data) != content {
		t.Errorf("Expected content %q, got %q", content, string(data))
	}

	cleanup(path)
}

