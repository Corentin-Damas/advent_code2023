package day3part2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

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

// type SpeGear struct {
// 	specialGears
// }

func Day3part2() int {

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
	var is_partNumber bool = false
	var is_nextCharNum bool
	var accumulator int = 0
	var gear_acc int = 0
	var accStr string
	var speGear map[string][]int
	pSpeGear := &speGear

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

	// Decompose the text into an array or array to have a 2D representation of the puzzle
	lines := strings.Split(s, "\n")
	arrFullText := puzzleMapToArr(lines)
	pArrFullText := &arrFullText

	for idxLine, linearr := range arrFullText {
		for idxCol, potentialNum := range linearr {
			// Init & Reset Part
			is_nextCharNum = false

			found_number := rune_is_Number(potentialNum)
			if found_number {
				accStr += string(potentialNum)
				defineBounds(pMapBound, idxLine, len(arrFullText)-1, idxCol, len(linearr)-1) // Define the bound for the rune we are checking
			} else {
				continue
			}

			if !is_partNumber {
				is_partNumber = check9around(pArrFullText, idxLine, idxCol, directions, pMapBound, pSpeGear)
			}

			if idxCol < len(linearr)-1 {
				is_nextCharNum = checkNextChar(pArrFullText, idxLine, idxCol)
			}

			if is_nextCharNum {
				// there is a next num
				continue
			}

			// next Char is NOT a num
			if is_partNumber && !is_nextCharNum {
				// End of the Num -> hand the full number to the list of valid num
				partNumbers = append(partNumbers, strtoInt(accStr))
			}
			// End of the Num || is not a part Number -> reset params
			accStr = ""
			is_partNumber = false
		}

	}
	for _, num := range partNumbers {
		// fmt.Println(num)
		accumulator += num
	}

	fmt.Println(speGear)

	for _, tuple := range speGear {
		fmt.Println(tuple[0] * tuple[1])
		gear_acc += (tuple[0] * tuple[1])
	}

	// Count everything as normal gear, create a map of special gear, deduce all this number from (a,b)
	// Multiply all the tuples and add to the final result.
	// Map[line int]tuple(num1, num2)
	return gear_acc

}

// ==========================   Convertions   ==========================

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

// Take a String and return an int
func strtoInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
func intToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

// ==========================   Rune Check   ==========================

// Check if a rune is a number
func rune_is_Number(r rune) bool {
	return unicode.IsDigit(r)
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

func checkNextChar(pFullMap *[][]rune, line int, col int) bool {

	nextChar := (*pFullMap)[line][col+1]
	return rune_is_Number(nextChar)
}

// ==========================   Mapping   ==========================

func check9around(pFullMap *[][]rune, line int, col int, directions map[string][]int, bound *Bounds, speGear *map[string][]int) bool {

	is_symbole := true
	for key, dir := range directions {
		i := dir[0]
		j := dir[1]

		if checkBound(key, bound) {
			continue
		}

		potentialRune := (*pFullMap)[line+i][col+j]
		is_sym := rune_is_Symbole(potentialRune)
		if is_sym && potentialRune != 42 {
			return is_symbole
		}

		if is_sym && potentialRune == 42 {
			gearArr, key_arr := checkAroundStar(pFullMap, line+i, col+j, directions)

			if len(gearArr) != 2 {
				return is_symbole
			}

			new_tuple(speGear, gearArr, key_arr)
		}
	}

	return !is_symbole
}

func checkAroundStar(pFullMap *[][]rune, line int, col int, directions map[string][]int) ([]int, string) {

	gearArr := []int{}
	key_arr := []int{}
	orderedKey := []int{}
	maxLine := len((*pFullMap))
	maxCol := len((*pFullMap)[line])

	starBound := Bounds{}
	pStarBound := &starBound

	defineBounds(pStarBound, line, maxLine, col, maxCol) // define the bound fo

	for key, dir := range directions {
		i := dir[0]
		j := dir[1]

		if checkBound(key, pStarBound) {
			// If a bound is find will pass the value
			continue
		}

		r := (*pFullMap)[line+i][col+j]
		is_num := rune_is_Number(r)

		if is_num {
			f_num, num_key := full_number(pFullMap, line+i, col+j)

			if !slices.Contains(key_arr, num_key) {
				key_arr = append(key_arr, num_key)
			}

			gearArr = append(gearArr, f_num)
		}
	}
	if len(key_arr) == 2 {
		orderedKey = order2Key(key_arr)
	}
	strKey := intToStr(orderedKey[0]) + "_" + intToStr(orderedKey[1])
	fmt.Println("found new tuple; ", gearArr, strKey)
	return gearArr, strKey
}

func full_number(pFullMap *[][]rune, line int, start_col int) (int, int) {

	var numStart int
	var numString string
	var key string

	len_arr := len((*pFullMap)[line]) - 1

	// Find the start of the number
	for i := start_col; i >= 0; i-- {
		if rune_is_Number((*pFullMap)[line][i]) {
			numStart = i
		} else {
			break
		}
	}
	// Get the full number into a string (easier to concat)
	for j := numStart; j <= len_arr; j++ {
		if rune_is_Number((*pFullMap)[line][j]) {
			numString += string((*pFullMap)[line][j])
			key = intToStr(line) + intToStr(numStart) + intToStr(j)
		} else {
			break
		}
	}
	result := strtoInt(numString)
	i_key := strtoInt(key)

	return result, i_key
}

func new_tuple(speGear *map[string][]int, tuple []int, key string) {
	for _, arr := range *speGear {
		if slices.Contains(arr, tuple[0]) && slices.Contains(arr, tuple[1]) {
			return
		}
	}

	(*speGear)[key] = tuple
}

func order2Key(arr []int) []int {
	if arr[0] > arr[1] {
		temp := arr[1]
		arr[1] = arr[0]
		arr[0] = temp
	}
	return arr
}

// ==========================   Bouderies   ==========================

func defineBounds(bound *Bounds, line int, max_Lines int, col int, max_Cols int) {

	if line == 0 {
		bound.isNoTop = true
	} else {
		bound.isNoTop = false
	}

	if line == max_Lines {
		bound.isNoBot = true
	} else {
		bound.isNoBot = false
	}

	if col == 0 {
		bound.isNoLeft = true
	} else {
		bound.isNoLeft = false
	}

	if col == max_Cols {
		bound.isNoRight = true
	} else {
		bound.isNoRight = false
	}
}

func checkBound(key string, bound *Bounds) bool {
	if bound.isNoTop && (key == "topleft" || key == "topmid" || key == "topright") {
		return true
	}
	if bound.isNoBot && (key == "botleft" || key == "botmid" || key == "botright") {
		return true
	}
	if bound.isNoLeft && (key == "botleft" || key == "midleft" || key == "topleft") {
		return true
	}
	if bound.isNoRight && (key == "topright" || key == "midright" || key == "botright") {
		return true
	}
	return false
}
