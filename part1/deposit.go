package main

import "fmt"

func main() {
	var initialDeposit, percents, goal, currentWithCents int
	fmt.Scan(&initialDeposit, &percents, &goal)
	var currentBalance int = initialDeposit * 100 // in cents
	var yearsCount int = 0
	for currentBalance < goal*100 {
		yearsCount++
		currentWithCents = currentBalance * (100 + percents) / 100
		currentBalance = currentWithCents - (currentWithCents % 100)
	}
	fmt.Print(yearsCount)
}
