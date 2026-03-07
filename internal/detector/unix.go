//go:build darwin || linux || freebsd || openbsd || netbsd

package detector

import "os"

type PlatformDetector struct{}

func (PlatformDetector) IsExecutable(path string, info os.FileInfo) bool {
	return info.Mode()&0111 != 0
}
