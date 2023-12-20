package day2

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	red   int
	blue  int
	green int
}

func Day2result() int {

	testContent, err := os.ReadFile("./Day2/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	strContent := string(testContent)
	linesResults := readlines(strContent)

	fmt.Println("Day 2 result is : ")
	return linesResults
}

func readlines(s string) int {
	resultPossibleGames := 0

	bag := Bag{
		red: 12,
		green: 13,
		blue: 14,
	}
	var pBag *Bag = &bag

	lines := strings.Split(s, "\n")
	for _, line := range lines {

		gameId, err := readLineInfo(line, pBag)
		if err != nil {
			fmt.Errorf("game %d not valid", gameId)
		} else {
			resultPossibleGames += gameId
		}

	}
	return resultPossibleGames

}

func readLineInfo(s string, pBag *Bag) (int, error) {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green || break at :
	arrString := strings.Split(s, ": ")
	id := gameId(arrString[0])
	result := readGameResults(arrString[1], pBag)
	return id, result

}

func gameId(s string) int {
	re := regexp.MustCompile(`(Game)\s(\w+)`)
	groups := re.FindStringSubmatch(s)

	idStr := groups[2]
	idInt := convertStrtoInt(idStr)
	return idInt
}

func readGameResults(s string, pBag *Bag) error {
	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	arrStringComplete := strings.Split(s, ";")

	for _, hand := range arrStringComplete{
		//3 blue, 4 red;
		arrHand := strings.Split(hand, ",")

		for _, color := range arrHand{
			//3 blue
			err := colorCheck(color, pBag)
			if err != nil {
				return err 
			}
		}
	}
	return nil
}


func colorCheck(s string, pBag *Bag) error{
	//3 blue
	re := regexp.MustCompile(`(\w+)\s(blue|green|red)`)
	groups := re.FindStringSubmatch(s)

	colorNum, err := strconv.Atoi(groups[1])
	if err != nil {
		log.Fatal(err)
	}

	if (groups[2]=="blue") && (colorNum <= pBag.blue) {
		return nil
	}
	if (groups[2]== "green") && (colorNum <= pBag.green) {
		return nil
	}
	if (groups[2]== "red" && colorNum <= pBag.red) {
		return nil
	}
	return errors.New("too many ball in hand compared to bag capacity")

}

func convertStrtoInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
