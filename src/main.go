package main

import (
	"github.com/awesome-gocui/gocui"
	"log"
	"math/rand"
	"time"
)

const gameViewName = "game"

var (
	game         *gocui.Gui
	r            = rand.New(rand.NewSource(time.Now().UnixNano()))
	running      = true
	gameFinished = false
	tickInterval = 50 * time.Millisecond
)

func main() {
	game = initGameView()
	defer game.Close()

	if err := initKeybindings(game); err != nil {
		log.Panicln(err)
	}

	if err := game.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func initGameView() *gocui.Gui {
	game, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}

	game.Highlight = true
	game.SelFgColor = gocui.ColorRed
	game.SetManagerFunc(manageGame)

	return game
}

func manageGame(game *gocui.Gui) error {
	maxX, maxY := game.Size()

	if err := initKeybindingsView(game); err != nil {
		log.Panicln(err)
	}
	if err := initStatsView(game); err != nil {
		log.Panicln(err)
	}

	if v, err := game.SetView(gameViewName, 0, 0, maxX-26, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			log.Panicln(err)
		}
		v.Title = "Snek"

		if _, err := game.SetViewOnBottom(gameViewName); err != nil {
			log.Panicln(err)
		}

		if err := setViewAtRandom(game, snekHead.viewName, true); err != nil {
			log.Panicln(err)
		}
		if err := setViewAtRandom(game, boxViewName, false); err != nil {
			log.Panicln(err)
		}
		go updateMovement(game)
	}
	if err := initPauseView(game); err != nil {
		log.Panicln(err)
	}
	if err := initGameOverView(game); err != nil {
		log.Panicln(err)
	}
	return nil
}

func updateMovement(game *gocui.Gui) {
	for {
		time.Sleep(tickInterval)
		if !running {
			continue
		}
		game.Update(func(g *gocui.Gui) error {
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
