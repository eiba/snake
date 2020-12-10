package a_star

import (
	"container/heap"
	"github.com/eiba/snake/game"
	"math"
)

func AStar(startPosition game.Position, goalPosition game.Position, bodyPositionSet map[game.Position]bool, positionMatrix [][]game.Position) []main.node {
	openSet := make(main.PriorityQueue, 1)
	openSet[0] = &main.PriorityNode{startPosition, 0 + distance(startPosition, goalPosition), 0}
	heap.Init(&openSet)

	cameFrom := make(map[game.Position]game.Position)

	gScore := make(map[game.Position]int)
	gScore[startPosition] = 0

	for openSet.Len() > 0 {
		var current = heap.Pop(&openSet).(*main.PriorityNode)
		if current.position == goalPosition {
			return reconstructPath(cameFrom, current.position)
		}

		for _, neighbour := range getNeighbours(current.position, bodyPositionSet, positionMatrix) {
			tentativeGScore := gScore[current.position] + 1
			if tentativeGScore < getScore(gScore, neighbour) {
				cameFrom[neighbour] = current.position
				gScore[neighbour] = tentativeGScore
				fScore := gScore[neighbour] + distance(neighbour, goalPosition)

				if priorityNode, exist := openSet.Exist(neighbour); exist {
					openSet.update(priorityNode, priorityNode.position, fScore)
				} else {
					heap.Push(&openSet,
						&main.PriorityNode{
							position: neighbour,
							fScore:   fScore,
						})
				}
			}
		}
	}
	return nil
}

func getScore(gScore map[game.Position]int, position game.Position) int {
	if score, exist := gScore[position]; exist {
		return score
	}
	return math.MaxInt32
}

func getNeighbours(currentPosition game.Position, bodyPositionSet map[game.Position]bool, positionMatrix [][]game.Position) []game.Position {
	positionCol := currentPosition.x0 / main.deltaX
	positionRow := currentPosition.y0 / main.deltaY

	var neighbours []game.Position
	if positionCol < len(positionMatrix)-1 {
		neighbour := positionMatrix[positionCol+1][positionRow]
		if !bodyPositionSet[neighbour] {
			neighbours = append(neighbours, neighbour)
		}
	}
	if positionCol > 0 {
		neighbour := positionMatrix[positionCol-1][positionRow]
		if !bodyPositionSet[neighbour] {
			neighbours = append(neighbours, neighbour)
		}
	}
	if positionRow < len(positionMatrix[0])-1 {
		neighbour := positionMatrix[positionCol][positionRow+1]
		if !bodyPositionSet[neighbour] {
			neighbours = append(neighbours, neighbour)
		}
	}
	if positionRow > 0 {
		neighbour := positionMatrix[positionCol][positionRow-1]
		if !bodyPositionSet[neighbour] {
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func reconstructPath(cameFrom map[game.Position]game.Position, current game.Position) []main.node {
	totalPath := []main.node{{position: current}}
	for position, exist := cameFrom[current]; exist; {
		totalPath = append(totalPath, main.node{getDirection(position, totalPath[len(totalPath)-1].position), position})
		position, exist = cameFrom[position]
	}
	reverseArray(totalPath)
	return totalPath
}

func getDirection(currentDirection game.Position, nextDirection game.Position) main.direction {
	currentCol, currentRow := currentDirection.x0/main.deltaX, currentDirection.y0/main.deltaY
	nextCol, nextRow := nextDirection.x0/main.deltaX, nextDirection.y0/main.deltaY

	if currentCol < nextCol {
		return main.directions.right
	}
	if currentCol > nextCol {
		return main.directions.left
	}
	if currentRow < nextRow {
		return main.directions.down
	}
	return main.directions.up
}

func reverseArray(positions []main.node) {
	for i, j := 0, len(positions)-1; i < j; i, j = i+1, j-1 {
		positions[i], positions[j] = positions[j], positions[i]
	}
}

func distance(position1 game.Position, position2 game.Position) int {
	position1Col, position1Row := position1.x0/main.deltaX, position1.y0/main.deltaY
	position2Col, position2Row := position2.x0/main.deltaX, position2.y0/main.deltaY

	return int(math.Abs(float64(position1Col-position2Col)) + math.Abs(float64(position1Row-position2Row)))
}
