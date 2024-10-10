package gui

//TODO: Add movement between tabs

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

// GUI struct
// this represents the graphical user interface of the application.
// It holds the gocui.Gui instance and manages the layout, keybindings, and navigation.
type GUI struct {
	gui *gocui.Gui
}

// Initialize new GUI
func NewGUI() *GUI {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	return &GUI{gui: g}
}

// Start the GUI
func (g *GUI) Start() {
	defer g.gui.Close()

	g.gui.SetManagerFunc(g.layout)

	if err := g.setKeyBindings(); err != nil {
		log.Panicln(err)
	}

	g.gui.Update(func(gui *gocui.Gui) error {
		// Set initial tab
		if _, err := g.gui.SetCurrentView("settings"); err != nil {
			log.Panicln("Error setting current view: ", err)
		}

		initialView, err := g.gui.View("settings")
		if err != nil {
			log.Panicln("Error getting initial view: ", err)
		}
		initialView.FgColor = gocui.ColorGreen
		initialView.SelFgColor = gocui.ColorGreen
		return nil
	})

	// quit keys
	//TODO: Move these into a keybindings functions and add movement keys hjkl & arrows
	// I wanna put this in menu naviagation

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

	// Settings tab
	if settingsTab, err := gui.SetView("settings", 0, 0, maxX/2-20, 14); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout
		settingsTab.Title = "Settings"
		fmt.Fprintln(settingsTab)
		fmt.Fprintln(settingsTab, "Host: 127.0.0.1")
		fmt.Fprintln(settingsTab, "Port: 12345")
		fmt.Fprintln(settingsTab, "Seed: XXXXX")
	}

	// Conncetions tab
	if connectionsTab, err := gui.SetView("connections", 0, 15, maxX/2-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout
		connectionsTab.Title = "Connections"

		fmt.Fprintln(connectionsTab)
		fmt.Fprintln(connectionsTab, " [1] 127.0.0.1")
		fmt.Fprintln(connectionsTab, " [2] 127.0.0.1")
		fmt.Fprintln(connectionsTab, " [3] 127.0.0.1")
	}

	// chat tab
	if chatboxTab, err := gui.SetView("chat", maxX/2-18, 0, maxX-1, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		chatboxTab.Title = "Chat"
		fmt.Fprintln(chatboxTab)
		fmt.Fprintln(chatboxTab, " ")
	}

	// type message tab
	if messageTab, err := gui.SetView("message", maxX/2-18, maxY-5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		messageTab.Title = "Message"
		fmt.Fprintln(messageTab)
		fmt.Fprintln(messageTab, " ")
	}

	return nil
}

// Navigation map for tabs
var viewNavigation = map[string]map[string]string{
	"chat":        {"left": "settings", "down": "chat"},
	"message":     {"left": "connections", "up": "message"},
	"settings":    {"right": "chat", "down": "connections"},
	"connections": {"right": "chat", "up": "settings"},
}

// Movement function
func (g *GUI) move(gui *gocui.Gui, v *gocui.View, direction string) error {
	if v == nil {
		return nil
	}

	currentView := v.Name()

	v.FgColor = gocui.ColorWhite
	// Check for available direction

	if targetViewName, ok := viewNavigation[currentView][direction]; ok {
		nextView, err := gui.View(targetViewName)
		if err != nil {
			return err
		}

		nextView.FgColor = gocui.ColorGreen

		_, err = gui.SetCurrentView(targetViewName)
		if err != nil {
			return fmt.Errorf("Erorr setting view %s %v", direction, currentView)
		}
	}

	return nil
}

// Initialize keybindings
func (g *GUI) setKeyBindings() error {

	if err := g.gui.SetKeybinding("", 'h', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "left")
	}); err != nil {
		return err
	}
	if err := g.gui.SetKeybinding("", 'j', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "down")
	}); err != nil {
		return err
	}
	if err := g.gui.SetKeybinding("", 'k', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "up")
	}); err != nil {
		return err
	}
	if err := g.gui.SetKeybinding("", 'l', gocui.ModNone, func(gui *gocui.Gui, v *gocui.View) error {
		return g.move(gui, v, "right")
	}); err != nil {
		return err
	}
	return nil
}

// closes the GUI and returns to main.go
func (g *GUI) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
