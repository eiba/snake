package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

func addBodyPartToEnd(g *gocui.Gui, currentLastSnekBodyPart snekBodyPart) error {
	x0, y0, x1, y1, err := g.ViewPosition(currentLastSnekBodyPart.viewName); if err != nil {return err}
	offsetX, offsetY := calculateBodyPartOffsets(currentLastSnekBodyPart)

	name := fmt.Sprintf("s%v", len(snekBodyParts))
	_, err = g.SetView(name, x0+offsetX, y0+offsetY, x1+offsetX, y1+offsetY, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
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

func moveSnekHead(g *gocui.Gui, snekBodyPart *snekBodyPart) error {
	snekBodyPart.previousDirection = snekBodyPart.currentDirection
	snekBodyPart.currentDirection = headDirection

	err := moveHeadView(g, snekBodyPart); if err != nil {return err}

	headToMainViewCollision, err := checkHeadToMainViewCollision(g, *snekBodyPart); if err != nil {return err}
	if headToMainViewCollision {
		return gameOver(g)
	}

	headToBodyCollision, err := checkHeadToBodyCollision(g); if err != nil {return err}
	if headToBodyCollision {
		return gameOver(g)
	}

	headToBoxCollision, err := checkViewCollision(g, snekBodyPart.viewName, boxViewName); if err != nil {return err}
	if headToBoxCollision {
		return collideWithBox(g)
	}
	return nil
}

func collideWithBox(g *gocui.Gui) error {
	err := addBodyPartToEnd(g, snekBodyParts[len(snekBodyParts)-1]); if err != nil {return err}
	return setViewAtRandom(g, boxViewName, false)
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
	xG0, yG0, xG1, yG1, err := g.ViewPosition(gameViewName)
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

func moveSnekBodyParts(g *gocui.Gui) error  {
	for i := 1; i < len(snekBodyParts); i++ {
		err := moveSnekBodyPart(g,  &snekBodyParts[i-1],  &snekBodyParts[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func moveSnekBodyPart(g *gocui.Gui, previousSnekBodyPart *snekBodyPart, currentSnekBodyPart *snekBodyPart) error {
	pX0, pY0, pX1, pY1, err := g.ViewPosition(previousSnekBodyPart.viewName); if err != nil {return err}
	offsetX, offsetY := calculateBodyPartOffsets(*previousSnekBodyPart)

	_, err = g.SetView(currentSnekBodyPart.viewName, pX0+offsetX, pY0+offsetY, pX1+offsetX, pY1+offsetY, 0)
	if err != nil && !gocui.IsUnknownView(err) {
			return err
	}

	currentSnekBodyPart.previousDirection = currentSnekBodyPart.currentDirection
	currentSnekBodyPart.currentDirection = previousSnekBodyPart.previousDirection
	return nil
}

func moveHeadView(g *gocui.Gui, snekHead *snekBodyPart) error {
	x0, y0, x1, y1, err := g.ViewPosition(snekHead.viewName); if err != nil {return err}
	offsetX, offsetY := calculateBodyPartOffsets(*snekHead)

	newX0, newY0, newX1, newY1 := x0-offsetX, y0-offsetY, x1-offsetX, y1-offsetY
	_, err = g.SetView(snekHead.viewName, newX0, newY0, newX1, newY1, 0); if err != nil {return err}
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
