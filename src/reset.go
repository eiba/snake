package main

import (
	"github.com/awesome-gocui/gocui"
)

func reset() error {
	running = true

	if err := deleteSnekBody(); err != nil {
		return err
	}

	headDirection = direction(r.Intn(4))
	snekHead = &snekBodyPart{headDirection, headDirection, "s0",position{0,0,0,0}}
	snekBodyParts = []*snekBodyPart{snekHead}

	position, err := setViewAtRandom(snekHead.viewName, true); if err != nil {
		return err
	}
	snekHead.position = position
	if _, err := setViewAtRandom(boxViewName, false); err != nil {
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

func deleteSnekBody() error {
	for i := 1; i < len(snekBodyParts); i++ {
		if err := gui.DeleteView(snekBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	return nil
}
