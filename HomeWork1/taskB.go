package main

import (
	"fmt"
	"sort"
)

func minStepsForSet(a, b, c, d int) (int, int) {
	if a == 0 && c == 0 || b == 0 && d == 0 {
		return 1, 1
	}

	if a == 0 {
		return 1, c + 1
	} else if b == 0 {
		return 1, d + 1
	} else if c == 0 {
		return a + 1, 1
	} else if d == 0 {
		return b + 1, 1
	}

	if a == b {
		if c == d {
			if a <= c {
				return a + 1, 1
			} else {
				return 1, c + 1
			}
		} else {
			return a + 1, 1
		}
	} else if c == d {
		return 1, c + 1
	}

	variants := make(map[int][2]int)
	maxShirts := max(a, b)
	maxSocks := max(c, d)

	variants[maxShirts+1+1] = [2]int{maxShirts + 1, 1}
	variants[1+maxSocks+1] = [2]int{1, maxSocks + 1}
	variants[a+1+c+1] = [2]int{a + 1, c + 1}
	variants[b+1+d+1] = [2]int{b + 1, d + 1}

	keys := make([]int, 0, len(variants))
	for k := range variants {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return variants[keys[0]][0], variants[keys[0]][1]
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(minStepsForSet(a, b, c, d))
}
