package internal

import (
	"io"
	"os"
	"regexp"
	"strings"
)

type Grep struct {
}

func NewGrep() *Grep {
	return &Grep{}
}

func (g *Grep) Grep(pattern string, file *os.File) ([]string, error) {
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if pattern == "" {
		return []string{string(content)}, nil
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var matches []string
	for _, line := range lines {
		if re.MatchString(line) {
			matches = append(matches, line)
		}
	}

	return matches, nil
}

func (g *Grep) RecursiveGrep() {
}
