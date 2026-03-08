package icon

type DefaultIconProvider struct{}

func (DefaultIconProvider) ForExecutable(ext string) string {
	return "\U000f107b"
}
