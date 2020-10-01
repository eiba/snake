package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

type snekBodyPart struct {
	currentDirection  direction
	previousDirection direction
	viewName          string
	position          position
}

type position struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

type direction int
type movementDirections struct {
	up    direction
	right direction
	down  direction
	left  direction
}

const (
	deltaX = 2
	deltaY = 1
)

var (
	directions    = movementDirections{0, 1, 2, 3}
	headDirection = direction(r.Intn(4))
	snekHead      = &snekBodyPart{headDirection, headDirection, "s0", position{0, 0, 0, 0}}
	snekBodyParts = []*snekBodyPart{snekHead}
	boxView       = viewProperties{"box", "", "", position{0, 0, 0, 0}}
)

func addBodyPartToEnd(currentLastSnekBodyPart snekBodyPart) error {
	x0, y0, x1, y1, err := gui.ViewPosition(currentLastSnekBodyPart.viewName)
	if err != nil {
		return err
	}
	offsetX, offsetY := calculateBodyPartOffsets2(currentLastSnekBodyPart.currentDirection)

	name := fmt.Sprintf("s%v", len(snekBodyParts))
	position := position{x0 + offsetX, y0 + offsetY, x1 + offsetX, y1 + offsetY}
	_, err = gui.SetView(name, x0+offsetX, y0+offsetY, x1+offsetX, y1+offsetY, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	snekBodyParts = append(
		snekBodyParts,
		&snekBodyPart{
			currentLastSnekBodyPart.currentDirection,
			currentLastSnekBodyPart.previousDirection,
			name, position})

	if err := updateStat(&lengthStat, lengthStat.value+1); err != nil {
		return err
	}
	return nil
}

//Checks collision between view1 and view2, returning true for collision and false otherwise.
func checkViewCollision(view1 string, view2 string) (bool, error) {
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

//Checks collision between position1 and position2, returning true for collision and false otherwise.
func positionOverlap(position1 position, position2 position) bool {
	Ax, Ay, Aw, Ah := position1.x0, position1.y0, position1.x1-position1.x0, position1.y1-position1.y0
	Bx, By, Bw, Bh := position2.x0, position2.y0, position2.x1-position2.x0, position2.y1-position2.y0

	if Bx+Bw > Ax &&
		By+Bh > Ay &&
		Ax+Aw > Bx &&
		Ay+Ah > By {
		return true
	}
	return false
}

func moveSnekHead() error {
	snekHead.previousDirection = snekHead.currentDirection
	snekHead.currentDirection = headDirection

	err := moveHeadView(snekHead)
	if err != nil {
		return err
	}

	fatalCollision, err := fatalCollision(snekHead.position)
	if err != nil {
		return err
	}
	if fatalCollision {
		return gameOver()
	}

	/*headToBodyCollision := checkBodyCollision(snekHead.position)

	if headToBodyCollision {
		return gameOver()
	}*/

	headToBoxCollision := positionOverlap(snekHead.position, boxView.position)
	/*if err != nil {
		return err
	}*/
	if headToBoxCollision {
		return collideWithBox()
	}
	return nil
}

func fatalCollision(position position) (bool, error) {
	mainViewCollision, err := mainViewCollision(position)
	if err != nil {
		return true,err
	}
	if mainViewCollision {
		return true, nil
	}
	if checkBodyCollision(position) {return true, nil}

	return false,nil
}

func collideWithBox() error {
	err := addBodyPartToEnd(*snekBodyParts[len(snekBodyParts)-1])
	if err != nil {
		return err
	}
	boxView.position, err = setViewAtRandom(boxView.name, false)
	return err
}

func checkBodyCollision(position position) bool {
	for i := 1; i < len(snekBodyParts); i++ {
		collision := positionOverlap(position, snekBodyParts[i].position)
		/*collision, err := checkViewCollision(snekBodyPart.viewName, snekBodyParts[i].viewName)
		if err != nil {
			return false, err
		}*/
		if collision {
			return true
		}
	}
	return false
}

func collision(position position, positions []position) bool {
	for i := 1; i < len(positions); i++ {
		if positionOverlap(position, positions[i]){
			return true
		}
	}
	return false
}

func mainViewCollision(position position) (bool, error) {
	xG0, yG0, xG1, yG1, err := gui.ViewPosition(gameViewName)
	if err != nil {
		return false, err
	}

	//position := snekBodyPart.position
	xH0, yH0, xH1, yH1 := position.x0, position.y0, position.x1, position.y1

	maxX, maxY, minX, minY := xG1-xG0, yG1-yG0, 0, 0
	if xH0 >= minX && yH0 >= minY && xH1 <= maxX && yH1 <= maxY {
		return false, nil
	}
	return true, nil
}

func moveSnekBodyParts() error {
	for i := 1; i < len(snekBodyParts); i++ {
		err := moveSnekBodyPart(snekBodyParts[i-1], snekBodyParts[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func moveSnekBodyPart(previousSnekBodyPart *snekBodyPart, currentSnekBodyPart *snekBodyPart) error {
	/*pX0, pY0, pX1, pY1, err := gui.ViewPosition(previousSnekBodyPart.viewName)
	if err != nil {
		return err
	}
	offsetX, offsetY := calculateBodyPartOffsets(*previousSnekBodyPart)*/

	currentSnekBodyPart.position = getPositionOfNextMove2(previousSnekBodyPart.currentDirection,previousSnekBodyPart.position)
	_, err := gui.SetView(
		currentSnekBodyPart.viewName,
		currentSnekBodyPart.position.x0,
		currentSnekBodyPart.position.y0,
		currentSnekBodyPart.position.x1,
		currentSnekBodyPart.position.y1,
		0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}

	currentSnekBodyPart.previousDirection = currentSnekBodyPart.currentDirection
	currentSnekBodyPart.currentDirection = previousSnekBodyPart.previousDirection
	return nil
}

func moveHeadView(snekHead *snekBodyPart) error {
	x0, y0, x1, y1, err := gui.ViewPosition(snekHead.viewName)
	if err != nil {
		return err
	}
	offsetX, offsetY := calculateBodyPartOffsets2(snekHead.currentDirection)

	snekHead.position = position{x0-offsetX, y0-offsetY, x1-offsetX, y1-offsetY}
	_, err = gui.SetView(
		snekHead.viewName,
		snekHead.position.x0,
		snekHead.position.y0,
		snekHead.position.x1,
		snekHead.position.y1,
		0)
	if err != nil {
		return err
	}
	return nil
}

func getPositionOfNextMove(snekBodyPart snekBodyPart) position  {
	offsetX, offsetY := calculateBodyPartOffsets2(snekBodyPart.currentDirection)
	currentPosition := snekBodyPart.position
	return position{currentPosition.x0+offsetX, currentPosition.y0+offsetY, currentPosition.x1+offsetX, currentPosition.y1+offsetY}
}

func getPositionOfNextMove2(currentDirection direction, currentPosition position) position  {
	offsetX, offsetY := calculateBodyPartOffsets2(currentDirection)
	return position{currentPosition.x0+offsetX, currentPosition.y0+offsetY, currentPosition.x1+offsetX, currentPosition.y1+offsetY}
}

func getPositionOfNextMove3(currentDirection direction, currentPosition position) position  {
	offsetX, offsetY := calculateBodyPartOffsets2(currentDirection)
	return position{currentPosition.x0-offsetX, currentPosition.y0-offsetY, currentPosition.x1-offsetX, currentPosition.y1-offsetY}
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

func calculateBodyPartOffsets2(direction direction) (int, int) {
	offsetX := 0
	offsetY := deltaY
	switch direction {
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
