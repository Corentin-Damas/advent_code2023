package day14

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Day14p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day14/datatest.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 14 part 1 : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}

	puzzleArr := lineToAtt(puzzleLines)

	puzzleResult := rollTherock(puzzleArr)

	
	printPuzzle(puzzleResult)
	return resultPuzzle(puzzleResult)
}

func lineToAtt(puzzle []string) [][]string {
	puzzleTranspose := [][]string{}

	for idx, line := range puzzle {
		puzzleTranspose = append(puzzleTranspose, []string{})

		for _, c := range line {
			puzzleTranspose[idx] = append(puzzleTranspose[idx], string(c))
		}
	}
	return puzzleTranspose
}

func printPuzzle(puzzle [][]string){
	for idx , line := range puzzle {
		fmt.Println(line, idx)
	}
	fmt.Println("")
}
func resultPuzzle(puzzle [][]string) int{
	total := 0

	for line := 0 ; line <= len(puzzle)-1 ; line ++ {
		for _, ch := range puzzle[line]{
			if ch == "O"{
				total += (len(puzzle) -line )
			}
		} 
	}
	return total
}


func rollTherock(puzzle [][]string)[][]string{
	puzzleArr := puzzle
	lenPuzzle := len(puzzleArr) - 1
	lenLinePuzzle := len(puzzleArr[0]) - 1

	for reverseIdxLine := lenPuzzle; reverseIdxLine > 0; reverseIdxLine-- {

		for ch := 0; ch <= lenLinePuzzle; ch++ {

			if puzzleArr[reverseIdxLine][ch] == "O" {

				// We are on the very bottom line no need to check bellow
				if reverseIdxLine == lenPuzzle && puzzleArr[reverseIdxLine-1][ch] == "." {
					temp := puzzleArr[reverseIdxLine][ch]
					puzzleArr[reverseIdxLine][ch] = puzzleArr[reverseIdxLine-1][ch]
					puzzleArr[reverseIdxLine-1][ch] = temp
					
				} else if puzzleArr[reverseIdxLine-1][ch] == "." && puzzleArr[reverseIdxLine+1][ch] == "O" {

					nextBottomLine := reverseIdxLine
					edge := false

					// Found the last "O" bellow
					if  reverseIdxLine + 1 == lenPuzzle{
						temp := puzzleArr[reverseIdxLine + 1][ch]
						puzzleArr[reverseIdxLine+ 1 ][ch] = puzzleArr[reverseIdxLine-1][ch]
						puzzleArr[reverseIdxLine-1][ch] = temp

					} else {
						for nextBottomLine <= lenPuzzle && !edge {
							nextBottomLine++
							
							if puzzleArr[nextBottomLine][ch] != "O" || nextBottomLine == lenPuzzle {
								edge = true
							}
						}
						// edge is found && next nextBottomLine is = the last O of the serie
						// Bubble up all O untile reverseIdxLine-1
						temp := puzzleArr[nextBottomLine -1][ch]
						puzzleArr[nextBottomLine -1][ch] = puzzleArr[reverseIdxLine-1][ch]
						puzzleArr[reverseIdxLine-1][ch] = temp
					}

				} else if puzzleArr[reverseIdxLine-1][ch] == "." {
					// We are sure that no O is on top so we just move one O upward if the top is a .
					temp := puzzleArr[reverseIdxLine][ch]
					puzzleArr[reverseIdxLine][ch] = puzzleArr[reverseIdxLine-1][ch]
					puzzleArr[reverseIdxLine-1][ch] = temp

				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
	return puzzleArr
}