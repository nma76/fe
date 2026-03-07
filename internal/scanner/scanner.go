package scanner

import (
	"fe/internal/detector"
	"os"
	"path/filepath"
)

type Scanner struct {
	Detector detector.ExecutableDetector
}

func (s Scanner) Scan(path string) ([]string, error) {
	var executables []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		info, _ := e.Info()
		full := filepath.Join(path, e.Name())

		if !info.IsDir() && s.Detector.IsExecutable(full, info) {
			executables = append(executables, full)
		}
	}
	return executables, nil
}
