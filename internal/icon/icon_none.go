package icon

type NoIconProvider struct{}

func (NoIconProvider) ForExecutable(ext string) string {
	return ""
}
