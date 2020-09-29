package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"time"
)

func initKeybindingsView(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if v, err := g.SetView("keybindings", maxX-25, 0, maxX-1, 7, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Keybindings"
		fmt.Fprintln(v, "Space: Restart")
		fmt.Fprintln(v, "← ↑ → ↓: Move")
		fmt.Fprintln(v, "W: Speed up")
		fmt.Fprintln(v, "S: Slow down")
		fmt.Fprintln(v, "P: Pause")
		fmt.Fprintln(v, "Esc: Exit")
	}
	return nil
}

func initKeybindings(g *gocui.Gui) error {
	if err := initQuitKey(g); err != nil {return err}
	if err := initSpaceKey(g); err != nil {return err}
	if err := initMovementKeys(g); err != nil {return err}
	if err := initTabKey(g); err != nil {return err}
	if err := initSpeedKeys(g); err != nil {return err}
	if err := initPauseKey(g); err != nil {return err}
	return nil
}

func initQuitKey(g *gocui.Gui) error{
	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}

func initMovementKeys(g *gocui.Gui) error {
	if err := initMovementKey(g, gocui.KeyArrowUp, directions.up); err != nil {return err}
	if err := initMovementKey(g, gocui.KeyArrowRight, directions.right); err != nil {return err}
	if err := initMovementKey(g, gocui.KeyArrowDown, directions.down); err != nil {return err}
	if err := initMovementKey(g, gocui.KeyArrowLeft, directions.left); err != nil {return err}
	return nil
}

func initMovementKey(g *gocui.Gui, key gocui.Key, keyDirection direction) error {
	if err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if snekHead.currentDirection == getOppositeDirection(keyDirection) {
				return nil
			}
			headDirection = keyDirection
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initTabKey(g *gocui.Gui) error{
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			err := addBodyPartToEnd(g, *snekBodyParts[len(snekBodyParts)-1])
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initSpaceKey(g *gocui.Gui) error{
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return reset(g)
		}); err != nil {
		return err
	}
	return nil
}

func initSpeedKeys(g *gocui.Gui) error{
	if err := initSpeedKey(g, 'w', -10); err != nil {return err}
	if err := initSpeedKey(g, 's', 10); err != nil {return err}
	return nil
}

func initSpeedKey(g *gocui.Gui, key rune, speedChange time.Duration) error{
	if err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			tickInterval +=  speedChange * time.Millisecond
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initPauseKey(g *gocui.Gui) error {
	if err := g.SetKeybinding("", 'p', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return pause(g)
		}); err != nil {
		return err
	}
	return nil
}
