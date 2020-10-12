package main

var (
	foodPath  []node
	pathIndex = -1
)

func autopilot3() error {
	if pathIndex == len(foodPath)-1 {
		pathIndex = 0
		snekBodyPartSet := getSnekPositionSet(snekBodyParts)
		foodPath = aStar(snekHead.position, foodView.position, snekBodyPartSet, positionMatrix)
	}
	if len(foodPath) == 0 {
		headDirection = getRandomValidDirection(snekHead.currentDirection)
		return nil
	}
	headDirection = foodPath[pathIndex].direction
	pathIndex++
	return nil
}

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
	/*
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
		validNextPositions := getPositionOfDirection(snekHead.currentDirection, snekHead.position, positionMatrix)*/

	/*for _, nextPosition := range validNextPositions {
		nextPositionCycleIndex := cycleIndexMap[nextPosition.position]
		//bestPathToFoodLength := math.MaxInt32
		if nextPositionCycleIndex > headCycleIndex && nextPositionCycleIndex > tailCycleIndex && nextPositionCycleIndex < foodCycleIndex {
			initiateAStar(foodView.position)
			/*pathToFood := aStar(nextPosition.position, foodView.position, getSnekPositionSet(snekBodyParts), positionMatrix)
			if  len(pathToFood) < bestPathToFoodLength {
				headDirection = nextPosition.direction
				bestPathToFoodLength = len(pathToFood)
			}
		}
	}*/
	return nil
}