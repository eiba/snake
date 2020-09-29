package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"log"
)

type view struct {
	name  string
	title string
	text  string
	x0    int
	x1    int
	y0    int
	y1    int
}

func getMaxXY(g *gocui.Gui, viewName string) (int, int, error) {
	x0, y0, x1, y1, err := g.ViewPosition(viewName)
	if err != nil {
		return 0, 0, err
	}
	return x1 - x0, y1 - y0, nil
}

func createView(g *gocui.Gui, view view, visible bool) error {
	if v, err := g.SetView(view.name, view.x0, view.y0, view.x1, view.y1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}

		v.Title = view.title
		v.Visible = visible
		fmt.Fprintln(v, "\n", view.text)
	}
	return nil
}

func setViewAtRandom(g *gocui.Gui, name string, setCurrent bool) error {
	x0, y0, x1, y1, err := g.ViewPosition(gameViewName)
	if err != nil {
		return err
	}

	maxX, maxY := x1-x0-3, y1-y0-2

	positionX, positionY := r.Intn(maxX)+1, r.Intn(maxY)+1

	lenX := 2
	lenY := 1
	_, err = g.SetView(name, positionX, positionY, positionX+lenX, positionY+lenY, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
	}

	if setCurrent {
		if _, err := g.SetCurrentView(name); err != nil {
			log.Panicln(err)
		}
	}
	return nil
}