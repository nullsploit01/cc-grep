package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func createTempDirWithFiles(t *testing.T, files map[string]string) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}

	for name, content := range files {
		filePath := filepath.Join(dir, name)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("failed to create temp file %s: %v", name, err)
		}
	}

	return dir
}

func TestReadFile(t *testing.T) {
	file := createTempFile(t, "Hello, World!")
	defer os.Remove(file.Name()) // Clean up
	defer file.Close()

	readFile, err := ReadFile(file.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer readFile.Close()

	// Check that the file is open
	stat, err := readFile.Stat()
	if err != nil {
		t.Fatalf("failed to stat file: %v", err)
	}
	if stat.Size() == 0 {
		t.Errorf("expected non-zero file size, got %d", stat.Size())
	}
}

func TestReadFilesFromGlob(t *testing.T) {
	dir := createTempDirWithFiles(t, map[string]string{
		"file1.txt": "Content of file 1",
		"file2.log": "Content of file 2",
		"file3.txt": "Content of file 3",
	})
	defer os.RemoveAll(dir) // Clean up

	pattern := filepath.Join(dir, "*.txt")
	files, err := ReadFilesFromGlob(pattern)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}

	expectedFiles := map[string]bool{
		"file1.txt": true,
		"file3.txt": true,
	}

	for _, file := range files {
		_, fileName := filepath.Split(file.Name())
		if !expectedFiles[fileName] {
			t.Errorf("unexpected file matched: %s", file.Name())
		}
	}
}

func TestReadFiles(t *testing.T) {
	file1 := createTempFile(t, "Content of file 1")
	file2 := createTempFile(t, "Content of file 2")
	defer os.Remove(file1.Name())
	defer os.Remove(file2.Name())

	// Test opening multiple files
	fileNames := []string{file1.Name(), file2.Name()}
	files, err := ReadFiles(fileNames)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}

	expectedFileNames := map[string]bool{
		file1.Name(): true,
		file2.Name(): true,
	}

	for _, file := range files {
		if !expectedFileNames[file.Name()] {
			t.Errorf("unexpected file opened: %s", file.Name())
		}
	}
}
