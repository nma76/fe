package output

type Theme struct {
	ExecutableColor func(string) string
}

var Classic = Theme{
	ExecutableColor: Green,
}
