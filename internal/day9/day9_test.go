package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtrapolate(t *testing.T) {

	cases := []struct {
		data         []int
		expected     int
		expectedLeft int
	}{
		{[]int{0, 3, 6, 9, 12, 15}, 18, -3},
		{[]int{1, 3, 6, 10, 15, 21}, 28, 0},
		{[]int{10, 13, 16, 21, 30, 45}, 68, 5},
	}

	for _, c := range cases {

		result := Play(c.data, false)
		assert.Equal(t, c.expected, result)

		result = Play(c.data, true)
		assert.Equal(t, c.expectedLeft, result)
	}
}

func TestPlayFile(t *testing.T) {

	result, err := PlayFile("testdata/input.txt", false)
	assert.Nil(t, err)
	assert.Equal(t, 114, result)

	result, err = PlayFile("testdata/input.txt", true)
	assert.Nil(t, err)
	assert.Equal(t, 2, result)

	_, err = PlayFile("testdata/no_file.txt", true)
	assert.NotNil(t, err)

	_, err = PlayFile("testdata/bad_file.txt", true)
	assert.NotNil(t, err)
}
