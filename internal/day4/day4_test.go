package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCard(t *testing.T) {

	card, err := ParseScratchcard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	assert.Nil(t, err)
	assert.Equal(t, 5, len(card.winning_nos))
	assert.Equal(t, 8, len(card.my_nos))

	assert.EqualValues(t, []int{41, 48, 83, 86, 17}, card.winning_nos)
	assert.EqualValues(t, []int{83, 86, 6, 31, 17, 9, 48, 53}, card.my_nos)

	for _, v := range []int{41, 48, 83, 86, 17} {
		assert.Equal(t, card.IsWinningNo(v), true)
	}

	for x := 0; x < 1000; x = x + 1 {
		switch x {
		case 41:
		case 48:
		case 83:
		case 86:
		case 17:
			assert.Equal(t, card.IsWinningNo(x), true)
			break
		default:
			assert.Equal(t, card.IsWinningNo(x), false)
		}
	}

	assert.Equal(t, 4, card.MatchingNos())
	assert.Equal(t, 8, card.Score())
}

func TestParseErrors(t *testing.T) {

	_, err := ParseScratchcard("")
	assert.NotNil(t, err)

	_, err = ParseScratchcard("Card")
	assert.NotNil(t, err)

	_, err = ParseScratchcard("Card YY")
	assert.NotNil(t, err)

	_, err = ParseScratchcard("Card 50 154")
	assert.NotNil(t, err)

	_, err = ParseScratchcard("Card 50: 154")
	assert.NotNil(t, err)

	_, err = ParseScratchcard("Card 50: abc")
	assert.NotNil(t, err)

	_, err = ParseScratchcard("Card 50: 1 2 3 | 1 2 7 8")
	assert.Nil(t, err)
}

func TestLoadGame(t *testing.T) {

	game, err := LoadGame("testdata/input.txt")
	assert.Nil(t, err)
	assert.Equal(t, 6, game.TotalCards())
	assert.Equal(t, 13, game.TotalScore())

	game.ApplyReverseRules()

	assert.Equal(t, 30, game.TotalCards())
}

func TestLoadErrors(t *testing.T) {
	_, err := LoadGame("testdata/bad_input.txt")
	assert.NotNil(t, err)
}
