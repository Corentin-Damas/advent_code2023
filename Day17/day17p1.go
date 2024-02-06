package day17

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "golang.org/x/exp/slices"
)

type Direction struct {
	north, south, east, west [2]int
}
type Point struct {
	currDir [2]int
	currPos [2]int
}
type Node struct {
	pointData      Point
	directionCount int
}

type State struct{
	coast int
	node Node
}

func newPoint(line int, col int, dir [2]int) Point {
	return Point{currPos: [2]int{line, col}, currDir: dir}
}

func (point *Point) nextInBoundPoints(grid [][]int) []Point {
	dir := Direction{
		north: [2]int{-1, 0},
		south: [2]int{1, 0},
		east:  [2]int{0, 1},
		west:  [2]int{0, -1},
	}
	inBoundPoints := []Point{}

	if point.currPos[0] > 0 && point.currDir != dir.east {
		inBoundPoints = append(inBoundPoints, newPoint(point.currPos[0], point.currPos[1]-1, dir.west))
	}
	if point.currPos[1] > 0 && point.currDir != dir.south {
		inBoundPoints = append(inBoundPoints, newPoint(point.currPos[0]-1, point.currPos[1], dir.north))
	}
	if point.currPos[0] < len(grid)-1 && point.currDir != dir.north {
		inBoundPoints = append(inBoundPoints, newPoint(point.currPos[0]+1, point.currPos[1], dir.south))
	}
	if point.currPos[1] < len(grid[0])-1 && point.currDir != dir.west {
		inBoundPoints = append(inBoundPoints, newPoint(point.currPos[0], point.currPos[1]+1, dir.east))
	}
	return inBoundPoints
}

func newNode(point Point, count int) Node {
	return Node{pointData: point, directionCount: count}
}
func neighbors(node *Node, grid [][]int) []Node {
	neighboors := []Node{}
	availableNextNode := node.pointData.nextInBoundPoints(grid)
	for _, avNode := range availableNextNode {
		if avNode.currDir != node.pointData.currDir {
			neighboors = append(neighboors, newNode(avNode, 1))
		} else if node.directionCount < 3 {
			neighboors = append(neighboors, newNode(avNode, node.directionCount+1))

		}
		// All other cases are invalide

	}
	return neighboors
}

// func (state *State)




func Day17p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day17/datatest.txt")
	if err != nil {
		log.Fatal(err)
	}
	puzzle := string(bytesText)
	result := analyzep1(puzzle)
	fmt.Println("Day 17 part 1  : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}
	numGrid := stringToGrid(puzzleLines)

	printGrid(numGrid)

	return 0
}



func stringToGrid(puzzle []string) [][]int {
	grid := [][]int{}
	for idxL, line := range puzzle {
		grid = append(grid, []int{})
		for _, ch := range line {
			grid[idxL] = append(grid[idxL], strtoInt(string(ch)))
		}
	}
	return grid
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

func printGrid(puzzle [][]int) {
	fmt.Println(" ")
	for _, line := range puzzle {
		fmt.Println(line)
	}
}
