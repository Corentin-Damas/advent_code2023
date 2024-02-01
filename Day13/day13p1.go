package day13

import (
	"fmt"
	"log"
	"os"
	"strings"
	// "golang.org/x/exp/slices"
)

func Day13p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day13/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	puzzle := string(bytesText)
	result := analyzep1(puzzle)
	fmt.Println("Day 13 part 1 result is : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r")
	}
	allPuzzles := getAllPuzzles(puzzleLines)

	total := 0
	for _, puzzle := range allPuzzles {
		idxHor, trueMirorHor := horizontalAnalyze(puzzle)
		if trueMirorHor {
			total += (idxHor * 100)
		} else {
			idxHor = 0
		}

		swapedPuzzle := swap(puzzle)
		idxVer, trueMiror := horizontalAnalyze(swapedPuzzle)
		if trueMiror {
			total += idxVer
		} else {
			idxVer = 0
		}
		fmt.Println(idxHor, idxVer)
	}
	return total
}

func horizontalAnalyze(puzzle []string) (int, bool) {
	for idx := 0; idx < len(puzzle)-1; idx++ {
		// fmt.Println(idx, puzzle[idx], idx+1, puzzle[idx+1], "max: ", len(puzzle)-1)

		if puzzle[idx] == puzzle[idx+1] {

			is_true_miror := is_RealMiror(idx, puzzle)

			if is_true_miror {
				return idx + 1, is_true_miror
			} else {
				continue
			}
		}
		if idx == len(puzzle)-1 {
			break
		}
	}
	return 0, false
}

func is_RealMiror(sep int, puzzle []string) bool {
	before := sep
	after := sep + 1
	for {
		if after <= len(puzzle)-1 {

			if puzzle[before] == puzzle[after] {
				before--
				after++
				// We found an edge
				if before < 0 || after > len(puzzle)-1 {
					return true
				}
			} else {
				break
			}
		} else {
			break
		}
	}
	return false

}

func swap(puzzle []string) []string {
	swapPuz := []string{}
	for idxL, line := range puzzle {
		for idxC, ch := range line {
			if idxL == 0 {
				swapPuz = append(swapPuz, string(""))
			}
			swapPuz[idxC] = swapPuz[idxC] + string(ch)
		}
	}
	for _, line := range swapPuz {
		strings.Trim(line, "\r")
		strings.Trim(line, "")
		strings.Trim(line, "\n")
	}

	return swapPuz
}

func getAllPuzzles(puzzleLines []string) [][]string {

	puzzleMap := [][]string{}
	mapCounter := 0
	puzzleStart := 0
	for i := 0; i < len(puzzleLines)-1; i++ {
		if strings.Contains(puzzleLines[i], string("#")) || strings.Contains(puzzleLines[i], ".") {
			continue
		} else {
			puzzleMap = append(puzzleMap, []string{})

			for ; puzzleStart < i; puzzleStart++ {
				puzzleMap[mapCounter] = append(puzzleMap[mapCounter], puzzleLines[puzzleStart])
			}
		}
		puzzleStart += 1
		mapCounter += 1
	}

	return puzzleMap
}
