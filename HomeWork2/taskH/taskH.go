package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData() []int {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, _ := strconv.Atoi(nStr)

	arrayStr, _ := reader.ReadString('\n')
	arrayStr = strings.TrimSpace(arrayStr)
	arrayStrElements := strings.Split(arrayStr, " ")

	nums := make([]int, n)
	for i, part := range arrayStrElements {
		nums[i], _ = strconv.Atoi(part)
	}

	return nums
}

func minTotalMoves(employees []int) int {
	n := len(employees)

	totalEmployees := 0
	for _, num := range employees {
		totalEmployees += num
	}

	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + employees[i-1]
	}

	halfTotal := (totalEmployees + 1) / 2
	p := 0
	for i := 1; i <= n; i++ {
		if prefixSum[i] >= halfTotal {
			p = i
			break
		}
	}

	totalMoves := 0
	for i := 1; i <= n; i++ {
		totalMoves += employees[i-1] * abs(i-p)
	}

	return totalMoves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	cabs := readData()
	fmt.Println(minTotalMoves(cabs))
}
