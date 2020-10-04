package main

type node struct {
	direction direction
	position  position
}

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
	//column := 0
	positionMatrix := make([][]position, totalCols)
	//positions := make([]position, totalCols*totalRows)
	//positionSet := make(map[position]bool)

	for col := range positionMatrix {
		positionMatrix[col] = make([]position, totalRows)
		for row := range positionMatrix[col] {
			x0 := col*deltaX
			y0 := row*deltaY
			position := position{x0, y0, x0+deltaX, y0 + deltaY}
			positionMatrix[col][row] = position
		}
	}
	/*for x := 0; x < totalCols; x += 1 {
		//log.Panicln(gameViewPosition, totalCols, totalRows)
		positionMatrix[x] = make([]position, totalRows)
		for row := 0; row < gameViewPosition.y1; row += deltaY {
			position := position{x*2, row, (x*2) + deltaX, row + deltaY}

			positionMatrix[x][row] = position
			//positions[(column*totalRows)+row] = position
			//positionSet[position] = true
		}
		//column++
	}*/
	return positionMatrix
}

func generateHamiltonianCycle(positionMatrix [][]position, snekHead *snekBodyPart) []node {
	vertexGraph := generateVertexGraph(positionMatrix)
	numNodes := len(positionMatrix) * len(positionMatrix[0])
	tour := make([]node, numNodes+1)

	/*if positionMatrix[startCol][startRow] == snekHead.position{
		log.Panicln("s",positionMatrix[startCol][startRow],snekHead.position)
	}
	log.Panicln("s",positionMatrix[startCol][startRow],snekHead.position)*/
	emptyNode := node{
		direction: 0,
		position:  position{},
	}
	//for /*!isNeighbours(tour[len(tour)-1].position,tour[0].position) &&*/ tour[len(tour)-1] == emptyNode {
		//snekHead.position, _ = setViewAtRandomPosition(snekHead.viewName, positionMatrix, true)
		startCol, startRow := snekHead.position.x0/deltaX, snekHead.position.y0/deltaY
		startPosition := positionMatrix[startCol][startRow]
		directions := getPositionVertices(startCol,startRow,len(positionMatrix),len(positionMatrix[0]))

		for i := range directions {
			tour[0] = node{directions[i], startPosition}
			usedPositions := make(map[position]bool)
			usedPositions[startPosition] = true
			tour = h(usedPositions, tour, 1, numNodes, vertexGraph, positionMatrix)

			if /*isNeighbours(tour[len(tour)-1].position,tour[0].position) &&*/ tour[len(tour)-1] != emptyNode {
				break
			}
		}
	//}

	//for tour[numNodes-1] == emptyNode {


	//}
	/*tour[0] = node{randomValidDirection(startCol, startRow, vertexGraph), startPosition}
	usedPositions := make(map[position]bool)
	usedPositions[startPosition] = true*/

	/*tour := h(usedPositions,tour,1,numNodes,vertexGraph,positionMatrix)

	addedNodes := 1
	for addedNodes < numNodes {
		tour[addedNodes] = getNextNode(positionMatrix, vertexGraph, tour[addedNodes-1])
		addedNodes++
	}*/

	/*for i := range vertexGraph {
		for j := range vertexGraph[i] {
			vertices := vertexGraph[i][j]
		}
	}
	tour[0] = snekHead.currentDirection

	for i := 1; i < numNodes; i++ {
		tour[i] = getNextDirection(tour[i-1])
	}*/
	return tour
}

func isNeighbours(position1 position, position2 position)bool  {
	if position1.y0+deltaY == position2.y0 && position1.x0 == position2.x0 {
		return true
	}
	if position1.y0-deltaY == position2.y0 && position1.x0 == position2.x0 {
		return true
	}
	if position1.x0+deltaX == position2.x0 && position1.y0 == position2.y0{
		return true
	}
	if position1.x0-deltaX == position2.x0 && position1.y0 == position2.y0{
		return true
	}
	return false
}

func h(usedPositions map[position]bool, tour []node, moveNumber int, totalMoves int, vertexGraph [][][]direction, positionMatrix [][]position) []node {
	if moveNumber == totalMoves && isNeighbours(tour[totalMoves-1].position,tour[0].position){
		tour[totalMoves] = tour[0]
		return tour
	}

	previousNode := tour[moveNumber-1]
	nextCol, nextRow := getNextPosition(previousNode)
	nextPosition := positionMatrix[nextCol][nextRow]

	if usedPositions[nextPosition] {
		//tour[moveNumber-1] =
		//delete(usedPositions, previousNode.position)
		//return h(usedPositions, tour, moveNumber-1, totalMoves, vertexGraph, positionMatrix)
		return tour
	}
	usedPositions[nextPosition] = true
	for _, nextDirection := range vertexGraph[nextCol][nextRow] {
		if nextDirection == getOppositeDirection(previousNode.direction) {
			continue
		}
		nextNode := node{nextDirection, nextPosition}
		tour[moveNumber] = nextNode
		h(usedPositions, tour, moveNumber+1, totalMoves, vertexGraph, positionMatrix)
	}
	return tour
}

func getNextPosition(currentNode node) (int, int) {
	currentCol, currentRow := currentNode.position.x0/deltaX, currentNode.position.y0/deltaY

	var nextCol, nextRow int
	switch currentNode.direction {
	case directions.right:
		nextCol = currentCol + 1
		nextRow = currentRow
	case directions.left:
		nextCol = currentCol - 1
		nextRow = currentRow
	case directions.up:
		nextCol = currentCol
		nextRow = currentRow - 1
	case directions.down:
		nextCol = currentCol
		nextRow = currentRow + 1
	}
	return nextCol, nextRow
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
