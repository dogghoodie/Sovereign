package panel

import (
	"Sovereign/pkg/visuals"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreatePanel() {
	fmt.Println(visuals.Colors.RED + "Sovereign Panel." + visuals.Colors.ANSI_RESET)

	for {
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		command = strings.ReplaceAll(command, "\n", "")

		switch command {
		case "ping":
			fmt.Println("pong")
			fmt.Println()
		case "connect":
			fmt.Println("Add connections menu here.")
			fmt.Println()
		case "quit":
			fmt.Println("Quitting...")
			fmt.Println()
			return
		}

		if err != nil {
			fmt.Print("There was an error:", err)
			continue
		}

	}
}
