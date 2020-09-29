package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"time"
)

func initKeybindingsView(gui *gocui.Gui) error {
	maxX, _ := gui.Size()
	if v, err := gui.SetView("keybindings", maxX-25, 0, maxX-1, 7, 0); err != nil {
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

func initKeybindings(gui *gocui.Gui) error {
	if err := initQuitKey(gui); err != nil {return err}
	if err := initSpaceKey(gui); err != nil {return err}
	if err := initMovementKeys(gui); err != nil {return err}
	if err := initTabKey(gui); err != nil {return err}
	if err := initSpeedKeys(gui); err != nil {return err}
	if err := initPauseKey(gui); err != nil {return err}
	return nil
}

func initQuitKey(gui *gocui.Gui) error{
	if err := gui.SetKeybinding("", gocui.KeyEsc, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}

func initMovementKeys(gui *gocui.Gui) error {
	if err := initMovementKey(gui, gocui.KeyArrowUp, directions.up); err != nil {return err}
	if err := initMovementKey(gui, gocui.KeyArrowRight, directions.right); err != nil {return err}
	if err := initMovementKey(gui, gocui.KeyArrowDown, directions.down); err != nil {return err}
	if err := initMovementKey(gui, gocui.KeyArrowLeft, directions.left); err != nil {return err}
	return nil
}

func initMovementKey(gui *gocui.Gui, key gocui.Key, keyDirection direction) error {
	if err := gui.SetKeybinding("", key, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
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

func initTabKey(gui *gocui.Gui) error{
	if err := gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			err := addBodyPartToEnd(gui, *snekBodyParts[len(snekBodyParts)-1])
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initSpaceKey(gui *gocui.Gui) error{
	if err := gui.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return reset(gui)
		}); err != nil {
		return err
	}
	return nil
}

func initSpeedKeys(gui *gocui.Gui) error{
	if err := initSpeedKey(gui, 'w', -10); err != nil {return err}
	if err := initSpeedKey(gui, 's', 10); err != nil {return err}
	return nil
}

func initSpeedKey(gui *gocui.Gui, key rune, speedChange time.Duration) error{
	if err := gui.SetKeybinding("", key, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			tickInterval +=  speedChange * time.Millisecond
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initPauseKey(gui *gocui.Gui) error {
	if err := gui.SetKeybinding("", 'p', gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return pause(gui)
		}); err != nil {
		return err
	}
	return nil
}
