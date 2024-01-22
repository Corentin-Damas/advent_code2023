package day6

import (
	"fmt"
)

type Record struct {
	time     []int
	distance []int
	solution []int
}

func Day6p1() int {
	// recTest := Record{
	// 	time:     []int{7, 15, 30},
	// 	distance: []int{9, 40, 200},
	// 	solution: []int{},
	// }
	records := Record{
		time:     []int{59, 79, 65, 75},
		distance: []int{597, 1234, 1034, 1328},
		solution: []int{},
	}
	pRec := &records

	for idx := range records.time {
		countWinningSolutions(pRec, idx)
	}

	acc := 1
	for _, sol := range records.solution {
		acc *= sol
	}

	fmt.Println("Result for day 6 is : ")
	return acc
}

// Start @ 0 Milmeter/milsec
// each second pressed := +1 milmeter/milSec but -1s of traveling

func countWinningSolutions(records *Record, idx int) {
	maxTime := records.time[idx]
	recordDistance := records.distance[idx]
	var numSolutions int = 0

	for pressedTime := 0; pressedTime <= maxTime; pressedTime++ {
		travelTime := maxTime - pressedTime
		speed := pressedTime

		if speed*travelTime > recordDistance {
			numSolutions++
		}

	}

	records.solution = append(records.solution, numSolutions)

}
