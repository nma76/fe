package output

import (
	"fe/internal/icon"
	"os"
)

func PrintFiles(files []os.FileInfo, c Colorizer, ip icon.IconProvider) {
	// If no files were found, return early
	if len(files) == 0 {
		return
	}

	// Create a slice to hold the colored file names with icons
	colored := make([]string, len(files))

	// Iterate over the found files and apply color and icons based on their extensions
	for i, f := range files {
		// Get the appropriate icon for the file based on its extension
		ic := ip.ForExecutable(f)

		// Get the appropriate color for the file based on its extension and combine it with the icon
		col := c.ForExecutable(f)
		coloredName := col.Sprint(f.Name())

		colored[i] = ic + coloredName
	}

	printColumns(colored)
}
