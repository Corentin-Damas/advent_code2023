package day18

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Day18p2() int {
	bytesText, err := os.ReadFile("./Day18/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)
	result := analyzep2(puzzle)

	fmt.Println("Day 18 part 1  : ")
	return result
}

func analyzep2(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}

	instructions := getInstruction(puzzleLines)

	bounds := 0
	vectorsPoints := [][2]int{{0, 0}}
	mapDirection := map[string][2]int{"U": {-1, 0}, "D": {1, 0}, "L": {0, -1}, "R": {0, 1}}
	// Split the instructions
	for _, instruc := range instructions {
		dirs := mapDirection[instruc.color.dir]
		dirLine, dirCol := dirs[0], dirs[1]
		actualPos := vectorsPoints[len(vectorsPoints)-1]
		bounds += instruc.color.distance
		nextPosX, nextPosY := actualPos[0], actualPos[1]
		vectorsPoints = append(vectorsPoints, [2]int{nextPosX + dirLine*instruc.color.distance, nextPosY + dirCol*instruc.color.distance})
	}

	perimeter := 0
	for idx := range vectorsPoints {
		if idx > 0 {
			perimeter += vectorsPoints[idx][0] * (vectorsPoints[idx-1][1] - vectorsPoints[(idx+1)%len(vectorsPoints)][1])

		}
	}
	perimeter = int(math.Abs(float64(perimeter))) / 2
	internArea := perimeter - (bounds / 2) + 1
	return internArea + bounds
}
