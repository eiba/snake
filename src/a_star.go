package main

import (
	"container/heap"
	"math"
)

func aStar(startPosition position, goalPosition position, positionMatrix [][]position){
	openSet := make(PriorityQueue, 1)
	openSet[0] = &PriorityNode{startPosition,0+distance(startPosition,goalPosition),0}
	heap.Init(&openSet)

	//cameFrom := make(map[position]position)

	gScore := make(map[position]int)
	gScore[startPosition] = 0

	fScore := make(map[position]int)
	fScore[startPosition] = distance(startPosition,goalPosition)

}

func distance(position1 position, position2 position) int  {
	position1Col, position1Row := position1.x0/deltaX, position1.y0/deltaY
	position2Col, position2Row := position2.x0/deltaX, position2.y0/deltaY

	return int(math.Abs(float64(position1Col-position2Col)) + math.Abs(float64(position1Row - position2Row)))
}

