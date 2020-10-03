package main

import (
	"github.com/awesome-gocui/gocui"
)

func reset() error {
	running = true

	if err := deleteSnekBody(); err != nil {
		return err
	}

	var err error
	snekHead.position, err = setViewAtRandomPosition(snekHead.viewName, positionMatrix, true)
	if err != nil {
		return err
	}
	foodView.position, err = setViewAtRandomPosition(foodView.name, positionMatrix, false)
	if err != nil {
		return err
	}

	headDirection = direction(r.Intn(4))
	snekHead.currentDirection = headDirection
	snekBodyParts = []*snekBodyPart{snekHead}

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
