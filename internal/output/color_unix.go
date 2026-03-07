//go:build darwin || linux || freebsd || openbsd || netbsd

package output

import "github.com/fatih/color"

type PlatformColorizer struct{}

func (PlatformColorizer) ForExecutable(ext string) *color.Color {
	return color.New(color.FgGreen)
}
