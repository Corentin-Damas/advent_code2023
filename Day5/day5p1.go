package day5

import (
	"fmt"
	"log"

	"os"
	"strconv"
	"strings"
	"unicode"
	// "golang.org/x/exp/slices"
)

type Filtering struct {
	soil        [][]int
	fertilizer  [][]int
	water       [][]int
	light       [][]int
	temperature [][]int
	humidity    [][]int
	location    [][]int
	instruction []string
}

func Day5p1() int {

	bytesText, err := os.ReadFile("./Day5/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	FullText := string(bytesText)
	result := analyzep1(FullText)

	fmt.Println("Day 5 part 1 result is : ")
	fmt.Println("This solution might take a lot of time to resolve... ")
	return result
}

func analyzep1(s string) int {
	lines := strings.Split(s, "\n")
	seeds := getSeeds(lines[0])
	lineMapping := lines[1:]
	fmt.Println(seeds)

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
	for idx, seed := range seeds {
		seeds[idx] = getThroughFilter(seed, pFilter)
	}
	
	fmt.Println(seeds)
	var min int 
	for i, seed := range seeds{
		
		if  i == 0 {
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
func getSeeds(s string) []int {
	splitS := strings.Split(s, ":")
	seedsString := splitS[1]
	return strToArr(seedsString)
}

func getStepTitle(s string) string {
	strSpaceSplit := strings.Split(s, " ")
	strMinusSplit := strings.Split(strSpaceSplit[0], "-")
	return strMinusSplit[2]
}

func (filter *Filtering) createFilter(step string, mapline string) {

	mapping := strToArr(mapline)

	sourceStart := mapping[1]
	destinationStart := mapping[0]
	lenghtMap := mapping[2]

	var payload []int
	payload = unMap(sourceStart, lenghtMap, destinationStart)

	switch step {
	case "soil":
		filter.soil = append(filter.soil, payload)
	case "fertilizer":
		filter.fertilizer = append(filter.fertilizer, payload)
	case "water":
		filter.water = append(filter.water, payload)
	case "light":
		filter.light = append(filter.light, payload)
	case "temperature":
		filter.temperature = append(filter.temperature, payload)
	case "humidity":
		filter.humidity = append(filter.humidity, payload)
	case "location":
		filter.location = append(filter.location, payload)
	}
}

// ReDo this so the map is [start, end, + filter]
func unMap(start int, len int, objective int) []int {
	var result []int

	dividend := start - objective

	result = append(result, start)
	result = append(result, start + (len -1))
	result = append(result, dividend)

	return result
}

func getThroughFilter(seed int, pFilter *Filtering) int {
	var seedTransformation int = seed

	for _, instruction := range pFilter.instruction {
		switch instruction {
		case "soil":
			seedTransformation = onFilter(pFilter.soil, seedTransformation)
		case "fertilizer":
			seedTransformation = onFilter(pFilter.fertilizer, seedTransformation)
		case "water":
			seedTransformation = onFilter(pFilter.water, seedTransformation)
		case "light":
			seedTransformation = onFilter(pFilter.light, seedTransformation)
		case "temperature":
			seedTransformation = onFilter(pFilter.temperature, seedTransformation)
		case "humidity":
			seedTransformation = onFilter(pFilter.humidity, seedTransformation)
		case "location":
			seedTransformation = onFilter(pFilter.location, seedTransformation)
		}
	}

	return seedTransformation
}

func onFilter(filter [][]int, seed int) int {
	
	for _, i := range filter{
		if seed >= i[0] && seed <= i[1] {
			return seed - i[2]
		}
	}
	return seed
}


// ================================   Utilities   ================================

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

// Seed that need to be planted
// They will go througt a serie of steps each steps has numbers  but soil 123 != fertilizer 123

// Each step are map / filter that convert numbers  Source -> filter -> destination
// Eg   seed number -> seed-to-soil -> soil number

// Destination range start | source range start | range lenght  // 50 98 2

//  source range start 98 lenght 2 : [ 98, 99 ] , destination range start 50 len 2 : [50, 51]
// seed 98 -> soil 50 // seed 99 -> soil 51
// repeeat fot the second line

// if a number is not mapped seed 10 -> soil 10
// They have to pas through all the steps and the LOWEST is the answer

// [[98 99 -48][......]]
