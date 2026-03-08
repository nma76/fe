package icon

import "os"

type DefaultIconProvider struct{}

func (DefaultIconProvider) ForExecutable(f os.FileInfo) string {
	// Executable generic icon
	return "\U000f107b" + " "
}
