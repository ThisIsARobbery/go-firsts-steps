package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if number/100000+(number/10000%10)+(number/1000%10) == (number/100%10)+(number/10%10)+number%10 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
