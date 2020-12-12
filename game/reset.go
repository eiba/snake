package game

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game/view"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func reset(gui *gocui.Gui, snakeBodyParts []*snakeBodyPart, positionMatrix [][]Position) error {
	//main.running = true

	if err := deleteSnakeBody(gui, snakeBodyParts); err != nil {
		return err
	}

	var err error
	snakeHead.position, err = view.SetViewAtRandomPosition(gui, snakeHead.viewName, positionMatrix, true)
	if err != nil {
		return err
	}
	foodView.Position, err = view.SetViewAtRandomPosition(gui, foodView.Name, positionMatrix, false)
	if err != nil {
		return err
	}

	headDirection = Direction(r.Intn(4))
	snakeHead.currentDirection = headDirection
	snakeBodyParts = []*snakeBodyPart{snakeHead}

	//main.gameOverView.Visible = false
	//main.pauseView.Visible = false
	//main.loadingView.Visible = false
	//main.gameFinished = false

	//main.foodPath = []main.node{}
	//main.pathIndex = -1

	if err := view.UpdateStat(&view.RestartStat, view.RestartStat.Value+1); err != nil {
		return err
	}
	if err := view.UpdateStat(&view.LengthStat, 1); err != nil {
		return err
	}
	return nil
}

func deleteSnakeBody(gui *gocui.Gui, snakeBodyParts []*snakeBodyPart) error {
	for i := 1; i < len(snakeBodyParts); i++ {
		if err := gui.DeleteView(snakeBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	return nil
}
