//go:build darwin || linux || freebsd || openbsd || netbsd

package output

import (
	"os"

	"github.com/fatih/color"
)

type PlatformColorizer struct{}

func (PlatformColorizer) ForExecutable(f os.FileInfo) *color.Color {
	return color.New(color.FgGreen)
}
