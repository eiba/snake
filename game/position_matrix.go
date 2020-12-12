package game

func initPositionMatrix(gameViewPosition Position, positionMatrix [][]Position) [][]Position {
	gameViewCols := gameViewPosition.X1 / DeltaX
	gameViewRows := gameViewPosition.Y1 / DeltaY
	if len(positionMatrix) == gameViewCols && len(positionMatrix[0]) == gameViewRows {
		return positionMatrix
	}
	return generatePositionMatrix(gameViewPosition)
}

func generatePositionMatrix(gameViewPosition Position) [][]Position {
	totalCols := gameViewPosition.X1 / DeltaX
	totalRows := gameViewPosition.Y1 / DeltaY
	positionMatrix := make([][]Position, totalCols)

	for col := range positionMatrix {
		positionMatrix[col] = make([]Position, totalRows)
		for row := range positionMatrix[col] {
			x0 := col * DeltaX
			y0 := row * DeltaY
			position := Position{x0, y0, x0 + DeltaX, y0 + DeltaY}
			positionMatrix[col][row] = position
		}
	}
	return positionMatrix
}
