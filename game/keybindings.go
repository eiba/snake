package game

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	snakeView "github.com/eiba/snake/game/view"
	"time"
)

func initKeybindingsView(gui *gocui.Gui, gameView snakeView.Properties) error {
	maxX  := gameView.Position.X1
	if v, err := gui.SetView("keybindings", maxX+1, 0, maxX+26, 8, 0); err != nil {
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

/*func initKeybindings(gui *gocui.Gui, snakeBodyParts []*snakeBodyPart, positionMatrix [][]Position, tickInterval *time.Duration, gameFinished bool, running bool, autoPilotEnabled bool) error {
	if err := initQuitKey(gui); err != nil {
		return err
	}
	if err := initSpaceKey(gui,snakeBodyParts, positionMatrix); err != nil {
		return err
	}
	if err := initMovementKeys(gui); err != nil {
		return err
	}
	if err := initTabKey(gui, snakeBodyParts); err != nil {
		return err
	}
	if err := initSpeedKeys(gui, tickInterval); err != nil {
		return err
	}
	if err := initPauseKey(gui, gameFinished, running); err != nil {
		return err
	}
	if err := initAutoPilotKey(gui, autoPilotEnabled); err != nil {
		return err
	}
	return nil
}*/

func initQuitKey(gui *gocui.Gui) error {
	if err := gui.SetKeybinding("", gocui.KeyEsc, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}

func initMovementKeys(gui *gocui.Gui) error {
	if err := initMovementKey(gui, gocui.KeyArrowUp, Directions.Up); err != nil {
		return err
	}
	if err := initMovementKey(gui, gocui.KeyArrowRight, Directions.Right); err != nil {
		return err
	}
	if err := initMovementKey(gui, gocui.KeyArrowDown, Directions.Down); err != nil {
		return err
	}
	if err := initMovementKey(gui, gocui.KeyArrowLeft, Directions.Left); err != nil {
		return err
	}
	return nil
}

func initMovementKey(gui *gocui.Gui, key gocui.Key, keyDirection Direction) error {
	if err := gui.SetKeybinding("", key, gocui.ModNone,
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

func initTabKey(gui *gocui.Gui, snakeBodyParts []*snakeBodyPart) error {
	if err := gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
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

func initSpaceKey(gui *gocui.Gui, snakeBodyParts []*snakeBodyPart, positionMatrix [][]Position) error {
	if err := gui.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return reset(gui, snakeBodyParts, positionMatrix)
		}); err != nil {
		return err
	}
	return nil
}

func initSpeedKeys(gui *gocui.Gui, tickInterval *time.Duration) error {
	if err := initSpeedKey(gui, 'w', tickInterval, -10); err != nil {
		return err
	}
	if err := initSpeedKey(gui, 's', tickInterval, 10); err != nil {
		return err
	}
	return nil
}

func initSpeedKey(gui *gocui.Gui, key rune, tickInterval *time.Duration, speedChange time.Duration) error {
	if err := gui.SetKeybinding("", key, gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			*tickInterval += speedChange * time.Millisecond
			if *tickInterval < time.Millisecond {
				*tickInterval = time.Millisecond
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func initPauseKey(gui *gocui.Gui, gameFinished bool, running bool) (error, bool) {
	if err := gui.SetKeybinding("", 'p', gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			return snakeView.Pause(gui, gameFinished, running)
		}); err != nil {
		return err, false
	}
	return nil, !running
}

func initAutoPilotKey(gui *gocui.Gui, autoPilotEnabled bool) (error, bool) {
	if err := gui.SetKeybinding("", 'a', gocui.ModNone,
		func(gui *gocui.Gui, view *gocui.View) error {
			autoPilotEnabled = !autoPilotEnabled
			return nil
		}); err != nil {
		return err, false
	}
	return nil, autoPilotEnabled
}