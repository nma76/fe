package detector

import "os"

// ExecutableDetector is an interface that defines a method to determine if a given file path and its associated file
// information correspond to an executable file. Implementations of this interface can use various criteria, such as file
// permissions, file extensions, or other heuristics, to make this determination.
type ExecutableDetector interface {
	IsExecutable(path string, info os.FileInfo) bool
}
