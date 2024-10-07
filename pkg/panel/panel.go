package panel

import (
	"Sovereign/pkg/visuals"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function creates our main client panel
func CreatePanel() {
	fmt.Println(visuals.Colors.RED + "Sovereign Panel." + visuals.Colors.ANSI_RESET)

	// Create the run loop
	for {
		// Initialize reader
		reader := bufio.NewReader(os.Stdin)

		// Initialize command reader as next string
		command, err := reader.ReadString('\n')
		// Clean up any white space after command
		command = strings.ReplaceAll(command, "\n", "")

		switch command {
		// Just a test
		case "ping":
			fmt.Println("pong")
			fmt.Println()
		// Start a connection with another user
		case "connect":
			fmt.Println("Add connections menu here.")
			fmt.Println()
		// Quit application
		case "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			// Returns out of loop
			return
		}

		// If there is an error, print
		if err != nil {
			fmt.Print("There was an error:", err)
			continue
		}

	}
}
