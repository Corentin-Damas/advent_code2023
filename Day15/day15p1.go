package day15

import (
	"fmt"
	"log"
	"os"
	"strings"

	// "golang.org/x/exp/slices"
)

func Day15p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day15/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 15 part 1  : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}

	// there is only one long thread line of char
	hashLine := puzzleLines[0]

	sequence := strings.Split(hashLine, ",")

	// Determine ASCII Code
	// runeChar := 'x' 
	// ascii := int(runeChar) 

	total := 0
	
	for _, pieceOfSequence := range sequence{
		currentValue := 0
		for _, runeChar := range pieceOfSequence {

			currentValue += int(runeChar)
			currentValue = (currentValue * 17) % 256
		}
		fmt.Println(pieceOfSequence, " becomes ", currentValue)
		total += currentValue

	}


	
	return total
}

