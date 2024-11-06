package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData() ([]int, int, []int) {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, _ := strconv.Atoi(nStr)

	arrayStr, _ := reader.ReadString('\n')
	arrayStr = strings.TrimSpace(arrayStr)
	arrayStrElements := strings.Split(arrayStr, " ")

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(arrayStrElements[i])
	}

	mkStr, _ := reader.ReadString('\n')
	mkStr = strings.TrimSpace(mkStr)
	mkStrElements := strings.Split(mkStr, " ")
	m, _ := strconv.Atoi(mkStrElements[0])
	k, _ := strconv.Atoi(mkStrElements[1])

	arrayStr, _ = reader.ReadString('\n')
	arrayStr = strings.TrimSpace(arrayStr)
	arrayStrElements = strings.Split(arrayStr, " ")

	x := make([]int, m)
	for i, part := range arrayStrElements {
		x[i], _ = strconv.Atoi(part)
	}

	return a, k, x
}

func findFinalPositions(a []int, k int, x []int) []int {
	l, r := 0, 1
	initialK := k
	finalPositions := make([]int, len(a))
	finalPositions[0] = 1

	for r < len(a) {
		if a[r] > a[r-1] {
			finalPositions[r] = l + 1
			r++
		} else if a[r] < a[r-1] {
			l = r
			r++
			finalPositions[l] = l + 1
			k = initialK
		} else {
			for k < 1 {
				if l != len(a) {
					if a[l] == a[l+1] {
						l++
						k++
					} else {
						l++
					}
				}
			}
			k--
			finalPositions[r] = l + 1
			r++
		}
	}

	result := make([]int, len(x))

	for i, el := range x {
		result[i] = finalPositions[el-1]
	}

	return result
}

func printArray(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
}

func main() {
	a, k, x := readData()
	result := findFinalPositions(a, k, x)

	printArray(result)
}
