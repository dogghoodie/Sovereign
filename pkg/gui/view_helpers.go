package gui

import (
	"github.com/jroimartin/gocui"
)

// Close any popup window
func ClosePopup(g *gocui.Gui, v *gocui.View) error {
	// Delete current popup
	if v != nil && len(v.Name()) > 6 && v.Name()[:6] == "popup_" {
		if err := g.DeleteView(v.Name()); err != nil {
			return err
		}

		if previousViewName != "" {
			if _, err := g.SetCurrentView(previousViewName); err != nil {
				return err
			}
		}
	}

	return nil
}
