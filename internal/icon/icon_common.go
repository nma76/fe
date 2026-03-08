package icon

type DefaultIconProvider struct{}

func (DefaultIconProvider) ForExecutable(ext string) string {
	// Executable generic icon
	return "\U000f107b" + " "
}
