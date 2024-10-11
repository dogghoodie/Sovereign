package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

//TODO: Look into changing tab border color for currentView
// and figure out floating tabs.

// GUI struct
// This represents the graphical user interface of the application.
// It holds the gocui.Gui instance and manages the layout, keybindings, and navigation.
type GUI struct {
	gui *gocui.Gui
}

// Initialize new GUI.
func NewGUI() *GUI {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	return &GUI{gui: g}
}

// Start the GUI.
func (g *GUI) Start() {
	// Ensure terminal is restored after gui closes.
	defer g.gui.Close()

	// Initialize the gui manager: layout.
	g.gui.SetManagerFunc(g.layout)

	// Initialize keybindings.
	if err := g.setKeyBindings(); err != nil {
		log.Panicln(err)
	}

	// Update the gui safely with
	// gocui's "thread-safe way".
	g.gui.Update(func(gui *gocui.Gui) error {
		// Set initial tab to settings (top left).
		if _, err := g.gui.SetCurrentView("settings"); err != nil {
			log.Panicln("Error setting current view: ", err)
		}
		// Retrieve initial view (settings).
		initialView, err := g.gui.View("settings")
		if err != nil {
			log.Panicln("Error getting initial view: ", err)
		}
		//TODO: Combine intial into selected? is that a thing?
		//
		// Set the intial foreground to green.
		initialView.FgColor = gocui.ColorGreen
		// Set the selected foreground to green.
		initialView.SelFgColor = gocui.ColorGreen
		return nil
	})

	// Keep gui running, and log any error other than quit.
	if err := g.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// Defines the layout for the GUI.
func (g *GUI) layout(gui *gocui.Gui) error {
	// Get max x,y values for gui grid.
	maxX, maxY := gui.Size()

	// Settings tab for setting seed and connection settings.
	if settingsTab, err := gui.SetView("settings", 0, 0, maxX/2-20, 14); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout.
		settingsTab.Title = "Settings"
		fmt.Fprintln(settingsTab)
		fmt.Fprintln(settingsTab, "Host: 127.0.0.1")
		fmt.Fprintln(settingsTab, "Port: 12345")
		fmt.Fprintln(settingsTab, "Seed: XXXXX")
	}

	// Conncetions tab for viewing chat partners.
	if connectionsTab, err := gui.SetView("connections", 0, 15, maxX/2-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout shit.
		connectionsTab.Title = "Connections"
		fmt.Fprintln(connectionsTab)
		fmt.Fprintln(connectionsTab)
		fmt.Fprintln(connectionsTab, " Active                Inactive")
		fmt.Fprintln(connectionsTab)
		fmt.Fprintln(connectionsTab, " [1] 127.0.0.1        [1] 127.0.0.1")
		fmt.Fprintln(connectionsTab, " [1] 127.0.0.1        [1] 127.0.0.1")
		fmt.Fprintln(connectionsTab, " [1] 127.0.0.1        [1] 127.0.0.1")
		fmt.Fprintln(connectionsTab, " [1] 127.0.0.1        [1] 127.0.0.1")
		fmt.Fprintln(connectionsTab, " [1] 127.0.0.1        [1] 127.0.0.1")
		fmt.Fprintln(connectionsTab, " [1] 127.0.0.1        [1] 127.0.0.1")
	}

	// Chat tab for viewing messages.
	if chatboxTab, err := gui.SetView("chat", maxX/2-18, 0, maxX-1, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		chatboxTab.Title = "Chat"
		fmt.Fprintln(chatboxTab)
		fmt.Fprintln(chatboxTab, " ")
	}

	// Message tab for writing to the chat.
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

// Navigation map for tabs.
var viewNavigation = map[string]map[string]string{
	"chat":        {"left": "settings", "down": "chat"},
	"message":     {"left": "connections", "up": "message"},
	"settings":    {"right": "chat", "down": "connections"},
	"connections": {"right": "chat", "up": "settings"},
}

// Movement function.
func (g *GUI) move(gui *gocui.Gui, v *gocui.View, direction string) error {
	// If the tab we're interacting with doesn't exist (nil) just chill. don't crash.
	if v == nil {
		return nil
	}

	// Retrieve current view (tab).
	currentView := v.Name()

	// Set the current tab to white before moving.
	v.FgColor = gocui.ColorWhite

	// Check if target tab exists and initalize.
	if targetViewName, ok := viewNavigation[currentView][direction]; ok {
		nextView, err := gui.View(targetViewName)
		if err != nil {
			return err
		}
		// Set the target view as the current view.
		_, err = gui.SetCurrentView(targetViewName)
		if err != nil {
			return fmt.Errorf("Erorr setting view %s %v", direction, currentView)
		}
		// Change the color of the current view.
		nextView.FgColor = gocui.ColorGreen
	}

	return nil
}

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

// closes the GUI and returns to main func.
func (g *GUI) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
