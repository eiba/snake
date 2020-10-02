package main


type slot struct {
	taken    bool
	position position
}

func initAutopilot(gameViewPosition position) {
	//gameViewPosition := gameView.position
	//snekHeadPosition := snekHead.position

	_ = getPositionMatrix(gameViewPosition)

	//log.Panicln(matrix, len(matrix),len(matrix[0]))
	//log.Panicln(gameViewPosition.x0,gameViewPosition.y0,gameViewPosition.x1,gameViewPosition.y1,len(positions))

	/*for i := 0; i < len(positions); i++ {
		var positionsWithoutI []position
		for j := 0; j < len(positions); j++ {
			if j==i {
				continue
			}
			positionsWithoutI = append(
				positionsWithoutI,
				positions[j])
		}
		if positionsOverlap(positions[i],positionsWithoutI){
			log.Panicln("overlapping")
		}
	}*/

}
func getPositionMatrix(gameViewPosition position) [][]position {
	totalCols := gameViewPosition.x1/deltaX
	totalRows := gameViewPosition.y1/deltaY
	column := 0
	positions := make([]position, totalCols*totalRows)
	positionMatrix := make([][]position, totalCols)

	for x := 0; x < gameViewPosition.x1; x += deltaX {
		positionMatrix[column] = make([]position, totalRows)
		for row := 0; row < gameViewPosition.y1; row += deltaY {
			position := position{x, row, x + deltaX, row + deltaY}
			positionMatrix[column][row] = position
			positions[(column*totalRows)+row] = position
		}
		column++
	}
	return positionMatrix
}

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
	if positionsOverlap(nextPosition, positions) || mainViewCollision(nextPosition) {
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
