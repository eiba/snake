package game

import "github.com/eiba/snake/game/view"

var foodView = view.Properties{"food", "", "", Position{}}

func eatFood() error {
	err := addBodyPartToEnd(*SnakeBodyParts[len(SnakeBodyParts)-1])
	if err != nil {
		return err
	}

	var foundEmptyPosition bool
	foodView.position, foundEmptyPosition, err = main.trySetViewAtRandomEmptyPosition(foodView.name, main.positionMatrix)
	if !foundEmptyPosition {
		return main.gameOver("Game Won!")
	}
	return err
}
