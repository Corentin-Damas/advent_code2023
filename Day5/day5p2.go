package day5

import (
	"fmt"
	"log"

	"os"
	"strings"
	"unicode"
)

func Day5p2() int {

	bytesText, err := os.ReadFile("./Day5/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	FullText := string(bytesText)
	result := analyzep2(FullText)

	fmt.Println("Day 5 part 2 result is : ")
	return result
}

func analyzep2(s string) int {
	lines := strings.Split(s, "\n")
	seedsInstruction := getSeeds(lines[0])

	lineMapping := lines[1:]
	fmt.Println(seedsInstruction)

	filter := Filtering{}
	filter.instruction = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	pFilter := &filter
	// Create Filter mapping
	var actualStep string
	for _, line := range lineMapping {

		if line == "\r" {
			continue
		}
		if !unicode.IsDigit(rune(line[0])) {
			actualStep = line
			continue
		}

		step := getStepTitle(actualStep)
		pFilter.createFilter(step, line)
	}

	// Pass the seeds throught the filters
	var seedRange [][2]int
	tempArr := [2]int{0, 0}
	for idx := range seedsInstruction {

		if idx%2 == 0 {
			tempArr[0] = seedsInstruction[idx]
			tempArr[1] = seedsInstruction[idx+1]
		} else {
			seedRange = append(seedRange, tempArr)
		}
	}
	fmt.Println(seedRange)

	var resultSeed []int
	for idx, seed := range seedsInstruction {
		if idx%2 == 0 {
			resultSeed = append(resultSeed, seed)
		}
	}

	for idx, seedR := range seedRange {
		for i := 0; i <= seedR[1]; i++ {
			potentialSeed := getThroughFilter(seedR[0]+i, pFilter)
			if potentialSeed < resultSeed[idx] {
				resultSeed[idx] = potentialSeed
			}
		}
	}

	fmt.Println(resultSeed)
	var min int
	for i, seed := range resultSeed {

		if i == 0 {
			min = seed
		} else {
			if seed < min {
				min = seed
			}
		}
	}
	return min

}

// getSeeds Receive the first line (String) of the text and extract the seeds numbers to an array of int ([]int )
