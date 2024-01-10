package day3

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	// "strconv"
	"golang.org/x/exp/slices"
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
				if symbol > 0 && symbol < lenArray {
					for _, line := range sliceFullText {
						line = line[symbol-1 : symbol+2]
						sliceSquare = append(sliceSquare, line)
					}
				}

				// will always send the line up and down (no need to send the line) because it will be always in the middle
				numsAroundSymbol(symbol, sliceSquare, sliceFullText) // column: symbol 
			}
		}
	}

	for _, num := range partNumbers {
		accumulator += num
	}

	return accumulator
}

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
	potentialNums := []int{}

	for i:= 0; i <3; i++ { // Line
		for j:= 0; j <3; j++{ // Cols
			if i != 1 && j != 1{ // Insure that we are not re analysing the symbole in the middle 
				is_num := is_Number(map3by3[i][j]) // Make sure we found a number and not an other symbole
				if is_num {
					number := checkFullNumber(mapText,i , col + (j-1)) //  col + (j-1) re caliber the y into the full length array 
					potentialNums = append(potentialNums, number)

				}
			}
		}
	}

	return potentialNums
}

func is_Number(r rune) bool {

	is_int := unicode.IsDigit(r)
	if is_int {
		return true
	}
	return false
}

func checkFullNumber(mapText [][]rune, line int, col int) int {
	// We have found a potential number now we want to map this number by checking bedore and after it
	// if there is other number 
	// Then we need to found a signature to make sure it's unique and we are not checking it multiple time



}

// func convertStrtoInt(s string) int {
// 	i, err := strconv.Atoi(s)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return i
// }
