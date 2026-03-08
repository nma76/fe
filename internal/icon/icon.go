package icon

type IconProvider interface {
	ForExecutable(ext string) string
}
