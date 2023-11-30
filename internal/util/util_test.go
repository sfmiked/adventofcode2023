package util

import (
	"errors"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileString(t *testing.T) {
	v, err := ReadFileAsStrings("testdata/file_lines.txt")
	assert.Nil(t, err)

	assert.Equal(t, len(v), 3)
	assert.Equal(t, v[0], "1000")
	assert.Equal(t, v[1], "2000")
	assert.Equal(t, v[2], "3000")
}

func TestFileInt(t *testing.T) {
	v, err := ReadFileAsInt("testdata/file_lines.txt")
	assert.Nil(t, err)

	assert.Equal(t, len(v), 3)
	assert.Equal(t, v[0], 1000)
	assert.Equal(t, v[1], 2000)
	assert.Equal(t, v[2], 3000)
}

func TestFileInt64(t *testing.T) {
	v, err := ReadFileAsInt64("testdata/file_lines_64.txt")
	assert.Nil(t, err)

	assert.Equal(t, len(v), 3)
	assert.Equal(t, v[0], int64(9223372036854775801))
	assert.Equal(t, v[1], int64(9223372036854775802))
	assert.Equal(t, v[2], int64(9223372036854775803))
}

func TestFileDoesNotExist(t *testing.T) {
	_, err := ReadFileAsStrings("testdata/file_not_existing.txt")
	assert.NotNil(t, err)
}

func TestRevereString(t *testing.T) {

	input, expected := "salamander", "rednamalas"
	result := ReverseString(input)
	assert.Equal(t, expected, result)
}

func TestRevereStringArray(t *testing.T) {

	input := []string{"abc", "def", "hij"}
	expected := []string{"hij", "def", "abc"}

	// in-place replace
	ReverseStringArray(input)
	assert.Equal(t, expected, input)
}

func TestFileNonInt(t *testing.T) {

	// should not parse as int
	_, err := ReadFileAsInt("testdata/file_nonint.txt")
	assert.NotNil(t, err)

	// should not parse as int64
	_, err = ReadFileAsInt64("testdata/file_nonint.txt")
	assert.NotNil(t, err)

	// should read ok as string
	_, err = ReadFileAsStrings("testdata/file_nonint.txt")
	assert.Nil(t, err)
}

// coverage for scanner.Err()
type flakyReader struct {
	inner  io.Reader
	broken *struct {
		value bool
	}
}

func (r flakyReader) Read(b []byte) (n int, err error) {
	log.Printf("Called Read: %d, broken? %v", len(b), r.broken)
	if r.broken != nil {
		return -1, errors.New("oh no!")
	}
	r.broken.value = true
	return r.inner.Read(b)
}

func TestFlakyReader(t *testing.T) {

	file, err := os.Open("testdata/file_vlong.txt")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	holder := struct{ value bool }{value: false}
	data := flakyReader{file, &holder}

	parser := func(_ string) (string, error) {
		// a function that will parse a string once
		//
		// we don't care about the data
		return "", nil
	}

	// read until broken = true

	_, err = readToArray(data, parser)
	log.Printf("error: %v", err)
	assert.NotNil(t, err)
}

func TestCloseFile(t *testing.T) {
	file, err := os.Open("testdata/file_lines.txt")
	if err != nil {
		t.Error(err)
	}
	// first time should work
	err = maybeCloseFile(file)
	assert.Nil(t, err)

	// should already be closed
	err = maybeCloseFile(file)
	assert.NotNil(t, err)
}
