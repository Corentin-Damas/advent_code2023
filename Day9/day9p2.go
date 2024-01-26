package day9

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Day9p2() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day9/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep2(puzzle)

	fmt.Println("Day 9 part 2 result is : ")
	return result
}

func analyzep2(s string) int {
	strLines := strings.Split(s, "\n")
	var numsArr [][]int
	var acc int = 0

	// Each line = historic report for a single value
	// Each line should? have logical increment at first sight

	//	0 3 6 9 12 15 X?
	//	1 3 6 10 15 21 X?
	//	10 13 16 21 30 45 X?
	// All Prediction for the next value of each line add up for final result

	// take a line > create a new line where the difference between 2 number is added in that new line
	// Repeat until last line is 0

	// then add a 0 at the end of the 0s line. Bubble up

	for _, line := range strLines {
		numsArr = append(numsArr, toNumsArr(line))
	}

	for _, numArr := range numsArr {
		predLine := getPrePrediction(numArr)
		acc += predLine

	}
	return acc
}

func getPrePrediction(puzzleLine []int) int {
	// take a line > create a new line where the difference between 2 number is added in that new line
	// Repeat until last line is 0
	var treeArr [][]int
	var idxTree int = 0

	treeArr = append(treeArr, puzzleLine)

	// Goes Down the tree Create regulare line to find line 0
	for {
		actualArr := treeArr[idxTree]

		if is_zeroArr(actualArr) {
			break
		}

		var nextArr []int
		treeArr = append(treeArr, nextArr)

		for idx := range actualArr {
			// Don't go on the last number idx+1 dosn't exist
			if idx == len(actualArr)-1 {
				break
			}

			newValueNextLine := actualArr[idx+1] - actualArr[idx]             // Create new value for next line
			treeArr[idxTree+1] = append(treeArr[idxTree+1], newValueNextLine) // Add that new val to the next line
		}

		idxTree += 1 // loop go to next line of the main tree

		// Should we considere that the 'last line' will never be 0 ?
	}

	// Goes up the tree calculate prediction
	len_tree := len(treeArr) - 1

	for i := len_tree; i >= 0; i-- {

		if i == len_tree {
			treeArr[i] = append( []int{0} ,treeArr[i]...)
			continue
		}

		newVal := treeArr[i][0] - treeArr[i+1][0] // first item of the actual list + first item next line

		treeArr[i] = append( []int{newVal} , treeArr[i]...)

	}

	return treeArr[0][0]
}
