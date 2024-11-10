package internal

import (
	"os"
	"path/filepath"
)

func ReadFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

func ReadFilesFromGlob(pattern string) ([]*os.File, error) {
	baseDir := filepath.Dir(pattern)
	filePattern := filepath.Base(pattern)
	var fileNames []string

	err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			matches, err := filepath.Match(filePattern, filepath.Base(path))
			if err != nil {
				return err
			}

			if matches {
				fileNames = append(fileNames, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return ReadFiles(fileNames)
}

func ReadFiles(fileNames []string) ([]*os.File, error) {
	var files []*os.File
	for _, fileName := range fileNames {
		file, err := ReadFile(fileName)
		if err != nil {
			for _, f := range files {
				f.Close()
			}

			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}
