package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	var max int = array[0]
	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	fmt.Print(max)
}
