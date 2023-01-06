package main

import "fmt"

func main() {
	var length, val int
	fmt.Scan(&length)
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&val)
		slice[i] = val
	}
	for i := 0; i < length; i = i + 2 {
		if i+2 >= length {
			fmt.Print(slice[i])
		} else {
			fmt.Print(slice[i], " ")
		}
	}
}
