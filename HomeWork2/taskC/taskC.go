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

	nrStr, _ := reader.ReadString('\n')
	nrStr = strings.TrimSpace(nrStr)
	nrStrElements := strings.Split(nrStr, " ")
	n, _ := strconv.Atoi(nrStrElements[0])
	r, _ := strconv.Atoi(nrStrElements[1])

	arrayStr, _ := reader.ReadString('\n')
	arrayStr = strings.TrimSpace(arrayStr)
	arrayStrElements := strings.Split(arrayStr, " ")

	nums := make([]int, n)
	for i, part := range arrayStrElements {
		nums[i], _ = strconv.Atoi(part)
	}

	return r, nums
}

func countVariants(number int, nums []int) int {
	count, r := 0, 0
	for l := range nums {
		for r < len(nums) && nums[r]-nums[l] <= number {
			r += 1
		}
		count += len(nums) - r
	}
	return count
}

func main() {
	r, nums := readData()
	fmt.Println(countVariants(r, nums))
}
