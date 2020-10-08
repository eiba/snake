package main

import (
	"container/heap"
	"math"
)

func aStar(startPosition position, goalPosition position, snekBodyParts []*snekBodyPart, positionMatrix [][]position) []position {
	openSet := make(PriorityQueue, 1)
	openSet[0] = &PriorityNode{startPosition, 0 + distance(startPosition, goalPosition), 0}
	heap.Init(&openSet)

	cameFrom := make(map[position]position)

	gScore := make(map[position]int)
	gScore[startPosition] = 0

	fScore := make(map[position]int)
	fScore[startPosition] = distance(startPosition, goalPosition)

	for openSet.Len() > 0 {
		var current = heap.Pop(&openSet).(*PriorityNode)

		if current.position == goalPosition {
			return reconstructPath(cameFrom, current.position)
		}
	}
	return nil
}

func reconstructPath(cameFrom map[position]position, current position) []position {
	totalPath := []position{current}
	for position, exist := cameFrom[current]; exist; {
		totalPath = append(totalPath, position)
		position, exist = cameFrom[position]
	}
	return totalPath
}

func distance(position1 position, position2 position) int {
	position1Col, position1Row := position1.x0/deltaX, position1.y0/deltaY
	position2Col, position2Row := position2.x0/deltaX, position2.y0/deltaY

	return int(math.Abs(float64(position1Col-position2Col)) + math.Abs(float64(position1Row-position2Row)))
}
