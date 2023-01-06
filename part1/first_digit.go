package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	if number/10000 > 0 {
		fmt.Print(number / 10000)
	} else if number/1000 > 0 {
		fmt.Print(number / 1000)
	} else if number/100 > 0 {
		fmt.Print(number / 100)
	} else if number/10 > 0 {
		fmt.Print(number / 10)
	} else {
		fmt.Print(number)
	}
}
