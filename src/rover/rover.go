package rover

import "strconv"

type Rover struct {
	xCoordinate   int
	yCoordinate   int
	headingDirection  int
}

func (r Rover) moveRover(move rune) {
	if move == 'L' || move == 'l' {
		r.moveRoverToLeft()
	} else if move == 'R' || move == 'r' {
		r.moveRoverToRight()
	} else if move == 'M' || move == 'm' {
		r.moveRoverForward()
	}
}

func (r Rover) moveRoverForward() {
	switch r.headingDirection {
	case N.int():
		r.yCoordinate++
		break
	case S.int():
		r.yCoordinate--
		break
	case W.int():
		r.xCoordinate++
		break
	case E.int():
		r.xCoordinate--
		break
	}
}

func (r Rover) moveRoverToLeft() {
	if r.headingDirection != N.int() {
		r.headingDirection--
	} else {
		r.headingDirection = W.int()
	}
}

func (r Rover) moveRoverToRight() {
	if r.headingDirection != W.int() {
		r.headingDirection++
	} else {
		r.headingDirection = N.int()
	}
}

func (r Rover) getPosition() string {
	return string(r.xCoordinate) + " " +string(r.yCoordinate) + " " + strconv.QuoteRune(directions[r.headingDirection])
}
