package day12

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

// https://www.youtube.com/watch?v=g3Ms5e7Jdqo

func Day12p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day12/datatest.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 12 part 1 result is : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}

	total := 0

	for _, line := range puzzleLines {

		puzzleInputLine, instructions := getLineInfo(line)
		fmt.Printf(" -> %s, %+v \n", puzzleInputLine, instructions)
		total += solutionsPuzzles(puzzleInputLine, instructions)

	}

	return total
}

// func solutionsPuzzles(puzzle []string, instruction []int) int {

// }

func solutionsPuzzles(puzzle []string, instruction []int) int {

	// Base Returns
	if len(puzzle) == 0 {
		if len(instruction) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(instruction) == 0 {
		if slices.Contains(puzzle, "#") {
			return 0
		} else {
			return 1
		}
	}
	count := 0

	// Recursion 1 where ? is = .
	if puzzle[0] == "." || puzzle[0] == "?" {
		if len(puzzle) > 1 {
			count += solutionsPuzzles(puzzle[1:], instruction)
		} else if len(puzzle) == 1 {
			count += solutionsPuzzles([]string{}, instruction)
		} else {
			count += solutionsPuzzles([]string{puzzle[1]}, instruction)
		}

	}

	// Recursion 2 where ? is = #
	if puzzle[0] == "#" || puzzle[0] == "?" {

		isEnougthGearLeft := instruction[0] <= len(puzzle)

		if isEnougthGearLeft && !slices.Contains(puzzle[:instruction[0]], ".") && (instruction[0] == len(puzzle) || puzzle[instruction[0]] != "#") {
			fmt.Println("Curently checking", puzzle, instruction)

			if len(instruction) == 1 && instruction[0] == len(puzzle) {
				count += solutionsPuzzles([]string{}, []int{})

			} else if len(puzzle)-len(puzzle[instruction[0]:]) == 1 {
				count += solutionsPuzzles([]string{puzzle[len(puzzle)-1]}, []int{})
			} else if len(instruction) == 1 {
				count += solutionsPuzzles(puzzle[instruction[0]+1:], []int{})

			}
			count += solutionsPuzzles(puzzle[instruction[0]+1:], instruction[1:])
			// puzzle[instruction[0]+1:] if 3 instructed spring where already analyse goes to the next valide place
		} else {
			count += 0
		}

	}

	return count
}

func getLineInfo(s string) ([]string, []int) {
	spaceDestruct := strings.Split(s, " ")

	puzzle := stringToArr(spaceDestruct[0])
	instructionLine := stringToArrNums(spaceDestruct[1])

	return puzzle, instructionLine
}

func stringToArr(s string) []string {

	arr := []string{}
	for _, c := range s {
		arr = append(arr, string(c))
	}
	return arr
}
func stringToArrNums(s string) []int {

	numsStr := strings.Split(s, ",")
	arr := []int{}
	for _, c := range numsStr {

		i, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		arr = append(arr, i)
	}
	return arr
}
