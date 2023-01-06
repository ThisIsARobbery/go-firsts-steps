package main

import "fmt"

func main() {
	var degree, hours, minutes int
	fmt.Scan(&degree)
	hours = degree / 30
	minutes = (degree - hours*30) * 2
	fmt.Print("It is ", hours, " hours ", minutes, " minutes.")
}
