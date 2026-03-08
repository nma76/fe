package cli

import (
	"fe/internal/debug"
	"flag"
)

type Options struct {
	Path string
}

func Parse() Options {
	var opts Options

	// path flag for the directory to scan for executables, defaulting to the current directory
	flag.StringVar(&opts.Path, "path", ".", "Directory to scan for executables")
	flag.StringVar(&opts.Path, "p", ".", "shorthand fpr --path")

	// debug flag to enable debug output
	flag.BoolVar(&debug.Enabled, "debug", false, "Enable debug output")

	flag.Parse()

	return opts
}
