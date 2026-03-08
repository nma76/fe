package icon

import "os"

type NoIconProvider struct{}

func (NoIconProvider) ForExecutable(f os.FileInfo) string {
	return ""
}
