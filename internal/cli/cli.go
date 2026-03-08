package cli

import "flag"

type Options struct {
	Path string
}

func Parse() Options {
	var opts Options

	flag.StringVar(&opts.Path, "path", ".", "Directory to scan for executables")
	flag.StringVar(&opts.Path, "p", ".", "shorthand fpr --path")
	flag.Parse()

	return opts
}
