package cli

import (
	"flag"
)

type Options struct {
	Path  string
	Help  bool
	Debug bool
}

func Parse() Options {
	var opts Options

	// path flag for the directory to scan for executables, defaulting to the current directory
	flag.StringVar(&opts.Path, "path", ".", "Directory to scan for executables")
	flag.StringVar(&opts.Path, "p", ".", "shorthand for --path")

	// help flag to show the help message
	flag.BoolVar(&opts.Help, "help", false, "Show help message")
	flag.BoolVar(&opts.Help, "h", false, "shorthand for --help")

	// debug flag to enable debug output
	flag.BoolVar(&opts.Debug, "debug", false, "Enable debug output")

	flag.Parse()

	return opts
}
