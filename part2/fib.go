package main

import "fmt"

func fibb(n int) int {
	var (
		prev = 1
		cur  = 1
	)
	if n == 1 || n == 2 {
		return 1
	}
	for i := 3; i < n+1; i++ {
		cur, prev = cur+prev, cur
	}
	return cur
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(fibb(N))
}
