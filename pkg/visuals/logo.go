package visuals

import (
	"bufio"
	"fmt"
	"os"
)

func Draw() {
	// Define relative path to the logo.txt file
	sourceFile := "resources/logo.txt"

	// Open the file
	file, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("Error openig file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println()

	// Old logo print method, line by line
	/*
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(Colors.RED + line + Colors.ANSI_RESET)
		}
	*/

	// New logo print method, one char at a time for more
	// accurate color control.
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			switch char {
			case '/':
				fmt.Print(Colors.COMMENT + string(char) + Colors.ANSI_RESET)
			case '|':
				fmt.Print(Colors.COMMENT + string(char) + Colors.ANSI_RESET)
			case '_':
				fmt.Print(Colors.COMMENT + string(char) + Colors.ANSI_RESET)
			case '-':
				fmt.Print(Colors.COMMENT + string(char) + Colors.ANSI_RESET)
			case '\\':
				fmt.Print(Colors.COMMENT + string(char) + Colors.ANSI_RESET)
			default:
				fmt.Print(Colors.WHITE + string(char) + Colors.ANSI_RESET)
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

}
