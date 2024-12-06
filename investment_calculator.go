package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, expectedReturnRate, years := 12000.0, 5.5, 5.0

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100.0, years)
	fmt.Print(futureValue, "\n")
	outputNullValues()
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
