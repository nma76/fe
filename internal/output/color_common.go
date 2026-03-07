package output

import "github.com/fatih/color"

func (PlatformColorizer) ForMessages() *color.Color {
	return color.New(color.Reset)
}
