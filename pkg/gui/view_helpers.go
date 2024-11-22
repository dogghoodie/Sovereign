package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// Close any popup window
func ClosePopup(g *gocui.Gui, v *gocui.View) error {
	// Delete current popup
	if v != nil && len(v.Name()) > 6 && v.Name()[:6] == "popup_" {
		if err := g.DeleteView(v.Name()); err != nil {
			return err
		}

		// Move back to previous view
		if previousViewName != "" {
			if _, err := g.SetCurrentView(previousViewName); err != nil {
				return err
			}
		}
	}

	return nil
}

func SendMessage(g *gocui.Gui, v *gocui.View) error {
	message := v.Buffer()
	AppendToChat(g, message)
	v.Clear()

	return nil
}

func AppendToChat(g *gocui.Gui, message string) error {
	g.Update(func(g *gocui.Gui) error {
		chatboxTab, err := g.View("chat")
		if err != nil {
			return err
		}
		fmt.Fprintln(chatboxTab, message)
		_, y := chatboxTab.Size()
		chatboxTab.SetOrigin(0, y)
		return nil
	})

	return nil
}
