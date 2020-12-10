package main

import "github.com/eiba/snake/game"

var foodView = viewProperties{"food", "", "", game.position{}}

func eatFood() error {
	err := game.addBodyPartToEnd(*game.snakeBodyParts[len(game.snakeBodyParts)-1])
	if err != nil {
		return err
	}

	var foundEmptyPosition bool
	foodView.position, foundEmptyPosition, err = trySetViewAtRandomEmptyPosition(foodView.name, positionMatrix)
	if !foundEmptyPosition {
		return gameOver("Game Won!")
	}
	return err
}
