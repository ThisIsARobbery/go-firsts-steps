package main

import "fmt"

func main() {
	var n, i, cur int
	var sum = 0
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&cur)
		if cur/100 == 0 && (cur/10 != 0) && cur%8 == 0 {
			sum += cur
		}
	}
	fmt.Print(sum)
}
