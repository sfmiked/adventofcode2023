package day4

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/miketzian/adventofcode2023/internal/util"
)

type Game struct {
	cards       []*Scratchcard
	card_counts []int
}

func (game *Game) ApplyReverseRules() {
	for ix, v := range game.cards {
		if matched := v.MatchingNos(); matched > 0 {
			// fmt.Printf("Game %d, matched %d\n", v.card_no, matched)
			for i := ix + 1; i < ix+matched+1 && ix < len(game.cards); i++ {
				game.card_counts[i] += game.card_counts[ix]
			}
			// fmt.Printf("%v", game.card_counts)
		}
	}
}

func (game *Game) TotalScore() int {
	total := 0
	for _, v := range game.cards {
		total += v.Score()
	}
	return total
}

func (game *Game) TotalCards() int {
	total := 0
	for _, v := range game.card_counts {
		total += v
	}
	return total
}

func LoadGame(input string) (*Game, error) {

	cards, err := util.ReadFileLines(input, func(card_line string) (*Scratchcard, error) {
		card, err := ParseScratchcard(card_line)
		if err != nil {
			return nil, err
		}
		return card, nil
	})

	if err != nil {
		return nil, err
	}
	card_counts := make([]int, len(cards))
	for ix := range card_counts {
		card_counts[ix] = 1
	}
	return &Game{
		cards, card_counts,
	}, nil

}

type Scratchcard struct {
	card_no     int
	winning_nos []int
	my_nos      []int
}

func (card *Scratchcard) MatchingNos() int {
	matching := 0
	for _, no := range card.my_nos {
		if card.IsWinningNo(no) {
			matching += 1
			fmt.Printf("%d is winning\n", no)
		}
	}
	return matching
}

func (card *Scratchcard) IsWinningNo(num int) bool {
	return util.IndexOf(card.winning_nos, num) != -1
}

func (card *Scratchcard) Score() int {
	score := 0
	for _, no := range card.my_nos {
		if card.IsWinningNo(no) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func ParseScratchcard(input string) (*Scratchcard, error) {

	var s scanner.Scanner

	s.Init(strings.NewReader(input))
	// card
	if s.Scan() == scanner.EOF {
		return nil, errors.New("empty input string")
	}
	// fmt.Printf("Token [%s]\n", s.TokenText())
	if s.Scan() == scanner.EOF {
		return nil, errors.New("invalid format")
	}
	// fmt.Printf("Token [%s]\n", s.TokenText())

	token := s.TokenText()

	card_no, err := strconv.ParseInt(token, 10, 32)
	if err != nil {
		return nil, err
	}

	if s.Scan() != ':' {
		return nil, errors.New("invalid input")
	}

	winning_nos := make([]int, 0)
	my_nos := make([]int, 0)
	found_pipe := false

	// for tok := s.Scan(); tok != scanner.EOF {
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {

		if tok == '|' {
			found_pipe = true
			continue
		}

		// fmt.Printf("Card Token [%s] [%v]\n", s.TokenText(), found_pipe)

		no, err := strconv.ParseInt(s.TokenText(), 10, 32)
		if err != nil {
			return nil, err
		}
		if found_pipe {
			my_nos = append(my_nos, int(no))
		} else {
			winning_nos = append(winning_nos, int(no))
		}
	}

	if !found_pipe {
		return nil, errors.New("text ended early")
	}

	card := Scratchcard{
		int(card_no), winning_nos, my_nos,
	}
	return &card, nil
	// return nil, errors.New("not implemented")
}
