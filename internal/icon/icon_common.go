package icon

import (
	"os"
	"path/filepath"
	"strings"
)

type DefaultIconProvider struct{}

func (DefaultIconProvider) ForExecutable(f os.FileInfo) string {
	// Check for shebang first
	if sb, _ := readShebang(f); sb != "" {
		if strings.Contains(sb, "python") {
			return "\U000f0818 " // Python
		}
		if strings.Contains(sb, "bash") || strings.Contains(sb, "sh") {
			return "\U000f0bc1 " // Shell
		}
		if strings.Contains(sb, "node") {
			return "\U000e7188 " // Node
		}
		if strings.Contains(sb, "perl") {
			return "\U000f03f8 " // Perl
		}
		// fler senare…
	}

	// If the file has a shebang, we can determine the type of executable
	ext := strings.ToLower(filepath.Ext(f.Name()))
	switch ext {
	case ".app", ".exe", ".com":
		return "\U000f107b" + " "
	case ".sh", ".bat", ".cmd", ".ps1":
		return "\U000f0bc1" + " "
	}

	// Executable generic icon
	return "\U000f107b" + " "
}
