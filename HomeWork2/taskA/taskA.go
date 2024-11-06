package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculatePrefixSums(nums []int) []int {
	prefixSums := make([]int, len(nums))
	prefixSums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i]
	}
	return prefixSums
}

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

func printArray(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
}

func main() {
	nums := readData()
	nums = calculatePrefixSums(nums)
	printArray(nums)
}
