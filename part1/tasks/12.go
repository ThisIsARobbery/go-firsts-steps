package main

import (
	"fmt"
	"math"
)

func main() {
	var N, pow int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		pow = int(math.Pow(2, float64(i)))
		if pow >= 1 && pow <= N {
			fmt.Printf("%d ", pow)
		}
	}
}
