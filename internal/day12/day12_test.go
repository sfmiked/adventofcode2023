package day12

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrangements(t *testing.T) {

	cases := []struct {
		springs  string
		cases    []int
		expected int
	}{
		{"???.###", []int{1, 1, 3}, 1},
		{".??..??...?##.", []int{1, 1, 3}, 4},
		{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1},
		{"????.#...#...", []int{4, 1, 1}, 1},
		{"????.######..#####.", []int{1, 6, 5}, 4},
		{"?###????????", []int{3, 2, 1}, 10},
	}

	for ix, c := range cases {
		result := ValidArrangements(c.springs, c.cases)
		log.Printf("%d: %s %v -> %d", ix, c.springs, c.cases, result)

		assert.Equal(t, c.expected, result)
	}

}

func TestArrangementsFromFile(t *testing.T) {

	// first part
	result, err := CheckArrangementsFromFile("testdata/input.txt", false)
	assert.Nil(t, err)
	assert.Equal(t, 21, result)

	// second part
	result, err = CheckArrangementsFromFile("testdata/input.txt", true)
	assert.Nil(t, err)
	assert.Equal(t, 525152, result)

	_, err = CheckArrangementsFromFile("testdata/bad_file.txt", false)
	assert.NotNil(t, err)
}

func TestExpand(t *testing.T) {

	s1, g1, e1 := parseInput("???.### 1,1,3", true)
	s2, g2, e2 := parseInput("???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3", false)

	assert.Equal(t, s1, s2)
	assert.Equal(t, g1, g2)
	assert.Nil(t, e1)
	assert.Nil(t, e2)

}
