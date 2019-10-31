package main

import (
	"fmt"
	"gitlab.com/harihara95/mars-rover/rover"
	)

func main() {
	var roverCount int
	fmt.Scan(&roverCount)
	if roverCount > 0 {
		for i := 0; i < roverCount; i++ {
			var roverxCoOrd,roveryCoOrd int
			var roverDir rune
			var roverMoves string

			fmt.Scan(&roverxCoOrd)
			fmt.Scan(&roveryCoOrd)
			fmt.Scan(&roverDir)
			fmt.Scan(&roverMoves)

			rover := rover.Rover {
				xCoordinate: roverxCoOrd,
				yCoordinate: roveryCoOrd,
				headingDirection: getDirConstant(roverDir).int(),
			}
			for move := range rover1Moves {
				rover.moveRover(move)
			}
			fmt.Println(rover.getPosition())
		}
	}
}
