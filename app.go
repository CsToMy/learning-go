package main

import (
	"fmt"

	"example.hu/app/bank"
	"example.hu/app/investment"
)

func main() {
	fmt.Println("Welcome to the practicing app!")
	var mainChoice int

	for {
		fmt.Println("What would you like to do?")
		printMainOptions()
		fmt.Scan(&mainChoice)
		if mainChoice < 1 || mainChoice > 3 {
			fmt.Println("Invalid choice. Choose a number between 1 and 3.")
			continue
		}

		if mainChoice == 1 {
			bank.Banking()
		} else if mainChoice == 2 {
			investment.Investment()
		} else {
			break
		}
	}
	fmt.Println("Good bye!")
}

func printMainOptions() {
	fmt.Println("----------------------------")
	fmt.Println("1. Banking app")
	fmt.Println("2. Investment app")
	fmt.Println("3. Exit")
	fmt.Println("----------------------------")
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
