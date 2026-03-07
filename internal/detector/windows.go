//go:build windows

package detector

import (
	"os"
	"path/filepath"
	"strings"
)

type PlatformDetector struct{}

// IsExecutable checks if the file at the given path is executable by checking its file extension against common
// executable extensions on Windows.
func (PlatformDetector) IsExecutable(path string, info os.FileInfo) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".exe", ".bat", ".cmd", ".com":
		return true
	default:
		return false
	}
}
