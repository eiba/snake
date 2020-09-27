package main

import (
	"fmt"
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
	deltaX            = 2
	deltaY            = 1
	gameView, boxView = "game", "box"
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
	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func getOppositeDirection(direction direction) direction {
	return (direction + 2) % 4
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := initKeybindingsView(g); err != nil {return err}

	if v, err := g.SetView(gameView, 0, 0, maxX-26, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}

		if _, err := g.SetViewOnBottom(gameView); err != nil {
			return err
		}

		if err := setViewAtRandom(g, snekBodyParts[0].viewName, true); err != nil {
			log.Panicln(err)
		}

		go updateMovement(g)

		if err := setViewAtRandom(g, boxView, false); err != nil {
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

			snekBodyParts[0].previousDirection = snekBodyParts[0].currentDirection
			snekBodyParts[0].currentDirection = headDirection
			err := moveSnekHead(g, snekBodyParts[0])
			if err != nil {
				return err
			}
			for i := 1; i < len(snekBodyParts); i++ {
				currentSnekBodyPart := snekBodyParts[i]
				previousSnekBodyPart := snekBodyParts[i-1]
				previousSnekBodyPartPreviousDirection := snekBodyParts[i-1].previousDirection
				err := moveSnekBodyPart(g, previousSnekBodyPart, currentSnekBodyPart)
				if err != nil {
					return err
				}
				snekBodyParts[i].previousDirection = snekBodyParts[i].currentDirection
				snekBodyParts[i].currentDirection = previousSnekBodyPartPreviousDirection
			}
			return nil
		})
	}
}

func reset(g *gocui.Gui) error {
	headDirection = direction(r.Intn(4))
	running = true
	tickInterval = 50 * time.Millisecond

	for i := 1; i < len(snekBodyParts); i++ {
		if err := g.DeleteView(snekBodyParts[i].viewName); err != nil && !gocui.IsUnknownView(err) {
			return err
		}
	}
	snekBodyParts = []snekBodyPart{{headDirection, headDirection, "s0"}}

	if err := setViewAtRandom(g, snekBodyParts[0].viewName, true); err != nil {
		return err
	}
	if err := setViewAtRandom(g, boxView, false); err != nil {
		return err
	}
	if err := g.DeleteView("gameOver"); err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	return nil
}

func gameOver(g *gocui.Gui) error {
	running = false
	x0, y0, x1, y1, err := g.ViewPosition(gameView)
	if err != nil {
		return err
	}
	maxX, maxY := x1-x0, y1-y0

	positionX, positionY := (maxX/2)-5, (maxY/2)-2

	lenX := 12
	lenY := 4
	name := "gameOver"
	if v, err := g.SetView(name, positionX, positionY, positionX+lenX, positionY+lenY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}

		v.Title = "game over"
		fmt.Fprintln(v, "\n  u lose")

		if _, err := g.SetCurrentView(name); err != nil {
			return err
		}
		if _, err := g.SetViewOnTop(name); err != nil {
			return err
		}
	}
	return nil
}

func pause(g *gocui.Gui) error {
	if running{
		return pauseGame(g)
	}
	return resumeGame(g)
}

func pauseGame(g *gocui.Gui) error {
	running = false
	x0, y0, x1, y1, err := g.ViewPosition(gameView)
	if err != nil {
		return err
	}
	maxX, maxY := x1-x0, y1-y0

	positionX, positionY := (maxX/2)-10, (maxY/2)-2

	lenX := 20
	lenY := 4
	name := "pause"
	if v, err := g.SetView(name, positionX, positionY, positionX+lenX, positionY+lenY, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}

		v.Title = "pause"
		fmt.Fprintln(v, "\n press p to resume")

		if _, err := g.SetCurrentView(name); err != nil {
			return err
		}
		if _, err := g.SetViewOnTop(name); err != nil {
			return err
		}
	}
	return nil
}

func resumeGame(g *gocui.Gui) error  {
	if err := g.DeleteView("pause"); err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	running = true
	return nil
}

func setViewAtRandom(g *gocui.Gui, name string, setCurrent bool) error {
	x0, y0, x1, y1, err := g.ViewPosition(gameView)
	if err != nil {
		return err
	}

	maxX, maxY := x1-x0-3, y1-y0-2

	positionX, positionY := r.Intn(maxX)+1, r.Intn(maxY)+1

	lenX := 2
	lenY := 1
	_, err = g.SetView(name, positionX, positionY, positionX+lenX, positionY+lenY, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
	}

	if setCurrent {
		if _, err := g.SetCurrentView(name); err != nil {
			log.Panicln(err)
		}
	}
	return nil
}

