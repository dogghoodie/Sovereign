package gui

// TODO :
// figure out why ESC isn't binding properly or pick another
// or pick another key to close popups.

import (
	"github.com/jroimartin/gocui"
	"log"
)

// Initialize keybindings.
func setKeyBindings(g *gocui.Gui) error {

	// NAVIGATION

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

	// SETTINGS

	// Bind s to open settings.
	if err := g.SetKeybinding("", 's', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return ShowSettings(g, v)
	}); err != nil {
		return err
	}

	// CLOSE POPUPS

	// Bind b to close popups. (b for back I guess. ESC isn't workign)
	if err := g.SetKeybinding("", 'b', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return ClosePopup(g, v)
	}); err != nil {
		return err
	}

	// CLOSE APPLICATION

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
