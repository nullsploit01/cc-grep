package internal

import (
	"io"
	"regexp"
	"strings"
)

type Grep struct {
}

func NewGrep() *Grep {
	return &Grep{}
}

func (g *Grep) Grep(pattern string, fileName string, caseInsensetive bool) ([]string, error) {
	file, err := ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if pattern == "" {
		return []string{string(content)}, nil
	}

	if caseInsensetive {
		pattern = "(?i)" + pattern
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
