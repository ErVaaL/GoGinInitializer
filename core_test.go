package main

import (
	"os"
	"testing"
)

func cleanup(t *testing.T, paths ...string) {
	for _, path := range paths {
		if err := os.RemoveAll(path); err != nil {
			t.Errorf("Failed to clean up %s: %v", path, err)
		}
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

	cleanup(t, "testdir1", "testdir2")
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

	cleanup(t, path)
}
