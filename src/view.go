package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

type viewProperties struct {
	name     string
	title    string
	text     string
	position position
}

func getLenXY(viewName string) (int, int, error) {
	x0, y0, x1, y1, err := gui.ViewPosition(viewName)
	if err != nil {
		return 0, 0, err
	}
	return x1 - x0, y1 - y0, nil
}

func createView(viewProperties viewProperties, visible bool) (*gocui.View, error) {
	view, err := gui.SetView(viewProperties.name, viewProperties.position.x0, viewProperties.position.y0, viewProperties.position.x1, viewProperties.position.y1, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return nil, err
		}

		view.Title = viewProperties.title
		view.Visible = visible
		fmt.Fprintln(view, "\n", viewProperties.text)
	}
	return view, nil
}

func setViewAtRandom(name string, positionMatrix [][]position, setCurrent bool) (position, error) {
	//randomPosition := positionMatrix[r.Intn(len(positionMatrix))][r.Intn(len(positionMatrix[0]))]
	_, randomPosition := tryGetRandomEmptyPosition(positionMatrix)
	if err := setViewPosition(name, randomPosition); err != nil {
		return position{}, err
	}

	if setCurrent {
		if err := setCurrentView(name); err != nil {
			return position{}, err
		}
	}
	return randomPosition, nil
}

func tryGetRandomEmptyPosition(positionMatrix [][]position) (bool, position) {
	randomCol := r.Intn(len(positionMatrix))
	randomRow := r.Intn(len(positionMatrix[0]))
	snekPositionSet := getSnekPositionSet(snekBodyParts)
	foundEmptyPosition, emptyPosition := tryGetEmptyPosition(snekPositionSet, positionMatrix, randomCol, randomRow)
	return foundEmptyPosition, emptyPosition
}

func tryGetEmptyPosition(snekPositionSet map[position]bool, positionMatrix [][]position, randomCol int, randomRow int) (bool, position) {
	foundEmptyPosition, position := lookForEmptyPosition(snekPositionSet, positionMatrix, randomCol, len(positionMatrix), randomRow, len(positionMatrix[0]))
	if !foundEmptyPosition {
		foundEmptyPosition, position = lookForEmptyPosition(snekPositionSet, positionMatrix, 0, randomCol, 0, randomRow)
	}
	return foundEmptyPosition, position
}

func lookForEmptyPosition(snekPositionSet map[position]bool, positionMatrix [][]position, startCol int, endCol int, startRow int, endRow int) (bool, position) {
	for i := startCol; i < endCol; i++ {
		for j := startRow; j < endRow; j++ {
			position := positionMatrix[i][j]
			if !snekPositionSet[position] {
				return true, position
			}
		}
	}
	return false, position{}
}

func setViewPosition(name string, position position) error {
	_, err := gui.SetView(name, position.x0, position.y0, position.x1, position.y1, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	return nil
}

func setCurrentView(name string) error {
	if _, err := gui.SetCurrentView(name); err != nil {
		return err
	}
	return nil
}
