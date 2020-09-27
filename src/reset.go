package main

import (
	"github.com/awesome-gocui/gocui"
	"time"
)

func reset(g *gocui.Gui) error {
	running = true
	headDirection = direction(r.Intn(4))
	tickInterval = 50 * time.Millisecond

	if err := deleteSnekBody(g); err != nil {
		return err
	}
	snekBodyParts = []snekBodyPart{{headDirection, headDirection, "s0"}}

	if err := setViewAtRandom(g, snekBodyParts[0].viewName, true); err != nil {
		return err
	}
	if err := setViewAtRandom(g, boxViewName, false); err != nil {
		return err
	}
	if err := g.DeleteView(gameOverViewName); err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	return nil
}

func deleteSnekBody(g *gocui.Gui) error {
	for i := 1; i < len(snekBodyParts); i++ {
		if err := g.DeleteView(snekBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	return nil
}