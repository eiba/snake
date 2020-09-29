package main

import (
	"github.com/awesome-gocui/gocui"
)

const gameOverViewName = "gameOver"

func gameOver(g *gocui.Gui) error {
	running = false

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
	return createView(g, gameOverView)
}
