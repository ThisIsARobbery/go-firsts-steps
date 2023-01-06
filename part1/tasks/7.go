package main

import "fmt"

func main() {
	var N, val int
	fmt.Scan(&N)
	var cnt int = 0

	for i := 0; i < N; i++ {
		fmt.Scan(&val)
		if val == 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
