package main

import (
	"github.com/awesome-gocui/gocui"
)

const pauseViewName = "pause"

var pauseView *gocui.View

func initPauseView() error {
	maxX, maxY, err := getMaxXY(gameViewName)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (maxX/2)-10, (maxY/2)-2
	viewLenX := 20
	viewLenY := 4

	pauseViewText := "Press P to resume"
	pauseViewProps := viewProperties{
		pauseViewName,
		"Pause",
		pauseViewText,
		viewPositionX,
		viewPositionX + viewLenX,
		viewPositionY,
		viewPositionY + viewLenY}
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
