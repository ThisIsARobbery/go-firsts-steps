package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	switch {
	case number > 0:
		fmt.Print("Число положительное")
	case number < 0:
		fmt.Print("Число отрицательное")
	case number == 0:
		fmt.Print("Ноль")
	}
}
