package main

import (
	"Sovereign/pkg/gui"
	"Sovereign/pkg/visuals"
	"bufio"
	"fmt"
	"github.com/muesli/termenv"
	"os"
	"strings"
)

// Moved panel into main.

func main() {
	// Animation stuff
	out := termenv.NewOutput(os.Stdout)
	// Hide the cursor
	out.HideCursor()

	visuals.ClearScreen()
	visuals.Draw_Launch_Animation()
	// Launch visuals
	visuals.ClearScreen()

	// Print Sovereign ascii logo to the screen, 2 sec
	visuals.Draw_Logo()

	out.ShowCursor()

	visuals.ClearScreen()

	// "dev panel" xD
	fmt.Println("**************")
	fmt.Println("* Dev Panel: *")
	fmt.Println("**************")
	fmt.Println()

	// Testing all colors for fun
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

	fmt.Println("commands: test gui, test conncetion, quit")

	reader := bufio.NewReader(os.Stdin)

	// Create the run loop
	for {
		// Text prompt formatting
		fmt.Println("Enter command: ")
		fmt.Print("$ ")
		// Take in user input and check for error
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}
		command = strings.TrimSpace(command)

		switch command {
		// Start a connection with another user
		case "test gui":
			//TODO: Make this actually do some shit.
			fmt.Println("Testing GUI.")
			fmt.Println()
			gui := gui.NewGUI()
			gui.Start()
		case "test connection":
			//TODO: Make this actually do some shit.
			fmt.Println(" - Add net code here.")
			fmt.Println()
		// Quit application
		case "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			// Exits loop and program
			return
		default:
			fmt.Println("Invalid command.")
			fmt.Println()
		}

	}
}
