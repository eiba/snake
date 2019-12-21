package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/carlescere/scheduler"
	"log"
	"math/rand"
	"time"
)

const delta = 1
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	views   = []string{}
	snekViews = []string{}
	curView = -1
	idxView = 0
	gameView = "game"
)

func main() {


	run()

}

func run()  {
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

	if v, err := g.SetView("help", maxX-25, 0, maxX-1, 8, 0); err != nil{
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Keybindings"
		fmt.Fprintln(v, "Space: New thing")
		fmt.Fprintln(v, "Tab: Next thing")
		fmt.Fprintln(v, "← ↑ → ↓: Move thing")
		fmt.Fprintln(v, "Backspace: Delete thing")
		fmt.Fprintln(v, "t: Set thing on top")
		fmt.Fprintln(v, "b: Set thing on bottom")
		fmt.Fprintln(v, "^C: Exit")
	}

	if v, err := g.SetView(gameView, 0, 0, maxX-26, maxY-1, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		if _, err := g.SetViewOnBottom(gameView); err != nil{
			return err
		}
		if err := newView(g,true); err != nil {
			log.Panicln(err)
		}
		v.Title = "Snek"
	}

	return nil
}

func initKeybindings(g *gocui.Gui) error {
	job := func() {
		g.Update(func(g *gocui.Gui) error {
			err := newView(g,false)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return newView(g,false)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyBackspace2, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return delView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return nextView(g, true)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, -delta, 0)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, delta, 0)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, delta)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, -delta)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 't', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			_, err := g.SetViewOnTop(views[curView])
			return err
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'b', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			_, err := g.SetViewOnBottom(views[curView])
			return err
		}); err != nil {
		return err
	}

	scheduler.Every(1).Seconds().NotImmediately().Run(job)

	return nil
}



func newView(g *gocui.Gui, setCurrent bool) error {
	x0, y0, x1, y1, err := g.ViewPosition(gameView)
	if err != nil {
		return err
	}

	maxX, maxY := x1-x0-3,y1-y0-2

	positionX, positionY := r.Intn(maxX)+1,r.Intn(maxY)+1

	lenX := 2
	lenY := 1
	name := fmt.Sprintf("v%v", idxView)
	_, err = g.SetView(name, positionX, positionY, positionX+lenX, positionY+lenY, 0)
	if err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
	}

	views = append(views, name)

	if setCurrent {
		if _, err := g.SetCurrentView(name); err != nil {
			log.Panicln(err)
		}
		curView = len(views) - 1

	}
	idxView += 1
	return nil
}

func addView(g *gocui.Gui,v *gocui.View) error  {

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
		fmt.Fprintln(v, "",idxView+1)
	}

	views = append(views, name)
	idxView += 1

	return nil
}

func delView(g *gocui.Gui) error {
	if len(views) <= 1 {
		return nil
	}

	if err := g.DeleteView(views[curView]); err != nil {
		return err
	}
	views = append(views[:curView], views[curView+1:]...)

	return nextView(g, false)
}

func nextView(g *gocui.Gui, disableCurrent bool) error {
	next := curView + 1
	if next > len(views)-1 {
		next = 0
	}

	if _, err := g.SetCurrentView(views[next]); err != nil {
		return err
	}

	curView = next
	return nil
}

func checkCollision(g *gocui.Gui, v *gocui.View)  {

}

func moveView(g *gocui.Gui, v *gocui.View, dx, dy int) error {
	name := v.Name()
	x0, y0, x1, y1, err := g.ViewPosition(name)
	if err != nil {
		return err
	}
	xg0, yg0, xg1, yg1, err := g.ViewPosition(gameView)
	if err != nil {
		return err
	}

	maxX, maxY, minX, minY := xg1-xg0,yg1-yg0,0,0
	newX0,newY0,newX1,newY1 := x0+dx,y0+dy,x1+dx,y1+dy
	if newX0 >= minX && newY0 >= minY && newX1 <= maxX && newY1 <= maxY {
		if _, err := g.SetView(name, newX0, newY0, newX1, newY1, 0); err != nil {
			return err
		}
	}

	return nil
}
