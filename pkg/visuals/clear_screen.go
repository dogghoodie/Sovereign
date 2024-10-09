// This might be stupid? but I like having it.

package visuals

import "fmt"

func ClearScreen() {
	// Print ANSI escape sequence
	fmt.Print("\033[H\033[2J")
}
