package view

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game"
)

const loadingViewName = "loading"

var loadingView *gocui.View

func initLoadingView(gui *gocui.Gui, gameView Properties) error {
	lenX, lenY, err := getLenXY(gui, gameView.Name)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (lenX/2)-13, (lenY/2)-2
	viewLenX := 26
	viewLenY := 4

	loadingViewText := "Initiating autopilot..."
	loadingViewProps := Properties{
		Name:  loadingViewName,
		Title: "Loading",
		Text:  loadingViewText,
		Position: game.Position{
			X0: viewPositionX,
			Y0: viewPositionY,
			X1: viewPositionX + viewLenX,
			Y1: viewPositionY + viewLenY}}
	loadingView, err = createView(gui, loadingViewProps, false)
	return err
}

func Loading(gui *gocui.Gui, gameFinished bool, running bool, loading bool) error {
	if gameFinished && !running {
		return nil
	}
	loadingView.Visible = loading
	if loading {
		if _, err := gui.SetCurrentView(loadingViewName); err != nil {
			return err
		}
		if _, err := gui.SetViewOnTop(loadingViewName); err != nil {
			return err
		}
	}
	return nil
}
