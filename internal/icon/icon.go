package icon

import "os"

type IconProvider interface {
	ForExecutable(f os.FileInfo) string
}
