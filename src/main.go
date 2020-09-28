package main

import (
	"github.com/awesome-gocui/gocui"
	"log"
	"math/rand"
	"time"
)

const gameViewName = "game"

var (
	r            = rand.New(rand.NewSource(time.Now().UnixNano()))
	running      = true
	tickInterval = 50 * time.Millisecond
)

func main() {
	g := initGameView()
	defer g.Close()

	if err := initKeybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func initGameView() *gocui.Gui {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.SetManagerFunc(manageGame)
	return g
}

func manageGame(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := initKeybindingsView(g); err != nil {
		log.Panicln(err)
	}

	if v, err := g.SetView(gameViewName, 0, 0, maxX-26, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			log.Panicln(err)
		}
		v.Title = "Snek"

		if _, err := g.SetViewOnBottom(gameViewName); err != nil {
			log.Panicln(err)
		}
		if err := setViewAtRandom(g, snekBodyParts[0].viewName, true); err != nil {
			log.Panicln(err)
		}
		if err := setViewAtRandom(g, boxViewName, false); err != nil {
			log.Panicln(err)
		}

		go updateMovement(g)
	}
	return nil
}

func updateMovement(g *gocui.Gui) {
	for {
		time.Sleep(tickInterval)
		if !running {
			continue
		}
		g.Update(func(g *gocui.Gui) error {
			if err := moveSnekHead(g, &snekBodyParts[0]); err != nil {
				log.Panicln(err)
			}
			if err := moveSnekBodyParts(g); err != nil {
				log.Panicln(err)
			}
			return nil
		})
	}
}
