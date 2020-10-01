package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"log"
)

type viewProperties struct {
	name     string
	title    string
	text     string
	position position
}

func getMaxXY(viewName string) (int, int, error) {
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

func setViewAtRandom(name string, setCurrent bool) (position, error) {
	x0, y0, x1, y1, err := gui.ViewPosition(gameViewName)
	if err != nil {
		return position{0,0,0,0}, err
	}

	maxX, maxY := x1-x0-3, y1-y0-2
	positionX, positionY := r.Intn(maxX)+1, r.Intn(maxY)+1
	viewPosition := position{positionX, positionY, positionX + deltaX, positionY + deltaY}

	_, err = gui.SetView(name, viewPosition.x0, viewPosition.y0, viewPosition.x1, viewPosition.y1, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return position{0,0,0,0}, err
	}

	if setCurrent {
		if _, err := gui.SetCurrentView(name); err != nil {
			log.Panicln(err)
		}
	}
	return viewPosition, nil
}
