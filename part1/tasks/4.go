package main

import "fmt"

func main() {
	var a, b, c uint
	fmt.Scan(&a, &b, &c)
	if a*a+b*b == c*c {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}
