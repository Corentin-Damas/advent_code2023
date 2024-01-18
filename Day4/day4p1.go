package day4

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"strconv"

	"golang.org/x/exp/slices"
)

func Day4() int {

	bytesText, err := os.ReadFile("./Day4/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	FullText := string(bytesText)
	result := analyze(FullText)

	fmt.Println("Day 4 part 1 result is : ")
	return result
}

func analyze(s string) int {

	var points int
	var accumulator int = 0

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		points = 0
		winningNums, scratchedNums := lineToArr(line)

		for _, scrNum := range winningNums {
			if slices.Contains(scratchedNums, scrNum) {
				points += 1
			}
		}

		if points > 0 {
			accumulator += int(math.Pow(2, float64(points-1)))
		}

	}
	return accumulator
}

func lineToArr(line string) ([]int, []int) {
	var winningNums string
	var onGameNums string

	midSplit := strings.Split(line, " | ")
	winningNums = midSplit[0]
	onGameNums = midSplit[1]
	winningSplit := strings.Split(winningNums, ": ")
	winningNums = winningSplit[1]

	return strToArr(winningNums), strToArr(onGameNums)
}

func strToArr(s string) []int {

	strArr := strings.Split(s, " ")
	intArr := []int{}
	for _, num := range strArr {
		if num == " " || num == "" {
			continue
		}
		intnum, err := strtoInt(num)
		if err != nil {
			continue
		}
		intArr = append(intArr, intnum)
	}

	return intArr
}

func strtoInt(s string) (int, error) {
	s = strings.Trim(s, "\r")
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return i, nil
}
