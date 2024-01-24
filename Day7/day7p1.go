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
	bytesText, err := os.ReadFile("./Day7/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(bytesText)

	result := analyzep1(puzzle)

	fmt.Println("Day 7 part 2 result is : ")
	return result
}

func analyzep1(s string) int {
	lines := strings.Split(s, "\n")

	playerData := getPlayersDatas(lines)

	pPlayerData := &playerData

	acc := getRank(pPlayerData)

	// for _, hand := range playerData {
	// 	fmt.Printf("hand: %s bid:%d type: %s handpoint: %d rank %d \n", hand.hand, hand.bid, hand.handType, hand.handPoint, hand.rank)
	// }

	return acc

}

func singleCardValues() map[string]int {
	labels := []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

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
	num_Joker := 0
	var similarCard []int

	for _, c := range hand {
		cards[c] += 1
		if string(c) == "J"{
			num_Joker +=1
		}
	}
	for _, val := range cards {
		similarCard = append(similarCard, val)
	}

	if slices.Contains(similarCard, 5) {
		return "Five of a kind", 7
	}
	if slices.Contains(similarCard, 4) {

		if num_Joker == 1 || num_Joker == 4{
			return "Five of a kind", 7
		}
		return "Four of a kind", 6
	}
	if slices.Contains(similarCard, 3) && slices.Contains(similarCard, 2) {
		if num_Joker == 2 || num_Joker == 3{
			return "Five of a kind", 7
		}
		return "Full house", 5
	}
	if slices.Contains(similarCard, 3) {
		if num_Joker == 1 || num_Joker == 3{
			return "Four of a kind", 6
		}
		return "Three of a kind", 4
	}
	if slices.Contains(similarCard, 2) && len(similarCard) == 3 {
		if num_Joker == 1{
			return "Full house", 5
		}
		if num_Joker == 2{
			return "Four of a kind", 6
		}

		return "Two pair", 3
	}
	if slices.Contains(similarCard, 2) && len(similarCard) == 4 {
		if num_Joker == 1 || num_Joker == 2{
			return "Three of a kind", 4
		}

		if num_Joker == 2{
			return "Four of a kind", 6
		}
		return "One pair", 2
	}

	if num_Joker == 1 {
		return "One pair", 2
	}
	return "High card", 1
}

func getRank(players *PlayerData) int {
	accumulator := 0

	results := make(map[int][]Player)
	maxRank := len(*players)
	fmt.Println("max Rank: ", maxRank)

	var currRank int = 1

	//Map each player by the Type of hand he had ( and the point it gave)
	for _, player := range *players {
		results[player.handPoint] = append(results[player.handPoint], player)
	}

	// We will compare hand by hand 1 point = High card , 2 point = One pair ... 7 point = Five of a kind
	for point := 1; point <= 7; point++ {
		// fmt.Println("== Checking for Hand point = ", point)

		// Break when all players has been ranked
		if currRank > maxRank {
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
			for _, pl := range *players {
				if pl.hand == arrPLayers[0].hand {
					pl.rank = currRank
					// fmt.Printf("Looking at: %s rank %d \n", pl.hand, pl.rank)
					accumulator += (pl.rank*pl.bid)
					break
				}
			}
			currRank += 1
			continue
		}

		// fmt.Printf("Array of players at: %d , value: %v \n", point, arrPLayers)
		// We need to compare each player hands card by card
		var orderedPlayerArray []Player
		pOrderedPlayerArray := &orderedPlayerArray
		
		err := orderPlayerPerCardValue(arrPLayers, pOrderedPlayerArray, 0)
		if err != nil {
			fmt.Errorf("we got an error in the ranking per card")
		}
		for i, j := 0, len(orderedPlayerArray)-1; i < j; i, j = i+1, j-1 {
			orderedPlayerArray[i], orderedPlayerArray[j] = orderedPlayerArray[j], orderedPlayerArray[i]
		}
		
		for _, player := range orderedPlayerArray {
			// fmt.Println("Player from arrPLayers ", player)
			for _, pl := range *players {
				if pl.hand == player.hand {
					pl.rank = currRank
					accumulator += (pl.rank*pl.bid)
					break
				}
			}
			currRank += 1
		}
	}

	return accumulator

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
func orderPlayerPerCardValue(notOrderPlayerArr []Player, inOrderPlayerArr *[]Player,start int) error {

	mapPerCard := make(map[int][]Player)

	if start > 5 {
		err := errors.New("to much recursion")
		return err
	}

	for _, player := range notOrderPlayerArr {
		mapPerCard[player.handValuePerCard[start]] = append(mapPerCard[player.handValuePerCard[start]], player)
	}
	// fmt.Println("Ordering : ", mapPerCard)

	for i := 15; i > 0; i-- {

		arr, ok := mapPerCard[i]
		if !ok {
			continue
		}

		if len(arr) == 1 {
			*inOrderPlayerArr= append(*inOrderPlayerArr, mapPerCard[i][0])
		} else {
			err := orderPlayerPerCardValue(arr, inOrderPlayerArr,start+1)
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


