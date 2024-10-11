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

//TODO: 1. Fix GUI coloring on selection.
//		2. Add functionality to messagebox.
//		3. Figure out globe animation.
//		4. Start working on connectivity.
//		5. Integrate chat encryption.
//		6. Package as stand-alone.
//		7. Package as nvim plugin.

// Entry function
func main() {
	// Animation stuff.
	// Initialize output handler (formatting).
	outputHandler := termenv.NewOutput(os.Stdout)
	// Hide the cursor.
	outputHandler.HideCursor()
	// Clear the terminal.
	visuals.ClearScreen()
	// Start the launch animation.
	visuals.Draw_Launch_Animation()
	// Clear the terminal.
	visuals.ClearScreen()
	// Start the logo animation.
	visuals.Draw_Logo()
	// Bring back terminal cursor.
	outputHandler.ShowCursor()
	// Clear the terminal.
	visuals.ClearScreen()

	// Draw the "dev panel". xD
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

	// Initialize bufio reader for input
	reader := bufio.NewReader(os.Stdin)

	// Create the run loop
	for {
		// Text prompt formatting.
		fmt.Println("Enter command: ")
		fmt.Print("$ ")

		// Initialize command as next string from user.
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			// Prompt user again if error.
			continue
		}

		// Trim extra white space off of input.
		command = strings.TrimSpace(command)

		// Handle the different inputs.
		switch command {
		// Test gui functionality.
		case "test gui":
			//TODO: Make this actually do some shit.
			fmt.Println("Testing GUI.")
			fmt.Println()
			gui := gui.NewGUI()
			gui.Start()
		// Test user connectivity.
		case "test connection":
			//TODO: Make this actually do some shit.
			fmt.Println(" - Add net code here.")
			fmt.Println()
		// Quit application.
		case "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			// Exits loop and program.
			return
		default:
			// There's a fuckin list.
			fmt.Println("Invalid command.")
			fmt.Println()
		}

	}
}
