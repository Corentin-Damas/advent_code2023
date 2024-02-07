package day18

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	dir      string
	move     int
	colorStr string
	color    Color
}

type Color struct {
	distance int
	dir      string
}

func getColorInfo(s string) Color {
	fullNumStr := strings.Trim(s, "()#")
	hexStr := fullNumStr[:len(fullNumStr)-1]
	direction := string(fullNumStr[len(fullNumStr)-1])
	value, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Printf("conversion faled: %s\n", err)
	} else if direction == "0" {
		direction = "R"
	} else if direction == "1" {
		direction = "D"
	} else if direction == "2" {
		direction = "L"
	} else if direction == "3" {
		direction = "U"
	}

	return Color{int(value), direction}
}

func Day18p1() int {
	bytesText, err := os.ReadFile("./Day18/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)
	result := analyzep1(puzzle)

	fmt.Println("Day 18 part 1  : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}

	instructions := getInstruction(puzzleLines)

	bounds := 0
	vectorsPoints := [][2]int{{0, 0}}
	mapDirection := map[string][2]int{"U": {-1, 0}, "D": {1, 0}, "L": {0, -1}, "R": {0, 1}}

	for _, instruc := range instructions {
		dirs := mapDirection[instruc.dir]
		dirLine, dirCol := dirs[0], dirs[1]
		actualPos := vectorsPoints[len(vectorsPoints)-1]
		bounds += instruc.move
		nextPosX, nextPosY := actualPos[0], actualPos[1]
		vectorsPoints = append(vectorsPoints, [2]int{nextPosX + dirLine*instruc.move, nextPosY + dirCol*instruc.move})
	}

	// Area of polygone use of Shoelace Formula https://en.wikipedia.org/wiki/Shoelace_formula
	// + https://en.wikipedia.org/wiki/Pick%27s_theorem
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

func getInstruction(puzzleLines []string) []Instruction {
	instructionArr := []Instruction{}

	for _, line := range puzzleLines {
		lineSplit := strings.Split(line, " ")
		instructionArr = append(instructionArr, Instruction{dir: lineSplit[0], move: strtoInt(lineSplit[1]), colorStr: lineSplit[2], color: getColorInfo(lineSplit[2])})

	}
	return instructionArr
}

func strtoInt(s string) int {
	s = strings.Trim(s, "\r")
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i
}

func printGrid(grid [][]string) {
	for _, line := range grid {
		fmt.Println(line)
	}
}
