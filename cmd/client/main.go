package main

import (
	"Sovereign/pkg/panel"
	"Sovereign/pkg/visuals"
	"time"
)

func main() {
	// Clear screen, check class for the voodoo printed
	visuals.ClearScreen()
	// Print Sovereign ascii logo to the screen, 2 sec
	visuals.Draw()
	time.Sleep(2000 * time.Millisecond)
	visuals.ClearScreen()

	// Initialize client interface
	panel.CreatePanel()

	//TODO: Go over plans written in each file.
	//		Discuss.

}
