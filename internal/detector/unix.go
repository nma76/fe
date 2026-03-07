//go:build darwin || linux || freebsd || openbsd || netbsd

package detector

import "os"

type PlatformDetector struct{}

// IsExecutable checks if the file at the given path is executable by checking its permission bits.
func (PlatformDetector) IsExecutable(path string, info os.FileInfo) bool {
	return info.Mode()&0111 != 0
}
