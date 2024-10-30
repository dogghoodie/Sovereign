package gui

// This setup can probably just stay the same.

import (
	"github.com/jroimartin/gocui"
	"log"
)

// Initialize keybindings.
func (g *GUI) setKeyBindings() error {
	// Bind h to move left.
	if err := g.gui.SetKeybinding("", 'h', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "left")
	}); err != nil {
		return err
	}

	// Bind j to move down.
	if err := g.gui.SetKeybinding("", 'j', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "down")
	}); err != nil {
		return err
	}

	// Bind k to move up.
	if err := g.gui.SetKeybinding("", 'k', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "up")
	}); err != nil {
		return err
	}

	// Bind l to move right.
	if err := g.gui.SetKeybinding("", 'l', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "right")
	}); err != nil {
		return err
	}

	// Bind Ctrl C to quit.
	if err := g.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, g.quit); err != nil {
		log.Panicln(err)
	}

	// Bind q to quit.
	if err := g.gui.SetKeybinding("", 'q', gocui.ModNone, g.quit); err != nil {
		log.Panicln(err)
	}

	return nil
}
