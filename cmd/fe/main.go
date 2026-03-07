package main

import (
	"fe/internal/cli"
	"fe/internal/detector"
	"fe/internal/output"
	"fe/internal/scanner"
	"path/filepath"
)

func main() {
	// Handles colored output for messages and executables
	colorizer := output.PlatformColorizer{}

	// Print a welcome message
	c := colorizer.ForMessages()
	c.Println(("This is fe... 2026 edition!"))

	// Parse command-line options
	opts := cli.Parse()

	// Create a new platform detector
	det := detector.PlatformDetector{}

	// Create a new scanner with the platform detector and scan the specified path for executables
	scan := scanner.Scanner{Detector: det}
	files, err := scan.Scan(opts.Path)

	// Handle any errors that occur during scanning
	if err != nil {
		panic(err)
	}

	// Print the found executables with appropriate colors
	for _, f := range files {
		c := colorizer.ForExecutable(filepath.Ext(f))
		c.Println(f)
	}
}
