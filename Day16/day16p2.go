package day16

import (
	"fmt"
	"log"
	"os"
	"strings"
	// "golang.org/x/exp/slices"
)

type Starters struct {
	startPos    [2]int
	orientation string
}

func Day16p2() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day16/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)
	result := analyzep2(puzzle)

	fmt.Println("Day 16 part 2 : ")
	return result
}

func analyzep2(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}
	puzzleArray := puzzleMapping(puzzleLines)
	
	// printMap(puzzleArray)

	dir := Direction{
		top:   [2]int{-1, 0},
		bot:   [2]int{1, 0},
		right: [2]int{0, 1},
		left:  [2]int{0, -1},
	}


	// Check for every tile on the outer edge
	// Get list of outer edge starter Pointer

	arrayStarterInstruction := getAllStarter(puzzleArray)
	mostEnergized := 0

	for _, starter := range arrayStarterInstruction {
		// Reset Energize map and queu and mapPointer(for loops detection)
		energizedArray := getEnergizeMap(puzzleArray)
		qu := Queue{
			[]Pointer{},
		}
		mapPointer := make(map[Pointer]int)
		newConfig(starter.startPos, starter.orientation, &dir, &qu)
		
		fmt.Println(starter, mostEnergized)
		for len(qu.queue) > 0 {

			fThread := qu.queue[0]
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
					break
				}
			}

		}
		energized := countEnergizeTiles(energizedArray)
		if energized > mostEnergized{
			mostEnergized = energized
		}
	}
	// printMap(energizedArray)
	return mostEnergized
}

func getAllStarter(puzzle [][]string) []Starters {
	listStarter := []Starters{}

	for idxL, line := range puzzle {
		for idxC := range line {

			if idxL == 0 {
				listStarter = append(listStarter, Starters{[2]int{idxL, idxC}, "bot"})
			} else if idxC == 0 {
				listStarter = append(listStarter, Starters{[2]int{idxL, idxC}, "right"})
			} else if idxL == len(puzzle)-1 {
				listStarter = append(listStarter, Starters{[2]int{idxL, idxC}, "top"})
			} else if idxC == len(line)-1 {
				listStarter = append(listStarter, Starters{[2]int{idxL, idxC}, "left"})
			}

		}
	}

	return listStarter
}
