package help

import "fmt"

func PrintHelp() {
	fmt.Println("fe - find executables in a directory")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  fe [options] [directory]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println(" -h, --help              Show this help message")
	fmt.Println(" -i, --icon              Show icon for each file")
	fmt.Println(" -f, --filter            Filter files on their names")
	fmt.Println(" -v, --version           Show version information")
	fmt.Println(" -d, --debug             Enable debug output")
}
