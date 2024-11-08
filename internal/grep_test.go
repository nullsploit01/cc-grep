package internal

import (
	"os"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	fileContent := "This is a test line\nAnother line\nLine with pattern\nPattern line again\n"

	createTempFile := func(content string) (*os.File, error) {
		tmpFile, err := os.CreateTemp("", "testfile")
		if err != nil {
			return nil, err
		}
		if _, err := tmpFile.Write([]byte(content)); err != nil {
			tmpFile.Close()
			return nil, err
		}
		if _, err := tmpFile.Seek(0, 0); err != nil {
			tmpFile.Close()
			return nil, err
		}
		return tmpFile, nil
	}

	defer func() {
		files, _ := os.ReadDir(os.TempDir())
		for _, file := range files {
			if strings.HasPrefix(file.Name(), "testfile") {
				os.Remove(os.TempDir() + "/" + file.Name())
			}
		}
	}()

	tests := []struct {
		name           string
		pattern        string
		expectedOutput []string
	}{
		{
			name:           "Non-empty pattern",
			pattern:        "pattern",
			expectedOutput: []string{"Line with pattern", "Pattern line again"},
		},
		{
			name:           "Empty pattern - returns all content",
			pattern:        "",
			expectedOutput: []string{fileContent},
		},
		{
			name:           "No matches",
			pattern:        "nomatch",
			expectedOutput: []string{},
		},
		{
			name:           "Full word match",
			pattern:        `\btest\b`,
			expectedOutput: []string{"This is a test line"},
		},
	}

	g := &Grep{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := createTempFile(fileContent)
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer tmpFile.Close()

			matches, err := g.Grep(tt.pattern, tmpFile)
			if err != nil {
				t.Fatalf("Grep() error = %v", err)
			}

			if len(matches) != len(tt.expectedOutput) {
				t.Errorf("expected %d matches, got %d", len(tt.expectedOutput), len(matches))
			}
			for i, line := range matches {
				if line != tt.expectedOutput[i] {
					t.Errorf("expected line %d to be %q, got %q", i, tt.expectedOutput[i], line)
				}
			}
		})
	}
}
