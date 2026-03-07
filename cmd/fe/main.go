package main

import (
	"fe/internal/cli"
	"fe/internal/detector"
	"fe/internal/output"
	"fe/internal/scanner"
	"fmt"
	"path/filepath"
)

func main() {
	colorizer := output.PlatformColorizer{}

	fmt.Println(("This is fe..."))
	opts := cli.Parse()

	det := detector.PlatformDetector{}

	scan := scanner.Scanner{Detector: det}
	files, err := scan.Scan(opts.Path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		c := colorizer.ForExecutable(filepath.Ext(f))
		c.Println(f)
	}
}
