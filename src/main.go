package main

import (
	"github.com/awesome-gocui/gocui"
	"log"
	"math/rand"
	"time"
)

const gameViewName = "game"

var (
	gui              *gocui.Gui
	r                = rand.New(rand.NewSource(time.Now().UnixNano()))
	running          = true
	gameFinished     = false
	autoPilotEnabled = false
	tickInterval     = 50 * time.Millisecond
)

func main() {
	gui = initGameView()
	defer gui.Close()

	if err := initKeybindings(); err != nil {
		log.Panicln(err)
	}

	if err := gui.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func initGameView() *gocui.Gui {
	gui, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}

	gui.Highlight = true
	gui.SelFgColor = gocui.ColorRed
	gui.SetManagerFunc(manageGame)

	return gui
}

func manageGame(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()

	if err := initKeybindingsView(); err != nil {
		log.Panicln(err)
	}
	if err := initStatsView(); err != nil {
		log.Panicln(err)
	}

	if v, err := gui.SetView(gameViewName, 0, 0, maxX-26, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			log.Panicln(err)
		}
		v.Title = "Snek"

		if _, err := gui.SetViewOnBottom(gameViewName); err != nil {
			log.Panicln(err)
		}

		if err := setViewAtRandom(snekHead.viewName, true); err != nil {
			log.Panicln(err)
		}
		if err := setViewAtRandom(boxViewName, false); err != nil {
			log.Panicln(err)
		}
		go updateMovement()
	}

	if err := initPauseView(); err != nil {
		log.Panicln(err)
	}
	if err := initGameOverView(); err != nil {
		log.Panicln(err)
	}
	return nil
}

func updateMovement() {
	for {
		time.Sleep(tickInterval)
		if !running {
			continue
		}
		if autoPilotEnabled {
			err := autopilot(); if err != nil {log.Panicln(err)}
		}
		gui.Update(func(gui *gocui.Gui) error {
			if err := moveSnekHead(); err != nil {
				log.Panicln(err)
			}
			if err := moveSnekBodyParts(); err != nil {
				log.Panicln(err)
			}
			return nil
		})
	}
}
