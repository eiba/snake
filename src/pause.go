package main

import (
	"github.com/awesome-gocui/gocui"
)

const pauseViewName = "pause"

func pause(g *gocui.Gui) error {
	if running {
		return pauseGame(g)
	}
	return resumeGame(g)
}

func pauseGame(g *gocui.Gui) error {
	running = false

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
	return createView(g, pauseView)
}

func resumeGame(g *gocui.Gui) error {
	if err := g.DeleteView(pauseViewName); err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	running = true
	return nil
}
