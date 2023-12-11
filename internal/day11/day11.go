package day11

import (
	"errors"
	"fmt"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
	"github.com/samber/lo"
)

type Observations struct {
	grid        [][]string
	galaxies    map[int]lo.Tuple2[int, int]
	expandValue int
}

func (obs *Observations) Galaxies() []int {
	return lo.Keys(obs.galaxies)
}

func NewObservations(inputFile string, expandValue int) (*Observations, error) {

	// load the grid from the input file
	grid, err := util.ReadFileLines(inputFile, func(item string) ([]string, error) {
		out := make([]string, len(item))
		for ix := range item {
			out[ix] = item[ix : ix+1]
		}
		return out, nil
	})
	if err != nil {
		return nil, err
	}

	// replace expanding regions with e
	markExpandingRegions(grid)

	// search the grid for #, replacing with galaxy number
	galaxies := mapGalaxies(grid)

	obs := &Observations{
		grid, galaxies, expandValue,
	}

	return obs, nil
}

func (obs *Observations) GalaxyDistance(g1 int, g2 int) (int, error) {

	loc1, ok := obs.galaxies[g1]
	if !ok {
		return 0, errors.New("not found")
	}
	loc2, ok := obs.galaxies[g2]
	if !ok {
		return -1, errors.New("not found")
	}

	dist := 0

	xby := 1
	if loc1.A > loc2.A {
		xby = -1
	}
	x := loc1.A
	for x != loc2.A {
		if obs.grid[x][loc1.B] == "e" {
			dist += obs.expandValue
		} else {
			dist += 1
		}
		x += xby
	}

	yby := 1
	if loc1.B > loc2.B {
		yby = -1
	}
	y := loc1.B
	for y != loc2.B {
		if obs.grid[loc2.A][y] == "e" {
			dist += obs.expandValue
		} else {
			dist += 1
		}
		y += yby
	}
	return dist, nil
}

func gridFromString(input string) [][]string {
	return lo.Map(strings.Split(input, "\n"), func(item string, ix int) []string {
		out := make([]string, len(item))
		for ix := range item {
			out[ix] = item[ix : ix+1]
		}
		return out
	})
}

func markExpandingRegions(input [][]string) {
	for _, row := range input {
		if lo.EveryBy(row, func(s string) bool { return s == "." }) {
			lo.ForEach(row, func(_ string, i int) {
				// for i := range row {
				row[i] = "e"
			})
		}
	}
	for ix := range input[0] {
		if lo.EveryBy(input, func(row []string) bool { return row[ix] == "." || row[ix] == "e" }) {
			lo.ForEach(input, func(row []string, _ int) {
				row[ix] = "e"
			})
		}
	}
}

func mapGalaxies(input [][]string) map[int]lo.Tuple2[int, int] {

	galaxies := make(map[int]lo.Tuple2[int, int], 0)
	index := 1
	for ix, row := range input {
		for iy, v := range row {
			if v == "#" {
				row[iy] = fmt.Sprintf("%d", index)
				galaxies[index] = lo.T2(ix, iy)
				index++
			}

		}
	}

	return galaxies
}
