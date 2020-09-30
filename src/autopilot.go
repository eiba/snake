package main

func autopilot() error {
	xH0, yH0, _, _, err := gui.ViewPosition(snekHead.viewName)
	if err != nil {
		return err
	}
	xB0, yB0, _, _, err := gui.ViewPosition(boxViewName)
	if err != nil {
		return err
	}

	if xH0 < xB0 && directionIsValid(directions.right) {
		headDirection = directions.right
	} else if xH0 > xB0 && directionIsValid(directions.left) {
		headDirection = directions.left
	} else if yH0 < yB0 && directionIsValid(directions.down) {
		headDirection = directions.down
	} else if yH0 > yB0 && directionIsValid(directions.up) {
		headDirection = directions.up
	}


	return nil
}

func directionIsValid(direction direction) bool {
	if getOppositeDirection(snekHead.currentDirection) == direction {
		return false
	}
	return true
}
