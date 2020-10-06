package main

import "math"

func aStar(startPosition position, goalPosition position, positionMatrix [][]position){
	openSet := []position{startPosition}

}

func distance(position1 position, position2 position) int  {
	position1Col, position1Row := position1.x0/deltaX, position1.y0/deltaY
	position2Col, position2Row := position2.x0/deltaX, position2.y0/deltaY

	return int(math.Abs(float64(position1Col-position2Col)) + math.Abs(float64(position1Row - position2Row)))
}

