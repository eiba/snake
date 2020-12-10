package game

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

type snakeBodyPart struct {
	currentDirection  direction
	previousDirection direction
	viewName          string
	position          Position
}

type Position struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
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
	headDirection = direction(main.r.Intn(4))
	snakeHead      = &snakeBodyPart{headDirection, headDirection, "s0", Position{}}
	snakeBodyParts = []*snakeBodyPart{snakeHead}
)

func addBodyPartToEnd(currentLastsnakeBodyPart snakeBodyPart) error {
	offsetX, offsetY := calculateOffsets(currentLastsnakeBodyPart.currentDirection, false)

	name := fmt.Sprintf("s%v", len(snakeBodyParts))
	position := Position{
		currentLastsnakeBodyPart.position.x0 + offsetX,
		currentLastsnakeBodyPart.position.y0 + offsetY,
		currentLastsnakeBodyPart.position.x1 + offsetX,
		currentLastsnakeBodyPart.position.y1 + offsetY,
	}

	_, err := main.gui.SetView(name, position.x0, position.y0, position.x1, position.y1, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	snakeBodyParts = append(
		snakeBodyParts,
		&snakeBodyPart{
			currentLastsnakeBodyPart.currentDirection,
			currentLastsnakeBodyPart.previousDirection,
			name,
			position,
		})
	return main.updateStat(&main.lengthStat, main.lengthStat.value+1)
}

//Checks if there is a collision between Position and all positions in positions
func positionsOverlap(position Position, positions []Position) bool {
	for i := 0; i < len(positions); i++ {
		if positionOverlap(position, positions[i]) {
			return true
		}
	}
	return false
}

//Checks collision between position1 and position2, returning true for collision and false otherwise.
func positionOverlap(position1 Position, position2 Position) bool {
	if position1 == position2 {
		return true
	}
	return false
}

func movesnakeHead() error {
	err := moveHeadView(snakeHead)
	if err != nil {
		return err
	}

	if fatalCollision(snakeHead.position) {
		return main.gameOver("Game Over")
	}

	if positionOverlap(snakeHead.position, main.foodView.position) {
		return main.eatFood()
	}
	return nil
}

func fatalCollision(position Position) bool {
	if mainViewCollision(position) || bodyCollision(position) {
		return true
	}
	return false
}

func bodyCollision(position Position) bool {
	for i := 1; i < len(snakeBodyParts); i++ {
		collision := positionOverlap(position, snakeBodyParts[i].position)
		if collision {
			return true
		}
	}
	return false
}

func mainViewCollision(position Position) bool {
	xG0, yG0, xG1, yG1 := main.gameView.position.x0, main.gameView.position.y0, main.gameView.position.x1, main.gameView.position.y1
	xH0, yH0, xH1, yH1 := position.x0, position.y0, position.x1, position.y1

	maxX, maxY, minX, minY := xG1-xG0, yG1-yG0, 0, 0
	if xH0 >= minX && yH0 >= minY && xH1 <= maxX && yH1 <= maxY {
		return false
	}
	return true
}

func movesnakeBodyParts() error {
	for i := 1; i < len(snakeBodyParts); i++ {
		err := movesnakeBodyPart(snakeBodyParts[i-1], snakeBodyParts[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func movesnakeBodyPart(previoussnakeBodyPart *snakeBodyPart, currentsnakeBodyPart *snakeBodyPart) error {
	currentsnakeBodyPart.position = getPositionOfNextMove(previoussnakeBodyPart.currentDirection, previoussnakeBodyPart.position, false)
	_, err := main.gui.SetView(
		currentsnakeBodyPart.viewName,
		currentsnakeBodyPart.position.x0,
		currentsnakeBodyPart.position.y0,
		currentsnakeBodyPart.position.x1,
		currentsnakeBodyPart.position.y1,
		0)
	if err != nil {
		return err
	}

	currentsnakeBodyPart.previousDirection = currentsnakeBodyPart.currentDirection
	currentsnakeBodyPart.currentDirection = previoussnakeBodyPart.previousDirection
	return nil
}

func moveHeadView(snakeHead *snakeBodyPart) error {
	snakeHead.previousDirection = snakeHead.currentDirection
	snakeHead.currentDirection = headDirection

	snakeHead.position = getPositionOfNextMove(snakeHead.currentDirection, snakeHead.position, true)
	_, err := main.gui.SetView(
		snakeHead.viewName,
		snakeHead.position.x0,
		snakeHead.position.y0,
		snakeHead.position.x1,
		snakeHead.position.y1,
		0)
	if err != nil {
		return err
	}
	return nil
}

func getPositionOfNextMove(currentDirection direction, currentPosition Position, isHead bool) Position {
	offsetX, offsetY := calculateOffsets(currentDirection, isHead)
	return Position{currentPosition.x0 + offsetX, currentPosition.y0 + offsetY, currentPosition.x1 + offsetX, currentPosition.y1 + offsetY}
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

func getsnakePositionSet(snake []*snakeBodyPart) map[Position]bool {
	snakePositionSet := make(map[Position]bool)
	for _, bodyPart := range snake {
		snakePositionSet[bodyPart.position] = true
	}
	return snakePositionSet
}

func getOppositeDirection(direction direction) direction {
	return (direction + 2) % 4
}

func getValidDirections(currentDirection direction) []direction {
	return []direction{currentDirection, (currentDirection + 1) % 4, (currentDirection + 3) % 4}
}
