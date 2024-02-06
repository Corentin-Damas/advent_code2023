package day16

import (
	"fmt"
	"log"
	"os"
	"strings"
	// "golang.org/x/exp/slices"
)

type Pointer struct {
	currDir  [2]int
	currlPos [2]int
}

func (limits *Pointer) isOutlimits(accMap [][]string) bool {
	heightMapMap := len(accMap) - 1
	widthMap := len(accMap[0]) - 1
	nextLine := limits.currlPos[0]
	nextColumn := limits.currlPos[1]

	if nextLine > heightMapMap || nextColumn > widthMap || nextLine < 0 || nextColumn < 0 {
		return true
	}
	return false
}

type Queue struct {
	queue []Pointer
}

type Direction struct {
	left  [2]int
	right [2]int
	top   [2]int
	bot   [2]int
}

func Day16p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day16/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)
	result := analyzep1(puzzle)

	fmt.Println("Day 16 part 1  : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}
	puzzleArray := puzzleMapping(puzzleLines)
	energizedArray := getEnergizeMap(puzzleArray)
	// printMap(puzzleArray)

	dir := Direction{
		top:   [2]int{-1, 0},
		bot:   [2]int{1, 0},
		right: [2]int{0, 1},
		left:  [2]int{0, -1},
	}
	qu := Queue{
		[]Pointer{},
	}

	newConfig([2]int{0, 0}, "right", &dir, &qu)

	//check for loops
	mapPointer := make(map[Pointer]int)

	for len(qu.queue) > 0  {

		fThread := qu.queue[0]
		// fmt.Println(" pos", fThread.currlPos, " direction", fThread.currDir, " queue", qu)

		if puzzleArray[fThread.currlPos[0]][fThread.currlPos[1]] == "." {
			energizedArray[fThread.currlPos[0]][fThread.currlPos[1]] = "#"

			// Go to the next location for this thread

			qu.queue[0].currlPos = [2]int{fThread.currlPos[0] + fThread.currDir[0], fThread.currlPos[1] + fThread.currDir[1]}

		} else if puzzleArray[fThread.currlPos[0]][fThread.currlPos[1]] == "|" {
			energizedArray[fThread.currlPos[0]][fThread.currlPos[1]] = "#"

			if fThread.currDir == dir.top || fThread.currDir == dir.bot {
				// Go to the next location
				qu.queue[0].currlPos = [2]int{fThread.currlPos[0] + fThread.currDir[0], fThread.currlPos[1] + fThread.currDir[1]}

			} else {
				// Reflect the light up and down && check if they are out of bound or not
				newConfig(fThread.currlPos, "top", &dir, &qu)
				if qu.queue[len(qu.queue)-1].isOutlimits(puzzleArray) {
					qu.queue = qu.queue[0 : len(qu.queue)-1]
				}

				newConfig(fThread.currlPos, "bot", &dir, &qu)
				if qu.queue[len(qu.queue)-1].isOutlimits(puzzleArray) {
					qu.queue = qu.queue[0 : len(qu.queue)-1]
				} // The first thread stop on the Glass and created 2 new threads
				qu.queue = qu.queue[1:]
			}

		} else if puzzleArray[fThread.currlPos[0]][fThread.currlPos[1]] == "-" {
			energizedArray[fThread.currlPos[0]][fThread.currlPos[1]] = "#"

			if fThread.currDir == dir.left || fThread.currDir == dir.right {
				// Go to the next location
				qu.queue[0].currlPos = [2]int{fThread.currlPos[0] + fThread.currDir[0], fThread.currlPos[1] + fThread.currDir[1]}

			} else {
				// Reflect the light up and down && check if they are out of bound or not
				newConfig(fThread.currlPos, "left", &dir, &qu)
				if qu.queue[len(qu.queue)-1].isOutlimits(puzzleArray) {
					qu.queue = qu.queue[0 : len(qu.queue)-1]
				}
				newConfig(fThread.currlPos, "right", &dir, &qu)
				if qu.queue[len(qu.queue)-1].isOutlimits(puzzleArray) {
					qu.queue = qu.queue[0 : len(qu.queue)-1]
				}
				qu.queue = qu.queue[1:]
			}

		} else if puzzleArray[fThread.currlPos[0]][fThread.currlPos[1]] == "/" {
			energizedArray[fThread.currlPos[0]][fThread.currlPos[1]] = "#"

			if fThread.currDir == dir.left {
				qu.queue[0].currDir = dir.bot
			} else if fThread.currDir == dir.bot {
				qu.queue[0].currDir = dir.left
			} else if fThread.currDir == dir.right {
				qu.queue[0].currDir = dir.top
			} else if fThread.currDir == dir.top {
				qu.queue[0].currDir = dir.right
			}
			qu.queue[0].currlPos = [2]int{qu.queue[0].currlPos[0] + qu.queue[0].currDir[0], qu.queue[0].currlPos[1] + qu.queue[0].currDir[1]}

		} else if puzzleArray[fThread.currlPos[0]][fThread.currlPos[1]] == "\\" {
			energizedArray[fThread.currlPos[0]][fThread.currlPos[1]] = "#"

			if fThread.currDir == dir.left {
				qu.queue[0].currDir = dir.top
			} else if fThread.currDir == dir.bot {
				qu.queue[0].currDir = dir.right
			} else if fThread.currDir == dir.right {
				qu.queue[0].currDir = dir.bot
			} else if fThread.currDir == dir.top {
				qu.queue[0].currDir = dir.left
			}
			qu.queue[0].currlPos = [2]int{qu.queue[0].currlPos[0] + qu.queue[0].currDir[0], qu.queue[0].currlPos[1] + qu.queue[0].currDir[1]}

		} else {
			fmt.Println("New Symbole found BREAK")
			break
		}

		// check if Pointer was already checked
		count, ok := mapPointer[qu.queue[0]]
		if !ok {
			mapPointer[qu.queue[0]] = 1
		} else if count > 2 {
			if len(qu.queue) > 1 {
				qu.queue = qu.queue[1:]
			} else {
				break
			}
		} else {
			mapPointer[qu.queue[0]]++
		}

		if qu.queue[0].isOutlimits(puzzleArray) {
			if len(qu.queue) > 1 {
				qu.queue = qu.queue[1:]
				continue
			} else if len(qu.queue) <= 1 {
				fmt.Println("BREAK")
				break
			}
		}

	}
	// printMap(energizedArray)

	total := countEnergizeTiles(energizedArray)

	return total
}

