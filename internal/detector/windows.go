//go:build windows

package detector

import (
	"os"
	"path/filepath"
	"strings"
)

type PlatformDetector struct{}

func (PlatformDetector) IsExecutable(path string, info os.FileInfo) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".exe", ".bat", ".cmd", ".com":
		return true
	default:
		return false
	}
}
