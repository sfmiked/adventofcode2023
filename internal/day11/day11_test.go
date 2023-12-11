package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	obs, err := NewObservations("testdata/input.txt", 2)

	assert.Nil(t, err)
	assert.Equal(t, 9, len(obs.Galaxies()))

	// from galaxy doesn't exist
	_, err = obs.GalaxyDistance(125, 3)
	assert.NotNil(t, err)

	// to galaxy doesn't exist
	_, err = obs.GalaxyDistance(5, 72)
	assert.NotNil(t, err)

	cases := []struct {
		fromGalaxy int
		toGalaxy   int
		distance   int
		eValue     int
	}{
		{5, 9, 9, 2},
		{9, 5, 9, 2},
		{1, 7, 15, 2},
		{3, 6, 17, 2},
		{8, 9, 5, 2},
		{2, 8, 19, 2},
	}

	for _, test := range cases {
		obs.expandValue = test.eValue

		dist, err := obs.GalaxyDistance(test.fromGalaxy, test.toGalaxy)
		assert.Nil(t, err)
		assert.Equal(t, test.distance, dist)
	}

}

func TestParse(t *testing.T) {

	input := `...
..#
.#.`

	parsed := [][]string{
		{".", ".", "."},
		{".", ".", "#"},
		{".", "#", "."},
	}

	result := gridFromString(input)

	assert.Equal(t, parsed, result)

	expected := [][]string{
		{"e", "e", "e"},
		{"e", ".", "#"},
		{"e", "#", "."},
	}

	markExpandingRegions(result)

	assert.Equal(t, expected, result)

}

func TestDistance(t *testing.T) {

	input := gridFromString(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`)

	markExpandingRegions(input)

	exp := gridFromString(`..e#.e..e.
..e..e.#e.
#.e..e..e.
eeeeeeeeee
..e..e#.e.
.#e..e..e.
..e..e..e#
eeeeeeeeee
..e..e.#e.
#.e.#e..e.`)

	// verify that this maps right
	assert.Equal(t, input, exp)

	mapGalaxies(input)

	withg := gridFromString(`..e1.e..e.
..e..e.2e.
3.e..e..e.
eeeeeeeeee
..e..e4.e.
.5e..e..e.
..e..e..e6
eeeeeeeeee
..e..e.7e.
8.e.9e..e.`)

	assert.Equal(t, input, withg)

}

func TestFileErrors(t *testing.T) {
	_, err := NewObservations("testdata/not_exist.txt", 2)
	assert.NotNil(t, err)
}
