package main

import "fmt"

func main() {
	var workArray [10]uint8
	var i, ind1, ind2 int
	for i = 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	for i = 0; i < 3; i++ {
		fmt.Scan(&ind1, &ind2)
		workArray[ind1], workArray[ind2] = workArray[ind2], workArray[ind1]
	}
	for ind, elem := range workArray {
		if ind == len(workArray)-1 {
			fmt.Print(elem)
		} else {
			fmt.Print(elem, " ")
		}
	}
}
