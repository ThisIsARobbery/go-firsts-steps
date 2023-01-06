package main

import "fmt"

func main() {
	var initNumber, currentRoot, tempSum, tempNum int

	fmt.Scan(&initNumber)
	currentRoot = initNumber
	tempSum = 0
	for currentRoot%10 != currentRoot {
		tempNum = currentRoot
		for tempNum > 0 {
			tempSum += tempNum % 10
			tempNum /= 10
		}
		currentRoot = tempSum
		tempSum = 0
	}
	fmt.Print(currentRoot)
}
