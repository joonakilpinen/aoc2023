package day7part2

import (
	"aoc2023/utils"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

type Card int

const (
	T Card = 10
	J Card = 1
	Q Card = 12
	K Card = 13
	A Card = 14
)

func convertToCard(input string) Card {
	num, err := strconv.Atoi(input)
	if err == nil {
		return Card(num)
	}
	switch input {
	case "T":
		return T
	case "J":
		return J
	case "Q":
		return Q
	case "K":
		return K
	case "A":
		return A
	default:
		panic(errors.New(fmt.Sprintf("unknown card value %s", input)))
	}
}

type Hand struct {
	Cards []Card
	Bid   int
}

type HandType int

// Every hand is exactly one type. From strongest to weakest, they are:
//   - Five of a kind, where all five cards have the same label: AAAAA
//   - Four of a kind, where four cards have the same label and one card
//     has a different label: AA8AA
//   - Full house, where three cards have the same label, and the
//     remaining two cards share a different label: 23332
//   - Three of a kind, where three cards have the same label, and the
//     remaining two cards are each different from any other card in the
//     hand: TTT98
//   - Two pair, where two cards share one label, two other cards share a
//     second label, and the remaining card has a third label: 23432
//   - One pair, where two cards share one label, and the other three
//     cards have a different label from the pair and each other: A23A4
//   - High card, where all cards' labels are distinct: 23456
const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func cardMapHas(cardMap map[Card]int, val int) bool {
	for _, v := range cardMap {
		if v == val {
			return true
		}
	}
	return false
}

func (h Hand) getType() HandType {
	cardMap := map[Card]int{}
	for _, card := range h.Cards {
		cardMap[card]++
	}
	jokers := cardMap[J]
	var bestCard Card
	highestVal := 0
	for card, val := range cardMap {
		if card == J {
			continue
		}
		if val > highestVal {
			highestVal = val
			bestCard = card
		}
	}
	if jokers < 5 {
		cardMap[bestCard] += jokers
		delete(cardMap, J)
	}
	if len(cardMap) == 5 {
		return HighCard
	}
	if len(cardMap) == 1 {
		return FiveOfAKind
	}
	if len(cardMap) == 4 {
		return OnePair
	}
	if cardMapHas(cardMap, 4) {
		return FourOfAKind
	}
	if cardMapHas(cardMap, 3) && cardMapHas(cardMap, 2) {
		return FullHouse
	}
	if cardMapHas(cardMap, 3) {
		return ThreeOfAKind
	}
	return TwoPair
}

func (h Hand) compareTo(other Hand) bool {
	if h.getType() != other.getType() {
		return h.getType() > other.getType()
	}
	for i, card := range h.Cards {
		if card != other.Cards[i] {
			return card > other.Cards[i]
		}
	}
	return false
}

func parseHands(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		var cards []Card
		for _, s := range strings.Split(split[0], "") {
			cards = append(cards, convertToCard(s))
		}
		hands = append(hands, Hand{Cards: cards, Bid: bid})
	}
	return hands
}

func (Solver) Solve(input string) string {
	hands := parseHands(utils.GetLines(input))
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].compareTo(hands[j])
	})
	rank := len(hands)
	result := 0
	for _, hand := range hands {
		result += hand.Bid * rank
		rank--
	}
	return strconv.Itoa(result)
}

//
//    To make things a little more interesting, the Elf introduces one
//    additional rule. Now, J cards are jokers - wildcards that can act like
//    whatever card would make the hand the strongest type possible.
//
//    To balance this, J cards are now the weakest individual cards, weaker
//    even than 2. The other cards stay in the same order: A, K, Q, T, 9, 8,
//    7, 6, 5, 4, 3, 2, J.
//
//    J cards can pretend to be whatever card is best for the purpose of
//    determining hand type; for example, QJJQ2 is now considered four of a
//    kind. However, for the purpose of breaking ties between two hands of
//    the same type, J is always treated as J, not the card it's pretending
//    to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.
//
//    Now, the above example goes very differently:
// 32T3K 765
// T55J5 684
// KK677 28
// KTJJT 220
// QQQJA 483
//
//      * 32T3K is still the only one pair; it doesn't contain any jokers, so
//        its strength doesn't increase.
//      * KK677 is now the only two pair, making it the second-weakest hand.
//      * T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank
//        3, QQQJA gets rank 4, and KTJJT gets rank 5.
//
//    With the new joker rule, the total winnings in this example are 5905.
//
//    Using the new joker rule, find the rank of every hand in your set. What
//    are the new total winnings?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
