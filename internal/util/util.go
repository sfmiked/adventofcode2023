package util

// https://github.com/miketzian/adventofcode2020/blob/main/util.go

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func ReverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}
	// return the reversed string.
	return string(rns)
}
func ReverseStringArray(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func ReadFileAsInt(name string) ([]int, error) {
	return ReadFileLines(name, func(value string) (int, error) {
		result, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return 0, err
		}
		// result is actually int64 type
		return int(result), nil
	})
}

func ReadFileAsInt64(name string) ([]int64, error) {
	return ReadFileLines(name, func(v string) (int64, error) {
		result, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return result, nil
	})
}

func ReadFileAsStrings(name string) ([]string, error) {
	return ReadFileLines(name, func(v string) (string, error) { return v, nil })
}

func ReadFileLines[V comparable](name string, parser func(string) (V, error)) ([]V, error) {

	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer maybeCloseFile(file)
	return readToArray(file, parser)
}

func maybeCloseFile(file *os.File) error {
	if err := file.Close(); err != nil {
		log.Printf("failed to close file %v", err)
		return err
	}
	return nil
}

func readToArray[V comparable](input io.Reader, parser func(string) (V, error)) ([]V, error) {
	scanner := bufio.NewScanner(input)

	output := make([]V, 0)
	for scanner.Scan() {
		result, err := parser(scanner.Text())
		if err != nil {
			return nil, err
		}
		output = append(output, result)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	// log.Printf("Output: %v", output)
	return output, nil
}