func puzzleMapping(puzzle []string) [][]string {
	puzzleMap := [][]string{}

	for idxL, line := range puzzle {
		puzzleMap = append(puzzleMap, []string{})
		for _, ch := range line {
			puzzleMap[idxL] = append(puzzleMap[idxL], string(ch))
		}
	}
	return puzzleMap
}

func newConfig(actualPos [2]int, orientation string, dir *Direction, queue *Queue) {

	var actualDir [2]int
	if orientation == "right" {
		actualDir = dir.right
	} else if orientation == "left" {
		actualDir = dir.left
	} else if orientation == "bot" {
		actualDir = dir.bot
	} else {
		actualDir = dir.top
	}

	newPointer := Pointer{
		currDir:  actualDir,
		currlPos: actualPos,
	}
	queue.queue = append(queue.queue, newPointer)
}

func getEnergizeMap(originalMap [][]string) [][]string {

	copiedMap := [][]string{}

	for idxL, line := range originalMap {
		copiedMap = append(copiedMap, []string{})
		for idx := 0; idx <= len(line)-1; idx++ {
			copiedMap[idxL] = append(copiedMap[idxL], ".")
		}
	}
	return copiedMap

}

// func printMap(puzzle [][]string) {
// 	fmt.Println(" ")
// 	for _, line := range puzzle {
// 		fmt.Println(line)
// 	}

// }

func countEnergizeTiles(puzzle [][]string) int {
	total := 0
	for _, line := range puzzle {
		for _, ch := range line {
			if ch == "#"{
				total ++
			}
		}
	}
	return total
}