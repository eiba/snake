package main

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
			x0 := col * deltaX
			y0 := row * deltaY
			position := position{x0, y0, x0 + deltaX, y0 + deltaY}
			positionMatrix[col][row] = position
		}
	}
	return positionMatrix
}
