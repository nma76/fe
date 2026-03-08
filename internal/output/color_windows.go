//go:build windows

package output

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type PlatformColorizer struct{}

func (PlatformColorizer) ForExecutable(f os.FileInfo) *color.Color {
	ext := strings.ToLower(filepath.Ext(f.Name()))

	switch ext {
	case ".exe":
		return color.New(color.FgGreen)
	case ".com":
		return color.New(color.FgYellow)
	case ".bat", ".cmd", ".ps1":
		return color.New(color.FgCyan)
	default:
		return color.New(color.FgGreen)
	}
}
