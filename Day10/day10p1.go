package day10

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type CardinalPoint struct {
	north [2]int
	south [2]int
	east  [2]int
	west  [2]int
}

func Day10p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day10/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 10 part 1 result is : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	pPuzzeLines := &puzzleLines

	startingPos, startingsLoop := getStartingPos(pPuzzeLines)

	totalPipeUsed := followPipe(startingPos, startingsLoop, pPuzzeLines)

	return totalPipeUsed / 2
}

func getStartingPos(lines *[]string) ([2]int, [2]int) {

	startPos := [2]int{}
	startPipe := [2]int{}
	for idxY, lineY := range *lines {
		for idxX, char := range lineY {
			if string(char) == "S" {
				startPos = [2]int{idxX, idxY}
			}
		}

	}

	if string((*lines)[startPos[1]-1][startPos[0]]) == "|" || string((*lines)[startPos[1]-1][startPos[0]]) == "F" || string((*lines)[startPos[1]-1][startPos[0]]) == "7" {
		startPipe = [2]int{startPos[0], startPos[1] - 1}

	} else if string((*lines)[startPos[1]+1][startPos[0]]) == "|" || string((*lines)[startPos[1]+1][startPos[0]]) == "J" || string((*lines)[startPos[1]+1][startPos[0]]) == "L" {
		startPipe = [2]int{startPos[0], startPos[1] + 1}

	} else if string((*lines)[startPos[1]][startPos[0]-1]) == "-" || string((*lines)[startPos[1]][startPos[0]-1]) == "L" || string((*lines)[startPos[1]][startPos[0]-1]) == "F" {
		startPipe = [2]int{startPos[0] + 1, startPos[1]}
	} else {
		startPipe = [2]int{startPos[0] - 1, startPos[1]}
	}

	return startPos, startPipe

}

func followPipe(startPos [2]int, StartLoop [2]int, puzzleLines *[]string) int {

	// registPipUsed:= [][2]int{}
	count := 0
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
		count += 1

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

		if nextPos == startPos {
			count += 1
			break
		}

	}

	return count
}
