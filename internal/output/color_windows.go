//go:build windows

package output

import (
	"strings"

	"github.com/fatih/color"
)

type PlatformColorizer struct{}

func (PlatformColorizer) ForExecutable(ext string) *color.Color {
	ext = strings.ToLower(ext)

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
