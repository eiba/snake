package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

const statsViewName = "stats"

type stat struct {
	name  string
	line  int
	value int
}

var (
	lengthStat  = stat{"Length", 0, 1}
	restartStat = stat{"Restarts", 1, 0}
)

func initStatsView(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if v, err := g.SetView(statsViewName, maxX-25, 8, maxX-1, 11, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Stats"

		fmt.Fprintln(v, fmt.Sprint(lengthStat.name, ":",lengthStat.value))
		fmt.Fprintln(v, fmt.Sprint(restartStat.name, ":",restartStat.value))
	}
	return nil
}

func updateStat(g *gocui.Gui, stat *stat, value int) error {
	v, err := g.View(statsViewName)
	if err != nil {
		return err
	}

	stat.value = value
	if err := v.SetLine(stat.line, fmt.Sprint(stat.name, ":", stat.value)); err != nil {
		return err
	}
	return nil
}
