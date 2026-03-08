package output

import (
	"os"

	"github.com/fatih/color"
)

type Colorizer interface {
	ForExecutable(f os.FileInfo) *color.Color
	ForMessages() *color.Color
}
