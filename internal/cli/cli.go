package cli

import (
	"flag"
)

type Options struct {
	Path    string
	Help    bool
	Debug   bool
	Icon    bool
	Filter  string
	Version bool
}

func Parse() Options {
	var opts Options

	// help flag to show the help message
	flag.BoolVar(&opts.Help, "help", false, "Show help message")
	flag.BoolVar(&opts.Help, "h", false, "shorthand for --help")

	// icon flag to enable icon output
	flag.BoolVar(&opts.Icon, "icon", false, "Enable icon output")
	flag.BoolVar(&opts.Icon, "i", false, "shorthand for --icon")

	// filter on filename
	flag.StringVar(&opts.Filter, "filter", "", "Filter files by name")
	flag.StringVar(&opts.Filter, "f", "", "Shorthand for --filter")

	// version flag to show version information
	flag.BoolVar(&opts.Version, "version", false, "Show version information")
	flag.BoolVar(&opts.Version, "v", false, "Shorthand for --version")

	// debug flag to enable debug output
	flag.BoolVar(&opts.Debug, "debug", false, "Enable debug output")
	flag.BoolVar(&opts.Debug, "d", false, "shorthand for --debug")

	// Parse the command-line flags
	flag.Parse()

	// If user provided a positional argument, treat it as the path
	if flag.NArg() > 0 {
		opts.Path = flag.Arg(0)
	} else {
		opts.Path = "."
	}

	return opts
}
