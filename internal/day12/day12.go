package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
)

func parseInput(line string, expand bool) (string, []int, error) {
	first := strings.Split(line, " ")

	if expand {
		// each repeat is split by ?
		first[0] = strings.Repeat(fmt.Sprintf("%s?", first[0]), 5)
		// drop the final ?
		first[0] = first[0][:len(first[0])-1]
		// these are comma split, so add a comma after each group
		// 1,2,3 -> 1,2,3,
		first[1] = strings.Repeat(fmt.Sprintf("%s,", first[1]), 5)
		// remove the final duplicated comma
		// 1,2,3,1,2,3, -> 1,2,3,1,2,3
		first[1] = first[1][:len(first[1])-1]
	}

	second := strings.Split(first[1], ",")
	groups := make([]int, len(second))
	for ix, v := range second {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return "", nil, err
		}
		groups[ix] = int(val)
	}
	return first[0], groups, nil
}

func CheckArrangementsFromFile(inputFile string, expand bool) (int, error) {

	result := 0

	err := util.ReadFileForEachLine(inputFile, func(line string) error {

		spring, groups, err := parseInput(line, expand)
		if err != nil {
			return err
		}
		result += ValidArrangements(spring, groups)

		return nil
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func ValidArrangements(spring string, groups []int) int {

	// need to declare so that we can recurse
	var arrangements func(spring string, groups []int, size int, debug string) int
	// inner function

	cache := make(map[string]int, 0)

	arrangements = func(spring string, groups []int, size int, debug string) int {

		if len(spring) == 0 {
			// we have consumed all the spring
			// and we are able to use it for the last group
			if len(groups) == 1 && groups[0] == size {
				// if debug
				// fmt.Printf("%s%s\n", debug, strings.Repeat("#", size))
				return 1

				// there is no more spring, but also
				// no more groups, so this is also fine
			} else if len(groups) == 0 && size == 0 {
				// if debug
				// fmt.Println(debug)
				return 1

				// no more spring, but there is groups
				// this is not a match
			} else {
				return 0
			}
		}

		// check if we are in the cache
		if cached, ok := cache[fmt.Sprintf("%s^^%v^^%d", spring, groups, size)]; ok {
			return cached
		}

		switch spring[0:1] {
		case "?":
			r := arrangements(fmt.Sprintf(".%s", spring[1:]), groups, size, debug)
			r += arrangements(fmt.Sprintf("#%s", spring[1:]), groups, size, debug)
			cache[fmt.Sprintf("%s^^%v^^%d", spring, groups, size)] = r
			return r
		case "#":
			if len(groups) > 0 {
				// this block is too large for this group, so this path is invalid.
				if size > groups[0] {
					return 0
				}
			}
			// increase size
			r := arrangements(spring[1:], groups, size+1, debug)
			cache[fmt.Sprintf("%s^^%v^^%d", spring, groups, size)] = r
			return r
		// case ".":
		default:
			// if size is zero, then we are not in a block, so we can just move forward.
			if size == 0 {
				r := arrangements(spring[1:], groups, 0, fmt.Sprintf("%s.", debug))
				cache[fmt.Sprintf("%s^^%v^^%d", spring, groups, size)] = r
				return r
			}
			// we are a the end of a previous # block since size > 0
			// so we need to see if the next group fits within that block
			if len(groups) > 0 {
				if size == groups[0] {
					r := arrangements(spring[1:], groups[1:], 0, fmt.Sprintf("%s%s.", debug, strings.Repeat("#", size)))
					cache[fmt.Sprintf("%s^^%v^^%d", spring, groups, size)] = r
					return r
				}
			}
			// size doesn't match the block, so this path is invalid
			cache[fmt.Sprintf("%s^^%v^^%d", spring, groups, size)] = 0
			return 0
		}
	}
	return arrangements(spring, groups, 0, "")
}
