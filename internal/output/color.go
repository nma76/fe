package output

import "github.com/fatih/color"

type Colorizer interface {
	ForExecutable(ext string) *color.Color
	ForMessages() *color.Color
}
