package main

import "fmt"

func whereToSwim(x1, y1, x2, y2, x, y int8) string {
	var result string

	if y < y1 {
		result += "S"
	} else if y > y2 {
		result += "N"
	}

	if x < x1 {
		result += "W"
	} else if x > x2 {
		result += "E"
	}

	return result
}

func main() {
	var x1, y1, x2, y2, x, y int8
	fmt.Scan(&x1, &y1, &x2, &y2, &x, &y)
	fmt.Println(whereToSwim(x1, y1, x2, y2, x, y))
}
