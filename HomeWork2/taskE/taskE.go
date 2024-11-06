package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianHeap struct {
	low  *MaxHeap
	high *MinHeap
}

func NewMedianHeap() *MedianHeap {
	return &MedianHeap{
		low:  &MaxHeap{},
		high: &MinHeap{},
	}
}

func (mh *MedianHeap) Add(num int) {
	if mh.low.Len() == 0 || num <= (*mh.low)[0] {
		heap.Push(mh.low, num)
	} else {
		heap.Push(mh.high, num)
	}

	if mh.low.Len() > mh.high.Len()+1 {
		heap.Push(mh.high, heap.Pop(mh.low))
	} else if mh.high.Len() > mh.low.Len() {
		heap.Push(mh.low, heap.Pop(mh.high))
	}
}

func (mh *MedianHeap) RemoveMedian() int {
	var median int
	if mh.low.Len() >= mh.high.Len() {
		median = heap.Pop(mh.low).(int)
	} else {
		median = heap.Pop(mh.high).(int)
	}
	return median
}

func (mh *MedianHeap) GetMedian() int {
	if mh.low.Len() > mh.high.Len() {
		return (*mh.low)[0]
	} else {
		return (*mh.high)[0]
	}
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

func main() {
	nums := readData()
	mh := NewMedianHeap()

	for _, num := range nums {
		mh.Add(num)
	}

	for mh.low.Len() > 0 || mh.high.Len() > 0 {
		fmt.Print(mh.RemoveMedian(), " ")
	}
}
