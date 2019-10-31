package main

import (
	"fmt"
)

func main() {
	var roverCount int
	fmt.Println("Enter the no of rovers")
	fmt.Scan(&roverCount)

	if roverCount > 0 {
		for i := 0; i < roverCount; i++ {
			var roverxCoOrd,roveryCoOrd int
			var roverDir string
			var roverMoves string
			fmt.Println("Enter the x co-ordinate")
			fmt.Scan(&roverxCoOrd)
			fmt.Println("Enter the y co-ordinate")
			fmt.Scan(&roveryCoOrd)
			fmt.Println("Enter the direction of the rover")
			fmt.Scan(&roverDir)
			fmt.Println("Enter the moves to be performed")
			fmt.Scan(&roverMoves)

			rover :=  Rover{
				xCoordinate: roverxCoOrd,
				yCoordinate: roveryCoOrd,
				headingDirection: getDirConstant(roverDir).int(),
			}

			moves :=  []rune(roverMoves)
			for j :=0; j < len(moves); j++ {
			rover.moveRover(moves[j])
			}
			rover.printPosition()
		}
	}
}
