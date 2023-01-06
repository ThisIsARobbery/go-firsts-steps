package main

import "fmt"

func main() {
	var N, min, val, count int

	fmt.Scan(&N)

	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&val)
		numbers[i] = val
	}
	min = numbers[0]
	count = 1
	for i := 1; i < N; i++ {
		if numbers[i] == min {
			count++
		} else if numbers[i] < min {
			min = numbers[i]
			count = 1
		}
	}
	fmt.Print(count)
}
