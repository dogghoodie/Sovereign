package visuals

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Draw() {

	var sourceFile [5]string

	// Define relative path to the logo.txt file
	sourceFile[0] = "resources/logo.txt"
	sourceFile[1] = "resources/logo2.txt"
	sourceFile[2] = "resources/logo3.txt"
	sourceFile[3] = "resources/logo4.txt"
	sourceFile[4] = "resources/logo5.txt"

	for i := 0; i < 5; i++ {

		// Open the file
		file, err := os.Open(sourceFile[i])
		if err != nil {
			fmt.Println("Error openig file: ", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// This method is for more accurate
		// control over each char for coloring
		/*	line := scanner.Text()
			for _, char := range line {
				switch char {
				case '/':
					fmt.Print(string(char))
				case '|':
					fmt.Print(string(char))
				case '_':
					fmt.Print(string(char))
				case '-':
					fmt.Print(string(char))
				case '\\':
					fmt.Print(string(char))
				default:
					fmt.Print(string(char))
			}
		*/

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
		if i < 2 {
			time.Sleep(120 * time.Millisecond)
			ClearScreen()
		} else if i < 4 {
			time.Sleep(220 * time.Millisecond)
			ClearScreen()
		} else if i < 5 {
			time.Sleep(1000 * time.Millisecond)
			ClearScreen()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file: ", err)
		}
	}
}
