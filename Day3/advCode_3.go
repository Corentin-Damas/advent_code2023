package day3

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"strconv"

	"golang.org/x/exp/slices"
)

// Any number adjacent to a symbol, even diagonally, is a "part number"
// and should be included in your sum. (Periods (.) do not count as a symbol.)
// add up all the part numbers

func Day3result() int {

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

func analyseText(s string) int {
	var partNumbers []int
	var symboleArrLoc []int
	var lenArray int
	accumulator := 0
	arrFullText := [][]rune{}

	lines := strings.Split(s, "\n")

	// Decompose the text into an array or array to have a 2D representation of the puzzle
	for _, line := range lines {
		arr, lengthArray := lineToarray(line)
		lenArray = lengthArray
		arrFullText = append(arrFullText, arr)
	}

	var dummySlice = []rune{}
	for i := 0; i < lenArray; i++ {
		dummySlice = append(dummySlice, 46)
	}

	// fmt.Println("full text:")
	// fmt.Println(arrFullText)

	// Analyse line by line, if a symbole is found look at the 9 directions if a num is there
	for idx, linearr := range arrFullText {
		// symboleLoc get back an []int with the idx of the symboles in a line // Colomn idx
		symboleArrLoc = symboleLoc(linearr)

		if len(symboleArrLoc) != 0 {

			for _, symbol := range symboleArrLoc {

				sliceFullText := [][]rune{}
				sliceSquare := [][]rune{}

				// Line slice
				if idx > 0 && idx < len(arrFullText)-1 {
					sliceFullText = arrFullText[idx-1 : idx+2]
				}

				if idx == 0 {

					TempSlice := arrFullText[idx : idx+2]
					firstLine := TempSlice[0]
					secondLine := TempSlice[1]
					sliceFullText = [][]rune{}
					sliceFullText = append(sliceFullText, dummySlice)
					sliceFullText = append(sliceFullText, firstLine)
					sliceFullText = append(sliceFullText, secondLine)
				}
				if idx == len(arrFullText)-1 {
					TempSlice := arrFullText[idx-1 : idx+1]
					firstLine := TempSlice[0]
					secondLine := TempSlice[1]
					sliceFullText = [][]rune{}
					sliceFullText = append(sliceFullText, firstLine)
					sliceFullText = append(sliceFullText, secondLine)
					sliceFullText = append(sliceFullText, dummySlice)
				}

				// col Slice
				if symbol > 0 && symbol <= lenArray-1 {
					for _, line := range sliceFullText {
						line = line[symbol-1 : symbol+2]
						sliceSquare = append(sliceSquare, line)
					}
				}
				if symbol == 0 {
					for _, line := range sliceFullText {
						subline := line[symbol : symbol+2]
						var assemble = []rune{46}
						for _, char := range subline {
							assemble = append(assemble, char)
						}
						sliceSquare = append(sliceSquare, assemble)
					}
				}

				// will always send the line up and down (no need to send the line) because it will be always in the middle
				numberAround := numsAroundSymbol(symbol, sliceSquare, sliceFullText) // column: symbol
				partNumbers = append(partNumbers, numberAround...)
			}

			// Then we need to found a signature to make sure it's unique and we are not checking it multiple time

		}
	}

	for _, num := range partNumbers {
		accumulator += num
	}

	return accumulator
}

// Correct
func lineToarray(s string) ([]rune, int) {
	arr := []rune{}

	for _, char := range s {
		if char == 13 {
			continue
		}
		arr = append(arr, char)
	}
	lenArray := len(arr)

	return arr, lenArray
}

// Correct
func symboleLoc(arr []rune) []int {
	passChar := "0123456789."
	var checkChar = []rune{}
	for _, c := range passChar {
		checkChar = append(checkChar, c)
	}

	symbolLoc := []int{}

	for idx, char := range arr {
		if !slices.Contains(checkChar, char) {
			symbolLoc = append(symbolLoc, idx)
		}
	}
	return symbolLoc
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

func is_Number(r rune) bool {

	is_int := unicode.IsDigit(r)
	return is_int

}

// checkFullNumber Receive an array of Rune with a starting possition, return the full number
func checkFullNumber(mapText []rune, col int) int {
	// We have found a potential number now we want to map this number by checking bedore and after it
	// if there is other number
	var numStart int
	var numString string
	// Find the start of the number
	for i := col; i >= 0; i-- {
		if is_Number(mapText[i]) {
			numStart = i
		} else {
			break
		}
	}
	// Get the full number into a string (easier to concat)
	for j := numStart; j <= len(mapText)-1; j++ {
		if is_Number(mapText[j]) {
			numString += string(mapText[j])
		} else {
			break
		}
	}
	result := convertStrtoInt(numString)
	return result
}

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
