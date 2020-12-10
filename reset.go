package main

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game"
	"github.com/eiba/snake/game/view"
)

func reset() error {
	running = true

	if err := deletesnakeBody(); err != nil {
		return err
	}

	var err error
	game.snakeHead.position, err = view.setViewAtRandomPosition(game.snakeHead.viewName, positionMatrix, true)
	if err != nil {
		return err
	}
	game.foodView.position, err = view.setViewAtRandomPosition(game.foodView.name, positionMatrix, false)
	if err != nil {
		return err
	}

	game.headDirection = game.direction(r.Intn(4))
	game.snakeHead.currentDirection = game.headDirection
	game.snakeBodyParts = []*game.snakeBodyPart{game.snakeHead}

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
	for i := 1; i < len(game.snakeBodyParts); i++ {
		if err := gui.DeleteView(game.snakeBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	return nil
}
