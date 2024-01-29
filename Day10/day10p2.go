package day10

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func Day10p2() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day10/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep2(puzzle)

	fmt.Println("Day 10 part 1 result is : ")
	return result
}

func analyzep2(s string) int {
	puzzleLines := strings.Split(s, "\n")
	pPuzzeLines := &puzzleLines

	startingPos, startingsLoop := getStartingPos(pPuzzeLines)

	regexUsedPipe := RegPipe(startingPos, startingsLoop, pPuzzeLines)

	cleanLines := cleanLines(pPuzzeLines, regexUsedPipe)

	for _, line := range cleanLines {
		fmt.Println(line)
	}

	nbOutsiders := getOusiders(cleanLines, regexUsedPipe)

	return nbOutsiders
}

func RegPipe(startPos [2]int, StartLoop [2]int, puzzleLines *[]string) [][2]int {

	registPipUsed := [][2]int{}
	registPipUsed = append(registPipUsed, startPos)

	nextPos := StartLoop
	actuPos := startPos

	fmt.Printf("actual starting pose %v , next pipe %v \n", actuPos, nextPos)

	cardinals := CardinalPoint{
		north: [2]int{0, -1},
		south: [2]int{0, 1},
		west:  [2]int{-1, 0},
		east:  [2]int{1, 0},
	}

	pipesInstructions := map[string][2][2]int{
		"|": [2][2]int{cardinals.north, cardinals.south},
		"-": [2][2]int{cardinals.east, cardinals.west},
		"L": [2][2]int{cardinals.north, cardinals.east},
		"J": [2][2]int{cardinals.north, cardinals.west},
		"7": [2][2]int{cardinals.south, cardinals.west},
		"F": [2][2]int{cardinals.south, cardinals.east},
	}

	for {

		actualPip := string((*puzzleLines)[nextPos[1]][nextPos[0]])
		cardinal_1 := pipesInstructions[actualPip][0]
		cardinal_2 := pipesInstructions[actualPip][1]

		if [2]int{nextPos[0] + cardinal_1[0], nextPos[1] + cardinal_1[1]} == actuPos {
			temp := nextPos
			actuPos = nextPos
			nextPos = [2]int{temp[0] + cardinal_2[0], temp[1] + cardinal_2[1]}
		} else {
			temp := nextPos
			actuPos = nextPos
			nextPos = [2]int{temp[0] + cardinal_1[0], temp[1] + cardinal_1[1]}
		}

		registPipUsed = append(registPipUsed, actuPos)

		if nextPos == startPos {
			break
		}

	}

	return registPipUsed
}

func cleanLines(oldLines *[]string, regex [][2]int) []string {

	newPuzzle := []string{}

	startingSahpe := getStartingShape([2][2]int{regex[1], regex[len(regex)-1]})

	for idxY, line := range *oldLines {
		newLine := ""
		for idxX, char := range line {

			if slices.Contains(regex, [2]int{idxX, idxY}) {

				if idxY == regex[0][1] && idxX == regex[0][0] {

					newLine += string(startingSahpe)
				} else {

					newLine += string(char)
				}

			} else {
				char = rune('.')
				newLine += string(char)
			}

		}

		newPuzzle = append(newPuzzle, newLine)
	}

	return newPuzzle
}

func getStartingShape(shapes [2][2]int) string {

	start := shapes[0]
	end := shapes[1]

	difference := [2]int{start[0] - end[0], start[1] - end[1]}

	fmt.Println(shapes, "diff: ", difference)

	if difference == [2]int{-1, 1} {
		return "F"
	} else if difference == [2]int{-1, -1} {
		return "7"
	} else if difference == [2]int{1, 1} {
		return "L"
	} else if difference == [2]int{1, -1} {
		return "J"
	} else if difference == [2]int{-2, 0} || difference == [2]int{2, 0} {
		return "|"
	} else {
		return "-"
	}

}

func getOusiders(lines []string, regexPipe [][2]int) int {

	insiders := [][2]int{}

	for idxY, line := range lines {

		within := false
		pipStack := []string{}

		for idxX, char := range line {

			// fmt.Println("Currently checking line ", line, "char", idxX )

			if string(char) == "|" {
				within = !within // Flip the value of within so that next chars are considered in or out

			} else if string(char) == "L" {
				pipStack = append(pipStack, "L")

			} else if string(char) == "F" {
				pipStack = append(pipStack, "F")

			} else if len(pipStack) > 0 && string(char) == "J" || string(char) == "7" {

				if pipStack[0] == "L" && string(char) == "7" {
					within = !within
					pipStack = []string{}
	
				} else if pipStack[0] == "F" && string(char) == "J" {
					within = !within
					pipStack = []string{}
	
				} else if pipStack[0] == "L" && string(char) == "J" {
					pipStack = []string{}
	
				} else if pipStack[0] == "F" && string(char) == "7" {
					pipStack = []string{}
				}
			}
			
			if within {
				insiders = append(insiders, [2]int{idxX, idxY})
			}
		}
	}
	fmt.Println("insiders: ", insiders)

	finalList := [][2]int{}
	for _, inside := range insiders{
		if slices.Contains(regexPipe, inside){
			continue
		} else  {
			finalList = append(finalList, inside)
		}
	}

	return len(finalList)
}
