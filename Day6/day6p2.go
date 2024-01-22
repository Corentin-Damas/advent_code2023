package day6

import (
	"fmt"
)

type Race struct {
	time     int
	distance int
	solution int
}

func Day6p2() int {

	oneRace := Race{
		time:     59796575,
		distance: 597123410341328,
	}

	pRec := &oneRace
	winningSolutions(pRec)

	fmt.Println("Result for day 6 is : ")
	return oneRace.solution
}

func winningSolutions(records *Race) {
	maxTime := records.time
	recordDistance := records.distance
	var numSolutions int = 0

	for pressedTime := 0; pressedTime <= maxTime; pressedTime++ {
		travelTime := maxTime - pressedTime
		speed := pressedTime

		if speed*travelTime > recordDistance {
			numSolutions++
		}

	}
	records.solution = numSolutions
}
