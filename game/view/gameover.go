package view

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game"
)

const gameOverViewName = "gameOver"

var gameOverView *gocui.View

func initGameOverView(gui *gocui.Gui, gameView Properties) error {
	lenX, lenY, err := getLenXY(gui, gameView.Name)
	if err != nil {
		return err
	}

	viewPositionX, viewPositionY := (lenX/2)-12, (lenY/2)-2
	viewLenX := 25
	viewLenY := 4

	gameOverViewProperties := Properties{
		Name: gameOverViewName,
		Text: "Press space to restart",
		Position: game.Position{
			X0: viewPositionX,
			Y0: viewPositionY,
			X1: viewPositionX + viewLenX,
			Y1: viewPositionY + viewLenY}}

	gameOverView, err = createView(gui, gameOverViewProperties, false)
	return err
}

func GameOver(gui *gocui.Gui, title string) error {
	gameOverView.Visible = true
	gameOverView.Title = title
	if _, err := gui.SetCurrentView(gameOverViewName); err != nil {
		return err
	}
	if _, err := gui.SetViewOnTop(gameOverViewName); err != nil {
		return err
	}
	return nil
}
