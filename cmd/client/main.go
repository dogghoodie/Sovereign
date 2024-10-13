package main

import (
	"Sovereign/pkg/gui"
	"Sovereign/pkg/visuals"
	"bufio"
	"fmt"
	"github.com/muesli/termenv"
	"os"
	"strings"
	"Sovereign/pkg/encryption"
	"strconv"
)

//TODO: 1. Fix GUI coloring on selection.
//		2. Add functionality to messagebox.
//		3. Figure out globe animation.
//		4. Polish the GUI.
//		5. Start working on connectivity.
//		6. Integrate chat encryption.
//		7. Package as stand-alone.
//		8. Package as nvim plugin.


var seed int64 = 12  // Default global seed
var cipher = *encryption.SetSeed(seed)
var encryptedMessage = "default_msg"
var testMes = "abc"

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
	fmt.Println("commands: test gui, test conncetion, quit, set seed, encrypt, decrypt")

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
		switch {
		// Test gui functionality.
		case command == "test gui":
			//TODO: Make this actually do some shit.
			fmt.Println("Testing GUI.")
			fmt.Println()
			gui := gui.NewGUI()
			gui.Start()
		// Test user connectivity.
		case command == "test connection":
			//TODO: Make this actually do some shit.
			fmt.Println(" - Add net code here.")
			fmt.Println()
		// Quit application.
		case command == "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			// Exits loop and program.
			return
		// Encrypts message to chinese
		case len(command) >= 6 && command[:7] == "encrypt":
			message := strings.TrimSpace(command[7:])
			// call method from cipher
			encryptedMessage = cipher.EncryptMessage(message)
			fmt.Println("Encrypted message:", encryptedMessage)
			testMes = "def"
		// Decrypts message from chinese, could add handling for no encrypted message,
		// but the encrypt and decrypt commands wont be manually called in final version anyway
		case len(command) >= 6 && command[:7] == "decrypt":
			decryptedMessage := cipher.DecryptMessage(encryptedMessage)
			fmt.Println("Decrypted message:", decryptedMessage)
			fmt.Println("testMes:", testMes)
		// Sets seed for chinese encryption
		case len(command) >= 7 && command[:8] == "set seed":
			// turn seed string into int64
			seed, err := strconv.ParseInt(strings.TrimSpace(command[8:]), 10, 64)
			if err != nil { panic(err) }
			// set new cipher based on new seed
			cipher = *encryption.SetSeed(seed)
			fmt.Println("Seed set")

		default:
			// There's a fuckin list.
			fmt.Println("Invalid command.")
			fmt.Println()
		}

	}
}
