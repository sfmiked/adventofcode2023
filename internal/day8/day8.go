package day8

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
	"github.com/samber/lo"
)

type Guidebook struct {
	moves   []rune
	options map[string]lo.Tuple2[string, string]
}

// play the game, return the number of turns until
// check returns true
func (book *Guidebook) Play(start string, check func(end string) bool) (int, error) {

	log.Printf("[%d] Start %s", 0, start)
	loc := start
	max_moves := len(book.moves) - 1
	count := 0
	for i := 0; ; i++ {
		opts, ok := book.options[loc]
		if !ok {
			return -1, errors.New(fmt.Sprintf("Not found: %s", loc))
		}
		count++
		switch book.moves[i] {
		case 'L':
			loc = opts.A
			break
		case 'R':
			loc = opts.B
			break
		}
		// log.Printf("[%d] Move %s", count, loc)
		if check(loc) {
			// we're done!
			return count, nil
		}
		if i == max_moves {
			i = -1
		}
	}
}

func (book *Guidebook) PlayEnding(end string) (int, error) {
	match := func(input string) bool {
		return strings.HasSuffix(input, end)
	}
	lcm := 1
	for k := range book.options {
		if strings.HasSuffix(k, "A") {
			result, err := book.Play(k, match)
			if err != nil {
				return -1, err
			}
			lcm = lcm * result / util.GCD(lcm, result)
		}
	}
	return lcm, nil
}

func LoadFile(input string) (*Guidebook, error) {

	var book *Guidebook
	err := util.ReadFileForEachLine(input, func(line string) error {
		if book == nil {
			book = &Guidebook{[]rune(line), make(map[string]lo.Tuple2[string, string], 0)}
		} else if line != "" {
			if len(line) != 16 {
				return errors.New("Unexpected Line format, should be 'AAA = (BBB, CCC)'")
			}
			code := line[0:3]
			left_code := line[7:10]
			right_code := line[12:15]
			log.Printf("%s -> (%s, %s)", code, left_code, right_code)
			book.options[code] = lo.T2(left_code, right_code)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return book, nil
}
