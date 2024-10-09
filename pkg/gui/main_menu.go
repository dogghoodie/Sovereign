package gui

//TODO: Add movement between tabs

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

type GUI struct {
	gui *gocui.Gui
}

func NewGUI() *GUI {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	return &GUI{gui: g}
}

func (g *GUI) Start() {
	defer g.gui.Close()

	g.gui.SetManagerFunc(g.layout)

	// Bind Ctrl+C to quit program
	if err := g.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, g.quit); err != nil {
		log.Panicln(err)
	}

	// Bind q to quit gui
	if err := g.gui.SetKeybinding("", 'q', gocui.ModNone, g.quit); err != nil {
		log.Panicln(err)
	}

	if err := g.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// defines the layout for the GUI
func (g *GUI) layout(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()

	if v, err := gui.SetView("settings", 0, 0, maxX/2-20, 14); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout
		v.Title = "Settings"
		fmt.Fprintln(v)
		fmt.Fprintln(v, "Host: 127.0.0.1")
		fmt.Fprintln(v, "Port: 12345")
		fmt.Fprintln(v, "Seed: XXXXX")
	}

	if v, err := gui.SetView("connections", 0, 15, maxX/2-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout
		v.Title = "Connections"
		fmt.Fprintln(v)
		fmt.Fprintln(v, " [1] 127.0.0.1")
		fmt.Fprintln(v)
		fmt.Fprintln(v, " [2] 127.0.0.1")
		fmt.Fprintln(v)
		fmt.Fprintln(v, " [3] 127.0.0.1")
	}

	if v, err := gui.SetView("chatbox", maxX/2-18, 0, maxX-1, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Conversation"
		fmt.Fprintln(v)
		fmt.Fprintln(v, " ")
	}

	if v, err := gui.SetView("messagebox", maxX/2-18, maxY-5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Message"
		fmt.Fprintln(v)
		fmt.Fprintln(v, " ")
	}
	return nil
}

// closes the GUI and returns to main.go
func (g *GUI) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
