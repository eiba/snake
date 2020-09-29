package main

import (
	"github.com/awesome-gocui/gocui"
)

const pauseViewName = "pause"
var pauseView *gocui.View

func initPauseView(g *gocui.Gui) error {
	maxX, maxY, err := getMaxXY(g, gameViewName)
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
	v, err := createView(g, pauseViewProps, false)
	pauseView = v
	return err
}

func pause(g *gocui.Gui) error {
	if gameFinished {
		return nil
	}

	if running {
		pauseView.Visible = true
		if _, err := g.SetCurrentView(pauseViewName); err != nil {
			return err
		}
		if _, err := g.SetViewOnTop(pauseViewName); err != nil {
			return err
		}
	} else {
		pauseView.Visible = false
	}
	running = !running
	return nil
}
