package scanner

import (
	"fe/internal/detector"
	"os"
	"path/filepath"
	"strings"
)

type Scanner struct {
	Detector detector.ExecutableDetector
	Filter   string
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

		if info.IsDir() {
			continue
		}

		if s.Filter != "" && !strings.Contains(strings.ToLower(info.Name()), strings.ToLower(s.Filter)) {
			continue
		}

		if s.Detector.IsExecutable(full, info) {
			executables = append(executables, info)
		}
	}
	return executables, nil
}
