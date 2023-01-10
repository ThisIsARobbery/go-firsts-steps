package main

import "fmt"

func sumInt(args ...int) (int, int) {
	count := 0
	sum := 0
	for _, elem := range args {
		sum += elem
		count++
	}
	return count, sum
}

func main() {
	fmt.Print(sumInt(1, 2, 3, 4))
}
