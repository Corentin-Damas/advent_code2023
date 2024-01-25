package day8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction struct {
	left        int
	right       int
	start       string
	end         string
	instruction string
}

func Day8p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day8/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 8 part 1 result is : ")
	return result
}

func analyzep1(s string) int {
	lines := strings.Split(s, "\n")

	directions := lines[0]
	directions = strings.Trim(directions, "\r\n")
	mapInstruction := getMapTranslation(lines[2:])

	lrDir := Direction{
		left:        0,
		right:       1,
		start:       "AAA",
		end:         "ZZZ",
		instruction: directions,
	}

	fmt.Println("Direction are: ", directions)

	// for key, val := range mapInstruction {
	// 	fmt.Println(key, " : ", val)
	// }

	plrDir := &lrDir

	acc := travelTime(mapInstruction, plrDir)

	return acc

}

func getMapTranslation(puzzleLines []string) map[string][2]string {

	puzzleMap := make(map[string][2]string)

	for _, line := range puzzleLines {
		keyValue := strings.Split(line, " = ")
		key := keyValue[0]
		values := strings.Split(keyValue[1], ", ")
		leftVal := values[0][1:4]
		rigVal := values[1][0:3]

		puzzleMap[key] = [2]string{leftVal, rigVal}
	}
	return puzzleMap
}

func travelTime(puzzleMap map[string][2]string, direction *Direction) int {
	var acc int = 0

	actualPos := direction.start

	for actualPos != direction.end {

		for _, c := range direction.instruction {

			if string(c) == "L" {
				actualPos = puzzleMap[actualPos][direction.left]
			}
			if string(c) == "R" {
				actualPos = puzzleMap[actualPos][direction.right]
			}
			acc += 1
			if actualPos == direction.end{
				break
			}
		}
	}
	return acc
}
