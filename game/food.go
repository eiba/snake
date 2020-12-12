package game

import (
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game/view"
)

var foodView = view.Properties{"food", "", "", Position{}}

func eatFood(gui *gocui.Gui, positionMatrix [][]Position) (error, bool) {
	err := addBodyPartToEnd(*SnakeBodyParts[len(SnakeBodyParts)-1])
	if err != nil {
		return err, false
	}

	var foundEmptyPosition bool
	foodView.Position, foundEmptyPosition, err = view.TrySetViewAtRandomEmptyPosition(gui, foodView.Name, positionMatrix)
	if !foundEmptyPosition {
		return view.GameOver(gui, "Game Won!"), false
	}
	return err, true
}
