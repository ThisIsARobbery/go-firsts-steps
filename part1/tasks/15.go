package main

import "fmt"

func main() {
	var N, rightPart, digit, digitToDelete int
	divider := 10
	fmt.Scan(&N, &digitToDelete)
	for N%(divider/10) != N {
		digit = 10 * (N % divider) / divider
		if digit == digitToDelete {
			rightPart = N % (divider / 10)
			N /= divider
			N *= divider / 10
			N += rightPart
			divider = 10
		} else {
			divider *= 10
		}
	}
	fmt.Print(N)
}
