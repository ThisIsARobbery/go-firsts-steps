package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)
	fmt.Printf("%d%d%d", num%10, (num%100)/10, num/100)
}
