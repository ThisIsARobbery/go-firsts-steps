package main

import "fmt"

func main() {
	var A, B int

	fmt.Scan(&A, &B)
	for i := B; i >= A; i-- {
		if i%7 == 0 {
			fmt.Print(i)
			return
		}
	}
	fmt.Print("NO")
}
