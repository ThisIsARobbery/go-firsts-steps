package main

import "fmt"

func main() {
	var (
		length  int
		counter int = 0
		val     int
	)

	fmt.Scan(&length)

	slice := make([]int, length)

	for i := 0; i < length; i++ {
		fmt.Scan(&val)
		slice[i] = val
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] > 0 {
			counter++
		}
	}

	fmt.Print(counter)
}
