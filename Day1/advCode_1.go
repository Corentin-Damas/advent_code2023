package day1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day1result() int{
	testContent, err := os.ReadFile("./Day1/data.txt")
	if err != nil {
		log.Fatal(err)
	} 

	strContent := string(testContent)
	frstline := readlines(strContent) 
	return  frstline
}


func readlines(s string) int{
	resultCoords := 0
	lines := strings.Split(s, "\n")
	for _, line := range lines{
		strCoords := readNumbers(line)
		intCoords  := convertStrtoInt(strCoords)
		resultCoords += intCoords

	}
	return resultCoords
	
}

func readNumbers(s string) string{
	num :="0123456789"
	strNums := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	arrNums := []string{}
	resultInt := ""


	for idx, n := range strNums{ 

		regex := regexp.MustCompile(n)
		s = regex.ReplaceAllString(s, n + string(num[idx]) +n )
		
	}

	for _ ,char := range s{
		for _, n := range num{
			if char == n {
				arrNums = append(arrNums, string(n))
			}

		}
	}

	if len(arrNums) != 0 {
		tempArr := []string{}
		tempArr = append(tempArr, arrNums[0])
		tempArr = append(tempArr, arrNums[len(arrNums)-1])
		arrNums = tempArr
	}

	for _, lt := range arrNums {
		resultInt += lt
	}
	
	return resultInt
}

func convertStrtoInt(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		} 
	
		return i
}