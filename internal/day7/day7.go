package day7

import (
	"errors"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
	"github.com/samber/lo"
)

type HandType uint8

// http://golang.org/ref/spec#Iota
const (
	high_card    HandType = 1 << iota
	one_pair     HandType = 1 << iota
	two_pair     HandType = 1 << iota
	three_o_kind HandType = 1 << iota
	full_house   HandType = 1 << iota
	four_o_kind  HandType = 1 << iota
	five_o_kind  HandType = 1 << iota
)

func rules(withJokers bool) map[rune]int {

	var arr []string
	if withJokers {
		arr = strings.Split("A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J, 0", ", ")
	} else {
		arr = strings.Split("A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2, 0", ", ")
	}
	return lo.Reduce(arr, func(agg map[rune]int, item string, index int) map[rune]int {
		agg[[]rune(item)[0]] = len(arr) - index // element to the left are worth more.
		return agg
	}, make(map[rune]int, 0))
}

type Hand struct {
	hand []rune
	bet  int
	kind HandType
}

type CamelCards struct {
	withJokers bool
	rules      map[rune]int
	hands      []*Hand
}

// these methods make the CamelCards struct itself sortable
// rather than the inner array of hands
// this works better as we have access to the 'withJokers' variable anyway
// based loosely on the planetSorter
// https://cs.opensource.google/go/go/+/refs/tags/go1.21.5:src/sort/example_keys_test.go

func (s *CamelCards) Len() int {
	return len(s.hands)
}

func (s *CamelCards) Swap(i, j int) {
	s.hands[i], s.hands[j] = s.hands[j], s.hands[i]
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (s *CamelCards) Less(i, j int) bool {
	// return s.by(s.hands[i], s.hands[j])

	if s.hands[i].kind < s.hands[j].kind {
		return true
	}
	if s.hands[i].kind > s.hands[j].kind {
		return false
	}

	// same kind
	g1, g2 := s.hands[i].hand, s.hands[j].hand

	for i, g1v := range g1 {
		g2v := g2[i]

		if g1v == g2v {
			// same character
			continue
		}
		// else pick the highest one
		g1vi := s.rules[g1v]
		g2vi := s.rules[g2v]
		if g1vi < g2vi {
			return true
		}
		return false
	}
	return false
}

func (game *CamelCards) TotalWinnings() int {

	// sorted on load, otherwise you could do it here
	// sort.Sort(game)

	return lo.Reduce(game.hands, func(agg int, item *Hand, ix int) int {
		// log.Printf("hand: %v, ix: %d", item, ix)
		return agg + ((ix + 1) * item.bet)
	}, 0)
}

func NewGame(withJokers bool) *CamelCards {

	game := CamelCards{withJokers, rules(withJokers), make([]*Hand, 0)}

	return &game
}

func (game *CamelCards) LoadHands(input string) error {

	hands, err := util.ReadFileLines(input, func(line string) (*Hand, error) {
		segs := strings.Split(line, " ")
		if len(segs) != 2 {
			return nil, errors.New("invalid file")
		}
		bet, err := strconv.ParseInt(segs[1], 10, 32)
		if err != nil {
			return nil, err
		}
		return game.NewHand(segs[0], int(bet), game.withJokers), nil
	})
	if err != nil {
		return err
	}
	game.hands = hands
	log.Printf("Loaded %d hands\n", len(hands))

	sort.Sort(game)

	return nil
}

func (game *CamelCards) NewHand(cards string, bet int, jokers bool) *Hand {
	data := []rune(cards)
	return &Hand{hand: data, bet: bet, kind: game.handType(data)}
}

func (game *CamelCards) handType(hand []rune) HandType {

	_, hasJoker := lo.Find(hand, func(item rune) bool {
		return item == 'J'
	})

	counts := lo.Reduce(hand, func(agg map[rune]int, item rune, _ int) map[rune]int {
		if c, ok := agg[item]; ok {
			agg[item] = c + 1
		} else {
			agg[item] = 1
		}
		return agg
	}, make(map[rune]int))

	// number of different cards
	cards := len(counts)
	if cards == 1 {
		return five_o_kind
	} else if cards == 2 {
		// JJJJ X or XXXX J or JJJ XX or YYY JJ
		if game.withJokers && hasJoker {
			return five_o_kind
		}
		for _, v := range counts {
			if v == 4 {
				return four_o_kind
			}
		}
		return full_house
	} else if cards == 3 {
		for _, v := range counts {
			if v == 3 {
				// JJJ X Y or YYY J X
				if game.withJokers && hasJoker {
					return four_o_kind
				} else {
					/// XXX Y Z
					return three_o_kind
				}
			}
		}
		if game.withJokers && hasJoker {
			if counts['J'] == 2 {
				// XX JJ Z
				return four_o_kind
			} else {
				// XX YY J or  -> XXJ YY
				return full_house
			}
		} else {
			// // XX YY Z
			return two_pair
		}
	} else if cards == 4 {
		if game.withJokers && hasJoker {
			// J X Y ZZ or JJ X Y Z
			return three_o_kind
		}
		return one_pair
	} else {
		// cards == 5
		if game.withJokers && hasJoker {
			// W X Y Z J
			return one_pair
		}
		return high_card
	}
}
