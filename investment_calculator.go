package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 12000 // explicitly define type here
	var expectedReturnRate = 5.5         // infered type here
	var years float64 = 5

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
