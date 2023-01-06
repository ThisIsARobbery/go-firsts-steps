package main

import "fmt"

func main() {
	var N, prev, cur int

	fmt.Scan(&N)

	var count = 2

	prev = 1
	cur = 1

	if N == 1 {
		fmt.Print(1)
		return
	}

	for cur != N {
		if cur > N {
			fmt.Print(-1)
			return
		}
		prev, cur = cur, prev+cur
		count++
	}
	fmt.Print(count)
}
