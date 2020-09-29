package main

import (
	"github.com/awesome-gocui/gocui"
)

const gameOverViewName = "gameOver"
var gameOverView *gocui.View

func initGameOverView(g *gocui.Gui) error {
	maxX, maxY, err := getMaxXY(g, gameViewName)
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

	v, err := createView(g, gameOverViewProperties, false)
	gameOverView = v
	return err
}

func gameOver(g *gocui.Gui) error {
	gameOverView.Visible = true
	if _, err := g.SetCurrentView(gameOverViewName); err != nil {
		return err
	}
	if _, err := g.SetViewOnTop(gameOverViewName); err != nil {
		return err
	}
	running = false
	gameFinished = true
	return nil
}
