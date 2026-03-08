package help

import "fmt"

func PrintHelp() {
	fmt.Println("fe - find executables in a directory")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  fe [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -p, --path <directory>   Directory to scan for executables (default: current directory)")
	fmt.Println("  -h, --help              Show this help message")
	fmt.Println("      --debug             Enable debug output")
}
