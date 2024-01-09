package day3

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Any number adjacent to a symbol, even diagonally, is a "part number"
// and should be included in your sum. (Periods (.) do not count as a symbol.)
// add up all the part numbers

func Day3result() int {

	bytesText, err := os.ReadFile("./Day3/datatest.txt")
	if err != nil {
		log.Fatal(err)
	}
	FullText := string(bytesText)
	result := analyseText(FullText)

	fmt.Println("Day 3 result is : ")
	return result
}

func analyseText(s string) int {
	var partNumbers []int
	accumulator := 0
	arrFullText := [][]rune{}

	lines := strings.Split(s, "\n")

	// Decompose the text into an array or array to have a 2D representation of the puzzle
	for _, line := range lines {
		arr := lineToarray(line)
		arrFullText = append(arrFullText, arr)
	}

	// Analyse line by line, if a symbole is found look at the 9 directions if a num is there
	for idx, arr := range arrFullText {

		// symboleLoc get back an []int with the idx of the symboles in a line // Colomn idx
		symboleArrLoc := symboleLoc(arr)
		if len(symboleArrLoc) != 0 {
			for _, symbol := range symboleArrLoc {

				sliceFullText := [][]rune{}

				if idx > 0 {
					sliceFullText = arrFullText[idx-1 : idx+1]
				}

				if idx == 0 {
					sliceFullText = arrFullText[idx : idx+1]
				}
				if idx == len(arrFullText)-1 {
					sliceFullText = arrFullText[idx-1 : idx]
				}
				numVicinity := numsAroundSymbol(idx, symbol, sliceFullText) // line: idx, column: symbol
				println(numVicinity)
			}
		}
	}

	for _, num := range partNumbers {
		accumulator += num
	}

	return accumulator
}

func lineToarray(s string) []rune {
	arr := []rune{}

	for _, char := range s {
		arr = append(arr, char)
	}
	fmt.Println(arr)

	return arr
}

func symboleLoc(arr []rune) []int {
	passChar := "0123456789."
	symbolLoc := []int{}

	for idx, char := range arr {
		for _, p := range passChar {
			if char != p {
				symbolLoc = append(symbolLoc, idx)

			}
		}
	}
	return symbolLoc
}

// numsAroundSymbol recive a coordniate of a symbole we should check all 9 location around
// But also take care if it's in a corner
func numsAroundSymbol(line int, col int, mapText [][]rune) []int {
	potentialNums := []int{}
	lineTocheck := len(mapText)

	if lineTocheck == 2 && col != len(mapText[0]) && col != 0 { // Proced normally
		for x := 0; x >= 1; x-- { // Line
			for y := -1; y <= 1; y-- { // Cols
				err, locX, locY := is_Number(mapText[line+x][col+y], line+x, col+y)
			}
		}

	}

	if lineTocheck == 3 && col != len(mapText[0]) && col != 0 { // // Proced normally
		for x := -1; x >= 1; x++ { // Line
			for y := -1; y <= 1; y-- { // Cols
				err, locX, locY := is_Number(mapText[line+x][col+y], line+x, col+y)
				if err == nil {
					number := checkFullNumber(mapText, locX, locY)
					potentialNums = append(potentialNums, number)
				}
			}
		}
	}

	return potentialNums
}

func is_Number(r rune, line int, col int) (error, int, int) {
	is_int := unicode.IsDigit(r)
	if is_int {
		return nil, line, col
	} else {
		return errors.New("not a num"), 0, 0
	}

}

func checkFullNumber(mapText [][]rune, line int, col int) int {

}

func convertStrtoInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
