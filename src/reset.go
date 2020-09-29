package main

import (
	"github.com/awesome-gocui/gocui"
	"time"
)

func reset(g *gocui.Gui) error {
	running = true
	tickInterval = 50 * time.Millisecond

	if err := deleteSnekBody(g); err != nil {
		return err
	}

	headDirection = direction(r.Intn(4))
	snekHead = &snekBodyPart{headDirection, headDirection, "s0"}
	snekBodyParts = []*snekBodyPart{snekHead}

	if err := setViewAtRandom(g, snekHead.viewName, true); err != nil {
		return err
	}
	if err := setViewAtRandom(g, boxViewName, false); err != nil {
		return err
	}
	if err := g.DeleteView(gameOverViewName); err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	if err := g.DeleteView(pauseViewName); err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	if err := updateStat(g, &restartStat, restartStat.value+1); err != nil {
		return err
	}
	if err := updateStat(g, &lengthStat, 1); err != nil {
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
