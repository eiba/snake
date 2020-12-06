package main

var (
	foodPath  []node
	pathIndex = -1
)

func initiateAStar(goal position) []node {
	foodPath = aStar(snekHead.position, goal, getSnekPositionSet(snekBodyParts), positionMatrix)
	if len(foodPath) == 0 {
		pathIndex = -1
		return foodPath
	}
	pathIndex = 0
	headDirection = foodPath[pathIndex].direction
	pathIndex++
	return foodPath
}

func getNextPositionInAStarPath() bool {
	if pathIndex == len(foodPath)-1 {
		return false
	}
	headDirection = foodPath[pathIndex].direction
	pathIndex++
	return true
}

func autopilot4() error {
	if getNextPositionInAStarPath() {
		return nil
	}
	pathToFood := initiateAStar(foodView.position)
	if len(pathToFood) == 0 {
		headPosition := snekHead.position
		headCycleIndex := cycleIndexMap[headPosition]
		headCycleNode := hCycle[headCycleIndex]

		if headCycleNode.direction != getOppositeDirection(snekHead.currentDirection) {
			headDirection = headCycleNode.direction
		}

		for i := 1; i < 100; i++ {
			if validDirection(headDirection) {
				break
			}
			headDirection = getRandomValidDirection(snekHead.currentDirection)
		}
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

func getRandomValidDirection(currentDirection direction) direction {
	oppositeDirection := getOppositeDirection(currentDirection)

	for {
		direction := direction(r.Intn(4))
		if direction != oppositeDirection && direction != headDirection {
			return direction
		}
	}
}