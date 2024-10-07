package panel

import (
	"Sovereign/pkg/visuals"
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TODO: Fuck this and get gocui working

// Function creates our main client panel
func CreatePanel() {
	// Temporary design just for now
	fmt.Println("***********************")
	fmt.Println("* Client Panel Switch *")
	fmt.Println("***********************")
	fmt.Println()
	fmt.Println(
		visuals.Colors.CYAN+"x",
		visuals.Colors.GREEN+"x",
		visuals.Colors.ORANGE+"x",
		visuals.Colors.PINK+"x",
		visuals.Colors.PURPLE+"x",
		visuals.Colors.RED+"x",
		visuals.Colors.YELLOW+"x",
		visuals.Colors.BLACK+"x",
		visuals.Colors.ANSI_RESET,
	)
	fmt.Println()
	fmt.Println("commands: ping, connect, quit")

	// Create the run loop
	for {
		// Text prompt formatting
		fmt.Print("$ ")
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
			//TODO: Make this actually do some shit.
			fmt.Println("Add connections menu here.")
			fmt.Println()
		// Quit application
		case "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			// Returns out of loop
			return
		default:
			fmt.Println("Invalid command.")
			fmt.Println()
		}

		// If there is an error, print
		if err != nil {
			fmt.Print("There was an error:", err)
			continue
		}

	}
}
