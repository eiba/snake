package view

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake"
	"github.com/eiba/snake/game"
)

const pauseViewName = "pause"

var pauseView *gocui.View

func initPauseView() error {
	lenX, lenY, err := getLenXY(main.gameView.name)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (lenX/2)-10, (lenY/2)-2
	viewLenX := 20
	viewLenY := 4

	pauseViewText := "Press P to resume"
	pauseViewProps := viewProperties{
		pauseViewName,
		"Pause",
		pauseViewText,
		game.position{
			viewPositionX,
			viewPositionY,
			viewPositionX + viewLenX,
			viewPositionY + viewLenY}}
	pauseView, err = createView(pauseViewProps, false)
	return err
}

func pause() error {
	if main.gameFinished {
		return nil
	}

	if main.running {
		pauseView.Visible = true
		if _, err := main.gui.SetCurrentView(pauseViewName); err != nil {
			return err
		}
		if _, err := main.gui.SetViewOnTop(pauseViewName); err != nil {
			return err
		}
	} else {
		pauseView.Visible = false
	}
	main.running = !main.running
	return nil
}
