package main

import (
	"github.com/awesome-gocui/gocui"
)

const gameOverViewName = "gameOver"
var gameOverView *gocui.View

func initGameOverView() error {
	maxX, maxY, err := getMaxXY(gameViewName)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (maxX/2)-12, (maxY/2)-2
	viewLenX := 25
	viewLenY := 4

	gameOverViewProperties := viewProperties{
		gameOverViewName,
		"Game over",
		"Press space to restart",
		position{viewPositionX,
		viewPositionY,
		viewPositionX + viewLenX,
		viewPositionY + viewLenY}}

	gameOverView, err = createView(gameOverViewProperties, false)
	return err
}

func gameOver() error {
	gameOverView.Visible = true
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
