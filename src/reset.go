package main

import (
	"github.com/awesome-gocui/gocui"
)

func reset() error {
	running = true

	if err := deletesnakeBody(); err != nil {
		return err
	}

	var err error
	snakeHead.position, err = setViewAtRandomPosition(snakeHead.viewName, positionMatrix, true)
	if err != nil {
		return err
	}
	foodView.position, err = setViewAtRandomPosition(foodView.name, positionMatrix, false)
	if err != nil {
		return err
	}

	headDirection = direction(r.Intn(4))
	snakeHead.currentDirection = headDirection
	snakeBodyParts = []*snakeBodyPart{snakeHead}

	gameOverView.Visible = false
	pauseView.Visible = false
	loadingView.Visible = false
	gameFinished = false

	foodPath = []node{}
	pathIndex = -1

	if err := updateStat(&restartStat, restartStat.value+1); err != nil {
		return err
	}
	if err := updateStat(&lengthStat, 1); err != nil {
		return err
	}
	return nil
}

func deletesnakeBody() error {
	for i := 1; i < len(snakeBodyParts); i++ {
		if err := gui.DeleteView(snakeBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	return nil
}
