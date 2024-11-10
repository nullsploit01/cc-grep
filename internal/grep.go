package internal

import (
	"io"
	"regexp"
	"strings"
)

type Grep struct {
}

type RecursiveGrepResult struct {
	FileName string
	Matches  []string
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

	return g.GetMatches(string(content), pattern, caseInsensetive)
}

func (g *Grep) RecursiveGrep(filePattern string, pattern string, caseInsensetive bool) ([]RecursiveGrepResult, error) {
	files, err := ReadFilesFromGlob(filePattern)
	if err != nil {
		return nil, err
	}

	var result []RecursiveGrepResult
	for _, file := range files {
		matches, err := g.Grep(pattern, file.Name(), caseInsensetive)
		if err != nil {
			return nil, err
		}

		result = append(result, RecursiveGrepResult{
			FileName: file.Name(),
			Matches:  matches,
		})
	}

	return result, nil
}

func (g *Grep) GetMatches(content string, pattern string, caseInsensetive bool) ([]string, error) {
	if pattern == "" {
		return []string{content}, nil
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
