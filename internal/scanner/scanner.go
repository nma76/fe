package scanner

import (
	"fe/internal/detector"
	"os"
	"path/filepath"
)

type Scanner struct {
	Detector detector.ExecutableDetector
}

func (s Scanner) Scan(path string) ([]os.FileInfo, error) {
	var executables []os.FileInfo

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		info, _ := e.Info()
		full := filepath.Join(path, e.Name())

		if !info.IsDir() && s.Detector.IsExecutable(full, info) {
			executables = append(executables, info)
		}
	}
	return executables, nil
}
