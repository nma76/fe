package detector

import "os"

type ExecutableDetector interface {
	IsExecutable(path string, info os.FileInfo) bool
}
