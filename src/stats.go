package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

const statsViewName = "stats"

var statsView *gocui.View

type stat struct {
	name  string
	line  int
	value int
}

var (
	lengthStat  = stat{"Length", 0, 1}
	restartStat = stat{"Restarts", 1, 0}
)

func initStatsView() error {
	maxX, _ := gui.Size()

	var err error
	statsView, err = gui.SetView(statsViewName, maxX-25, 9, maxX-1, 12, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		statsView.Title = "Stats"

		fmt.Fprintln(statsView, fmt.Sprint(lengthStat.name, ":", lengthStat.value))
		fmt.Fprintln(statsView, fmt.Sprint(restartStat.name, ":", restartStat.value))
	}
	return nil
}

func updateStat(stat *stat, value int) error {
	stat.value = value
	if err := statsView.SetLine(stat.line, fmt.Sprint(stat.name, ":", stat.value)); err != nil {
		return err
	}
	return nil
}
