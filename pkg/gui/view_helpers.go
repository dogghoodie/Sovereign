package gui

import (
	"github.com/jroimartin/gocui"
)

// Close any popup window
func ClosePopup(g *gocui.Gui, v *gocui.View) error {
	if v != nil && len(v.Name()) > 6 && v.Name()[:6] == "popup_" {
		return g.DeleteView(v.Name())
	}
	return nil
}
