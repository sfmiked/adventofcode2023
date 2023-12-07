package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeHeirarchy(t *testing.T) {

	assert.Less(t, high_card, one_pair)
	assert.Less(t, one_pair, two_pair)
	assert.Less(t, two_pair, three_o_kind)
	assert.Less(t, three_o_kind, full_house)
	assert.Less(t, full_house, four_o_kind)
	assert.Less(t, four_o_kind, five_o_kind)
}

func TestRules(t *testing.T) {

	// if we're not playing with Jokers, 'J' is worth more than 9
	testRule := rules(false)
	assert.Greater(t, testRule['A'], testRule['K'])
	assert.Greater(t, testRule['K'], testRule['Q'])
	assert.Greater(t, testRule['Q'], testRule['J'])
	assert.Greater(t, testRule['J'], testRule['T'])
	assert.Greater(t, testRule['T'], testRule['9'])
	assert.Greater(t, testRule['9'], testRule['8'])
	assert.Greater(t, testRule['8'], testRule['7'])
	assert.Greater(t, testRule['7'], testRule['6'])
	assert.Greater(t, testRule['6'], testRule['5'])
	assert.Greater(t, testRule['5'], testRule['4'])
	assert.Greater(t, testRule['4'], testRule['3'])
	assert.Greater(t, testRule['3'], testRule['2'])
	assert.Greater(t, testRule['2'], testRule['0'])

	// if we're playing with Jokers, 'J' is worth less than 1
	testRule = rules(true)
	assert.Less(t, testRule['J'], testRule['2'])

	assert.Greater(t, testRule['A'], testRule['K'])
	assert.Greater(t, testRule['K'], testRule['Q'])
	assert.Greater(t, testRule['Q'], testRule['T'])
	assert.Greater(t, testRule['T'], testRule['9'])
	assert.Greater(t, testRule['9'], testRule['8'])
	assert.Greater(t, testRule['8'], testRule['7'])
	assert.Greater(t, testRule['7'], testRule['6'])
	assert.Greater(t, testRule['6'], testRule['5'])
	assert.Greater(t, testRule['5'], testRule['4'])
	assert.Greater(t, testRule['4'], testRule['3'])
	assert.Greater(t, testRule['3'], testRule['2'])
	assert.Greater(t, testRule['2'], testRule['J'])
	assert.Greater(t, testRule['J'], testRule['0'])
}

func TestHands(t *testing.T) {

	cases := []struct {
		hand          string
		kind          HandType
		kind_if_joker HandType
	}{
		{"AAAAA", five_o_kind, five_o_kind},
		{"AAAAJ", four_o_kind, five_o_kind},
		{"AAAJJ", full_house, five_o_kind},
		{"AAJJJ", full_house, five_o_kind},

		{"AAAAX", four_o_kind, four_o_kind},
		{"AAAJX", three_o_kind, four_o_kind},
		{"JJJAX", three_o_kind, four_o_kind},

		{"AAKKJ", two_pair, full_house},
		{"AAKKX", two_pair, two_pair},
		{"AAKJJ", two_pair, four_o_kind},

		{"AAQTJ", one_pair, three_o_kind},

		{"AKQTJ", high_card, one_pair},
		{"AKQT9", high_card, high_card},
	}

	game := NewGame(false)
	jokerGame := NewGame(true)

	for ix, testCase := range cases {

		assert.Equal(t, testCase.kind, game.handType([]rune(testCase.hand)), "%d / %s failed", ix, testCase.hand)
		assert.Equal(t, testCase.kind_if_joker, jokerGame.handType([]rune(testCase.hand)), "%d / %s failed (joker)", ix, testCase.hand)
	}
}

func TestLoad(t *testing.T) {

	game := NewGame(false)
	err := game.LoadHands("testdata/bad_file.txt")
	assert.NotNil(t, err)

	err = game.LoadHands("testdata/bad_file_2.txt")
	assert.NotNil(t, err)

	err = game.LoadHands("testdata/weird_but_ok.txt")
	assert.Nil(t, err)

	err = game.LoadHands("testdata/input.txt")
	assert.Nil(t, err)

}

func TestPlay(t *testing.T) {

	game := NewGame(false)
	err := game.LoadHands("testdata/input.txt")
	assert.Nil(t, err)

	assert.Equal(t, 6440, game.TotalWinnings())

	// playing with jokers
	game = NewGame(true)
	err = game.LoadHands("testdata/input.txt")
	assert.Nil(t, err)

	assert.Equal(t, 5905, game.TotalWinnings())
}
