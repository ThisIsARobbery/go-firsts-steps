package main

import "fmt"

func main() {
	var cur int

	for fmt.Scan(&cur); cur < 101; fmt.Scan(&cur) {
		if cur < 10 {
			continue
		}
		fmt.Println(cur)
	}
}
