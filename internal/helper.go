package internal

import (
	"os"
	"path/filepath"
)

func ReadFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

func ReadFilesFromGlob(pattern string) ([]*os.File, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	return ReadFiles(files)
}

func ReadFiles(fileNames []string) ([]*os.File, error) {
	var files []*os.File
	for _, fileName := range fileNames {
		file, err := ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}
