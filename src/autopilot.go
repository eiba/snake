package main

type node struct {
	direction direction
	position  position
}

var (
	hCycle []node
	cycleIndexMap map[position]int
)

func initPositionMatrix(gameViewPosition position) {
	gameViewCols := gameViewPosition.x1 / deltaX
	gameViewRows := gameViewPosition.y1 / deltaY
	if len(positionMatrix) == gameViewCols && len(positionMatrix[0]) == gameViewRows {
		return
	}
	positionMatrix = generatePositionMatrix(gameViewPosition)
}

func generatePositionMatrix(gameViewPosition position) [][]position {
	totalCols := gameViewPosition.x1 / deltaX
	totalRows := gameViewPosition.y1 / deltaY
	positionMatrix := make([][]position, totalCols)

	for col := range positionMatrix {
		positionMatrix[col] = make([]position, totalRows)
		for row := range positionMatrix[col] {
			x0 := col*deltaX
			y0 := row*deltaY
			position := position{x0, y0, x0+deltaX, y0 + deltaY}
			positionMatrix[col][row] = position
		}
	}
	return positionMatrix
}

func generateHamiltonianCycle(positionMatrix [][]position, snekHead *snekBodyPart) []node {
	vertexGraph := generateVertexGraph(positionMatrix)
	numNodes := len(positionMatrix) * len(positionMatrix[0])
	//log.Panicln(len(positionMatrix),len(positionMatrix[0]))
	startCol, startRow := 0,0//snekHead.position.x0/deltaX, snekHead.position.y0/deltaY
	startPosition := positionMatrix[startCol][startRow]
	directions := getPositionVertices(startCol, startRow, len(positionMatrix), len(positionMatrix[0]))
	var tour []node
	//for i := range directions {
	tour = make([]node, numNodes+1)
	tour[0] = node{directions[0], startPosition}
	tour[numNodes] = tour[0]
	usedPositions := make(map[position]bool)
	usedPositions[startPosition] = true
	tour = hamiltonianCycle(usedPositions, tour, 1, numNodes, vertexGraph, positionMatrix)
	//}
	return tour
}
func generateHamiltonianCycleIndexMap(hamiltonianCycle []node) map[position]int{
	indexMap := make(map[position]int)
	for i := 0; i < len(hamiltonianCycle)-1; i++ {
		indexMap[hamiltonianCycle[i].position] = i
	}
	return indexMap
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

func copyPositionMapMap(positionMap map[position]bool) map[position]bool  {
	positionMapCopy := make(map[position]bool)
	for key,value := range positionMap {
		positionMapCopy[key] = value
	}
	return positionMapCopy
}

func randomValidDirection(col int, row int, vertexGraph [][][]direction) direction {
	return vertexGraph[col][row][r.Intn(len(vertexGraph[col][row]))]
}

func getNextNode(positionMatrix [][]position, vertexGraph [][][]direction, previousNode node) node {
	previousCol, previousRow := previousNode.position.x0/deltaX, previousNode.position.y0/deltaY
	previousDirection := previousNode.direction

	var currentCol, currentRow int
	switch previousDirection {
	case directions.right:
		currentCol = previousCol + 1
		currentRow = previousRow
	case directions.left:
		currentCol = previousCol - 1
		currentRow = previousRow
	case directions.up:
		currentCol = previousCol
		currentRow = currentRow - 1
	case directions.down:
		currentCol = previousCol
		currentRow = currentRow + 1

	}
	direction := randomValidDirection(currentCol, currentRow, vertexGraph)
	position := positionMatrix[currentCol][currentRow]
	return node{direction, position}
}

func generateVertexGraph(positionMatrix [][]position) [][][]direction {
	cols := len(positionMatrix)
	rows := len(positionMatrix[0])
	//numNodes := len(positionMatrix)*len(positionMatrix[0])
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

func autopilot() error {
	xH0, yH0, _, _, err := gui.ViewPosition(snekHead.viewName)
	if err != nil {
		return err
	}
	xB0, yB0, _, _, err := gui.ViewPosition(foodView.name)
	if err != nil {
		return err
	}

	if xH0 < xB0 && directionIsValid(directions.right) {
		headDirection = directions.right
	}
	if xH0 > xB0 && directionIsValid(directions.left) {
		headDirection = directions.left
	}
	if yH0 < yB0 && directionIsValid(directions.down) {
		headDirection = directions.down
	}
	if yH0 > yB0 && directionIsValid(directions.up) {
		headDirection = directions.up
	}
	for i := 1; i < 100; i++ {
		if validDirection(headDirection) {
			break
		}
		headDirection = getRandomValidDirection(snekHead.currentDirection)
	}
	return nil
}

func validDirection(direction direction) bool {
	positions := make([]position, len(snekBodyParts)-1)
	for i := 1; i < len(snekBodyParts); i++ {
		positions[i-1] = getPositionOfNextMove(snekBodyParts[i-1].currentDirection, snekBodyParts[i-1].position, false)
	}

	nextPosition := getPositionOfNextMove(direction, snekHead.position, true)
	if positionsOverlap(nextPosition, positions) || mainViewCollision(nextPosition) {
		return false
	}
	return true
}

func directionIsValid(direction direction) bool {
	if getOppositeDirection(snekHead.currentDirection) == direction {
		return false
	}
	return true
}

func getRandomValidDirection(currentDirection direction) direction {
	oppositeDirection := getOppositeDirection(currentDirection)

	for {
		direction := direction(r.Intn(4))
		if direction != oppositeDirection && direction != headDirection {
			return direction
		}
	}
}
