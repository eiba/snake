package view

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

const statsViewName = "stats"

var statsView *gocui.View

type stat struct {
	name  string
	line  int
	Value int
}

var (
	LengthStat  = stat{"Length", 0, 1}
	RestartStat = stat{"Restarts", 1, 0}
)

func initStatsView(gui *gocui.Gui, gameView Properties) error {
	maxX  := gameView.Position.X1

	var err error
	statsView, err = gui.SetView(statsViewName, maxX+1, 9, maxX+26, 12, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		statsView.Title = "Stats"

		fmt.Fprintln(statsView, fmt.Sprint(LengthStat.name, ":", LengthStat.Value))
		fmt.Fprintln(statsView, fmt.Sprint(RestartStat.name, ":", RestartStat.Value))
	}
	return nil
}

func UpdateStat(stat *stat, value int) error {
	stat.Value = value
	if err := statsView.SetLine(stat.line, fmt.Sprint(stat.name, ":", stat.Value)); err != nil {
		return err
	}
	return nil
}