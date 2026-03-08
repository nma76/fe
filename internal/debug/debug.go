package debug

import "fmt"

var Enabled bool

func Println(v ...any) {
	if Enabled {
		fmt.Println(v...)
	}
}
