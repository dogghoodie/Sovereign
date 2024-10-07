package main

import (
	"Sovereign/pkg/panel"
	"Sovereign/pkg/visuals"
	"fmt"
	"time"
)

func main() {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// Print Sovereign ascii logo to the screen
	visuals.Draw()
	time.Sleep(2000 * time.Millisecond)

	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// Initialize client interface potentially rename to gocui terms
	panel.CreatePanel()
}
