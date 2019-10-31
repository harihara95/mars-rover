package main

import (
	"fmt"
	"strconv"
)

type Rover struct {
	xCoordinate   int
	yCoordinate   int
	headingDirection  int
}

func (r *Rover) moveRover(move rune) {

	if move == 'L'  {
		r.moveRoverToLeft()
	} else if move == 'R' {
		r.moveRoverToRight()
	} else if move == 'M' {
		r.moveRoverForward()
	}
}

func (r *Rover) moveRoverForward() {
	switch r.headingDirection {
	case N.int():
		r.yCoordinate++
		break
	case S.int():
		r.yCoordinate--
		break
	case W.int():
		r.xCoordinate--
		break
	case E.int():
		r.xCoordinate++
		break
	}
}

func (r *Rover) moveRoverToLeft() {
	if r.headingDirection != N.int() {
		r.headingDirection--
	} else {
		r.headingDirection = W.int()
	}
}

func (r *Rover) moveRoverToRight() {
	if r.headingDirection != W.int() {
		r.headingDirection++
	} else {
		r.headingDirection = N.int()
	}
}

func (r Rover) printPosition() {

	fmt.Print(r.xCoordinate)
	fmt.Print(" ")
	fmt.Print(r.yCoordinate)
	fmt.Print(" ")
	fmt.Print(strconv.QuoteRune(directions[r.headingDirection]))
}
