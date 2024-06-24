package main

import (
	"fmt"
	// "math"
)

func main() {

	revenue := outputUser("Revenue: ")
	expenses := outputUser("Expenses: ")
	taxRate := outputUser("taxRate: ")

	netIncome, ebt := ebtCalc(revenue, expenses, taxRate)
	fmt.Println("Expected before tax: $" + fmt.Sprint(ebt))

	fmt.Println("Expected after tax: $" + fmt.Sprint(netIncome))

}

func outputUser(infotext string) float64 {
	var userInput float64
	fmt.Print(infotext)
	fmt.Scan(&userInput)
	return userInput
}

func ebtCalc(revenue, expenses, taxRate float64) (float64, float64) {
	ebt := revenue - expenses
	netIncome := ebt * (1 - taxRate/100.0)
	return netIncome, ebt
}
