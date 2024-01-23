package day7

import (
	"errors"
	"fmt"
	"log"

	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Player struct {
	hand             string
	handValuePerCard []int
	bid              int
	handType         string
	handPoint        int
	subSorted        bool
	rank             int
}

type PlayerData []Player

func Day7p1() int {

	// cardsMapValues := createCards()
	bytesText, err := os.ReadFile("./Day7/datatest.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 7 part 1 result is : ")
	return result
}

func analyzep1(s string) int {
	lines := strings.Split(s, "\n")

	playerData := getPlayersDatas(lines)

	orderPlayers := getRank(playerData)

	for _, hand := range orderPlayers {
		fmt.Printf("hand: %s bid:%d type: %s handpoint: %d rank %d \n", hand.hand, hand.bid, hand.handType, hand.handPoint, hand.rank)
	}

	return 0

}

func singleCardValues() map[string]int {
	labels := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	cardValues := make(map[string]int)

	for i := len(labels) - 1; i <= 0; i-- {
	}

	for idx, lab := range labels {
		cardValues[lab] = len(labels) - idx
	}
	return cardValues
}

func getPlayersDatas(lines []string) PlayerData {
	var playerData PlayerData

	for _, line := range lines {
		d := strings.Split(line, " ")
		hand := d[0]
		handCardValue := getArrValuePerCard(hand)
		bid := strtoInt(d[1])
		handType, handPoint := scanHand(hand)

		playerData = append(playerData, Player{hand, handCardValue, bid, handType, handPoint, false, 0})
	}
	return playerData
}

func scanHand(hand string) (string, int) {
	cards := make(map[rune]int)
	var similarCard []int
	for _, c := range hand {
		cards[c] += 1
	}
	for _, val := range cards {
		similarCard = append(similarCard, val)
	}

	if slices.Contains(similarCard, 5) {
		return "Five of a kind", 7
	}
	if slices.Contains(similarCard, 4) {
		return "Four of a kind", 6
	}
	if slices.Contains(similarCard, 3) && slices.Contains(similarCard, 2) {
		return "Full house", 5
	}
	if slices.Contains(similarCard, 3) {
		return "Three of a kind", 4
	}
	if slices.Contains(similarCard, 2) && len(similarCard) == 3 {
		return "Two pair", 3
	}
	if slices.Contains(similarCard, 2) && len(similarCard) == 4 {
		return "One pair", 2
	}
	return "High card", 1
}

func getRank(players PlayerData) []Player {

	results := make(map[int][]Player)
	maxRank := len(players)
	fmt.Println("max Rank: ", maxRank)

	var rank int = 1

	//Map each player by the Type of hand he had ( and the point it gave)
	for _, hand := range players {
		results[hand.handPoint] = append(results[hand.handPoint], hand)
	}

	// We will compare hand by hand 1 point = High card , 2 point = One pair ... 7 point = Five of a kind
	for point := 1; point <= 7; point++ {

		// Break when all players has been ranked
		if rank > maxRank {
			break
		}

		// Pass if no one got a type of card combo
		arrPLayers, ok := results[point]
		if !ok {
			continue
		}

		// fmt.Printf("Looking at: %d , value: %v \n", point , arrPLayers)

		// The player is alone to have a type of card -> we can give him a rank and pass to the next Hand Type
		if len(arrPLayers) == 1 {
			arrPLayers[0].rank = rank
			rank += 1
			// fmt.Printf("Looking at: %d , value: %v, rank %d \n", point , arrPLayers, arrPLayers[0].rank)
			continue
		}

		fmt.Printf("Looking at: %d , value: %v \n", point, arrPLayers)
		// We need to compare each player hands card by card

		pArrPLayers := &arrPLayers
		err := orderPlayerPerCardValue(pArrPLayers, 0)
		if err != nil {
			fmt.Errorf("we got an error in the ranking per card")
		}

		for _, player := range arrPLayers {
			player.rank = rank
			rank += 1
		}
	}
	return players

}

func getArrValuePerCard(strCard string) []int {
	intCardArr := []int{}
	mapCardVal := singleCardValues()
	for _, c := range strCard {
		intCardArr = append(intCardArr, mapCardVal[string(c)])
	}
	return intCardArr
}

// orderPlayerPerCardValue Take an array of Player and will sort them depending of there cards labels, considere that many players can get involved
func orderPlayerPerCardValue(orderArr *[]Player, start int) error {

	mapArr := make(map[int][]Player)

	if start > 5 {
		err := errors.New("to much recursion")
		return err
	}

	for _, player := range *orderArr {
		mapArr[player.handValuePerCard[start]] = append(mapArr[player.handValuePerCard[start]], player)

	}
	fmt.Println(mapArr)

	for i := 15; i > 0; i-- {
		arr, ok := mapArr[i]
		if !ok {
			continue
		}
		if len(arr) == 1 {
			*orderArr = append(*orderArr, mapArr[i][0])
		} else {
			err := orderPlayerPerCardValue(orderArr, start+1)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
