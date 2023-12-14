package day14

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTilt(t *testing.T) {

	rows := strings.Split(`...
.#.
OOO`, "\n")

	expected := `O.O
.#.
.O.`

	dish := DishFromRows(rows)

	dish.Tilt(North)

	assert.Equal(t, expected, dish.ToString())

	assert.Equal(t, 7, dish.Load())

}

func TestSpin(t *testing.T) {

	p, err := LoadPlatform("testdata/input.txt")
	assert.Nil(t, err)

	after_1, after_2, after_3 := /* 1 */ `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`, /* 2 */ `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`, /* 3 */ `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`

	p.Spin()
	assert.Equal(t, after_1, p.ToString())

	p.Spin()
	assert.Equal(t, after_2, p.ToString())

	p.Spin()
	assert.Equal(t, after_3, p.ToString())
}

func TestLoad(t *testing.T) {
	_, err := LoadPlatform("nothing.txt")
	assert.NotNil(t, err)
}

func TestPredictor(t *testing.T) {

	platform, err := LoadPlatform("testdata/input.txt")
	assert.Nil(t, err)

	predictor := platform.Predictor()
	result := predictor.PredictLoad(1_000_000_000)

	assert.Equal(t, 64, result)
}

func TestPredictorWithoutPrediction(t *testing.T) {

	platform, err := LoadPlatform("testdata/input.txt")
	assert.Nil(t, err)

	predictor := platform.Predictor()
	result := predictor.PredictLoad(1)

	assert.Equal(t, 87, result)
}
