package main

import (
	"fe/internal/cli"
	"fe/internal/debug"
	"fe/internal/detector"
	"fe/internal/output"
	"fe/internal/scanner"
)

func main() {
	// Handles colored output for messages and executables
	colorizer := output.PlatformColorizer{}

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

	debug.Println("Path:", opts.Path)

	// Print the found executables with appropriate colors
	output.PrintFiles(files, colorizer)
}
