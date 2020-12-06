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