package main

func randomValidDirection(col int, row int, vertexGraph [][][]direction) direction {
	return vertexGraph[col][row][r.Intn(len(vertexGraph[col][row]))]
}

func getNextNode(positionMatrix [][]position, vertexGraph [][][]direction, previousNode node) node {
	previousCol, previousRow := previousNode.position.x0/deltaX, previousNode.position.y0/deltaY
	previousDirection := previousNode.direction

	var currentCol, currentRow int
	switch previousDirection {
	case directions.right:
		currentCol = previousCol + 1
		currentRow = previousRow
	case directions.left:
		currentCol = previousCol - 1
		currentRow = previousRow
	case directions.up:
		currentCol = previousCol
		currentRow = currentRow - 1
	case directions.down:
		currentCol = previousCol
		currentRow = currentRow + 1

	}
	direction := randomValidDirection(currentCol, currentRow, vertexGraph)
	position := positionMatrix[currentCol][currentRow]
	return node{direction, position}
}

func autopilot2() error {
	headPosition := snekHead.position
	headCycleIndex := cycleIndexMap[headPosition]
	headCycleNode := hCycle[headCycleIndex]

	if headCycleNode.direction != getOppositeDirection(snekHead.currentDirection) {
		headDirection = headCycleNode.direction
	}

	foodPosition := foodView.position
	foodCycleIndex := cycleIndexMap[foodPosition]

	tailPosition := snekBodyParts[len(snekBodyParts)-1].position
	tailCycleIndex := cycleIndexMap[tailPosition]

	validNextPositions := getPositionOfDirection(snekHead.currentDirection, snekHead.position, positionMatrix)

	for _, nextPosition := range validNextPositions {
		nextPositionCycleIndex := cycleIndexMap[nextPosition.position]
		highestValidIndex := headCycleIndex
		if nextPositionCycleIndex > headCycleIndex && nextPositionCycleIndex > tailCycleIndex && nextPositionCycleIndex > highestValidIndex && nextPositionCycleIndex < foodCycleIndex {
			highestValidIndex = nextPositionCycleIndex
			headDirection = nextPosition.direction
		} else if headCycleIndex > foodCycleIndex && nextPositionCycleIndex < foodCycleIndex && nextPositionCycleIndex < tailCycleIndex && validDirection(nextPosition.direction) {
			headDirection = nextPosition.direction
			break
		}
	}
	if !validDirection(headDirection) {
		headDirection = getRandomValidDirection(snekHead.currentDirection)
	}

	return nil
}

func getPositionOfDirection(currentDirection direction, currentPosition position, positionMatrix [][]position) []node {
	currentCol, currentRow := currentPosition.x0/deltaX, currentPosition.y0/deltaY

	positionVetrices := getPositionVertices(currentCol, currentRow, len(positionMatrix), len(positionMatrix[0]))

	var possibleNextPositions []node
	for _, possibleDirection := range positionVetrices {
		if possibleDirection == getOppositeDirection(currentDirection) {
			continue
		}
		nextCol, nextRow := currentCol, currentRow
		switch possibleDirection {
		case directions.up:
			nextRow = currentRow - 1
		case directions.right:
			nextCol = currentCol + 1
		case directions.down:
			nextRow = currentRow + 1
		case directions.left:
			nextCol = currentCol - 1
		}
		possibleNextPositions = append(
			possibleNextPositions,
			node{possibleDirection, positionMatrix[nextCol][nextRow]})
	}
	return possibleNextPositions
}

func autopilot() error {
	xH0, yH0, _, _, err := gui.ViewPosition(snekHead.viewName)
	if err != nil {
		return err
	}
	xB0, yB0, _, _, err := gui.ViewPosition(foodView.name)
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
