package main

import (
	"github.com/awesome-gocui/gocui"
)

const pauseViewName = "pause"

func initPauseView(g *gocui.Gui) error {
	maxX, maxY, err := getMaxXY(g, gameViewName)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (maxX/2)-10, (maxY/2)-2
	viewLenX := 20
	viewLenY := 4

	pauseViewText := "Press P to resume"
	pauseView := view{
		pauseViewName,
		"Pause",
		pauseViewText,
		viewPositionX,
		viewPositionX + viewLenX,
		viewPositionY,
		viewPositionY + viewLenY}
	return createView(g, pauseView, false)
}

func pause(g *gocui.Gui) error {
	if gameFinished {
		return nil
	}

	v, err := g.View(pauseViewName)
	if err != nil {
		return err
	}

	if running {
		v.Visible = true
		if _, err := g.SetCurrentView(pauseViewName); err != nil {
			return err
		}
		if _, err := g.SetViewOnTop(pauseViewName); err != nil {
			return err
		}
	} else {
		v.Visible = false
	}
	running = !running
	return nil
}
