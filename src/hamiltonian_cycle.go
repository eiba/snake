package main

type node struct {
	direction direction
	position  position
}

var (
	hCycle        []node
	cycleIndexMap map[position]int
)

func initHamiltonianCycle(gameViewPosition position) error {
	gameViewCols := gameViewPosition.x1 / deltaX
	gameViewRows := gameViewPosition.y1 / deltaY
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

func generateHamiltonianCycle(positionMatrix [][]position) []node {
	numNodes := len(positionMatrix) * len(positionMatrix[0])

	startCol, startRow := 0, 0
	startPosition := positionMatrix[startCol][startRow]
	directions := getPositionVertices(startCol, startRow, len(positionMatrix), len(positionMatrix[0]))

	var tour []node
	tour = make([]node, numNodes+1)
	tour[0] = node{directions[0], startPosition}
	tour[numNodes] = tour[0]

	usedPositions := make(map[position]bool)
	usedPositions[startPosition] = true

	tour = hamiltonianCycle(usedPositions, tour, 1, numNodes, generateVertexGraph(positionMatrix), positionMatrix)
	return tour
}

func hamiltonianCycle(usedPositions map[position]bool, tour []node, moveNumber int, totalMoves int, vertexGraph [][][]direction, positionMatrix [][]position) []node {
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

func copyPositionMap(positionMap map[position]bool) map[position]bool {
	positionMapCopy := make(map[position]bool)
	for key, value := range positionMap {
		positionMapCopy[key] = value
	}
	return positionMapCopy
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

func generateHamiltonianCycleIndexMap(hamiltonianCycle []node) map[position]int {
	indexMap := make(map[position]int)
	for i := 0; i < len(hamiltonianCycle)-1; i++ {
		indexMap[hamiltonianCycle[i].position] = i
	}
	return indexMap
}
