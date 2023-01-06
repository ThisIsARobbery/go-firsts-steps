package main

import "fmt"

func main() {
	var a, b, index int
	var sum int = 0
	fmt.Scan(&a, &b)
	for index = a; index < b+1; index++ {
		sum += index
	}
	fmt.Print(sum)
}
