package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Print(n, " ")

	if n%10 == 0 {
		fmt.Print("korov")
	} else if n >= 11 && n <= 19 {
		fmt.Print("korov")
	} else if n%10 == 1 {
		fmt.Print("korova")
	} else if n%10 >= 2 && n%10 <= 4 {
		fmt.Print("korovy")
	} else {
		fmt.Print("korov")
	}
}
