package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

func pause(g *gocui.Gui) error {
	if running {
		return pauseGame(g)
	}
	return resumeGame(g)
}

func pauseGame(g *gocui.Gui) error {
	running = false
	x0, y0, x1, y1, err := g.ViewPosition(gameView)
	if err != nil {
		return err
	}
	maxX, maxY := x1-x0, y1-y0

	positionX, positionY := (maxX/2)-10, (maxY/2)-2

	lenX := 20
	lenY := 4
	name := "pause"
	if v, err := g.SetView(name, positionX, positionY, positionX+lenX, positionY+lenY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}

		v.Title = "pause"
		fmt.Fprintln(v, "\n press p to resume")

		if _, err := g.SetCurrentView(name); err != nil {
			return err
		}
		if _, err := g.SetViewOnTop(name); err != nil {
			return err
		}
	}
	return nil
}

func resumeGame(g *gocui.Gui) error {
	if err := g.DeleteView("pause"); err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	running = true
	return nil
}
