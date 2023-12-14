package day14

import (
	"log"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
	"github.com/samber/lo"
)

type DishPlatform struct {
	field [][]string
}

func (dish *DishPlatform) Load() int {

	return lo.Sum(lo.Map(dish.field, func(row []string, index int) int {
		rowLoad := len(dish.field) - index
		rocks := lo.Count(row, round_rock)
		return rowLoad * rocks
	}))
}

func (dish *DishPlatform) ToString() string {

	var b strings.Builder
	b.Grow(len(dish.field)*len(dish.field[0]) + 1)
	for ix, row := range dish.field {
		if ix != 0 {
			b.WriteString("\n")
		}
		for _, v := range row {
			b.WriteString(v)
		}
	}
	return b.String()
}

func DishFromRows(rows []string) *DishPlatform {
	data := lo.Map(rows, func(item string, index int) []string {
		out := make([]string, len(item))
		for ix := range item {
			out = append(out, item[ix:ix+1])
		}
		return out
	})
	return &DishPlatform{data}
}

type Axis = uint8

// http://golang.org/ref/spec#Iota
const (
	x_axis Axis = 1 << iota
	y_axis Axis = 1 << iota
)

const (
	round_rock  string = "O"
	square_rock string = "#"
	empty       string = "."
)

type Direction = lo.Tuple2[Axis, int]

var North Direction = lo.T2(x_axis, -1)
var South Direction = lo.T2(x_axis, 1)
var East Direction = lo.T2(y_axis, 1)
var West Direction = lo.T2(y_axis, -1)

func (plat *DishPlatform) Spin() {
	plat.Tilt(North)
	plat.Tilt(West)
	plat.Tilt(South)
	plat.Tilt(East)
}

func (plat *DishPlatform) InRange(x int, y int) bool {
	if x >= 0 && x < len(plat.field) {
		if y >= 0 && y < len(plat.field[0]) {
			return true
		}
	}
	return false
}

func (plat *DishPlatform) Tilt(dir Direction) {

	if dir.A == x_axis {
		move_x := dir.B
		move := func(ix int) {
			for iy := 0; iy < len(plat.field[0]); iy++ {
				pos_x, v := ix, plat.field[ix][iy]
				if v != round_rock {
					continue
				}
				for plat.InRange(pos_x+move_x, iy) && plat.field[pos_x+move_x][iy] == empty {
					pos_x += move_x
				}
				// swap the empty and round rocks
				if pos_x != ix {
					plat.field[pos_x][iy], plat.field[ix][iy] = plat.field[ix][iy], plat.field[pos_x][iy]
				}
			}

		}
		if dir.B < 0 {
			// move rocks > 0 -> 0
			for ix := 1; ix < len(plat.field); ix++ {
				move(ix)
			}
		} else {
			// move rocks <= 0 away from 0
			for ix := len(plat.field) - 1; ix >= 0; ix-- {
				move(ix)
			}
		}
	} else if dir.A == y_axis {
		move_y := dir.B
		move := func(iy int) {
			for ix := 0; ix < len(plat.field); ix++ {
				pos_y, v := iy, plat.field[ix][iy]
				if v != round_rock {
					continue
				}
				for plat.InRange(ix, pos_y+move_y) && plat.field[ix][pos_y+move_y] == empty {
					pos_y += move_y
				}
				if pos_y != iy {
					plat.field[ix][pos_y], plat.field[ix][iy] = plat.field[ix][iy], plat.field[ix][pos_y]
				}
			}
		}
		if dir.B < 0 {
			for iy := 0; iy < len(plat.field[0]); iy++ {
				move(iy)
			}
		} else {
			for iy := len(plat.field[0]) - 1; iy >= 0; iy-- {
				move(iy)
			}
		}
	}
}

func LoadPlatform(inputFile string) (*DishPlatform, error) {
	rows, err := util.ReadFileAsStrings(inputFile)
	if err != nil {
		return nil, err
	}

	return DishFromRows(rows), nil
}

type LoadPredictor struct {
	// cache   map[string][][]string
	dish    *DishPlatform
	results []string
}

func (dish *DishPlatform) Predictor() *LoadPredictor {
	return &LoadPredictor{dish, nil}
}

// / predict what the load would be after x spins
func (p *LoadPredictor) PredictLoad(spins int) int {

	// this should copy the dish
	// so that this is repeatable

	for i := 0; i < spins; i++ {

		dishStr := p.dish.ToString()

		_, index, found := lo.FindIndexOf(p.results, func(item string) bool {
			return item == dishStr
		})
		// log.Printf("i=%sindex: %d, found: %v\n", index, found)
		if found {
			// we don't care about the other results

			log.Printf("found loop at %d\n", index)
			p.results = p.results[index:]
			loop_size := len(p.results)
			total := spins
			remaining := total - i
			extra := remaining % loop_size

			// spin extra amount of times
			for e := 0; e < extra; e++ {
				p.dish.Spin()
			}

			return p.dish.Load()
		}
		p.results = append(p.results, dishStr)
		p.dish.Spin()
	}
	// if we have not determined a pattern, then
	// return the load after the proper amount of spins
	return p.dish.Load()
}
