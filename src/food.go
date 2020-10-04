package main

var foodView = viewProperties{"food", "", "", position{}}

func eatFood() error {
	err := addBodyPartToEnd(*snekBodyParts[len(snekBodyParts)-1])
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