package main

import "fmt"

func minimumFromFour() int {
	var min, val int
	fmt.Scan(&val)
	min = val
	for i := 0; i < 3; i++ {
		fmt.Scan(&val)
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	min := minimumFromFour()
	fmt.Print(min)
}
