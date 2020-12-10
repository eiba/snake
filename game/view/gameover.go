package view

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake"
	"github.com/eiba/snake/game"
)

const gameOverViewName = "gameOver"

var gameOverView *gocui.View

func initGameOverView() error {
	lenX, lenY, err := getLenXY(main.gameView.name)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (lenX/2)-12, (lenY/2)-2
	viewLenX := 25
	viewLenY := 4

	gameOverViewProperties := viewProperties{
		gameOverViewName,
		"",
		"Press space to restart",
		game.position{
			viewPositionX,
			viewPositionY,
			viewPositionX + viewLenX,
			viewPositionY + viewLenY}}

	gameOverView, err = createView(gameOverViewProperties, false)
	return err
}

func gameOver(title string) error {
	gameOverView.Visible = true
	gameOverView.Title = title
	if _, err := main.gui.SetCurrentView(gameOverViewName); err != nil {
		return err
	}
	if _, err := main.gui.SetViewOnTop(gameOverViewName); err != nil {
		return err
	}
	main.running = false
	main.gameFinished = true
	return nil
}
