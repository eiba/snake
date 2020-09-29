package main

import (
	"github.com/awesome-gocui/gocui"
)

const gameOverViewName = "gameOver"
var gameOverView *gocui.View

func initGameOverView(gui *gocui.Gui) error {
	maxX, maxY, err := getMaxXY(gui, gameViewName)
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
		viewPositionX,
		viewPositionX + viewLenX,
		viewPositionY,
		viewPositionY + viewLenY}

	gameOverView, err = createView(gui, gameOverViewProperties, false)
	return err
}

func gameOver(gui *gocui.Gui) error {
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
