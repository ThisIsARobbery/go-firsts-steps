package main

import "fmt"

// func main() {
// 	var number, first, second, third int
// 	fmt.Scan(&number)
// 	first = number / 100
// 	second = (number / 10) % 10
// 	third = number % 10
// 	if first != second && first != third && second != third {
// 		fmt.Print("YES")
// 	} else {
// 		fmt.Print("NO")
// 	}
// }

func main() {
	var a, b, c, i int
	fmt.Scan(&i)
	a = i / 100
	b = (i % 100) / 10
	c = (i % 100) % 10
	fmt.Println(a, b, c)
	if a != b && b != c && c != a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
