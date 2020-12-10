package game

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/eiba/snake"
	view2 "github.com/eiba/snake/game/view"
	"time"
)

func initKeybindingsView() error {
	maxX  := main.gameView.position.x1
	if v, err := main.gui.SetView("keybindings", maxX+1, 0, maxX+26, 8, 0); err != nil {
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
	if err := main.gui.SetKeybinding("", gocui.KeyEsc, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}

func initMovementKeys() error {
	if err := initMovementKey(gocui.KeyArrowUp, Directions.Up); err != nil {
		return err
	}
	if err := initMovementKey(gocui.KeyArrowRight, Directions.Right); err != nil {
		return err
	}
	if err := initMovementKey(gocui.KeyArrowDown, Directions.Down); err != nil {
		return err
	}
	if err := initMovementKey(gocui.KeyArrowLeft, Directions.Left); err != nil {
		return err
	}
	return nil
}

func initMovementKey(key gocui.Key, keyDirection Direction) error {
	if err := main.gui.SetKeybinding("", key, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			if snakeHead.currentDirection == GetOppositeDirection(keyDirection) {
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
	if err := main.gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			err := addBodyPartToEnd(*snakeBodyParts[len(snakeBodyParts)-1])
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
	if err := main.gui.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
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
	if err := main.gui.SetKeybinding("", key, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			main.tickInterval += speedChange * time.Millisecond
			if main.tickInterval < time.Millisecond {
				main.tickInterval = time.Millisecond
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initPauseKey() error {
	if err := main.gui.SetKeybinding("", 'p', gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return view2.pause()
		}); err != nil {
		return err
	}
	return nil
}

func initAutoPilotKey() error {
	if err := main.gui.SetKeybinding("", 'a', gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			main.autoPilotEnabled = !main.autoPilotEnabled
			return nil
		}); err != nil {
		return err
	}
	return nil
}