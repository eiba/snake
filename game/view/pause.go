package view

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game"
)

const pauseViewName = "pause"

var pauseView *gocui.View

func initPauseView(gui *gocui.Gui, gameView Properties) error {
	lenX, lenY, err := getLenXY(gui, gameView.Name)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (lenX/2)-10, (lenY/2)-2
	viewLenX := 20
	viewLenY := 4

	pauseViewText := "Press P to resume"
	pauseViewProps := Properties{
		pauseViewName,
		"Pause",
		pauseViewText,
		game.Position{
			X0: viewPositionX,
			Y0: viewPositionY,
			X1: viewPositionX + viewLenX,
			Y1: viewPositionY + viewLenY}}
	pauseView, err = createView(gui, pauseViewProps, false)
	return err
}

func Pause(gui *gocui.Gui, gameFinished bool, running bool) error {
	if gameFinished {
		return nil
	}

	if running {
		pauseView.Visible = true
		if _, err := gui.SetCurrentView(pauseViewName); err != nil {
			return err
		}
		if _, err := gui.SetViewOnTop(pauseViewName); err != nil {
			return err
		}
	} else {
		pauseView.Visible = false
	}
	return nil
}
