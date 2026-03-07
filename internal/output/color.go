package output

func Green(s string) string {
	return "\033[32m" + s + "\033[0m"
}
