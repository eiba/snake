package main

import "github.com/eiba/snake/game"

func initPositionMatrix(gameViewPosition game.position) {
	gameViewCols := gameViewPosition.x1 / game.deltaX
	gameViewRows := gameViewPosition.y1 / game.deltaY
	if len(positionMatrix) == gameViewCols && len(positionMatrix[0]) == gameViewRows {
		return
	}
	positionMatrix = generatePositionMatrix(gameViewPosition)
}

func generatePositionMatrix(gameViewPosition game.position) [][]game.position {
	totalCols := gameViewPosition.x1 / game.deltaX
	totalRows := gameViewPosition.y1 / game.deltaY
	positionMatrix := make([][]game.position, totalCols)

	for col := range positionMatrix {
		positionMatrix[col] = make([]game.position, totalRows)
		for row := range positionMatrix[col] {
			x0 := col * game.deltaX
			y0 := row * game.deltaY
			position := game.position{x0, y0, x0 + game.deltaX, y0 + game.deltaY}
			positionMatrix[col][row] = position
		}
	}
	return positionMatrix
}
