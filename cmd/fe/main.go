package main

import (
	"fe/internal/cli"
	"fe/internal/detector"
	"fe/internal/output"
	"fe/internal/scanner"
	"fmt"
)

func main() {
	opts := cli.Parse()

	det := detector.PlatformDetector{}

	scan := scanner.Scanner{Detector: det}
	files, err := scan.Scan(opts.Path)
	if err != nil {
		panic(err)
	}

	theme := output.Classic

	for _, f := range files {
		fmt.Println(theme.ExecutableColor(f))
	}
}
