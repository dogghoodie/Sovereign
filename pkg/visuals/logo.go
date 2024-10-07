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

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(Colors.RED + line + Colors.ANSI_RESET)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

}
