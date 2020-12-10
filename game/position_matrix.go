package game

import "github.com/eiba/snake"

func initPositionMatrix(gameViewPosition position) {
	gameViewCols := gameViewPosition.x1 / DeltaX
	gameViewRows := gameViewPosition.y1 / DeltaY
	if len(main.positionMatrix) == gameViewCols && len(main.positionMatrix[0]) == gameViewRows {
		return
	}
	main.positionMatrix = generatePositionMatrix(gameViewPosition)
}

func generatePositionMatrix(gameViewPosition position) [][]position {
	totalCols := gameViewPosition.x1 / DeltaX
	totalRows := gameViewPosition.y1 / DeltaY
	positionMatrix := make([][]position, totalCols)

	for col := range positionMatrix {
		positionMatrix[col] = make([]position, totalRows)
		for row := range positionMatrix[col] {
			x0 := col * DeltaX
			y0 := row * DeltaY
			position := position{x0, y0, x0 + DeltaX, y0 + DeltaY}
			positionMatrix[col][row] = position
		}
	}
	return positionMatrix
}
