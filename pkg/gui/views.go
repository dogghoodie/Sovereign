package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// Navigation map for tabs.
var viewNavigation = map[string]map[string]string{
	"chat":        {"left": "connections", "down": "message"},
	"message":     {"left": "connections", "up": "chat"},
	"connections": {"right": "message"},
}

var previousViewName string

func InitializeViews(g *gocui.Gui) error {
	if err := SetupMainViews(g); err != nil {
		return err
	}
	return nil
}

// Defines the layout for the GUI.
func SetupMainViews(g *gocui.Gui) error {
	// Get max x,y values for gui grid.
	maxX, maxY := g.Size()

	// Conncetions tab for viewing chat partners.
	if connectionsTab, err := g.SetView("connections", 0, 0, maxX/2-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout shit.
		connectionsTab.Title = "Connections"
		fmt.Fprintln(connectionsTab)
	}

	// Chat tab for viewing messages.
	if chatboxTab, err := g.SetView("chat", maxX/2-18, 0, maxX-1, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		chatboxTab.Title = "Chat"
		fmt.Fprintln(chatboxTab)
		fmt.Fprintln(chatboxTab, " ")
	}

	// Message tab for writing to the chat.
	if messageTab, err := g.SetView("message", maxX/2-18, maxY-5, maxX-1, maxY-1); err != nil {
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

func ShowSettings(g *gocui.Gui, v *gocui.View) error {
	// If we're already in settings, don't reopen settings
	if existingPopup, _ := g.View("popup_settings"); existingPopup != nil {
		return nil
	}

	// Get max x,y values for gui grid.
	maxX, maxY := g.Size()

	// Save previous view
	if v != nil {
		previousViewName = v.Name()
	}

	// Popup dimensions
	width, height := 70, 20

	x0, y0 := (maxX-width)/2, (maxY-height)/2
	x1, y1 := x0+width, y0+height

	// Settings tab for setting seed and connection settings.
	if settingsPopup, err := g.SetView("popup_settings", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Testing layout.
		settingsPopup.Title = "Settings"
		fmt.Fprintln(settingsPopup)
	}

	if _, err := g.SetCurrentView("popup_settings"); err != nil {
		return err
	}

	return nil
}
