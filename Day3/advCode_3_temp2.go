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

type Bounds struct {
	isNoTop   bool
	isNoBot   bool
	isNoLeft  bool
	isNoRight bool
}

func Day3temp2() int {

	bytesText, err := os.ReadFile("./Day3/datatest.txt")
	// bytesText, err := os.ReadFile("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	FullText := string(bytesText)
	result := analyze(FullText)

	fmt.Println("Day 3 result is : ")
	return result
}

func analyze(s string) int {

	var partNumbers []int
	var is_partNumber bool

	directions := map[string][]int{
		"topleft":  {-1, -1},
		"topmid":   {-1, 0},
		"topright": {-1, 1},
		"midleft":  {0, -1},
		"midright": {0, 1},
		"botleft":  {1, -1},
		"botmid":   {1, 0},
		"botright": {1, 1},
	}

	mapbound := Bounds{
		isNoTop:   false,
		isNoBot:   false,
		isNoLeft:  false,
		isNoRight: false,
	}

	pMapBound := &mapbound

	accumulator := 0

	lines := strings.Split(s, "\n")
	// Decompose the text into an array or array to have a 2D representation of the puzzle
	arrFullText := puzzleMapToArr(lines)
	pArrFullText := &arrFullText

	var accStr string

	for idxLine, linearr := range arrFullText {
		for idxCol, potentialNum := range linearr {
			// Init Reset Part
			is_partNumber = false
			mapbound.isNoTop = false
			mapbound.isNoBot = false
			mapbound.isNoLeft = false
			mapbound.isNoRight = false

			fmt.Println("current numbers; ", partNumbers)

			found_number := rune_is_Number(potentialNum)

			// If symbole turn "accepted number one"
			// if next rune is also a number check it too

			if found_number {

				accStr += string(potentialNum)

				fmt.Println(accStr)

				if idxLine == 0 {
					mapbound.isNoTop = true
				}

				if idxLine == len(arrFullText)-1 {
					mapbound.isNoBot = true
				}

				if idxCol == 0 {
					mapbound.isNoLeft = true
				}

				if idxCol == len(linearr)-1 {
					mapbound.isNoRight = true
				}

			} else {
				continue
			}

			if !is_partNumber {
				fmt.Println("checking for partNumber", is_partNumber)
				is_partNumber = check9around(pArrFullText, idxLine, idxCol, directions, pMapBound)
			}

			is_nextCharNum := checkNextChar(pArrFullText, idxLine, idxCol)

			//FIXEME :  Bug here numbers are not splitting at the right moment

			if is_nextCharNum {
				// there is a next num
				continue
			}

			// next Char is NOT a num

			if is_partNumber {
				// End of the Num -> hand the full number to the list of valid num
				partNumbers = append(partNumbers, strtoInt(accStr))
			} else {
				// End of the Num && is not a part Number -> reset params
				accStr = ""
				is_partNumber = false
			}

		}

	}
	for _, num := range partNumbers {
		accumulator += num
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

func rune_is_Symbole(r rune) bool {
	notSymbole := "0123456789."
	var checkChar = []rune{}
	for _, char := range notSymbole {
		checkChar = append(checkChar, char)
	}

	// True = the rune is or a num or a dot || False = symbole
	is_num_or_dot := slices.Contains(checkChar, r)

	// We want to return true if a symbole is found
	return !is_num_or_dot

}

// it also receive the line up and below
func check9around(pFullMap *[][]rune, line int, col int, directions map[string][]int, bound *Bounds) bool {

	is_symbole := true
	fmt.Println("currently checking: ", string((*pFullMap)[line][col]))

	// directions := map[string][]int{
	// 	"topleft":  {-1, -1},
	// 	"topmid":   {-1, 0},
	// 	"topright": {-1, 1},
	// 	"midleft":  {0, -1},
	// 	"midright": {0, 1},
	// 	"botleft":  {1, -1},
	// 	"botmid":   {1, 0},
	// 	"botright": {1, 1},
	// }

	for key, dir := range directions {
		i := dir[0]
		j := dir[1]

		if bound.isNoTop && (key == "topleft" || key == "topmid" || key == "topright") {
			continue
		}
		if bound.isNoBot && (key == "botleft" || key == "botmid" || key == "botright") {
			continue
		}
		if bound.isNoLeft && (key == "botleft" || key == "midleft" || key == "topleft") {
			continue
		}
		if bound.isNoRight && (key == "topright" || key == "midright" || key == "botright") {
			continue
		}

		potentialRune := (*pFullMap)[line+i][col+j]

		if rune_is_Symbole(potentialRune) {
			fmt.Println("potential symbole ", potentialRune)
			return is_symbole
		}
	}
	return !is_symbole
}

func checkNextChar(pFullMap *[][]rune, line int, col int) bool {

	nextChar := (*pFullMap)[line][col+1]
	return rune_is_Number(nextChar)
}

func strtoInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

// checkFullNumber Receive an array of Rune with a starting possition, return the full number

// should found 540887
