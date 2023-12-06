package day6

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/miketzian/adventofcode2023/internal/util"
	"github.com/samber/lo"
)

type Record struct {
	time     int
	distance int
}

func (rec *Record) Compete() int {
	var low int
	var high int
	for hold := 1; hold < rec.time; hold++ {
		distance := hold * (rec.time - hold)
		if distance > rec.distance {
			low = hold
			break
		}
	}
	for hold := rec.time - 1; hold > 1; hold-- {
		distance := hold * (rec.time - hold)
		if distance > rec.distance {
			high = hold
			break
		}
	}

	ret := high - low + 1
	log.Printf("%v: low=%d, high=%d, ret=%d", rec, low, high, ret)
	// there is a curve, so anything between the high and low
	// winning distances is also a win
	return high - low + 1
}

func LoadFile(name string, combineNumbers bool) ([]*Record, error) {

	input, err := util.ReadFileAsStrings(name)
	if err != nil {
		return nil, err
	}
	if len(input) != 2 {
		return nil, errors.New("expected two lines")
	}

	times_distances, err := parseRecords(input[0], input[1])
	if err != nil {
		return nil, err
	}

	if combineNumbers {
		res := lo.Reduce(times_distances, func(agg []string, rec *Record, _ int) []string {
			agg[0] = fmt.Sprintf("%s%d", agg[0], rec.time)
			agg[1] = fmt.Sprintf("%s%d", agg[1], rec.distance)
			return agg
		}, make([]string, 2))

		time, _ := strconv.ParseInt(res[0], 10, 64)
		distance, _ := strconv.ParseInt(res[1], 10, 64)

		return []*Record{{int(time), int(distance)}}, nil
	}
	return times_distances, nil

}

func parseRecords(time string, distance string) ([]*Record, error) {

	readInts := func(input string) []int {
		var s scanner.Scanner
		out := make([]int, 0)
		s.Init(strings.NewReader(input))
		for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
			if tok == scanner.Int {
				v, _ := strconv.ParseInt(s.TokenText(), 10, 32)
				out = append(out, int(v))
			}
		}
		return out
	}

	times := readInts(time)
	distances := readInts(distance)

	if len(times) != len(distances) {
		return nil, errors.New(fmt.Sprintf("times[%v] != distances[%v]", times, distances))
	}
	return lo.Map(lo.Zip2(times, distances), func(item lo.Tuple2[int, int], ix int) *Record {
		return &Record{item.A, item.B}
	}), nil
}

func RecordProduct(records []*Record) int {
	return lo.Reduce(records, func(agg int, r *Record, _ int) int {
		if agg == 0 {
			return r.Compete()
		}
		return agg * r.Compete()
	}, 0)
}
