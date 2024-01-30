package day11

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func Day11p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day11/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 11 part 1 & part 2 result is (for this time we just have tochange the scale of one param) : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}
	// 9 galaxie : 8+7+6+5+4+3+2+1  // (9*8)/2

	// Identifying empty row and empty col
	idxEmptyRow, idxEmptyCol, idxPlanets := getExpendingGalaxies(puzzleLines)

	total := 0
	scale := 1000000 // part 1 is scale = 2

	for p, scannedPlanetCoor := range idxPlanets[:len(idxPlanets)-1] {
		for _, targetPlanetCoor := range idxPlanets[p+1:] {
			// for i:= scannedPlanetCoor[0] - targetPlanetCoor[0]; i >= 0; i --{
			for h := scannedPlanetCoor[0]; h < targetPlanetCoor[0]; h++ {
				if slices.Contains(idxEmptyRow, h) {
					total += scale
				} else {
					total += 1
				}
			}
			if targetPlanetCoor[1] > scannedPlanetCoor[1] {
				for h := scannedPlanetCoor[1]; h < targetPlanetCoor[1]; h++ {
					if slices.Contains(idxEmptyCol, h) {
						total += scale
					} else {
						total += 1
					}
				}
			} else {
				for h := scannedPlanetCoor[1]; h > targetPlanetCoor[1]; h-- {
					if slices.Contains(idxEmptyCol, h) {
						total += scale
					} else {
						total += 1
					}
				}
			}

		}
	}

	return total
}

func getExpendingGalaxies(s []string) ([]int, []int, [][2]int) {
	emptyLine := []int{}
	emptyCol := []int{}
	planets := [][2]int{}
	verticalSwap := []string{}

	for l, line := range s {
		if !strings.Contains(line, "#") {
			emptyLine = append(emptyLine, l)
		}

		for c, letter := range line {
			if l == 0 {
				verticalSwap = append(verticalSwap, string(letter))
			} else {
				verticalSwap[c] = verticalSwap[c] + string(letter)
			}
			if string(letter) == "#" {
				planets = append(planets, [2]int{l, c})
			}

		}
	}
	for l, line := range verticalSwap {
		if !strings.Contains(line, "#") {
			emptyCol = append(emptyCol, l)
		}
	}
	return emptyLine, emptyCol, planets
}
