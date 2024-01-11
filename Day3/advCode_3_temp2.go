package day3

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"strconv"

	"github.com/ydb-platform/ydb-go-sdk/v3/topic"
	"golang.org/x/exp/slices"
)

// Any number adjacent to a symbol, even diagonally, is a "part number"
// and should be included in your sum. (Periods (.) do not count as a symbol.)
// add up all the part numbers

func Day3temp2() int {

	bytesText, err := os.ReadFile("./Day3/datatest.txt")
	// bytesText, err := os.ReadFile("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	FullText := string(bytesText)
	result := analyseText(FullText)

	fmt.Println("Day 3 result is : ")
	return result
}

func analyze(s string) int {
	arrFullText := [][]rune{}

	var partNumbers []int
	var numString string
	var isNoTop bool
	var isNoBot bool
	var isNoLeft bool
	var isNoRight bool
	
	accumulator := 0

	lines := strings.Split(s, "\n")
	// Decompose the text into an array or array to have a 2D representation of the puzzle
	arrFullText = puzzleMapToArr(lines)

	for idxLine, linearr := range arrFullText {
		for idxCol, potentialNum := range linearr {
			found_number := rune_is_Number(potentialNum)

			// If symbole turn "accepted number one" 
			// if next rune is also a number check it too 

			if found_number{
				numString += string(arrFullText[idxLine][idxCol])

				is_top := line_exist(arrFullText[idxLine-1])
				if is_top != nil {
					isNoTop = true
				}
				is_bot := line_exist(arrFullText[idxLine+1])
				if is_bot != nil {
					isNoBot = true
				}

				is_left := col_exist(arrFullText[idxLine][idxCol-1])
				if is_left != nil {
					isNoLeft = true
				}
				is_right := col_exist(arrFullText[idxLine][idxCol+1])
				if is_right != nil {
					isNoRight = true
				}
			
				// Check all


				// Need to check the surronding of that rune

			}
			


		}

		for _, num := range partNumbers {
			accumulator += num
		}

	}
	return accumulator
}

// Transfmor the lines of string in Arrays of runes 2D
func puzzleMapToArr(lines []string) [][]rune {
	result := [][]rune{}
	for _, line := range lines {
		arr := strToarray(line)
		result = append(result, arr)
	}
	return result
}

// Take one line of string into runes and make sure that the \n is not counted rune: 13
func strToarray(s string) []rune {
	arr := []rune{}
	for _, char := range s {
		if char == 13 {
			continue
		}
		arr = append(arr, char)
	}
	return arr
}

// Check if a rune is a number 
func rune_is_Number(r rune) bool {

	is_int := unicode.IsDigit(r)
	return is_int

}

// checkFullNum Give back 
func checkFullNum(arr []rune, start int) (int, int) {
	
	var numString string
	var idxNumEnd int

	// Get the full number into a string (easier to concat)
	for j := start; j <= len(arr)-1; j++ {
		if is_Number(arr[j]) {
			numString += string(arr[j])
			idxNumEnd = j
		} else {
			break
		}
	}
	result := convertStrtoInt(numString)
	 
	return result, idxNumEnd
}

func line_exist(arr []rune) error{
	if len(arr) > 0 {
		return nil
	}else {
		return fmt.Errorf("Out of bound: Line")
	}
} 
func col_exist(r rune) error{
	if rune_is_Number(r) {
		return nil
	}else {
		return fmt.Errorf("Out of bound: Col")
	}
} 




// numsAroundSymbol recive a 9x9 square with the symbole in the middle
// it also receive the line up and below
func numsAroundSymbol(col int, map3by3 [][]rune, mapText [][]rune) []int {
	// var tempMap map[int]int
	potentialNums := []int{}

	var number int

	for i := 0; i < 3; i++ { // Line

		for j := 0; j < 3; j++ { // Cols
			if is_Number(map3by3[i][j]) {
				number = checkFullNumber(mapText[i], col+(j-1)) //  col + (j-1) re caliber the y into the full length array

				// Need to found a way to identify a number so that a number is not count 2time and
				// Give the posibility to have 2 same number around the same symbole
				is_already_found := numInSlice(number, potentialNums)

				if !is_already_found {
					potentialNums = append(potentialNums, number)
				}

			}

		}
	}
	return potentialNums
}

// checkFullNumber Receive an array of Rune with a starting possition, return the full number

func numInSlice(num int, list []int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}

func convertStrtoInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

// should found 540887
