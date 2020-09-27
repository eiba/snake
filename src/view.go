package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
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

func createView(g *gocui.Gui, view view) error {
	if v, err := g.SetView(view.name, view.x0, view.y0, view.x1, view.y1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}

		v.Title = view.title
		fmt.Fprintln(v, "\n", view.text)

		if _, err := g.SetCurrentView(view.name); err != nil {
			return err
		}
		if _, err := g.SetViewOnTop(view.name); err != nil {
			return err
		}
	}
	return nil
}
