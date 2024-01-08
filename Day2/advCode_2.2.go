package day2

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag2 struct {
	red   int
	blue  int
	green int
}

func Day2p2result() int {

	testContent, err := os.ReadFile("./Day2/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	strContent := string(testContent)
	linesResults := readlines2(strContent)

	fmt.Println("Day 2 part 2 result is : ")
	return linesResults
}

func readlines2(s string) int {
	resultPossibleGames := 0

	bag := Bag2{
		red: 0,
		green: 0,
		blue: 0,
	}
	var pBag *Bag2 = &bag

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		fmt.Println(line)

		readLineCubes(line, pBag)
		powerOfCube := (pBag.red * pBag.blue * pBag.green)
		resultPossibleGames += powerOfCube

		pBag.red = 0
		pBag.blue = 0
		pBag.green = 0

	}
	return resultPossibleGames

}


func readLineCubes(s string, pBag *Bag2)  {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green || break at :
	arrString := strings.Split(s, ": ")
	err := readGameResults2(arrString[1], pBag)
	if err != nil {
		fmt.Errorf("game disturbed not valid")
	}
}

func readGameResults2(s string, pBag *Bag2) error {
	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	arrStringComplete := strings.Split(s, ";")

	for _, hand := range arrStringComplete{
		//3 blue, 4 red;
		arrHand := strings.Split(hand, ",")

		for _, color := range arrHand{
			//3 blue
			colorCheck2(color, pBag)

		}
	}

	return nil
}


func colorCheck2(s string, pBag *Bag2) {
	//3 blue
	re := regexp.MustCompile(`(\w+)\s(blue|green|red)`)
	groups := re.FindStringSubmatch(s)

	colorNum, err := strconv.Atoi(groups[1])
	if err != nil {
		log.Fatal(err)
	}

	if (groups[2]=="blue") && (colorNum > pBag.blue) {
		pBag.blue = colorNum
	}

	if (groups[2]== "green") && (colorNum > pBag.green) {
		pBag.green = colorNum
	}

	if (groups[2]== "red" && colorNum > pBag.red) {
		pBag.red = colorNum
	}
}
