package day13

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Day13p3() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day13/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	puzzle := string(bytesText)
	result := analyzep2(puzzle)
	fmt.Println("Day 13 part 2 result is : ")
	fmt.Println("We should find : 29276")
	return result
}

func analyzep3(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r")
	}
	allPuzzles := getAllPuzzles(puzzleLines)

	total := 0
	for _, puzzle := range allPuzzles {
		idxHor, foundSmudgehr := horizontalAnalyzep3(puzzle)
		if foundSmudgehr {
			total += (idxHor * 100)
		} else {
			idxHor = 0
		}
		swapedPuzzle := swap(puzzle)
		idxVer, foundSmudgeVr:= horizontalAnalyzep3(swapedPuzzle)

		if foundSmudgeVr {
			total += idxVer
		} else {
			idxVer = 0
		}

		fmt.Println(idxHor, idxVer)

	}
	return total
}

func horizontalAnalyzep3(puzzle []string) (int, bool) {

	for idx := 0; idx < len(puzzle)-1; idx++ {
		// fmt.Println(idx, puzzle[idx], idx+1, puzzle[idx+1], "max: ", len(puzzle)-1)

		if checkForSmudge(puzzle[idx], puzzle[idx+1]) {
			return idx + 1, true
		}
		if puzzle[idx] == puzzle[idx+1] {

			is_smudgeMiror := is_RealMirorp3(idx, puzzle)

			if is_smudgeMiror {
				return idx + 1, is_smudgeMiror
			} 
		}

		if idx == len(puzzle)-1 {
			break
		}
	}
	return 0, false
}

func is_RealMirorp3(sep int, puzzle []string) (bool) {
	before := sep
	after := sep + 1
	smudgeAlreadyFound := false

	for {
		if after <= len(puzzle)-1 {

			if checkForSmudge(puzzle[before], puzzle[after]) && !smudgeAlreadyFound {
				before--
				after++
				smudgeAlreadyFound = true
				// We found an edge
				if before < 0 || after > len(puzzle)-1 {
					return true
				}
			} else if puzzle[before] == puzzle[after] {
				before--
				after++
				// We found an edge
				if before < 0 || after > len(puzzle)-1 {
					return true
				}

			} else {
				return false
			}

		} else {
			break
		}
	}
	return false

}

