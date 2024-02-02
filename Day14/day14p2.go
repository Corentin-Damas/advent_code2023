package day14

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type void struct{}

var member void

func Day14p2() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day14/datatest.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep2(puzzle)

	fmt.Println("Day 14 part 2 : ")
	return result
}

func analyzep2(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}

	puzzleArr := lineToAtt(puzzleLines)

	iter := 0

	var firstLoopStart int
	seen := make(map[string]int)
	var allGrids [][][]string

	actualCycle := puzzleArr

	for iter <= 1000000000 {
		iter += 1
		var keyCycle string
		tempGrid := actualCycle

		actualCycle, keyCycle = completeCycle(actualCycle)
		
		fmt.Println("iteration: ", iter)
		allGrids = append(allGrids, tempGrid)
		loopStartKey, exists := seen[keyCycle]
		if exists {
			firstLoopStart = loopStartKey
			break
		}
		seen[keyCycle] = iter 
		
	}

	printPuzzle(allGrids[4])
	printPuzzle(allGrids[8])

	fmt.Println("First loop detected at iter : ", iter, " goes back to loop: ", firstLoopStart)

	idxTry := (100000000 - firstLoopStart) % ((iter - firstLoopStart) + firstLoopStart)
	fmt.Println("final result idx:  ", idxTry)

	finalGrid := allGrids[(100000000-firstLoopStart)%(iter-firstLoopStart)+firstLoopStart-1]

	fmt.Println("result ")
	printPuzzle(finalGrid)
	// printPuzzle(puzzleArr)

	return resultPuzzle(finalGrid)
}


func completeCycle(puzzle [][]string) ([][]string, string) {
	cycledPuzzle := puzzle
	for i := 1; i <= 4; i++ {

		// Twice to secure the fact that all rock rolled correctly
		cycledPuzzle = rollTherock(cycledPuzzle)
		cycledPuzzle = rollTherock(cycledPuzzle)

		cycledPuzzle = flipthemap(cycledPuzzle)
	}

	keyCycle := puzzleToString(cycledPuzzle)

	return cycledPuzzle, keyCycle
}

func puzzleToString(puzzle [][]string) string {
	keyCycle := ""
	for _, line := range puzzle {
		for _, ch := range line {
			keyCycle += ch
		}
	}
	return keyCycle
}

func flipthemap(puzzle [][]string) [][]string {
	flipPuzzle := puzzle

	//diagonal flip
	for idxL, line := range flipPuzzle {
		for idxC := range line {
			if idxL != idxC && idxL < idxC {
				temp := flipPuzzle[idxL][idxC]
				flipPuzzle[idxL][idxC] = flipPuzzle[idxC][idxL]
				flipPuzzle[idxC][idxL] = temp
			}
		}
	}
	// horizontal flip ( top is bottom bottom is top)
	// i top j bottom while i<j (equal stop) i++ j--

	// J'AI ROLL EAST ET PAS OEST
	for idxL := range flipPuzzle {

		for i, j := 0, len(flipPuzzle[idxL])-1; i < j; i, j = i+1, j-1 {
			flipPuzzle[idxL][i], flipPuzzle[idxL][j] = flipPuzzle[idxL][j], flipPuzzle[idxL][i]
		}
	}

	return flipPuzzle
}

// Brutforce use less but if we see a cycle we can probably get out

// grid = array[(100000000 -first)%(iter - first) +first ]
// how far we went arount the loop %  Lenght of the loop = arrive at the "end of the "100000000" iterations
