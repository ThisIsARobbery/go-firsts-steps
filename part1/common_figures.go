package main

import "fmt"

func main() {
	var num1,
		num2,
		num1digit,
		num2digit,
		num2cursor,
		num1digitCount,
		num1forCount,
		rank int
	var firstTyped bool = false
	rank = 1
	fmt.Scan(&num1, &num2)
	num1forCount = num1
	num1digitCount = 0
	for num1forCount > 0 {
		num1digitCount++
		num1forCount /= 10
		rank *= 10
	}
	rank /= 10
	// O(n^2), ah shit
	num2cursor = num2
	for num1 > 0 {
		num1digit = num1 / rank // first digit
		for num2cursor > 0 {
			num2digit = num2cursor % 10
			if num1digit == num2digit {
				if firstTyped {
					fmt.Print(" ", num1digit)
				} else {
					fmt.Print(num1digit)
					firstTyped = true
				}
			}
			num2cursor /= 10
		}
		num2cursor = num2  // renew cursor
		num1 = num1 % rank // removing first digit
		rank /= 10
	}
}
