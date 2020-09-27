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

	viewPositionX, viewPositionY := (maxX/2)-5, (maxY/2)-2
	viewLenX := 12
	viewLenY := 4

	gameOverViewText := " u lose"
	gameOverView := view{
		gameOverViewName,
		gameOverViewName,
		gameOverViewText,
		viewPositionX,
		viewPositionX + viewLenX,
		viewPositionY,
		viewPositionY + viewLenY}
	return createView(g, gameOverView)
}
