package output_test

import (
	"fe/internal/output"
	"testing"
)

func TestVisibleLen(t *testing.T) {
	get := output.VisibleLen("hello")
	want := 5

	if get != want {
		t.Errorf("VisibleLen %d, want %d", get, want)
	}
}

func TestVisibleLenWithANSI(t *testing.T) {
	got := output.VisibleLen("\033[31mgello\033[0m")
	want := 5

	if got != want {
		t.Errorf("VisibleLen() with ANSI = %d, want %d", got, want)
	}
}
