package cli_test

import (
	"fe/internal/cli"
	"flag"
	"os"
	"reflect"
	"testing"
)

func withArgs(args []string, fn func()) {
	origArgs := os.Args
	origCommandLine := flag.CommandLine

	os.Args = args
	flag.CommandLine = flag.NewFlagSet(args[0], flag.ExitOnError)

	defer func() {
		os.Args = origArgs
		flag.CommandLine = origCommandLine
	}()

	fn()
}

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want cli.Options
	}{
		{
			name: "default values",
			args: []string{"fe"},
			want: cli.Options{
				Path:   ".",
				Help:   false,
				Debug:  false,
				Icon:   false,
				Filter: "",
			},
		},
		{
			name: "with path argument",
			args: []string{"fe", "/test/path"},
			want: cli.Options{
				Path:   "/test/path",
				Help:   false,
				Debug:  false,
				Icon:   false,
				Filter: "",
			},
		},
		{
			name: "with help argument",
			args: []string{"fe", "--help"},
			want: cli.Options{
				Path:   ".",
				Help:   true,
				Debug:  false,
				Icon:   false,
				Filter: "",
			},
		},
		{
			name: "with icon argument",
			args: []string{"fe", "--icon"},
			want: cli.Options{
				Path:   ".",
				Help:   false,
				Debug:  false,
				Icon:   true,
				Filter: "",
			},
		},
		{
			name: "with filter argument",
			args: []string{"fe", "--filter", "test"},
			want: cli.Options{
				Path:   ".",
				Help:   false,
				Debug:  false,
				Icon:   false,
				Filter: "test",
			},
		},
		{
			name: "with debug argument",
			args: []string{"fe", "--debug"},
			want: cli.Options{
				Path:   ".",
				Help:   false,
				Debug:  true,
				Icon:   false,
				Filter: "",
			},
		},
		{
			name: "with arguments and path",
			args: []string{"fe", "-i", "/test/path"},
			want: cli.Options{
				Path:   "/test/path",
				Help:   false,
				Debug:  false,
				Icon:   true,
				Filter: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withArgs(tt.args, func() {
				got := cli.Parse()

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Parse() = %v, want %v", got, tt.want)
				}
			})
		})
	}
}
