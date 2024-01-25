package day8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type GhostDirection struct {
	left        int
	right       int
	start       []string
	end         string
	instruction string
}

func Day8p2() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day8/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep2(puzzle)

	fmt.Println("Day 8 part 2 result is : ")
	return result
}

func analyzep2(s string) int {
	lines := strings.Split(s, "\n")

	directions := lines[0]
	directions = strings.Trim(directions, "\r\n")
	mapInstruction, startingPoints := getGhostMapTranslation(lines[2:])

	lrGhostDir := GhostDirection{
		left:        0,
		right:       1,
		start:       startingPoints,
		end:         "ZZZ",
		instruction: directions,
	}

	fmt.Println("Direction are: ", directions)
	fmt.Println("Starting Points: ", startingPoints)

	// for key, val := range mapInstruction {
	// 	fmt.Println(key, " : ", val)
	// }

	pGhostDir := &lrGhostDir

	acc := travelTimeForGhost(mapInstruction, pGhostDir)

	return acc

}

func getGhostMapTranslation(puzzleLines []string) (map[string][2]string, []string) {

	puzzleMap := make(map[string][2]string)
	startingPoints := []string{}

	for _, l := range puzzleLines {
		keyValue := strings.Split(l, " = ")
		key := keyValue[0]
		if string(key[2]) == "A" {
			startingPoints = append(startingPoints, key)
		}

		values := strings.Split(keyValue[1], ", ")
		leftVal := values[0][1:4]
		rigVal := values[1][0:3]

		puzzleMap[key] = [2]string{leftVal, rigVal}
	}
	return puzzleMap, startingPoints
}

func travelTimeForGhost(puzzleMap map[string][2]string, direction *GhostDirection) int {
	actualPos := direction.start

	var cyclePerStartingPos [][]int
	// Do a cache

	for _, currPos := range actualPos {

		// How much step does a starting point take to come back at the start for every Direction loop it takes
		currentSteps := direction.instruction
		cycle := []int{}
		step_count := 0
		var first_z string = ""

		for {
			for step_count == 0 || string(currPos[2]) != "Z" {
				step_count += 1

				if string(currentSteps[0]) == "L" {
					currPos = puzzleMap[currPos][direction.left]
				}
				if string(currentSteps[0]) == "R" {
					currPos = puzzleMap[currPos][direction.right]
				}

				// Infinite loop on the instructions .. It keep regenerating by puling the accutal intruction to the end
				newStart := currentSteps[1:]
				oldStart := string(currentSteps[0])
				currentSteps = newStart + oldStart

			}
			// Get the number of step to arrive from A -> Z for a particular path and for a complete loop
			cycle = append(cycle, step_count)
			if first_z == "" {
				first_z = currPos
				step_count = 0
			} else if currPos == first_z {
				break
			}
		}
		// Cycle has 2 number the number of time ir take to go from A-> Z from start to finish with a normal loop and and if it has to start from current pos of Z 
		// Where they should be equals
		cyclePerStartingPos = append(cyclePerStartingPos, cycle)
	}

	var nums []int

	// cyclePerStartingPos hold in memory the number of step it take for each pass to do a A->Z
	for _, cy := range cyclePerStartingPos {
		nums = append(nums, cy[0])
	}
	fmt.Println("Doing the Calc of the LCM (Least commmon Multiplier)", nums)

	// we need a starting point to calc the LCM so we take the last elem of the least and remove it so we don't calculate twice
	lenNum := len(nums) - 1
	lcm := nums[lenNum]
	nums = nums[:lenNum]

	for _, num := range nums {
		// Function and Calc found on internet
		lcm = lcm * num / GCD(lcm, num)
	}

	return lcm

}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
