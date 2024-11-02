package gui

// This setup can probably just stay the same.
// TODO : Add keybind for settings pop up menu.

import (
	"github.com/jroimartin/gocui"
	"log"
)

// Initialize keybindings.
func setKeyBindings(g *gocui.Gui) error {
	// Bind h to move left.
	if err := g.SetKeybinding("", 'h', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return Move(gui, v, "left")
	}); err != nil {
		return err
	}

	// Bind j to move down.
	if err := g.SetKeybinding("", 'j', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return Move(g, v, "down")
	}); err != nil {
		return err
	}

	// Bind k to move up.
	if err := g.SetKeybinding("", 'k', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return Move(g, v, "up")
	}); err != nil {
		return err
	}

	// Bind l to move right.
	if err := g.SetKeybinding("", 'l', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return Move(g, v, "right")
	}); err != nil {
		return err
	}

	// Bind s to open settings.
	if err := g.SetKeybinding("", 's', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return ShowSettings(g, v)
	}); err != nil {
		return err
	}

	// Bind ESC to close popups.
	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		// Why isn't this deleting??
		// could be escape key not binding
		return ClosePopup(g, v)
	}); err != nil {
		return err
	}

	// Bind Ctrl C to quit.
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	// Bind q to quit.
	if err := g.SetKeybinding("", 'q', gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	return nil
}
