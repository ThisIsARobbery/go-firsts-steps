package main

import "fmt"

func main() {
	var inputNumber float64

	fmt.Scan(&inputNumber)

	if inputNumber <= 0 {
		fmt.Printf("число %2.2f не подходит", inputNumber)
	} else if inputNumber > 10000 {
		fmt.Printf("%e", inputNumber)
	} else {
		fmt.Printf("%.4f", inputNumber*inputNumber)
	}
}
