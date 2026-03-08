package icon

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func readShebang(fi os.FileInfo) (string, error) {
	f, err := os.Open(fi.Name())
	if err != nil {
		return "", err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	if strings.HasPrefix(line, "#!") {
		return strings.TrimSpace(line[2:]), nil
	}

	return "", nil
}
