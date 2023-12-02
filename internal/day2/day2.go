package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
)

// / Represents one game, which may have many plays
type Game struct {
	id    int
	plays []Play
}

func LoadGameFile(name string) ([]*Game, error) {
	return util.ReadFileLines(name, func(input string) (*Game, error) {
		game, err := ParseGame(input)
		if err != nil {
			return nil, err
		}
		return game, nil
	})
}

// / Represents one play
type Play struct {
	red   int
	blue  int
	green int
}

func (game *Game) Id() int {
	return game.id
}

func (game *Game) DicesRequired() (int, int, int) {
	min := game.MinPlayOfPlays()
	return min.red, min.green, min.blue
}

func (game *Game) Power() int {
	min := game.MinPlayOfPlays()
	return min.red * min.blue * min.green
}

// find the minimal set of color dices required for
// all plays in the game.
func (game *Game) MinPlayOfPlays() *Play {

	minPlay := Play{0, 0, 0}

	for _, play := range game.plays {
		if play.red > minPlay.red {
			minPlay.red = play.red
		}
		if play.blue > minPlay.blue {
			minPlay.blue = play.blue
		}
		if play.green > minPlay.green {
			minPlay.green = play.green
		}
	}
	return &minPlay
}

func ParseGame(line string) (*Game, error) {

	elements := strings.Split(line, " ")
	if len(elements) < 3 || len(elements)%2 != 0 {
		return nil, errors.New(fmt.Sprintf("Invalid Line: %s", line))
	}

	game_id, err := strconv.ParseInt(elements[1][:len(elements[1])-1], 10, 32)
	if err != nil {
		return nil, err
	}
	// fmt.Println(line)

	game := Game{
		id:    int(game_id),
		plays: make([]Play, 0),
	}

	play := Play{0, 0, 0}
	for ix, unit := range elements[2:] {

		if ix%2 != 0 {
			// lets only act on every second element
			continue
		}

		value, err := strconv.ParseInt(unit, 10, 32)
		if err != nil {
			return nil, err
		}

		// next element
		color := elements[ix+3]
		// fmt.Printf("%d %s \n", value, color)
		if strings.HasPrefix(color, "red") {
			play.red = int(value)
		}
		if strings.HasPrefix(color, "blue") {
			play.blue = int(value)
		}
		if strings.HasPrefix(color, "green") {
			play.green = int(value)
		}
		if strings.HasSuffix(color, ";") {
			// log.Printf("%v\n", play)
			game.plays = append(game.plays, play)
			play = Play{0, 0, 0}
		}
	}
	// log.Printf("%v\n", play)
	game.plays = append(game.plays, play)
	return &game, nil
}
