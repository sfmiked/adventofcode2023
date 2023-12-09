package day8

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaySingleGame(t *testing.T) {

	// happy path, single game
	book, err := LoadFile("testdata/input.txt")
	assert.Nil(t, err)
	result, err := book.Play("AAA", func(input string) bool {
		return input == "ZZZ"
	})
	assert.Nil(t, err)
	assert.Equal(t, 2, result)
}

func TestPlayWithInvalidMove(t *testing.T) {
	// single game with an invalid initial move (YYY is not in the deck)
	book, err := LoadFile("testdata/input.txt")
	assert.Nil(t, err)
	_, err = book.Play("YYY", func(input string) bool {
		return input == "ZZZ"
	})
	assert.NotNil(t, err)
}

func TestPlayWithSuffixMatch(t *testing.T) {

	// single game
	book, err := LoadFile("testdata/alternate.txt")
	assert.Nil(t, err)
	result, err := book.Play("11A", func(input string) bool {
		return strings.HasSuffix(input, "Z")
	})
	assert.Nil(t, err)
	assert.Equal(t, 2, result)

	// single game
	result, err = book.Play("22A", func(input string) bool {
		return strings.HasSuffix(input, "Z")
	})
	assert.Nil(t, err)
	assert.Equal(t, 3, result)

	// combined games
	result, err = book.PlayEnding("Z")

	assert.Nil(t, err)
	assert.Equal(t, 6, result)
}

func TestInvalidFiles(t *testing.T) {

	// this file contains bad moves, which will error out
	book, err := LoadFile("testdata/bad.txt")
	assert.Nil(t, err)
	_, err = book.Play("11A", func(input string) bool {
		return strings.HasSuffix(input, "Z")
	})
	assert.NotNil(t, err)

	// should also fail for multiple-kinds
	_, err = book.PlayEnding("Z")
	assert.NotNil(t, err)

	// this file is complete, and should not load correctly.
	_, err = LoadFile("testdata/incomplete.txt")
	assert.NotNil(t, err)
}
