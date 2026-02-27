package verbose

import "fmt"

var Enabled bool

func Verbosef(format string, a ...any) {
	if Enabled {
		fmt.Printf("[v] "+format+"\n", a...)
	}
}
