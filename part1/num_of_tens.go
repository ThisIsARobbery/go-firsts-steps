package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	var result int = (number%100 - (number % 10)) / 10
	fmt.Print(result)
}
