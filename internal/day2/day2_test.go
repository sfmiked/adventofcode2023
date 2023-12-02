package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartAsInt(t *testing.T) {

	g := Game{0, []Play{
		{4, 3, 0},
		{1, 6, 2},
		{green: 2},
	}}

	min := g.MinPlayOfPlays()

	assert.Equal(t, 4, min.red)
	assert.Equal(t, 6, min.blue)
	assert.Equal(t, 2, min.green)
}

func TestParseGame(t *testing.T) {

	_, err := ParseGame("Game")
	assert.NotNil(t, err)

	_, err = ParseGame("Game XX: 3 blue, 4 red")
	assert.NotNil(t, err)

	_, err = ParseGame("Game 5: YY blue, 4 red")
	assert.NotNil(t, err)

	_, err = ParseGame("Game 5: 4 blue, 4")
	assert.NotNil(t, err)

	game, err := ParseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	assert.Nil(t, err)

	min := game.MinPlayOfPlays()

	assert.Equal(t, 4, min.red)
	assert.Equal(t, 6, min.blue)
	assert.Equal(t, 2, min.green)
}

func TestSumCalc(t *testing.T) {

	game, err := ParseGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")
	assert.Nil(t, err)

	assert.Equal(t, 12, game.Power())
	assert.Equal(t, 2, game.Id())

	red, green, blue := game.DicesRequired()

	assert.Equal(t, 1, red)
	assert.Equal(t, 4, blue)
	assert.Equal(t, 3, green)

}

func TestLoadGameFile(t *testing.T) {

	games, err := LoadGameFile("testdata/input.txt")
	assert.Nil(t, err)

	id_sum, power_sum := 0, 0
	for _, g := range games {
		id_sum += g.id
		power_sum += g.Power()
	}
	assert.Equal(t, 15, id_sum)
	assert.Equal(t, 2286, power_sum)

	_, err = LoadGameFile("testdata/bad_input.txt")
	assert.NotNil(t, err)
}
