package internal

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func createTempFile(t *testing.T, content string) *os.File {
	t.Helper()
	file, err := ioutil.TempFile("", "grep_test")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := file.WriteString(content); err != nil {
		file.Close()
		t.Fatalf("failed to write to temp file: %v", err)
	}
	file.Seek(0, 0)
	return file
}

func TestGrep_GetMatches(t *testing.T) {
	grep := NewGrep()
	content := "Hello World\nhello Go\nHELLO gophers\n"
	pattern := "hello"

	tests := []struct {
		name            string
		caseInsensetive bool
		invert          bool
		expectedMatches []string
	}{
		{
			name:            "Case-sensitive match",
			caseInsensetive: false,
			invert:          false,
			expectedMatches: []string{"hello Go"},
		},
		{
			name:            "Case-insensitive match",
			caseInsensetive: true,
			invert:          false,
			expectedMatches: []string{"Hello World", "hello Go", "HELLO gophers"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches, err := grep.GetMatches(content, pattern, tt.caseInsensetive, tt.invert)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(matches) != len(tt.expectedMatches) {

				t.Fatalf("expected %d matches, got %d", len(tt.expectedMatches), len(matches))
			}

			for i, match := range matches {
				if match != tt.expectedMatches[i] {
					t.Errorf("expected match %q, got %q", tt.expectedMatches[i], match)
				}
			}
		})
	}
}

func TestGrep_Grep(t *testing.T) {
	grep := NewGrep()
	content := "Hello World\nhello Go\nHELLO gophers\n"
	file := createTempFile(t, content)
	defer os.Remove(file.Name())
	defer file.Close()

	matches, err := grep.Grep("hello", file.Name(), true, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"Hello World", "hello Go", "HELLO gophers"}
	if len(matches) != len(expected) {
		t.Fatalf("expected %d matches, got %d", len(expected), len(matches))
	}

	for i, match := range matches {
		if match != expected[i] {
			t.Errorf("expected match %q, got %q", expected[i], match)
		}
	}
}

func TestGrep_RecursiveGrep(t *testing.T) {
	grep := NewGrep()

	tempDir, err := ioutil.TempDir("", "recursive_grep_test")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	file1 := createTempFileInDir(t, tempDir, "file1.txt", "Go is fun\nGo gophers\n")
	file2 := createTempFileInDir(t, tempDir, "file2.txt", "Gophers are great\nGo language\n")
	log.Printf("Created test files: %s, %s", file1.Name(), file2.Name())

	filesPattern := filepath.Join(tempDir, "*")
	log.Printf("Using filesPattern: %s", filesPattern)

	results, err := grep.RecursiveGrep(filesPattern, "Go", true, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	expectedMatches := [][]string{
		{"Go is fun", "Go gophers"},
		{"Gophers are great", "Go language"},
	}

	for i, result := range results {
		if len(result.Matches) != len(expectedMatches[i]) {
			t.Fatalf("expected %d matches for file %s, got %d", len(expectedMatches[i]), result.FileName, len(result.Matches))
		}
		for j, match := range result.Matches {
			if match != expectedMatches[i][j] {
				t.Errorf("expected match %q, got %q", expectedMatches[i][j], match)
			}
		}
	}
}

func createTempFileInDir(t *testing.T, dir, name, content string) *os.File {
	t.Helper()
	filePath := filepath.Join(dir, name)
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := file.WriteString(content); err != nil {
		file.Close()
		t.Fatalf("failed to write to temp file: %v", err)
	}
	file.Seek(0, 0)
	return file
}
