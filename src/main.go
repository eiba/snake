package main

import (
	"github.com/awesome-gocui/gocui"
	"log"
	"math/rand"
	"time"
)

type snekBodyPart struct {
	currentDirection  direction
	previousDirection direction
	viewName          string
}

type direction int
type movementDirections struct {
	up    direction
	right direction
	down  direction
	left  direction
}

const (
	deltaX                    = 2
	deltaY                    = 1
	gameViewName, boxViewName = "game", "box"
)

var (
	r             = rand.New(rand.NewSource(time.Now().UnixNano()))
	directions    = movementDirections{0, 1, 2, 3}
	headDirection = direction(r.Intn(4))
	snekBodyParts = []snekBodyPart{{headDirection, headDirection, "s0"}}
	running       = true
	tickInterval  = 50 * time.Millisecond
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.SetManagerFunc(manageGame)

	if err := initKeybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func manageGame(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := initKeybindingsView(g); err != nil {return err}

	if v, err := g.SetView(gameViewName, 0, 0, maxX-26, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}

		if _, err := g.SetViewOnBottom(gameViewName); err != nil {
			return err
		}

		if err := setViewAtRandom(g, snekBodyParts[0].viewName, true); err != nil {
			log.Panicln(err)
		}

		go updateMovement(g)

		if err := setViewAtRandom(g, boxViewName, false); err != nil {
			log.Panicln(err)
		}
		v.Title = "Snek"
	}
	return nil
}

func updateMovement(g *gocui.Gui) error {
	for {
		time.Sleep(tickInterval)
		if !running {
			continue
		}
		g.Update(func(g *gocui.Gui) error {
			err := moveSnekHead(g, &snekBodyParts[0]); if err != nil {
				return err
			}
			return moveSnekBodyParts(g)
		})
	}
}