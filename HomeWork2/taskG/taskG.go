package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData() (int, []string) {
	reader := bufio.NewReader(os.Stdin)

	nsStr, _ := reader.ReadString('\n')
	nsStr = strings.TrimSpace(nsStr)
	nsStrElements := strings.Split(nsStr, " ")
	n, _ := strconv.Atoi(nsStrElements[0])
	k, _ := strconv.Atoi(nsStrElements[1])

	arrayStr, _ := reader.ReadString('\n')
	arrayStr = strings.TrimSpace(arrayStr)
	arrayStrElements := strings.Split(arrayStr, "")

	nums := make([]string, n)
	for i, part := range arrayStrElements {
		nums[i] = part
	}

	return k, nums
}

func maxLiteraryWorkLength(c int, s []string) int {
	n := len(s)
	prefixA := make([]int, n+1)
	prefixB := make([]int, n+1)
	coarseness := make([]int, n+1)

	for i := 1; i <= n; i++ {
		prefixA[i] = prefixA[i-1]
		prefixB[i] = prefixB[i-1]
		coarseness[i] = coarseness[i-1]
		if s[i-1] == "a" {
			prefixA[i]++
		} else if s[i-1] == "b" {
			prefixB[i]++
			coarseness[i] += prefixA[i-1]
		}
	}

	maxLength := 0
	L := 1

	getCoarseness := func(L, R int) int {
		totalCoarseness := coarseness[R] - coarseness[L-1]
		totalCoarseness -= prefixA[L-1] * (prefixB[R] - prefixB[L-1])
		return totalCoarseness
	}

	for R := 1; R <= n; R++ {
		for L <= R && getCoarseness(L, R) > c {
			L++
		}
		if R-L+1 > maxLength {
			maxLength = R - L + 1
		}
	}

	return maxLength
}

func main() {
	c, s := readData()
	fmt.Println(maxLiteraryWorkLength(c, s))
}
