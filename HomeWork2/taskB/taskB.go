package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData() (int, []int) {
	reader := bufio.NewReader(os.Stdin)

	nkStr, _ := reader.ReadString('\n')
	nkStr = strings.TrimSpace(nkStr)
	nkStrElements := strings.Split(nkStr, " ")
	n, _ := strconv.Atoi(nkStrElements[0])
	k, _ := strconv.Atoi(nkStrElements[1])

	arrayStr, _ := reader.ReadString('\n')
	arrayStr = strings.TrimSpace(arrayStr)
	arrayStrElements := strings.Split(arrayStr, " ")

	nums := make([]int, n)
	for i, part := range arrayStrElements {
		nums[i], _ = strconv.Atoi(part)
	}

	return k, nums
}

func countSets(k int, nums []int) int {
	count, currentSum, r := 0, 0, 0

	for l := range nums {
		for r < len(nums) && currentSum < k {
			currentSum += nums[r]
			r += 1
		}
		if currentSum == k {
			count += 1
		}
		currentSum -= nums[l]
	}

	return count
}

func main() {
	k, nums := readData()
	fmt.Println(countSets(k, nums))
}
