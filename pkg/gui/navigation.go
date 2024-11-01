package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// Movement function.
func Move(gui *gocui.Gui, v *gocui.View, direction string) error {
	// If the tab we're interacting with doesn't exist (nil) just chill. don't crash.
	if v == nil {
		return nil
	}

	// Retrieve current view (tab).
	currentView := v.Name()

	// Check if target tab exists and initalize.
	if targetViewName, ok := viewNavigation[currentView][direction]; ok {
		nextView, err := gui.View(targetViewName)
		if err != nil {
			return err
		}
		// Set the current tab to white before moving.
		v.FgColor = gocui.ColorDefault
		// Set the target view as the current view.
		_, err = gui.SetCurrentView(targetViewName)
		if err != nil {
			return fmt.Errorf("Erorr setting view %s %v", direction, currentView)
		}
		// Change the color of the current view.
		nextView.FgColor = gocui.ColorGreen
		// nextView.TitleColor = gocui.ColorGreen
	}

	return nil
}
