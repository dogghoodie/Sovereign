// This might be stupid? but I like having it.

package visuals

import "fmt"

// Function cleanrs terminal screen. Yes I'm commenting this, too.
func ClearScreen() {
	// Print ANSI escape sequence
	fmt.Print("\033[H\033[2J")
}
