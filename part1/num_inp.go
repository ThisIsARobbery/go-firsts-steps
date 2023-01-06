package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	number *= 2
	number += 100
	fmt.Print(number)
}
