package visuals

import (
	"github.com/muesli/termenv"
	"os"
)

func Intro() {
	// Initialize output handler (formatting).
	outputHandler := termenv.NewOutput(os.Stdout)
	// Hide the cursor.
	outputHandler.HideCursor()
	// Clear the terminal.
	ClearScreen()
	// Start the launch animation.
	Draw_Launch_Animation()
	// Clear the terminal.
	ClearScreen()
	// Start the logo animation.
	Draw_Logo()
	// Bring back terminal cursor.
	outputHandler.ShowCursor()
	// Clear the terminal.
	ClearScreen()

}
