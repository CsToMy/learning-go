package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit := EBT * (taxRate / 100)
	ratio := EBT / profit

	fmt.Println("EBT:", EBT)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}

func investment_calculator() {
	const inflationRate float64 = 2.5

	var investmentAmount float64
	var expectedReturnRate float64 = 2.9 // default value
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount) // passing a pointer instead of a reference

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100.0, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Print("Future value: ", futureValue, "\n")
	fmt.Print("Future real value: ", futureRealValue, "\n")
}

func outputNullValues() {
	var a int
	var b float64
	var c string
	var d bool
	var e complex64

	fmt.Print("Int: ", a, "\n")
	fmt.Print("Float: ", b, "\n")
	fmt.Print("String: ", c, "\n")
	fmt.Print("Boolean: ", d, "\n")
	fmt.Print("Complex: ", e, "\n")
}
