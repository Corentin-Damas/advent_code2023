package day4

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func Day4p2() int {

	bytesText, err := os.ReadFile("./Day4/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	FullText := string(bytesText)
	result := analyzeDay2(FullText)

	fmt.Println("Day 4 part 1 result is : ")
	return result
}

func analyzeDay2(s string) int {

	mapScratchedCard := make(map[int]int)

	var points int
	var accumulator int = 0

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		points = 0
		cardID, winningNums, scratchedNums := extractLine(line)

		for _, scrNum := range winningNums {
			if slices.Contains(scratchedNums, scrNum) {
				points += 1
			}
		}
		mapScratchedCard[cardID] += 1

		if points > 0 {
			tempId := cardID

			for i := points; i > 0; i-- {
				tempId += 1
				mapScratchedCard[tempId] += (1 * mapScratchedCard[cardID])
			}
		}

	}
	for _, val := range mapScratchedCard {
		accumulator += val
	}
	return accumulator
}

func extractLine(line string) (int, []int, []int) {
	var winningNums string
	var onGameNums string

	midSplit := strings.Split(line, " | ")
	winningNums = midSplit[0]
	onGameNums = midSplit[1]
	winningSplit := strings.Split(winningNums, ": ")
	winningNums = winningSplit[1]

	cardStr := winningSplit[0]

	cardIDSplit := strings.Split(cardStr, "Card ")
	cardID := cardIDSplit[1]

	trimedCardId := trimId(cardID)
	

	fmt.Println("Card id:", trimedCardId)

	return trimedCardId, strToArr(winningNums), strToArr(onGameNums)
}

func trimId (s string) int {
	strAcc := ""
	for _, char := range s{
		if string(char) == " " || string(char) == ""{
			continue
		} else {
			strAcc += string(char)
		}
	}

	id, err := strtoInt(strAcc)
	if err != nil {
		fmt.Println(err)
	}

	return id
}