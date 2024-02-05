package day15

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	// "golang.org/x/exp/slices"
)

func Day15p2() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day15/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep2(puzzle)

	fmt.Println("Day 15 part 2  : ")
	return result
}

func analyzep2(s string) int {
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

	hashMapLenze := make(map[int][]string)

	for _, pieceOfSequence := range sequence {
		currentValue := 0
		labelToCheck := ""
		for _, runeChar := range pieceOfSequence {

			if string(runeChar) == "=" {
				is_existing := false

				for idx, values := range hashMapLenze[currentValue] {

					valuesLabel := strings.Split(values, "=")
					if valuesLabel[0] == labelToCheck {
						hashMapLenze[currentValue][idx] = pieceOfSequence
						is_existing = true
						break
					}
				}
				if !is_existing{
					hashMapLenze[currentValue] = append(hashMapLenze[currentValue], pieceOfSequence)

				}
				break

			} else if string(runeChar) == "-" {
				for idx, values := range hashMapLenze[currentValue] {
					valuesLabel := strings.Split(values, "=")
					if valuesLabel[0] == labelToCheck {
						hashMapLenze[currentValue] = removeFromSlice(hashMapLenze[currentValue], idx)
					}
				}
				break
			}

			labelToCheck += string(runeChar)

			currentValue += int(runeChar)
			currentValue = (currentValue * 17) % 256

		}
	}
	total := 0

	for box, listOflense := range hashMapLenze {
		fmt.Println("Box:", box, " lense label:", listOflense)
		
		for slot, lense := range listOflense{
			lenseData := strings.Split(lense, "=")
			total += strtoInt(lenseData[1]) * (slot + 1) * (box +1)

		} 
		 
		
	}

	return total
}

func removeFromSlice(slice []string, idx int) []string {
	return append(slice[:idx], slice[idx+1:]...)
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
