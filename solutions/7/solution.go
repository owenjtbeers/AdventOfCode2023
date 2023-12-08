package sol7

import (
	mylib "aoc/lib"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var handRanks = map[string]int{
	"fiveOfAKind":  6,
	"fourOfAKind":  5,
	"fullHouse":    4,
	"threeOfAKind": 3,
	"twoPair":      2,
	"onePair":      1,
	"highCard":     0,
}

var cardRankOrder = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type HandInfo struct {
	cards string
	rank  int
	bid   int
}

func Sol7() int {
	scanner := mylib.ReadLines("./solutions/7/input.txt")

	index := 0
	mapHandRanks := make(map[int][]HandInfo)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := strings.Split(line, " ")
		hand := parsedLine[0]
		bid, _ := strconv.Atoi(parsedLine[1])
		handRank := getHandRank(hand)
		handInfo := HandInfo{hand, handRank, bid}
		mapHandRanks[handRank] = append(mapHandRanks[handRank], handInfo)
		index++
	}
	handRankInts := getSortedValuesAscending(handRanks)
	rankedHands := []HandInfo{}
	for _, value := range handRankInts {
		sortedHands := sortHands(mapHandRanks[value])
		rankedHands = append(rankedHands, sortedHands...)
	}
	sum := 0
	for rank, handInfo := range rankedHands {
		sum += handInfo.bid * (rank + 1)
		fmt.Println(rank+1, handInfo)
	}
	return sum
}

type CardCount struct {
	card  string
	count int
}

func getHandRank(hand string) int {
	handMap := make(map[string]int)
	for i := 0; i < len(hand); i++ {
		handMap[string(hand[i])]++
	}
	// Jokers act as the highest winning card
	// We can just sort the hand and then increment the count of the highest card by the number of jokers
	handMatches := []int{}
	jokerCount := handMap["J"]
	sortedListOfCardCounts := []CardCount{}
	for key, value := range handMap {
		sortedListOfCardCounts = append(sortedListOfCardCounts, CardCount{key, value})
	}
	sort.Slice(sortedListOfCardCounts, func(i, j int) bool { return sortedListOfCardCounts[i].count > sortedListOfCardCounts[j].count })
	if sortedListOfCardCounts[0].card != "J" {
		handMap[sortedListOfCardCounts[0].card] += jokerCount
		delete(handMap, "J")
	} else if jokerCount != 5 {
		handMap[sortedListOfCardCounts[1].card] += jokerCount
		delete(handMap, "J")
	}
	// fmt.Println(sortedListOfCardCounts)
	for key, value := range handMap {
		if value == 5 {
			handMatches = append(handMatches, handRanks["fiveOfAKind"])
		}
		if value == 4 {
			handMatches = append(handMatches, handRanks["fourOfAKind"])
		}
		if value == 3 {
			for _, value2 := range handMap {
				if value2 == 2 {
					handMatches = append(handMatches, handRanks["fullHouse"])
				}
			}
			handMatches = append(handMatches, handRanks["threeOfAKind"])
		}
		if value == 2 {
			for key2, value2 := range handMap {
				if value2 == 2 && key2 != key {
					handMatches = append(handMatches, handRanks["twoPair"])
				}
			}
			handMatches = append(handMatches, handRanks["onePair"])
		}

	}
	if len(handMatches) == 0 {
		return handRanks["highCard"]
	}
	return sortIntsDescending(handMatches)[0]
}

func sortHands(hands []HandInfo) []HandInfo {
	sort.Slice(hands, func(i, j int) bool {
		handOne := hands[i].cards
		handTwo := hands[j].cards
		for cardIndex := 0; cardIndex < len(handOne); cardIndex++ {
			handOneCard := string(handOne[cardIndex])
			handTwoCard := string(handTwo[cardIndex])
			if cardRankOrder[handOneCard] < cardRankOrder[handTwoCard] {
				return true
			} else if cardRankOrder[handOneCard] != cardRankOrder[handTwoCard] {
				return false
			}
		}
		return false
	})
	return hands
}
func sortIntsDescending(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return nums
}

func sortIntsAscending(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums
}

func getSortedValuesAscending(handMap map[string]int) []int {
	values := []int{}
	for _, value := range handMap {
		values = append(values, value)
	}
	return sortIntsAscending(values)
}