func addBodyPartToEnd(g *gocui.Gui, currentLastSnekBodyPart snekBodyPart) error {
	x0, y0, x1, y1, err := g.ViewPosition(currentLastSnekBodyPart.viewName); if err != nil {return err}
	offsetX, offsetY := calculateOffsets(currentLastSnekBodyPart)

	name := fmt.Sprintf("s%v", len(snekBodyParts))
	_, err = g.SetView(name, x0+offsetX, y0+offsetY, x1+offsetX, y1+offsetY, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
	}

	snekBodyParts = append(
		snekBodyParts,
		snekBodyPart{
			currentLastSnekBodyPart.currentDirection,
			currentLastSnekBodyPart.previousDirection,
			name})
	return nil
}

//Checks collision between view1 and view2, returning true for collision and false otherwise.
func checkViewCollision(g *gocui.Gui, view1 string, view2 string) (bool, error) {
	x10, y10, x11, y11, err := g.ViewPosition(view1); if err != nil {return false, err}
	x20, y20, x21, y21, err := g.ViewPosition(view2); if err != nil {return false, err}

	Ax, Ay, Aw, Ah := x10, y10, x11-x10, y11-y10
	Bx, By, Bw, Bh := x20, y20, x21-x20, y21-y20

	if Bx+Bw > Ax &&
		By+Bh > Ay &&
		Ax+Aw > Bx &&
		Ay+Ah > By {return true, nil}
	return false, nil
}

func moveSnekHead(g *gocui.Gui, snekBodyPart snekBodyPart) error {
	err := moveHeadView(g, snekBodyPart); if err != nil {return err}

	headToMainViewCollision, err := checkHeadToMainViewCollision(g, snekBodyPart); if err != nil {return err}
	if headToMainViewCollision {
		return gameOver(g)
	}

	headToBodyCollision, err := checkHeadToBodyCollision(g); if err != nil {return err}
	if headToBodyCollision {
		return gameOver(g)
	}

	headToBoxCollision, err := checkViewCollision(g, snekBodyParts[0].viewName, boxView); if err != nil {return err}
	if headToBoxCollision {
		return collideWithBox(g)
	}
	return nil
}

func collideWithBox(g *gocui.Gui) error {
	err := addBodyPartToEnd(g, snekBodyParts[len(snekBodyParts)-1]); if err != nil {return err}
	return setViewAtRandom(g, boxView, false)
}

func checkHeadToBodyCollision(g *gocui.Gui) (bool, error) {
	for i := 1; i < len(snekBodyParts); i++ {
		collision, err := checkViewCollision(g, snekBodyParts[0].viewName, snekBodyParts[i].viewName)
		if err != nil {
			return false, err
		}
		if collision {
			return true, nil
		}
	}
	return false, nil
}

func checkHeadToMainViewCollision(g *gocui.Gui, snekHead snekBodyPart) (bool, error) {
	xG0, yG0, xG1, yG1, err := g.ViewPosition(gameView)
	if err != nil {
		return false, err
	}

	xH0, yH0, xH1, yH1, err := g.ViewPosition(snekHead.viewName)
	if err != nil {
		return false, err
	}

	maxX, maxY, minX, minY := xG1-xG0, yG1-yG0, 0, 0
	if xH0 >= minX && yH0 >= minY && xH1 <= maxX && yH1 <= maxY {
		return false, nil
	}
	return true, nil
}

func moveSnekBodyPart(g *gocui.Gui, previousSnekBodyPart snekBodyPart, currentSnekBodyPart snekBodyPart) error {
	pX0, pY0, pX1, pY1, err := g.ViewPosition(previousSnekBodyPart.viewName); if err != nil {return err}
	offsetX, offsetY := calculateOffsets(previousSnekBodyPart)

	_, err = g.SetView(currentSnekBodyPart.viewName, pX0+offsetX, pY0+offsetY, pX1+offsetX, pY1+offsetY, 0)
	return nil
}

func moveHeadView(g *gocui.Gui, snekHead snekBodyPart) error {
	x0, y0, x1, y1, err := g.ViewPosition(snekHead.viewName); if err != nil {return err}
	offsetX, offsetY := calculateOffsets(snekHead)

	newX0, newY0, newX1, newY1 := x0-offsetX, y0-offsetY, x1-offsetX, y1-offsetY
	_, err = g.SetView(snekHead.viewName, newX0, newY0, newX1, newY1, 0); if err != nil {return err}
	return nil
}

func calculateOffsets(snekBodyPart snekBodyPart) (int, int) {
	offsetX := 0
	offsetY := deltaY
	switch snekBodyPart.currentDirection {
	case directions.right:
		offsetX = -deltaX
		offsetY = 0
	case directions.down:
		offsetX = 0
		offsetY = -deltaY
	case directions.left:
		offsetX = deltaX
		offsetY = 0
	}
	return offsetX, offsetY
}
