package main

import (
	"github.com/awesome-gocui/gocui"
	"log"
	"math/rand"
	"time"
)

var (
	gui              *gocui.Gui
	r                = rand.New(rand.NewSource(time.Now().UnixNano()))
	running          = true
	gameFinished     = false
	autoPilotEnabled = true
	tickInterval     = 50 * time.Millisecond
	gameView         = viewProperties{"game", "Snek", "", position{}}
	positionMatrix   [][]position
)

func main() {
	gui = initGUI()
	defer gui.Close()

	if err := initKeybindings(); err != nil {
		log.Panicln(err)
	}

	if err := gui.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func initGUI() *gocui.Gui {
	gui, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	gui.Highlight = true
	gui.SelFgColor = gocui.ColorRed
	gui.SetManagerFunc(manageGame)
	return gui
}

func initGameView(maxX int, maxY int) (position, error) {
	gameViewPosition := calculateGameViewPosition(maxX, maxY)
	if v, err := gui.SetView(gameView.name, gameViewPosition.x0, gameViewPosition.x0, gameViewPosition.x1, gameViewPosition.y1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return gameViewPosition, err
		}
		v.Title = "Snek"
		if _, err := gui.SetViewOnBottom(gameView.name); err != nil {
			return gameViewPosition, err
		}
		initPositionMatrix(gameViewPosition)
		return gameViewPosition, initGame()
	}
	return gameViewPosition, nil
}

func calculateGameViewPosition(maxX int, maxY int) position {
	defaultPosition := position{0, 0, maxX - 26, maxY - 1}

	if defaultPosition.x1%2 != 0 {
		defaultPosition.x1--
	}
	if (defaultPosition.x1/deltaX)%2 != 0 {
		defaultPosition.x1 = defaultPosition.x1 - deltaX
	}

	if defaultPosition.y1%2 != 0 {
		defaultPosition.y1--
	}
	if (defaultPosition.y1/deltaY)%2 != 0 {
		defaultPosition.y1 = defaultPosition.y1 - deltaY
	}
	return defaultPosition
}

func initGame() error {
	var err error
	snekHead.position, err = setViewAtRandomPosition(snekHead.viewName, positionMatrix, true)
	if err != nil {
		return err
	}
	foodView.position, err = setViewAtRandomPosition(foodView.name, positionMatrix, false)
	if err != nil {
		return err
	}
	go updateMovement()
	return nil
}

func manageGame(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()

	if err := initKeybindingsView(); err != nil {
		log.Panicln(err)
	}

	if err := initStatsView(); err != nil {
		log.Panicln(err)
	}

	var err error
	gameView.position, err = initGameView(maxX, maxY)
	if err != nil {
		log.Panicln(err)
	}

	if err := initPauseView(); err != nil {
		log.Panicln(err)
	}

	if err := initGameOverView(); err != nil {
		log.Panicln(err)
	}

	if err := initLoadingView(); err != nil {
		log.Panicln(err)
	}
	return nil
}

func updateMovement() {
	for {
		initPositionMatrix(gameView.position)
		initHamiltonianCycle(gameView.position)
		time.Sleep(tickInterval)
		if !running {
			continue
		}
		if autoPilotEnabled {
			err := autopilot2()
			if err != nil {
				log.Panicln(err)
			}
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
