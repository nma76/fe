package main

import (
	"fe/internal/cli"
	"fe/internal/debug"
	"fe/internal/detector"
	"fe/internal/help"
	"fe/internal/icon"
	"fe/internal/output"
	"fe/internal/scanner"
)

func main() {
	// Parse command-line options
	opts := cli.Parse()

	// Handles colored output for messages and executables
	colorizer := output.PlatformColorizer{}

	// Provides icons for executables based on their file extension
	var iconProvider icon.IconProvider
	if opts.Icon {
		iconProvider = icon.DefaultIconProvider{}
	} else {
		iconProvider = icon.NoIconProvider{}
	}

	// Enable debug output if the debug flag is set
	if opts.Debug {
		debug.Enabled = true
		debug.Println("Debug: enabled")
	}

	// If the help flag is set, print the help message and exit
	if opts.Help {
		help.PrintHelp()
		return
	}

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
	output.PrintFiles(files, colorizer, iconProvider)
}
