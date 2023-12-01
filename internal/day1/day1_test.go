package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartAsInt(t *testing.T) {

	intSymbols := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	value, err := findFirstSymbol("1abc2", intSymbols)
	assert.Nil(t, err)
	assert.Equal(t, "1", value)

	value, err = findFirstSymbol("pqr3stu8vwx", intSymbols)
	assert.Equal(t, "3", value)
	assert.Nil(t, err)

	value, err = findFirstSymbol("a1b2c3d4e5f", intSymbols)
	assert.Equal(t, "1", value)
	assert.Nil(t, err)

	value, err = findFirstSymbol("treb7uchet", intSymbols)
	assert.Equal(t, "7", value)
	assert.Nil(t, err)

	_, err = findFirstSymbol("nonumbers", intSymbols)
	assert.NotNil(t, err)

	_, err = findFirstSymbol("something", []string{"1"})
	assert.NotNil(t, err)

	_, err = findFirstSymbol("", intSymbols)
	assert.NotNil(t, err)

	_, err = findFirstSymbol("something", nil)
	assert.NotNil(t, err)

	_, err = findFirstSymbol("something", []string{})
	assert.NotNil(t, err)
}

func TestEndAsInt(t *testing.T) {

	intSymbols := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	value, err := findLastSymbol("1abc2", intSymbols)
	assert.Nil(t, err)
	assert.Equal(t, "2", value)

	value, err = findLastSymbol("pqr3stu8vwx", intSymbols)
	assert.Equal(t, "8", value)
	assert.Nil(t, err)

	value, err = findLastSymbol("a1b2c3d4e5f", intSymbols)
	assert.Equal(t, "5", value)
	assert.Nil(t, err)

	value, err = findLastSymbol("treb7uchet", intSymbols)
	assert.Equal(t, "7", value)
	assert.Nil(t, err)

	_, err = findLastSymbol("", intSymbols)
	assert.NotNil(t, err)

	_, err = findLastSymbol("something", nil)
	assert.NotNil(t, err)

	_, err = findLastSymbol("something", []string{})
	assert.NotNil(t, err)

	_, err = findLastSymbol("something", []string{"1"})
	assert.NotNil(t, err)
}

func TestParseNumeric(t *testing.T) {

	config := make(map[string]int)

	config["0"] = 0
	config["1"] = 1
	config["2"] = 2
	config["3"] = 3

	doc := newDoc(config)

	value, err := doc.parseLine("1abc2")
	assert.Nil(t, err)
	assert.Equal(t, value, 12)

	_, err = doc.parseLine("4or5")
	assert.NotNil(t, err)
}

func TestLoadNumberFile(t *testing.T) {

	doc := newDoc(numberSymbols)
	lines, err := doc.LoadFile("testdata/numbers.txt")

	assert.Nil(t, err)
	assert.Equal(t, 4, lines)
	assert.Equal(t, 142, doc.Sum())
}

func TestLoadNumberWordFile(t *testing.T) {

	doc := NewDoc(false)
	lines, err := doc.LoadFile("testdata/numbers_and_text.txt")

	fmt.Println(doc.values)

	assert.Nil(t, err)
	assert.Equal(t, 7, lines)
	assert.Equal(t, 281, doc.Sum())
}

func TestLoadBadFile(t *testing.T) {

	doc := NewDoc(true)
	_, err := doc.LoadFile("testdata/numbers.txt")
	assert.Nil(t, err)

	_, err = doc.LoadFile("testdata/numbers.txt")
	// already loaded
	assert.NotNil(t, err)

	// should fail
	_, err = doc.parseLine("")
	assert.NotNil(t, err)

	doc = newDoc(map[string]int{"1": 9223372036854775802})

	// out of range
	v, err := doc.parseLine("1")
	fmt.Println(v)
	assert.NotNil(t, err)

	doc = newDoc(numberSymbols)
	_, err = doc.LoadFile("testdata/bad_data.txt")
	assert.NotNil(t, err)
}
