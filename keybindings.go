package main

import (
	"github.com/awesome-gocui/gocui"
	"time"
)

func initKeybindings(g *gocui.Gui) error {
	if err := initQuitKey(g); err != nil {return err}
	if err := initSpaceKey(g); err != nil {return err}
	if err := initMovementKeys(g); err != nil {return err}
	if err := initTabKey(g); err != nil {return err}
	if err := initSpeedKeys(g); err != nil {return err}
	return nil
}

func initQuitKey(g *gocui.Gui) error{
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
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
			if snekBodyParts[0].currentDirection == (keyDirection + 2) % 4 {
				return nil
			}
			currentDirection = keyDirection
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initTabKey(g *gocui.Gui) error{
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			err := addView(g, snekBodyParts[len(snekBodyParts)-1].viewName, snekBodyParts[len(snekBodyParts)-1].currentDirection)
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
	if err := initSpeedKey(g, gocui.KeyCtrlW, -10); err != nil {return err}
	if err := initSpeedKey(g, gocui.KeyCtrlS, 10); err != nil {return err}
	return nil
}

func initSpeedKey(g *gocui.Gui, key gocui.Key, speedChange time.Duration) error{
	if err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			tickInterval +=  speedChange * time.Millisecond
			return nil
		}); err != nil {
		return err
	}
	return nil
}
