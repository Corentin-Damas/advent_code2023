package day13

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Day13p2() int {

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

func analyzep2(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r")
	}
	allPuzzles := getAllPuzzles(puzzleLines)

	total := 0
	for _, puzzle := range allPuzzles {
		idxHor, trueMirorHor, smudgeFoundhr := horizontalAnalyzep2(puzzle)

		swapedPuzzle := swap(puzzle)
		idxVer, trueMirorVer, smudgeFoundve := horizontalAnalyzep2(swapedPuzzle)
		if trueMirorHor && smudgeFoundhr{
			total += (idxHor * 100)
		}else {
			idxHor = 0
		}
		
		if trueMirorVer && smudgeFoundve {
			total += idxVer
		} else {
			idxVer = 0
		}


		if idxHor == 0 && idxVer == 0{
			idxHor2, trueMirorHorp1 := horizontalAnalyze(puzzle)
			if trueMirorHorp1 {
				total += (idxHor2 * 100)
			} else {
				idxHor2 = 0
			}

			idxVer2, trueMiror2 := horizontalAnalyze(swapedPuzzle)
			if trueMiror2 {
				total += idxVer2
			} else {
				idxVer2 = 0
			}

			fmt.Println(idxHor2 , idxVer2)

		} else{

			fmt.Println(idxHor , idxVer)
		}

	}
	return total
}

func horizontalAnalyzep2(puzzle []string) (int, bool, bool) {

	for idx := 0; idx < len(puzzle)-1; idx++ {
		// fmt.Println(idx, puzzle[idx], idx+1, puzzle[idx+1], "max: ", len(puzzle)-1)

		if checkForSmudge(puzzle[idx], puzzle[idx+1]){

			is_smudgeMiror, smudgeFound := is_RealMirorp2(idx, puzzle, true)

			if is_smudgeMiror {
				return idx + 1, is_smudgeMiror ,smudgeFound
			} else {
				continue
			}
		}

		if puzzle[idx] == puzzle[idx+1] {

			is_smudgeMiror, smudgeFound := is_RealMirorp2(idx, puzzle, false)

			if is_smudgeMiror {
				return idx + 1, is_smudgeMiror , smudgeFound
			} else {
				continue
			}
		}
		if idx == len(puzzle)-1 {
			break
		}
	}
	return 0, false, false
}

func is_RealMirorp2(sep int, puzzle []string, isAlreadyFound bool) (bool, bool) {
	before := sep
	after := sep + 1
	smudgeAlreadyFound := isAlreadyFound

	for {
		if after <= len(puzzle)-1 {

			if checkForSmudge(puzzle[before], puzzle[after]) && !smudgeAlreadyFound{
				before--
				after++
				smudgeAlreadyFound = true
				// We found an edge
				if (before < 0 || after > len(puzzle)-1)  {
					return true, smudgeAlreadyFound
				}
			} else if puzzle[before] == puzzle[after] {
				before--
				after++
				// We found an edge
				if (before < 0 || after > len(puzzle)-1)  {
					return true, smudgeAlreadyFound
				}
				
			} else {
				return false, smudgeAlreadyFound
			}

		} else {
			break
		}
	}
	return false, smudgeAlreadyFound

}

// Every Mirror has exaclty one SMUDGE one . or # should be reversed
// Locate and fix the smudge that cause a different reflection line to be valid
// The old reflection line won't necessarily continue being valid after the smudge is fixed

func checkForSmudge(s1 string, s2 string) bool {
	
	count := 0
	for i := 0; i <= len(s1)-1; i++ {
		if s1[i] != s2[i] {
			count++
		}
	}

	if count == 1 {
		return true
		
	} 
	return false

}