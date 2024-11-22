package main

import (
	"Sovereign/pkg/encryption"
	"Sovereign/pkg/gui"
	"Sovereign/pkg/visuals"
	"bufio"
	"fmt"
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
	visuals.Print_Panel()
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
			continue
		}

		// Trim extra white space off of input.
		command = strings.TrimSpace(command)

		switch command {

		// Test the animation scene
		case "test animation":
			visuals.Intro()
			visuals.Print_Panel()

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

		// Encrypts message
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

		// Decrypt message
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

		// Sets seed for random chinese encryption
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

		// Clear screen
		case "clear":
			visuals.ClearScreen()
			// Draw the "dev panel". xD
			visuals.Print_Panel()

		// Quit application.
		case "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			// Exits loop and program.
			return

		// Invalid command
		default:
			// There's a fuckin list.
			fmt.Println("Invalid command.")
			fmt.Println()
		}

	}
}
