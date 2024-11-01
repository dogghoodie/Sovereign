package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// Navigation map for tabs.
var viewNavigation = map[string]map[string]string{
	"chat":        {"left": "settings", "down": "message"},
	"message":     {"left": "connections", "up": "chat"},
	"settings":    {"right": "chat", "down": "connections"},
	"connections": {"right": "message", "up": "settings"},
}

// Defines the layout for the GUI.
func Layout(gui *gocui.Gui) error {
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
	}

	// Conncetions tab for viewing chat partners.
	if connectionsTab, err := gui.SetView("connections", 0, 15, maxX/2-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout shit.
		connectionsTab.Title = "Connections"
		fmt.Fprintln(connectionsTab)
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

		// TODO: Make messageTab editable, save input somewhere as 'message'
		// reading this for now:
		// https://gist.github.com/jroimartin/3b2e943a3811d795e0718b4a95b89bec

		messageTab.Editable = true
		messageTab.Wrap = true
		fmt.Fprintln(messageTab)
		fmt.Fprintln(messageTab, ">")
	}

	return nil
}
