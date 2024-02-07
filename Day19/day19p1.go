package day19

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "golang.org/x/exp/slices"
)

type Rule struct {
	pieceCat    string
	condition   string
	threshold   int
	destination string
	isDefault   bool
}

type Workflow struct {
	setRule []Rule
}

type Piece struct {
	data map[string]int

	isAccepted bool
}

func Day19p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day19/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)
	result := analyzep1(puzzle)

	fmt.Println("Day 19 part 1  : ")
	return result
}

func analyzep1(s string) int {
	puzzleLines := strings.Split(s, "\n")
	for idx := range puzzleLines {
		puzzleLines[idx] = strings.Trim(puzzleLines[idx], "\r\n")
	}

	setWorkflow := make(map[string]Workflow)
	setPieces := []Piece{}

	for _, line := range puzzleLines {
		if len(line) > 1 {

			if string(line[0]) != "{" {
				key, newWorkflow := getworkflow(line)
				setWorkflow[key] = newWorkflow
			} else {
				setPieces = append(setPieces, getPieceData(line))
			}
		}
	}

	total := 0
	for _, piece := range setPieces {
		piece.isAccepted = checkIsAcceped(piece, &setWorkflow)
		if piece.isAccepted {
			for _, val := range piece.data {
				total += val
			}
		}

	}

	return total
}

func getPieceData(s string) Piece {
	newPiece := Piece{make(map[string]int), false}

	s = strings.Trim(s, "{}")
	numsStr := strings.Split(s, ",")
	for _, n := range numsStr {
		arrData := strings.Split(n, "=")
		newPiece.data[arrData[0]] = strtoInt(arrData[1])

	}
	return newPiece
}

func checkIsAcceped(piece Piece, allWorkflows *map[string]Workflow) bool {
	accKeyWf := "in"
	testLoop := 0
	accted := false
	for testLoop < 20 {
		testLoop++
		workflow := (*allWorkflows)[accKeyWf]

		for _, rule := range workflow.setRule {
			if !rule.isDefault {
				numTocheck := piece.data[rule.pieceCat]
				if rule.condition == ">" && numTocheck > rule.threshold {
					if rule.destination == "A" {
						accted = true
						return true
					}
					if rule.destination == "R" {
						return false
					}
					accKeyWf = rule.destination
					break
				}
				if rule.condition == ">" && numTocheck < rule.threshold {
					continue
				}

				if rule.condition == "<" && numTocheck < rule.threshold {
					if rule.destination == "A" {
						accted = true
						return true
					} else if rule.destination == "R" {
						return false
					} else {
						accKeyWf = rule.destination
						break
					}
				}
				if rule.condition == "<" && numTocheck > rule.threshold {
					continue
				}
			} else {
				if rule.destination == "A" {
					accted = true
					return true
				} else if rule.destination == "R" {
					return false
				} else {
					accKeyWf = rule.destination
				}
			}
		}
	}
	return accted
}

func getworkflow(s string) (string, Workflow) {

	strSplit := strings.Split(s, "{")
	key := strSplit[0]
	strRules := strSplit[1]
	strRules = strings.Trim(strRules, "}")

	workf := Workflow{[]Rule{}}
	arrRules := strings.Split(strRules, ",")

	for idx, rule := range arrRules {
		if idx != len(arrRules)-1 {
			workf.setRule = append(workf.setRule, getRule(rule))
		} else {
			workf.setRule = append(workf.setRule, defaultRule(rule))
		}
	}

	return key, workf
}

func getRule(s string) Rule {
	cat := string(s[0])
	cond := string(s[1])
	rest := s[2:]
	restSplit := strings.Split(rest, ":")
	thresh := strtoInt(restSplit[0])
	destin := restSplit[1]

	newCat := Rule{pieceCat: cat, condition: cond, threshold: thresh, destination: destin, isDefault: false}
	return newCat
}

func defaultRule(s string) Rule {
	return Rule{destination: s, isDefault: true}
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

// Each rule seperated by "," ;
// Condition < > after condition : indicate the workflow to follow
// A : accepted
// R : Rejected
// Else (last element) follow a workflow or Get Accepted / rejected
