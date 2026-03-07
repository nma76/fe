package output

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

var ansi = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

func visibleLen(s string) int {
	clean := ansi.ReplaceAllString(s, "")
	return utf8.RuneCountInString(clean)
}

func printColumns(items []string) {
	if len(items) == 0 {
		return
	}

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width < 20 {
		width = 80 // fallback
	}

	// hitta längsta synliga längden
	max := 0
	for _, it := range items {
		l := visibleLen(it)
		if l > max {
			max = l
		}
	}

	padding := 2
	colWidth := max + padding

	if colWidth < 4 {
		colWidth = 4
	}

	cols := width / colWidth
	if cols < 1 {
		cols = 1
	}

	for i, it := range items {
		// skriv ut färgad text
		fmt.Print(it)

		// räkna synlig längd
		pad := colWidth - visibleLen(it)
		if pad < 1 {
			pad = 1
		}

		// lägg till mellanslag
		fmt.Print(strings.Repeat(" ", pad))

		if (i+1)%cols == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

}
