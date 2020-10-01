package main

func autopilot() error {
	xH0, yH0, _, _, err := gui.ViewPosition(snekHead.viewName)
	if err != nil {
		return err
	}
	xB0, yB0, _, _, err := gui.ViewPosition(boxView.name)
	if err != nil {
		return err
	}

	if xH0 < xB0 && directionIsValid(directions.right) {
		headDirection = directions.right
	}
	if xH0 > xB0 && directionIsValid(directions.left) {
		headDirection = directions.left
	}
	if yH0 < yB0 && directionIsValid(directions.down) {
		headDirection = directions.down
	}
	if yH0 > yB0 && directionIsValid(directions.up) {
		headDirection = directions.up
	}
	for i := 1; i < 100; i++ {
		if validDirection(headDirection) {
			break
		}
		headDirection = getRandomValidDirection(snekHead.currentDirection)
	}
	return nil
}

func validDirection(direction direction) bool {
	positions := make([]position, len(snekBodyParts)-1)
	for i := 1; i < len(snekBodyParts); i++ {
		positions[i-1] = getPositionOfNextMove(snekBodyParts[i-1].currentDirection, snekBodyParts[i-1].position, false)
	}
	nextPosition := getPositionOfNextMove(direction, snekHead.position, true)

	mainViewCollision, _ := mainViewCollision(nextPosition)
	if positionsOverlap(nextPosition, positions) || mainViewCollision {
		return false
	}
	return true
}

func directionIsValid(direction direction) bool {
	if getOppositeDirection(snekHead.currentDirection) == direction {
		return false
	}
	return true
}

func getRandomValidDirection(currentDirection direction) direction {
	oppositeDirection := getOppositeDirection(currentDirection)

	for {
		direction := direction(r.Intn(4))
		if direction != oppositeDirection && direction != headDirection {
			return direction
		}
	}
}
