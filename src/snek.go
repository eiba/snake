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
	snekHead      = &snekBodyPart{headDirection, headDirection, "s0", position{}}
	snekBodyParts = []*snekBodyPart{snekHead}
	foodView      = viewProperties{"food", "", "", position{}}
)

func addBodyPartToEnd(currentLastSnekBodyPart snekBodyPart) error {
	offsetX, offsetY := calculateOffsets(currentLastSnekBodyPart.currentDirection, false)

	name := fmt.Sprintf("s%v", len(snekBodyParts))
	position := position{
		currentLastSnekBodyPart.position.x0 + offsetX,
		currentLastSnekBodyPart.position.y0 + offsetY,
		currentLastSnekBodyPart.position.x1 + offsetX,
		currentLastSnekBodyPart.position.y1 + offsetY,
	}

	_, err := gui.SetView(name, position.x0, position.y0, position.x1, position.y1, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	snekBodyParts = append(
		snekBodyParts,
		&snekBodyPart{
			currentLastSnekBodyPart.currentDirection,
			currentLastSnekBodyPart.previousDirection,
			name,
			position,
		})
	return updateStat(&lengthStat, lengthStat.value+1)
}

//Checks if there is a collision between position and all positions in positions
func positionsOverlap(position position, positions []position) bool {
	for i := 0; i < len(positions); i++ {
		if positionOverlap(position, positions[i]) {
			return true
		}
	}
	return false
}

//Checks collision between position1 and position2, returning true for collision and false otherwise.
func positionOverlap(position1 position, position2 position) bool {
	if position1 == position2{
		 return true
	}
	return false
	/*Ax, Ay, Aw, Ah := position1.x0, position1.y0, position1.x1-position1.x0, position1.y1-position1.y0
	Bx, By, Bw, Bh := position2.x0, position2.y0, position2.x1-position2.x0, position2.y1-position2.y0

	if Bx+Bw > Ax &&
		By+Bh > Ay &&
		Ax+Aw > Bx &&
		Ay+Ah > By {
		return true
	}
	return false*/
}

func moveSnekHead() error {
	err := moveHeadView(snekHead)
	if err != nil {
		return err
	}

	if fatalCollision(snekHead.position) {
		return gameOver()
	}

	if positionOverlap(snekHead.position, foodView.position) {
		return eatFood()
	}
	return nil
}

func fatalCollision(position position) bool {
	if mainViewCollision(position) || bodyCollision(position) {
		return true
	}
	return false
}

func eatFood() error {
	err := addBodyPartToEnd(*snekBodyParts[len(snekBodyParts)-1])
	if err != nil {
		return err
	}
	foodView.position, err = setViewAtRandom(foodView.name, positionMatrix, false)
	return err
}

func bodyCollision(position position) bool {
	for i := 1; i < len(snekBodyParts); i++ {
		collision := positionOverlap(position, snekBodyParts[i].position)
		if collision {
			return true
		}
	}
	return false
}

func mainViewCollision(position position) bool {
	xG0, yG0, xG1, yG1 := gameView.position.x0, gameView.position.y0, gameView.position.x1, gameView.position.y1
	xH0, yH0, xH1, yH1 := position.x0, position.y0, position.x1, position.y1

	maxX, maxY, minX, minY := xG1-xG0, yG1-yG0, 0, 0
	if xH0 >= minX && yH0 >= minY && xH1 <= maxX && yH1 <= maxY {
		return false
	}
	return true
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
	currentSnekBodyPart.position = getPositionOfNextMove(previousSnekBodyPart.currentDirection, previousSnekBodyPart.position, false)
	_, err := gui.SetView(
		currentSnekBodyPart.viewName,
		currentSnekBodyPart.position.x0,
		currentSnekBodyPart.position.y0,
		currentSnekBodyPart.position.x1,
		currentSnekBodyPart.position.y1,
		0)
	if err != nil {
		return err
	}

	currentSnekBodyPart.previousDirection = currentSnekBodyPart.currentDirection
	currentSnekBodyPart.currentDirection = previousSnekBodyPart.previousDirection
	return nil
}

func moveHeadView(snekHead *snekBodyPart) error {
	snekHead.previousDirection = snekHead.currentDirection
	snekHead.currentDirection = headDirection

	snekHead.position = getPositionOfNextMove(snekHead.currentDirection, snekHead.position, true)
	_, err := gui.SetView(
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

func getPositionOfNextMove(currentDirection direction, currentPosition position, isHead bool) position {
	offsetX, offsetY := calculateOffsets(currentDirection, isHead)
	return position{currentPosition.x0 + offsetX, currentPosition.y0 + offsetY, currentPosition.x1 + offsetX, currentPosition.y1 + offsetY}
}

func calculateOffsets(direction direction, isHead bool) (int, int) {
	modifier := 1
	if isHead {
		modifier = -1
	}

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
	return modifier * offsetX, modifier * offsetY
}

func getOppositeDirection(direction direction) direction {
	return (direction + 2) % 4
}
