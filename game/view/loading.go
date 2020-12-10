package view

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake"
	"github.com/eiba/snake/game"
)

const loadingViewName = "loading"

var loadingView *gocui.View

func initLoadingView() error {
	lenX, lenY, err := getLenXY(main.gameView.name)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (lenX/2)-13, (lenY/2)-2
	viewLenX := 26
	viewLenY := 4

	loadingViewText := "Initiating autopilot..."
	loadingViewProps := viewProperties{
		loadingViewName,
		"Loading",
		loadingViewText,
		game.position{
			viewPositionX,
			viewPositionY,
			viewPositionX + viewLenX,
			viewPositionY + viewLenY}}
	loadingView, err = createView(loadingViewProps, false)
	return err
}

func loading(loading bool) error {
	if main.gameFinished && !main.running {
		return nil
	}
	loadingView.Visible = loading
	if loading {
		if _, err := main.gui.SetCurrentView(loadingViewName); err != nil {
			return err
		}
		if _, err := main.gui.SetViewOnTop(loadingViewName); err != nil {
			return err
		}
	}
	return nil
}
