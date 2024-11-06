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

func getSumMul(arr []int) int {
	module := 1000000007
	sum1, sum2, sum3 := 0, 0, 0

	for _, a := range arr {
		sum3 += (a % module * sum2 % module) % module
		sum3 %= module
		sum2 += (a % module * sum1 % module) % module
		sum2 %= module
		sum1 += a % module
		sum1 %= module
	}

	return sum3
}

func main() {
	nums := readData()
	fmt.Println(getSumMul(nums))
}
