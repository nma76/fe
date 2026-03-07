package output

import "path/filepath"

func PrintFiles(files []string, c Colorizer) {
	if len(files) == 0 {
		return
	}

	colored := make([]string, len(files))
	for i, f := range files {
		col := c.ForExecutable(filepath.Ext(f))
		colored[i] = col.Sprint(f)
	}

	printColumns(colored)
}
