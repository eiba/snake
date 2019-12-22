package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"log"
	"math/rand"
	"time"
)

const delta = 1

var (
	snekViews = []string{}

	direction                   = 0
	idxView                     = 0
	gameView, boxView, snekView = "game", "box", "snek"
	running                     = true
	tickInterval                = 100 * time.Millisecond
	r                           = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	run()
}

func run() {
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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("help", maxX-25, 0, maxX-1, 6, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Keybindings"
		fmt.Fprintln(v, "Space: Restart")
		fmt.Fprintln(v, "← ↑ → ↓: Move thing")
		fmt.Fprintln(v, "Ctrl+W: Speed up")
		fmt.Fprintln(v, "Ctrl+S: Slow down")
		fmt.Fprintln(v, "^C: Exit")
	}

	if v, err := g.SetView(gameView, 0, 0, maxX-26, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		if _, err := g.SetViewOnBottom(gameView); err != nil {
			return err
		}
		if err := setViewAtRandom(g, snekView, true); err != nil {
			log.Panicln(err)
		}
		go updateMovement(g, snekView)
		if err := setViewAtRandom(g, boxView, false); err != nil {
			log.Panicln(err)
		}
		v.Title = "Snek"
	}

	return nil
}

func updateMovement(g *gocui.Gui, viewName string) {
	for {
		time.Sleep(tickInterval)
		if !running {
			continue
		}
		g.Update(func(g *gocui.Gui) error {
			var err error
			switch direction {
			case 0: //up
				err = moveView(g, viewName, 0, -delta)
			case 1: //right
				err = moveView(g, viewName, delta+1, 0)
			case 2: //down
				err = moveView(g, viewName, 0, delta)
			case 3: //left
				err = moveView(g, viewName, -delta-1, 0)
			}
			return err
		})
	}
}

func reset(g *gocui.Gui) error {
	direction = 0
	running = true
	tickInterval = 100 * time.Millisecond
	if err := setViewAtRandom(g, snekView, true); err != nil {
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

		v.Title = "Game over"
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

func initKeybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return reset(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if direction == 1 {
				return nil
			}
			direction = 3
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if direction == 3 {
				return nil
			}
			direction = 1
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if direction == 0 {
				return nil
			}
			direction = 2
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if direction == 2 {
				return nil
			}
			direction = 0
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlW, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			tickInterval -= 10 * time.Millisecond
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("",  gocui.KeyCtrlS, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			tickInterval += 10 * time.Millisecond
			return nil
		}); err != nil {
		return err
	}

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

func addView(g *gocui.Gui, v *gocui.View) error {

	name := v.Name()
	x0, y0, x1, y1, err := g.ViewPosition(name)
	if err != nil {
		return err
	}
	name = fmt.Sprintf("v%v", idxView)
	lenX := 4
	lenY := 2
	if idxView >= 9 {
		lenX++
	}
	v, err = g.SetView(name, x0, y0+lenY, x1, y1+lenY, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Wrap = true
		fmt.Fprintln(v, "", idxView+1)
	}

	idxView += 1

	return nil
}

func checkCollision(g *gocui.Gui, view1 string, view2 string) (bool, error) {
	x10, y10, x11, y11, err := g.ViewPosition(view1)
	if err != nil {
		return false, err
	}

	x20, y20, x21, y21, err := g.ViewPosition(view2)
	if err != nil {
		return false, err
	}

	Ax, Ay, Aw, Ah := x10, y10, x11-x10, y11-y10
	Bx, By, Bw, Bh := x20, y20, x21-x20, y21-y20

	if Bx+Bw > Ax &&
		By+Bh > Ay &&
		Ax+Aw > Bx &&
		Ay+Ah > By {
		return true, nil

	}

	return false, nil
}

func moveView(g *gocui.Gui, viewName string, dx, dy int) error {
	x0, y0, x1, y1, err := g.ViewPosition(viewName)
	if err != nil {
		return err
	}
	xg0, yg0, xg1, yg1, err := g.ViewPosition(gameView)
	if err != nil {
		return err
	}

	maxX, maxY, minX, minY := xg1-xg0, yg1-yg0, 0, 0
	newX0, newY0, newX1, newY1 := x0+dx, y0+dy, x1+dx, y1+dy
	if newX0 >= minX && newY0 >= minY && newX1 <= maxX && newY1 <= maxY {
		if _, err := g.SetView(viewName, newX0, newY0, newX1, newY1, 0); err != nil {
			return err
		}

		collision, err := checkCollision(g, snekView, boxView)
		if err != nil {
			return err
		}

		if collision {
			return setViewAtRandom(g, boxView, false)
		}
	} else {
		return gameOver(g)
	}

	return nil
}
