package main

import (
	"math/rand"
	"time"
)

type snekBodyPart struct {
	currentDirection  direction
	previousDirection direction
	viewName          string
	position          position
}

type position struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

type direction int
type movementDirections struct {
	up    direction
	right direction
	down  direction
	left  direction
}

type node struct {
	direction direction
	position  position
}

const (
	deltaX = 1
	deltaY = 1
)

var (
	gameViewPosition = position{0, 0, 100 * deltaX, 100}
	positionMatrix   = generatePositionMatrix(gameViewPosition)
	snekHead         = &snekBodyPart{directions.up, directions.up, "s0", positionMatrix[0][0]}
	directions       = movementDirections{0, 1, 2, 3}
	r                = rand.New(rand.NewSource(time.Now().UnixNano()))
	k                = 0
)

func main() {
	//positionMatrix := generatePositionMatrix(gameViewPosition)*
	//vertexGraph := generateVertexGraph(positionMatrix)

	_ = generateHamiltonianCycle(positionMatrix, snekHead)

	/*directions := getPositionVertices(1,1,10,10)
	directions2 := getPositionVertices(1,1,10,10)
	log.Println(directions)
	shuffleDirections(directions)
	log.Println(directions)
	log.Println(directions2)*/
}

func test(list []int)  {
	list[0] = 1
}
func generatePositionMatrix(gameViewPosition position) [][]position {
	totalCols := gameViewPosition.x1 / deltaX
	totalRows := gameViewPosition.y1 / deltaY
	positionMatrix := make([][]position, totalCols)

	for col := range positionMatrix {
		positionMatrix[col] = make([]position, totalRows)
		for row := range positionMatrix[col] {
			x0 := col * deltaX
			y0 := row * deltaY
			position := position{x0, y0, x0 + deltaX, y0 + deltaY}
			positionMatrix[col][row] = position
		}
	}
	return positionMatrix
}

func generateVertexGraph(positionMatrix [][]position) [][][]direction {
	cols := len(positionMatrix)
	rows := len(positionMatrix[0])
	vertexGraph := make([][][]direction, cols)

	for col := range positionMatrix {
		vertexGraph[col] = make([][]direction, rows)
		for row := range vertexGraph[col] {
			vertexGraph[col][row] = getPositionVertices(col, row, cols, rows)
		}
	}
	return vertexGraph
}

func generateHamiltonianCycle(positionMatrix [][]position, snekHead *snekBodyPart) []node {
	vertexGraph := generateVertexGraph(positionMatrix)
	numNodes := len(positionMatrix) * len(positionMatrix[0])
	startCol, startRow := snekHead.position.x0/deltaX, snekHead.position.y0/deltaY
	startPosition := positionMatrix[startCol][startRow]
	directions := getPositionVertices(startCol, startRow, len(positionMatrix), len(positionMatrix[0]))
	var tour []node
	//for i := range directions {
		tour = make([]node, numNodes+1)
		tour[0] = node{directions[0], startPosition}
		tour[numNodes] = tour[0]
		usedPositions := make(map[position]bool)
		usedPositions[startPosition] = true
		k=0
		tour = hamiltonianCycle(usedPositions, tour, 1, numNodes, vertexGraph, positionMatrix)
	//}
	return tour
}

func hamiltonianCycle(usedPositions map[position]bool, tour []node, moveNumber int, totalMoves int, vertexGraph [][][]direction, positionMatrix [][]position) []node {
	previousNode := tour[moveNumber-1]
	nextCol, nextRow := getNextPosition(previousNode)
	nextPosition := positionMatrix[nextCol][nextRow]

	if usedPositions[nextPosition] {
		return tour
	}

	usedPositionsCopy := copyPositionMapMap(usedPositions)
	usedPositionsCopy[nextPosition] = true
	validVertices := vertexGraph[nextCol][nextRow]

	for _, nextDirection := range validVertices {
		if nextDirection == getOppositeDirection(previousNode.direction) {
			continue
		}
		if isNeighbours(tour[totalMoves-1].position, tour[0].position) {
			return tour
		} else {
			nextNode := node{nextDirection, nextPosition}
			tour[moveNumber] = nextNode
			tour = hamiltonianCycle(usedPositionsCopy, tour, moveNumber+1, totalMoves, vertexGraph, positionMatrix)
		}
	}
	return tour
}

func copyPositionMapMap(positionMap map[position]bool) map[position]bool  {
	positionMapCopy := make(map[position]bool)
	for key,value := range positionMap {
		positionMapCopy[key] = value
	}
	return positionMapCopy
}

func getNextPosition(currentNode node) (int, int) {
	currentCol, currentRow := currentNode.position.x0/deltaX, currentNode.position.y0/deltaY

	switch currentNode.direction {
	case directions.up:
		currentRow--
	case directions.right:
		currentCol++
	case directions.down:
		currentRow++
	case directions.left:
		currentCol--
	}
	return currentCol, currentRow
}

func getOppositeDirection(direction direction) direction {
	return (direction + 2) % 4
}

func getPositionVertices(col int, row int, cols int, rows int) []direction {
	if col == 0 && row == 0 {
		return []direction{directions.right, directions.down}
	}
	if col == 0 && row == rows-1 {
		return []direction{directions.up, directions.right}
	}
	if col == cols-1 && row == 0 {
		return []direction{directions.down, directions.left}
	}
	if col == cols-1 && row == rows-1 {
		return []direction{directions.up, directions.left}
	}
	if col == 0 {
		return []direction{directions.up, directions.right, directions.down}
	}
	if col == cols-1 {
		return []direction{directions.up, directions.down, directions.left}
	}
	if row == 0 {
		return []direction{directions.right, directions.down, directions.left}
	}
	if row == rows-1 {
		return []direction{directions.up, directions.right, directions.left}
	}
	return []direction{directions.up, directions.right, directions.down, directions.left}
}

func shuffleDirections(directions []direction) {
	r.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })
}

func isNeighbours(position1 position, position2 position) bool {
	emptyPosition := position{}
	if position1 == emptyPosition || position2 == emptyPosition {
		return false
	}
	if position1.y0+deltaY == position2.y0 && position1.x0 == position2.x0 {
		return true
	}
	if position1.y0-deltaY == position2.y0 && position1.x0 == position2.x0 {
		return true
	}
	if position1.x0+deltaX == position2.x0 && position1.y0 == position2.y0 {
		return true
	}
	if position1.x0-deltaX == position2.x0 && position1.y0 == position2.y0 {
		return true
	}
	return false
}
