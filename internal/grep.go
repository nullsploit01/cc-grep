package internal

import (
	"bufio"
	"os"
	"regexp"
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

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var matches []string
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			matches = append(matches, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func (g *Grep) RecursiveGrep() {
}
