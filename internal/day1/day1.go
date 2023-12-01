package day1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
)

// calibration document, has a list of lines

var numberSymbols = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
	"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

var numberWordSymbols = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
	"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

type CalibrationDocument struct {
	values    []int
	symbolMap map[string]int
	symbols   []string
}

func NewDoc(onlyDigits bool) CalibrationDocument {
	if onlyDigits {
		return newDoc(numberSymbols)
	} else {
		return newDoc(numberWordSymbols)
	}
}

func newDoc(symbols map[string]int) CalibrationDocument {
	keys := make([]string, 0, len(symbols))
	for k := range symbols {
		keys = append(keys, k)
	}
	return CalibrationDocument{
		[]int{}, symbols, keys,
	}
}

func (doc *CalibrationDocument) LoadFile(name string) (int, error) {

	if len(doc.values) != 0 {
		return 0, errors.New("Document already initialized!")
	}
	values, err := util.ReadFileLines(name, func(v string) (int, error) {
		result, err := doc.parseLine(v)
		if err != nil {
			return 0, err
		}
		return result, nil
	})
	if err != nil {
		return 0, err
	}
	doc.values = values
	return len(doc.values), nil
}

// parse a line, take the first and last found symbol, combine the values, and return
func (doc *CalibrationDocument) parseLine(line string) (int, error) {

	if line == "" {
		return 0, errors.New("no input provided")
	}

	firstSymbol, err := findFirstSymbol(line, doc.symbols)
	if err != nil {
		return 0, err
	}
	lastSymbol, _ := findLastSymbol(line, doc.symbols)
	// unreachable as if you can find a symbol above,
	// then you won't find an error here.
	// if err != nil {
	// 	return 0, err
	// }

	result := fmt.Sprintf("%d%d", doc.symbolMap[firstSymbol], doc.symbolMap[lastSymbol])

	v, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func findFirstSymbol(value string, symbols []string) (string, error) {

	if value == "" {
		return "", errors.New("no input provided")
	}
	if len(symbols) == 0 {
		return "", errors.New("no symbols provided")
	}

	find := func(input string) (string, bool) {
		for _, symbol := range symbols {
			if strings.HasPrefix(input, symbol) {
				return symbol, true
			}
		}
		return "", false
	}

	index := 0
	for index < len(value) {
		if result, ok := find(value[index:]); ok {
			return result, nil
		}
		index += 1
	}
	return "", errors.New("not found!")
}

func findLastSymbol(value string, symbols []string) (string, error) {

	if value == "" {
		return "", errors.New("no input provided")
	}
	if len(symbols) == 0 {
		return "", errors.New("no symbols provided")
	}

	find := func(input string) (string, bool) {
		for _, symbol := range symbols {
			if strings.HasPrefix(input, symbol) {
				return symbol, true
			}
		}
		return "", false
	}

	index := len(value) - 1
	for index >= 0 {
		if result, ok := find(value[index:]); ok {
			return result, nil
		}
		index -= 1
	}
	return "", errors.New("not found!")
}

func (doc *CalibrationDocument) Sum() int {
	total := 0
	for _, value := range doc.values {
		total += value
	}
	return total
}
