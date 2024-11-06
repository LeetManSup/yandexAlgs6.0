package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getMinDays(k int, tasks []int) int {
	n := len(tasks)
	sort.Ints(tasks)
	daysNeeded := 0
	l := 0

	for r := 0; r < n; r++ {
		for tasks[r]-tasks[l] > k {
			l++
		}
		overlap := r - l + 1
		if overlap > daysNeeded {
			daysNeeded = overlap
		}
	}

	return daysNeeded
}
func main() {
	k, nums := readData()
	fmt.Println(getMinDays(k, nums))
}
