package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
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
	deltaX      = 2
	deltaY      = 1
	boxViewName = "box"
)

var (
	directions    = movementDirections{0, 1, 2, 3}
	headDirection = direction(r.Intn(4))
	snekHead      = &snekBodyPart{headDirection, headDirection, "s0"}
	snekBodyParts = []*snekBodyPart{snekHead}
)

func addBodyPartToEnd(gui *gocui.Gui, currentLastSnekBodyPart snekBodyPart) error {
	x0, y0, x1, y1, err := gui.ViewPosition(currentLastSnekBodyPart.viewName)
	if err != nil {
		return err
	}
	offsetX, offsetY := calculateBodyPartOffsets(currentLastSnekBodyPart)

	name := fmt.Sprintf("s%v", len(snekBodyParts))
	_, err = gui.SetView(name, x0+offsetX, y0+offsetY, x1+offsetX, y1+offsetY, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	snekBodyParts = append(
		snekBodyParts,
		&snekBodyPart{
			currentLastSnekBodyPart.currentDirection,
			currentLastSnekBodyPart.previousDirection,
			name})

	if err := updateStat(&lengthStat, lengthStat.value+1); err != nil {
		return err
	}
	return nil
}

//Checks collision between view1 and view2, returning true for collision and false otherwise.
func checkViewCollision(gui *gocui.Gui, view1 string, view2 string) (bool, error) {
	x10, y10, x11, y11, err := gui.ViewPosition(view1)
	if err != nil {
		return false, err
	}
	x20, y20, x21, y21, err := gui.ViewPosition(view2)
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

func moveSnekHead(gui *gocui.Gui, snekBodyPart *snekBodyPart) error {
	snekBodyPart.previousDirection = snekBodyPart.currentDirection
	snekBodyPart.currentDirection = headDirection

	err := moveHeadView(gui, snekBodyPart)
	if err != nil {
		return err
	}

	headToMainViewCollision, err := checkHeadToMainViewCollision(gui, *snekBodyPart)
	if err != nil {
		return err
	}
	if headToMainViewCollision {
		return gameOver(gui)
	}

	headToBodyCollision, err := checkHeadToBodyCollision(gui)
	if err != nil {
		return err
	}
	if headToBodyCollision {
		return gameOver(gui)
	}

	headToBoxCollision, err := checkViewCollision(gui, snekBodyPart.viewName, boxViewName)
	if err != nil {
		return err
	}
	if headToBoxCollision {
		return collideWithBox(gui)
	}
	return nil
}

func collideWithBox(gui *gocui.Gui) error {
	err := addBodyPartToEnd(gui, *snekBodyParts[len(snekBodyParts)-1])
	if err != nil {
		return err
	}
	return setViewAtRandom(gui, boxViewName, false)
}

func checkHeadToBodyCollision(gui *gocui.Gui) (bool, error) {
	for i := 1; i < len(snekBodyParts); i++ {
		collision, err := checkViewCollision(gui, snekHead.viewName, snekBodyParts[i].viewName)
		if err != nil {
			return false, err
		}
		if collision {
			return true, nil
		}
	}
	return false, nil
}

func checkHeadToMainViewCollision(gui *gocui.Gui, snekHead snekBodyPart) (bool, error) {
	xG0, yG0, xG1, yG1, err := gui.ViewPosition(gameViewName)
	if err != nil {
		return false, err
	}

	xH0, yH0, xH1, yH1, err := gui.ViewPosition(snekHead.viewName)
	if err != nil {
		return false, err
	}

	maxX, maxY, minX, minY := xG1-xG0, yG1-yG0, 0, 0
	if xH0 >= minX && yH0 >= minY && xH1 <= maxX && yH1 <= maxY {
		return false, nil
	}
	return true, nil
}

func moveSnekBodyParts(gui *gocui.Gui) error {
	for i := 1; i < len(snekBodyParts); i++ {
		err := moveSnekBodyPart(gui, snekBodyParts[i-1], snekBodyParts[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func moveSnekBodyPart(gui *gocui.Gui, previousSnekBodyPart *snekBodyPart, currentSnekBodyPart *snekBodyPart) error {
	pX0, pY0, pX1, pY1, err := gui.ViewPosition(previousSnekBodyPart.viewName)
	if err != nil {
		return err
	}
	offsetX, offsetY := calculateBodyPartOffsets(*previousSnekBodyPart)

	_, err = gui.SetView(currentSnekBodyPart.viewName, pX0+offsetX, pY0+offsetY, pX1+offsetX, pY1+offsetY, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}

	currentSnekBodyPart.previousDirection = currentSnekBodyPart.currentDirection
	currentSnekBodyPart.currentDirection = previousSnekBodyPart.previousDirection
	return nil
}

func moveHeadView(gui *gocui.Gui, snekHead *snekBodyPart) error {
	x0, y0, x1, y1, err := gui.ViewPosition(snekHead.viewName)
	if err != nil {
		return err
	}
	offsetX, offsetY := calculateBodyPartOffsets(*snekHead)

	newX0, newY0, newX1, newY1 := x0-offsetX, y0-offsetY, x1-offsetX, y1-offsetY
	_, err = gui.SetView(snekHead.viewName, newX0, newY0, newX1, newY1, 0)
	if err != nil {
		return err
	}
	return nil
}

func calculateBodyPartOffsets(snekBodyPart snekBodyPart) (int, int) {
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

func getOppositeDirection(direction direction) direction {
	return (direction + 2) % 4
}
