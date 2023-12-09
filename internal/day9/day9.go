package day9

import (
	"log"
	"strconv"
	"strings"

	"github.com/miketzian/adventofcode2023/internal/util"
	"github.com/samber/lo"
)

func extrapolate(arrs [][]int, left bool) int {
	count := 0
	for aix := len(arrs) - 1; aix > -1; aix-- {
		if count == 0 {
			if left {
				// simplest
				arrs[aix] = append([]int{0}, arrs[aix]...)

				// copy destination source  (less memory?)
				// arrs[aix] = append(arrs[aix], 0)
				// copy(arrs[aix][1:], arrs[aix][:len(arrs[aix])-2])
				// arrs[aix][0] = 0
			} else {
				// append to the right
				arrs[aix] = append(arrs[aix], 0)
			}
		} else {

			thisArr := arrs[aix]
			thatArr := arrs[aix+1]

			if left {
				diff := thatArr[0] * -1
				last := thisArr[0]
				arrs[aix] = append([]int{last + diff}, thisArr...)
			} else {
				thisLast := len(thisArr) - 1
				thatLast := len(thatArr) - 1

				diff := thatArr[thatLast]
				last := thisArr[thisLast]

				arrs[aix] = append(arrs[aix], last+diff)
			}
		}
		count += 1
	}
	// debug
	/*
		spaces := ""
		for _, arr := range arrs {
			log.Printf("%s%v", spaces, arr)
			spaces += " "
		}
	*/

	if left {
		// the first on the left
		return arrs[0][0]
	} else {
		// the one on the right
		return arrs[0][len(arrs[0])-1]
	}

}

func Play(game []int, addToLeft bool) int {

	diff := func(arr []int) []int {
		d1 := make([]int, len(arr)-1)
		for i := 1; i < len(arr); i++ {
			// ie d[0] = arr[1] - arr[0]
			d1[i-1] = arr[i] - arr[i-1]
		}
		return d1
	}

	allZeros := func(line []int) bool {
		return lo.Count(line, 0) == len(line)
	}

	check := make([][]int, 1)
	check[0] = game

	next := game

	for !allZeros(next) {
		// log.Printf("Not all zeros: %v\n", next)
		next = diff(next)
		check = append(check, next)
	}

	return extrapolate(check, addToLeft)
}

func PlayFile(inputFile string, addToLeft bool) (int, error) {

	result := 0
	err := util.ReadFileForEachLine(inputFile, func(line string) error {

		num_str := strings.Split(line, " ")
		nums := make([]int, len(num_str))
		for i, value := range num_str {
			result, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				log.Printf("Error Parsing [%s] to int: %v", value, err)
				return err
			}
			nums[i] = int(result)
		}
		res := Play(nums, addToLeft)
		log.Printf("Input: %v, Result: %d\n", nums, res)
		result += res
		return nil
	})
	if err != nil {
		return -1, err
	}
	return result, nil
}
