package main

import (
	"fe/internal/cli"
	"fe/internal/debug"
	"fe/internal/detector"
	"fe/internal/help"
	"fe/internal/icon"
	"fe/internal/output"
	"fe/internal/scanner"
	"fmt"
	"os"
)

// Version information, set at build time using -ldflags
var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func main() {
	// Parse command-line options
	opts := cli.Parse()

	// If the help flag is set, print the help message and exit
	if opts.Help {
		help.PrintHelp()
		return
	}

	// If the version flag is set, print version information and exit
	if opts.Version {
		fmt.Printf("fe version %s (commit: %s, built: %s)\n", Version, Commit, BuildDate)
		return
	}

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

	// Create a new platform detector
	det := detector.PlatformDetector{}

	// Create a new scanner with the platform detector and scan the specified path for executables
	scan := scanner.Scanner{Detector: det, Filter: opts.Filter}
	files, err := scan.Scan(opts.Path)
	// Handle any errors that occur during scanning
	if err != nil {
		//panic(err)
		fmt.Fprintf(os.Stderr, "fe: %v\n", err)
		os.Exit(1)
	}

	// Print the found executables with appropriate colors
	output.PrintFiles(files, colorizer, iconProvider)
}
