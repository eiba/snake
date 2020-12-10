package main

import (
	"github.com/eiba/snake/a-star"
	"github.com/eiba/snake/game"
)

var (
	foodPath  []node
	pathIndex = -1
)

func initiateAStar(goal game.position) []node {
	foodPath = a_star.AStar(game.snakeHead.position, goal, game.getsnakePositionSet(game.snakeBodyParts), positionMatrix)
	if len(foodPath) == 0 {
		pathIndex = -1
		return foodPath
	}
	pathIndex = 0
	game.headDirection = foodPath[pathIndex].direction
	pathIndex++
	return foodPath
}

func getNextPositionInAStarPath() bool {
	if pathIndex == len(foodPath)-1 {
		return false
	}
	game.headDirection = foodPath[pathIndex].direction
	pathIndex++
	return true
}

func autopilot() error {
	if getNextPositionInAStarPath() {
		return nil
	}
	pathToFood := initiateAStar(game.foodView.position)
	if len(pathToFood) == 0 {
		headPosition := game.snakeHead.position
		headCycleIndex := cycleIndexMap[headPosition]
		headCycleNode := hCycle[headCycleIndex]

		if headCycleNode.direction != game.getOppositeDirection(game.snakeHead.currentDirection) {
			game.headDirection = headCycleNode.direction
		}

		for i := 1; i < 100; i++ {
			if validDirection(game.headDirection) {
				break
			}
			game.headDirection = getRandomValidDirection(game.snakeHead.currentDirection)
		}
	}
	return nil
}

func validDirection(direction game.direction) bool {
	positions := make([]game.position, len(game.snakeBodyParts)-1)
	for i := 1; i < len(game.snakeBodyParts); i++ {
		positions[i-1] = game.getPositionOfNextMove(game.snakeBodyParts[i-1].currentDirection, game.snakeBodyParts[i-1].position, false)
	}

	nextPosition := game.getPositionOfNextMove(direction, game.snakeHead.position, true)
	if game.positionsOverlap(nextPosition, positions) || game.mainViewCollision(nextPosition) {
		return false
	}
	return true
}

func getRandomValidDirection(currentDirection game.direction) game.direction {
	oppositeDirection := game.getOppositeDirection(currentDirection)

	for {
		direction := game.direction(r.Intn(4))
		if direction != oppositeDirection && direction != game.headDirection {
			return direction
		}
	}
}