package main

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game"
	"github.com/eiba/snake/game/view"
)

const gameOverViewName = "gameOver"

var gameOverView *gocui.View

func initGameOverView() error {
	lenX, lenY, err := view.getLenXY(gameView.name)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (lenX/2)-12, (lenY/2)-2
	viewLenX := 25
	viewLenY := 4

	gameOverViewProperties := view.viewProperties{
		gameOverViewName,
		"",
		"Press space to restart",
		game.position{
			viewPositionX,
			viewPositionY,
			viewPositionX + viewLenX,
			viewPositionY + viewLenY}}

	gameOverView, err = view.createView(gameOverViewProperties, false)
	return err
}

func gameOver(title string) error {
	gameOverView.Visible = true
	gameOverView.Title = title
	if _, err := gui.SetCurrentView(gameOverViewName); err != nil {
		return err
	}
	if _, err := gui.SetViewOnTop(gameOverViewName); err != nil {
		return err
	}
	running = false
	gameFinished = true
	return nil
}
