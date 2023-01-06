package main

import "fmt"

func main() {
	var n, c, d, cur int
	fmt.Scan(&n, &c, &d)
	for cur = 1; cur < n+1; cur++ {
		if cur%c == 0 && cur%d != 0 {
			fmt.Print(cur)
			break
		}
	}
}
