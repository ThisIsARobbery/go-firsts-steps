package main

import "fmt"

func main() {
	var (
		current int
		max     int
		count   int = 1
	)
	fmt.Scan(&current)
	max = current
	for fmt.Scan(&current); current != 0; fmt.Scan(&current) {
		if current > max {
			max = current
			count = 1
		} else if current == max {
			count++
		}
	}
	fmt.Print(count)
}
