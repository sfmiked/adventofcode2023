package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompete(t *testing.T) {

	cases := []struct {
		record   *Record
		expected int
	}{
		{&Record{7, 9}, 4},
		{&Record{15, 40}, 8},
		{&Record{30, 200}, 9},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, c.record.Compete())
	}
}

func TestParse(t *testing.T) {

	records, err := parseRecords("Time:     7", "Distance:    9")
	assert.Nil(t, err)
	assert.Equal(t, &Record{7, 9}, records[0])

}

func TestLoad(t *testing.T) {

	_, err := LoadFile("testdata/non_existing.txt", false)
	assert.NotNil(t, err)

	_, err = LoadFile("testdata/bad_file.txt", false)
	assert.NotNil(t, err)

	_, err = LoadFile("testdata/only_time.txt", false)
	assert.NotNil(t, err)

	records, err := LoadFile("testdata/input.txt", false)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(records))
	assert.Equal(t, &Record{7, 9}, records[0])
	assert.Equal(t, &Record{15, 40}, records[1])
	assert.Equal(t, &Record{30, 200}, records[2])

	result := RecordProduct(records)

	assert.Equal(t, 288, result)

	records, err = LoadFile("testdata/input.txt", true)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(records))
	assert.Equal(t, &Record{71530, 940200}, records[0])

}
