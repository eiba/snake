package view

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake/game"
	"math/rand"
	"time"
)
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

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

func createView(gui *gocui.Gui, viewProperties Properties, visible bool) (*gocui.View, error) {
	view, err := gui.SetView(viewProperties.Name, viewProperties.Position.X0, viewProperties.Position.Y0, viewProperties.Position.X1, viewProperties.Position.Y1, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return nil, err
		}

		view.Title = viewProperties.Title
		view.Visible = visible
		fmt.Fprintln(view, "\n", viewProperties.Text)
	}
	return view, nil
}

func setViewPosition(gui *gocui.Gui, name string, position game.Position) error {
	_, err := gui.SetView(name, position.X0, position.Y0, position.X1, position.Y1, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	return nil
}

func setCurrentView(gui *gocui.Gui,name string) error {
	if _, err := gui.SetCurrentView(name); err != nil {
		return err
	}
	return nil
}

func SetViewAtRandomPosition(gui *gocui.Gui, name string, positionMatrix [][]game.Position, setCurrent bool) (game.Position, error) {
	randomPosition := getRandomPosition(positionMatrix)
	if err := setViewPosition(gui, name, randomPosition); err != nil {
		return game.Position{}, err
	}

	if setCurrent {
		if err := setCurrentView(gui, name); err != nil {
			return game.Position{}, err
		}
	}
	return randomPosition, nil
}

func getRandomPosition(positionMatrix [][]game.Position) game.Position {
	return positionMatrix[r.Intn(len(positionMatrix))][r.Intn(len(positionMatrix[0]))]
}

func TrySetViewAtRandomEmptyPosition(gui *gocui.Gui, name string, positionMatrix [][]game.Position) (game.Position, bool, error) {
	randomPosition, foundEmptyPosition := tryGetRandomEmptyPosition(positionMatrix)
	if !foundEmptyPosition {
		return randomPosition, foundEmptyPosition, nil
	}
	if err := setViewPosition(gui, name, randomPosition); err != nil {
		return game.Position{}, foundEmptyPosition, err
	}
	return randomPosition, foundEmptyPosition, nil
}

func tryGetRandomEmptyPosition(positionMatrix [][]game.Position) (game.Position, bool) {
	randomCol := r.Intn(len(positionMatrix))
	randomRow := r.Intn(len(positionMatrix[0]))
	snakePositionSet := game.GetsnakePositionSet(game.SnakeBodyParts)
	emptyPosition, foundEmptyPosition := tryGetEmptyPosition(snakePositionSet, positionMatrix, randomCol, randomRow)
	return emptyPosition, foundEmptyPosition
}

func tryGetEmptyPosition(snakePositionSet map[game.Position]bool, positionMatrix [][]game.Position, randomCol int, randomRow int) (game.Position, bool) {
	position, foundEmptyPosition := lookForEmptyPosition(snakePositionSet, positionMatrix, randomCol, len(positionMatrix), randomRow, len(positionMatrix[0]))
	if !foundEmptyPosition {
		position, foundEmptyPosition = lookForEmptyPosition(snakePositionSet, positionMatrix, 0, randomCol, 0, randomRow)
	}
	return position, foundEmptyPosition
}

func lookForEmptyPosition(snakePositionSet map[game.Position]bool, positionMatrix [][]game.Position, startCol int, endCol int, startRow int, endRow int) (game.Position, bool) {
	for i := startCol; i < endCol; i++ {
		for j := startRow; j < endRow; j++ {
			position := positionMatrix[i][j]
			if !snakePositionSet[position] {
				return position, true
			}
		}
	}
	return game.Position{}, false
}
