package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

// TODO :
// Add black/darkgrey color for inactive view.

func configuration(g *gocui.Gui) {
	g.ASCII = false
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorDefault
}

func initializeView(g *gocui.Gui) error {
	g.Update(func(gui *gocui.Gui) error {
		// Set initial tab to settings (top left).
		if _, err := g.SetCurrentView("settings"); err != nil {
			log.Panicln("Error setting current view: ", err)
		}

		g.FgColor = gocui.ColorDefault
		g.SelFgColor = gocui.ColorGreen

		return nil
	})

	return nil
}

// Start the gui.
func Start() {
	// Initialize new gui
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	// Ensure GUI closes properly.
	defer g.Close()

	g.SetManagerFunc(Layout)
	configuration(g)
	initializeView(g)

	// Initialize keybindings.
	if err := setKeyBindings(g); err != nil {
		log.Panicln(err)
	}
	// Keep gui running, and log any error other than quit.
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// Quit the gui.
func Quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
