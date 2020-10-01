package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"time"
)

func initKeybindingsView() error {
	maxX, _ := gui.Size()
	if v, err := gui.SetView("keybindings", maxX-25, 0, maxX-1, 8, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Keybindings"
		fmt.Fprintln(v, "Space: Restart")
		fmt.Fprintln(v, "← ↑ → ↓: Move")
		fmt.Fprintln(v, "W: Speed up")
		fmt.Fprintln(v, "S: Slow down")
		fmt.Fprintln(v, "P: Pause")
		fmt.Fprintln(v, "A: Toggle autopilot")
		fmt.Fprintln(v, "Esc: Exit")
	}
	return nil
}

func initKeybindings() error {
	if err := initQuitKey(); err != nil {
		return err
	}
	if err := initSpaceKey(); err != nil {
		return err
	}
	if err := initMovementKeys(); err != nil {
		return err
	}
	if err := initTabKey(); err != nil {
		return err
	}
	if err := initSpeedKeys(); err != nil {
		return err
	}
	if err := initPauseKey(); err != nil {
		return err
	}
	if err := initAutoPilotKey(); err != nil {
		return err
	}
	return nil
}

func initQuitKey() error {
	if err := gui.SetKeybinding("", gocui.KeyEsc, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}

func initMovementKeys() error {
	if err := initMovementKey(gocui.KeyArrowUp, directions.up); err != nil {
		return err
	}
	if err := initMovementKey(gocui.KeyArrowRight, directions.right); err != nil {
		return err
	}
	if err := initMovementKey(gocui.KeyArrowDown, directions.down); err != nil {
		return err
	}
	if err := initMovementKey(gocui.KeyArrowLeft, directions.left); err != nil {
		return err
	}
	return nil
}

func initMovementKey(key gocui.Key, keyDirection direction) error {
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

func initTabKey() error {
	if err := gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			err := addBodyPartToEnd(*snekBodyParts[len(snekBodyParts)-1])
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initSpaceKey() error {
	if err := gui.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return reset()
		}); err != nil {
		return err
	}
	return nil
}

func initSpeedKeys() error {
	if err := initSpeedKey('w', -10); err != nil {
		return err
	}
	if err := initSpeedKey('s', 10); err != nil {
		return err
	}
	return nil
}

func initSpeedKey(key rune, speedChange time.Duration) error {
	if err := gui.SetKeybinding("", key, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			tickInterval += speedChange * time.Millisecond
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initPauseKey() error {
	if err := gui.SetKeybinding("", 'p', gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return pause()
		}); err != nil {
		return err
	}
	return nil
}

func initAutoPilotKey() error {
	if err := gui.SetKeybinding("", 'a', gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			autoPilotEnabled = !autoPilotEnabled
			return nil
		}); err != nil {
		return err
	}
	return nil
}
