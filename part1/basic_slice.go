package main

import "fmt"

func main() {
	var length, i, elem int
	fmt.Scan(&length)
	slice := make([]int, length, length)
	for i = 0; i < length; i++ {
		fmt.Scan(&elem)
		slice[i] = elem
	}
	fmt.Print(slice[3])
}
