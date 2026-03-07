package cli

type Options struct {
	Path  string
	Theme string
}

func Parse() Options {
	// TODO: parse flags
	return Options{
		Path:  ".",
		Theme: "classic",
	}
}
