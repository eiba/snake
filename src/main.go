package main

import (
	"github.com/awesome-gocui/gocui"
	"log"
	"math/rand"
	"time"
)

const gameViewName = "game"

var (
	gui          *gocui.Gui
	r            = rand.New(rand.NewSource(time.Now().UnixNano()))
	running      = true
	gameFinished = false
	tickInterval = 50 * time.Millisecond
)

func main() {
	gui = initGameView()
	defer gui.Close()

	if err := initKeybindings(gui); err != nil {
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

	if err := initKeybindingsView(gui); err != nil {
		log.Panicln(err)
	}
	if err := initStatsView(gui); err != nil {
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

		if err := setViewAtRandom(gui, snekHead.viewName, true); err != nil {
			log.Panicln(err)
		}
		if err := setViewAtRandom(gui, boxViewName, false); err != nil {
			log.Panicln(err)
		}
		go updateMovement(gui)
	}
	if err := initPauseView(gui); err != nil {
		log.Panicln(err)
	}
	if err := initGameOverView(gui); err != nil {
		log.Panicln(err)
	}
	return nil
}

func updateMovement(gui *gocui.Gui) {
	for {
		time.Sleep(tickInterval)
		if !running {
			continue
		}
		gui.Update(func(g *gocui.Gui) error {
			if err := moveSnekHead(g, snekHead); err != nil {
				log.Panicln(err)
			}
			if err := moveSnekBodyParts(g); err != nil {
				log.Panicln(err)
			}
			return nil
		})
	}
}
