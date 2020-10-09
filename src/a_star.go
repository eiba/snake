package main

import (
	"container/heap"
	"math"
)

func aStar(startPosition position, goalPosition position, bodyPositionSet map[position]bool, positionMatrix [][]position) []node {
	openSet := make(PriorityQueue, 1)
	openSet[0] = &PriorityNode{startPosition, 0 + distance(startPosition, goalPosition), 0}
	heap.Init(&openSet)

	cameFrom := make(map[position]position)

	gScore := make(map[position]int)
	gScore[startPosition] = 0

	for openSet.Len() > 0 {
		var current = heap.Pop(&openSet).(*PriorityNode)
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
						&PriorityNode{
							position: neighbour,
							fScore:   fScore,
						})
				}
			}
		}
	}
	return nil
}

func getScore(gScore map[position]int, position position) int {
	if score, exist := gScore[position]; exist {
		return score
	}
	return math.MaxInt32
}

func getNeighbours(currentPosition position, bodyPositionSet map[position]bool, positionMatrix [][]position) []position {
	positionCol := currentPosition.x0 / deltaX
	positionRow := currentPosition.y0 / deltaY

	var neighbours []position
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

func reconstructPath(cameFrom map[position]position, current position) []node {
	totalPath := []node{{position: current}}
	for position, exist := cameFrom[current]; exist; {
		totalPath = append(totalPath, node{getDirection(position, totalPath[len(totalPath)-1].position), position})
		position, exist = cameFrom[position]
	}
	reverseArray(totalPath)
	return totalPath
}

func getDirection(currentDirection position, nextDirection position) direction {
	currentCol, currentRow := currentDirection.x0/deltaX, currentDirection.y0/deltaY
	nextCol, nextRow := nextDirection.x0/deltaX, nextDirection.y0/deltaY

	if currentCol < nextCol {
		return directions.right
	}
	if currentCol > nextCol {
		return directions.left
	}
	if currentRow < nextRow {
		return directions.down
	}
	return directions.up
}

func reverseArray(positions []node) {
	for i, j := 0, len(positions)-1; i < j; i, j = i+1, j-1 {
		positions[i], positions[j] = positions[j], positions[i]
	}
}

func distance(position1 position, position2 position) int {
	position1Col, position1Row := position1.x0/deltaX, position1.y0/deltaY
	position2Col, position2Row := position2.x0/deltaX, position2.y0/deltaY

	return int(math.Abs(float64(position1Col-position2Col)) + math.Abs(float64(position1Row-position2Row)))
}
