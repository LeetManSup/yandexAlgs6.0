package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Algorithm struct {
	idx int
	a   int
	b   int
}

func readData() (int, []Algorithm, []int) {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, _ := strconv.Atoi(nStr)

	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)
	aValues := strings.Split(aStr, " ")

	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)
	bValues := strings.Split(bStr, " ")

	pStr, _ := reader.ReadString('\n')
	pStr = strings.TrimSpace(pStr)
	pValues := strings.Split(pStr, " ")

	algorithms := make([]Algorithm, n)
	moods := make([]int, n)

	for i := 0; i < n; i++ {
		ai, _ := strconv.Atoi(aValues[i])
		bi, _ := strconv.Atoi(bValues[i])
		pi, _ := strconv.Atoi(pValues[i])
		algorithms[i] = Algorithm{
			idx: i,
			a:   ai,
			b:   bi,
		}
		moods[i] = pi
	}

	return n, algorithms, moods
}

func getAlgOrder(n int, algorithms []Algorithm, moods []int) string {
	interestingnessOrder := make([]Algorithm, n)
	usefulnessOrder := make([]Algorithm, n)
	copy(interestingnessOrder, algorithms)
	copy(usefulnessOrder, algorithms)

	sort.Slice(interestingnessOrder, func(i, j int) bool {
		if interestingnessOrder[i].a != interestingnessOrder[j].a {
			return interestingnessOrder[i].a > interestingnessOrder[j].a
		}
		if interestingnessOrder[i].b != interestingnessOrder[j].b {
			return interestingnessOrder[i].b > interestingnessOrder[j].b
		}
		return interestingnessOrder[i].idx < interestingnessOrder[j].idx
	})

	sort.Slice(usefulnessOrder, func(i, j int) bool {
		if usefulnessOrder[i].b != usefulnessOrder[j].b {
			return usefulnessOrder[i].b > usefulnessOrder[j].b
		}
		if usefulnessOrder[i].a != usefulnessOrder[j].a {
			return usefulnessOrder[i].a > usefulnessOrder[j].a
		}
		return usefulnessOrder[i].idx < usefulnessOrder[j].idx
	})

	interestingPtr := 0
	usefulPtr := 0
	studied := make([]bool, n)
	result := make([]int, n)

	for day := 0; day < n; day++ {
		if moods[day] == 0 {
			for interestingPtr < n && studied[interestingnessOrder[interestingPtr].idx] {
				interestingPtr++
			}
			if interestingPtr < n {
				idx := interestingnessOrder[interestingPtr].idx
				result[day] = idx + 1
				studied[idx] = true
				interestingPtr++
			}
		} else {
			for usefulPtr < n && studied[usefulnessOrder[usefulPtr].idx] {
				usefulPtr++
			}
			if usefulPtr < n {
				idx := usefulnessOrder[usefulPtr].idx
				result[day] = idx + 1
				studied[idx] = true
				usefulPtr++
			}
		}
	}

	output := make([]string, n)
	for i, v := range result {
		output[i] = strconv.Itoa(v)
	}
	return strings.Join(output, " ")
}

func main() {
	n, algorithms, moods := readData()
	fmt.Println(getAlgOrder(n, algorithms, moods))
}
