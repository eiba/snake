package main

import (
	"github.com/awesome-gocui/gocui"
)

const pauseViewName = "pause"

var pauseView *gocui.View

func initPauseView() error {
	lenX, lenY, err := getLenXY(gameView.name)
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
		position{
			viewPositionX,
			viewPositionY,
			viewPositionX + viewLenX,
			viewPositionY + viewLenY}}
	pauseView, err = createView(pauseViewProps, false)
	return err
}

func pause() error {
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
	running = !running
	return nil
}
