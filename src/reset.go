package main

import (
	"github.com/awesome-gocui/gocui"
	"time"
)

func reset(gui *gocui.Gui) error {
	running = true
	tickInterval = 50 * time.Millisecond

	if err := deleteSnekBody(gui); err != nil {
		return err
	}

	headDirection = direction(r.Intn(4))
	snekHead = &snekBodyPart{headDirection, headDirection, "s0"}
	snekBodyParts = []*snekBodyPart{snekHead}

	if err := setViewAtRandom(gui, snekHead.viewName, true); err != nil {
		return err
	}
	if err := setViewAtRandom(gui, boxViewName, false); err != nil {
		return err
	}

	gameOverView.Visible = false
	pauseView.Visible = false
	gameFinished = false
	if err := updateStat(&restartStat, restartStat.value+1); err != nil {
		return err
	}
	if err := updateStat(&lengthStat, 1); err != nil {
		return err
	}
	return nil
}

func deleteSnekBody(gui *gocui.Gui) error {
	for i := 1; i < len(snekBodyParts); i++ {
		if err := gui.DeleteView(snekBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	return nil
}
