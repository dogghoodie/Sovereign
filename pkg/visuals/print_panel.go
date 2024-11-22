package visuals

import (
	"fmt"
)

func Print_Panel() {
	// Clear the terminal.
	ClearScreen()
	// Draw the "dev panel". xD
	fmt.Println("**************")
	fmt.Println("* Dev Panel: *")
	fmt.Println("**************")
	fmt.Println()
	// Testing all colors for fun
	fmt.Println(
		Colors.CYAN+"x",
		Colors.GREEN+"x",
		Colors.ORANGE+"x",
		Colors.PINK+"x",
		Colors.PURPLE+"x",
		Colors.RED+"x",
		Colors.YELLOW+"x",
		Colors.BLACK+"x",
		Colors.ANSI_RESET,
	)
	fmt.Println()
	fmt.Println("commands: test animation, test gui, test conncetion, quit, set seed, encrypt, decrypt")
}
