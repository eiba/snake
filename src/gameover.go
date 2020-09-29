package main

import (
	"github.com/awesome-gocui/gocui"
)

const gameOverViewName = "gameOver"

func initGameOverView(g *gocui.Gui) error {
	maxX, maxY, err := getMaxXY(g, gameViewName)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (maxX/2)-12, (maxY/2)-2
	viewLenX := 25
	viewLenY := 4

	gameOverView := view{
		gameOverViewName,
		"Game over",
		"Press space to restart",
		viewPositionX,
		viewPositionX + viewLenX,
		viewPositionY,
		viewPositionY + viewLenY}
	return createView(g, gameOverView, false)
}

func gameOver(g *gocui.Gui) error {
	v, err := g.View(gameOverViewName)
	if err != nil {
		return err
	}
	v.Visible = true
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
