package main

import (
	"Sovereign/pkg/encryption"
	"Sovereign/pkg/gui"
	"Sovereign/pkg/visuals"
	"bufio"
	"fmt"
	"github.com/muesli/termenv"
	"os"
	"strconv"
	"strings"
)

// Global variables.
var seed int64 = 12 // default
var cipher = *encryption.SetSeed(seed)
var encryptedMessage = "default_msg"
var testMes = "abc"

// Entry function
func main() {
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
	fmt.Println("commands: test animation, test gui, test conncetion, quit, set seed, encrypt, decrypt")

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
		// Test the animation scene
		// vvv OHHH THE HORRORRRR !!
		case "test animation":
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
			fmt.Println("commands: test animation, test gui, test conncetion, quit, set seed, encrypt, decrypt")

		// Test gui functionality.
		case "test gui":
			fmt.Println("Testing GUI.")
			fmt.Println()
			gui.Start()

		// Test user connectivity.
		case "test connection":
			//TODO: Make this actually do some shit.
			fmt.Println(" - Add net code here.")
			fmt.Println()

		// Encrypts message to chinese
		case "encrypt":
			fmt.Printf("> ")
			// enter message
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}
			// call method from cipher
			encryptedMessage = cipher.EncryptMessage(message)
			fmt.Println("Encrypted message:\n< ", encryptedMessage)

		// Decrypts message from chinese, could add handling for no encrypted message,
		// but the encrypt and decrypt commands wont be manually called in final version anyway
		case "decrypt":
			fmt.Printf("> ")
			// enter message
			encryptedMessage, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}

			decryptedMessage := cipher.DecryptMessage(encryptedMessage)
			fmt.Println("Decrypted message:\n< ", decryptedMessage)
			fmt.Println()

		// Sets seed for chinese encryption
		case "set seed":
			fmt.Printf("> ")
			// enter seed
			seedString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}
			// turn seed string into int64
			seed, err = strconv.ParseInt(strings.TrimSpace(seedString), 10, 64)
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}
			// set new cipher based on new seed
			cipher = *encryption.SetSeed(seed)
			fmt.Println("\nSeed set")
			fmt.Println()

		// Quit application.
		case "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			// Exits loop and program.
			return

		// Clear screen
		case "clear":
			// AAAAHHHHH I HATE ITTT I HATE IT
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
			fmt.Println("commands: test animation, test gui, test conncetion, quit, set seed, encrypt, decrypt")

		// Invalid command
		default:
			// There's a fuckin list.
			fmt.Println("Invalid command.")
			fmt.Println()
		}

	}
}
