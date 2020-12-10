package hamiltonian_cycle

import (
	"github.com/eiba/snake/game"
	"github.com/eiba/snake/game/view"
)

type node struct {
	direction game.Direction
	Position  game.Position
}

var (
	hCycle        []node
	cycleIndexMap map[game.Position]int
)

func initHamiltonianCycle(gameViewPosition game.Position, PositionMatrix [][]game.Position, autoPilot bool) error {
	gameViewCols := gameViewPosition.X1 / game.DeltaX
	gameViewRows := gameViewPosition.X1 / game.DeltaY
	if len(hCycle)-1 == gameViewCols*gameViewRows || !autoPilot {
		return nil
	}

	if err := view.Loading(true); err != nil {
		return err
	}
	hCycle = generateHamiltonianCycle(PositionMatrix)
	cycleIndexMap = generateHamiltonianCycleIndexMap(hCycle)
	if err := view.Loading(false); err != nil {
		return err
	}
	return nil
}

func generateHamiltonianCycle(PositionMatrix [][]game.Position) []node {
	numNodes := len(PositionMatrix) * len(PositionMatrix[0])

	startCol, startRow := 0, 0
	startPosition := PositionMatrix[startCol][startRow]
	directions := getPositionVertices(startCol, startRow, len(PositionMatrix), len(PositionMatrix[0]))

	var tour []node
	tour = make([]node, numNodes+1)
	tour[0] = node{directions[0], startPosition}
	tour[numNodes] = tour[0]

	usedPositions := make(map[game.Position]bool)
	usedPositions[startPosition] = true

	tour = hamiltonianCycle(usedPositions, tour, 1, numNodes, generateVertexGraph(PositionMatrix), PositionMatrix)
	return tour
}

func hamiltonianCycle(usedPositions map[game.Position]bool, tour []node, moveNumber int, totalMoves int, vertexGraph [][][]game.Direction, PositionMatrix [][]game.Position) []node {
	previousNode := tour[moveNumber-1]
	nextCol, nextRow := getNextPosition(previousNode)
	nextPosition := PositionMatrix[nextCol][nextRow]

	if usedPositions[nextPosition] {
		return tour
	}

	usedPositionsCopy := copyPositionMap(usedPositions)
	usedPositionsCopy[nextPosition] = true
	validVertices := vertexGraph[nextCol][nextRow]

	for _, nextDirection := range validVertices {
		if nextDirection == game.GetOppositeDirection(previousNode.direction) {
			continue
		}
		if isNeighbours(tour[totalMoves-1].Position, tour[0].Position) {
			return tour
		} else {
			nextNode := node{nextDirection, nextPosition}
			tour[moveNumber] = nextNode
			tour = hamiltonianCycle(usedPositionsCopy, tour, moveNumber+1, totalMoves, vertexGraph, PositionMatrix)
		}
	}
	return tour
}

func generateVertexGraph(PositionMatrix [][]game.Position) [][][]game.Direction {
	cols := len(PositionMatrix)
	rows := len(PositionMatrix[0])
	vertexGraph := make([][][]game.Direction, cols)

	for col := range PositionMatrix {
		vertexGraph[col] = make([][]game.Direction, rows)
		for row := range vertexGraph[col] {
			vertexGraph[col][row] = getPositionVertices(col, row, cols, rows)
		}
	}
	return vertexGraph
}

func getPositionVertices(col int, row int, cols int, rows int) []game.Direction {
	if col == 0 && row == 0 {
		return []game.Direction{game.Directions.Right, game.Directions.Down}
	}
	if col == 0 && row == rows-1 {
		return []game.Direction{game.Directions.Up, game.Directions.Right}
	}
	if col == cols-1 && row == 0 {
		return []game.Direction{game.Directions.Down, game.Directions.Left}
	}
	if col == cols-1 && row == rows-1 {
		return []game.Direction{game.Directions.Up, game.Directions.Left}
	}
	if col == 0 {
		return []game.Direction{game.Directions.Up, game.Directions.Right, game.Directions.Down}
	}
	if col == cols-1 {
		return []game.Direction{game.Directions.Up, game.Directions.Down, game.Directions.Left}
	}
	if row == 0 {
		return []game.Direction{game.Directions.Right, game.Directions.Down, game.Directions.Left}
	}
	if row == rows-1 {
		return []game.Direction{game.Directions.Up, game.Directions.Right, game.Directions.Left}
	}
	return []game.Direction{game.Directions.Up, game.Directions.Right, game.Directions.Down, game.Directions.Left}
}

func getNextPosition(currentNode node) (int, int) {
	currentCol, currentRow := currentNode.Position.X0/game.DeltaX, currentNode.Position.Y0/game.DeltaY

	switch currentNode.direction {
	case game.Directions.Up:
		currentRow--
	case game.Directions.Right:
		currentCol++
	case game.Directions.Down:
		currentRow++
	case game.Directions.Left:
		currentCol--
	}
	return currentCol, currentRow
}

func copyPositionMap(PositionMap map[game.Position]bool) map[game.Position]bool {
	PositionMapCopy := make(map[game.Position]bool)
	for key, value := range PositionMap {
		PositionMapCopy[key] = value
	}
	return PositionMapCopy
}

func isNeighbours(Position1 game.Position, Position2 game.Position) bool {
	emptyPosition := game.Position{}
	if Position1 == emptyPosition || Position2 == emptyPosition {
		return false
	}
	if Position1.Y0+game.DeltaY == Position2.Y0 && Position1.X0 == Position2.X0 {
		return true
	}
	if Position1.Y0-game.DeltaY == Position2.Y0 && Position1.X0 == Position2.X0 {
		return true
	}
	if Position1.X0+game.DeltaX == Position2.X0 && Position1.Y0 == Position2.Y0 {
		return true
	}
	if Position1.X0-game.DeltaX == Position2.X0 && Position1.Y0 == Position2.Y0 {
		return true
	}
	return false
}

func generateHamiltonianCycleIndexMap(hamiltonianCycle []node) map[game.Position]int {
	indexMap := make(map[game.Position]int)
	for i := 0; i < len(hamiltonianCycle)-1; i++ {
		indexMap[hamiltonianCycle[i].Position] = i
	}
	return indexMap
}
