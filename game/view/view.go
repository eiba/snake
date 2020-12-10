package view

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake"
	"github.com/eiba/snake/game"
)

type Properties struct {
	Name     string
	Title    string
	Text     string
	Position game.Position
}

func getLenXY(gui *gocui.Gui, viewName string) (int, int, error) {
	x0, y0, x1, y1, err := gui.ViewPosition(viewName)
	if err != nil {
		return 0, 0, err
	}
	return x1 - x0, y1 - y0, nil
}

func createView(viewProperties Properties, visible bool) (*gocui.View, error) {
	view, err := main.gui.SetView(viewProperties.Name, viewProperties.Position.x0, viewProperties.position.y0, viewProperties.position.x1, viewProperties.position.y1, 0)
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

func setViewPosition(name string, position game.position) error {
	_, err := main.gui.SetView(name, position.x0, position.y0, position.x1, position.y1, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	return nil
}

func setCurrentView(name string) error {
	if _, err := main.gui.SetCurrentView(name); err != nil {
		return err
	}
	return nil
}

func setViewAtRandomPosition(name string, positionMatrix [][]game.position, setCurrent bool) (game.position, error) {
	randomPosition := getRandomPosition(positionMatrix)
	if err := setViewPosition(name, randomPosition); err != nil {
		return game.position{}, err
	}

	if setCurrent {
		if err := setCurrentView(name); err != nil {
			return game.position{}, err
		}
	}
	return randomPosition, nil
}

func getRandomPosition(positionMatrix [][]game.position) game.position {
	return positionMatrix[main.r.Intn(len(positionMatrix))][main.r.Intn(len(positionMatrix[0]))]
}

func trySetViewAtRandomEmptyPosition(name string, positionMatrix [][]game.position) (game.position, bool, error) {
	randomPosition, foundEmptyPosition := tryGetRandomEmptyPosition(positionMatrix)
	if !foundEmptyPosition {
		return randomPosition, foundEmptyPosition, nil
	}
	if err := setViewPosition(name, randomPosition); err != nil {
		return game.position{}, foundEmptyPosition, err
	}
	return randomPosition, foundEmptyPosition, nil
}

func tryGetRandomEmptyPosition(positionMatrix [][]game.position) (game.position, bool) {
	randomCol := main.r.Intn(len(positionMatrix))
	randomRow := main.r.Intn(len(positionMatrix[0]))
	snakePositionSet := game.getsnakePositionSet(game.snakeBodyParts)
	emptyPosition, foundEmptyPosition := tryGetEmptyPosition(snakePositionSet, positionMatrix, randomCol, randomRow)
	return emptyPosition, foundEmptyPosition
}

func tryGetEmptyPosition(snakePositionSet map[game.position]bool, positionMatrix [][]game.position, randomCol int, randomRow int) (game.position, bool) {
	position, foundEmptyPosition := lookForEmptyPosition(snakePositionSet, positionMatrix, randomCol, len(positionMatrix), randomRow, len(positionMatrix[0]))
	if !foundEmptyPosition {
		position, foundEmptyPosition = lookForEmptyPosition(snakePositionSet, positionMatrix, 0, randomCol, 0, randomRow)
	}
	return position, foundEmptyPosition
}

func lookForEmptyPosition(snakePositionSet map[game.position]bool, positionMatrix [][]game.position, startCol int, endCol int, startRow int, endRow int) (game.position, bool) {
	for i := startCol; i < endCol; i++ {
		for j := startRow; j < endRow; j++ {
			position := positionMatrix[i][j]
			if !snakePositionSet[position] {
				return position, true
			}
		}
	}
	return game.position{}, false
}
