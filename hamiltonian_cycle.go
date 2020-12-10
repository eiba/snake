package main

import "github.com/eiba/snake/game"

type node struct {
	direction game.direction
	position  game.position
}

var (
	hCycle        []node
	cycleIndexMap map[game.position]int
)

func initHamiltonianCycle(gameViewPosition game.position) error {
	gameViewCols := gameViewPosition.x1 / game.deltaX
	gameViewRows := gameViewPosition.y1 / game.deltaY
	if len(hCycle)-1 == gameViewCols*gameViewRows || !autoPilotEnabled {
		return nil
	}

	if err := loading(true); err != nil {
		return err
	}
	hCycle = generateHamiltonianCycle(positionMatrix)
	cycleIndexMap = generateHamiltonianCycleIndexMap(hCycle)
	if err := loading(false); err != nil {
		return err
	}
	return nil
}

func generateHamiltonianCycle(positionMatrix [][]game.position) []node {
	numNodes := len(positionMatrix) * len(positionMatrix[0])

	startCol, startRow := 0, 0
	startPosition := positionMatrix[startCol][startRow]
	directions := getPositionVertices(startCol, startRow, len(positionMatrix), len(positionMatrix[0]))

	var tour []node
	tour = make([]node, numNodes+1)
	tour[0] = node{directions[0], startPosition}
	tour[numNodes] = tour[0]

	usedPositions := make(map[game.position]bool)
	usedPositions[startPosition] = true

	tour = hamiltonianCycle(usedPositions, tour, 1, numNodes, generateVertexGraph(positionMatrix), positionMatrix)
	return tour
}

func hamiltonianCycle(usedPositions map[game.position]bool, tour []node, moveNumber int, totalMoves int, vertexGraph [][][]game.direction, positionMatrix [][]game.position) []node {
	previousNode := tour[moveNumber-1]
	nextCol, nextRow := getNextPosition(previousNode)
	nextPosition := positionMatrix[nextCol][nextRow]

	if usedPositions[nextPosition] {
		return tour
	}

	usedPositionsCopy := copyPositionMap(usedPositions)
	usedPositionsCopy[nextPosition] = true
	validVertices := vertexGraph[nextCol][nextRow]

	for _, nextDirection := range validVertices {
		if nextDirection == game.getOppositeDirection(previousNode.direction) {
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

func generateVertexGraph(positionMatrix [][]game.position) [][][]game.direction {
	cols := len(positionMatrix)
	rows := len(positionMatrix[0])
	vertexGraph := make([][][]game.direction, cols)

	for col := range positionMatrix {
		vertexGraph[col] = make([][]game.direction, rows)
		for row := range vertexGraph[col] {
			vertexGraph[col][row] = getPositionVertices(col, row, cols, rows)
		}
	}
	return vertexGraph
}

func getPositionVertices(col int, row int, cols int, rows int) []game.direction {
	if col == 0 && row == 0 {
		return []game.direction{game.directions.right, game.directions.down}
	}
	if col == 0 && row == rows-1 {
		return []game.direction{game.directions.up, game.directions.right}
	}
	if col == cols-1 && row == 0 {
		return []game.direction{game.directions.down, game.directions.left}
	}
	if col == cols-1 && row == rows-1 {
		return []game.direction{game.directions.up, game.directions.left}
	}
	if col == 0 {
		return []game.direction{game.directions.up, game.directions.right, game.directions.down}
	}
	if col == cols-1 {
		return []game.direction{game.directions.up, game.directions.down, game.directions.left}
	}
	if row == 0 {
		return []game.direction{game.directions.right, game.directions.down, game.directions.left}
	}
	if row == rows-1 {
		return []game.direction{game.directions.up, game.directions.right, game.directions.left}
	}
	return []game.direction{game.directions.up, game.directions.right, game.directions.down, game.directions.left}
}

func getNextPosition(currentNode node) (int, int) {
	currentCol, currentRow := currentNode.position.x0/game.deltaX, currentNode.position.y0/game.deltaY

	switch currentNode.direction {
	case game.directions.up:
		currentRow--
	case game.directions.right:
		currentCol++
	case game.directions.down:
		currentRow++
	case game.directions.left:
		currentCol--
	}
	return currentCol, currentRow
}

func copyPositionMap(positionMap map[game.position]bool) map[game.position]bool {
	positionMapCopy := make(map[game.position]bool)
	for key, value := range positionMap {
		positionMapCopy[key] = value
	}
	return positionMapCopy
}

func isNeighbours(position1 game.position, position2 game.position) bool {
	emptyPosition := game.position{}
	if position1 == emptyPosition || position2 == emptyPosition {
		return false
	}
	if position1.y0+game.deltaY == position2.y0 && position1.x0 == position2.x0 {
		return true
	}
	if position1.y0-game.deltaY == position2.y0 && position1.x0 == position2.x0 {
		return true
	}
	if position1.x0+game.deltaX == position2.x0 && position1.y0 == position2.y0 {
		return true
	}
	if position1.x0-game.deltaX == position2.x0 && position1.y0 == position2.y0 {
		return true
	}
	return false
}

func generateHamiltonianCycleIndexMap(hamiltonianCycle []node) map[game.position]int {
	indexMap := make(map[game.position]int)
	for i := 0; i < len(hamiltonianCycle)-1; i++ {
		indexMap[hamiltonianCycle[i].position] = i
	}
	return indexMap
}
