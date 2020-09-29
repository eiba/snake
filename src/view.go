package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"log"
)

type viewProperties struct {
	name  string
	title string
	text  string
	x0    int
	x1    int
	y0    int
	y1    int
}

func getMaxXY(viewName string) (int, int, error) {
	x0, y0, x1, y1, err := gui.ViewPosition(viewName)
	if err != nil {
		return 0, 0, err
	}
	return x1 - x0, y1 - y0, nil
}

func createView(viewProperties viewProperties, visible bool) (*gocui.View, error) {
	view, err := gui.SetView(viewProperties.name, viewProperties.x0, viewProperties.y0, viewProperties.x1, viewProperties.y1, 0)
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

func setViewAtRandom(name string, setCurrent bool) error {
	x0, y0, x1, y1, err := gui.ViewPosition(gameViewName)
	if err != nil {
		return err
	}

	maxX, maxY := x1-x0-3, y1-y0-2

	positionX, positionY := r.Intn(maxX)+1, r.Intn(maxY)+1

	lenX := 2
	lenY := 1
	_, err = gui.SetView(name, positionX, positionY, positionX+lenX, positionY+lenY, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}

	if setCurrent {
		if _, err := gui.SetCurrentView(name); err != nil {
			log.Panicln(err)
		}
	}
	return nil
}
