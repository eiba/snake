package game

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake"
	"github.com/eiba/snake/game/view"
)

func reset() error {
	main.running = true

	if err := deletesnakeBody(); err != nil {
		return err
	}

	var err error
	snakeHead.position, err = view.setViewAtRandomPosition(snakeHead.viewName, main.positionMatrix, true)
	if err != nil {
		return err
	}
	foodView.position, err = view.setViewAtRandomPosition(foodView.name, main.positionMatrix, false)
	if err != nil {
		return err
	}

	headDirection = Direction(main.r.Intn(4))
	snakeHead.currentDirection = headDirection
	snakeBodyParts = []*snakeBodyPart{snakeHead}

	main.gameOverView.Visible = false
	main.pauseView.Visible = false
	main.loadingView.Visible = false
	main.gameFinished = false

	main.foodPath = []main.node{}
	main.pathIndex = -1

	if err := main.updateStat(&main.restartStat, main.restartStat.value+1); err != nil {
		return err
	}
	if err := main.updateStat(&main.lengthStat, 1); err != nil {
		return err
	}
	return nil
}

func deletesnakeBody() error {
	for i := 1; i < len(snakeBodyParts); i++ {
		if err := main.gui.DeleteView(snakeBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	return nil
}
